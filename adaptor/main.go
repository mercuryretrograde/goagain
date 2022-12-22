package main

import (
	"sample.com"
)

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUSBInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdaptor := &windowsAdaptor{
		windowMachine: windowsMachine,
	}
	client.insertSquareUSBInComputer(windowsMachineAdaptor)
}
