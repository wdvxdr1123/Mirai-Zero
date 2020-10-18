package zero

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"reflect"
)

func initZeroEvents() {
	zero.Client.OnGroupMessage(zeroOnGroupMessage)
}

func zeroOnGroupMessage(_ *client.QQClient,m *message.GroupMessage)  {
	zero.Events.Emit(GroupMessageEvent, zero, NewBaseSession(
		Group&Message&Base,
		m.Time,
		convertStructRaw(m),
		))
	fmt.Println(convertStructRaw(m))
}

func convertStructRaw(s interface{}) MSG {
	var (
		t = reflect.TypeOf(s).Elem()
		v = reflect.ValueOf(s).Elem()
		msg = MSG{}
	)
	for i := 0; i < t.NumField(); i++ {
		jsonStr := t.Field(i).Tag.Get("json")
		if jsonStr == "" {
			continue
		}
		var val = v.Interface()
		if reflect.TypeOf(val).Kind() == reflect.Struct {
			msg[jsonStr] = convertStructRaw(val)
		} else {
			msg[jsonStr] = val
		}
	}
	return msg
}

