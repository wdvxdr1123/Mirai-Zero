package utils

import (
	"github.com/manifoldco/promptui"
)

func Check(i interface{}, err error) interface{} {
	if err != nil {
		return nil
	}
	return i
}

func ReadLine(info string) (string, error) {
	prompt := &promptui.Prompt{Label: info}
	return prompt.Run()
}
