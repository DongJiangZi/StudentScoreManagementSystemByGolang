package async

import (
	"fmt"
	"os"
	"time"
)

type LogEntry struct {
	Level   string
	Message string
	Time    time.Time
}

type AsyncLogger struct {
	logChan chan LogEntry
	file    *os.File
}

func NewAsyncLogger(filePath string) (*AsyncLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := &AsyncLogger{
		logChan: make(chan LogEntry, 100),
		file:    file,
	}

	go logger.run()

	return logger, nil
}

func (l *AsyncLogger) run() {
	for entry := range l.logChan {
		line := fmt.Sprintf("[%s] [%s] %s",
			entry.Time.Format("2006-01-02 15:04:05"),
			entry.Level,
			entry.Message,
		)

		l.file.WriteString(line)
	}
}

func (l *AsyncLogger) Info(msg string) {
	l.logChan <- LogEntry{
		Level:   "INFO",
		Message: msg,
		Time:    time.Now(),
	}
}

func (l *AsyncLogger) Error(msg string) {
	l.logChan <- LogEntry{
		Level:   "ERROR",
		Message: msg,
		Time:    time.Now(),
	}
}

func (l *AsyncLogger) Close() {
	close(l.logChan)
	l.file.Close()
}
