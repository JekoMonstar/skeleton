package main

import "flag"

type Flag struct {
	Module string
}

func (f *Flag) Parse() {
	flag.StringVar(&f.Module, "m", "", "Module")
	flag.Parse()
}
