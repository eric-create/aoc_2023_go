package vectors

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(other Vector) Vector {
	new := Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
	return new
}

// Returns `0`.
func X() int {
	return 0
}

// Returns `1`.
func Y() int {
	return 1
}

func Up() Vector {
	return Vector{0, -1}
}

func RightUp() Vector {
	return Vector{1, -1}
}

func Right() Vector {
	return Vector{1, 0}
}

func RightDown() Vector {
	return Vector{1, 1}
}

func Down() Vector {
	return Vector{0, 1}
}

func LeftDown() Vector {
	return Vector{-1, 1}
}

func Left() Vector {
	return Vector{-1, 0}
}

func LeftUp() Vector {
	return Vector{-1, -1}
}

// Up, RightUp, Right, RightDown, Down, LeftDown, Left, LeftUp
func AllDirections() []Vector {
	return []Vector{
		Up(),
		RightUp(),
		Right(),
		RightDown(),
		Down(),
		LeftDown(),
		Left(),
		LeftUp(),
	}
}

// Up, Right, Down, Left
func ManhattanDirections() []Vector {
	return []Vector{
		Up(),
		Right(),
		Down(),
		Left(),
	}
}

// Right, Left
func Horizontal() []Vector {
	return []Vector{
		Right(),
		Left(),
	}
}

// Up, Down
func Vertical() []Vector {
	return []Vector{
		Up(),
		Down(),
	}
}
