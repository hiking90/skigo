#define SK_RELEASE

#include <src/gpu/gl/GrGLUtil.h>
#include <include/gpu/gl/GrGLInterface.h>
#include <include/gpu/GrContext.h>
#include <include/core/SkSurface.h>
#include <include/core/SkCanvas.h>
#include <include/core/SkTextBlob.h>

#include "skigo.h"

sk_gr_context_t * sk_gr_context_make_gl() {
    // // setup GrContext
    auto interface = GrGLMakeNativeInterface();
    // // setup contexts
    sk_sp<GrContext> grContext(GrContext::MakeGL(interface));
    SkASSERT(grContext);

    return (sk_gr_context_t *) grContext.release();
}

void sk_gr_context_unref(sk_gr_context_t *context) {
    SkSafeUnref((GrContext*)context);
}

sk_surface_t *sk_surface_make_onscreen_gl(sk_gr_context_t *context, int width, int height) {
    // setup GrContext
    auto interface = GrGLMakeNativeInterface();

    // Wrap the frame buffer object attached to the screen in a Skia render target so Skia can
    // render to it
    GrGLint buffer;
    GR_GL_GetIntegerv(interface.get(), GR_GL_FRAMEBUFFER_BINDING, &buffer);

    GrGLFramebufferInfo info;
    info.fFBOID = (GrGLuint) buffer;
    info.fFormat = GR_GL_RGBA8;

    SkColorType colorType;
    colorType = kRGBA_8888_SkColorType;

    GrGLint stencil;
    GR_GL_GetIntegerv(interface.get(), GR_GL_STENCIL_BITS, &stencil);

    GrBackendRenderTarget target(width, height, 0, stencil, info);

    // setup SkSurface
    // To use distance field text, use commented out SkSurfaceProps instead
    // SkSurfaceProps props(SkSurfaceProps::kUseDeviceIndependentFonts_Flag,
    //                      SkSurfaceProps::kLegacyFontHost_InitType);
    SkSurfaceProps props(SkSurfaceProps::kLegacyFontHost_InitType);

    sk_sp<SkSurface> surface(SkSurface::MakeFromBackendRenderTarget((GrContext*) context, target,
                                                                    kBottomLeft_GrSurfaceOrigin,
                                                                    colorType, nullptr, &props));
    return (sk_surface_t *) surface.release();
}

static inline SkCanvas* AsCanvas(sk_canvas_t* ccanvas) {
    return reinterpret_cast<SkCanvas*>(ccanvas);
}

static inline SkFont* AsFont(sk_font_t* cfont) {
    return reinterpret_cast<SkFont*>(cfont);
}

static inline SkTextBlob* AsTextBlob(sk_text_blob_t* cblob) {
    return reinterpret_cast<SkTextBlob*>(cblob);
}

static inline SkPaint* AsPaint(sk_paint_t* cpaint) {
    return reinterpret_cast<SkPaint*>(cpaint);
}

void sk_canvas_flush(sk_canvas_t *canvas) {
    AsCanvas(canvas)->flush();
}

void sk_canvas_clear(sk_canvas_t *canvas, sk_color_t color) {
    AsCanvas(canvas)->clear((SkColor) color);
}

void sk_canvas_draw_string(sk_canvas_t *canvas, const char *str, float x, float y,
    sk_font_t *font, sk_paint_t *paint) {
    AsCanvas(canvas)->drawString(str, x, y, *AsFont(font), *AsPaint(paint));
}

void sk_canvas_draw_text_blob(sk_canvas_t *canvas, const sk_text_blob_t *blob, float x, float y,
    sk_paint_t *paint) {
    AsCanvas(canvas)->drawTextBlob(AsTextBlob((sk_text_blob_t *) blob), x, y, *AsPaint(paint));
}

sk_text_blob_t * sk_text_blob_make_string(const char *text, sk_font_t *font) {
    return (sk_text_blob_t *) SkTextBlob::MakeFromString(text, *AsFont(font)).release();
}

void sk_text_blob_bounds(sk_text_blob_t *blob, sk_rect_t *rect) {
    SkRect r = AsTextBlob(blob)->bounds();
    rect->left = r.fLeft;
    rect->top = r.fTop;
    rect->right = r.fRight;
    rect->bottom = r.fBottom;
}


sk_font_t * sk_font_new() {
    return (sk_font_t *) new SkFont();
}

void sk_font_set_size(sk_font_t *font, float size) {
    AsFont(font)->setSize(size);
}

void sk_font_set_scalex(sk_font_t *font, float scale) {
    AsFont(font)->setScaleX(scale);
}

float sk_font_measure_text(sk_font_t *font, const char *text, int len, sk_rect_t *bounds, sk_paint_t *paint) {
    return AsFont(font)->measureText(text, len, SkTextEncoding::kUTF8, (SkRect *) bounds, AsPaint(paint));
}


void sk_font_delete(sk_font_t *font) {
    delete AsFont(font);
}
