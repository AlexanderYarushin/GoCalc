package calculator

type SignType struct {
	Index    uint8
	Priority uint8
}

var Signs = map[string]SignType{
	"+": {1, 1},
	"-": {2, 1},
	"*": {3, 2},
	"/": {4, 2},
	"(": {5, 0},
	")": {6, 0},
}
