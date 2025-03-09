package main

import (
	"context"
	t "powbox/internal/tools"

	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const EMPTY_STRING = ""
const ERR_UNEXPECTED = "um erro inesperado ocorreu, por-favor reporte para @kaiotellure. %v"



func compare(q Query, l *Login) {
	// Add up the total logins count
	q.Metrics.Total++
	
	// Check if login domain contains the search term
	if !strings.Contains(l.Host, q.FilterKey) {
		return
	}
	
	// Write login line to output text file
	line := fmt.Sprintf(q.OutFormat, l.Host, l.Username, l.Password)
	q.OutFile.WriteString(line + "\n")
	
	// Add up the matched logins count
	q.Metrics.Matches++

	/* q.Progress.Send(fmt.Sprintf(
		"\r%s há %s pesquisando por: (%s) txt°→(%s) encontrados→(%s/%s)",
		yellow(t.Spinner()),
		purple(fmt.Sprintf("%ds", time.Now().Unix()-q.StartSecond)),
		green(q.FilterKey),
		purple(fmt.Sprintf("%d", q.StepIndex)),
		purple(fmt.Sprintf("%d", matches_count)),
		yellow(fmt.Sprintf("%d", total_count)),
	)) */
}

func main() {
	// The initial meta informations displayed on startup
	s := fmt.Sprintf("Lofin %s 1.2\n", muted("●"))
	s += muted("Ferramenta para filtrar e transformar combo lists")
	fmt.Println(border(s))

	// Ask user for the input file or a folder of files
	input_files := step_ask_input_files()

	// Use the first's file directory as working dir
	iwd := t.Dir(input_files[0])

	var metrics Metrics
	// This will later be used to display the task's duration
	metrics.StartSecond = time.Now().Unix()

	var query Query
	query.Metrics = &metrics

	query.FilterKey = step_ask_filter_key()
	query.OutFormat = step_ask_out_format()

	// eg. results-2000-03-13.txt
	outFilename := t.TimedFilename(query.FilterKey, ".txt")

	// Concatenate path and ensure all anscenstor directories exist
	absoluteOutpath := path.Join(iwd, "resultados", outFilename)
	t.Ensure(absoluteOutpath)

	// The file filtered credentials will be appended to
	query.OutFile = t.Appendable(absoluteOutpath)
	defer query.OutFile.Close()

	ctx, cancel := context.WithCancel(context.Background())
	query.Context = ctx

	fmt.Print("\n")

	query.Progress = tea.NewProgram(
		NewProgress(&metrics, cancel),
		tea.WithContext(ctx),
	)

	go func() {
		if _, err := query.Progress.Run(); err != nil {
			panic(err)
		}
	}()

	for i, filepath := range input_files {
		query.StepIndex = i + 1
		process(filepath, query)
	}

	cancel()

	fmt.Printf(
		"%s %s %s %s\n",
		green("●"),
		purple("Concluído"),
		(fmt.Sprintf(t.LINK_FORMAT, "file://"+absoluteOutpath, "("+outFilename+")")),
		green("√"),
	)
}

func process(filepath string, query Query) {
	reader, file := t.Reader(filepath)
	defer file.Close()

	var clip string
	var args []string

	lastCheck := time.Now().Unix()

loop:
	for {
		select {
		case <-query.Context.Done():
			return
		default:
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
					compare(query, login)
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
				break loop
			}
		}
	}
}
