package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok/biz/dal"
	"tiktok/pkg/constants"
)

func main() {
	h := server.New(server.WithHostPorts(constants.SERVEER_URL))

	dal.InitDB()

	h.Spin()
}
