//+build custom

package main

import (
	"bytes"
	"testing"

	hcodec "github.com/hashicorp/go-msgpack/codec"
	"github.com/stretchr/testify/require"
	"github.com/ugorji/go/codec"
)

func TestCustom_EmptyArray_InHashicorp(t *testing.T) {
	h := &hcodec.MsgpackHandle{}

	type MyStruct struct {
		A []byte
		B []byte

		C []string
		D []string
	}

	data := MyStruct{
		A: []byte{},
		B: nil,
		C: []string{},
		D: nil,
	}

	var buf bytes.Buffer
	encoder := hcodec.NewEncoder(&buf, h)
	require.NoError(t, encoder.Encode(data))

	var dest MyStruct
	decoder := hcodec.NewDecoder(&buf, h)
	require.NoError(t, decoder.Decode(&dest))

	require.Equal(t, data, dest)
}

func TestCustom_EmptyArray_Ugorji(t *testing.T) {
	h := &codec.MsgpackHandle{}

	type MyStruct struct {
		A []byte
		B []byte

		C []string
		D []string
	}

	data := MyStruct{
		A: []byte{},
		B: nil,
		C: []string{},
		D: nil,
	}

	var buf bytes.Buffer
	encoder := codec.NewEncoder(&buf, h)
	require.NoError(t, encoder.Encode(data))

	var dest MyStruct
	decoder := codec.NewDecoder(&buf, h)
	require.NoError(t, decoder.Decode(&dest))

	require.Equal(t, data, dest)
}
