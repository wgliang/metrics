package metrics

import (
	"testing"
)

func TestMeterZero(t *testing.T) {
	m := NewMeter()
	if count := m.Count(); 0 != count {
		t.Errorf("m.Count(): 0 != %v\n", count)
	}
}

func TestMeterNonzero(t *testing.T) {
	m := NewMeter()
	m.Mark(3)
	if count := m.Count(); 3 != count {
		t.Errorf("m.Count(): 3 != %v\n", count)
	}
}

func TestMeterRate1(t *testing.T) {
	m := NewMeter()
	m.Mark(3)
	m.Tick()
	const expected = 0.6
	if r1 := m.Rate1(); r1 != expected {
		t.Errorf("m.Rate1(): %v != %v\n", expected, r1)
	}
}
