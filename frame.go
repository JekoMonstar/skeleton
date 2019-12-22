package main

import (
	"fmt"
	"github.com/JekoMonstar/skeleton/fstr"
	"os"
	"strings"
)

type Frame struct {
	Module string
	Name   string
	Flag   *Flag
}

func NewFrame() *Frame {
	f := &Flag{}
	f.Parse()
	m := strings.Split(f.Module, "/")
	ins := &Frame{
		Module: f.Module,
		Name:   m[len(m)-1],
		Flag:   f,
	}
	return ins
}

func (f *Frame) pathWithPrefix(pathb string) string {
	return fmt.Sprintf("%s/%s", f.Name, pathb)
}

func (f *Frame) Mkdir() {
	os.MkdirAll(f.pathWithPrefix("app"), os.ModePerm)
	os.MkdirAll(f.pathWithPrefix("cmd/cli"), os.ModePerm)
	os.MkdirAll(f.pathWithPrefix("conf"), os.ModePerm)
	os.MkdirAll(f.pathWithPrefix("pub"), os.ModePerm)
}

func (f *Frame) autoTouch(path, fstr string) {
	if FileIsExist(path) {
		os.Remove(path)
	}
	file, _ := os.Create(path)
	defer file.Close()
	file.WriteString(fstr)
}

func (f *Frame) Touch() {
	f.autoTouch(f.pathWithPrefix("cmd/cli/main.go"), fmt.Sprintf(fstr.MainStr, f.Module))
}
