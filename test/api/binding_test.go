package test

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/batchcorp/thrift-iterator/test/api/binding_test"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_binding(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	transport := thrift.NewTFramedTransport(buf)
	proto := thrift.NewTBinaryProtocol(transport, true, true)
	proto.WriteStructBegin("hello")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1024)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	transport.Flush()
	var val binding_test.TestObject
	should.NoError(api.Unmarshal(buf.Bytes()[4:], &val))
	should.Equal(int64(1024), val.Field1)
}
