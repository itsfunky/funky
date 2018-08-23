package serve

import (
	"io"
	"log"
	"strings"
)

type logWriter struct {
	logger *log.Logger
}

// Write prints to the logger.
func (w logWriter) Write(b []byte) (int, error) {
	s := strings.TrimSuffix(string(b), "\n")
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		w.logger.Println(line)
	}

	return len(b), nil
}

func createLogWriter(prefix string, writer io.Writer) *logWriter {
	return &logWriter{
		logger: log.New(writer, prefix+" ", 0),
	}
}
