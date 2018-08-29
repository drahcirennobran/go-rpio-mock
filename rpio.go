package rpio

type Pin uint8

func Open() (err error) {
	return nil
}

func (pin Pin) Output() {
}

func Close() error {
	return nil
}

func (pin Pin) High() {
	//fmt.Println("pin ", pin, " up")
}

func (pin Pin) Low() {
	//fmt.Println("pin ", pin, " low")
}
