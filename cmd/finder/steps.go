package main

import (
	t "banklistreader/internal/tools"
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

func step_ask_input_files() []string {
	var input string

	err := huh.NewInput().
		Title("Onde esta o arquivo ou pasta com os logins?").
		Placeholder("Dica: Arraste o arquivo ou a pasta para esta janela!").
		Validate(func(s string) error {
			if !t.Exists(strings.Trim(s, "\"")) {
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

	// On Windows when files are dragged to the terminal window and they contain space in their name, they come as an absolute path surrounded by double quotes. eg: "C://myfiles/folder/file.txt"
	input = strings.Trim(input, "\"")

	if !t.Exists(input) {
		color.Red("● O arquivo/pasta (%s) não existe!", input)
		return step_ask_input_files()
	}

	if t.IsDir(input) {
		text_files := t.Files(input, ".txt")

		if len(text_files) == 0 {
			color.Red("● Nenhum arquivo .txt encontrado na pasta fornecida!", input)
			return step_ask_input_files()
		}

		color.Yellow("● %d arquivos .txt foram encontrados na pasta!", len(text_files))
		return text_files
	}

	return []string{input}
}

func step_ask_filter_key() string {
	var input string

	err := huh.NewInput().
		Title("Qual o termo da pesquisa?").
		Description("Nota: prefira letras minúscula").
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

	return input
}

var FORMATS = map[string]string{
	"<Site>:<Usuário>:<Senha>": "%s:%s:%s",
	"<Usuário>:<Senha>":        "%[2]s:%[3]s",
}

func step_ask_out_format() string {
	var input string

	err := huh.NewSelect[string]().
		Options(huh.NewOptions(t.Keys(FORMATS)...)...).
		Title("Com qual formato o resultado deve ser salvo?").
		Value(&input).
		Run()

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_out_format()
	}

	return FORMATS[input]
}
