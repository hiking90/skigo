package app

import (
    "testing"

    sk "github.com/hiking90/skigo"
    skw "github.com/hiking90/skigo/widgets"
    skd "github.com/hiking90/skigo/drawables"
)

func TestMain(m *testing.M) {
    Init()
    RunApp(&App {
        Width: 640,
        Height: 480,
        Title: "Skigo Application Test",
        Body: &skw.Center {
            Widget: skw.Widget {
                BackgroundColor: sk.ColorRED,
            },
            Body: &skw.Drawable {
                Body: &skd.Text {
                    Text: "Hello World gj",
                    Size: 30,
                    Color: sk.ColorBLACK,
                },
            },
        },
    })
}