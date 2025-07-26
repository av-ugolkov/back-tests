package main

import "log"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Fatal("fatal") // 'defer' doesn't run because log.Fatal uses os.Exit which affects 'defer'.
	// panic("panic")
}
