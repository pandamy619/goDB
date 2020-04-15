package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogObj struct {
	Path string
	FormatLogFile string
	file *os.File
}

func (l *LogObj) checkDir() {
	_, err := os.Stat(l.Path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(l.Path,0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func (l *LogObj) nameLog() string{
	timeNow := time.Now()
	name := fmt.Sprintf("%s%s.log",
		l.Path,
		timeNow.Format(l.FormatLogFile))
	return name
}

func (l *LogObj) createFile() {
	l.checkDir()
	f, err := os.OpenFile(l.nameLog(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		log.Println(err)
	}
	l.file = f
}

func NewLog(path string, formatLogFile string, prefix string) (*log.Logger, *os.File){
	if prefix == "" {
		prefix = "prefix"
	}
	l := LogObj{
		Path: path,
		FormatLogFile: formatLogFile,
		file: nil,
	}
	l.createFile()
	logger := log.New(l.file, prefix, log.LstdFlags)
	return logger, l.file
}

func CloseLog(file *os.File) {
	_ = file.Close()
}
