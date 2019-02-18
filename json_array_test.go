package fastjson

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseArray(t *testing.T) {
	arr := ParseArray("[{\"bid\": null},{\"bid\": null}]")
	assert.Equal(t, 2, arr.Size())
	arr.Put(JSONObject{"a": 10, "b": "10"})
	assert.Equal(t, 3, arr.Size())
	obj := arr.GetJSONObject(2)
	v, ok := obj.GetInt("a")
	v2, ok := obj.GetInt("b")
	assert.Equal(t, true, ok)
	assert.Equal(t, 10, v)
	assert.Equal(t, 10, v2)
}

func TestParseArrayT(t *testing.T) {
	arr := ParseArrayT("[10,12]", reflect.Int)
	assert.Equal(t, 10, arr.Get(0))
}
