package domain

import _const "github.com/easysoft/zagent/pkg/const"

type ValidRequest struct {
	Method _const.ValidMethod `json:"method"`
	Value  string             `json:"value"`

	Id   int    `json:"id"`
	Type string `json:"type"`
}

type ValidResp struct {
	Pass bool `json:"pass"`
	Msg  bool `json:"msg"`
}
