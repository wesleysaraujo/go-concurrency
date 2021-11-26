package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var m sync.Mutex

func main() {
	go runProcess("P1", 30)
	go runProcess("P2", 20)

	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	/*
		Race Condition é um comportamento que ocorre quando utilizo GoRoutine e tenho um mesmo processo rodando em concorrência/paralelo
		Nesse exemplo abaixo, eu estou atribuindo a execução de cada iteração do Loop a variável resul
		O problema é que ao executar P1->1 e P2->1, o valor do meu result será atribuído como 1 duas vezes,
		No nosso exemplo eu tenho 50 processos executados, porém o número máximo do meu result será 30(P1)

		Obs: Para verificar se meu código está com race condition, eu rodo:
		go run -race race_condition.go

		Para resolver o problema de racecondition usamos sync.Mutex, onde temos os métodos Lock() e Unlock()
	 */
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", i, "Partial result: ", result)
		m.Unlock()
	}
}

