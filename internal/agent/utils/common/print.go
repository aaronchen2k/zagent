package agentUtils

import (
	"fmt"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_commonUtils "github.com/easysoft/zagent/internal/pkg/libs/common"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	agentRes "github.com/easysoft/zagent/res/agent"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
)

var (
	usageFile = filepath.Join("res", "doc", "usage.txt")
)

func PrintUsage() {
	_logUtils.PrintColor("Usage: ", color.FgCyan)

	usage := ReadResData(usageFile)

	app := agentConst.AppName
	if _commonUtils.IsWin() {
		app += ".exe"
	}
	usage = fmt.Sprintf(usage, app)
	fmt.Printf("%s\n", usage)
}

func ReadResData(path string) string {
	isRelease := _commonUtils.IsRelease()

	var jsonStr string
	if isRelease {
		data, _ := agentRes.Asset(path)
		jsonStr = string(data)
	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			jsonStr = "fail to read " + path
		} else {
			str := string(buf)
			jsonStr = _commonUtils.RemoveBlankLine(str)
		}
	}

	return jsonStr
}
