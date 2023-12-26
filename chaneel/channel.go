package main

import (
	"fmt"
	"time"
)

func sleepTropper(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("las tropas, el numero %d, esta roncando\n", id)
	c <- id
}
func main() {
	c := make(chan int)

	for i := 0; i < 6; i++ {
		go sleepTropper(i, c)

	}
	for i := 0; i < 6; i++ {
		stormTroper := <-c
		fmt.Printf("las stormtrooper, numero %d esta derpierto!\n", stormTroper)
	}
	close(c)

}
