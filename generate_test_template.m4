define(`_preamble', `//+build all generated $2

package main

import (
"bytes"
"fmt"
"testing"
"time"

hcodec "github.com/hashicorp/go-msgpack/codec"
"github.com/stretchr/testify/require"
"github.com/ugorji/go/codec"

$1

)

var _ *time.Time = nil
')dnl
define(`_test_type', `
func TestSerialization_$2_$1(t *testing.T) {
s := true
    for i := 0; s && i < cases; i++ {
    s = t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
    var data $1
        fakeStruct(&data)

logStruct(t, data)
var buf bytes.Buffer
encoder := hcodec.NewEncoder(&buf, hMsgpack)
require.NoError(t, encoder.Encode(data))

var oroundtripped $1
decoder := hcodec.NewDecoder(&buf, hMsgpack)
require.NoError(t, decoder.Decode(&oroundtripped))

if checkHashicorpMismatch {
assertEqual(t, data, oroundtripped, "hashicorp did not preserve data")
}


        r1 := t.Run(fmt.Sprintf("from ugorji to ugorji: %v", i), func(t *testing.T) {
        logStruct(t, data)
            var buf bytes.Buffer
            encoder := codec.NewEncoder(&buf, uMsgpack)
            require.NoError(t, encoder.Encode(data))

      var dest $1
            decoder := codec.NewDecoder(&buf, uMsgpack)
            require.NoError(t, decoder.Decode(&dest))

requireEqualsToEither(t, dest, data, oroundtripped)
        })

        r2 := t.Run(fmt.Sprintf("from hashicorp to ugorji: %v", i), func(t *testing.T) {
        logStruct(t, data)
            var buf bytes.Buffer
            encoder := hcodec.NewEncoder(&buf, hMsgpack)
            require.NoError(t, encoder.Encode(data))

      var dest $1
            decoder := codec.NewDecoder(&buf, uMsgpack)
            require.NoError(t, decoder.Decode(&dest))

requireEqualsToEither(t, dest, data, oroundtripped)
        })

        r3 := t.Run(fmt.Sprintf("from ugorji to hashicorp: %v", i), func(t *testing.T) {
        logStruct(t, data)
            var buf bytes.Buffer
            encoder := codec.NewEncoder(&buf, uMsgpack)
            require.NoError(t, encoder.Encode(data))

      var dest $1
            decoder := hcodec.NewDecoder(&buf, hMsgpack)
            require.NoError(t, decoder.Decode(&dest))

requireEqualsToEither(t, dest, data, oroundtripped)
        })

        if !(r1 && r2 && r3) {
            t.FailNow()
        }
        })
    }
}')dnl
