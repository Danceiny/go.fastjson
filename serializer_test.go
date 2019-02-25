package fastjson

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestToJSONBv2(t *testing.T) {
    var u = User2{}
    var bytes []byte
    var ja *JSONObject
    bytes = ToJSONBv2(u, &JSONSerializerFeature{MethodValue: true})
    fmt.Println(string(bytes))
    ja = ParseObjectB(bytes)
    assert.Equal(t, true, ja.GetBool("isReal2"))

    bytes = ToJSONBv2(&u, &JSONSerializerFeature{MethodValue: true})
    ja = ParseObjectB(bytes)
    fmt.Println(string(bytes))
    assert.Equal(t, true, ja.GetBool("isReal2"))
}
