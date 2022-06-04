package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// get inputted text
	input := os.Args

	// remove first item in input list (because it's basically like some path thing)
	input = input[1:]

	// put every single string together into one large joined string, and then separate each word with a space
	joinedInput := strings.Join(input, " ")

	// Brainfuck2 code that will be outputted
	output := ""

	// iterate through every character in joinedInput
	for i := 0; i < len(joinedInput); i++ {
		// if the character is a space
		if string(joinedInput[i]) == " " {
			output += "!-;" // just add a space inserting Brainfuck2 command to the output
		} else { // otherwise
			// get binary (byte) form of the character value, and keep it as a string
			byteValue := fmt.Sprintf("%.8b", []byte(string(joinedInput[i])))

			// now we have a problem, the string is surrounded with []

			// so take them out
			byteValue = strings.Replace(byteValue, "[", "", 1)
			byteValue = strings.Replace(byteValue, "]", "", 1)

			// add a character command with the new stringified byte value
			output += "!" + byteValue + ";"
			// so it will look something like:
			// !xxxxxxxx;
			// x is either a 1 or 0 obviously
		}
	}

	// attempt to create an output file
	file, err := os.Create("output.bf2")

	// if error isn't nil
	if err != nil {
		fmt.Print("\nYou fucking idiot.\n\n") // say "You fucking idiot."
		panic(err)                            // panic the error
	}

	// try to write the output to the output Brainfuck2 file
	_, writeErr := file.WriteString(output)

	// if there is a write error
	if writeErr != nil {
		fmt.Print("\nYou fucking idiot.\n\n") // say "You fucking idiot."
		panic(err)                            // panic the error
	}

	// close the file
	file.Close()

	// print success message
	fmt.Println("Created output Brainfuck2 file successfully, now go away")
}
