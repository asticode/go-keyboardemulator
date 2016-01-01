package keyboardemulator

import (
	"errors"
	"strconv"
	"time"

	"github.com/AllenDang/w32"
)

// Event types
const (
	eventKeyDown    = 0x0000
	eventKeyUp      = 0x0002
	eventKeyUnicode = 0x0004
)

func (ke keyboardEmulator) down(keys Keys) error {
	// Initialize
	var inputs []w32.INPUT

	// Add key down
	for _, key := range keys {
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         keymap[key],
				WScan:       0,
				DwFlags:     eventKeyDown,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	// Send inputs
	r := w32.SendInput(inputs)

	// Check return
	if int(r) != len(keys) {
		return errors.New("Invalid return status " + strconv.Itoa(int(r)))
	}

	// Return
	return nil
}

func (ke keyboardEmulator) up(keys Keys) error {
	// Initialize
	var inputs []w32.INPUT

	// Add key down
	for _, key := range keys {
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         keymap[key],
				WScan:       0,
				DwFlags:     eventKeyUp,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	// Send inputs
	r := w32.SendInput(inputs)

	// Check return
	if int(r) != len(keys) {
		return errors.New("Invalid return status " + strconv.Itoa(int(r)))
	}

	// Return
	return nil
}

func (ke keyboardEmulator) press(keys Keys) error {
	// Initialize
	var inputs []w32.INPUT

	// Add key down
	for _, key := range keys {
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         keymap[key],
				WScan:       0,
				DwFlags:     eventKeyDown,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	// Add key up
	for _, key := range keys {
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         keymap[key],
				WScan:       0,
				DwFlags:     eventKeyUp,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	// Send inputs
	r := w32.SendInput(inputs)

	// Check return
	if int(r) != 2*len(keys) {
		return errors.New("Invalid return status " + strconv.Itoa(int(r)))
	}

	// Return
	return nil
}

func (ke keyboardEmulator) write(i string) error {
	// Loop through characters
	for k, c := range i {
		// Initialize
		inputs := []w32.INPUT{}

		// Add real life delay between presses
		if k > 0 && vk.addRealLifeDelayBetweenPresses {
			time.Sleep(ke.randomRealLifeDelay())
		}

		// Add key down
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(c),
				DwFlags:     eventKeyDown | eventKeyUnicode,
				Time:        0,
				DwExtraInfo: 0,
			},
		})

		// Add key up
		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(c),
				DwFlags:     eventKeyUp | eventKeyUnicode,
				Time:        0,
				DwExtraInfo: 0,
			},
		})

		// Send inputs
		r := w32.SendInput(inputs)

		// Check return
		if r != 2 {
			return errors.New("Invalid return status " + strconv.Itoa(int(r)))
		}
	}

	// Return
	return nil
}
