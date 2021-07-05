package agentService

import (
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_gitUtils "github.com/easysoft/zagent/internal/pkg/lib/git"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mholt/archiver/v3"
	"github.com/satori/go.uuid"
	"os"
	"strings"
)

type ScmService struct {
}

func NewScmService() *ScmService {
	return &ScmService{}
}

func (s *ScmService) GetTestScript(build *commDomain.Build) _domain.RpcResp {
	if build.ScmAddress != "" {
		CheckoutCodes(build)
	} else if strings.Index(build.ScriptUrl, "http://") == 0 {
		DownloadCodes(build)
	} else {
		build.ProjectDir = _fileUtils.AddPathSepIfNeeded(build.ScriptUrl)
	}

	result := _domain.RpcResp{}
	result.Success("")
	return result
}

func CheckoutCodes(task *commDomain.Build) {
	url := task.ScmAddress
	userName := task.ScmAccount
	password := task.ScmPassword

	projectDir := task.WorkDir + _gitUtils.GetGitProjectName(url) + _const.PthSep

	_fileUtils.MkDirIfNeeded(projectDir)

	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	if userName != "" {
		options.Auth = &http.BasicAuth{
			Username: userName,
			Password: password,
		}
	}
	_, err := git.PlainClone(projectDir, false, &options)
	if err != nil {
		return
	}

	task.ProjectDir = projectDir
}

func DownloadCodes(build *commDomain.Build) {
	zipPath := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.ScriptUrl)
	_fileUtils.Download(build.ScriptUrl, zipPath)

	scriptFolder := _fileUtils.GetZipSingleDir(zipPath)
	if scriptFolder != "" { // single dir in zip
		build.ProjectDir = build.WorkDir + scriptFolder
		archiver.Unarchive(zipPath, build.WorkDir)
	} else { // more then one dir, unzip to a folder
		fileNameWithoutExt := _fileUtils.GetFileNameWithoutExt(zipPath)
		build.ProjectDir = build.WorkDir + fileNameWithoutExt + _const.PthSep
		archiver.Unarchive(zipPath, build.ProjectDir)
	}
}
