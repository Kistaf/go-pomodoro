package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	timer, mode := handleArguments()
	m := model{
		progress:  progress.New(progress.WithDefaultGradient()),
		countdown: Countdown{left: timer, total: timer, mode: mode},
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Something went wrong!", err)
		os.Exit(1)
	}
}

const (
	padding  = 2
	maxWidth = 30
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type model struct {
	progress  progress.Model
	countdown Countdown
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// quit
	case tea.KeyMsg:
		return m, tea.Quit

	// resize bar according to window
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	// update values
	case tickMsg:
		if m.countdown.left == 0 {
			return m, tea.Quit
		}

		incValue := m.countdown.next()
		cmd := m.progress.IncrPercent(incValue)
		return m, tea.Batch(tickCmd(), cmd)

		// animate
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)

	// display the bar and timer + mode
	return "\n" +
		pad + m.countdown.prettyStatus() + " | " + m.countdown.mode + "\n\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}
