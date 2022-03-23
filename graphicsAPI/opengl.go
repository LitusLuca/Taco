package graphicsapi

import (
	"log"
)

type OpenGLGraphicsAPI struct {
	test int
}

func newOpenGlAPI() *OpenGLGraphicsAPI {
	api := new(OpenGLGraphicsAPI)
	api.Init()
	gApi = api
	return api
}

func (api *OpenGLGraphicsAPI) Init() {
	api.test = 0
	log.Println(api.test)
}

func (api *OpenGLGraphicsAPI) DrawTest() {
	log.Println(api.test)
}
