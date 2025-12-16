// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestOffendingIE(t *testing.T) {
	structured := ie.NewOffendingIE(ie.Cause)
	decoded := ie.Cause

	got, err := structured.OffendingIE()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, decoded); diff != "" {
		t.Error(diff)
	}
}

func TestApplyActionIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     *ie.ApplyActionFields
		decoderFunc func(*ie.IE) (*ie.ApplyActionFields, error)
	}{
		{
			description: "ApplyAction/pre-16.3.0",
			structured:  ie.NewApplyAction(0x04), // Flag BUFF is set
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat",
			structured:  ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0",
			structured:  ie.NewApplyAction(0x04, 0x02), //Flags BUFF and BDPN are set
			decoded:     &ie.ApplyActionFields{Buff: true, Bdpn: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/pre-16.3.0/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04), // Flag BUFF is set
			),
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			),
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x02), // Flags BUFF and BDPN are set
			),
			decoded:     &ie.ApplyActionFields{Buff: true, Bdpn: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/pre-16.3.0/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04), // Flag BUFF is set
			),
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			),
			decoded:     &ie.ApplyActionFields{Buff: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x02), // Flags BUFF and BDPN are set
			),
			decoded:     &ie.ApplyActionFields{Buff: true, Bdpn: true},
			decoderFunc: func(i *ie.IE) (*ie.ApplyActionFields, error) { return i.ApplyAction() },
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

func TestGateStatusIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     ie.Gate
		decoderFunc func(*ie.IE) (ie.Gate, error)
	}{

		{
			description: "GateStatus/OpenOpen",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenClosed",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusClosed),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedOpen",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusOpen),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedClosed",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusClosed),
			decoded:     5,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenOpen/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenOpen/UpdateQER",
			structured: ie.NewUpdateQER(
				ie.NewQERID(0xffffffff),
				ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (ie.Gate, error) { return i.GateStatus() },
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

func TestDSCPToPPIMappingInformationIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     *ie.DSCPToPPIMappingInformationFields
		decoderFunc func(*ie.IE) (*ie.DSCPToPPIMappingInformationFields, error)
	}{

		{
			description: "DSCPToPPIMappingInformationFields",
			structured:  ie.NewDSCPToPPIMappingInformation(1),
			decoded: &ie.DSCPToPPIMappingInformationFields{
				PPIValue: 1,
			},
			decoderFunc: func(i *ie.IE) (*ie.DSCPToPPIMappingInformationFields, error) { return i.DSCPToPPIMappingInformation() },
		}, {
			description: "DSCPToPPIMappingInformationFields/DSCPToPPIControlInformationFields",
			structured: ie.NewDSCPToPPIControlInformation(
				ie.NewDSCPToPPIMappingInformation(1, 2, 3),
			),
			decoded: &ie.DSCPToPPIMappingInformationFields{
				PPIValue: 1,
				DSCP:     []uint8{2, 3},
			},
			decoderFunc: func(i *ie.IE) (*ie.DSCPToPPIMappingInformationFields, error) { return i.DSCPToPPIMappingInformation() },
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
