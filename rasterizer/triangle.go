package rasterizer

import (
	"github.com/nilese1/Ascii-Rasterizer/vector"
)

type Triangle struct {
	A vector.Vec3
	B vector.Vec3
	C vector.Vec3

	normal_vec vector.Vec3
}

func CreateTriangle(world_a vector.Vec3, world_b vector.Vec3, world_c vector.Vec3, normal vector.Vec3) Triangle {
	return Triangle{world_a, world_b, world_c, normal.Normalise()}
}

func (tri *Triangle) GetWorldCenter() vector.Vec3 {
	x := (tri.A.X + tri.B.X + tri.C.X) / 3
	y := (tri.A.Y + tri.B.Y + tri.C.Y) / 3
	z := (tri.A.Z + tri.B.Z + tri.C.Z) / 3

	return vector.Vec3{X: x, Y: y, Z: z}
}

func (tri *Triangle) GetNormal() vector.Vec3 {
	return tri.normal_vec
}

func (tri *Triangle) Rotate(rot_x float64, rot_y float64, rot_z float64) Triangle {
	new_a := tri.A.Rotate(rot_x, rot_y, rot_z)
	new_b := tri.B.Rotate(rot_x, rot_y, rot_z)
	new_c := tri.C.Rotate(rot_x, rot_y, rot_z)
	new_normal := tri.normal_vec.Rotate(rot_x, rot_y, rot_z)

	return CreateTriangle(new_a, new_b, new_c, new_normal)
}

func (tri *Triangle) Translate(x float64, y float64, z float64) Triangle {
	translation := vector.Vec3{X: x, Y: y, Z: z}

	new_a := vector.Add(tri.A, translation)
	new_b := vector.Add(tri.B, translation)
	new_c := vector.Add(tri.C, translation)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}

func (tri *Triangle) Scale(scale_factor float64) Triangle {
	new_a := tri.A.Mul(scale_factor)
	new_b := tri.B.Mul(scale_factor)
	new_c := tri.C.Mul(scale_factor)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}
