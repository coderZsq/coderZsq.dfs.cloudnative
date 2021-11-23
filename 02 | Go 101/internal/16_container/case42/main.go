package main

type Buffer struct {
	start, end int
	data       [1024]byte
}

func fa(buffers []Buffer) int {
	numUnreads := 0
	for _, buf := range buffers {
		numUnreads += buf.end - buf.start
	}
	return numUnreads
}

func fb(buffers []Buffer) int {
	numUnreads := 0
	for i := range buffers {
		numUnreads += buffers[i].end - buffers[i].start
	}
	return numUnreads
}

func main() {

}
