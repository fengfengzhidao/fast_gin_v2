package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
)

type MyLog struct {
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(MyLog{})
	logrus.AddHook(&MyHook{
		logPath: "logs",
	})
}

type MyHook struct {
	file     *os.File
	errFile  *os.File
	fileDate string
	logPath  string
	mu       sync.Mutex
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()

	timer := entry.Time.Format("2006-01-02")
	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("failed to format log entry: %v", err)
	}

	if hook.fileDate != timer {
		if err := hook.rotateFiles(timer); err != nil {
			return err
		}
	}

	if _, err := hook.file.Write([]byte(line)); err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	if entry.Level <= logrus.ErrorLevel {
		if _, err := hook.errFile.Write([]byte(line)); err != nil {
			return fmt.Errorf("failed to write to error log file: %v", err)
		}
	}

	return nil
}

// rotateFiles 日志轮换
func (hook *MyHook) rotateFiles(timer string) error {
	if hook.file != nil {
		if err := hook.file.Close(); err != nil {
			return fmt.Errorf("failed to close log file: %v", err)
		}
	}
	if hook.errFile != nil {
		if err := hook.errFile.Close(); err != nil {
			return fmt.Errorf("failed to close error log file: %v", err)
		}
	}

	dirName := fmt.Sprintf("%s/%s", hook.logPath, timer)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	infoFilename := fmt.Sprintf("%s/info.log", dirName)
	errFilename := fmt.Sprintf("%s/err.log", dirName)

	var err error
	hook.file, err = os.OpenFile(infoFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	hook.errFile, err = os.OpenFile(errFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %v", err)
	}

	hook.fileDate = timer
	return nil
}

// Levels 哪些级别的日志能走 Fire 方法
func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
