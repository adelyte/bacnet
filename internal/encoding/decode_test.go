package encoding

import (
	"bacnet/internal/types"
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestDecodeTag(t *testing.T) {
	ttc := []struct {
		data     string //hex string
		expected tag
	}{
		{
			data: "09",
			expected: tag{
				ID:      0,
				Value:   1,
				Context: true,
			},
		},
		{
			data: "1a",
			expected: tag{
				ID:      1,
				Value:   2,
				Context: true,
			},
		},
		{
			data: "c4",
			expected: tag{
				ID:    12,
				Value: 4,
			},
		},
		{
			data: "22",
			expected: tag{
				ID:    2,
				Value: 2,
			},
		},
		{
			data: "91",
			expected: tag{
				ID:    9,
				Value: 1,
			},
		},
	}
	for _, tc := range ttc {
		t.Run(fmt.Sprintf("Tag %s", tc.data), func(t *testing.T) {
			is := is.New(t)
			b, err := hex.DecodeString(tc.data)
			is.NoErr(err)
			buf := bytes.NewBuffer(b)
			length, tag, err := decodeTag(buf)
			is.NoErr(err)
			is.Equal(length, len(b))
			if !tagEqual(tag, tc.expected) {
				t.Errorf("Tag differ: expected %+v got %+v", tc.expected, tag)
			}
		})
	}
}

func tagEqual(t1, t2 tag) bool {
	return t1.ID == t2.ID &&
		t1.Context == t2.Context &&
		t1.Value == t2.Value &&
		t1.Opening == t2.Opening &&
		t1.Closing == t2.Closing
}

func TestDecodeAppData(t *testing.T) {
	ttc := []struct {
		data     string //hex string
		from     interface{}
		expected interface{}
	}{
		{
			data: "c4020075e9",
			from: types.ObjectID{},
			expected: types.ObjectID{
				Type:     8,
				Instance: 30185,
			},
		},
		{
			data:     "2205c4",
			from:     uint32(0),
			expected: uint32(1476),
		},
		{
			data:     "9100",
			from:     types.SegmentationSupport(0),
			expected: types.SegmentationSupportBoth,
		},
		{
			data:     "22016c",
			from:     uint32(0),
			expected: uint32(364),
		},
	}
	for _, tc := range ttc {
		t.Run(fmt.Sprintf("AppData %s", tc.data), func(t *testing.T) {
			is := is.New(t)
			b, err := hex.DecodeString(tc.data)
			is.NoErr(err)
			buf := bytes.NewBuffer(b)
			switch tc.from.(type) {
			case types.ObjectID:
				x := types.ObjectID{}
				err = DecodeAppData(buf, &x)
				is.NoErr(err)
				is.Equal(x, tc.expected)
			case uint32:
				var x uint32
				err = DecodeAppData(buf, &x)
				is.NoErr(err)
				is.Equal(x, tc.expected)
			case types.SegmentationSupport:
				var x types.SegmentationSupport
				err = DecodeAppData(buf, &x)
				is.NoErr(err)
				is.Equal(x, tc.expected)
			default:
				t.Errorf("Invalid from type %T", tc.from)
			}

		})
	}
}
