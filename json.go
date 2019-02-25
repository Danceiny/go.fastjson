package fastjson

import (
    log "github.com/sirupsen/logrus"
)

var json = FastJson

func init() {
    log.SetLevel(log.ErrorLevel)
    // acquiescently, we registered custom time extension
    // and you can register some new extensions (for our implemented extensions, just call `RegisterExt`,
    // or else you implement your extension and register it to `FastJson` which is exposed）
    RegisterExt(CUSTOM_TIME_EXT)
}

type JSON interface {
    ToJSON() []byte
    ToJSONString() string
}

type JSONSerializerFeature struct {
    MethodValue bool
}
