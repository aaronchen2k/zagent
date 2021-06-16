package main

import (
	"bytes"
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	gateway "github.com/rpcx-ecosystem/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/log"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	cc := &codec.MsgpackCodec{}

	args := &_domain.ArithArgs{
		A: 1,
		B: 2,
	}

	data, _ := cc.Encode(args)
	// request
	url := fmt.Sprintf("http://127.0.0.1:%d/", _const.RpcPort)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		_logUtils.Errorf("failed to create request, error: %s", err.Error())
		return
	}

	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "arith")
	h.Set(gateway.XServiceMethod, "Add")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("failed to call: ", err)
	}
	defer res.Body.Close()
	// 获取结果
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("failed to read response: ", err)
	}
	// 解码
	reply := &_domain.ArithReply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Errorf("failed to decode reply: ", err)
	}
	log.Infof("%d + %d = %d", args.A, args.B, reply.C)
}
