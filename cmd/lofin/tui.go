package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	border = lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("250")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).Render
	
	muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color("238")).Render
	
	warn = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Bold(true).Render
	
	yellow = lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).Render

	green = lipgloss.NewStyle().
		Foreground(lipgloss.Color("50")).Render

	purple = lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).Render

	pink = lipgloss.NewStyle().
		Foreground(lipgloss.Color("200")).Render
)
