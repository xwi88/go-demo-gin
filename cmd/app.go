// Package main
// budget control command
package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kdpujie/log4go"
	"github.com/spf13/cobra"
	"github.com/xwi88/kit4go/datetime"
	"github.com/xwi88/kit4go/utils"

	"github.com/xwi88/go-demo-gin/api/handlers"
	"github.com/xwi88/go-demo-gin/api/middleware"
	"github.com/xwi88/go-demo-gin/configs"
	"github.com/xwi88/go-demo-gin/constants"
	"github.com/xwi88/go-demo-gin/resources"
)

// startCMD real time budget control service
var startCMD = &cobra.Command{
	Use:       "start",
	Short:     "Start the app service",
	Long:      ``,
	Example:   "app start\n  app start -c [file]\n  app start --config [file]",
	ValidArgs: []string{"start"},
	PreRun: func(cmd *cobra.Command, args []string) {
		// TIPS: init configs and resources here
		log.Printf("[start-PreRun] load config file:%v", *budgetConfFile)
		err := configs.Init(*budgetConfFile)
		if err != nil {
			log.Panicf(fmt.Sprintf("[start-PreRun] parse config err, %v", err.Error()))
		}
		cfg := configs.GetCfg()
		// include log init
		err = fillAndFixConfig(*budgetConfFile)
		if err != nil {
			log.Panicf("[start-PreRun] fillAndFixConfig err:%v", err.Error())
		}
		// set GOMAXPROCS
		numCPU := runtime.NumCPU()
		runtime.GOMAXPROCS(numCPU)
		cfg.NodeInfo.GOMAXPROCS = numCPU

		log4go.Info("[start-PreRun] runtime config:%v", configs.GetCfg())

		exportConfigForce := cfg.APP.ExportConfigForce
		exportConfigPath := cfg.APP.ExportConfigPath
		exportConfig := cfg.APP.ExportConfig
		if exportConfig {
			existExportCfgPath := utils.IsExist(exportConfigPath)
			if !existExportCfgPath && exportConfigForce {
				err = os.Mkdir(exportConfigPath, os.ModePerm)
				if err != nil {
					log4go.Error(err.Error())
				}
			}
			if cfg.APP.ExportConfigUnique {
				err = configs.WriteConfigWithPath(exportConfigPath, *budgetConfFile)
			} else {
				err = configs.SafeWriteConfigWithPath(exportConfigPath, *budgetConfFile)
			}
			if err != nil {
				log4go.Error(err.Error())
			}
		}

		err = resources.Init()
		if err != nil {
			log4go.Error(err.Error())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TIPS: start main service
		pid := syscall.Getpid()
		log4go.Info("[start-Run] pid:%v, at:%v", pid, datetime.GetNowWithZone(nil))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// TIPS: block until, the resources destroy
		// log4go.Info("[start-PostRun] PostRun")
		// for temp
		ctx, cancelFunc := context.WithCancel(context.Background())

		catchSignal := utils.NewCatchSignal()
		catchSignal.RegisterSigFunc(func() {
			log4go.Warn("[start-PostRun] resources close start")
			cancelFunc()
			resources.Close()
			log4go.Warn("[start-PostRun] resources close end")
		}).Start()
		go tickerService(ctx)

		corsConfig := cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: false,
			MaxAge:           24 * time.Hour,
		}
		cfg := configs.GetCfg()
		corsConfig.AllowAllOrigins = true
		gin.SetMode(cfg.APP.Mode)
		r := gin.New()
		r.ForwardedByClientIP = false
		r.HandleMethodNotAllowed = false // 严格限制路由请求方法；不设置，NotFound handler
		r.Use(middleware.Logger())
		r.Use(gin.Recovery())
		r.Use(cors.New(corsConfig))

		// routers
		// 接口服务
		apiGroup := r.Group("/api/v1")
		apiGroup.GET("/simple", handlers.SimpleHandler)

		adminGroup := r.Group("/admin")
		// /admin/version
		adminGroup.GET("/version", handlers.VersionHandler)

		routersInfo := bytes.NewBufferString("")
		for index, route := range r.Routes() {
			routersInfo.WriteString(fmt.Sprintf("\n%d|%v|%v|%v", index, route.Method, route.Path, route.Handler))
		}
		log4go.Debug("routes_info: %v", routersInfo.String())
		log4go.Debug("Listening and serving HTTP on %v, pid: %v", cfg.APP.Addr, os.Getpid())

		s := &http.Server{
			Addr:           cfg.APP.Addr,
			Handler:        r,
			ReadTimeout:    cfg.APP.ReadTimeout,
			WriteTimeout:   cfg.APP.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}
		err := s.ListenAndServe()
		if err != nil {
			log4go.Error(err.Error())
		}
		// block here
		// select {}
	},
}

func tickerService(ctx context.Context) {
	cfg := configs.GetCfg()
	ti := cfg.APP.TickerInterval
	if ti <= 0 {
		log4go.Warn("[tickerService] interval shall less than 0, exit")
		return
	}
	tk := time.NewTicker(ti)
	defer tk.Stop()
loop:
	for {
		select {
		case <-tk.C:
			log4go.Debug("ticker")
		case <-ctx.Done():
			break loop
		}
	}
}

// fill and fix partial config
func fillAndFixConfig(file string) (err error) {
	cfg := configs.GetCfg()
	if cfg == nil {
		return errors.New("nil config")
	}
	publicIP, _ := utils.PublicIP()
	serverIP := utils.LocalIP()

	// NodeInfo check and set
	cfg.NodeInfo.PID = syscall.Getpid()
	cfg.NodeInfo.ServerIP = serverIP
	cfg.NodeInfo.PublicIP = publicIP
	cfg.NodeInfo.NumCPU = runtime.NumCPU()
	cfg.NodeInfo.HostName, _ = os.Hostname()
	cfg.NodeInfo.StartTime = datetime.GetNowWithZone(nil)
	cfg.NodeInfo.WorkDir, _ = os.Getwd()
	cfg.NodeInfo.ConfigFile = file

	if len(cfg.APP.ExportConfigPath) == 0 {
		cfg.APP.ExportConfigPath = constants.ExportConfigPath
	}

	if mode, exist := constants.AppModeMap[cfg.APP.Mode]; exist {
		cfg.APP.Mode = mode
	} else {
		cfg.APP.Mode = constants.ReleaseMode
	}

	// log4go init
	extraFields := cfg.Log4go.KafKaWriter.MSG.ExtraFields
	if extraFields == nil {
		extraFields = make(map[string]interface{})
		cfg.Log4go.KafKaWriter.MSG.ExtraFields = extraFields
	}
	cfg.Log4go.KafKaWriter.MSG.PublicIP = publicIP
	cfg.Log4go.KafKaWriter.MSG.ServerIP = serverIP

	// 自定义字段，如首层不存在，则提取到日志信息第一层!
	extraFields["app_env"] = cfg.APP.AppEnv
	extraFields["app_name"] = cfg.APP.AppName
	extraFields["hostname"], _ = os.Hostname()
	cfg.Log4go.KafKaWriter.MSG.ExtraFields = extraFields
	err = log4go.SetupLog(cfg.Log4go)
	return
}
