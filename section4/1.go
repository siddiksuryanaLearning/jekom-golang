package section4

import (
	"math"
)

func hitungLuasPermukaanTabung(jariJari, tinggi float64) float64 {
	luasPermukaan := 2 * math.Pi * jariJari * (jariJari + tinggi)
	return luasPermukaan
}
