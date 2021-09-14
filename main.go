package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/controller"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routes"
	"web_app/settings"

	"go.uber.org/zap"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
// GO Web开发比较通用的脚手架模版
func main() {
	if len(os.Args) < 2 {
		fmt.Print("请指定config文件！")
		return
	}

	// 1. 加载配置
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings failed, err: %v \n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v \n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3. 初始化MySql连接
	if err := mysql.Init(settings.Conf.MySqlConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v \n", err)
		return
	}
	defer mysql.Close()

	// 3. 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v \n", err)
		return
	}
	defer redis.Close()

	// 雪花算法生成用户id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed, err: %v \n", err)
		return
	}

	// 初始化校验器的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init InitTrans failed, err: %v \n", err)
		return
	}

	// 5. 注册路由
	r := routes.Setup(settings.Conf.Mode)

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
