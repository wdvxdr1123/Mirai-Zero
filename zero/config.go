package zero

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	log "github.com/sirupsen/logrus"
	"github.com/wdvxdr1123/mirai-zero/utils"
	"io/ioutil"
	"os"
)

type Config struct {
	QQ       int64  `json:"qq"`
	Password string `json:"password"`
}

func LoadConfig() *Config {
	if utils.FileExist("config.json") {
		config := &Config{}
		cfg, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatal("无法读取配置文件 ", err)
		}
		err = json.Unmarshal(cfg, config)
		if err != nil {
			log.Fatal("读取配置文件失败 ", err)
		}
		return config
	}
	return nil
}

// 生成随机设备信息
func LoadRandomDevice() ([]byte, error) {
	if !utils.FileExist("device.json") {
		log.Warn("device.json 不存在, 将随机生成设备信息...")
		client.GenRandomDevice()
		err := ioutil.WriteFile("device.json", client.SystemDeviceInfo.ToJson(), os.FileMode(0755))
		if err != nil {
			log.Fatal("无法写入设备信息 device.json: ", err)
		}
	}
	return ioutil.ReadFile("device.json")
}
