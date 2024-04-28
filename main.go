package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Bro %v + %v is %v \n", v, v, v+v)
	case string:
		fmt.Printf("Bro \"%v\" is %v bytes long\n", v, len(v))
	default:
		fmt.Println("I don't know you bro")
	}
}

func main() {
	do(10)
	do("Bugatti")
	do(false)

}
