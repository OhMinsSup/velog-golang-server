# story-server
🙏🏻 velog-server clone to Golang (Gin)

## EP.1 ORM gorm에서 ent로 변경
일단 오늘은 데이터베이스 orm을 변경했다. 내가 이걸 변경하는 이유는 개발을 진행하면서 gorm에서 모델간의 관계와 model의 
필수값 nullable을 표시하는게 상당히 불편했다. 그리고 무엇보다 변경을 하자고 생각한 이유는 뭔가 유지보수가 안되는 것 같았다.
그래서 변경을 결심했고 ent로 변경을 했다.

현재 User 모델만 생성을 했지만 계속 변경해 나갈 예정이다.

### 오늘 한 일
- ent orm 적용
- gin server 구조 변경및 docker 파일 분리 (redis, postgresql)

## EP.2 유저 회원가입을 기능 구현
entgo에서 대해서 조금 찾아봤는데 code frist로 개발되어서 그런지 원시적으로 raw query를 지워하지 않는다.
만약에 사용할려면 빌드인 sql 모듈과 연동해서 사용해야 한다.

이메일 발송 API를 생성했는데 아주 잘 작동하고 있다.

### 오늘 한 일

- ent orm (emailAuth, User, UserProfile) 모델 구현
- ent orm에서 raw query가 지원이 안되기 때문에 빌드인 sql을 연동해서 사용
- sendEmail API 생성

## EP.3 유저 코드 인증 기능 구현
오늘은 유저 회원가입및 코드 인증부분을 구현했다. 일단 코드로 register token을 발급 받는 부분은 정상적으로 실행을 하고 있지만
로그인 부분은 오류가 있을 것 같다.

왜냐하며 로컬 회원가입중 authToken 테이블을 생성하는 부분에서 sql 에러가 발생했다.
```bash
insert node to table "auth_tokens": pq: insert or update on table "auth_tokens" violates foreign key constraint "auth_tokens_users_auth_token"
```
내가 table 설계를 잘못한 것 같다.... 음... 정상적으로 한 것 같은데... 이건 다음에 수정해야겠다.

### 오늘 한 일

- 로컬 회원가입 API
- 코드 인증 API

## EP.4 유저 회원가입 API
오늘은 저번 이야기에서 유저 토큰이 발급이 안되는 이슈를 해결했다. 이슈의 원인은 entgo에서는 1:N N:M등 다대다 관계에서는
column을 설정할 때 s를 붙여야하고 user 모델에서 토큰을 추가해줘야하는 이슈였다.

하지만 한가지 이슈가 더 발생했는데, authToken을 생성하면 fk_user_id에 값이 들어가야하는데 왠지 모르게 값이 안들어간다...
내가 생각했을 때는 내가 잘못한 것 같은데 한번 더 이유를 찾아봐야겠다.

### 오늘 한 일

- 로컬 회원가입 API - (저번 이야기 이슈 처리)
