package fastjson

import (
    "fmt"
    "strconv"
)

type JSONObject map[string]interface{}

func (jsonObject *JSONObject) Put(k string, v interface{}) {
    (*jsonObject)[k] = v
}

func (jsonObject *JSONObject) Get(k string) interface{} {
    return (*jsonObject)[k]
}

func (jsonObject *JSONObject) GetString(k string) (string, bool) {
    v := jsonObject.Get(k)
    if v == nil {
        return "", false
    } else {
        return v.(string), true
    }
}
func (jsonObject *JSONObject) getInt(k string, bitSize int) (int64, bool) {
    v := jsonObject.Get(k)
    if v == nil {
        return 0, false
    } else {
        switch v.(type) {
        case int32:
            return int64(v.(int32)), true
        case int:
            return int64(v.(int)), true
        case string:
            ret, err := strconv.ParseInt(v.(string), 10, bitSize)
            if err == nil {
                return ret, true
            } else {
                return 0, true
            }
        default:
            ret, err := strconv.ParseInt(fmt.Sprint(v), 10, bitSize)
            if err == nil {
                return ret, true
            } else {
                return 0, true
            }
        }
    }
}

func (jsonObject *JSONObject) GetInt64(k string) (int64, bool) {
    v, ok := jsonObject.getInt(k, 64)
    return int64(v), ok
}
func (jsonObject *JSONObject) GetInt32(k string) (int32, bool) {
    v, ok := jsonObject.getInt(k, 32)
    return int32(v), ok
}
func (jsonObject *JSONObject) GetInt(k string) (int, bool) {
    v, ok := jsonObject.getInt(k, 0)
    return int(v), ok
}

func (jsonObject *JSONObject) PutFluent(k string, v interface{}) *JSONObject {
    (*jsonObject)[k] = v
    return jsonObject
}

func (jsonObject *JSONObject) ToJSONString() string {
    return ToJSONString(jsonObject)
}

func (jsonObject *JSONObject) ToJSON() []byte {
    return ToJSON(jsonObject)
}
