// WPU Coding challenge
// 12/366

// Question :

// Answer :
package soal

func arrayPlusArray(array1 []int, array2 []int) int {
	result := 0
	for _, arr1 := range array1 {
		result += arr1
	}
	for _, arr2 := range array2 {
		result += arr2
	}
	return result
}

// Done
// di codewars nya gada golang, jadi pake js :

// function arrayPlusArray(arr1, arr2) {
// 	let result = 0
// 	for (let i = 0; i < arr1.length; i++){
// 	  result += arr1[i]
// 	}
// 	for (let i = 0; i < arr2.length; i++){
// 	  result += arr2[i]
// 	}
// 	return result
//   }

// function arrayPlusArray(arr1, arr2) {
// 	let result = 0
// 	for (let i = 0; i < arguments.length; i++){
// 	  for (let j = 0; j < arguments[i].length; j++){
// 		result += arguments[i][j]
// 	  }
// 	}
	
// 	return result
//   }