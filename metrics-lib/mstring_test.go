package metrics

import (
	"fmt"
	"testing"
)

func TestPlain(t *testing.T) {
	m := NewMString()
	fmt.Println(m.Value())
	m.Update("sklar")
	fmt.Println(m.Value())
	m.Update("sklar2")
	fmt.Println(m.Value())
	if v := m.Value(); "sklar2" != v {
		t.Errorf("g.Value(): 47 != %v\n", v)
	}
}
