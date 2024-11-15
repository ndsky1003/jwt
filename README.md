# jwt
便捷使用jwt

以下是 JWT（JSON Web Token）中各个字段的意义：

1. **`iss` (Issuer)**：
   - 表示 JWT 的发行者。通常是一个标识符，用于确认谁创建了这个令牌。

2. **`sub` (Subject)**：
   - 表示 JWT 的主题，通常是用户的唯一标识符（如用户ID），用于指明该令牌所针对的用户。

3. **`aud` (Audience)**：
   - 表示令牌的受众，指明该 JWT 是为哪个用户或系统设计的。可以是一个或多个受众的标识符。

4. **`exp` (Expiration Time)**：
   - 表示令牌的过期时间，过期后该令牌将不再有效。通常以 Unix 时间戳的形式表示。

5. **`nbf` (Not Before)**：
   - 表示令牌在何时之前不可用。只有在指定的时间之后，令牌才会被视为有效。

6. **`iat` (Issued At)**：
   - 表示令牌的签发时间，通常用于记录令牌何时被创建。

7. **`jti` (JWT ID)**：
   - 表示 JWT 的唯一标识符，用于防止重放攻击。每个令牌应该有一个唯一的 ID。

这些字段有助于在身份验证和授权过程中提供必要的信息，确保令牌的使用安全可靠。
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
			SetAudience("aa", "bb").
			SetID("12341234").
			SetExpiresAt(time.Now().Add(100*time.Second)).
			SetNotBefore(time.Now().Add(-1*time.Second)).
			SetIssuedAt(time.Now().Add(1000*time.Second)),
	)
```

3. 解析jwt
```go
    v, err := Parse[Person](tokenStr) //就算有错,令牌过期这些逻辑错误,都会返回这个token所承载的内容
```
