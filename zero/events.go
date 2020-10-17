package zero

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/wdvxdr1123/mirai-zero/events"
	"github.com/wdvxdr1123/mirai-zero/types/session"
)

func initZeroEvents() {
	zero.Client.OnGroupMessage(zeroOnGroupMessage)
}

func zeroOnGroupMessage(c *client.QQClient,m *message.GroupMessage)  {
	zero.Events.Emit(events.ZeroEventGroupMessage, zero, session.NewBaseSession(
		session.Group & session.Message & session.Base,
		m.Time,
		))
}
