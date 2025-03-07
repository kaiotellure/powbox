package main

import (
	t "powbox/internal/tools"

	"strings"
	"errors"
	"path"
	"time"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

const EMPTY_STRING = ""
const ERR_UNEXPECTED = "um erro inesperado ocorreu, por-favor reporte para @kaiotellure. %v"

var matches_count int
var total_count int

func compare(q Query, l *Login) {
	total_count++

	if !strings.Contains(l.Host, q.FilterKey) {
		return
	}

	line := fmt.Sprintf(q.OutFormat, l.Host, l.Username, l.Password)
	q.OutFile.WriteString(line + "\n")

	matches_count++

	fmt.Printf(
		"\r  %s há %s pesquisando por: (%s) txt°→(%s) encontrados→(%s/%s)",
		color.YellowString(t.Spinner()),
		color.YellowString(fmt.Sprintf("%ds", time.Now().Unix()-q.StartSecond)),
		color.YellowString(q.FilterKey),
		color.YellowString(fmt.Sprintf("%d", q.StepIndex)),
		color.GreenString(fmt.Sprintf("%d", matches_count)),
		color.YellowString(fmt.Sprintf("%d", total_count)),
	)
}

func main() {
	input_files := step_ask_input_files()

	// root directory of the input file
	iwd := t.Dir(input_files[0])

	var query Query
	query.FilterKey = step_ask_filter_key()
	query.OutFormat = step_ask_out_format()

	query.StartSecond = time.Now().Unix()

	// eg. resultados-2000-03-13.txt
	outFilename := t.TimedFilename(query.FilterKey, ".txt")

	absoluteOutpath := path.Join(iwd, "resultados", outFilename)
	t.Ensure(absoluteOutpath)

	// the appendable result output file
	query.OutFile = t.Appendable(absoluteOutpath)
	defer query.OutFile.Close()

	for i, filepath := range input_files {
		query.StepIndex = i + 1
		process(filepath, query)
	}

	color.Green(
		"\n\n● Concluído → %s %s",
		color.CyanString(t.LINK_FORMAT, "file://"+absoluteOutpath, "("+outFilename+")"),
		color.GreenString("√"),
	)
}

func process(filepath string, query Query) {
	reader, file := t.Reader(filepath)
	defer file.Close()

	var clip string
	var args []string

	lastCheck := time.Now().Unix()

	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			if errors.Is(err, io.EOF) {
				// When we hit the End Of File there may be a login still being captured, so its left over password is left on clip, we just do the same as an end of line.
				args = append(args, clip)
				login := coerceLogin(args)

				if login != nil {
					compare(query, login)
				}
				break
			}

			fmt.Printf(ERR_UNEXPECTED, err)
			os.Exit(1)
		}

		switch char {
		case ':':
			args = append(args, clip)
			clip = EMPTY_STRING
		case '\n':
			args = append(args, clip)

			login := coerceLogin(args)
			if login != nil {
				go compare(query, login)
				lastCheck = time.Now().Unix()
			}

			// clear vars for next line
			clip = EMPTY_STRING
			args = []string{}
		default:
			// store character
			clip += string(char)
		}

		// when inactive for more than 1 second, means we're finished!
		if time.Now().Unix()-lastCheck > 1 {
			break
		}
	}
}
