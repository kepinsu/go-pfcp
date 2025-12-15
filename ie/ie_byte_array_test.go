// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestByteArrayIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     []byte
		decoderFunc func(*ie.IE) ([]byte, error)
	}{
		{
			description: "CPFunctionFeatures",
			structured:  ie.NewCPFunctionFeatures(0x3f),
			decoded:     []byte{0x3f},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.CPFunctionFeatures() },
		}, {
			description: "CPFunctionFeatures/2bytes",
			structured:  ie.NewCPFunctionFeatures(0x3f, 0x01),
			decoded:     []byte{0x3f, 0x01},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.CPFunctionFeatures() },
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.decoderFunc(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.decoded); diff != "" {
				t.Error(diff)
			}
		})
	}
}
