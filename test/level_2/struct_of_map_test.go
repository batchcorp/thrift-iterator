package test

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/batchcorp/thrift-iterator/general"
	"github.com/batchcorp/thrift-iterator/protocol"
	"github.com/batchcorp/thrift-iterator/test"
	"github.com/batchcorp/thrift-iterator/test/level_2/struct_of_map_test"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_skip_struct_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.MAP, 1)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_unmarshal_general_struct_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.MAP, 1)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val general.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Map{
			int32(2): int64(2),
		}, val[protocol.FieldId(1)])
	}
}

func Test_unmarshal_struct_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.MAP, 1)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val struct_of_map_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_of_map_test.TestObject{
			map[int32]int64{2: 2},
		}, val)
	}
}

func Test_marshal_general_struct_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		m := general.Struct{
			protocol.FieldId(1): general.Map{
				int32(2): int64(2),
			},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Map{
			int32(2): int64(2),
		}, val[protocol.FieldId(1)])
	}
}

func Test_marshal_struct_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		m := struct_of_map_test.TestObject{
			map[int32]int64{2: 2},
		}

		output, err := c.Marshal(m)
		should.NoError(err)
		output1, err := c.Marshal(&m)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Map{
			int32(2): int64(2),
		}, val[protocol.FieldId(1)])
	}
}
