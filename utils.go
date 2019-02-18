package fastjson

import (
    "errors"
    "fmt"
    "os"
    "reflect"
    "strconv"
)

// override go.utils.GetEntOrDefault
func GetEnvOrDefault(k string, dv interface{}) interface{} {
    v, ok := os.LookupEnv(k)
    if !ok {
        return dv
    } else {
        if dv == nil {
            return v
        }
        var err error
        switch dv.(type) {
        case string:
            return v
        case bool:
            var ret bool
            ret, err = strconv.ParseBool(v)
            if err == nil {
                return ret
            }
            break
        case int:
            var ret int
            ret, err = strconv.Atoi(v)
            if err == nil {
                return ret
            }
            break
        case int64:
            var ret int64
            ret, err = strconv.ParseInt(v, 10, 64)
            if err == nil {
                return ret
            }
            break
        case int32:
            var ret int64
            ret, err = strconv.ParseInt(v, 10, 32)
            if err == nil {
                return ret
            }
            break
        case float64:
            var ret float64
            ret, err = strconv.ParseFloat(v, 64)
            if err == nil {
                return ret
            }
            break
        case float32:
            var ret float64
            ret, err = strconv.ParseFloat(v, 32)
            if err == nil {
                return ret
            }
            break
        case map[string]interface{}:
            ret := make(map[string]interface{})
            err = json.UnmarshalFromString(v, &ret)
            if err == nil {
                return ret
            }
        case JSONObject:
            ret := ParseObject(v)
            if ret == nil {
                err = errors.New("parse to json failed")
            } else {
                return *ret
            }
            break
        case JSONArray:
            ret := ParseArray(v)
            if ret == nil {
                err = errors.New("parse to json failed")
            } else {
                return *ret
            }
            break
        default:
            fmt.Printf("type: %v", reflect.TypeOf(dv))
            return v
        }
        if err != nil {
            // TODO: 此处返回值可商榷
            return dv
        }
        return nil
    }
}
