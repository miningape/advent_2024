package util

type Matrix [2][2]int

func LinearEquations(first Vector, second Vector) Matrix {
	return Matrix{ 
		{ first.X, second.X }, 
		{ first.Y, second.Y },
	}
}

func (matrix Matrix) Determinant() int {
	return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
}

func (matrix Matrix) Inverse() (int, Matrix) {
	return matrix.Determinant(), Matrix{
		{ matrix[1][1], -matrix[0][1] },
		{ -matrix[1][0], matrix[0][0] },
	}
}

func (matrix Matrix) Mul(vector Vector) Vector {
	return Vector{
		X: (vector.X * matrix[0][0]) + (vector.Y * matrix[0][1]),
		Y: (vector.X * matrix[1][0]) + (vector.Y * matrix[1][1]),
	}
}

func (matrix Matrix) SolveFor(solution Vector) (Vector, bool) {
	determinant, inverse := matrix.Inverse()

	if determinant == 0 {
		return Vector{}, false
	}

	p := inverse.Mul(solution)
	if p.X % determinant != 0 || p.Y % determinant != 0 {
		return Vector{}, false
	}
	
	return Vector{
		X: p.X / determinant,
		Y: p.Y / determinant,
	}, true
}
