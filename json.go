package fastjson

import (
	. "github.com/Danceiny/go.utils"
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
	} else {
		return &o
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
	} else {
		return postParseArray(&o, reflect.Invalid)
	}
}

func ParseArrayB(s []byte) *JSONArray {
	var o JSONArray
	if err := json.Unmarshal(s, &o.arr); err != nil {
		log.Warning(err)
		return nil
	}
	return postParseArray(&o, reflect.Invalid)
}

func ParseArrayT(s string, t reflect.Kind) *JSONArray {
	var o JSONArray
	if err := json.UnmarshalFromString(s, &o.arr); err != nil {
		log.Warning(err)
		return nil
	}
	return postParseArray(&o, t)
}

func ParseArrayBT(s []byte, t reflect.Kind) *JSONArray {
	var o JSONArray
	if err := json.Unmarshal(s, &o.arr); err != nil {
		log.Warning(err)
		return nil
	} else {
		return postParseArray(&o, t)
	}
}

func postParseArray(o *JSONArray, t reflect.Kind) *JSONArray {
	o.size = len(o.arr)
	if t != reflect.Invalid {
		TransferInterfaces(&o.arr, t)
	}
	return o
}
