package message

type GiftMessage struct {
	productId uint32
}

func NewGiftMessage(ProductId uint32) *GiftMessage {
	return &GiftMessage{productId: ProductId}
}

func (g *GiftMessage) Send() {
	panic("impl me")
}
