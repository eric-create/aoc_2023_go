package vectors

// Returns `0`.
func X() int {
	return 0
}

// Returns `1`.
func Y() int {
	return 1
}

func Up() [2]int {
	return [2]int{0, -1}
}

func RightUp() [2]int {
	return [2]int{1, -1}
}

func Right() [2]int {
	return [2]int{1, 0}
}

func RightDown() [2]int {
	return [2]int{1, 1}
}

func Down() [2]int {
	return [2]int{0, 1}
}

func LeftDown() [2]int {
	return [2]int{-1, 1}
}

func Left() [2]int {
	return [2]int{-1, 0}
}

func LeftUp() [2]int {
	return [2]int{-1, -1}
}

// Up, RightUp, Right, RightDown, Down, LeftDown, Left, LeftUp
func AllDirections() [][2]int {
	return [][2]int{
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
func ManhattanDirections() [][2]int {
	return [][2]int{
		Up(),
		Right(),
		Down(),
		Left(),
	}
}

// Right, Left
func Horizontal() [][2]int {
	return [][2]int{
		Right(),
		Left(),
	}
}

func Add(a [2]int, b [2]int) *[2]int {
	newVector := [2]int{a[0] + b[0], a[1] + b[1]}
	return &newVector
}
