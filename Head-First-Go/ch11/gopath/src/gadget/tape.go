package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("TapePlayer Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("TapePlayer Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("TapeRecorder Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("TapeRecorder Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("TapeRecorder Recording...")
}
