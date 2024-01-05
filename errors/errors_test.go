package errors

import "testing"

func TestNew(t *testing.T) {
	err := New(500, "lihuizhegegb")
	t.Error(err)
}

func TestNewBusinessError(t *testing.T) {
	err := NewBusinessError(501, "lowzhixiangailuozi")
	t.Error(err)
}

func TestFormatGRPCError(t *testing.T) {
	err := FormatGRPCError(New(502, "rpc error: code = 500 desc = 500-hanjunlilaobnvren"))
	t.Error(err)
}
