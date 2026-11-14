// Command engine은 GSMSV v2의 Go 실행 엔진이다.
// Kotlin API 서버만 부르는 내부 HTTP API를 연다.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("engine stopped", "err", err)
		os.Exit(1)
	}
}

func run() error {
	// 공개 저장소라 주소·포트 기본값을 코드에 두지 않는다.
	addr := os.Getenv("GSMSV_INFRA_ADDR")
	if addr == "" {
		return errors.New("GSMSV_INFRA_ADDR is required")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 포트를 먼저 잡아야 바인드 실패가 「listening」 로그 없이 바로 드러난다.
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Handler:           http.NewServeMux(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() { errCh <- srv.Serve(ln) }()
	slog.Info("engine listening", "addr", ln.Addr().String())

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	// 처리 중인 HTTP 요청만 기다린다. 작업 상태는 휘발성이고 Kotlin이 실제 상태로 복구한다.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
