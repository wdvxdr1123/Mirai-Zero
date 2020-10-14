package message

func NewGiftMessage(ProductId uint32) *GiftMessage {
	return &GiftMessage{productId: ProductId}
}

func (g *GiftMessage) Send() {
	panic("impl me")
}
