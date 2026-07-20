package rasterizer

type Colour struct {
	r, g, b uint32
}

type Model struct {
	Triangles []Triangle
	Colour    *Colour
}

func clampColourVal(x uint32) uint32 {
	var minVal, maxVal uint32
	minVal, maxVal = 0, 255

	x = min(x, maxVal)
	x = max(x, minVal)

	return x
}

func (c *Colour) GetR() uint32 {
	return c.r
}

func (c *Colour) GetG() uint32 {
	return c.g
}

func (c *Colour) GetB() uint32 {
	return c.b
}

func (c *Colour) Scale(f float64) *Colour {
	r := uint32(float64(c.r) * f)
	g := uint32(float64(c.g) * f)
	b := uint32(float64(c.b) * f)

	return CreateColour(r, g, b)
}

func CreateColour(r uint32, g uint32, b uint32) *Colour {
	r = clampColourVal(r)
	g = clampColourVal(g)
	b = clampColourVal(b)

	return &Colour{r, g, b}
}

func (m *Model) Rotate(rot_x float64, rot_y float64, rot_z float64) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Rotate(rot_x, rot_y, rot_z)
	}
}

func (m *Model) Translate(x float64, y float64, z float64) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Translate(x, y, z)
	}
}

func (m *Model) Scale(scale_factor float64) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Scale(scale_factor)
	}
}
