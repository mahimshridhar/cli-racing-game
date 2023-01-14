package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const CAR = "#"
const NOCAR = " "
const WIDTH = 13
const HEIGHT = 20
const WALL = "|"

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
	traffic   []car
	myCar     MyCar
}

func initModel() model {
	return model{
		racetrack: [][]string{},
		gameOver:  false,
		score:     0,
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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.myCar.x-1 > 1 {
				m.myCar.x = m.myCar.x - 1
			}
		case "down":
			if m.myCar.x+1 < HEIGHT {
				m.myCar.x = m.myCar.x + 1
			}
			return m, nil
		case "right":
			if m.myCar.y+2 <= WIDTH-1 {
				m.myCar.y = m.myCar.y + 2
			}
		case "left":
			if m.myCar.y-2 >= 1 {
				m.myCar.y = m.myCar.y - 2
			}
		}
	case TickMsg:
		r := NewTraffic()

		if HEIGHT == len(m.traffic) {
			m.traffic = m.traffic[:len(m.traffic)-1]
		}
		m.traffic = append([]car{r}, m.traffic...)

		for i, randomCar := range m.traffic {
			if randomCar.position == m.myCar.y && i == m.myCar.x && randomCar.shape == CAR {
				m.gameOver = true
				return m, tea.Quit
			}
		}
		m.score += 1
		return m, m.tick()

	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	s.WriteString(RenderTitle())

	s.WriteRune('\n')

	RenderRaceTrack(&m)
	for i, val := range m.traffic {
		if val.shape == CAR {
			m.racetrack[i][val.position] = val.shape
		}
	}
	if m.gameOver {
		m.racetrack[m.myCar.x][m.myCar.y] = RenderCarCrashed()

	} else {
		m.racetrack[m.myCar.x][m.myCar.y] = RenderMyCar()
	}

	for _, row := range m.racetrack {
		s.WriteString(strings.Join(row, ""))
		s.WriteRune('\n')
	}
	s.WriteRune('\n')
	s.WriteString(RenderScore(m.score))
	s.WriteRune('\n')
	s.WriteString(RenderHelp())
	s.WriteRune('\n')
	s.WriteString(RenderQuitcommand())
	s.WriteRune('\n')
	s.WriteRune('\n')

	if m.gameOver {
		s.WriteString(RenderGameOver())

	}
	return s.String()
}
