package cmd

import (
	"context"
	"fmt"
	"gin-project/api"
	"gin-project/internal/config"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		start(confFilePath)
	},
}

func start(confFilePath string) {
	err := initCommonPart(confFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := config.GetConfig()

	// 启动gin server，将阻塞执行直到server退出，
	// 所以这行代码放在函数最后！！！
	startServers(cfg.APIServerConf)
}

func startServers(cf config.APIServerConf) {
	errChan := make(chan error, 1)

	// 加载路由
	r := api.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cf.Ip, cf.Port),
		Handler: r,
	}

	go func() {
		// 启动HTTP服务器，ListenAndServe()会阻塞当前goroutine直到服务停止
		// http.ErrServerClosed是调用Shutdown()后产生的正常关闭错误
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("[HTTP] server err: %v", err)
			errChan <- err
			return
		}
	}()

	// 处理错误和系统信号
	handleErr(srv, errChan, cf.GracefulShutdownTimeSec)
}

func handleErr(srv *http.Server, errChan chan error, gracefulShutdownTime int) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
	for {
		select {
		case sig := <-sigCh:
			fmt.Printf("[http] catch signal(%+v), stop server\n", sig)
			shutdownServer(srv, gracefulShutdownTime)
			return
		case err := <-errChan:
			fmt.Printf("[http] catch server err:%v", err)
			shutdownServer(srv, gracefulShutdownTime)
			return
		}
	}
}

func shutdownServer(srv *http.Server, gracefulShutdownTime int) {
	// 创建一个超时ctx，若关闭操作超过此时间，会强制终止。
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(gracefulShutdownTime)*time.Second)
	defer cancel()
	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("[http] server shutdown err:%v", err)
	}
	fmt.Println("[http] Server shutdown completed")
}
