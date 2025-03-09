package main

import (
	"context"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Metrics struct {
	Total   int
	Matches int

	StartSecond int64
}

type Query struct {
	Metrics     *Metrics
	FilterKey   string
	OutFormat   string
	OutFile     *os.File
	StepIndex   int
	Progress    *tea.Program
	Context     context.Context
}

type Login struct {
	Host     string
	Username string
	Password string
}

func coerceLogin(args []string) (l *Login) {
	if len(args) < 3 {
		return nil
	}
	l = &Login{}

	host := args[:len(args)-2]
	l.Host = strings.Join(host, ":")
	user := args[len(args)-2:]
	l.Username = user[0]
	l.Password = user[1]

	return
}
