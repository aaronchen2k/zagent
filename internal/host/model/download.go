package agentModel

import (
	consts "github.com/easysoft/zagent/internal/pkg/const"
	"time"
)

type Task struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`

	// for download
	Url      string            `json:"url,omitempty"`
	Md5      string            `json:"md5,omitempty"`
	Path     string            `json:"path,omitempty"`
	Status   consts.TaskStatus `json:"status"`
	ExecInfo string            `json:"execInfo" gorm:"column:execInfo"`
	Retry    int               `json:"retry"`
	Rate     float64           `json:"rate"`
	Speed    float64           `json:"speed,omitempty"`

	// for export vm
	Vm      string `json:"vm,omitempty"`
	Backing string `json:"backing,omitempty"`
	Xml     string `json:"xml,omitempty"`
	//Path    string `json:"path,omitempty"`

	StartDate   *time.Time `json:"startDate" gorm:"column:startDate"`
	EndDate     *time.Time `json:"endDate" gorm:"column:endDate"`
	TimeoutDate *time.Time `json:"timeoutDate,omitempty" gorm:"column:timeoutDate"`
	CancelDate  *time.Time `json:"cancelDate,omitempty" gorm:"column:cancelDate"`

	Task int             `json:"task" gorm:"column:task"`
	Type consts.TaskType `json:"type" gorm:"type"`
}

func (Task) TableName() string {
	return "biz_task"
}
