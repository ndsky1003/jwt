package jwt

import (
	"testing"
	"time"
)

type Person struct {
	Name string
	Age  int64
}

func Test_New(t *testing.T) {
	n := time.Now().UnixNano() //223
	p := Person{
		"张三",
		n,
	}
	ss, err := New(
		p,
		Option().
			SetExpiresAt(time.Now().Add(-100*time.Second)).
			SetNotBefore(time.Now().Add(-1*time.Second)).
			SetIssuedAt(time.Now().Add(1000*time.Second)),
	)
	t.Log(len(ss), ss, err)
	v, err2 := Parse[Person](ss)
	t.Log(v.GetAudience())
	t.Logf("err2:%v,v:%+v", err2, v)
}
