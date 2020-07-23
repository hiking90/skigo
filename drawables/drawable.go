package drawable

import (
    // "fmt"
    sk "github.com/hiking90/skigo"
)

var (
    XSCALE      float32 = 1.0
    YSCALE      float32 = 1.0
)

type IDrawable interface {
    Init()
    GetBound() *sk.Rect
    OnDraw(*sk.Canvas)
}

type Drawable struct {
    Bound       sk.Rect
}

func (self *Drawable) GetBound() *sk.Rect {
    return &self.Bound
}


type Text struct {
    Drawable
    Text        string
    FontFace    string
    Size        float32
    Color       sk.Color

    _Font       *sk.Font
    _Paint      *sk.Paint
    _TextBlob   *sk.TextBlob
    _TextBounds *sk.Rect
}

func (self *Text) Init() {
    self._Font = sk.NewFont()
    if self.FontFace != "" {
    }

    if self.Size != 0 {
        self._Font.SetSize(self.Size * YSCALE)
        // self._Font.SetScaleX(XSCALE)
    }

    self._TextBlob = sk.NewTextBlob(self.Text, self._Font)
    _, self._TextBounds = self._Font.MeasureText(self.Text, nil)
    self.Drawable.Bound.Right = self._TextBounds.Width()
    self.Drawable.Bound.Bottom = self._TextBounds.Height()
    self._Paint = sk.NewPaint()
    self._Paint.SetColor(self.Color)
}

func (self *Text) OnDraw(canvas *sk.Canvas) {
    _bounds := self.GetBound()
    self._Paint.SetColor(sk.ColorWHITE)
    canvas.DrawRect(_bounds, self._Paint)

    self._Paint.SetColor(self.Color)
    canvas.DrawTextBlob(self._TextBlob,
        _bounds.Left - self._TextBounds.Left,
        _bounds.Top - self._TextBounds.Top, self._Paint)
}