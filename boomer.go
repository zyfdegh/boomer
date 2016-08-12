package boomer

import "time"

// Boomer is a count down timer and can trigger a function.
type Boomer struct {
	SECONDS_INIT uint64
	seconds      uint64
	f            interface{}
}

// NewBoomer creates a boomer
func NewBoomer(seconds uint64, f interface{}) *Boomer {
	p := &Boomer{}
	p.SECONDS_INIT = seconds
	p.seconds = seconds
	p.f = f
	return p
}

// Let timer start to count down
func (p *Boomer) Arm() {
	go func() {
		for ; p.seconds >= 0; p.seconds-- {
			if p.seconds-1 == 0 {
				p.boom(p.f)
			}
			time.Sleep(time.Second)
		}
	}()
	return
}

// Cancel boomer
func (p *Boomer) Unarm() {
	p.SECONDS_INIT = 0
	p.seconds = 0
	return
}

// Reset time to init
func (p *Boomer) Rewind() {
	p.seconds = p.SECONDS_INIT
	return
}

// Trigger function
func (p *Boomer) boom(f interface{}) {
	f.(func())()
	return
}
