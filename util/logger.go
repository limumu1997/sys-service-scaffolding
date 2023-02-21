package util

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

type MyFormatter struct{}

type MyErrorHook struct{}

func init() {
	h := &MyErrorHook{}
	logf, err := rotateLogs("app")
	writers := []io.Writer{
		logf,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err != nil {
		log.Error("failed to create rotatelogs: %s", err)
	}
	log.SetFormatter(&MyFormatter{})
	log.AddHook(h)
	log.SetOutput(fileAndStdoutWriter)
}

// Implement Formatter interface
// Format renders a single log entry
func (m *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	formatEntry := formatEntry(entry)
	return formatEntry, nil
}

func formatEntry(entry *log.Entry) []byte {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02T15:04:05.000")
	fileVal := "[" + timestamp + "] " + "[" + entry.Level.String() + "] " + entry.Message + "\n"
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
	logf, err := rotateLogs("err")
	formatEntry := formatEntry(entry)
	if err != nil {
		return err
	}
	if _, err := logf.Write(formatEntry); err != nil {
		return err
	}
	return nil
}

func rotateLogs(filename string) (*rotatelogs.RotateLogs, error) {
	res, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	curlog := filepath.Join(res, "logs")
	return rotatelogs.New(
		filepath.Join(curlog, filename+".%Y%m%d.log"),
		rotatelogs.WithLinkName(filepath.Join(curlog, filename+".log")),
		// 180 day Time out auto delete
		rotatelogs.WithMaxAge(time.Duration(180*24)*time.Hour),
	)
}
