package main

import (
	"os"
	"time"

	"github.com/Srlion/glua"
)

func init() {
	glua.GMOD13_OPEN = gmod13_open
	glua.GMOD13_CLOSE = gmod13_close
}

func shutdown(L glua.State) int {
	L.GetGlobal("Msg")
	L.PushString("[Server] Safe shutdown triggered...\n")
	L.Call(1, 0)

	L.GetGlobal("hook")
	L.GetField(-1, "Run")
	L.PushString("ShutDown")
	L.Call(1, 0)
	L.Pop()

	glua.Go(func() {
		time.Sleep(200 * time.Millisecond)
		os.Exit(1)
	})

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
	L.PushString("exit")
	L.PushNil()
	L.SetTable(-3)
	L.Pop()

	return 0
}

func main() {}
