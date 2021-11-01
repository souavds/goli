package gcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArgs(t *testing.T) {

	t.Run("should build subcommand and args into a single array", func(t *testing.T) {
		args := BuildArgs("test", []string{"1", "2", "3"})

		assert.Equal(t, []string{"test", "1", "2", "3"}, args)
	})
}
