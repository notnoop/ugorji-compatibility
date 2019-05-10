package main

import (
	"encoding/json"
	"reflect"
	"testing"

	hcodec "github.com/hashicorp/go-msgpack/codec"
	"github.com/kr/pretty"
	"github.com/ugorji/go/codec"
)

const cases = 50

const skipHashicorpMismatch = true

var (
	uMsgpack = &codec.MsgpackHandle{}
	hMsgpack = &hcodec.MsgpackHandle{}
)

func init() {
	uMsgpack = &codec.MsgpackHandle{}
	uMsgpack.RawToString = true
	uMsgpack.BasicHandle.TimeNotBuiltin = true
	uMsgpack.BasicHandle.TypeInfos = codec.NewTypeInfos([]string{"codec"})
	uMsgpack.MapType = reflect.TypeOf(map[string]interface{}(nil))

	hMsgpack.RawToString = true
	hMsgpack.MapType = reflect.TypeOf(map[string]interface{}(nil))
}

func logStruct(t *testing.T, data interface{}) {
	if b, err := json.Marshal(data); err == nil {
		t.Logf("tring out: %v", string(b))
	} else {
		t.Logf("trying out: %# v", pretty.Formatter(data))
	}
}
