package hostHandler

import (
	v1 "github.com/easysoft/zagent/cmd/host/router/v1"
	"github.com/easysoft/zagent/internal/host/service/kvm"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	natHelper "github.com/easysoft/zagent/internal/pkg/utils/net"
	_httpUtils "github.com/easysoft/zagent/pkg/lib/http"
	"github.com/kataras/iris/v12"
)

type KvmCtrl struct {
	KvmService     *kvmService.KvmService     `inject:""`
	LibvirtService *kvmService.LibvirtService `inject:""`
}

func NewKvmCtrl() *KvmCtrl {
	return &KvmCtrl{}
}

//// ListTmpl
//// @summary 获取KVM虚拟机模板信息
//// @Produce json
//// @Success 200 {object} _domain.Response{data=[]v1.KvmRespTempl} "code = success | fail"
//// @Router /api/v1/kvm/listTempl [get]
//func (c *KvmCtrl) ListTmpl(ctx iris.Context) {
//	templs, err := c.LibvirtService.ListTmpl()
//
//	if err != nil {
//		ctx.JSON(_httpUtils.RespData(consts.ResultFail, "fail to list vm tmpl", err))
//		return
//	}
//
//	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to list vm tmpl", templs))
//
//	return
//}

// Create
// @summary 创建KVM虚拟机
// @Accept json
// @Produce json
// @Param CreateVmReq body v1.CreateVmReq true "Create Kvm Request Object"
// @Success 200 {object} _domain.Response{data=v1.CreateVmResp} "code = success | fail"
// @Router /api/v1/kvm/create [post]
func (c *KvmCtrl) Create(ctx iris.Context) {
	req := v1.CreateVmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	dom, macAddress, vmVncPort, err :=
		c.KvmService.CreateVmFromImage(&req, true)

	vmStatus := consts.VmLaunch
	if err != nil || dom == nil {
		vmStatus = consts.VmFailCreate
	}

	vm := v1.CreateVmResp{
		Mac:    macAddress,
		Vnc:    vmVncPort,
		Status: vmStatus,
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to create vm", vm))

	return
}

// ExportVm
// @summary 导出KVM虚拟机为模板镜像
// @Accept json
// @Produce json
// @Param ExportVmReq body v1.ExportVmReq true "Export Kvm Request Object"
// @Success 200 {object} _domain.Response "code = success | fail"
// @Router /api/v1/kvm/exportVm [post]
func (c *KvmCtrl) ExportVm(ctx iris.Context) {
	req := v1.ExportVmReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err := c.KvmService.AddExportVmTask(req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to add export vm task", nil))
	return
}

//// Clone
//// @summary 克隆KVM虚拟机
//// @Accept json
//// @Produce json
//// @Param CloneVmReq body v1.CloneVmReq true "Kvm Clone Request Object"
//// @Success 200 {object} _domain.Response{data=v1.KvmResp} "code = success | fail"
//// @Router /api/v1/kvm/clone [post]
//func (c *KvmCtrl) Clone(ctx iris.Context) {
//	req := v1.CloneVmReq{}
//	err := ctx.ReadJSON(&req)
//	if err != nil {
//		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
//		return
//	}
//
//	if req.VmSrc == "" {
//		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "request vmSrc field can not be empty.", nil))
//		return
//	}
//
//	dom, vmIp, vmVncPort, vmAgentPortMapped, vmRawPath, vmBackingPath, err := c.LibvirtService.CloneVm(&req, true)
//
//	if err != nil {
//		ctx.JSON(_httpUtils.RespData(consts.ResultFail, "fail to clone vm", err))
//		return
//	}
//
//	vmName, _ := dom.GetName()
//	vm := v1.KvmResp{
//		Name:    vmName,
//		Ip:      vmIp,
//		Mac:     req.VmMacAddress,
//		Agent:   vmAgentPortMapped,
//		Vnc:     vmVncPort,
//		Image:   vmRawPath,
//		Backing: vmBackingPath,
//	}
//
//	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to clone vm", vm))
//
//	return
//}

// Destroy
// @summary 摧毁KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _domain.Response{data=string} "code = success | fail"
// @Router /api/v1/kvm/{name}/destroy [post]
func (c *KvmCtrl) Destroy(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	req := v1.DestroyVmReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err = c.LibvirtService.SafeDestroyVmByName(name)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	err = natHelper.RemoveForward(req.Ip, 0, consts.All)
	if err != nil {
		ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to destroy vm", name))
	return
}

// Reboot
// @summary 重启KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _domain.Response{data=string} "code = success | fail"
// @Router /api/v1/kvm/{name}/reboot [post]
func (c *KvmCtrl) Reboot(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.RebootVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to reboot vm", name))
	return
}

// Suspend
// @summary 暂停KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _domain.Response{data=string} "code = success | fail"
// @Router /api/v1/kvm/{name}/suspend [post]
func (c *KvmCtrl) Suspend(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	err := c.LibvirtService.SuspendVmByName(name)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to suspend vm", name))
	return
}

// Resume
// @summary 恢复KVM虚拟机
// @Accept json
// @Produce json
// @Param name path string true "Kvm Name"
// @Success 200 {object} _domain.Response{data=string} "code = success | fail"
// @Router /api/v1/kvm/{name}/resume [post]
func (c *KvmCtrl) Resume(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	err := c.LibvirtService.ResumeVmByName(name)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to resume vm", name))
	return
}

func (c *KvmCtrl) Boot(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.BootVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to boot vm", name))
	return
}

func (c *KvmCtrl) Shutdown(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	if name == "" {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "vm name is empty", nil))
		return
	}

	c.LibvirtService.ShutdownVmByName(name)

	ctx.JSON(_httpUtils.RespData(consts.ResultPass, "success to shutdown vm", nil))
	return
}
