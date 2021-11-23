package main

func main() {
	for v := range aChannel {
		// 使用v
	}

	for {
		v, ok = <-aChannel
		if !ok {
			break
		}
		// 使用v
	}
}
