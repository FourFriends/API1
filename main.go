package main

import (
	"API1/config"
	"API1/router"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config","c","","apiserver config file path")
)



func main()  {
	pflag.Parse()

	err := config.Init(*cfg)
	if err != nil{
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	router.Load(g)

	http.ListenAndServe(viper.GetString("port"),g)

	go check()
}

//自检程序

func check(){
	err :=pingServer()
	if err != nil{
		log.Fatal("The router has no response")
	}
	log.Println("The router has been deployed successfully.")
}


func pingServer() (error){
	for i:=0; i<viper.GetInt("max_ping_count"); i++ {
		resp,error := http.Get(viper.GetString("url") + "/v1/health")
		if error == nil && resp.StatusCode == 200{
			return nil
		}
		log.Println("waiting for the router,retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
