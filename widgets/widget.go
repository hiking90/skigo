package widget

import (
    "fmt"
    sk "github.com/hiking90/skigo"
    skd "github.com/hiking90/skigo/drawables"
)

type IWidget interface {
    Init()
    GetBound() *sk.IntRect
    Relayout(*sk.IntRect) *sk.IntRect

    OnDrawBackground(*sk.Canvas)
    OnDrawContent(*sk.Canvas)
}

type Widget struct {
    BackgroundColor     sk.Color
    Bound               sk.IntRect
}

func (self *Widget) Init() {
}

func (self *Widget) GetBound() *sk.IntRect {
    return &self.Bound
}

func (self *Widget) Relayout(rect *sk.IntRect) *sk.IntRect {
    return rect
}

func (self *Widget) OnDrawBackground(canvas *sk.Canvas) {
    canvas.Clear(self.BackgroundColor)
}

func (self *Widget) OnDrawContent(canvas *sk.Canvas) {
    fmt.Println("Widget.OnDrawContent")
}

type Center struct {
    Widget
    Body    IWidget
}

func (self *Center) Init() {
    if self.Body == nil {
        panic("widget.Center Body is nil.")
    }
    self.Body.Init()
}

func (self *Center) Relayout(bound *sk.IntRect) *sk.IntRect {
    _bound := self.Body.GetBound()

    x := bound.Left + (bound.Width() - _bound.Width()) / 2
    y := bound.Top + (bound.Height() - _bound.Height()) / 2

    fmt.Printf("%d, %d, %d, %d\n", x, y, _bound.Width(), _bound.Height())

    self.Body.Relayout(sk.NewIntRectXYWH(x, y, _bound.Width(), _bound.Height()))
    return bound
}

func (self *Center) OnDrawContent(canvas *sk.Canvas) {
    self.Body.OnDrawContent(canvas)
}

type Drawable struct {
    Widget
    Body    skd.IDrawable
}

func (self *Drawable) Init() {
    if self.Body == nil {
        panic("widget.Drawable Body is nil.")
    }

    self.Body.Init()
    bound := self.GetBound()
    bound.SetRect(self.Body.GetBound())
    b := self.Body.GetBound()
    fmt.Printf("Drawable.Init %f, %f, %d, %d\n", b.Width(), b.Height(), bound.Width(), bound.Height())
}

func (self *Drawable) Relayout(bound *sk.IntRect) *sk.IntRect {
    _bound := self.Body.GetBound()
    _bound.SetIntRect(bound)
    *self.GetBound() = *bound
    fmt.Printf("Drawable.Relayout %f, %f\n", _bound.Width(), _bound.Height())
    return bound
}

func (self *Drawable) OnDrawContent(canvas *sk.Canvas) {
    self.Body.OnDraw(canvas)
}
