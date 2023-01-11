package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type RaceTrack [][]string

type model struct {
	racetrack [][]string
	gameOver  bool
	score     int
	width     int
	height    int
}

func initModel() model {
	return model{
		racetrack: [][]string{},
		gameOver:  false,
		score:     0,
		width:     13,
		height:    20,
	}
}

type TickMsg time.Time

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second/2, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return m.tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		//quit game
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case TickMsg:
		r := NewTraffic()
		m.racetrack = append([][]string{r}, m.racetrack...)
		m.racetrack = m.racetrack[1:]
		return m, m.tick()

	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	for i := 0; i < m.height; i++ {
		m.racetrack = append(m.racetrack, strings.Split("|"+strings.Repeat(" |", 6), ""))
	}

	m.racetrack[0] = NewTraffic()

	for _, row := range m.racetrack {
		s.WriteString(strings.Join(row, ""))
		s.WriteRune('\n')
	}
	return s.String()
}
