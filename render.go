package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderTitle() string {
	ts := lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63")).
		Width(WIDTH).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Underline(true)
	return ts.Render("CAR RACING")
}

func RenderRaceTrack(m *model) {
	for i := 0; i < HEIGHT; i++ {
		m.racetrack = append(m.racetrack, strings.Split(WALL+strings.Repeat(" "+WALL, 6), ""))
	}
}

func RenderHelp() string {
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	return ts.Render("Press arrow keys to move your car.")
}

func RenderQuitcommand() string {
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	return ts.Render("Press ctrl+c to quit.")
}

func RenderScore(score int) string {
	scoreStr := fmt.Sprintf("Score: %d ", score)
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))

	return ts.Render(scoreStr)
}

func RenderGameOver() string {
	return lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Width(WIDTH).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Render("Game Over!")
}

func RenderCarCrashed() string {
	ts := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000"))
	return ts.Render(CAR)
}

func RenderMyCar() string {
	ts := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00"))
	return ts.Render(CAR)
}
