package main

func main() {
	for key := range m {
		delete(m, key)
	}
}
