package structures

import (
	"errors"
	"fmt"
	"log"
)

type Person struct {
	name      string
	lastName  string
	birthDate string
	cpf       string
	renamed   bool
}

var me = Person{
	name:      "Victor",
	lastName:  "Santos",
	birthDate: "23092006",
	cpf:       "000.000.000-00",
	renamed:   false,
}

func GetDefault() Person {
	fmt.Println("||||||||||||||| STRUCTURES |||||||||||||||")
	fmt.Printf("Nome inicial: %s \n", me.name)
	log.SetFlags(0)
	log.SetPrefix("Structures error: ")
	renomear(&me, "Novo nome 1")
	fmt.Printf("Renomeado: %s \n", me.name)
	//Como passa o endereço, ele renomeia DIRETO, sem ser uma cópia
	isRenamed, err  := renomear(&me, "Novo nome 2")
	if !isRenamed {
		log.Fatal(err)
	}
	
	//Endereço do objeto "me"
	fmt.Printf("%p \n", &me)
	return me
}

// Com ponteiro: recebe o ENDEREÇO, original muda
func renomear(p *Person, name string) (bool, error) {

	//Se ja foi renomeado, nao pode renomear denovo, então lança um erro
	if me.renamed {
		return false, errors.New("Struct has been renamed before.")
	}

	//Renomeia a variável em si, nao somente uma cópia
	p.name = name
	p.renamed = true

	return true, nil
}
