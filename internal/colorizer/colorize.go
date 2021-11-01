package colorizer

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Colorize struct {
	TextStyles []TextStyle
}

type TextStyle struct {
	Key        string
	CheckKey   string
	Foreground string
	Background string
}

var defaultStyles = []TextStyle{
	{Key: "RUN", CheckKey: "=== RUN", Foreground: "#121212", Background: "#57aeff"},
	{Key: "PASS", CheckKey: "--- PASS", Foreground: "#121212", Background: "#34eb71"},
	{Key: "FAIL", CheckKey: "--- FAIL", Foreground: "#121212", Background: "#ff5757"},
}

func New() *Colorize {
	return &Colorize{
		TextStyles: defaultStyles,
	}
}

func NewWithCustomColors(textStyles []TextStyle) *Colorize {
	return &Colorize{
		TextStyles: textStyles,
	}
}

func (c *Colorize) Format(str string) string {
	for _, s := range c.TextStyles {
		hasKey := strings.Contains(str, s.CheckKey)

		if !hasKey {
			continue
		}

		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(s.Foreground)).
			Background(lipgloss.Color(s.Background)).
			PaddingLeft(1).
			PaddingRight(1)

		str = strings.Replace(str, s.Key, style.Render(s.Key), 1)
	}
	return str
}
