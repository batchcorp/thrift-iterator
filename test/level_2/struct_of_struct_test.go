package test

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/batchcorp/thrift-iterator/general"
	"github.com/batchcorp/thrift-iterator/protocol"
	"github.com/batchcorp/thrift-iterator/test"
	"github.com/batchcorp/thrift-iterator/test/level_2/struct_of_struct_test"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_skip_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_unmarshal_general_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val general.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): "abc",
		}, val[protocol.FieldId(1)])
	}
}

func Test_unmarshal_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val struct_of_struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_of_struct_test.TestObject{
			struct_of_struct_test.EmbeddedObject{"abc"},
		}, val)
	}
}

func Test_marshal_general_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		obj := general.Struct{
			protocol.FieldId(1): general.Struct{
				protocol.FieldId(1): "abc",
			},
		}

		output, err := c.Marshal(obj)
		should.NoError(err)
		output1, err := c.Marshal(&obj)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): "abc",
		}, val[protocol.FieldId(1)])
	}
}

func Test_marshal_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		obj := struct_of_struct_test.TestObject{
			struct_of_struct_test.EmbeddedObject{"abc"},
		}

		output, err := c.Marshal(obj)
		should.NoError(err)
		output1, err := c.Marshal(&obj)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): "abc",
		}, val[protocol.FieldId(1)])
	}
}
