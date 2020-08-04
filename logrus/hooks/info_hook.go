/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  info_hook
 * @Version: 1.0.0
 * @Date: 2020/8/4 10:48 上午
 */
package hooks

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type AppHook struct {
	Skip    int
	AppName string
}

func NewContextHook(appName string) *AppHook {
	return &AppHook{AppName: appName, Skip: 5}
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	entry.Data["line"] = findCaller(h.Skip)
	return nil
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
