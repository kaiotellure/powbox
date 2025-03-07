package main

import (
	t "banklistreader/internal/tools"

	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func step_ask_input_files() []string {
	fmt.Printf(
		"%s Arraste o arquivo %s ou uma pasta com os logins para esta janela!\n",
		color.YellowString("→"),
		color.MagentaString(".txt"),
	)

	// On Windows when files are dragged to the terminal window and they contain space in their name, they come as an absolute path surrounded by double quotes. eg: "C://myfiles/folder/file.txt"
	input := strings.Trim(t.Ask(), "\"")

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
	fmt.Printf(
		"%s Qual termo quer achar? exemplos: %s, %s ou %s\n",
		color.YellowString("→"),
		color.MagentaString("max"),
		color.RedString("netflix"),
		color.BlueString("paramount"),
	)

	input := t.Ask()

	if input == EMPTY_STRING {
		color.Red("● O termo não pode ser vazio!")
		return step_ask_filter_key()
	}

	if len(input) < 3 {
		color.Red("● O termo não pode ser tão curto!")
		return step_ask_filter_key()
	}

	return input
}

var FORMATS = map[string]string{
	"site:usuário:senha": "%s:%s:%s",
	"usuário:senha":      "%[2]s:%[3]s",
}

func step_ask_out_format() string {
	format_select := promptui.Select{
		Label:    "Como os logins encontrados devem ser guardados?",
		Items:    t.Keys(FORMATS),
		HideHelp: true,
	}

	_, format, err := format_select.Run()
	if err != nil {
		color.Red("● Error ao ler formato: %v", err)
		return step_ask_out_format()
	}

	return FORMATS[format]
}
