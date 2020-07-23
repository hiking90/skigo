package skigo

type IntRect struct {
    Left    int32
    Top     int32
    Right   int32
    Bottom  int32
}

func NewIntRectXYWH(x, y, w, h int32) *IntRect {
    return &IntRect{ x, y, x + w, y + h }
}

func (self *IntRect) Width() int32 {
    return self.Right - self.Left
}

func (self *IntRect) Height() int32 {
    return self.Bottom - self.Top
}

func (self *IntRect) SetRect(rect *Rect) {
    self.Left = int32(rect.Left)
    self.Top = int32(rect.Top)
    self.Right = self.Left + int32(rect.Width())
    self.Bottom = self.Top + int32(rect.Height())
}


type Rect struct {
    Left    float32
    Top     float32
    Right   float32
    Bottom  float32
}

func NewRectXYWH(x, y, w, h float32) *Rect {
    return &Rect{ x, y, x + w, y + h }
}

func (self *Rect) Width() float32 {
    return self.Right - self.Left
}

func (self *Rect) Height() float32 {
    return self.Bottom - self.Top
}

func (self *Rect) SetIntRect(rect *IntRect) {
    self.Left = float32(rect.Left)
    self.Top = float32(rect.Top)
    self.Right = float32(rect.Right)
    self.Bottom = float32(rect.Bottom)
}
