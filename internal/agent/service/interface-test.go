package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
)

type InterfaceTestService struct {
	CommonService
	ExecService *InterfaceExecService `inject:""`
}

func NewInterfaceTestService() *InterfaceTestService {
	return &InterfaceTestService{}
}

func (s *InterfaceTestService) ExecScenario(build *commDomain.IntfTest) (result commDomain.TestResult) {
	scenario := build.TestScenario
	s.ExecService.ExecProcessor(build, &scenario.Processor)

	// TODO: deal with result with logs in scenario.Processor

	result.Name = scenario.Name

	return
}

func (s *InterfaceTestService) ExecSet(build *commDomain.IntfTest, result *commDomain.TestResult) {
	set := build.TestSet

	// TODO:

	result.Name = set.Name
}
