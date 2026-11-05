package main

import (
	"fmt"
	"time"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {

	go temp2()
	time.Sleep(time.Second * 12)
	fmt.Println("DONE")

}

func temp2() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func temp1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "  ")
		time.Sleep(time.Second * 3)
	}
}
