package pipeline

// Concurrency in Go

// or done channel
// 合并一系列Done Channel,只要有一个完成就退出
func Or(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 递归选取
		switch len(chans) {
		case 2:
			select {
			case <-chans[0]:
			case <-chans[1]:
			}
		default:
			select {
			case <-chans[0]:
			case <-chans[1]:
			case <-chans[2]:
			case <-Or(chans[3:]...):
			}
		}
	}()
	return orDone
}

// or-done-channel
func OrDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

// 将离散值生成pipeline
func Generator(done <-chan interface{}, its ...interface{}) <-chan interface{} {
	itStream := make(chan interface{})
	go func() {
		defer close(itStream)
		for _, it := range its {
			select {
			case <-done:
				return
			case itStream <- it:
			}
		}
	}()
	return itStream
}

// 将离散值无限生成pipeline
func Repeat(done <-chan interface{}, its ...interface{}) <-chan interface{} {
	itStream := make(chan interface{})
	go func() {
		defer close(itStream)
		for {
			for _, it := range its {
				select {
				case <-done:
					return
				case itStream <- it:
				}
			}
		}
	}()
	return itStream
}

// 从pipeline中挑选前 num 个
func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- valueStream:
			}
		}
	}()
	return takeStream
}

// 桥接 Channel
func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range OrDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
