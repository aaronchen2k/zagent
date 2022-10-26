package hostRepo

import (
	agentModel "github.com/easysoft/zv/internal/host/model"
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/server/model"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"gorm.io/gorm"
	"time"
)

type TaskRepo struct {
	DB *gorm.DB `inject:""`
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r *TaskRepo) Query() (ret map[consts.DownloadProgress][]agentModel.Download, err error) {
	ret = map[consts.DownloadProgress][]agentModel.Download{}

	pos := make([]agentModel.Download, 0)

	err = r.DB.Model(&agentModel.Download{}).
		Where("NOT deleted").
		Find(&pos).Error

	for _, po := range pos {
		//_, ok := ret[po.Status]
		//
		//if !ok {
		//	ret[po.Status] = make([]agentModel.Start, 0)
		//}

		ret[po.Status] = append(ret[po.Status], po)
	}

	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *TaskRepo) Get(id uint) (po agentModel.Download) {
	r.DB.Model(&agentModel.Download{}).Preload("Environments", "NOT deleted").Where("id = ?", id).First(&po)

	return
}
func (r *TaskRepo) GetDetail(id uint) (po agentModel.Download) {
	r.DB.Model(&agentModel.Download{}).
		Preload("Environments", "NOT deleted").
		Preload("Histories", "NOT deleted").
		Where("id = ?", id).First(&po)

	return
}

func (r *TaskRepo) Save(po *agentModel.Download) (err error) {
	err = r.DB.Model(&agentModel.Download{}).Create(&po).Error
	return
}

func (r *TaskRepo) Update(po *agentModel.Download) (err error) {
	err = r.DB.Model(&agentModel.Download{}).Where("task_id = ?", po.ID).Delete(&model.Environment{}).Error
	err = r.DB.Model(&agentModel.Download{}).Where("id = ?", po.ID).
		Session(&gorm.Session{FullSaveAssociations: true}).Updates(&po).Error
	return
}

func (r *TaskRepo) EndTask(id uint) (err error) {
	err = r.DB.Model(&agentModel.Download{}).Where("id = ?", id).
		Updates(map[string]interface{}{"status": consts.End, "end_time": time.Now()}).Error

	return
}

func (r *TaskRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&agentModel.Download{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *TaskRepo) SetTimeout(po agentModel.Download) (err error) {
	r.DB.Model(&agentModel.Download{}).Where("id=?", po.ID).Updates(
		map[string]interface{}{"status": consts.Timeout, "timeout_time": time.Now()})
	return
}
