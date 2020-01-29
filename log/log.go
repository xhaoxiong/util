/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  log
 * @Version: 1.0.0
 * @Date: 10/24/19 12:54 PM
 */

package log

type Xlogger struct {
	level int
}

func NewXlogger() *Xlogger {
	return &Xlogger{}
}

func (xl *Xlogger) SetLevel(level int) {
	xl.level = level
}
