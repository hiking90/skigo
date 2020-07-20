package skigo

/*
#cgo darwin,!ios LDFLAGS: -Lprebuilt/darwin.amd64 -lpng -lskia -lz -framework Cocoa -framework Foundation -framework CoreGraphics
#cgo linux,!android LDFLAGS: -lpng -lskia -lz -lfreetype -lfontconfig -lm -ldl
#cgo CXXFLAGS: -std=c++17 -Iskia.m85
#cgo CFLAGS: -Iskia.m85
*/
/*
#include <stdlib.h>
#include "skigo.h"
*/
import "C"

import (
    "unsafe"
)

type GrContext struct {
    skGrContext *C.sk_gr_context_t
}

type Surface struct {
    skSurface *C.sk_surface_t
}

type Canvas struct {
    skCanvas *C.sk_canvas_t
}

type Paint struct {
    skPaint *C.sk_paint_t
}

type Font struct {
    skFont *C.sk_font_t
}

type TextBlob struct {
    skTextBlob *C.sk_text_blob_t
}

type Rect struct {
    Left    float32
    Top     float32
    Right   float32
    Bottom  float32
}

type IntRect struct {
    Left    int32
    Top     int32
    Right   int32
    Bottom  int32
}

type Color = uint32


var (
    ColorBLACK      = SetARGB(0xFF, 0x00, 0x00, 0x00)
    ColorDKGRAY     = SetARGB(0xFF, 0x44, 0x44, 0x44)
    ColorGRAY       = SetARGB(0xFF, 0x88, 0x88, 0x88)
    ColorLTGRAY     = SetARGB(0xFF, 0xCC, 0xCC, 0xCC)
    ColorWHITE      = SetARGB(0xFF, 0xFF, 0xFF, 0xFF)
    ColorRED        = SetARGB(0xFF, 0xFF, 0x00, 0x00)
    ColorGREEN      = SetARGB(0xFF, 0x00, 0xFF, 0x00)
    ColorBLUE       = SetARGB(0xFF, 0x00, 0x00, 0xFF)
    ColorYELLOW     = SetARGB(0xFF, 0xFF, 0xFF, 0x00)
    ColorCYAN       = SetARGB(0xFF, 0x00, 0xFF, 0xFF)
    ColorMAGENTA    = SetARGB(0xFF, 0xFF, 0x00, 0xFF)
)

func SetARGB(a, r, g, b uint8) Color {
    return (uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
}

func GrContextMakeGL() *GrContext {
    return &GrContext{ skGrContext: C.sk_gr_context_make_gl() }
}

func (self *GrContext) Unref() {
    C.sk_gr_context_unref(self.skGrContext)
}

func MakeOnScreenGLSurface(context *GrContext, width, height int) *Surface {
    return &Surface{ skSurface: C.sk_surface_make_onscreen_gl(context.skGrContext, C.int(width), C.int(height)) }
}

func (self *Surface) Unref() {
    C.sk_surface_unref(self.skSurface)
}

func (self *Surface) GetCanvas() *Canvas {
    return &Canvas{ skCanvas: C.sk_surface_get_canvas(self.skSurface) }
}


func (self *Canvas) Scale(sx, sy float32) {
    C.sk_canvas_scale(self.skCanvas, C.float(sx), C.float(sy))
}

func (self *Canvas) Clear(color Color) {
    C.sk_canvas_clear(self.skCanvas, C.sk_color_t(color))
}

func (self *Canvas) DrawString(str string, x, y float32, font *Font, paint *Paint) {
    cstr := C.CString(str)
    C.sk_canvas_draw_string(self.skCanvas, cstr, C.float(x), C.float(y), font.skFont, paint.skPaint)
    C.free(unsafe.Pointer(cstr))
}

func (self *Canvas) DrawRect(rect *Rect, paint *Paint) {
    crect := &C.sk_rect_t{ C.float(rect.Left), C.float(rect.Top), C.float(rect.Right), C.float(rect.Bottom) }
    C.sk_canvas_draw_rect(self.skCanvas, crect, paint.skPaint)
}

func (self *Canvas) DrawTextBlob(blob *TextBlob, x, y float32, paint *Paint) {
    C.sk_canvas_draw_text_blob(self.skCanvas, blob.skTextBlob, C.float(x), C.float(y), paint.skPaint)
}

func (self *Canvas) Save() {
    C.sk_canvas_save(self.skCanvas)
}

func (self *Canvas) Restore() {
    C.sk_canvas_restore(self.skCanvas)
}

func (self *Canvas) Flush() {
    C.sk_canvas_flush(self.skCanvas)
}

func NewPaint() *Paint {
    return &Paint{ skPaint: C.sk_paint_new() }
}

func (self *Paint) Unref() {
    C.sk_paint_delete(self.skPaint)
    self.skPaint = nil
}

func (self *Paint) SetColor(color Color) {
    if self.skPaint != nil {
        C.sk_paint_set_color(self.skPaint, C.sk_color_t(color))
    }
}

func NewFont() *Font {
    return &Font{ skFont: C.sk_font_new() }
}

func (self *Font) SetSize(size float32) {
    C.sk_font_set_size(self.skFont, C.float(size))
}

func (self *Font) Unref() {
    C.sk_font_delete(self.skFont)
    self.skFont = nil
}

func NewTextBlob(str string, font *Font) *TextBlob {
    cstr := C.CString(str)
    blob := C.sk_text_blob_make_string(cstr, font.skFont)
    C.free(unsafe.Pointer(cstr))
    return &TextBlob{ skTextBlob: blob }
}

func (self *TextBlob) Bounds() *Rect {
    var left, right, top, bottom float32
    C.sk_text_blob_bounds(self.skTextBlob,
        (*C.float)(&left), (*C.float)(&top), (*C.float)(&right), (*C.float)(&bottom))
    return &Rect{ left, top, right, bottom}
}
