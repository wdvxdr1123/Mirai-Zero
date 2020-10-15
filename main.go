package main

import "github.com/wdvxdr1123/mirai-zero/zero"

func main() {
	zero.Init()

	zero.Start()

	select {}
}
