# GO로 API 서버 만들어보기

### 로직별 폴더/파일 분리하기
    - config: DB connect, ORM AutoMigrate
    - controller: API 요청 처리
    - DAO: DB 접근, CRUD ...
    - middleware: jwtMiddleware
    - models: ORM 정의
    - routes: 라우팅, 컨트롤러 매핑
    - utils: 도구
    - views: 테스트 페이지

### DB, ORM 연결
    dotenv로 환경변수 관리
    gorm 라이브러리 사용
        - gorm.io/gorm
        - gorm.io/driver/mysql
    

### API 생성, ORM 사용
### 회원가입, 로그인 기능 구현
### jwt 인증 구현