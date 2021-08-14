package helper

func JoinText(input []string) (output string) {

	for key, value := range input {
		if key == 0 {
			output = value
		} else {
			output = output + " " + value
		}
	}

	return
}
