package asciiraster

type Triangle struct {
	A Vec3
	B Vec3
	C Vec3

	normal_vec Vec3
}

func CreateTriangle(world_a Vec3, world_b Vec3, world_c Vec3, normal Vec3) Triangle {
	return Triangle{world_a, world_b, world_c, normal.Normalise()}
}

func (tri *Triangle) GetWorldCenter() Vec3 {
	x := (tri.A.X + tri.B.X + tri.C.X) / 3
	y := (tri.A.Y + tri.B.Y + tri.C.Y) / 3
	z := (tri.A.Z + tri.B.Z + tri.C.Z) / 3

	return Vec3{X: x, Y: y, Z: z}
}

func (tri *Triangle) GetNormal() Vec3 {
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
	translation := Vec3{X: x, Y: y, Z: z}

	new_a := Add(tri.A, translation)
	new_b := Add(tri.B, translation)
	new_c := Add(tri.C, translation)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}

func (tri *Triangle) Scale(scale_factor float64) Triangle {
	new_a := tri.A.Mul(scale_factor)
	new_b := tri.B.Mul(scale_factor)
	new_c := tri.C.Mul(scale_factor)

	return CreateTriangle(new_a, new_b, new_c, tri.normal_vec)
}
