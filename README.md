# Important

`go-keyboardemulator` is only available for Windows for now. Linux will come soon and MAC OSX too I hope :D

# About

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/asticode/go-keyboardemulator/keyboardemulator)

`go-keyboardemulator` is a cross-OS keyboard emulator that can take control of your keyboard through code  for the GO programming language (http://golang.org).

# Install `go-keyboardemulator`

Run the following command:

    $ go get github.com/asticode/go-keyboardemulator/keyboardemulator
    
# Example

To make this test effective, create a new text document, open it, come back to your shell with ALT+TAB and execute the following code

    // Create a keyboard emulator
    vk := keyboardemulator.NewKeyboardEmulator().SetAddRealLifeDelayBetweenPresses(true)

    // Switch focus to the text document and select all its content
    vk.Press(
        // Switch focus
        keyboardemulator.Keys{
            keyboardemulator.KeyAlt,
            keyboardemulator.KeyTab,
        },
        // Select all
        keyboardemulator.Keys{
            keyboardemulator.KeyControl,
            keyboardemulator.KeyA,
        },
    )

    // Write
    vk.Write("go-keyboardemulator is awesome!")
