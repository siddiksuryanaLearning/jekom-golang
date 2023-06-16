package section4

import (

)

func konversiGrade(nilai int) string {
	if nilai >= 0 && nilai <= 34 {
		return "E"
	} else if nilai >= 35 && nilai <= 49 {
		return "D"
	} else if nilai >= 50 && nilai <= 64 {
		return "C"
	} else if nilai >= 65 && nilai <= 79 {
		return "B"
	} else if nilai >= 80 && nilai <= 100 {
		return "A"
	} else {
		return "Nilai tidak valid"
	}
}