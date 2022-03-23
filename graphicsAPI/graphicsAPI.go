package graphicsapi

import "fmt"

var gApi IGraphicsAPI

const (
	OpenGLAPI API = 0
	OtherAPI  API = 1
)

type API uint

type IGraphicsAPI interface {
	Init()
	DrawTest()
}

func NewAPI(api API) (IGraphicsAPI, error) {
	if gApi != nil {
		return nil, fmt.Errorf("graphics api is already initialized")
	}
	switch api {
	case OpenGLAPI:
		i := IGraphicsAPI(newOpenGlAPI())
		gApi = i
		return i, nil
	default:
		return nil, fmt.Errorf("no api specified")
	}

}
