package main

import (
	"fmt"
	"time"
)

//function
func plus(a int, b int) int {
	return a + b
}

func calc(a int, b int) (int, int) {
	return a / b, a * b
}

type node struct {
	mapping map[int]int
	index   int
}

func main() {
	fmt.Println("Hello, World!")
	/*var a = "initial"
	var b, c int =1,2
	var d = true
	var e int
	e = 3
	f:="apple"  //automate get the type
	const s string = "constant"*/

	//loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for { //infinite loop
		fmt.Println("loop")
		break
	}

	//if else
	if 7%2 == 0 {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}

	//array
	var g [5]int
	g[4] = 100

	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}

	//Hash map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	//function
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	div, mul := calc(3, 7)
	fmt.Println("3/7 =", div)
	fmt.Println("3*7 =", mul)
	//goroutine have two thread to run tgt one is the function and one is the main
	go func(msg string) { // goroutine
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	//channel

	messages := make(chan string)      // chan is the channel the element in the channel is the string
	go func() { messages <- "ping" }() // put ping into the channel named messages

	msg := <-messages //wait to get the ping from the channel named messages
	fmt.Println(msg)

	//select
	msg1 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		msg1 <- "ping1"
	}()
	msg2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		msg1 <- "ping2"
	}()

	for i := 0; i < 2; i++ {
		select { //wait for the channel who will reply sin
		case recv1 := <-msg1: //msg1 receive sin do
			fmt.Println("received", recv1)
		case recv2 := <-msg2: //msg2 receive sin do
			fmt.Println("received", recv2)
		}
	}

	rn := node{
		mapping: make(map[int]int),
		index:   0,
	}
	rn.mapping[15] = 5 //it will auto create a key vlaue pair if there doesn't exist

	fmt.Println(rn.mapping[15])
	fmt.Println(rn.mapping[1]) //show 0 for the first value
	if rn.mapping[2] == 0 {
		fmt.Println("success in comparing an uninitialize thing")
	}
	fmt.Println(rn.index)

	slice := []string{"A", "B", "C"}
	sl := slice[1:3]
	fmt.Println(sl, slice[2])
}
