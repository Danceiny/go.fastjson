package fastjson

import (
    "fmt"
    "testing"
)

type User2 struct {
    Id      int  `json:"id"`
    IsReal  bool `json:"isReal" method:"-Get"`
    IsReal2 bool `json:"isReal2" method:"GetisReal"`
}

func (u *User2) GetisReal() bool {
    // todo: return u.Id == 10
    return true
}
func (u User2) GetIsReal() bool {
    // todo: return u.Id == 10
    return true
}

func TestMethodValueExtension(t *testing.T) {
    RegisterExt(METHOD_VALUE_EXT)
    var u = User2{
        Id: 10,
    }
    var js1 = ToJSON(&u)
    var js2 = ToJSON(&u)
    fmt.Println(js1)
    fmt.Println(js2)
}
