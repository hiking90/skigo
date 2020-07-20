#ifndef __SKIGO_H__
#define __SKIGO_H__

#include <include/c/sk_types.h>
#include <include/c/sk_paint.h>
#include <include/c/sk_surface.h>
#include <include/c/sk_canvas.h>
#include <include/c/sk_paint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct sk_font_t sk_font_t;
typedef struct sk_text_blob_t sk_text_blob_t;

typedef struct sk_gr_context_t sk_gr_context_t;
sk_gr_context_t * sk_gr_context_make_gl();
void sk_gr_context_unref(sk_gr_context_t*);

sk_surface_t *sk_surface_make_onscreen_gl(sk_gr_context_t *context, int width, int height);

void sk_canvas_flush(sk_canvas_t *canvas);
void sk_canvas_clear(sk_canvas_t *canvas, sk_color_t color);
void sk_canvas_draw_string(sk_canvas_t *canvas, const char *str, float x, float y,
    sk_font_t *font, sk_paint_t *paint);
void sk_canvas_draw_text_blob(sk_canvas_t *canvas, const sk_text_blob_t *blob, float x, float y,
    sk_paint_t *paint);

sk_text_blob_t * sk_text_blob_make_string(const char *text, sk_font_t *font);
void sk_text_blob_bounds(sk_text_blob_t *, float *left, float *top, float *right, float *bottom);


sk_font_t * sk_font_new();
void sk_font_set_size(sk_font_t *font, float size);
void sk_font_delete(sk_font_t *font);

#ifdef __cplusplus
}
#endif

#endif