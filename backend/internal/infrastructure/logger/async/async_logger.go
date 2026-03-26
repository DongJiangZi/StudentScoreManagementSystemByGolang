package async

import (
	"fmt"
	"os"
	"sync"
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
	wg      sync.WaitGroup
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

	logger.wg.Add(1)
	go logger.run()

	return logger, nil
}

func (l *AsyncLogger) run() {
	defer l.wg.Done()

	for entry := range l.logChan {
		line := fmt.Sprintf("[%s] [%s] %s\n",
			entry.Time.Format("2006-01-02 15:04:05"),
			entry.Level,
			entry.Message,
		)

		if _, err := l.file.WriteString(line); err != nil {
			fmt.Fprintf(os.Stderr, "write log failed: %v\n", err)
		}
	}
}

func (l *AsyncLogger) Info(msg string) {
	l.logChan <- LogEntry{
		Level:   "INFO",
		Message: msg,
		Time:    time.Now(),
	}
}

func (l *AsyncLogger) Warn(msg string) {
	l.logChan <- LogEntry{
		Level:   "WARN",
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
	l.wg.Wait()
	_ = l.file.Sync()
	l.file.Close()
}
