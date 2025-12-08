package secretentrance

type Dial struct {
	Count           int
	PasswordCounter int
}

const minDialCount = 0
const maxDialCount = 99

const normalizeCount = 100

func (d *Dial) TurnLeft(number int) {
	d.Count = -number

	if d.Count < minDialCount {
		d.Count += normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) TurnRight(number int) {
	d.Count += number

	if d.Count > maxDialCount {
		d.Count %= normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) plusPasswordCounterIfNeeded() {
	if d.Count == minDialCount {
		d.PasswordCounter += 1
	}
}

func LoadData() {

}
