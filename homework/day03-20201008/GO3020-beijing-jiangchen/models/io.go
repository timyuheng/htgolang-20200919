package models

import (
	"github.com/peterh/liner"
)

//DIYLiner ...
type DIYLiner struct {
	State  *liner.State
	tmode  liner.ModeApplier
	lmode  liner.ModeApplier
	paused bool
}

//NewLiner ...
//wrap of liner.NewLiner
func NewLiner() (ret *DIYLiner) {
	ret = &DIYLiner{}
	ret.tmode, _ = liner.TerminalMode()

	line := liner.NewLiner()
	line.SetMultiLineMode(true)
	line.SetCtrlCAborts(true)

	ret.State = line
	return
}

//Pause ...
func (dl *DIYLiner) Pause() (err error) {
	if dl.paused {
		panic("DIYLiner already paused")
	}
	dl.paused = true
	err = dl.tmode.ApplyMode()
	return
}

//Resume ...
func (dl *DIYLiner) Resume() (err error) {
	if !dl.paused {
		panic("DIYLiner is not paused")
	}
	dl.paused = false
	err = dl.lmode.ApplyMode()
	return
}

//Close ...
func (dl *DIYLiner) Close() (err error) {
	err = dl.State.Close()
	if err != nil {
		return err
	}
	return nil
}

//RunInput ...
func RunInput() {
	line := NewLiner()
	defer line.Close()

	Name, err := line.State.Prompt("Please input Name > ")
	if err != nil {
		return
	}
	Contact, err := line.State.Prompt("Please input Contact Number > ")
	if err != nil {
		return
	}
	Address, err := line.State.Prompt("Please input Address > ")
	if err != nil {
		return
	}
	AddElement(&Users, GenerateElement(GetMaxIDPlusOne(&Users), Name, Contact, Address))
}
