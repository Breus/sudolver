package main

func main() {
	sudoku := getInput()

	for !(validateInput(sudoku)) {
		sudoku = getInput()
	}
}
