package helper

func RepeatString(input string, max int) (output string, output2 string) {

	for i := 0; i < max; i++ {
		output = output + input
	}
	output2 = output + output
	return
}
