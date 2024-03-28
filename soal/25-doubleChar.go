// WPU Coding challenge
// 24/366

// Question :

// Given a string, you have to return a string in which each character (case-sensitive) is repeated once.
// Examples (Input -> Output):
// * "String"      -> "SSttrriinngg"
// * "Hello World" -> "HHeelllloo  WWoorrlldd"
// * "1234!_ "     -> "11223344!!__  "
// Good Luck!

// Answer :
package soal

func DoubleChar(text string) (result string) {
	for _, t := range text {
		str := string(t)
		result += str + str
	}
	return
}