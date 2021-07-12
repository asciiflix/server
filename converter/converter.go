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

	var delay = gifToConvert.Delay
	var frames []image.Image

	imgWidth, imgHeigt := getGifDimensions(&gifToConvert)
	startFrame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeigt))

	draw.Draw(startFrame, startFrame.Bounds(), gifToConvert.Image[0], image.Point{X: 0, Y: 0}, draw.Src)

	frames = append(frames, startFrame)

	for i := 1; i < len(gifToConvert.Image); i++ {

		switch gifToConvert.Disposal[i] {
		case gif.DisposalBackground:
			frames = append(frames, getFrameRTB(frames[i-1], gifToConvert.Image[i-1], gifToConvert.Image[i]))
		case gif.DisposalPrevious:
			frames = append(frames, getFrameRTP(frames, gifToConvert.Disposal, gifToConvert.Image[i]))
		case gif.DisposalNone:
			frames = append(frames, getFrameDND(frames[i-1], gifToConvert.Image[i]))
		default:
			frames = append(frames, getFrameDND(frames[i-1], gifToConvert.Image[i]))
		}
	}
	for i := 0; i < len(frames); i++ {
		frames[i] = resize.Resize(uint(width), uint(height), frames[i], resize.Lanczos3)
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

//Do Not Dispose
func getFrameDND(previousFrame image.Image, nextFrame image.Image) image.Image {

	//Generate canvas
	canvas := image.NewRGBA(image.Rect(0, 0, previousFrame.Bounds().Max.X, previousFrame.Bounds().Max.Y))
	//Set last frame as base
	draw.Draw(canvas, canvas.Bounds(), previousFrame, image.Point{X: 0, Y: 0}, draw.Src)
	//Overdraw next frame
	draw.Draw(canvas, canvas.Bounds(), nextFrame, image.Point{X: 0, Y: 0}, draw.Over)

	return canvas
}

//Restore To Previous
func getFrameRTP(frames []image.Image, disposal []byte, nextFrame image.Image) image.Image {

	var canvas *image.RGBA
	//Search last frame without DisposalPrevious
	for i := len(frames) - 1; i > -1; i-- {
		if disposal[i] != gif.DisposalPrevious {
			canvas = image.NewRGBA(image.Rect(0, 0, frames[i].Bounds().Max.X, frames[i].Bounds().Max.Y))
			//Set last frame as base
			draw.Draw(canvas, canvas.Bounds(), frames[i], image.Point{X: 0, Y: 0}, draw.Src)

		}
	}
	//Overdraw next frame
	draw.Draw(canvas, canvas.Bounds(), nextFrame, image.Point{X: 0, Y: 0}, draw.Over)
	return canvas
}

//Restore to Background
func getFrameRTB(previousFrame image.Image, previousOverlay image.Image, nextFrame image.Image) image.Image {
	//Generate canvas
	canvas := image.NewRGBA(image.Rect(0, 0, previousFrame.Bounds().Max.X, previousFrame.Bounds().Max.Y))
	//Set last frame as base
	draw.Draw(canvas, canvas.Bounds(), previousFrame, image.Point{X: 0, Y: 0}, draw.Src)
	//Clear last changes to transparent.
	transparent := image.NewRGBA(image.Rect(previousOverlay.Bounds().Min.X, previousOverlay.Bounds().Min.Y, previousOverlay.Bounds().Max.X, previousOverlay.Bounds().Max.Y))
	draw.Draw(canvas, canvas.Bounds(), transparent, image.Point{X: 0, Y: 0}, draw.Src)

	//Overdraw next frame
	draw.Draw(canvas, canvas.Bounds(), nextFrame, image.Point{X: 0, Y: 0}, draw.Over)

	return canvas

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
