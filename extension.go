package fastjson

import (
    "github.com/json-iterator/go"
    "unsafe"
)

type ExtensionEnum uint8

const (
    _ ExtensionEnum = iota
    CUSTOM_TIME_EXT
    METHOD_VALUE_EXT
)

var (
    FastJson = jsoniter.ConfigCompatibleWithStandardLibrary
)

func RegisterExt(extension ExtensionEnum) {
    switch extension {
    case CUSTOM_TIME_EXT:
        FastJson.RegisterExtension(&CustomTimeExtension{})
        break
    case METHOD_VALUE_EXT:
        FastJson.RegisterExtension(&MethodValueExtension{})
        break
    default:
        return
    }
}

func UnregisterExt(extension ExtensionEnum) {
    // todo: json-iterator没有实现且没有暴露[]extension变量，可能需要自己fork来搞一下
}

type funcDecoder struct {
    fun jsoniter.DecoderFunc
}

func (decoder *funcDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    decoder.fun(ptr, iter)
}

type funcEncoder struct {
    fun         jsoniter.EncoderFunc
    isEmptyFunc func(ptr unsafe.Pointer) bool
}

func (encoder *funcEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    encoder.fun(ptr, stream)
}

func (encoder *funcEncoder) IsEmpty(ptr unsafe.Pointer) bool {
    if encoder.isEmptyFunc == nil {
        return false
    }
    return encoder.isEmptyFunc(ptr)
}
