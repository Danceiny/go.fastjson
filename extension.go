package fastjson

import (
    "github.com/json-iterator/go"
)

type ExtensionEnum uint8

const (
    FAST_JSON_EXTENSION = iota
    CUSTOM_TIME_EXT
)

var (
    FastJson = jsoniter.ConfigCompatibleWithStandardLibrary
)

func RegisterExt(extension ExtensionEnum) {
    switch extension {
    case CUSTOM_TIME_EXT:
        FastJson.RegisterExtension(&CustomTimeExtension{})
        break
    default:
        return
    }
}

func UnregisterExt(extension ExtensionEnum) {
    // todo: json-iterator没有实现且没有暴露[]extension变量，可能需要自己fork来搞一下
}
