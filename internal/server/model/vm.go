package model

import (
	v1 "github.com/easysoft/zv/cmd/host/router/v1"
	"github.com/easysoft/zv/internal/pkg/const"
	"time"
)

type Vm struct {
	BaseModel

	Name      string `json:"name"`
	Desc      string `json:"desc"`
	ImagePath string `json:"imagePath"`

	HostId   uint   `json:"hostId"`
	HostName string `json:"hostName"`

	TmplId   uint   `json:"tmplId"`
	TmplName string `json:"tmplName"`

	BackingId   uint   `json:"backingId"`
	BackingPath string `json:"backingPath"`
	BackingName string `json:"backingName"`
	Bridge      string `json:"backingName"`

	OsCategory consts.OsCategory `json:"osCategory" example:"windows"` // Enums consts.OsCategory
	OsType     consts.OsType     `json:"osType" example:"win10"`       // Enums consts.OsType
	OsVersion  string            `json:"osVersion"`
	OsLang     consts.OsLang     `json:"osLang" example:"zh_cn"` // Enums consts.OsLang

	Status    consts.VmStatus `json:"status" example:"ready"` // Enums consts.VmStatus
	DestroyAt *time.Time      `json:"destroyAt"`

	FirstDetectedTime *time.Time `json:"firstDetectedTime"`
	LastRegisterTime  *time.Time `json:"lastRegisterTime"`

	NodeIp     string `json:"nodeIp"`
	NodePort   int    `json:"nodePort"`
	MacAddress string `json:"macAddress"`
	RpcPort    int    `json:"rpcPort"`
	SshPort    int    `json:"sshPort"`
	VncPort    int    `json:"vncPort"`
	VncUrl     string `json:"vncUrl"`
	WorkDir    string `json:"workDir"`

	DefPath          string `json:"defPath"`
	DiskSize         uint   `json:"diskSize"`   // M
	MemorySize       uint   `json:"memorySize"` // M
	CdromSys         string `json:"cdromSys"`
	CdromDriver      string `json:"cdromDriver"`
	ResolutionHeight int    `json:"resolutionHeight"`
	ResolutionWidth  int    `json:"resolutionWidth"`

	Histories []History `json:"histories" gorm:"polymorphic:Owner;polymorphicValue:vm"`

	CloudInstId string `json:"cloudInstId"`
	CloudEipId  string `json:"cloudEipId"`
}

func GenKvmReq(po Vm) (req v1.KvmReq) {
	req = v1.KvmReq{
		VmMacAddress: po.MacAddress, VmUniqueName: po.Name,
		VmBacking: po.BackingPath, VmTemplate: po.TmplName,

		OsCategory: po.OsCategory,
		OsType:     po.OsType,
		OsVersion:  po.OsVersion,
		OsLang:     po.OsLang,

		VmDiskSize: po.DiskSize, VmMemorySize: po.MemorySize,
		//VmCdromSys: po.CdromSys, VmCdromDriver: po.CdromDriver
	}

	return
}

func GenVirtualBoxReq(vm Vm, backing VmBacking, host Host) (req v1.VirtualBoxReq) {
	req = v1.VirtualBoxReq{
		VmUniqueName: vm.Name,
		BackingName:  vm.BackingName,

		VmCpu:        backing.SuggestCpuCount,
		VmMemorySize: backing.SuggestMemorySize,
	}

	return
}

func GenVmWareReq(vmName, backingName, vmId string, processors, memory uint, userName, password string) (req v1.VmWareReq) {
	req = v1.VmWareReq{
		VmUniqueName:  vmName,
		VmBackingName: backingName,

		VmProcessors: processors,
		VmMemory:     memory,

		UserName: userName,
		Password: password,

		VmId: vmId,
	}

	return
}

func (Vm) TableName() string {
	return "biz_vm"
}
