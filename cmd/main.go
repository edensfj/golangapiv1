package main

import (
	"back_go/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.Routes(r)
	if err := http.ListenAndServe(":4000", r); err != nil {
		panic(err)
	}

}
