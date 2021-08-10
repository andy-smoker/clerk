package clerk

import (
	"fmt"
	"os"
	"time"
)

var levels []string = []string{"DEBUG", "INFO", "ERROR", "WARNING"}
var mods []string = []string{"trace", "info"}

type Printer struct {
	Mode     string
	FilePath string
}

type Message struct {
	Type string
	Date time.Time
	Log  string
}

func NewPrinter(mod, name, filepath string) *Printer {
	p := Printer{
		Mode:     mods[1],
		FilePath: name + ".log",
	}
	for _, v := range mods {
		if v == mod {
			p.Mode = v
		}
	}
	fmt.Sscan(filepath, &p.FilePath)
	if _, err := os.Stat(p.FilePath); os.IsNotExist(err) {
		os.Create(p.FilePath)
	}
	return &p
}

func (p *Printer) writeToFile(input string) error {
	// open file for append strings
	file, err := os.OpenFile(p.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(input)
	return err
}

// WriteLog lvl 0:"DEBUG", 1:"INFO", 2:"ERROR", 3:"WARNING"
func (p *Printer) WriteLog(lvl int, t time.Time, msg string) error {
	if p.Mode == mods[1] && lvl < 1 {
		return nil
	}
	return p.writeToFile(fmt.Sprintf("%s [%s] %s", t.Format("\n2006-01-02 15:04:05"), levels[lvl], msg))
}
