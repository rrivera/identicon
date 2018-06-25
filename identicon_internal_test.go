package identicon

import (
	"image"
	"testing"
)

func TestIdentIconNextPoint(t *testing.T) {
	type testS struct {
		b        byte
		p        image.Point
		width    int
		height   int
		expected image.Point
	}
	tt := []testS{
		testS{
			b:        (moveUp | moveRight),
			p:        image.Point{0, 0},
			width:    3,
			height:   5,
			expected: image.Point{1, 4},
		},
		testS{
			b:        moveUp,
			p:        image.Point{2, 3},
			width:    3,
			height:   5,
			expected: image.Point{2, 2},
		},
		testS{
			b:        (moveUp | moveLeft),
			p:        image.Point{2, 4},
			width:    3,
			height:   5,
			expected: image.Point{1, 3},
		},
		testS{
			b:        (moveDown | moveUp | moveLeft),
			p:        image.Point{0, 0},
			width:    3,
			height:   5,
			expected: image.Point{2, 0},
		},
		testS{
			b:        (moveRight),
			p:        image.Point{2, 0},
			width:    3,
			height:   5,
			expected: image.Point{0, 0},
		},
	}

	for i, ut := range tt {
		r := nextPoint(ut.b, ut.p, ut.width, ut.height)
		if r != ut.expected {
			t.Errorf("Unexpected Point at index %v: expected %v, actual %v", i, ut.expected, r)
		}
	}
}

func TestIdentIconInitialPoint(t *testing.T) {
	expected := image.Point{4, 2}
	ip := initialPoint(234, expected.X, expected.Y)
	if ip != expected {
		t.Errorf("Unexpected Initial Point: expected %v, actual %v", expected, ip)
	}
}

func TestIdentIconMirrorSymmetric(t *testing.T) {
	type testS struct {
		p        image.Point
		size     int
		expected image.Point
	}
	tt := []testS{
		testS{
			p:        image.Point{0, 0},
			size:     5,
			expected: image.Point{4, 0},
		},
		testS{
			p:        image.Point{1, 3},
			size:     5,
			expected: image.Point{3, 3},
		},
		testS{
			p:        image.Point{2, 3},
			size:     5,
			expected: image.Point{2, 3},
		},
		testS{
			p:        image.Point{3, 1},
			size:     5,
			expected: image.Point{1, 1},
		},
		testS{
			p:        image.Point{0, 0},
			size:     10,
			expected: image.Point{9, 0},
		},
		testS{
			p:        image.Point{20, 34},
			size:     100,
			expected: image.Point{79, 34},
		},
	}

	for i, ut := range tt {
		r := mirrorSymmetric(ut.p, ut.size)
		if r != ut.expected {
			t.Errorf("Unexpected Point at index %v: expected %v, actual %v", i, ut.expected, r)
		}
	}
}

func TestIdentIconGetFillValue(t *testing.T) {
	type testS struct {
		b        byte
		expected int
	}
	tt := []testS{
		testS{
			b:        0,
			expected: 0,
		},
		testS{
			b:        fillPoint,
			expected: 1,
		},
		testS{
			b:        0x2,
			expected: 1,
		},
		testS{
			b:        0x8,
			expected: 1,
		},
		testS{
			b:        0x1,
			expected: 0,
		},
		testS{
			b:        0xF,
			expected: 1,
		},
		testS{
			b:        0xA,
			expected: 1,
		},
		testS{
			b:        0x5,
			expected: 0,
		},
	}

	for i, ut := range tt {
		v := getFillValue(ut.b)
		if v != ut.expected {
			t.Errorf("Unexpected Value at index %v: expected %v, actual %v", i, ut.expected, v)
		}
	}
}
