package boomer

import (
	"errors"
	"time"
)

// Boomer is a count down timer and can trigger a function.
type Boomer struct {
	SECONDS_INIT uint64
	seconds      uint64
	f            interface{}
	armed        bool
	unarmed      bool
	boomed       bool
}

// NewBoomer creates a boomer
func NewBoomer(seconds uint64, f interface{}) (*Boomer, error) {
	if seconds <= 0 {
		return nil, errors.New("invalid seconds")
	}
	p := &Boomer{}
	p.SECONDS_INIT = seconds
	p.seconds = seconds
	p.f = f
	return p, nil
}

// Let timer start to count down
func (p *Boomer) Arm() {
	go func() {
		for ; p.seconds+1 > 0; p.seconds-- {
			// fmt.Printf("\r%ds ", p.seconds)
			if p.seconds == 0 {
				p.boom(p.f)
			}
			time.Sleep(time.Second)
		}
	}()
	p.armed = true
	return
}

// Cancel boomer
func (p *Boomer) Unarm() {
	p.SECONDS_INIT = 0
	p.seconds = 0
	p.unarmed = true
	return
}

// Reset time to init
func (p *Boomer) Rewind() error {
	if !p.armed {
		return errors.New("cannot rewind, boomer not armed")
	}
	if p.unarmed {
		return errors.New("cannot rewind, boomer unarmed")
	}
	if p.unarmed {
		return errors.New("cannot rewind, boomer has boomed")
	}
	p.seconds = p.SECONDS_INIT
	return nil
}

// Trigger function
func (p *Boomer) boom(f interface{}) {
	f.(func())()
	p.boomed = true
	return
}
