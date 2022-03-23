package window

import (
	"fmt"
	"runtime"

	graphicsapi "github.com/LitusLuca/Taco/graphicsAPI"

	"github.com/go-gl/gl/v4.6-core/gl"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	CursorNormal   = CursorMode(glfw.CursorNormal)
	CursorHidden   = CursorMode(glfw.CursorHidden)
	CursorDisabled = CursorMode(glfw.CursorDisabled)
)

type GlfwWindow struct {
	*glfw.Window

	graphicsAPI   graphicsapi.IGraphicsAPI
	title         string
	width, height int
	fullScreen    bool
	posX, posY    int
	vSync         bool
}

func Init(props WindowProps) error {
	if win != nil {
		panic(fmt.Errorf("can only init one window"))
	}

	runtime.LockOSThread()

	w := new(GlfwWindow)

	w.title = props.Title
	w.width = props.Width
	w.height = props.Height

	var err error

	err = glfw.Init()
	if err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Samples, 8)
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	w.Window, err = glfw.CreateWindow(props.Width, props.Height, props.Title, nil, nil)
	if err != nil {
		return err
	}
	w.MakeContextCurrent()

	w.graphicsAPI, err = graphicsapi.NewAPI(graphicsapi.OpenGLAPI)

	if err != nil {
		return err
	}

	err = gl.Init()

	if err != nil {
		return err
	}

	//TODO: set event callbacks

	win = w
	return nil
}

func (w *GlfwWindow) OnUpdate() {
	w.graphicsAPI.DrawTest()
	glfw.PollEvents()
	glfw.GetCurrentContext().SwapBuffers()
}

func (w *GlfwWindow) GetSize() (int, int) {
	return w.width, w.height
}

func (w *GlfwWindow) Destroy() {
	w.Window.Destroy()
}

func (w *GlfwWindow) IsFullScreen() bool {
	return w.fullScreen
}

func (w *GlfwWindow) ToggleFullScreen(fullScreen bool) {
	if w.fullScreen == fullScreen {
		return
	}
	if fullScreen {
		w.posX, w.posY = w.GetPos()
		monitor := glfw.GetPrimaryMonitor()
		videomode := monitor.GetVideoMode()
		w.SetMonitor(monitor, 0, 0, videomode.Width, videomode.Height, videomode.RefreshRate)
	} else {
		w.SetMonitor(nil, w.posX, w.posY, w.width, w.height, glfw.DontCare)
	}
	w.fullScreen = fullScreen
}

func (w *GlfwWindow) IsVSync() bool {
	return w.vSync
}

func (w *GlfwWindow) SetVSync(vsync bool) {
	if vsync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}
	w.vSync = vsync
}

func (w *GlfwWindow) SetCursorMode(mode CursorMode) {
	glfw.GetCurrentContext().SetInputMode(glfw.CursorMode, int(mode))
}

func (w *GlfwWindow) GetTime() float32{
	return float32(glfw.GetTime())
}