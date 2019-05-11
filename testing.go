package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	hcodec "github.com/hashicorp/go-msgpack/codec"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ugorji/go/codec"
)

const cases = 5000

const checkHashicorpMismatch = false

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

func jsonStr(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func logStruct(t *testing.T, data interface{}) {
	if b, err := json.Marshal(data); err == nil {
		t.Logf("tring out: %v", string(b))
	} else {
		t.Logf("trying out: %# v", pretty.Formatter(data))
	}
}

func assertEqual(t *testing.T, expected, actual interface{}, args ...interface{}) {
	spew.Config.Sprintf("%v", expected)
	spew.Config.Sprintf("%v", actual)

	assert.Equal(t, expected, actual, args...)

}

func requireEqual(t *testing.T, expected, actual interface{}) {
	spew.Config.Sprintf("%v", expected)
	spew.Config.Sprintf("%v", actual)

	require.Equal(t, expected, actual)

}

// assertEqualsToEither assets that actual value matches one of the expected values
func requireEqualsToEither(t *testing.T, actual, exp1, exp2 interface{}) {
	spew.Config.Sprintf("%v", actual)
	spew.Config.Sprintf("%v", exp1)
	spew.Config.Sprintf("%v", exp2)

	c1 := reflect.DeepEqual(exp1, actual)
	c2 := reflect.DeepEqual(exp2, actual)

	if c1 || c2 {
		return
	}

	assert.EqualValues(t, exp1, actual, "doesnt match first value")
	assert.EqualValues(t, exp2, actual, "doesnt match second value")
	t.FailNow()
}
