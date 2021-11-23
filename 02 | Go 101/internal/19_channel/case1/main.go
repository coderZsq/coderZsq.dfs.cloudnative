package main

func main() {
	close(ch)

	ch <- v

	<-ch

	v = <-ch
	v, sentBeforeClosed = <-ch

	cap(ch)

	len(ch)
}
