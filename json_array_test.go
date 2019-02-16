package fastjson

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParseArray(t *testing.T) {
    arr := ParseArray("[{\"bid\": null},{\"bid\": null}]")
    assert.Equal(t, 2, arr.Size())
    arr.Put(&JSONObject{"a": 10, "b": "10"})
    assert.Equal(t, 3, arr.Size())
    v, ok := arr.Get(2).GetInt("a")
    v2, ok := arr.Get(2).GetInt("b")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, v)
    assert.Equal(t, 10, v2)

}
