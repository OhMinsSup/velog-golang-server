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
