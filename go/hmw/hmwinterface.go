package main

import "fmt"

type Worker interface {
	work() string
}

func Printwork(k Worker) {

	fmt.Printf("your work is %v\n", k.work())

}

type workk struct {
	worrk string
}

func (l workk) work() string {

	return l.worrk

}

func main() {

	t := workk{
		worrk: "logist",
	}

	Printwork(t)

}
