package converter

import (
	"encoding/json"
	"image"
	"image/color"
	"image/gif"
	"math"
	"reflect"

	"github.com/nfnt/resize"
)

type video struct {
	Frames []frame `json:"Frames"`
}
type frame struct {
	Rows []string `json:"Rows"`
}

func ConvertGif(gifToConvert gif.GIF, width int, height int) (vidJson map[string]interface{}, err error) {

	//Build image slice
	var frames []image.Image

	for _, frame := range gifToConvert.Image {
		//Resize Image
		resizedFrame := resize.Resize(uint(width), uint(height), frame, resize.Lanczos3)

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
		for i := 0; i < newHeight; i++ {
			for j := 0; j < newWidth; j++ {
				//Get Color of pixel
				color := color.NRGBAModel.Convert(frames[frameIndex].At(j, i))
				r := reflect.ValueOf(color).FieldByName("R").Uint()
				g := reflect.ValueOf(color).FieldByName("G").Uint()
				b := reflect.ValueOf(color).FieldByName("B").Uint()
				a := reflect.ValueOf(color).FieldByName("A").Uint()

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
