// 1. To print a single digit created on data type as 'placeholder', which is array of 5 strings
//    Each array element has three places, either white space or ▉. e.g. "▉▉▉"/"▉ ▉"/" ▉ "/"   "
// 2. Using this 'placeholder' data type we will create digits like zero, one, tow, etc.
//    e.g. For digit zero
// 		zero = placeholder{
//	 		"▉▉▉",
// 			"▉ ▉",
// 			"▉ ▉",
// 			"▉ ▉",
// 			"▉▉▉",
// 		}
// 3. We will have a 'digits' array which will hold all above placeholders.
// 4. To print each digit side by side, for loop will run for 5 times. Inside that loop,
//    another loop iterate over all the digits one by one and then we will fetch elements of each
//    digit by index, starting from 0 to 4.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type placeholder [5]string

var (
	zero = placeholder{
		"▉▉▉",
		"▉ ▉",
		"▉ ▉",
		"▉ ▉",
		"▉▉▉",
	}
	one = placeholder{
		" ▉ ",
		" ▉ ",
		" ▉ ",
		" ▉ ",
		" ▉ ",
	}
	two = placeholder{
		"▉▉▉",
		"  ▉",
		"▉▉▉",
		"▉  ",
		"▉▉▉",
	}
	three = placeholder{
		"▉▉▉",
		"  ▉",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	four = placeholder{
		"▉  ",
		"▉  ",
		"▉▉▉",
		"  ▉",
		"  ▉",
	}
	five = placeholder{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	six = placeholder{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}
	seven = placeholder{
		"▉▉▉",
		"  ▉",
		"  ▉",
		"  ▉",
		"  ▉",
	}
	eight = placeholder{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}
	nine = placeholder{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}
	separater = placeholder{
		"   ",
		" ▉ ",
		"   ",
		" ▉ ",
		"   ",
	}
)

func printPh(clock []placeholder) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Println("Now time is:")
	for i := 0; i < len(clock[0]); i++ {
		for _, d := range clock {
			fmt.Printf("%s ", d[i])
		}
		fmt.Println()
	}
}

func setDigit(t int) placeholder {
	switch t {
	case 0:
		return zero
	case 1:
		return one
	case 2:
		return two
	case 3:
		return three
	case 4:
		return four
	case 5:
		return five
	case 6:
		return six
	case 7:
		return seven
	case 8:
		return eight
	case 9:
		return nine
	default:
		return zero
	}
}

func setPlaceholder(time int, pl []placeholder) {
	t1 := time / 10
	t2 := time % 10

	pl[0] = setDigit(t1)
	pl[1] = setDigit(t2)
}

func setClock(cl []placeholder) {
	tHr := time.Now().Hour()
	tMn := time.Now().Minute()
	tSc := time.Now().Second()

	setPlaceholder(tHr, cl[0:2])
	setPlaceholder(tMn, cl[3:5])
	setPlaceholder(tSc, cl[6:8])
}

func main() {
	var clock []placeholder = []placeholder{zero, zero, separater, zero, zero, separater, zero, zero}
	for {
		setClock(clock)
		printPh(clock)
		time.Sleep(1 * 1000 * 1000 * 1000)
	}
}
