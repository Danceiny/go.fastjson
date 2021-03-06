package fastjson

import (
    utils "github.com/Danceiny/go.utils"
    log "github.com/sirupsen/logrus"
    "reflect"
)

func ParseObjectB(bytes []byte) *JSONObject {
    var o JSONObject
    if err := json.Unmarshal(bytes, &o); err != nil {
        return nil
    }
    return &o
}

// o, &struct{}
func ParseObjectBT(bytes []byte, o interface{}) {
    if err := json.Unmarshal(bytes, o); err != nil {
        log.Warning(err)
    }
}

// o, &struct{}
func ParseObjectT(s string, o interface{}) {
    if err := json.UnmarshalFromString(s, o); err != nil {
        log.Warning(err)
    }
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
