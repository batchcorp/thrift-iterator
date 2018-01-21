package thrifter

import (
	"io"
	"github.com/thrift-iterator/go/spi"
)

type Protocol int

var ProtocolBinary Protocol = 1
var ProtocolCompact Protocol = 2

type Decoder interface {
	Decode(obj interface{}) error
	Reset(reader io.Reader, buf []byte)
}

type Encoder interface {
	Encode(obj interface{}) error
	Reset(writer io.Writer)
	Buffer() []byte
}

type Config struct {
	Protocol       Protocol
	IsFramed       bool
	DynamicCodegen bool
}

type API interface {
	// NewStream is low level streaming api
	NewStream(writer io.Writer, buf []byte) spi.Stream
	// NewIterator is low level streaming api
	NewIterator(reader io.Reader, buf []byte) spi.Iterator
	Unmarshal(buf []byte, obj interface{}) error
	Marshal(obj interface{}) ([]byte, error)
	NewDecoder(reader io.Reader, buf []byte) Decoder
	NewEncoder(writer io.Writer) Encoder
	// WillDecodeFromBuffer should only be used in generic.Declare
	WillDecodeFromBuffer(sample ...interface{})
	// WillDecodeFromReader should only be used in generic.Declare
	WillDecodeFromReader(sample ...interface{})
	// WillEncode should only be used in generic.Declare
	WillEncode(sample ...interface{})
}

var DefaultConfig = Config{Protocol: ProtocolBinary, IsFramed: true, DynamicCodegen: true}.Froze()

func NewStream(writer io.Writer, buf []byte) spi.Stream {
	return DefaultConfig.NewStream(writer, buf)
}

func NewIterator(reader io.Reader, buf []byte) spi.Iterator {
	return DefaultConfig.NewIterator(reader, buf)
}

func Unmarshal(buf []byte, obj interface{}) error {
	return DefaultConfig.Unmarshal(buf, obj)
}

func Marshal(obj interface{}) ([]byte, error) {
	return DefaultConfig.Marshal(obj)
}

func NewDecoder(reader io.Reader, buf []byte) Decoder {
	return DefaultConfig.NewDecoder(reader, buf)
}

func NewEncoder(writer io.Writer) Encoder {
	return DefaultConfig.NewEncoder(writer)
}
