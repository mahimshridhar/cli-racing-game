package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type RaceTrack [][]string

type car struct {
	position int
}

type model struct {
	racetrack [][]string
	gameOver  bool
	score     int
	width     int
	height    int
	traffic   []car
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
	return tea.Tick(time.Second/10, func(t time.Time) tea.Msg {
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

		if m.height == len(m.traffic) {
			m.traffic = m.traffic[:len(m.traffic)-1]
		}
		m.traffic = append([]car{r}, m.traffic...)

		return m, m.tick()

	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	for i := 0; i < m.height; i++ {
		m.racetrack = append(m.racetrack, strings.Split("|"+strings.Repeat(" |", 6), ""))
	}

	// m.racetrack[0] = NewTraffic()

	for i, val := range m.traffic {
		m.racetrack[i][val.position] = "#"
	}

	for _, row := range m.racetrack {
		s.WriteString(strings.Join(row, ""))
		s.WriteRune('\n')
	}
	return s.String()
}
