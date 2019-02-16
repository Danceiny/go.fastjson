package fastjson

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestJSONObject_Put(t *testing.T) {
	jsonObj := ParseObject("{\"bid\": null}")
	if bid, _ := jsonObj.GetString("bid"); bid != "" {
		t.Errorf("bid is: %v", jsonObj.Get("bid"))
	}
}

func TestJSONObject_getInt(t *testing.T) {
	o := JSONObject{"int": 32, "uint64": uint64(32), "int64": int64(32), "int32": int32(32), "string": "32"}
	var vint int
	var vint32 int32
	var vint64 int64
	vint, _ = o.GetInt("int")
	vint32, _ = o.GetInt32("int32")
	vint64, _ = o.GetInt64("int64")
	assert.Equal(t, 32, vint)
	assert.Equal(t, int32(32), vint32)
	assert.Equal(t, int64(32), vint64)

	vint, _ = o.GetInt("string")
	vint32, _ = o.GetInt32("string")
	vint64, _ = o.GetInt64("string")
	assert.Equal(t, 32, vint)
	assert.Equal(t, int32(32), vint32)
	assert.Equal(t, int64(32), vint64)
}

func TestNilIssue(t *testing.T) {

}

func TestTypeIssue(t *testing.T) {
	f := func() interface{} {
		var a int32
		return a
	}
	b := f()
	defer func() {
		err := recover()
		if err != nil {
			log.Print(err)
		}
	}()
	_ = b.(int64)
}
