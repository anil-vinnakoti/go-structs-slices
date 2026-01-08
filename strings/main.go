package main

import "fmt"

func main() {
	name := "Anil Kumar"

	char := name[0]
	fmt.Println(char)				// 65 (ASCI value)
	fmt.Println(string(char))		// A

	// rune 
	var runeLetter rune = 'V'			// rune & bytes should be declared using single quotes
	fmt.Println(runeLetter)				// 86
	fmt.Println(string(runeLetter))		// V

	// byte
	var byteLetter byte = 'K'			// rune & bytes should be declared using single quotes
	fmt.Println(byteLetter)				// 75
	fmt.Println(string(byteLetter))		// K

	// converting string to rune or byte slice
	runeName := []rune(name)
	fmt.Println(runeName)

	byteName := []byte(name)
	fmt.Println(byteName)


	// slicing a string	(same as slices)
	firstName:= string(name[:5])
	fmt.Println(firstName)					// Anil


	// how to deal with unicode characters
	chineseWord := "三分钟"
	firstChar := chineseWord[0]
	fmt.Println(string(firstChar))			// ä	=> wrong

	runeChineseWord := []rune(chineseWord)
	firstRuneChar := runeChineseWord[0]
	fmt.Println(string(firstRuneChar))		// 三


	// NOTE
	// If the string contains only normal ASCII characters, you do NOT need a []rune.
	// Convert to []rune only when you need character-level operations.
}