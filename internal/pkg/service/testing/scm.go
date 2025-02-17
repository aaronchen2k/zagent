package agentTestingService

import (
	commDomain "github.com/easysoft/zagent/internal/pkg/domain"
	_const "github.com/easysoft/zagent/pkg/const"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_gitUtils "github.com/easysoft/zagent/pkg/lib/git"
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

func (s *ScmService) GetTestScript(build *commDomain.Build) (err error) {
	if build.ScmAddress != "" {
		err = CheckoutCodes(build)
	} else if strings.Index(build.ScriptUrl, "http://") == 0 {
		err = DownloadCodes(build)
	}
	//else {
	//	build.ProjectDir = _fileUtils.AddPathSepIfNeeded(build.ScriptUrl)
	//}

	return
}

func CheckoutCodes(build *commDomain.Build) (err error) {
	url := build.ScmAddress
	userName := build.ScmAccount
	password := build.ScmPassword

	projectDir := build.WorkDir + _gitUtils.GetGitProjectName(url) + _const.PthSep
	build.ProjectDir = projectDir
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
	_, err = git.PlainClone(projectDir, false, &options)
	if err != nil {
		return
	}

	return
}

func DownloadCodes(build *commDomain.Build) (err error) {
	zipPath := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.ScriptUrl)
	err = _fileUtils.Download(build.ScriptUrl, zipPath)

	if err != nil {
		return
	}

	scriptFolder := _fileUtils.GetZipSingleDir(zipPath)
	if scriptFolder != "" { // single dir in zip
		build.ProjectDir = build.WorkDir + scriptFolder
		err = archiver.Unarchive(zipPath, build.WorkDir)
	} else { // more then one dir, unzip to a folder
		fileNameWithoutExt := _fileUtils.GetFileNameWithoutExt(zipPath)
		build.ProjectDir = build.WorkDir + fileNameWithoutExt + _const.PthSep
		err = archiver.Unarchive(zipPath, build.ProjectDir)
	}

	return
}
