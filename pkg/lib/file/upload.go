package _fileUtils

import (
	"bytes"
	_i118Utils "github.com/easysoft/zagent/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(url string, files []string, extraParams map[string]string) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	for _, file := range files {
		fw, _ := bodyWriter.CreateFormFile("file", file)
		f, _ := os.Open(file)
		defer f.Close()
		io.Copy(fw, f)
	}

	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuffer)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_upload_files", err.Error()))
	}
	if resp == nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_upload_files", "resp is nil"))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_logUtils.Error(_i118Utils.Sprintf("fail_to_parse_upload_response", err.Error()))
	}

	_logUtils.Info(_i118Utils.Sprintf("upload_status", resp.Status, string(respBody)))
}
