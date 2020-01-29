/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 10/24/19 12:38 PM
 */

package log

import "time"

type Logger interface {
	Init(config string) error
	WriteMsg(When time.Time, msg string, level int) error
	Destroy()
	Flush()
}
type newLoggerFunc func() Logger

var adapters = make(map[string]newLoggerFunc)

func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("logs: Register called twice for provider " + name)
	}
	adapters[name] = log
}


