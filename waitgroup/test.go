package main

import (
	"fmt"
	"time"
)

func funcaoDemorada(i int, c chan bool) (flag bool) {
	time.Sleep(1000000000)
	fmt.Printf("Iniciando processo %d \n", i)
	time.Sleep(1000000000)
	fmt.Printf("Retornando processo %d \n", i)

	if i%2 == 0 {
		c <- true
	} else {
		c <- false
	}

	return
}

func main() {
	fmt.Println("Programa inicializado")

	channel := make(chan bool)
	for i := 0; i < 5; i++ {
		go funcaoDemorada(i, channel)
		if <-channel {

		} else {

		}
	}

}
