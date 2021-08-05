package clerk

import (
	"fmt"
	"sync"
)

type Clerk struct {
	*sync.WaitGroup
	Printers []*Printer
}

type Printer struct {
	*sync.Mutex
	Mode     string
	FilePath string
	Input    chan string
}

func NewClerk() *Clerk {
	return &Clerk{}

}

func (c *Clerk) NewPrinter(mode, filepath string, input chan string) {
	c.WaitGroup.Add(1)
	p := Printer{
		Mode:     "LOG",
		FilePath: "./default.log",
		Input:    input,
	}
	fmt.Sscan(mode, &p.Mode)
	fmt.Sscan(filepath, &p.FilePath)
	c.Printers = append(c.Printers, &p)
}

func (c *Printer) ReadLog() {

}

func (c *Printer) WriteLog() {

}

func (c *Clerk) Start() {

}
