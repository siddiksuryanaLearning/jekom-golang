package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func main() {
	students := make([]Student, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan nama siswa %d: ", i+1)
		fmt.Scanln(&students[i].Name)

		fmt.Printf("Masukkan skor siswa %d: ", i+1)
		for j := 0; j < 3; j++ {
			var score int
			fmt.Scanln(&score)
			students[i].Score = append(students[i].Score, score)
		}
	}

	totalScore := 0
	for _, student := range students {
		for _, score := range student.Score {
			totalScore += score
		}
	}
	averageScore := float64(totalScore) / float64(len(students)*len(students[0].Score))

	minScore := math.MaxInt32
	var minStudent string
	for _, student := range students {
		sum := 0
		for _, score := range student.Score {
			sum += score
		}
		avg := float64(sum) / float64(len(student.Score))
		if avg < float64(minScore) {
			minScore = int(avg)
			minStudent = student.Name
		}
	}

	maxScore := math.MinInt32
	var maxStudent string
	for _, student := range students {
		sum := 0
		for _, score := range student.Score {
			sum += score
		}
		avg := float64(sum) / float64(len(student.Score))
		if avg > float64(maxScore) {
			maxScore = int(avg)
			maxStudent = student.Name
		}
	}

	fmt.Printf("\nSkor rata-rata: %.2f\n", averageScore)
	fmt.Printf("Siswa dengan skor minimum: %s (%d)\n", minStudent, minScore)
	fmt.Printf("Siswa dengan skor maksimum: %s (%d)\n", maxStudent, maxScore)
}
