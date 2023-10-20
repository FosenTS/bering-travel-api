package valueobject

var classOrder = map[uint8]uint{
	'U': 0,
	'T': 1,
	'X': 2,
	'V': 3,
	'Q': 4,
	'O': 5,
	'N': 6,
	'M': 7,
	'L': 8,
	'K': 9,
	'H': 10,
	'G': 11,
	'E': 12,
	'B': 13,
	'Y': 14,
	'S': 15,
	'W': 16,
}

type class struct {
	W, S, Y, B, E, G, H, K, L, M, N, O, Q, V, X, T, U uint8
}

var Class = &class{
	W: 'W',
	S: 'S',
	Y: 'Y',
	B: 'B',
	E: 'E',
	G: 'G',
	H: 'H',
	K: 'K',
	L: 'L',
	M: 'M',
	N: 'N',
	O: 'O',
	Q: 'Q',
	V: 'V',
	X: 'X',
	T: 'T',
	U: 'U',
}

func (c *class) Contains(class uint8) bool {
	_, ok := classOrder[class]
	return ok
}

func (c *class) Order(class uint8) uint {
	return classOrder[class]
}
