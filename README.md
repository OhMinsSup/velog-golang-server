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
