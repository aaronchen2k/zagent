package serverService

import (
	"github.com/easysoft/zagent/internal/comm/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_stringUtils "github.com/easysoft/zagent/internal/pkg/lib/string"
	"github.com/easysoft/zagent/internal/server/model"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/easysoft/zagent/internal/server/service/vendors"
	"strings"
)

type AliyunDockerService struct {
	HostRepo    *repo.HostRepo    `inject:""`
	BackingRepo *repo.BackingRepo `inject:""`
	VmRepo      *repo.VmRepo      `inject:""`
	QueueRepo   *repo.QueueRepo   `inject:""`

	VmCommonService  *VmCommonService          `inject:""`
	HistoryService   *HistoryService           `inject:""`
	AliyunEciService *vendors.AliyunEciService `inject:""`
}

func NewAliyunDockerService() *AliyunDockerService {
	return &AliyunDockerService{}
}

func (s AliyunDockerService) CreateRemote(hostId, queueId uint) (result _domain.RpcResp) {
	queue := s.QueueRepo.GetQueue(queueId)
	host := s.HostRepo.Get(hostId)

	client, _ := s.AliyunEciService.CreateEciClient(host.CloudKey, host.CloudSecret, host.CloudRegion)
	cmd := []string{
		"/bin/bash",
		"-c",
		strings.Join(strings.Split(queue.BuildCommands, "\n"), "; "),
	}

	image := queue.DockerImage
	jobName := queue.TaskName + "-" + _stringUtils.NewUuid()

	id, err := s.AliyunEciService.CreateInst(jobName, jobName, image, cmd, host.CloudRegion, client)
	if err != nil {
		result.Fail(err.Error())
		return
	}

	vm := model.Vm{
		HostId:      host.ID,
		HostName:    host.Name,
		Status:      consts.VmCreated,
		CouldInstId: id,
	}
	s.VmRepo.Save(&vm)

	return
}
func (s AliyunDockerService) DestroyRemote(vmId, queueId uint) {
	vm := s.VmRepo.GetById(vmId)
	jobName := vm.CouldInstId

	host := s.HostRepo.Get(vm.HostId)

	client, err := s.AliyunEciService.CreateEciClient(host.CloudKey, host.CloudSecret, host.CloudRegion)

	var status consts.VmStatus
	if err == nil {
		_, err = s.AliyunEciService.Destroy(jobName, host.CloudRegion, client)
	}

	if err != nil {
		status = consts.VmFailDestroy
	}

	s.VmRepo.UpdateStatusByCloudInstId([]string{vm.CouldInstId}, status)
	s.HistoryService.Create(consts.Vm, vmId, queueId, "", status.ToString())

	return
}
