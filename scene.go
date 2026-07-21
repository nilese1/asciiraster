package asciiraster

type Scene struct {
	SceneWidth      uint32
	SceneHeight     uint32
	CamFOV          float32
	ViewPlaneHeight float64

	SunDir Vec3
	Model  *Model
}


func (scene *Scene) ConvertToScreen(point Vec3) Vec2 {
	view_plane_point := convertToViewPlane(point, scene.ViewPlaneHeight)

	scale := float64(scene.SceneHeight) / 2

	scaled_y := view_plane_point.Y * scale
	scaled_x := view_plane_point.X * scale

	offset_y := float64(scene.SceneHeight)/2 - scaled_y
	offset_x := float64(scene.SceneWidth)/2 + scaled_x

	return Vec2{X: offset_x, Y: offset_y}
}

func (scene *Scene) GetTriInPixel(pixel_x uint32, pixel_y uint32, tris []Triangle) (bool, Triangle) {
	point := Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	hit := false
	var nearest_tri Triangle

	for _, i := range tris {
		if scene.isPointInTri(point, &i) {
			if !hit || nearest_tri.GetWorldCenter().Z < nearest_tri.GetWorldCenter().Z {
				nearest_tri = i
			}

			hit = true
		}
	}

	return hit, nearest_tri
}

func GetTotalHeight(scenes []Scene) uint32 {
	var totalHeight uint32

	for _, scene := range scenes {
		totalHeight += scene.SceneHeight
	}

	return totalHeight
}

func (scene *Scene) isPointInTri(point Vec2, tri *Triangle) bool {
	screen_a := scene.ConvertToScreen(tri.A)
	screen_b := scene.ConvertToScreen(tri.B)
	screen_c := scene.ConvertToScreen(tri.C)

	ap := Sub(point, screen_a)
	bp := Sub(point, screen_b)
	cp := Sub(point, screen_c)

	ab := Sub(screen_b, screen_a)
	bc := Sub(screen_c, screen_b)
	ca := Sub(screen_a, screen_c)

	ab_out := ab.Rot90()
	bc_out := bc.Rot90()
	ca_out := ca.Rot90()

	return VecsSameDir(ap, ab_out) && VecsSameDir(bp, bc_out) && VecsSameDir(cp, ca_out)
}

func convertToViewPlane(point Vec3, viewPlaneHeight float64) Vec2 {
	view_plane_y := point.Y / point.Z
	fraction_up_plane := view_plane_y / viewPlaneHeight

	view_plane_x := point.X / point.Z * 2 //multiply by 2 because ASCII char width is half of char height
	fraction_along_plane := view_plane_x / viewPlaneHeight

	return Vec2{X: fraction_along_plane, Y: fraction_up_plane}
}
