package skigo

import (
    "fmt"
    "testing"
    "runtime"
    "github.com/go-gl/glfw/v3.3/glfw"
    "github.com/go-gl/gl/v3.3-core/gl"
)

const (
    WIDTH   = 640
    HEIGHT  = 480
)

func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func TestMain(m *testing.M) {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    glfw.WindowHint(glfw.Resizable, glfw.False)
    window, err := glfw.CreateWindow(WIDTH, HEIGHT, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        panic(err)
    }

    fw, fh := window.GetFramebufferSize()

    context := GrContextMakeGL()
    defer context.Unref()
    surface := MakeOnScreenGLSurface(context, fw, fh)
    defer surface.Unref()

    canvas := surface.GetCanvas()
    canvas.Scale(1.0, 1.0)

    paint := NewPaint()

    font := NewFont()
    font.SetSize(120)

    textBlob := NewTextBlob("Hello World", font)
    textRect := textBlob.Bounds()

    for !window.ShouldClose() {
        fw, fh := window.GetFramebufferSize()
        gl.Viewport(0, 0, int32(fw), int32(fh))
        fmt.Printf("Window Framebuffer size: %d, %d\n", fw, fh)

        canvas.Clear(ColorWHITE)
        paint.SetColor(ColorRED)
        canvas.DrawRect(&Rect{0, 0, float32(fw)/2, float32(fh)/2}, paint)
        paint.SetColor(ColorBLACK)

        canvas.DrawTextBlob(textBlob,
            (float32(fw) - textRect.Right - textRect.Left)/2,
            (float32(fh) - textRect.Bottom - textRect.Top)/2, paint)

        // canvas.DrawString("Hello", float32(fw)/2, float32(fh)/2, font, paint)

        canvas.Flush();

        window.SwapBuffers()
        glfw.WaitEvents()
    }
}