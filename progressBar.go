package progressbar

import (
	"fmt"
)

type ProgressBar struct {
	value int
	r     string
	long  int
	begin string
	end   string
}

func (pb *ProgressBar) Add(v int) {
	pb.value += v
}

func (pb ProgressBar) Update() {
	fmt.Print("\r")
	fmt.Print(pb.begin)
	fmt.Print("  ")
	v := float64(pb.value) / 100
	for i := 0; i < int(v*float64(pb.long)); i++ {
		fmt.Print(pb.r)
	}
	fmt.Print("  ")
	fmt.Print(pb.end)
}

func (pb *ProgressBar) SetBegin(v string) {
	pb.begin = v
}

func (pb *ProgressBar) SetEnd(v string) {
	pb.end = v
	
}

func New(r string, long int, begin, end string) ProgressBar {
	return ProgressBar{r: r, long: long, begin: begin, end: end}
}