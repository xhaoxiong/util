/**
 * @Author xiaoxiao
 * @Description CREATE FILE config
 * @Date 2020/8/4 3:52 下午
 **/
package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

const (
	WithMaxAge        = "7*24h"
	WithRotationTime  = "24h"
	WithRotationCount = 5
	LogName           = "dayliy"
)

func InitLog() {
	withMaxAge, _ := time.ParseDuration(WithMaxAge)
	withRotationTime, _ := time.ParseDuration(WithRotationTime)
	writer, err := rotatelogs.New(
		LogName+"_%Y%m%d.log",
		rotatelogs.WithLinkName(LogName),
		rotatelogs.WithMaxAge(withMaxAge),
		rotatelogs.WithRotationTime(withRotationTime),
		rotatelogs.WithRotationCount(WithRotationCount),
	)

	writer1 := os.Stdout
	if err != nil {
		log.Fatalf("create file log faild:%v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer, writer1))
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&nested.Formatter{
		FieldsOrder:           []string{"component", "category"},
		HideKeys:              false,
		NoColors:              false,
		NoFieldsColors:        false,
		ShowFullLevel:         false,
		TrimMessages:          false,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
		TimestampFormat:       time.RFC3339,
	})
}
