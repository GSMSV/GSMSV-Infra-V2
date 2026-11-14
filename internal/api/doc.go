// Package api는 Kotlin API 서버만 부르는 내부 HTTP API다.
//
// 여기서 할 일: 라우팅, 서비스 사이 인증, request_id 전달, 에러 응답 형식.
//
// Go 실행 엔진 어디에도 넣지 않는 것: 권한 판단, 할당량 계산, 노드 선택,
// 만료 정책, 팀·사용자 개념, GSMSV DB 접근, 사용자에게 보일 문장.
// 권한 판단은 Kotlin 자원 로더에서만 한다(기능명세서 22번). 나머지는
// GSMSV-Contracts-V2 policies/kotlin-go-boundary.md의 레드라인이다.
package api
