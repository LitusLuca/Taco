package window

var win IWindow

type WindowProps struct {
	Title         string
	Width, Height int
}

type CursorMode int

type IWindow interface {
	OnUpdate()
	GetSize() (width, height int)
	Destroy()
	IsFullScreen() bool
	ToggleFullScreen(full bool)
	IsVSync() bool
	SetVSync(VSync bool)
	SetCursorMode(mode CursorMode)
	GetTime() float32
}

func Get() IWindow {
	if win == nil {
		panic("window.Init() must be called first!")
	}
	return win
}
