package main

import (
	"flag"
	"fmt"

	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_checkUtils "github.com/easysoft/zagent/internal/pkg/utils/check"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	_shellUtils "github.com/easysoft/zagent/pkg/lib/shell"
)

var (
	softList     string
	forceInstall bool
	check        bool
	secret       string
)

func main() {

	flag.StringVar(&softList, "s", "", "")
	flag.BoolVar(&forceInstall, "r", false, "")
	flag.BoolVar(&check, "c", false, "")
	flag.StringVar(&secret, "secret", "", "")
	flag.Parse()

	consts.PrintLog = false
	agentConf.Inst.RunMode = consts.RunModeHost
	agentConf.Init(consts.AppNameAgentVm)
	_logUtils.Init(consts.AppNameAgentHost)

	if check {
		status, _ := _checkUtils.CheckAgent()
		_checkUtils.CheckPrint("zagent", status)
		status, _ = _checkUtils.CheckNginx()
		_checkUtils.CheckPrint("nginx", status)
		status, _ = _checkUtils.CheckKvm()
		_checkUtils.CheckPrint("kvm", status)
		status, _ = _checkUtils.CheckNovnc()
		_checkUtils.CheckPrint("novnc", status)
		status, _ = _checkUtils.CheckWebsockify()
		_checkUtils.CheckPrint("websockify", status)
	} else if softList != "" {
		consts.PrintLog = true
		cmd := fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -s %s`, "https://pkg.qucheng.com/zenagent/zagent.sh", softList)

		if forceInstall {
			cmd = fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -s %s -r`, "https://pkg.qucheng.com/zenagent/zagent.sh", softList)
		}

		if secret != "" {
			cmd = fmt.Sprintf("%s -k %s", cmd, secret)
		}

		_shellUtils.ExeShellWithOutput(cmd)
	} else {
		consts.PrintLog = true
		cmd := fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s)`, "https://pkg.qucheng.com/zenagent/zagent.sh")

		if forceInstall {
			cmd = fmt.Sprintf(`/usr/bin/bash <(curl -s -S -L %s) -r`, "https://pkg.qucheng.com/zenagent/zagent.sh")
		}

		if secret != "" {
			cmd = fmt.Sprintf("%s -k %s", cmd, secret)
		}

		_shellUtils.ExeShellWithOutput(cmd)
	}
}
