package zero

import "github.com/Mrs4s/MiraiGo/client"

type Accessory struct {
	client *client.QQClient
	zero *Zero
}

// Get the mirai-zero object
func (a *Accessory) GetZero() *Zero {
	return a.zero
}

// Get the mirai cliet object
func (a *Accessory) GetClient() *client.QQClient {
	return a.client
}