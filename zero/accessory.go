package zero

import (
	"context"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/wdvxdr1123/mirai-zero/events"
	"reflect"
)

type Context struct {
	context.Context

	data map[string]interface{}
}

// 各种事件，方法的辅助结构
type Accessory struct {
	client  *client.QQClient
	zero    *Zero
	context Context
}

// Get the mirai-zero
func (a *Accessory) GetZero() *Zero {
	return a.zero
}

// Get the mirai-cliet
func (a *Accessory) GetClient() *client.QQClient {
	return a.client
}

// Get the global config
func (a *Accessory) GetConfig() *Config {
	return a.zero.config
}

// register event
func (a *Accessory) registerEvent(name string, f interface{}) {
	a.zero.events.On(events.EventName(name), func(data ...interface{}) {
		defer func() {

		}()
		values := make([]reflect.Value, 0, len(data))
		for _, v := range data {
			values = append(values, reflect.ValueOf(v))
		}
		_ = reflect.ValueOf(f).Call(values)
	})
}
