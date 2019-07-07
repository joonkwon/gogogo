package main

import "fmt"

//https://golang.org/ref/spec#Numeric_types
//https://golang.org/pkg/fmt/

// byte		   alias for uint8
// rune        alias for int32

func main() {

	//#########################################################

	fmt.Println("bytes: ", []byte("café"))
	fmt.Println("runes: ", []rune("café"))

	//#########################################################

	//print all the Unicode Characters in string from 0-999
	// for i := 0; i < 999; i++ {
	// 	fmt.Println(i, "-", string(i))
	// }

	//#########################################################

	//For strings, the range loop iterates over Unicode code points
	//The index is the first byte of a UTF-8-encoded code point
	//sentence := "Hello"
	//sentence := ("你好")

	//Byte
	//sentence := []byte("Hello")
	//sentence := []byte("你好")

	//Rune
	//sentence := []rune("Hello")
	//sentence := []rune("你好")
	// sentence := "你好"

	// counter := 0
	// for i, letter := range sentence {
	// 	fmt.Printf(" index %d Character: %c\n", i, letter)
	// 	//fmt.Printf("index %d string: %s\n", i, letter)
	// 	counter++
	// }

	// fmt.Println(sentence)
	// fmt.Printf("counter value: %v\n", counter)

	//#########################################################
	//Conversions

	//Convert String to rune
	// r := []rune("ABC€")
	// fmt.Println(r)
	// fmt.Printf("%U\n", r)

	//Convert Rune to String
	// s := string([]rune{'\u0041', 'B', '\u0043', '\u20AC'})
	// fmt.Println(s)

}

//Example adding space: https://play.golang.org/p/gDejvQbEUL
