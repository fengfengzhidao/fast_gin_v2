package api

import (
	"fast_gin/api/image_api"
	"fast_gin/api/user_api"
)

type Api struct {
	UserApi  user_api.UserApi
	ImageApi image_api.ImageApi
}

var App = new(Api)
