package main

//Made import this way so we can simply add imports if needed
import (
	"fmt"
	"math/rand"
)

//Var table
var (
	//Selects the type of calculation
	choice string
	//Definite the calculation
	hm_numbers int
	hl_numbers int
	//This is the loop definition
	Prog bool
	//The loop in the calculation simply called Prog2
	Prog2 bool
	//look if the user did the right input
	Prog3 bool
	//loop whiche is asking if the user like to continue
	ask bool
	//input from the user
	result int
)

//loop for ask if the want to continue run it everytime when the user is doen
func ask_user() {
	ask = true
	fmt.Println("Pls enter y if you like to continue if not enter q")
	for ask == true {
		fmt.Scanln(&choice)
		if choice == "q" {
			Prog3 = false
			ask = false
		}
		if choice == "y" {
			//fmt.Println("Continue with all the settings")
			ask = false
		}
	}
}

func mode1() {
	Prog2 = true
	for Prog2 == true {
		Prog3 = true
		//Definites the lenght of the number self 1 = 1-9 2 = 1-99....
		fmt.Println("Pls enter how long the numbers should be in max?")
		fmt.Scanln(&hl_numbers)
		//Definites the lenght of the calculation 1 = 1 + 1 2 = 1 + 1 + 1.....
		fmt.Println("Pls write how many numbers your want in total to work with max (5)")
		fmt.Scanln(&hm_numbers)
		//Making a loop again to ask the user if he like to continue
		for Prog3 == true {
			if hm_numbers == 1 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := nr1 + nr2
				fmt.Println("what is ", nr1, " + ", nr2)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr3 {
					fmt.Println("That awnser was correct")
				}
				if result != nr3 {
					fmt.Println("The awnser was incorrect right would be ", nr3)
				}
				ask_user()
			}

			if hm_numbers == 2 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := nr1 + nr2 + nr3
				fmt.Println("what is ", nr1, " + ", nr2, " + ", nr3)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr4 {
					fmt.Println("That awnser was correct")
				}
				if result != nr4 {
					fmt.Println("The awnser was incorrect right would be ", nr4)
				}
				ask_user()

			}

			if hm_numbers == 3 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := nr1 + nr2 + nr3 + nr4
				fmt.Println("what is ", nr1, " + ", nr2, " + ", nr3, " + ", nr4)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr5 {
					fmt.Println("That awnser was correct")
				}
				if result != nr5 {
					fmt.Println("The awnser was incorrect right would be ", nr5)
				}
				ask_user()
			}

			if hm_numbers == 4 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := nr1 + nr2 + nr3 + nr4 + nr5
				fmt.Println("what is ", nr1, " + ", nr2, " + ", nr3, " + ", nr4, " + ", nr5)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr6 {
					fmt.Println("That awnser was correct")
				}
				if result != nr6 {
					fmt.Println("The awnser was incorrect right would be ", nr6)
				}
				ask_user()
			}

			if hm_numbers == 5 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := rand.Intn(hl_numbers-0) + 0
				nr7 := nr1 + nr2 + nr3 + nr4 + nr5 + nr6
				fmt.Println("what is ", nr1, " + ", nr2, " + ", nr3, " + ", nr4, " + ", nr5, " + ", nr6)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr7 {
					fmt.Println("That awnser was correct")
				}
				if result != nr7 {
					fmt.Println("The awnser was incorrect right would be ", nr7)
				}
				ask_user()
			}
		}
	}
}

func mode2() {
	Prog2 = true
	for Prog2 == true {
		Prog3 = true
		//Definites the lenght of the number self 1 = 1-9 2 = 1-99....
		fmt.Println("Pls enter how long the numbers should be in max?")
		fmt.Scanln(&hl_numbers)
		//Definites the lenght of the calculation 1 = 1 + 1 2 = 1 + 1 + 1.....
		fmt.Println("Pls write how many numbers your want in total to work with max (5)")
		fmt.Scanln(&hm_numbers)
		//Making a loop again to ask the user if he like to continue
		for Prog3 == true {
			if hm_numbers == 1 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := nr1 * nr2
				fmt.Println("what is ", nr1, " x ", nr2)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr3 {
					fmt.Println("That awnser was correct")
				}
				if result != nr3 {
					fmt.Println("The awnser was incorrect right would be ", nr3)
				}
				ask_user()
			}

			if hm_numbers == 2 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := nr1 * nr2 * nr3
				fmt.Println("what is ", nr1, " x ", nr2, " x ", nr3)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr4 {
					fmt.Println("That awnser was correct")
				}
				if result != nr4 {
					fmt.Println("The awnser was incorrect right would be ", nr4)
				}
				ask_user()

			}

			if hm_numbers == 3 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := nr1 * nr2 * nr3 * nr4
				fmt.Println("what is ", nr1, " x ", nr2, " x ", nr3, " x ", nr4)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr5 {
					fmt.Println("That awnser was correct")
				}
				if result != nr5 {
					fmt.Println("The awnser was incorrect right would be ", nr5)
				}
				ask_user()
			}

			if hm_numbers == 4 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := nr1 * nr2 * nr3 * nr4 * nr5
				fmt.Println("what is ", nr1, " x ", nr2, " x ", nr3, " x ", nr4, " x ", nr5)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr6 {
					fmt.Println("That awnser was correct")
				}
				if result != nr6 {
					fmt.Println("The awnser was incorrect right would be ", nr6)
				}
				ask_user()
			}

			if hm_numbers == 5 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := rand.Intn(hl_numbers-0) + 0
				nr7 := nr1 * nr2 * nr3 * nr4 * nr5 * nr6
				fmt.Println("what is ", nr1, " x ", nr2, " x ", nr3, " x ", nr4, " x ", nr5, " x ", nr6)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr7 {
					fmt.Println("That awnser was correct")
				}
				if result != nr7 {
					fmt.Println("The awnser was incorrect right would be ", nr7)
				}
				ask_user()
			}
		}
	}
}

