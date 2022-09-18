package test

import (
	"github.com/v2pro/wombat/generic"
	"github.com/batchcorp/thrift-iterator"
	"github.com/batchcorp/thrift-iterator/test/api/binding_test"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
}.Froze()

//go:generate go install github.com/batchcorp/thrift-iterator/cmd/thrifter
//go:generate $GOPATH/bin/thrifter -pkg github.com/batchcorp/thrift-iterator/test/api
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*binding_test.TestObject)(nil),
		)
	})
}
