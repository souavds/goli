package gcli

// BuildArgs expect a subcommand and its args to join all together
func BuildArgs(subcommand string, args []string) []string {
	return append([]string{subcommand}, args...)
}
