package day2

import (
	"strconv"
	"strings"
)

type Color string

// const (
// 	R color = "red"
// 	G color = "green"
// 	B color = "blue"
// )

type CubesHand map[Color]uint32

// NewHandFromString from input string
//
// e.g.: 8 green, 6 blue, 20 red
func NewHandFromString(in string) CubesHand {
	out := CubesHand{}
	colorDraws := strings.Split(in, ",")
	for _, cd := range colorDraws {
		cd = strings.TrimSpace(cd)
		valStr, c, found := strings.Cut(cd, " ")
		if !found {
			continue
		}

		val, err := strconv.Atoi(valStr)
		if err != nil {
			continue
		}

		out[Color(c)] = uint32(val)
	}

	return out
}

func (h CubesHand) IsValid(rules CubesHand) bool {
	for c, val := range h {
		if rule := rules[c]; val > rule {
			return false
		}
	}

	return true
}

func (h CubesHand) LogMinimalVal(in CubesHand) {
	for c, val := range in {
		hc, ok := h[c]
		if !ok {
			h[c] = val
			continue
		}
		if hc < val {
			h[c] = val
		}
	}
}

func (h CubesHand) ToSlice() []uint32 {
	valsSlice := make([]uint32, 0, len(h))
	for _, val := range h {
		valsSlice = append(valsSlice, val)
	}

	return valsSlice
}

func (h CubesHand) Power() uint32 {
	valsSlice := h.ToSlice()
	if len(valsSlice) < 1 {
		return 0
	}

	out := valsSlice[0]
	for _, val := range valsSlice[1:] {
		out *= val
	}

	return out
}
