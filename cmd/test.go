/*
Copyright © 2021 Arthur Diniz <arthurvdinizs@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/arthurvdiniz/goli/internal/gcli"
	"github.com/charmbracelet/lipgloss"
	gcmd "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

type Styles struct {
	Key       string
	CheckKey  string
	TextStyle TextStyle
}

type TextStyle struct {
	Foreground string
	Background string
}

var styles = []Styles{
	{Key: "RUN", CheckKey: "=== RUN", TextStyle: TextStyle{Foreground: "#121212", Background: "#57aeff"}},
	{Key: "PASS", CheckKey: "--- PASS", TextStyle: TextStyle{Foreground: "#121212", Background: "#34eb71"}},
	{Key: "FAIL", CheckKey: "--- FAIL", TextStyle: TextStyle{Foreground: "#121212", Background: "#ff5757"}},
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:                "test",
	Short:              "go test {args} but with colorized output",
	Example:            "goli test './... -v'",
	Args:               cobra.ArbitraryArgs,
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	Run: func(cmd *cobra.Command, args []string) {
		arguments := gcli.BuildArgs("test", args)

		command := gcmd.NewCmd("go", arguments...)
		status := <-command.Start()

		for _, line := range status.Stdout {
			fmt.Println(colorizedLine(line))

		}

	},
}

func colorizedLine(line string) string {
	for _, s := range styles {
		hasKey := strings.Contains(line, s.CheckKey)

		if !hasKey {
			continue
		}

		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(s.TextStyle.Foreground)).
			Background(lipgloss.Color(s.TextStyle.Background)).
			PaddingLeft(1).
			PaddingRight(1)

		line = strings.Replace(line, s.Key, style.Render(s.Key), 1)
	}
	return line
}

func init() {
	rootCmd.AddCommand(testCmd)
}
