package zero

import (
	"bytes"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	message2 "github.com/Mrs4s/MiraiGo/message"
	log "github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"github.com/wdvxdr1123/mirai-zero/message"
	"github.com/wdvxdr1123/mirai-zero/utils"
	"github.com/yinghau76/go-ascii-art"
	"image"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
)

type Zero struct {
	Config *Config
	Client *client.QQClient

	Events EventEmitter
}

var zero *Zero // 主体服务实例 <目前没有支持多号登录的计划>

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[Zero][%time%] [%lvl%]: %msg% \n",
	})
	zero = &Zero{}
}

func Init() {
	zero.Config = LoadConfig()
	zero.Client = client.NewClient(zero.Config.Uin, zero.Config.Password)
	device, err := LoadRandomDevice()
	if err != nil {
		log.Fatal("加载设备信息失败 ", err)
	}
	_ = client.SystemDeviceInfo.ReadJson(device)
}

func Start() {
	cli := zero.Client
	zero.Client.AllowSlider = false
	rsp, err := cli.Login()
	for {
		if err != nil {
			log.Fatal("登录失败: ", err)
		}
		var text string
		if !rsp.Success {
			switch rsp.Error {
			case client.NeedCaptcha:
				_ = ioutil.WriteFile("captcha.jpg", rsp.CaptchaImage, 0644)
				img, _, _ := image.Decode(bytes.NewReader(rsp.CaptchaImage))
				fmt.Println(asciiart.New("image", img).Art)
				text, _ = utils.ReadLine("请输入验证码 (captcha.jpg)： (Enter 提交)")
				rsp, err = cli.SubmitCaptcha(strings.ReplaceAll(text, "\n", ""), rsp.CaptchaSign)
				continue
			case client.SMSNeededError:
				_, _ = utils.ReadLine(fmt.Sprintf("账号已开启设备锁, 按下 Enter 向手机 %v 发送短信验证码.", rsp.SMSPhone))
				if !cli.RequestSMS() {
					log.Warnf("发送验证码失败，可能是请求过于频繁.")
					time.Sleep(time.Second * 5)
					os.Exit(0)
				}
				text, _ = utils.ReadLine("请输入短信验证码： (Enter 提交)")
				rsp, err = cli.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
				continue
			case client.SMSOrVerifyNeededError:
				log.Warnf("账号已开启设备锁，请选择验证方式:")
				log.Warnf("1. 向手机 %v 发送短信验证码", rsp.SMSPhone)
				log.Warnf("2. 使用手机QQ扫码验证.")
				text, _ = utils.ReadLine("请输入(1 - 2):")
				if strings.Contains(text, "1") {
					if !cli.RequestSMS() {
						log.Warnf("发送验证码失败，可能是请求过于频繁.")
						time.Sleep(time.Second * 5)
						os.Exit(0)
					}
					text, _ = utils.ReadLine("请输入短信验证码： (Enter 提交)")
					rsp, err = cli.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
					continue
				}
				log.Warnf("请前往 -> %v <- 验证并重启Bot.", rsp.VerifyUrl)
				_, _ = utils.ReadLine("按 Enter 继续....")
				os.Exit(0)
				return
			case client.UnsafeDeviceError:
				log.Warnf("账号已开启设备锁，请前往 -> %v <- 验证并重启Bot.", rsp.VerifyUrl)
				_, _ = utils.ReadLine("按 Enter 继续....")
				cli.Disconnect()
				continue
			case client.OtherLoginError, client.UnknownLoginError:
				log.Warnf("登录失败: %v", rsp.ErrorMessage)
				_, _ = utils.ReadLine("按 Enter 继续....")
				os.Exit(0)
				return
			}
		}
		break
	}
	log.Info("登录成功！")
	initZeroEvents()
}

// 将插件注册到主服务
func RegisterPlugin(plugin IPlugin) {
	// todo:Log
	panic("impl me")
	// plugin.Init(new(Accessory))
}

// register event
func (z *Zero) registerEvent(name uint8, f interface{}) {
	z.Events.On(EventName(name), func(data ...interface{}) {
		defer func() {

		}()
		values := make([]reflect.Value, 0, len(data))
		for _, v := range data {
			values = append(values, reflect.ValueOf(v))
		}
		_ = reflect.ValueOf(f).Call(values)
	})
}

func (z *Zero) SendGroupMessage(groupId int64, m message.IMessage) int32 {
	switch e := m.(type) {
	case *message.RichMessage:
		_ = z.Client.SendGroupMessage(groupId, &message2.SendingMessage{
			Elements: z.GroupUpload(e, groupId).Elems,
		},
		)
	}
	return 0
}

func (z *Zero) GroupUpload(m *message.RichMessage, groupId int64) *message.RichMessage {
	for index, elem := range m.Elems {
		if i, ok := elem.(*message2.ImageElement); ok {
			gm, err := z.Client.UploadGroupImage(groupId, i.Data)
			if err != nil {
				log.Warnf("警告: 群 %v 消息图片上传失败: %v", groupId, err)
				continue
			}
			m.Elems[index] = gm
			continue
		}
	}
	return m
}
