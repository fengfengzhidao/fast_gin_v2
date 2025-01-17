package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/utils/jwts"
	"fmt"
)

func main() {
	core.InitLogger()
	flags.Parse()
	global.Config = core.ReadConfig()
	token, err := jwts.SetToken(jwts.Claims{
		UserID: 1,
		RoleID: 1,
	})
	fmt.Println(token, err)
	//claims, err := jwts.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsImlzcyI6ImZlbmdmZW5nIiwiZXhwIjoxNzM3MDYwNzY2fQ.MQ-uiLjU0UArChYUPx2EjqSWZYkpvNRdhk26LwaOjo1")
	//fmt.Println(claims, err)
}
