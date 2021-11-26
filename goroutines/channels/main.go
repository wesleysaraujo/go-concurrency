package main

import (
	"fmt"
	"sync"
)
// Nesse exemplo eu tenho 2 rotinas rodando em paralelo através do GoRoutines
func main() {
	channel := make(chan int)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// Nessa GoRoutine eu estou atribuindo o valor de i ao meu canal
	// Um Channel só pode rodar dentro de uma Go Routine, caso contrário o Go retorna Deadlock
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		// Esse wait quer dizer que eu aguardo enquanto as operações WaitGroup não são concluídas
		waitGroup.Wait()
		// Fecho o canal
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}
