package main

import "fmt"

func main() {
	name := "Jacques"
	age := 36
	version := 1.1
	fmt.Println("Hello, Sr.", name, "your age is", age)
	fmt.Println("This program is on version", version)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comand int
	fmt.Scan(&comand)

	fmt.Println("O endereço da minha variável é", &comand)
	fmt.Println("O comando escolhido foi", comand)

}
