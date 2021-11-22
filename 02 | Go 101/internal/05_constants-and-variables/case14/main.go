package main

const (
	Failed = iota - 1 // == -1
	Unknown           // == 0
	Succeeded         // == 1
)

const (
	Readable = 1 << iota // == 1
	Writable             // == 2
	Executable           // == 4
)

func main() {

}
