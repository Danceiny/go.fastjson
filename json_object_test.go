package fastjson

import (
	"testing"
)

func TestJSONObject_Put(t *testing.T) {
	jsonObj := ParseString("{\"bid\": null}")
	if bid, _ := jsonObj.GetString("bid"); bid != "" {
		t.Errorf("bid is: %v", jsonObj.Get("bid"))
	}
}
