package fastjson

import (
    utils "github.com/Danceiny/go.utils"
    log "github.com/sirupsen/logrus"
    "reflect"
)

var json = FastJson

func init() {
    // acquiescently, we registered custom time extension
    // and you can register some new extensions (for our implemented extensions, just call `RegisterExt`,
    // or else you implement your extension and register it to `FastJson` which is exposed）
    RegisterExt(CUSTOM_TIME_EXT)
}

type JSON interface {
    ToJSON() []byte
    ToJSONString() string
}

func ToJSON(v interface{}) []byte {
    if bytes, err := json.Marshal(v); err != nil {
        log.Warning(err)
        return nil
    } else {
        return bytes
    }
}

func ToJSONString(v interface{}) string {
    if s, err := json.MarshalToString(v); err != nil {
        log.Warning(err)
        return ""
    } else {
        return s
    }
}

func ParseObjectB(bytes []byte) *JSONObject {
    var o JSONObject
    if err := json.Unmarshal(bytes, &o); err != nil {
        log.Warning(err)
        return nil
    }
    return &o
}

func ParseObject(s string) *JSONObject {
    var o JSONObject
    if err := json.UnmarshalFromString(s, &o); err != nil {
        log.Warning(err)
        return nil
    } else {
        return &o
    }
}

func ParseArray(s string) *JSONArray {
    var o JSONArray
    if err := json.UnmarshalFromString(s, &o.arr); err != nil {
        log.Warning(err)
        return nil
    }
    return postParseArray(&o)
}

func ParseArrayB(s []byte) *JSONArray {
    var o JSONArray
    if err := json.Unmarshal(s, &o.arr); err != nil {
        log.Warning(err)
        return nil
    }
    return postParseArray(&o)
}

// t, only support primitive types
func ParseArrayT(s string, t reflect.Kind) interface{} {
    var o JSONArray
    if err := json.UnmarshalFromString(s, &o.arr); err != nil {
        log.Warning(err)
        return nil
    }
    return postParseArrayT(&o, t)
}

func ParseArrayBT(s []byte, t reflect.Kind) interface{} {
    var o JSONArray
    if err := json.Unmarshal(s, &o.arr); err != nil {
        log.Warning(err)
        return nil
    }
    return postParseArrayT(&o, t)
}

func postParseArrayT(o *JSONArray, t reflect.Kind) interface{} {
    o.size = len(o.arr)
    utils.CastPrimitiveSliceInplace(&o.arr, t)
    return utils.CastPrimitiveSlice(o.arr, t)
}

func postParseArray(o *JSONArray) *JSONArray {
    o.size = len(o.arr)
    return o
}
