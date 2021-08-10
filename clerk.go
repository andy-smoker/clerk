package clerk

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var levels []string = []string{"DEBUG", "INFO", "ERROR", "WARNING"}
var mods []string = []string{"trace", "info"}

type Clerk struct {
	*sync.WaitGroup
	Printers map[string]Printer
}

type Printer struct {
	*sync.Mutex
	Mode     string
	FilePath string
}

type Message struct {
	Type string
	Date time.Time
	Log  string
}

func NewClerk() *Clerk {
	return &Clerk{
		WaitGroup: &sync.WaitGroup{},
		Printers:  map[string]Printer{},
	}

}

func (c *Clerk) NewPrinter(mode, name, filepath string, input string) *Printer {
	c.WaitGroup.Add(1)
	p := Printer{
		Mode:     "LOG",
		FilePath: "./" + name + ".log",
	}
	fmt.Sscan(mode, &p.Mode)
	fmt.Sscan(filepath, &p.FilePath)
	if _, err := os.Stat(p.FilePath); os.IsNotExist(err) {
		os.Create(p.FilePath)
	}
	c.Printers[name] = p
	return &p
}

func (p *Printer) WriteLog(input string) {
	file, err := os.Open(p.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	p.Mutex.Lock()
	_, err = file.WriteString(input)
	p.Mutex.Unlock()
	if err != nil {
		fmt.Println(err)
	}
}
