package main

import (
	"github.com/go-ini/ini"
	"time"
)

type Embeded struct {
	Dates  []time.Time `delim:"|"`
	Places []string    `ini:"places,omitempty"`
	None   []int       `ini:",omitempty"`
}

type Author struct {
	Name      string `ini:"NAME"`
	Male      bool
	Age       int
	GPA       float64
	NeverMind string `ini:"-"`
	*Embeded
}

func main() {
	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"HangZhou", "Boston"},
			[]int{},
		}}
	cfg := ini.Empty()
	_ = ini.ReflectFrom(cfg, a)
	cfg.SaveTo("my.ini")
}
