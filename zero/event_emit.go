package zero

func initZeroEvents() {
	//zero.Client.OnGroupMessage(zeroOnGroupMessage)
}

/*
func zeroOnGroupMessage(_ *client.QQClient,m *message.GroupMessage)  {
	zero.Events.Emit(GroupMessageEvent, zero, NewBaseSession(
		Group|Message|Base,
		m.Time,
		convertStructRaw(m),
		))
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
		var val = v.Field(i).Interface()
		msg[jsonStr] = val
	}
	return msg
}
*/
