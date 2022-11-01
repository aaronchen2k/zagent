package _fileUtils

import (
	"fmt"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Download(url string, dst string) (err error) {
	fmt.Printf("DownloadToFile From: %s to %s.\n", url, dst)

	MkDirIfNeeded(filepath.Dir(dst))

	var data []byte
	data, err = HTTPDownload(url)
	if err == nil {
		_logUtils.Info(_i118Utils.Sprintf("file_downloaded", url))

		err = WriteDownloadFile(dst, data)
		if err == nil {
			_logUtils.Info(_i118Utils.Sprintf("file_download_saved", url, dst))
		}
	}

	return
}

func HTTPDownload(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	return d, err
}

func WriteDownloadFile(dst string, d []byte) error {
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		_logUtils.Error(err.Error())
	}
	return err
}
