package main

import (
	t "powbox/internal/tools"

	"fmt"
	"strconv"

	"errors"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

func step_ask_input_files() []string {
	var input string

	err := huh.NewInput().
		Placeholder("Dica: Arraste o arquivo ou a pasta para esta janela!").
		Title("Onde está o arquivo ou pasta com os logins?").
		Validate(func(s string) error {
			if !t.Exists(strings.Trim(s, "\"'")) {
				return errors.New("arquivo ou pasta não existe")
			}
			return nil
		}).
		Prompt("→ ").
		Value(&input).
		Run()

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_input_files()
	}

	// Remove any trailing quotes, (") Windows (') Linux
	input = strings.Trim(input, "\"'")

	if t.IsDir(input) {
		// Gathers all .txt files in that directory, non-recursively
		text_files := t.Files(input, ".txt")

		if len(text_files) == 0 {
			color.Red("● Nenhum arquivo .txt encontrado na pasta fornecida!", input)
			return step_ask_input_files()
		}

		color.Yellow("● %d arquivos .txt foram encontrados na pasta!", len(text_files))

		fmt.Printf("%s a pasta escolhida foi: %s(%s)\n", purple("●"), purple(input), pink(strconv.Itoa(len(text_files))))
		return text_files
	}

	fmt.Printf("%s o arquivo escolhido foi: %s\n", purple("●"), purple(input))

	// User provided a single input file
	return []string{input}
}

func step_ask_filter_key() string {
	var input string

	err := huh.NewInput().
		Title("Qual o termo da pesquisa?").
		Description(warn("NOTA") + " Letras maiúsculas e minúsculas não são iguais!").
		Placeholder("max.com, netflix.com, adobe.com e etc...").
		Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("não pode ser vazio")
			}
			return nil
		}).
		Prompt("→ ").
		Value(&input).
		Run()

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_filter_key()
	}

	fmt.Printf("%s o termo de pesquisa foi: %s\n", purple("●"), purple(input))
	return input
}

var FORMATS = map[string]string{
	"<site>:<usuário>:<senha>": "%s:%s:%s",
	"<usuário>:<senha>":        "%[2]s:%[3]s",
}

func step_ask_out_format() string {
	var input string

	err := huh.NewSelect[string]().
		Title("Com qual formato o resultado deve ser salvo?").
		Options(huh.NewOptions(t.Keys(FORMATS)...)...).
		Value(&input).
		Run()

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_out_format()
	}

	fmt.Printf("%s formato das credenciais: %s\n", purple("●"), purple(input))
	return FORMATS[input]
}
