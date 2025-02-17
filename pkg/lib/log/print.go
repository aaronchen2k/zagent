package _logUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	consts "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/fatih/color"
)

func Info(str string) {
	logger.Infoln(str)
	if consts.PrintLog {
		log.Println(str)
	}
}
func Infof(str string, args ...interface{}) {
	logger.Infof(str, args...)
	if consts.PrintLog {
		log.Printf(str+"\n", args...)
	}
}

func Warn(str string) {
	logger.Warnln(str)
	if consts.PrintLog {
		log.Println(str)
	}
}
func Warnf(str string, args ...interface{}) {
	logger.Warnf(str, args...)
	if consts.PrintLog {
		log.Printf(str+"\n", args...)
	}
}

func Error(str string) {
	logger.Errorln(str)
	if consts.PrintLog {
		log.Println(str)
	}
}
func Errorf(str string, args ...interface{}) {
	logger.Errorf(str, args...)
	if consts.PrintLog {
		log.Printf(str+"\n", args...)
	}
}

func Print(str string) {
	logger.Println(str)
}
func Printf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Print(msg)
}

func PrintColor(msg string, attr color.Attribute) {
	if attr < 0 {
		Print(msg)
	} else {
		color.New(attr).Fprintf(color.Output, msg+"\n")
	}
}

func PrintUnicode(str []byte) {
	Print(ConvertUnicode(str))
}

func ConvertUnicode(str []byte) string {
	var a interface{}

	temp := strings.Replace(string(str), "\\\\", "\\", -1)

	err := json.Unmarshal([]byte(temp), &a)

	var msg string
	if err == nil {
		msg = fmt.Sprint(a)
	} else {
		msg = temp
	}

	return msg
}
