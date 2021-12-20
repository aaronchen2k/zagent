package v1

import "github.com/easysoft/zv/internal/comm/const"

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

type KvmResp struct {
	Name        string `json:"name"`
	MacAddress  string `json:"macAddress"`
	VncPort     string `json:"vncPort"`
	ImagePath   string `json:"imagePath"`
	BackingPath string `json:"backingPath"`
}
