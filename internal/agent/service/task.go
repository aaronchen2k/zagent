package agentService

import (
	consts "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	"sync"
	"time"
)

var lock sync.Mutex

type TaskService struct {
	TimeStamp time.Time
	running   bool
	tasks     []commDomain.Build
}

func NewTaskService() *TaskService {
	service := &TaskService{}

	service.TimeStamp = time.Now()
	service.tasks = make([]commDomain.Build, 0)

	return service
}

func (s *TaskService) AddTask(task commDomain.Build) {
	lock.Lock()

	s.tasks = append(s.tasks, task)

	lock.Unlock()
}

func (s *TaskService) PeekTask() commDomain.Build {
	lock.Lock()
	defer lock.Unlock()

	return s.tasks[0]
}

func (s *TaskService) RemoveTask() {
	lock.Lock()

	if len(s.tasks) == 0 {
		return
	}
	s.tasks = s.tasks[1:]

	lock.Unlock()
}

func (s *TaskService) StartTask() {
	lock.Lock()

	s.TimeStamp = time.Now()
	s.running = true

	lock.Unlock()
}
func (s *TaskService) EndTask() {
	lock.Lock()

	s.running = false

	lock.Unlock()
}

func (s *TaskService) GetTaskSize() int {
	lock.Lock()
	defer lock.Unlock()

	return len(s.tasks)
}

func (s *TaskService) IsRunning() bool {
	lock.Lock()
	defer lock.Unlock()

	if time.Now().Unix()-s.TimeStamp.Unix() > consts.AgentRunTime*60*1000 {
		s.running = false
	}
	return s.running
}
