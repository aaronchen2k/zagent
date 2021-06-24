package service

import (
	"fmt"
	bizCasbin "github.com/easysoft/zagent/internal/server/biz/casbin"
	"github.com/easysoft/zagent/internal/server/repo"
	"github.com/fatih/color"
	"strconv"
)

type CommonService struct {
	CommonRepo    *repo.CommonRepo         `inject:""`
	CasbinService *bizCasbin.CasbinService `inject:""`
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

// GetPermissionsForUser 获取角色权限
func (s *CommonService) GetPermissionsForUser(uid uint) [][]string {
	return s.CasbinService.Enforcer.GetPermissionsForUser(strconv.FormatUint(uint64(uid), 10))
}

// GetRolesForUser 获取角色
func (s *CommonService) GetRolesForUser(uid uint) []string {
	uids, err := s.CasbinService.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}
