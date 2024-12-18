package util

type Vector struct {
	X, Y int
}

func (v Vector) RotateOrigin90() Vector {
	return Vector{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vector) Rotate90(about Vector) Vector {
	return v.Sub(about).RotateOrigin90().Add(about)
}

func (left Vector) Add(right Vector) Vector {
	return Vector{
		X: left.X + right.X,
		Y: left.Y + right.Y,
	}
}

func (v Vector) Opposite() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector) ManhattanOrigin() int {
	return Abs(v.X) + Abs(v.Y)
}

func (left Vector) Sub(right Vector) Vector {
	return Vector{
		X: left.X - right.X,
		Y: left.Y - right.Y,
	}
}

func (left Vector) Mul(right int) Vector {
	return Vector{
		X: left.X * right,
		Y: left.Y * right,
	}
}

func (v Vector) Collapse() int {
	return v.X + v.Y
}

func (left Vector) VectorMul(right Vector) Vector {
	return Vector{
		X: left.X * right.X,
		Y: left.Y * right.Y,
	}
}

var (
	NORTH = Vector{ 0, -1 }
	EAST = Vector{ 1, 0 }
	SOUTH = Vector{ 0, 1 }
	WEST = Vector{ -1, 0 }
	NORTH_EAST = NORTH.Add(EAST)
	NORTH_WEST = NORTH.Add(WEST)
	SOUTH_EAST = SOUTH.Add(EAST)
	SOUTH_WEST = SOUTH.Add(WEST)
)

func Cardinals() []Vector {
	return []Vector{
		EAST,
		NORTH,
		WEST,
		SOUTH,
	}
}

func Diagonals() []Vector {
	return []Vector{
		SOUTH_EAST,
		NORTH_EAST,
		NORTH_WEST,
		SOUTH_WEST,
	}
}

func Units() []Vector {
	units := make([]Vector, 0, 8)

	units = append(units, Diagonals()...)
	units = append(units, Cardinals()...)

	return units
}
