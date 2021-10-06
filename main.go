package main

import "github.com/status-im/tcp-pair-sync-prototype/app"

func main() {
	a := new(app.App)
	err := a.Init()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