func mode3() {
	Prog2 = true
	for Prog2 == true {
		Prog3 = true
		//Definites the lenght of the number self 1 = 1-9 2 = 1-99....
		fmt.Println("Pls enter how long the numbers should be in max?")
		fmt.Scanln(&hl_numbers)
		//Definites the lenght of the calculation 1 = 1 + 1 2 = 1 + 1 + 1.....
		fmt.Println("Pls write how many numbers your want in total to work with max (5)")
		fmt.Scanln(&hm_numbers)
		//Making a loop again to ask the user if he like to continue
		for Prog3 == true {
			if hm_numbers == 1 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := nr1 - nr2
				fmt.Println("what is ", nr1, " - ", nr2)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr3 {
					fmt.Println("That awnser was correct")
				}
				if result != nr3 {
					fmt.Println("The awnser was incorrect right would be ", nr3)
				}
				ask_user()
			}

			if hm_numbers == 2 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := nr1 - nr2 - nr3
				fmt.Println("what is ", nr1, " - ", nr2, " - ", nr3)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr4 {
					fmt.Println("That awnser was correct")
				}
				if result != nr4 {
					fmt.Println("The awnser was incorrect right would be ", nr4)
				}
				ask_user()

			}

			if hm_numbers == 3 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := nr1 - nr2 - nr3 - nr4
				fmt.Println("what is ", nr1, " - ", nr2, " - ", nr3, " - ", nr4)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr5 {
					fmt.Println("That awnser was correct")
				}
				if result != nr5 {
					fmt.Println("The awnser was incorrect right would be ", nr5)
				}
				ask_user()
			}

			if hm_numbers == 4 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := nr1 - nr2 - nr3 - nr4 - nr5
				fmt.Println("what is ", nr1, " - ", nr2, " - ", nr3, " - ", nr4, " - ", nr5)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr6 {
					fmt.Println("That awnser was correct")
				}
				if result != nr6 {
					fmt.Println("The awnser was incorrect right would be ", nr6)
				}
				ask_user()
			}

			if hm_numbers == 5 {
				//number definitons done it this way might be dumb but well it works lol
				nr1 := rand.Intn(hl_numbers-0) + 0
				nr2 := rand.Intn(hl_numbers-0) + 0
				nr3 := rand.Intn(hl_numbers-0) + 0
				nr4 := rand.Intn(hl_numbers-0) + 0
				nr5 := rand.Intn(hl_numbers-0) + 0
				nr6 := rand.Intn(hl_numbers-0) + 0
				nr7 := nr1 - nr2 - nr3 - nr4 - nr5 - nr6
				fmt.Println("what is ", nr1, " - ", nr2, " - ", nr3, " - ", nr4, " - ", nr5, " - ", nr6)
				//the input
				fmt.Scanln(&result)
				//Check if our input is the result or not
				if result == nr7 {
					fmt.Println("That awnser was correct")
				}
				if result != nr7 {
					fmt.Println("The awnser was incorrect right would be ", nr7)
				}
				ask_user()
			}
		}
	}
}

func main() {
	fmt.Println("Pls select one of the following options 1 = 1+1, 2 = 1x1, 3 = 1-1")
	Prog = true
	for Prog == true {
		fmt.Scanln(&choice)
		if choice == "1" {
			fmt.Println("Your choice was 1")
			Prog = false
			mode1()
		}
		if choice == "2" {
			fmt.Println("Your choice was 2")
			Prog = false
			mode2()
		}
		if choice == "3" {
			fmt.Println("Your choice was 3")
			Prog = false
			mode3()
		}
	}
}
