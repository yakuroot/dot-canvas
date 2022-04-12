package canvas

import (
	"github.com/Neoration/dot-canvas/src/config"
)

func SetDot(x, y int, color string) {
	canvas.SetHexColor(color)
	canvas.DrawRectangle(
		float64(x)*dotPerPixel, float64(y)*dotPerPixel, dotPerPixel, dotPerPixel)
	canvas.Fill()
}

func Init() {
	canvas.SetHexColor(BaseGray)
	canvas.DrawRectangle(
		0, 0, float64(config.CanvasWidth)*dotPerPixel, float64(config.CanvasHeigh)*dotPerPixel)
	canvas.Fill()
}
