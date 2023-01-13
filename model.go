package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type RaceTrack [][]string

type car struct {
	position int
	shape    string
}

type MyCar struct {
	x int
	y int
}

type model struct {
	racetrack [][]string
	gameOver  bool
	score     int
	width     int
	height    int
	traffic   []car
	myCar     MyCar
}

func initModel() model {
	return model{
		racetrack: [][]string{},
		gameOver:  false,
		score:     0,
		width:     13,
		height:    20,
		myCar:     MyCar{x: 19, y: 1},
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
		case "up":
			if m.myCar.x-1 > 1 {
				m.myCar.x = m.myCar.x - 1
			}
		case "down":
			if m.myCar.x+1 < m.height {
				m.myCar.x = m.myCar.x + 1
			}
			return m, nil
		case "right":
			if m.myCar.y > 0 {
				m.myCar.y = m.myCar.y + 2
			}
		case "left":
			if m.myCar.y > 0 {
				m.myCar.y = m.myCar.y - 2
			}
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

	for i, val := range m.traffic {
		m.racetrack[i][val.position] = val.shape
	}

	m.racetrack[m.myCar.x][m.myCar.y] = "X"

	for _, row := range m.racetrack {
		s.WriteString(strings.Join(row, ""))
		s.WriteRune('\n')
	}
	return s.String()
}
