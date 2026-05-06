package main

import (
	"fmt"
	"log"
	"vitozap/greetings"
)

func main() {

	//Propriedades básicas do Logger imbutido da linguagem
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("vitozap")

	//Se o campo erro for diferente de nil, ele escreve e finaliza o programa
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
