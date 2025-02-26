package main

type duck struct{}

func (d duck) Says() string {
	return "kray-kray"
}

var Animal duck
