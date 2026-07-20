package main

import (
	"math"
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

func triInPixel(scene *sc.Scene, pixel_x uint32, pixel_y uint32, tris []rasterizer.Triangle) (bool, rasterizer.Triangle) {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	hit := false
	var nearest_tri rasterizer.Triangle

	for _, i := range tris {
		if scene.IsPointInTri(point, &i) {
			if !hit || nearest_tri.GetWorldCenter().Z < nearest_tri.GetWorldCenter().Z {
				nearest_tri = i
			}

			hit = true
		}
	}

	return hit, nearest_tri
}

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

			hits, tri := triInPixel(scene, x, y, model.Triangles)
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

func spinningObjects(scenes []sc.Scene) {
	for {
		for i := range scenes {
			rotateModel(scenes[i].Model, MODEL_TRANSLATION)
			screen := renderModel(scenes[i].Model, &scenes[i])
			printScene(screen)
		}
		resetCursor(scenes)
	}
}

func cleanup(scenes []sc.Scene) {
	showCursor()
	moveCursor(sc.GetTotalHeight(scenes), false)
}

func main() {
	var scenes []sc.Scene
	var models []rasterizer.Model

	model := rasterizer.ParseModel("models/cube")

	model.Scale(MODEL_SCALE)
	model.Translate(MODEL_TRANSLATION.X, MODEL_TRANSLATION.Y, MODEL_TRANSLATION.Z)

	models = append(models, model)

	scene := sc.Scene{
		SceneWidth:      100,
		SceneHeight:     20,
		CamFOV:          math.Pi / 3,
		ViewPlaneHeight: math.Tan(math.Pi / 6),

		SunDir: vector.Vec3{X: 1, Y: 0.2, Z: -0.3}.Normalise(),
		Model:  &models[0],
	}

	scenes = append(scenes, scene)

	defer cleanup(scenes)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	hideCursor()
	go spinningObjects(scenes)

	<-sig
}
