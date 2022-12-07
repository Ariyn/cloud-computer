package cloud_computer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func readAsync(r *os.File) (status chan int) {
	status = make(chan int, 1)

	go func(reader *bufio.Reader) {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					continue
				}

				panic(err)
			}

			line = strings.TrimSpace(line)

			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			if !(value == 0 || value == 1) {
				panic(fmt.Errorf("incorrect input %d", value))
			}

			status <- value
		}
	}(bufio.NewReader(r))

	return
}

func writeAsync(w *os.File) (status chan int) {
	status = make(chan int, 1)

	go func() {
		for s := range status {
			_, err := w.WriteString(strconv.Itoa(s) + "\n")
			if err != nil {
				panic(err)
			}
		}
	}()

	return status
}

func openInput(path string) (file *os.File, err error) {
	return open(path, os.O_RDONLY|syscall.O_NONBLOCK)
}

func openOutput(path string) (file *os.File, err error) {
	return open(path, os.O_WRONLY)
}

func open(path string, flag int) (file *os.File, err error) {
	file, err = os.OpenFile(path, flag, 0600)
	return
}
