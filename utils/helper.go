package utils

func Check(i interface{}, err error) interface{} {
	if err != nil {
		return nil
	}
	return i
}
