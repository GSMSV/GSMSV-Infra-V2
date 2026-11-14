# GSMSV-Infra-V2

[ GSMSV-V2 | Infra ] GSMSV V2 Proxmox & kubernetes 연동 실행 엔진

GSMSV v2에서 Proxmox·Kubernetes를 다루는 실행 엔진(Go)이다. 아직 개발 전이며, 첫 코드는 인프라 담당이 만든다.

## GSMSV v2 저장소

| 저장소 | 내용 |
|---|---|
| [GSMSV-Server-V2](https://github.com/GSMSV/GSMSV-Server-V2) | API 서버 (Kotlin) |
| [GSMSV-Client-V2](https://github.com/GSMSV/GSMSV-Client-V2) | 웹 콘솔 (React) |
| [GSMSV-Infra-V2](https://github.com/GSMSV/GSMSV-Infra-V2) | Proxmox·Kubernetes를 다루는 실행 엔진 (Go) |
| [GSMSV-Contracts-V2](https://github.com/GSMSV/GSMSV-Contracts-V2) | 저장소끼리 함께 쓰는 약속 (API 명세, 자원 이름 규칙) |

## 참여하기

- 협업 규칙: [CONTRIBUTING](.github/CONTRIBUTING.md)
- 보안 문제 신고: [SECURITY](.github/SECURITY.md)

## 개발

- Go 버전은 `go.mod`의 `go` 줄을 따른다
- 실행: `GSMSV_INFRA_ADDR=127.0.0.1:8080 go run ./cmd/engine`
- 검사: `gofmt -l .`, `go vet ./...`, `go test -race ./...`. PR마다 CI가 같은 검사를 돌린다
- 빌드: `go build -o bin/ ./cmd/...`
- 패키지마다 맡는 일은 `internal/` 아래 각 폴더의 `doc.go`에 적었다
