package v1

import "github.com/easysoft/zagent/internal/comm/const"

type KvmReq struct {
	VmMacAddress   string `json:"vmMacAddress"`
	VmBackingPath  string `json:"vmBacking"`
	VmTemplateName string `json:"vmTemplate"`

	VmUniqueName  string `json:"vmUniqueName"`
	VmCpu         uint   `json:"vmCpu"`
	VmMemorySize  uint   `json:"vmMemorySize"`
	VmDiskSize    uint   `json:"vmDiskSize"`
	VmCdromSys    string `json:"vmCdromSys"`
	VmCdromDriver string `json:"vmCdromDriver"`

	OsCategory consts.OsCategory `json:"osCategory" example:"windows"` // Enums consts.OsCategory
	OsType     consts.OsType     `json:"osType" example:"win10"`       // Enums consts.OsType
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang" example:"zh_cn"` // Enums consts.OsLang

	StartAfterCreated bool `json:"startAfterCreated"`
}

type KvmResponse struct {
	Code       int    `json:"code"`
	Msg        int    `json:"msg"`
	Name       string `json:"name"`
	VncAddress string `json:"vncAddress"`
	Path       string `json:"path"`
	Mac        string `json:"mac"`
}
