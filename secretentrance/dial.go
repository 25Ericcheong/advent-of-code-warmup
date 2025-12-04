package secretentrance

const minDialCount = 0
const maxDialCount = 99

const normalizeCount = 100

type Dial struct {
	Number int

	PasswordCounter int
}

func (d *Dial) TurnLeft(number int) {
	d.Number -= number

	if d.Number < minDialCount {
		d.Number += normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) TurnRight(number int) {
	d.Number += number

	if d.Number > maxDialCount {
		d.Number %= normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) plusPasswordCounterIfNeeded() {
	if d.Number == minDialCount {
		d.PasswordCounter += 1
	}
}
