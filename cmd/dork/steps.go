package main

import (
	"banklistreader/internal/tools"
	"errors"
	"io"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

func step_ask_dork_file() []string {
	var input string

	err := huh.NewInput().
		Title("Onde esta o arquivo .txt com as dorks?").
		Placeholder("Dica: Arraste o arquivo ou a pasta para esta janela!").
		Validate(func(s string) error {
			if !tools.Exists(strings.Trim(s, "\"")) {
				return errors.New("arquivo ou pasta não existe")
			}
			return nil
		}).
		Prompt("→ ").
		Value(&input).
		Run()

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_dork_file()
	}
	
	reader, _ := tools.Reader(strings.Trim(input, "\""))
	b, err := io.ReadAll(reader)

	if err != nil {
		color.Red("● Um erro ocorreu, por-favor, tente novamente!")
		return step_ask_dork_file()
	}

	return strings.Split(string(b), "\n")
}
