package fastjson

import (
    log "github.com/sirupsen/logrus"
    "reflect"
    "strings"
)

func ToJSONB(v interface{}) []byte {
    if bytes, err := json.Marshal(v); err != nil {
        log.Warning(err)
        return nil
    } else {
        return bytes
    }
}

func ToJSONBv2(v interface{}, serializer *JSONSerializerFeature) []byte {
    if serializer.MethodValue {
        var t = reflect.TypeOf(v)
        var kind = t.Kind()
        if kind == reflect.Ptr {
            var ve = reflect.ValueOf(v).Elem()
            var fieldsNum = ve.NumField()
            var params = []reflect.Value{reflect.ValueOf(v)}
            for i := 0; i < fieldsNum; i++ {
                var field = t.Elem().Field(0)
                var tag = field.Tag
                var methodName, ok = tag.Lookup("method")
                if !ok || methodName == "-" {
                    continue
                }
                if methodName == "" {
                    methodName = "Get" + field.Name
                } else if strings.HasPrefix(methodName, "-") {
                    methodName = methodName[1:] + field.Name
                }
                var method, exist = t.MethodByName(methodName)
                if !exist {
                    continue
                }
                var vals = method.Func.Call(params)
                if len(vals) != 1 {
                    continue
                }
                ve.Field(i).Set(vals[0]) // panic if type not compatible
            }
        }
    }
    return ToJSONB(v)
}

func ToJSON(v interface{}) string {
    if s, err := json.MarshalToString(v); err != nil {
        log.Warning(err)
        return ""
    } else {
        return s
    }
}
