package errors

import (
	"fmt"
	"regexp"

	"github.com/Leo225/sdk/convert"
)

const (
	grpcErrPattern = "rpc error: code = (?P<rpc_code>.+) desc = (?P<code>[0-9]+)-(?P<msg>.+)"
)

type BusinessError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("%d-%s", e.Code, e.Msg)
}

func New(code int, msg string) error {
	return &BusinessError{
		Code: code,
		Msg:  msg,
	}
}

func NewBusinessError(code int, msg string) *BusinessError {
	return &BusinessError{
		Code: code,
		Msg:  msg,
	}
}

func FormatGRPCError(err error) (e *BusinessError) {
	if err == nil {
		e = nil
		return
	}

	errMap := convertMap(err, grpcErrPattern)
	_, rpcCodeOk := errMap["rpc_code"]
	if !rpcCodeOk {
		return
	}

	code, codeOk := errMap["code"]
	if !codeOk {
		return
	}

	msg, ok := errMap["msg"]
	if !ok {
		return
	}

	e = &BusinessError{
		Code: convert.ToInt(code),
		Msg:  msg,
	}
	return
}

func convertMap(err error, pattern string) map[string]string {
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(err.Error())
	paramsMap := make(map[string]string)

	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
