package main

func Join(input []string) (output string) {
	for index, item := range input {
		if index == 0 {
			output = item //output = Halo
			continue
		}
		output = output + " " + item
	}
	return
}
