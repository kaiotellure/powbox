package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	ColorPurple    = lipgloss.Color("99")
	ColorGray      = lipgloss.Color("245")
	ColorLightGray = lipgloss.Color("241")
)

type model struct {
	metrics *Metrics
	cancel  context.CancelFunc
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewProgress(metrics *Metrics, cancel context.CancelFunc) model {
	return model{metrics, cancel}
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.QuitMsg:
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.cancel()
			return m, tea.Quit
		}
	case TickMsg:
		return m, doTick()
	}

	return m, nil
}

func (m model) View() string {
	now := time.Now().Unix()

	s := table.New().
		BorderStyle(lipgloss.NewStyle().Foreground(ColorPurple)).
		StyleFunc(func(r,c int) lipgloss.Style {
			return lipgloss.NewStyle().
				Align(lipgloss.Center).
				Padding(0, 1).
				Bold(true)
		}).
		Headers("ANALIZADOS", "ENCONTRADOS", "DURAÇÃO").
		Row(
			purple(strconv.Itoa(m.metrics.Total)),
			green(strconv.Itoa(m.metrics.Matches)),
			yellow(fmt.Sprintf("%ds", now - m.metrics.StartSecond)),
		).
		String()

	s += "\n\nq / ctrl+c " + muted("parar")
	return s
}
