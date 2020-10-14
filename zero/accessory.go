package zero

type Accessory struct {
	zero *Zero
}

// Get the mirai-zero object
func (a *Accessory) GetZero() *Zero {
	return a.zero
}