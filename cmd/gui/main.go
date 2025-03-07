package main

import (
	"fmt"

	z "github.com/ncruces/zenity"
)

func main() {
	folder, _ :=z.SelectFile(z.Title("Selecione a pasta contendo os logins"),z.Directory())

	outFormat, _ := z.ListItems("Escolha como os logins devem serem salvos", "usuario:senha", "site:usuario:senha")

	fmt.Println(folder, outFormat)
}