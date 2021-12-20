package domain

import (
	"fmt"
	_const "github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type TestResult struct {
	TestSetId uint `json:"testSetId" yaml:"testSetId"`

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`

	StartTime time.Time `json:"startTime" yaml:"startTime"`
	EndTime   time.Time `json:"endTime" yaml:"endTime"`
	Duration  int       `json:"duration" yaml:"duration"` // sec

	TotalNum  int `json:"totalNum" yaml:"totalNum"`
	PassNum   int `json:"passNum" yaml:"passNum"`
	FailNum   int `json:"failNum" yaml:"failNum"`
	MissedNum int `json:"missedNum" yaml:"missedNum"`

	Payload interface{} `json:"payload"`
}

func (result *TestResult) Pass(msg string) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = msg
}
func (result *TestResult) Passf(str string, args ...interface{}) {
	result.Code = _const.ResultSuccess.Int()
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestResult) Fail(msg string) {
	result.Code = _const.ResultFail.Int()
	result.Msg = msg
}

func (result *TestResult) Failf(str string, args ...interface{}) {
	result.Code = _const.ResultFail.Int()
	result.Msg = fmt.Sprintf(str+"\n", args...)
}

func (result *TestResult) IsSuccess() bool {
	return result.Code == _const.ResultSuccess.Int()
}
