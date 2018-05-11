package main

import "fmt"

func hanoi(height int, from string, buffer string, to string) {
	if height == 0 {
		return
	}
	hanoi(height-1, from, to, buffer)
	fmt.Println(from, "=>", to)
	hanoi(height-1, buffer, from, to)
}
func main() {
	hanoi(3, "left", "middle", "right")

}
