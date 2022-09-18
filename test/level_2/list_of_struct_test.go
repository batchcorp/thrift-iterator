package test

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/batchcorp/thrift-iterator/general"
	"github.com/batchcorp/thrift-iterator/protocol"
	"github.com/batchcorp/thrift-iterator/test"
	"github.com/batchcorp/thrift-iterator/test/level_2/list_of_struct_test"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_skip_list_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.STRUCT, 2)
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteListEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipList(nil))
	}
}

func Test_unmarshal_general_list_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.STRUCT, 2)
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteListEnd()
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[0])
	}
}

func Test_unmarshal_list_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.STRUCT, 2)
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		proto.WriteListEnd()
		var val []list_of_struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([]list_of_struct_test.TestObject{
			{1024}, {1024},
		}, val)
	}
}

func Test_marshal_general_list_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		lst := general.List{
			general.Struct{
				protocol.FieldId(1): int64(1024),
			},
			general.Struct{
				protocol.FieldId(1): int64(1024),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[0])
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[1])
	}
}

func Test_marshal_list_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		lst := []list_of_struct_test.TestObject{
			{1024}, {1024},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.Equal(output, output1)
		should.NoError(err)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[0])
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[1])
	}
}
