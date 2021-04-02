package main

import (
	"whxph.com/xph_csv/communication"
	"whxph.com/xph_csv/fileoperation"
)

func main() {

	go communication.Start()

	go fileoperation.Start()

	select {}
}
