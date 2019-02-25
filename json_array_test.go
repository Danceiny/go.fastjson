package fastjson

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "os"
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
    assert.Equal(t, 10, arr.([]int)[0])
}

func TestGetEnvOrDefault(t *testing.T) {
    v := GetEnvOrDefault("KEY", &JSONObject{"KEY": 10})
    o := v.(*JSONObject)
    vv, ok := o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)

    _ = os.Setenv("KEY", "{\"KEY\": 10}")
    v = GetEnvOrDefault("KEY", &JSONObject{})
    o = v.(*JSONObject)
    vv, ok = o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)

    _ = os.Setenv("KEY", "[{\"KEY\": 10}]")
    v = GetEnvOrDefault("KEY", JSONArray{})
    o2 := v.(JSONArray)
    o2o := o2.GetJSONObject(0)
    fmt.Println("o2o is nil???", o2o)
    vv, ok = o2o.GetInt("KEY")
    assert.Equal(t, true, ok)
    assert.Equal(t, 10, vv)
}

func TestParseArrayT2Slice(t *testing.T) {
    arr := ParseArrayT("[10,\"12\"]", reflect.Int).([]int)
    assert.Equal(t, 2, len(arr))
    assert.Equal(t, 10, arr[0])
    arr2 := ParseArrayT("", reflect.Int)
    assert.Equal(t, nil, arr2)
    arr3 := ParseArrayT("{\"KEY\": 10}", reflect.Int)
    assert.Equal(t, nil, arr3)
}

func TestJSONArray_Next(t *testing.T) {
    ja := NewJSONArray()
    ja.Put(10)
    ja.Put(11)
    assert.Equal(t, 2, ja.Size())
    var cnt = 0
    for ja.Next() {
        ja.Current()
        cnt++
    }
    assert.Equal(t, 2, cnt)

    cnt = 0
    for ja.Prev() {
        ja.Current()
        cnt++
    }
    assert.Equal(t, 1, cnt)
    assert.Equal(t, 10, ja.Current().(int))
}

func TestJSONArray_Set(t *testing.T) {
    ja := NewJSONArrayLimited(3)
    ja.Put(10)
    ja.Put(11)
    ja.Set(1, 13)
    assert.Equal(t, 13, ja.Get(1).(int))
    ja.Set(2, 13)
    assert.Equal(t, 13, ja.Get(2).(int))
    ja.Set(26, 26)
    assert.Equal(t, 26, ja.Get(26).(int))
}
