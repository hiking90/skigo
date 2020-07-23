package app

import (
    // "fmt"
    "runtime"

    "github.com/go-gl/glfw/v3.3/glfw"
    "github.com/go-gl/gl/v3.3-core/gl"

    sk "github.com/hiking90/skigo"
    skw "github.com/hiking90/skigo/widgets"
    skd "github.com/hiking90/skigo/drawables"
)

func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

type App struct {
    Width, Height   int
    Title           string
    Body            skw.IWidget
}

func Init() {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
}

func RunApp(app *App) {
    mainWin, err := glfw.CreateWindow(app.Width, app.Height, app.Title, nil, nil)
    if err != nil {
        panic(err)
    }

    skd.XSCALE, skd.YSCALE = mainWin.GetContentScale()

    mainWin.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        panic(err)
    }

    fw, fh := mainWin.GetFramebufferSize()

    glContext := sk.GrContextMakeGL()
    defer glContext.Unref()

    mainSurface := sk.MakeOnScreenGLSurface(glContext, fw, fh)
    defer mainSurface.Unref()

    mainCanvas := mainSurface.GetCanvas()

    if app.Body == nil {
        panic("Application Body is nil.")
    }

    app.Body.Init()
    app.Body.Relayout(sk.NewIntRectXYWH(0, 0, int32(fw), int32(fh)))

    for !mainWin.ShouldClose() {
        mainCanvas.Save()
        app.Body.OnDrawBackground(mainCanvas)
        app.Body.OnDrawContent(mainCanvas)
        mainCanvas.Restore()

        mainCanvas.Flush()

        mainWin.SwapBuffers()
        glfw.WaitEvents()
    }

    glfw.Terminate()
}