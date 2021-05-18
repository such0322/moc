package logs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type fileLogWriter struct {
	Filename   string
	fileWriter *os.File

	Perm int

	fileNameOnly, suffix string
}

func NewFileWriter() Logger {
	w := &fileLogWriter{
		Filename: "logs/server.log",
		Perm:     0664,
	}
	return w
}

func (w *fileLogWriter) Init() error {
	if len(w.Filename) == 0 {
		return fmt.Errorf("must have filename")
	}

	w.suffix = filepath.Ext(w.Filename)
	w.fileNameOnly = strings.TrimSuffix(w.Filename, w.suffix)
	if w.suffix == "" {
		w.suffix = ".log"
	}

	file, err := w.createLogFile()
	if err != nil {
		return err
	}
	w.fileWriter = file

	return nil
}

func (w *fileLogWriter) createLogFile() (*os.File, error) {
	filepath := path.Dir(w.Filename)
	os.MkdirAll(filepath, os.FileMode(w.Perm))

	fd, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(w.Perm))
	if err == nil {
		os.Chmod(w.Filename, os.FileMode(w.Perm))
	}
	return fd, err
}

func (w *fileLogWriter) WriteMsg(msg string) error {
	_, err := w.fileWriter.Write([]byte(fmt.Sprintf("%s\n", msg)))
	if err != nil {
		return err
	}
	return nil
}

func (w *fileLogWriter) Flush() {
	w.fileWriter.Sync()
}

func (w *fileLogWriter) Destroy() {
	w.fileWriter.Close()
}
