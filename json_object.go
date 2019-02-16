package fastjson

import (
	log "github.com/sirupsen/logrus"
)

var json = ConfigWithCustomTimeFormat

type JSON interface {
	ToJSON() []byte
	ToJSONString()
	Parse(s string) JSONObject
}

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

func (jsonObject *JSONObject) PutFluent(k string, v interface{}) *JSONObject {
	(*jsonObject)[k] = v
	return jsonObject
}

func ParseString(s string) JSONObject {
	var o JSONObject
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		log.Warning(err)
		return nil
	} else {
		return o
	}
}

func Parse(bytes []byte) JSONObject {
	var o JSONObject
	if err := json.Unmarshal(bytes, &o); err != nil {
		log.Warning(err)
		return nil
	} else {
		return o
	}
}
func (jsonObject *JSONObject) ToJSONString() string {
	return string(jsonObject.ToJSON())
}

func (jsonObject *JSONObject) ToJSON() []byte {
	if bytes, err := json.Marshal(jsonObject); err != nil {
		log.Warning(err)
		return nil
	} else {
		return bytes
	}
}

func ToJSON(v interface{}) []byte {
	if bytes, err := json.Marshal(v); err != nil {
		log.Warning(err)
		return nil
	} else {
		return bytes
	}
}
