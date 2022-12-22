package main

type windowsAdaptor struct {
	windowsMachine *windows
}

func (w *windowsAdaptor) insertInSquarePort() {
	w.windowsMachine.insertInCirclePort()
}
