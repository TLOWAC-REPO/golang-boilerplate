
## Folder Structure Draft
.
├── README.md                     # 프로젝트 설명 및 사용법을 담고 있는 문서
├── api                           # API 관련 코드 모음
│   ├── controllers               # 클라이언트 요청을 처리하는 핸들러 (API 엔드포인트)
│   ├── repository                # 데이터베이스와의 상호작용을 담당하는 리포지토리 패턴 구현
│   ├── routes                    # API 라우팅 설정 (엔드포인트와 핸들러 매핑)
│   └── services                  # 비즈니스 로직을 포함하는 서비스 계층
├── cmd                           # 애플리케이션의 진입점
│   └── app                       # 애플리케이션 설정 및 실행을 담당하는 폴더
│       └── main.go               # 메인 함수, 애플리케이션 실행 시작점
├── go.mod                        # Go 모듈 의존성 관리 파일
├── makefile                      # 빌드, 테스트, 실행 명령어를 관리하는 파일
├── pkg                           # 공통 라이브러리 및 유틸리티 코드 모음
│   ├── config                    # 애플리케이션 설정 관련 코드 (환경 변수 및 설정 파일 로드)
│   │   └── config.go             # 설정 로드를 위한 구조체 및 메서드 구현
│   ├── db                        # 데이터베이스 연결 및 설정 관리
│   │   └── db.go                 # 데이터베이스 초기화 및 설정 관련 코드
│   ├── logger                    # 로깅 관련 설정 및 초기화 코드
│   │   └── logger.go             # 로깅 기능을 구현한 코드
│   ├── middleware                # 미들웨어 관련 코드
│   │   ├── auth.go               # 인증 미들웨어 (예: JWT 인증)
│   │   └── logging.go            # 요청 로깅을 위한 미들웨어
│   └── util                      # 자주 사용되는 공통 유틸리티 함수들
│       └── util.go               # 다양한 유틸리티 함수 구현
├── scripts                       # 스크립트 및 자동화 도구 (예: DB 마이그레이션)
└── test                          # 테스트 코드 모음
    ├── api_test.go               # API 관련 테스트 코드
    └── pkg_test.go               # pkg 관련 테스트 코드

## Lib

- [mocking](https://github.com/uber-go/mock)


## 구조 참고
- https://github.com/atharv-bhadange/go-api-template/tree/main/api

