package main

import (
	"github.com/Flokey82/aifiver"
)

func main() {
	t := aifiver.NewTraiter()
	aifiver.DefaultTraits(t)
	p := t.NewPersonalityFromPreset(aifiver.PresetUnstable)
	p.Log()

	p2 := t.NewPersonalityFromPreset(aifiver.PresetFool)
	p2.Log()
}
