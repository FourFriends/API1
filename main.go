package main

import (
	"API1/router"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
	"github.com/spf13/pflag"
)

func main()  {
	g := gin.New()

	router.Load(g)

	g.Run()

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
	for i:=0; i<2; i++ {
		resp,error := http.Get("http://127.0.0.1:8080"+"/v1/health")
		if error == nil && resp.StatusCode == 200{
			return nil
		}
		log.Println("waiting for the router,retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
