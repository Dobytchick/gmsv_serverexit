package main

import "github.com/Srlion/glua"

func init() {
	glua.GMOD13_OPEN = gmod13_open
	glua.GMOD13_CLOSE = gmod13_close
}

func shutdown(L glua.State) int {
	L.GetGlobal("engine")
	if !L.IsTable(-1) {
		L.Pop()
		L.ErrorNoHalt("serverexit: engine library is unavailable")
		return 0
	}
	L.GetField(-1, "CloseServer")
	if !L.IsFunc(-1) {
		L.PopN(2)
		L.ErrorNoHalt("serverexit: engine.CloseServer is unavailable")
		return 0
	}
	L.Call(0, 0)
	L.Pop()
	return 0
}

func gmod13_open(L glua.State) int {
	L.GetGlobal("server")
	if !L.IsTable(-1) {
		L.Pop()
		L.NewTable()
		L.SetGlobal("server")
		L.GetGlobal("server")
	}

	L.PushString("exit")
	L.PushGoFunc(shutdown)
	L.SetTable(-3)
	L.Pop()

	return 0
}

func gmod13_close(L glua.State) int {
	L.GetGlobal("server")
	if !L.IsTable(-1) {
		L.Pop()
		return 0
	}
	L.PushString("exit")
	L.PushNil()
	L.SetTable(-3)
	L.Pop()

	return 0
}

func main() {}
