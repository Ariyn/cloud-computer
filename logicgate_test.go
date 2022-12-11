package cloud_computer

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestFlagOrder(t *testing.T) {
	originalArgs := append(os.Args)
	os.Args = make([]string, 0)

	samples := make([]string, 0)
	for i := 0; i < 1000; i++ {
		samples = append(samples, "-inputs", fmt.Sprintf("%d", i))
	}
	os.Args = append(os.Args, samples...)
	defer func() {
		os.Args = originalArgs
	}()

	t.Run("", func(t *testing.T) {
		flag.Parse()

		for i, input := range Inputs {
			assert.Equal(t, strconv.Itoa(i), input)
		}
	})
}

func TestFlagReverseOrder(t *testing.T) {
	originalArgs := append(os.Args)
	os.Args = make([]string, 0)

	samples := make([]string, 0)
	for i := 0; i < 1000; i++ {
		samples = append(samples, "-inputs", fmt.Sprintf("%d", 1000-i))
	}
	os.Args = append(os.Args, samples...)
	defer func() {
		os.Args = originalArgs
	}()

	t.Run("", func(t *testing.T) {
		flag.Parse()

		for i, input := range Inputs {
			assert.Equal(t, strconv.Itoa(1000-i), input)
		}
	})
}
