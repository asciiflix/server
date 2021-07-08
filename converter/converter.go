package converter

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"reflect"

	"github.com/nfnt/resize"
)

type video struct {
	Frames []frame `json:"Frames"`
}
type frame struct {
	Rows   []string `json:"Rows"`
	Colors []string `json:"Color"`
	Delay  int      `json:"Delay"`
}

func ConvertGif(gifToConvert gif.GIF, width int, height int) (vidJson map[string]interface{}, err error) {

	imgWidth, imgHeight := getGifDimensions(&gifToConvert)
	var delay = gifToConvert.Delay
	//Get default image
	priorFrame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(priorFrame, priorFrame.Bounds(), gifToConvert.Image[0], image.Point{X: 0, Y: 0}, draw.Src)

	//Build image slice
	var frames []image.Image

	for _, frame := range gifToConvert.Image {
		//Draw over priorFrame
		draw.Draw(priorFrame, priorFrame.Bounds(), frame, image.Point{X: 0, Y: 0}, draw.Over)
		//Init actualFrame
		actualFrame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		//Add to actualFrame
		draw.Draw(actualFrame, actualFrame.Bounds(), priorFrame, image.Point{X: 0, Y: 0}, draw.Over)

		//Resize Image
		resizedFrame := resize.Resize(uint(width), uint(height), actualFrame, resize.Lanczos3)

		//Add to slice
		frames = append(frames, resizedFrame)
	}

	video := video{}
	video.Frames = make([]frame, len(frames))

	chars := []byte(" .,:;i1tfLCG08@")

	for frameIndex := 0; frameIndex < len(frames); frameIndex++ {
		//Get size of frame
		size := frames[frameIndex].Bounds()
		newWidth := size.Max.X
		newHeight := size.Max.Y
		video.Frames[frameIndex].Rows = make([]string, (newHeight))
		video.Frames[frameIndex].Colors = make([]string, (newHeight))

		//Save delay
		video.Frames[frameIndex].Delay = delay[frameIndex]

		for i := 0; i < newHeight; i++ {
			for j := 0; j < newWidth; j++ {
				//Get Color of pixel
				color := color.NRGBAModel.Convert(frames[frameIndex].At(j, i))
				r := reflect.ValueOf(color).FieldByName("R").Uint()
				g := reflect.ValueOf(color).FieldByName("G").Uint()
				b := reflect.ValueOf(color).FieldByName("B").Uint()
				a := reflect.ValueOf(color).FieldByName("A").Uint()

				//Save color
				video.Frames[frameIndex].Colors[i] += fmt.Sprintf("rgba(%d,%d,%d,%d);", r, g, b, a)

				//Calculate color intensity
				intensity := (r + g + b) * a / 255

				//Map color intensity to char and add it to row
				step := float64(255 * 3 / (len(chars) - 1))
				video.Frames[frameIndex].Rows[i] += getChar(chars[roundValue(float64(intensity)/step)])
			}
		}
	}

	//Generate json
	vid, err := json.Marshal(video)
	if err != nil {
		return nil, err
	}

	vidJson = make(map[string]interface{})
	err = json.Unmarshal(vid, &vidJson)
	if err != nil {
		return nil, err
	}

	return
}

func getChar(input byte) string {
	return string(input)
}
func roundValue(value float64) int {
	return int(math.Floor(value + 0.5))
}

func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int

	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}
