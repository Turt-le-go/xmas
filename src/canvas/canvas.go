package canvas

import (
	"fmt"
	"os/exec"
	"os"
	"log"
)

//type Canvas struct {
//	width int
//	height int
//}

func GetSize() (x,y int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	_,err = fmt.Sscanf(string(out),"%d %d",&y,&x)
	if err != nil {
		log.Fatal(err)
	}
	return x,y
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func DrawSymbol(x,y int, symbol string) {
	w,h := GetSize()
	if x > w || x < 0 {
		log.Fatalf("\nX (%v) is not between 0 and %v\n", x, w)
	}
	if y > h || y < 0 {
		log.Fatalf("\nY (%v) is not between 0 and %v\n", y, h)
	}
	fmt.Printf("\033[%d;%dH%s", y, x, symbol)
}
