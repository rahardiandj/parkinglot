package parkinglot

import "fmt"

func (mod *Mod) FileReader() {

}

func (m *Message) Formatted(argument string) {
	m.Message = fmt.Sprintf(m.Message, argument)
}
