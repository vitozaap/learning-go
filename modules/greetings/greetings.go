package greetings

import (
	"errors"
	"fmt"
)

// Função retorna string e erro, como possui dois returns, preciso setar em cada fluxo o estado de cada return
func Hello(name string) (string, error) {
	if name == "" {
		//Retorna "" em string e um Erro
		return "", errors.New("Name is empty")
	}

	//Usa a função len() que verificar o length do parâmetro, se for menor que 3, retorna um erro
	if len(name) < 3 {
		return "", errors.New("No way your name is this short, cmon!! The minimal is: 3")
	}

	message := fmt.Sprintf("Hello, %v! Nice to meet You!", name)

	//Retorna a mensagem, e um valor nil (significa literalmente: "sem erros") para o campo de return "error" da função
	return message, nil
}
