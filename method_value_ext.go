package fastjson

import (
    "github.com/json-iterator/go"
    "reflect"
    "strings"
    "unsafe"
)

// compatible with java constraint,
// serialize `GetXXX` return values with XXX's tag `method:"Get"`

type MethodValueExtension struct {
    jsoniter.DummyExtension
}

func (extension *MethodValueExtension) UpdateStructDescriptor(sd *jsoniter.StructDescriptor) {
    var t = sd.Type
    var newValReal = t.New()
    var t2 = reflect.TypeOf(newValReal)
    var isPtr = strings.HasPrefix(t2.String(), "*")
    var t1 = t.Type1()
    var methodsNum = t1.NumMethod()
    var methodsNum2 = t2.NumMethod()
    if methodsNum == 0 && methodsNum2 == 0 {
        return
    }
    // 作为方法调用的第一个参数（this, self）
    var stValue reflect.Value
    if isPtr {
        stValue = reflect.ValueOf(newValReal)
    } else {
        stValue = reflect.ValueOf(t1)
    }
    for _, binding := range sd.Fields {
        var variableName = binding.Field.Name()
        var methodName = binding.Field.Tag().Get("method")
        if methodName == "" || methodName == "-" {
            continue
        }
        if strings.HasPrefix(methodName, "-") {
            methodName = methodName[1:] + variableName
        }
        var method reflect.Method
        var exist bool
        if isPtr {
            method, exist = t2.MethodByName(methodName)
        } else {
            method, exist = t1.MethodByName(methodName)
        }
        if !exist {
            // method not found
            continue
        }
        binding.Encoder = &funcEncoder{fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
            var params = []reflect.Value{stValue}
            var vals = method.Func.Call(params)
            if len(vals) != 1 {
                // not support return multi value
                return
            }
            var val = vals[0]
            stream.WriteVal(val.Interface())
        }}
    }

}
