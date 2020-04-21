package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aimpsmile/x-apisix-example/handler"
	"github.com/micro/go-micro/v2/web"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

func startWeb(path string, h handlerFunc, options ...web.Option) {
	//  配置 web service
	service := web.NewService(options...)
	//	设置 http handler
	service.HandleFunc(path, h)
	service.HandleFunc("/stats", handler.StatusHandler())
	service.HandleFunc("/health", handler.HealthHandler())

	//	服务初始化
	if err := service.Init(); err != nil {
		log.Panicf("WEB init err: %+v", err)
	}
	// 运行服务
	if err := service.Run(); err != nil {
		log.Panicf("WEB run err: %+v", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ab := map[string]string{
		"code": "200",
		"msg":  "success",
	}
	data, _ := json.Marshal(ab)
	w.Write(data)
}

func main() {

	op := []web.Option{
		web.Name("aimgo.web.v1.test"),
		web.Version("v1.22.33"),
	}
	//	curl http://web.uqudu.com:8888/v1/test/hello
	//	由于demo配置会自动增加 v1/test 前辍，请根据自己的场景灵活的配置模板
	startWeb("/v1/test/hello", Hello, op...)
}
