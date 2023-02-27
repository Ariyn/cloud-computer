package cloud_computer

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"strconv"
	"syscall"
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

func Test_getSelectCaseSignals(t *testing.T) {
	t.Run("비어있는 시그널을 넣는 경우, 에러가 발생한다.", func(t *testing.T) {
		_, err := getSelectCaseSignals()
		assert.Error(t, err)
	})

	t.Run("시그널을 한 개 넣을경우, 해당 시그널을 기다리는 SelectCase가 나온다.", func(t *testing.T) {
		targetSignal := syscall.SIGHUP

		sc, err := getSelectCaseSignals(targetSignal)
		assert.NoError(t, err)
		assert.IsType(t, reflect.SelectCase{}, sc)
	})

	// TODO: 해당 시그널이 진짜 동일한 시그널인지 확인하는 테스트 필요
}
