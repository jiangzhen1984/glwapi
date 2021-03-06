
package glwapi_test


import (
	"testing"
	"com/glwapi"
	"encoding/json"
)


func TestCodeUrl(t * testing.T) {
	var g * glwapi.GlwapiConfig = glwapi.D()

	u, s, e := g.CodeURL("", "")
	if e == nil {
		t.Errorf("Should not genertate url without redirect url")
	} else {
		t.Logf(e.Error())
	}

	u, s, e = g.CodeURL("http://aa.bb.cc/a?a=sg&c=11", "a")
	if e != nil {
		t.Errorf(e.Error())
	}
	if s != "a" {
		t.Errorf("state code mismatch")
	}
	t.Logf("%s", u)
	t.Logf("%s", s)

}


func TestUserTokenURL(t * testing.T) {
	var g * glwapi.GlwapiConfig = glwapi.D()
	u, e := g.UserTokenURL("")
	if e == nil {
		t.Errorf("Should not genertate url without code")
	} else {
		t.Logf(e.Error())
	}

	t.Logf(u)

	u, e = g.UserTokenURL("1234")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		t.Logf(u)
	}
}


func TestHandleTokenURLResponse(t * testing.T) {
	var g * glwapi.GlwapiConfig = glwapi.D()
	var u * glwapi.WeChatUser = &glwapi.WeChatUser{"a", 2, "c", "d", "e", 0, nil}
	b, e :=	json.Marshal(u)
	if e == nil {
		_, e1 := g.HandleUserTokenURLResponse(b)
		if e1 != nil {
			t.Errorf(e1.Error())
		} else {
			t.Logf("Test ok Case 1  \n" )
		}

	} else {
		t.Errorf(e.Error())
	}

	u = &glwapi.WeChatUser{"a", 2, "c", "d", "e", 11, nil}
	b, e =	json.Marshal(u)
	if e == nil {
		_, e1 := g.HandleUserTokenURLResponse(b)
		if e1 != nil {
			t.Logf("Test ok Case 2")
		} else {
			t.Errorf("Test failed should get error ")
		}
	} else {
		t.Errorf(e.Error())
	}
}
