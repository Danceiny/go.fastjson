package fastjson

const DEFAULT_CAPACITY = 16

type JSONArray struct {
    arr  []interface{}
    size int
    cur  int
    err  error
}

func NewJSONArray() *JSONArray {
    arr := make([]interface{}, 0, DEFAULT_CAPACITY)
    return &JSONArray{arr: arr}
}

func NewJSONArrayLimited(capacity int) *JSONArray {
    arr := make([]interface{}, 0, capacity)
    return &JSONArray{arr: arr}
}

func (ja *JSONArray) Put(v interface{}) {
    ja.arr = append(ja.arr, v)
    ja.size++
}

func (ja *JSONArray) Set(index int, v interface{}) {
    var capacity = ja.cap()
    var size = ja.size
    if index >= size && index < capacity {
        ja.arr = append(ja.arr, v)
        return
    } else if index >= capacity {
        var newArr = make([]interface{}, index+1, index*2)
        copy(newArr, ja.arr)
        ja.arr = newArr
        ja.size = len(newArr)
    } else if index < 0 {
        index = size + index
    }
    ja.arr[index] = v
}

func (ja *JSONArray) Get(index int) interface{} {
    return (ja.arr)[index]
}

func (ja *JSONArray) GetJSONObject(index int) *JSONObject {
    var ret = (ja.arr)[index]
    switch ret.(type) {
    case *JSONObject:
        return ret.(*JSONObject)
    case JSONObject:
        var ja = ret.(JSONObject)
        return &ja
    case map[string]interface{}:
        return NewJSONObjectFromMap(ret.(map[string]interface{}))
    }
    return nil
}

func (ja *JSONArray) ToJSON() string {
    return ToJSON(ja.arr)
}

func (ja *JSONArray) ToJSONB() []byte {
    return ToJSONB(ja.arr)
}

func (ja *JSONArray) Size() int {
    return ja.size
}

func (ja *JSONArray) cap() int {
    return cap(ja.arr)
}

func (ja *JSONArray) Next() bool {
    if ja.cur == ja.size {
        return false
    }
    ja.cur++
    return true
}

func (ja *JSONArray) Current() interface{} {
    return ja.Get(ja.cur - 1)
}

func (ja *JSONArray) GetCurrentJSONObject() interface{} {
    return ja.GetJSONObject(ja.cur - 1)
}

func (ja *JSONArray) Error() error {
    return nil
}

func (ja *JSONArray) Prev() bool {
    if ja.cur <= 1 {
        return false
    }
    ja.cur--
    return true
}

func (ja *JSONArray) Values() []interface{} {
    return ja.arr
}
