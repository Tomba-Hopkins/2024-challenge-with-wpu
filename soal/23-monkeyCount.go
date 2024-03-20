// WPU Coding challenge
// 23/366

// Question :

// Answer :
package soal


func MonkeyCount(n int) (res []int) {

	for i := 1; i <= n; i++ {
		res = append(res, i)
	}

	return res

}

// Aslinya pake js :
// function monkeyCount(n) {
// 	let res = []
// 	for(let i = 1; i <= n; i++){
// 	  res.push(i)
// 	}
// 	return res
//   }