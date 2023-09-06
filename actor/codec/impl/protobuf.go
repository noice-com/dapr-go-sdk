package impl

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/dapr/go-sdk/actor/codec"
	"github.com/dapr/go-sdk/actor/codec/constant"
	"google.golang.org/protobuf/proto"
)

var (
	ErrNotProtoMessage = errors.New("not a proto.Message")
)

func init() {
	codec.SetActorCodec(constant.ProtobufSerializerType, func() codec.Codec {
		return &ProtobufCodec{}
	})
}

type ProtobufCodec struct{}

func (c *ProtobufCodec) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("%w, got %T", ErrNotProtoMessage, v)
	}

	return proto.Marshal(m)
}

func (c *ProtobufCodec) Unmarshal(data []byte, v interface{}) error {
	// Get the reflection value of the passed pointer.
	vValue := reflect.ValueOf(v)

	// Check if the passed pointer is a pointer.
	if vValue.Kind() != reflect.Ptr {
		return fmt.Errorf("ptr must be a pointer")
	}

	// Get the type of the underlying element that ptr points to.
	targetType := vValue.Elem().Type()

	newObjValue := reflect.New(targetType.Elem())

	fmt.Println("NEW OBJ VALUE:", newObjValue.Type().String())

	v = newObjValue.Interface()

	m, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("%w: got %T %#v", ErrNotProtoMessage, v, v)
	}

	// Assign the newBar to the pointer.
	vValue.Elem().Set(newObjValue)

	return proto.Unmarshal(data, m)
}
