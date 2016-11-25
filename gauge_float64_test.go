package metrics

import (
	"testing"
)

func TestGaugeFloat64(t *testing.T) {
	g := NewGaugeFloat64()
	g.Update(float64(47.0))
	if v := g.Value(); v-47 > 0.0001 || v-47 < -0.0001 {
		t.Errorf("g.Value(): 47.0 != %v\n", v)
	}
}
