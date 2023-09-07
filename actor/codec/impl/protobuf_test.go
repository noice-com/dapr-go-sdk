package impl

import (
	"testing"

	sample "github.com/dapr/go-sdk/actor/codec/impl/protosample"
	"google.golang.org/protobuf/proto"
)

func TestProtobufMarshal(t *testing.T) {
	inObj := &sample.Sample{
		IntValue: 123,
		StrValue: "abc",
	}

	c := &ProtobufCodec{}

	bytes, err := c.Marshal(inObj)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	var outObj *sample.Sample

	err = c.Unmarshal(bytes, &outObj)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if !proto.Equal(inObj, outObj) {
		t.Error("input and output does not match")
	}
}
