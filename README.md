# jwt
便捷使用jwt

##### usage
1. 定义一个Payload
```go
type Person struct {
    Name string
    Age  int
}
```
2. 创建jwt
```go
	p := Person{
		"张三",
		18,
	}
	tokenStr, err := New[Person](
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
```

3. 解析jwt
```go
    v, err := Parse[Person](tokenStr)
```
