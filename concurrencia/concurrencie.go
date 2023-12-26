package main

import (
	"fmt"
	"time"
)

func tarea(laGoroutine int, canal chan int) {
	time.Sleep(time.Second * time.Duration(laGoroutine))
	canal <- laGoroutine
}

func main() {
	canal := make(chan int)

	for i := 1; i <= 5; i++ {
		go tarea(i, canal) // Inicia goroutines para ejecutar la tarea
	}

	for i := 1; i <= 5; i++ {
		valor := <-canal // Recibe valores del canal
		fmt.Println("Goroutine", valor, "terminada.")
	}
}
