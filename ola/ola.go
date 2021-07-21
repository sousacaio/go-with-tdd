package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá,"
const prefixoOlaEspanhol = "Hola,"
const prefixoOlaFrances = "Salut,"

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixodeSaudacao(idioma) + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
