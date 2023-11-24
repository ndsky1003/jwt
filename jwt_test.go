package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ndsky1003/jwt/options"
)

type Person struct {
	Name string
	Age  int
}

func Test_New(t *testing.T) {
	p := Person{
		"张三",
		18,
	}
	ss, err := New[Person](
		p,
		options.New().
			SetIssuer("李四").
			SetSubject("sub").
			SetAudience(&jwt.ClaimStrings{"aa", "bb"}).
			SetID("12341234").
			SetExpiresAt(jwt.NewNumericDate(time.Now().Add(100*time.Second))).
			SetNotBefore(jwt.NewNumericDate(time.Now().Add(-1*time.Second))).
			SetIssuedAt(jwt.NewNumericDate(time.Now().Add(1000*time.Second))),
	)
	t.Log(ss, err)
	v, err2 := Parse[Person](ss)
	t.Logf("%+v,err2:%v", v, err2)
}
