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

	"github.com/arthurvdiniz/goli/internal/colorizer"
	"github.com/arthurvdiniz/goli/internal/gcli"
	gcmd "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:                "test",
	Short:              "go test {args} but with colorized output",
	Example:            "goli test ./... -v",
	Args:               cobra.ArbitraryArgs,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		color := colorizer.New()
		arguments := gcli.BuildArgs("test", args)

		command := gcmd.NewCmd("go", arguments...)
		status := <-command.Start()

		for _, line := range status.Stdout {
			fmt.Println(color.Format(line))
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
