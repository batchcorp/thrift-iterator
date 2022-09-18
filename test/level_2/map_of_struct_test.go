package test

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/batchcorp/thrift-iterator/general"
	"github.com/batchcorp/thrift-iterator/protocol"
	"github.com/batchcorp/thrift-iterator/test"
	"github.com/batchcorp/thrift-iterator/test/level_2/map_of_struct_test"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_skip_map_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.STRUCT, 1)
		proto.WriteI64(1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipMap(nil))
	}
}

func Test_unmarshal_general_map_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.STRUCT, 1)
		proto.WriteI64(1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteMapEnd()
		var val general.Map
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[int64(1)])
	}
}

func Test_unmarshal_map_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.STRUCT, 1)
		proto.WriteI64(1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.I64, 1)
		proto.WriteI64(1024)
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteMapEnd()
		var val map[int64]map_of_struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(map[int64]map_of_struct_test.TestObject{
			1: {1024},
		}, val)
	}
}

func Test_marshal_general_map_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		m := general.Map{
			int64(1): general.Struct{
				protocol.FieldId(1): int64(1024),
			},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Map
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[int64(1)])
	}
}

func Test_marshal_map_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		m := map[int64]map_of_struct_test.TestObject{
			1: {1024},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Map
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val[int64(1)])
	}
}
