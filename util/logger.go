package util

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

type MyFormatter struct{}

type MyErrorHook struct{}

func init() {
	h := &MyErrorHook{}
	log.SetFormatter(&MyFormatter{})
	log.AddHook(h)
	res, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	curlog := filepath.Join(res, "logs")
	logf, err := rotatelogs.New(
		filepath.Join(curlog, "app.%Y%m%d.log"),
		rotatelogs.WithLinkName(filepath.Join(curlog, "applog.log")),
		// 180 day Time out auto delete
		rotatelogs.WithMaxAge(time.Duration(180*24)*time.Hour),
	)

	writers := []io.Writer{
		logf,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err != nil {
		log.Error("failed to create rotatelogs: %s", err)
	}
	log.SetOutput(fileAndStdoutWriter)
}

// Implement Formatter interface
// Format renders a single log entry
func (m *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	formatEntry := FormatEntry(entry)
	return formatEntry, nil
}

func FormatEntry(entry *log.Entry) []byte {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02T15:04:05.000")
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(timestamp)
	sb.WriteString("]")
	sb.WriteString(" ")
	sb.WriteString("[")
	sb.WriteString(entry.Level.String())
	sb.WriteString("]")
	sb.WriteString(" ")
	sb.WriteString(entry.Message)
	sb.WriteString("\n")
	fileVal := sb.String()
	b.WriteString(fileVal)
	return b.Bytes()
}

// Implement Hook interface
// Only the log level of interest is required
func (h *MyErrorHook) Levels() []log.Level {
	return []log.Level{
		log.ErrorLevel,
		log.PanicLevel,
	}
}

func (h *MyErrorHook) Fire(entry *log.Entry) error {
	res, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	curlog := filepath.Join(res, "logs")
	logf, err := rotatelogs.New(
		filepath.Join(curlog, "err.%Y%m%d.log"),
		rotatelogs.WithLinkName(filepath.Join(curlog, "err.log")),
		// 180 day Time out auto delete
		rotatelogs.WithMaxAge(time.Duration(180*24)*time.Hour),
	)
	formatEntry := FormatEntry(entry)
	if err != nil {
		return err
	}
	if _, err := logf.Write(formatEntry); err != nil {
		return err
	}
	return nil
}
