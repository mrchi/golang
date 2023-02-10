package custom

import (
	"fmt"
	"gadget"
)

func AcceptAnything(thing interface{}) {
	player, ok := thing.(gadget.TapePlayer)
	if ok {
		player.Play("中国话")
		return
	}

	recorder, ok := thing.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
		return
	}

	fmt.Printf("%v, %#v, %T\n", thing, thing, thing)
}
