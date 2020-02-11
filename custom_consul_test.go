package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ugorji/go/codec"
)

type Entries struct {
	Resolvers map[string]*Entry
}

type Entry struct {
	Kind string
}

func TestMySerializing_DiscoveryChainConfigEntries(t *testing.T) {
	c := Entries{
		Resolvers: map[string]*Entry{
			"zwg": (*Entry)(nil),
		},
	}

	t.Logf("SERIALIZAING %#+v\n", c)

	var buf bytes.Buffer
	encoder := codec.NewEncoder(&buf, uMsgpack)
	require.NoError(t, encoder.Encode(&c))

	var f Entries

	decoder := codec.NewDecoder(&buf, uMsgpack)
	require.NoError(t, decoder.Decode(&f))

	require.Equal(t, c, f)

}
