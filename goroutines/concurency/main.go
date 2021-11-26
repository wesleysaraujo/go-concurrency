package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// WaitGroup nos permite controlar de uma forma mais eficientes processos que estão rodando em concorrência.
var waitGroup sync.WaitGroup

func init() {
	// Define quantos CPUs eu quero usar para minha aplicação
	// runtime.GOMAXPROCS(2)
	runtime.GOMAXPROCS(runtime.NumCPU()) // Nesse caso eu uso todas as minhas CPUs
}

func main() {
	waitGroup.Add(2)
	go runProcess("P1", 30)
	go runProcess("P2", 20)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}
