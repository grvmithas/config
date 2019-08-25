package main

import (
	"io/ioutil"
	"net/http"
	"personal/config/controller"
	"personal/config/domain"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := domain.Config{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}

	// set router
	router := httprouter.New()
	c := &controller.Controller{}
	c.Config = &config
	router.GET("/readConfig/:serviceName", c.ReadRequest)
	http.ListenAndServe(":8081", router)

}
