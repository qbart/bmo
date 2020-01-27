package bmo

// Devices list.
type Devices struct {
	yeeBulbs []*YeeBulb
}

// NewDevices inits device container.
func NewDevices() *Devices {
	return &Devices{
		make([]*YeeBulb, 0),
	}
}

// RegisterYeeBulb adds yee light.
func (d *Devices) RegisterYeeBulb(addr string) {
	d.yeeBulbs = append(d.yeeBulbs, NewYeeBulb(addr))
}

// BulbsPower powers on/off all bulbs.
func (d *Devices) BulbsPower(on bool) {
	for _, bulb := range d.yeeBulbs {
		go bulb.Power(on)
	}
}

// BulbsBrightness sets brightness for all bulbs.
func (d *Devices) BulbsBrightness(value int) {
	for _, bulb := range d.yeeBulbs {
		go bulb.Brightness(value)
	}
}

// BulbsColor sets color for all bulbs.
func (d *Devices) BulbsColor(r, g, b uint8) {
	for _, bulb := range d.yeeBulbs {
		go bulb.Color(r, g, b)
	}
}
