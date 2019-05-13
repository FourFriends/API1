package main

import (
	"API1/config"
	"API1/model"
	"API1/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	//不能在这里初始化log，因为log在这里还没有初始化
	//log.Info("App Start!")

	pflag.Parse()

	//初始化设置yaml
	err := config.Init(*cfg)
	if err != nil {
		panic(err)
	}

	//初始化db
	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	router.Load(g)

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("port"))

	//先进入检测，然后再调用主进程
	go check()

	http.ListenAndServe(viper.GetString("port"), g)

}

//自检程序

func check() {
	log.Info("Enter The router ")
	err := pingServer()
	if err != nil {
		log.Fatal("The router has no response, or it might took too long to start up.", err)
	}
	log.Info("The router has been deployed successfully,Leave The router")
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, error := http.Get(viper.GetString("url") + "/v1/health")
		if error == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
