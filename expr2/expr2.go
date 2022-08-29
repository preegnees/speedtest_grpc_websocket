package main

func main() {
	for {
		select {
		default:
			println("helllo")
			return
		}
	}
}