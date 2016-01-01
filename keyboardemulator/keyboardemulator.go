package keyboardemulator

import (
	"math/rand"
	"time"
)

// KeyboardEmulator represents a keyboard emulator
type KeyboardEmulator interface {
	// Events
	Down(keys Keys) error
	Press(keys ...Keys) error
	Up(keys Keys) error
	Write(i string) error

	// Setters and Getters
	AddRealLifeDelayBetweenPresses() bool
	SetAddRealLifeDelayBetweenPresses(v bool) KeyboardEmulator
}

// NewKeyboardEmulator creates a new keyboard emulator
func NewKeyboardEmulator() KeyboardEmulator {
	return &keyboardEmulator{}
}

type keyboardEmulator struct {
	addRealLifeDelayBetweenPresses bool
}

func (ke keyboardEmulator) Down(keys Keys) error {
	return ke.down(keys)
}

func (ke keyboardEmulator) Press(keys ...Keys) error {
	// Range over keys
	for _, k := range keys {
		// TODO Make it optionnal ?
		time.Sleep(time.Duration(50) * time.Millisecond)

		// Press key
		if e := ke.press(k); e != nil {
			return e
		}
	}

	// Return
	return nil
}

func (ke keyboardEmulator) Up(keys Keys) error {
	return ke.up(keys)
}

func (ke keyboardEmulator) Write(i string) error {
	return ke.write(i)
}

func (ke keyboardEmulator) AddRealLifeDelayBetweenPresses() bool {
	return ke.addRealLifeDelayBetweenPresses
}

func (ke *keyboardEmulator) SetAddRealLifeDelayBetweenPresses(v bool) KeyboardEmulator {
	ke.addRealLifeDelayBetweenPresses = v
	return ke
}

func (ke keyboardEmulator) randomRealLifeDelay() time.Duration {
	// Get random int
	max := 70
	min := 30
	i := min + rand.Intn(max-min)

	// Return duration
	return time.Duration(i) * time.Millisecond
}
