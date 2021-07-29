package applifecycle

import (
	"net/http"
	"testing"
)

func TestStartFailed(t *testing.T) {
	servers := []*Server{
		{
			Name:       "api1",
			httpServer: &http.Server{Addr: ":8080"},
		},
		{
			Name:       "api2",
			httpServer: &http.Server{Addr: ":8081"},
		},
		{
			Name:       "api3",
			httpServer: &http.Server{Addr: ":8080"}, // 端口冲突导致启动失败
		},
	}
	app := &App{Servers: servers}
	app.Run()
}

func TestSignShutdown(t *testing.T) {
	servers := []*Server{
		{
			Name:       "api1",
			httpServer: &http.Server{Addr: ":8080"},
		},
		{
			Name:       "api2",
			httpServer: &http.Server{Addr: ":8081"},
		},
		{
			Name:       "api3",
			httpServer: &http.Server{Addr: ":8082"},
		},
	}
	app := &App{Servers: servers}
	app.Run()
}
