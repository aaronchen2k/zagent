package _logUtils

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/var"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/log"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var logger *logrus.Logger

func Init(app string) {
	if logger != nil {
		return
	}

	usr, _ := user.Current()
	log.Info("RunRemote as user " + usr.Username)

	_var.WorkDir = addPathSepIfNeeded(filepath.Join(usr.HomeDir, consts.AppName))
	logDir := addPathSepIfNeeded("log")
	MkDirIfNeeded(logDir)
	log.Info("Log dir is " + logDir)

	logger = logrus.New()
	logger.Out = ioutil.Discard

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  logDir + "info.log",
		logrus.WarnLevel:  logDir + "info.log",
		logrus.ErrorLevel: logDir + "error.log",
	}

	logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&MyFormatter{},
	))

	logger.SetFormatter(&MyFormatter{})

	return
}

type MyFormatter struct {
	logrus.TextFormatter
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message + "\n"), nil
}

func addPathSepIfNeeded(pth string) string {
	sepa := _const.PthSep

	if strings.LastIndex(pth, sepa) < len(pth)-1 {
		pth += sepa
	}
	return pth
}

func MkDirIfNeeded(dir string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}

	return nil
}
func FileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
