package main

import (
	"math/rand"
	"os"
	"os/signal"

	ar "github.com/nilese1/asciiraster"
)

func randAngle(max_angle float64) float64 {
	rand_f := rand.Float64()

	return rand_f * max_angle
}

func rotateModel(model *ar.Model, translation ar.Vec3) {
	rot_x := randAngle(0.05)
	rot_y := randAngle(0.05)
	rot_z := randAngle(0.05)

	inverse_translate := translation.Mul(-1)

	model.Translate(inverse_translate.X, inverse_translate.Y, inverse_translate.Z)
	model.Rotate(rot_x, rot_y, rot_z)
	model.Translate(translation.X, translation.Y, translation.Z)
}

// meant as a method to demo the rasterizer
func spinningObjects(scenes []ar.Scene) {
	for {
		for i := range scenes {
			rotateModel(scenes[i].Model, ar.MODEL_TRANSLATION)
			screen := ar.RenderModel(scenes[i].Model, &scenes[i])
			ar.PrintScene(screen)
		}
		ar.ResetCursor(scenes)
	}
}

func main() {
	var scenes []ar.Scene

	model := ar.LoadObjFile("cube")

	scene := ar.CreateScene(model)

	scenes = append(scenes, *scene)

	defer ar.Cleanup(scenes)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ar.HideCursor()
	go spinningObjects(scenes)

	<-sig
}
