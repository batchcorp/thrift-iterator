package test

import (
	"testing"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/test"
	"github.com/thrift-iterator/go/test/level_1/struct_test"
)

func Test_decode_struct_by_iterator(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
		should.NoError(iter.Error())
		should.True(called)
	}
}

func Test_decode_struct_with_bool_by_iterator(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.BOOL, 1)
		proto.WriteBool(true)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeBool, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(true, iter.ReadBool())
		}
		should.True(called)
	}
}

func Test_encode_struct_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStructHeader()
		stream.WriteStructField(protocol.TypeI64, protocol.FieldId(1))
		stream.WriteInt64(1024)
		stream.WriteStructFieldStop()
		iter := c.CreateIterator(stream.Buffer())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
	}
}

func Test_encode_struct_with_bool_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStructHeader()
		stream.WriteStructField(protocol.TypeBool, protocol.FieldId(1))
		stream.WriteBool(true)
		stream.WriteStructFieldStop()
		iter := c.CreateIterator(stream.Buffer())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeBool, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(true, iter.ReadBool())
		}
		should.True(called)
	}
}

func Test_decode_struct_as_object(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		obj := iter.ReadStruct()
		should.Equal(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): int64(1024),
		}, obj)
	}
}

func Test_unmarshal_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_test.TestObject{1024}, val)
	}
}

func Test_encode_struct_from_object(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStruct(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): int64(1024),
		})
		iter := c.CreateIterator(stream.Buffer())
		obj := iter.ReadStruct()
		should.Equal(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): int64(1024),
		}, obj)
	}
}

func Test_skip_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_marshal_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(struct_test.TestObject{1024})
		should.NoError(err)
		iter := c.CreateIterator(output)
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
		should.True(called)
	}
}
