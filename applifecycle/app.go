package applifecycle

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	Name       string
	httpServer *http.Server
}

func (srv *Server) Start(ctx context.Context) error {
	fmt.Println("start", srv.Name)
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Stop(ctx context.Context) error {
	fmt.Println("stop", srv.Name)
	return srv.httpServer.Shutdown(ctx)
}

type App struct {
	Servers []*Server
}

func (app *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	for _, server := range app.Servers {
		server := server
		g.Go(func() error {
			<-ctx.Done()
			return server.Stop(ctx)
		})

		g.Go(func() error {
			return server.Start(ctx)
		})
	}

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	g.Go(func() error {
		select {
		case <-ctx.Done(): // 执行 server.Start(ctx) 返回错误时，errgroup 会触发 ctx cancel
		case <-quit: // 接收到kill信号时，主动触发cancel唤醒 server.Stop(ctx)
			cancel()
		}
		return nil
	})
	err := g.Wait()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
