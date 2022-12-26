# db 설계 관련

registerqueue 가 어떻게 생겨야 하는지

gorm의 field - 종속성을 중심으로 분류
json의 객체 - 클라이언트 요청과 직접적으로 관련이 있는 것만 포함
go의 struct - json 과 db 사이의 매개체

이걸 따로 정의할건지 같이 정의할건지의 문제인데

같이 정의함 -> 편의상 이득
따로 정의함 -> 처리 시간 단축

시간 이득은 굉장히 미미하므로 그냥 합쳐줘도 상관 없음
가입 신청만 하면 일단 큐에 올리는걸로 하자

아니 그러면 registerqueue 를 만들 필요가 없는데 애초에

userauth - for login & auth
personalinfo - for personal information

DailyMileage - for mileage status
mileage
mileagerequestqueue - for mileage request

이렇게 4개의 db table 을 만드는게 자연스러움

DB가 끊기면 panic 시켜야하는거 아닌가 굳이 예외 처리를 해야하나

db 쿼리를 두 번 던지는게 굉장히 불쾌하네

중복 체크랑 존재 체크랑

예측할 수 있는 예외랑 예측할 수 없는 예외랑 구분하자
syntax error

뭐가 바람직한지를 계속 고민하는중

get 과 post

get mileage request 를 하는데

문법적으로 get 뒤에는 대상을 확정해줄 수 있는 인자가 들어가고, return 값으로 원하는 정보가 들어오는게 바람직하다고 생각
매주마다 리셋되는 마일리지가 있고 누적되는 마일리지가 있기 때문에 이 둘의 성격이 좀 다르지 않나 싶음

WeekendMileage

WeekendMileage

개인 요청부터 시작
MileageRequest

axios 를 사용할때 쿠키 정보도 당연히 넘어오는데
admin 때문에 그런걸로 하자

지금 수준에서는 고민을 줄이는게 좋으니까 구현이 조잡해 보이더라도 그냥 하자

어떤 기준으로 함수들을 분류할건지가 중요한데
rest layer 에서는 기능, 사용자 중심으로 분류가 이뤄졌다면
db layer 에서는 table 별로 분류가 이뤄지는게 편할 것 같아

userauth - for login & auth
personalinfo - for personal information

DailyMileage - for mileage status
mileage
mileagerequestqueue - for mileage request

management - 관리에 필요한 함수들

excel 파일 읽기
일단은 userid 를 받는다고 생각하자
완성이 최우선이니까 요일별 업데이트만 할 수 있는걸로 하자

dgrijalva/jwt-go

redis 로 token 정보를 사용하고
mysql 로 사용자 정보 관리하고
너무 복잡해지네

db 도 전역으로 쓰는게 훨씬 깔끔한데 왜 이 생각을 못했는지

db, redis 이렇게 두 개가 존재하도록 하고
redis도 추상화를 거쳐야 함
redis 를 사용하는 유일한 위치인

userid 를 int 로 관리하는게 좋은가

중복 체크를 하기 위해

드디어 테스트 해볼 수 있는 시간

token 의 sign method 가 적절한지
token 이 expired 되지 않았는지

authority check
-> verify token : 권한 검사, 따로 함수를 만들 이유가 없음
-> extract token : 토큰을 반환
-> token valid : 사인 검사

함수 인자로 c 를 넘겨주는거 별거 아님

토큰 검사 방법
extract token

서명까지 정당하다면 그대로 사용해도 되는건가

extract token
token valid test

1. header 에서 bearToken 을 가져오기
2. bearToken 에서 method 를 검증하고 jwt.Token 을 추출
3. jwt.Token 에서 expired를 검증하고 claim 을 추출
4. claim 은 map 형태로, 필요한 정보들을 추출하여 AccessDetails 로 반환
5. jwt.Token 에서 사용자 정의 struct 추출하기

token 추출
token valid 검사
token 사용

이렇게 3가지 플로우로 구성됨

middleware 내부에 refresh 가 존재해야 함

1. Access 와 Refresh 가 모두 죽은 경우
2. Access 는 죽고 Refresh 는 유효한 경우
3. Access 가 살아있는 경우

이와 같은 세부적인 조작이 가능해야 하기 때문에 이걸 한꺼번에 해서는 안되는 것

middleware 함수를 봤을 때 어느 수준의 보안을 사용하고 있구나 정도의 정보를 알 수 있게끔
refresh, checkauth 두 가지가 드러나도록 한거 잘했어

너무 복잡해

middleware

1. RefreshToken :
   redis 14일짜리 토큰이랑 15분짜리 토큰이랑 둘 다 존재하며 일정 시간이 지난 후 사라짐
   access token 이 없는 경우에 한정
   refresh token 을 독해해서

이 레벨을 거친다면 다음 레벨에서부터 access token 만 확인하면 됨

2. AuthCheck :
   만약 권한이 없다면 Unauthorize 하고 끝내면 됨
   admin 과 user 권한 인증 어떻게 하지

RefreshToken
SetAuthority
CkeckAdmin, CkeckUser
이런 식으로 단계별로 구분

redis 는 refresh-token 을 통해 access-token 을 발급하는데에만 관여함
토큰을 헤체하면 권한을 확인할 수 있고 이걸 기록해두면 됨

토큰을 검증하는 과정

1. ParseTokenFromCookie : cookie 에서 값을 가져와서
2. VerifyTokenMethod : method 검증
3. VerifyExpiration :
4. ExtractToken : 토큰을 추출

5. VerifyToken : 이 전체 과정을 한꺼번에 함수로

middleware 에서 토큰을 다 해체해서 정보를 context 에 전부 기록해놓아야 함

userid 를 param 으로 넘기는게 별로 안좋은건가

userid params sound

그냥 하루빨리 테스트에 들어가고 싶은데 계속 이것저것 문제가 생겨서 뒤쳐지는 느낌
userid 를 int 로 통일
username 을 대용으로 사용하자
실제 사람 이름은 name 을 사용

오늘 무조권 다 끝내고 잔다

DB 쪽 구현을 덜했었구나

PersonalInfo 로 묶일 수 있는 정보 특징 : json 으로 보냄
Join 하는 법을 이참에 익히자

종속성을 한 테이블에서 두 개 가지지 말라고 하니까
Recommender

db 질문글 올려보자

token 이 없다
refresh token 이랑 signin 이랑 로직 똑같아

이이 다까먹었어
auth 만 뚫으면 다 된다는 생각으로 해보자
다시 졸속으로 익혀봐

redis까지 도입하면서 너무 복잡해져버렸어

ExtractTokenFromCookie

VerifyToken

TokenValid

gin.Context 의 Cookie에 String 형태로 저장돼있는거 추출해야 함
token 이 sign method가 맞는지, ACCESS_SECRET 로 뚫리는지 확인해야 함
token 이 유효한지 확인해야 함

jwt token 발행하는 것까지 완성
id가 int인게 더 이상하지 않은지

초기 유저는 master
