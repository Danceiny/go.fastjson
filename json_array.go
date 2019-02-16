package fastjson

const DEFAULT_CAPACITY = 16

type JSONArray struct {
	arr  []*JSONObject
	size int
	cur  int
}

func NewJSONArray() *JSONArray {
	arr := make([]*JSONObject, 0, DEFAULT_CAPACITY)
	return &JSONArray{arr: arr}
}

func NewJSONArrayLimited(capacity int) *JSONArray {
	arr := make([]*JSONObject, 0, capacity)
	return &JSONArray{arr: arr}
}

func (jsonArray *JSONArray) Put(v *JSONObject) {
	jsonArray.arr = append(jsonArray.arr, v)
	jsonArray.size++
}

func (jsonArray *JSONArray) Get(index int) *JSONObject {
	return (jsonArray.arr)[index]
}

func (jsonArray *JSONArray) ToJSONString() string {
	return ToJSONString(jsonArray.arr)
}

func (jsonArray *JSONArray) ToJSON() []byte {
	return ToJSON(jsonArray.arr)
}

func (jsonArray *JSONArray) Size() int {
	return jsonArray.size
}

func (jsonArray *JSONArray) cap() int {
	return cap(jsonArray.arr)
}
