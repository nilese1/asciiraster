package main

import (
	"math/rand"
	"os"
	"os/signal"

	"github.com/nilese1/Ascii-Rasterizer/rasterizer"
	sc "github.com/nilese1/Ascii-Rasterizer/scene"
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

var (
	MODEL_TRANSLATION = vector.Vec3{X: 0, Y: 0, Z: 5}
	MODEL_SCALE       = 1.5
)

// meant as a method to demo the rasterizer
func SpinningObjects(scenes []sc.Scene) {
	for {
		for i := range scenes {
			rotateModel(scenes[i].Model, MODEL_TRANSLATION)
			screen := renderModel(scenes[i].Model, &scenes[i])
			PrintScene(screen)
		}
		ResetCursor(scenes)
	}
}

func Cleanup(scenes []sc.Scene) {
	showCursor()
	moveCursor(sc.GetTotalHeight(scenes), false)
}

// TODO: delete this
func randAngle(max_angle float64) float64 {
	rand_f := rand.Float64()

	return rand_f * max_angle
}

func rotateModel(model *rasterizer.Model, translation vector.Vec3) {
	rot_x := randAngle(0.05)
	rot_y := randAngle(0.05)
	rot_z := randAngle(0.05)

	inverse_translate := translation.Mul(-1)

	model.Translate(inverse_translate.X, inverse_translate.Y, inverse_translate.Z)
	model.Rotate(rot_x, rot_y, rot_z)
	model.Translate(translation.X, translation.Y, translation.Z)
}

// this needs some optimization
func renderModel(model *rasterizer.Model, scene *sc.Scene) [][]pixel {
	screen := make([][]pixel, scene.SceneHeight)
	for i := range screen {
		screen[i] = make([]pixel, scene.SceneWidth)
	}

	for y := range scene.SceneHeight {
		for x := range scene.SceneWidth {
			p := pixel{0, 0, 0, 0}

			hits, tri := scene.GetTriInPixel(x, y, model.Triangles)
			if hits {
				normal := tri.GetNormal()
				light := (1 + vector.Dot3(&scene.SunDir, &normal)) * 0.5
				brightness_adjusted := model.Colour.Scale(light)

				p.r = int(brightness_adjusted.GetR())
				p.g = int(brightness_adjusted.GetB())
				p.b = int(brightness_adjusted.GetG())

				p.light = light
			}

			screen[y][x] = p
		}
	}

	return screen
}

func main() {
	var scenes []sc.Scene

	model := rasterizer.ParseModel("models/cube")

	model.Scale(MODEL_SCALE)
	model.Translate(MODEL_TRANSLATION.X, MODEL_TRANSLATION.Y, MODEL_TRANSLATION.Z)

	scene := sc.CreateScene(model)

	scenes = append(scenes, *scene)

	defer Cleanup(scenes)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	hideCursor()
	go SpinningObjects(scenes)

	<-sig
}
