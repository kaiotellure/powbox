package main

import (
	"os"
	"strings"
)

type Query struct {
	FilterKey   string
	OutFormat   string
	OutFile     *os.File
	StepIndex   int
	StartSecond int64
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
