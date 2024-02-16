// WPU Coding challenge
// 16/366

// Question :
// It's the academic year's end, fateful moment of your school report. The averages must be calculated. All the students come to you and entreat you to calculate their average for them. Easy ! You just need to write a script.
// Return the average of the given array rounded down to its nearest integer.
// The array will never be empty.

// function getAverage(marks){
// }

// Answer :
package soal

import "math"

func getAverage(marks []int) int {
	var result int
	for _, mark := range marks {
		result += mark
	}
	result /= len(marks)
	result = int(math.Floor(float64(result)))
	return result
}

// Done

// aslinya pake js, versi js :