
주요 메서드
- SignIn
- RefreshToken
- SetAuthority

SignIn
- Create Token
- Create Auth (redis)
- Token on Cookie


RefreshToken
- Get Token
- Token Valid
- Refresh Token

SetAuthority
- Get Token From Cookie
- Check if Token is Valid
- if Access Token Expired then Refresh Token
- Extract Token Metadata
- Set Authority

Refresh Token
Set Authority 에서 하는 역할이 중복됨

코드 수정이 간결하도록 verify token 위치에서 refresh 를 하도록 하자
verify :
1. sign method 가 정당한지
2. outdated 되지는 않았는지


redis 서버에 접속하는 비용 역시 존재하기 때문에 refresh 는 토큰이 만료된 경우에만 사용

refresh 문제가 해결이 되지 않고 있음
refresh 자체도 하나의 미들웨어로 들어가야 함

















