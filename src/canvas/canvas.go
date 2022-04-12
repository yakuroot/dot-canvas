package canvas

import (
	"context"
	"fmt"
	"log"

	"github.com/Neoration/dot-canvas/src/config"
	"github.com/Neoration/dot-canvas/src/database"
	"github.com/Neoration/dot-canvas/src/model"
	"github.com/fogleman/gg"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	canvas      *gg.Context
	dotPerPixel = float64(5)
	BaseGray    = "#373F52"
)

func init() {
	defer log.Println("Canvas Loaded.")
	defer SaveImage()

	log.Println("Canvas Loading...")

	canvas = gg.NewContext(config.CanvasWidth*int(dotPerPixel), config.CanvasHeigh*int(dotPerPixel))
	canvas.SetHexColor(BaseGray)
	canvas.DrawRectangle(
		0, 0, float64(config.CanvasWidth)*dotPerPixel, float64(config.CanvasHeigh)*dotPerPixel)
	canvas.Fill()

	res, err := database.Canvas.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}

	var canvasDatas []model.CanvasProps
	if err := res.All(context.Background(), &canvasDatas); err != nil {
		fmt.Println(err)
		return
	}

	for _, data := range canvasDatas {
		canvas.SetHexColor(data.Color)
		canvas.DrawRectangle(
			float64(data.X)*dotPerPixel, float64(data.Y)*dotPerPixel, dotPerPixel, dotPerPixel)
		canvas.Fill()
	}
}
