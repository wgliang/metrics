package poetry

import (
	"testing"
	"fmt"
)

func TestNewflow(t *testing.T) {
	fl := NewFlow(5)

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}

func TestRPush(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush("htwiu4bhnyimn")
	fl.RPush("sfewiqughrwughr")
	fl.RPush("fiqruwght")

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}

func TestLPop(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush("htwiu4bhnyimn")
	fl.LPush("fewnignr")
	fl.RPush("sfewiqughrwughr")
	fl.LPush("frewqgutbtr")
	fl.RPush("fiqruwght")
	fl.LPop()
	fl.LPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}

func TestLPush(t *testing.T) {
	fl := NewFlow(5)
	fl.LPush("htwiu4bhnyimn")
	fl.LPush("sfewiqughrwughr")
	fl.LPush("fiqruwght")

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}

func TestRPop(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush("htwiu4bhnyimn")
	fl.LPush("fewnignr")
	fl.RPush("sfewiqughrwughr")
	fl.LPush("frewqgutbtr")
	fl.RPush("fiqruwght")
	fl.RPop()
	fl.RPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}

func TestSize(t *testing.T) {
	fl := NewFlow(5)
	fmt.Println(fl.Size())
	fl.RPush("htwiu4bhnyimn")

	fmt.Println(fl.Size())
	fl.LPush("fewnignr")

	fmt.Println(fl.Size())
	fl.RPush("sfewiqughrwughr")
	fl.LPush("frewqgutbtr")
	fl.RPush("fiqruwght")
	fl.RPop()
	fl.RPop()
	fmt.Println(fl.Size())
}

func TestValues(t *testing.T) {
	fl := NewFlow(5)
	fmt.Println(fl.Values())
	fl.RPush("htwiu4bhnyimn")
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}

	fl.LPush("fewnignr")
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}

	fl.RPush("sfewiqughrwughr")
	fl.LPush("frewqgutbtr")
	fl.RPush("fiqruwght")
	fl.RPop()
	
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fl.RPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
}

func TestClear(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush("htwiu4bhnyimn")
	fl.LPush("fewnignr")
	fl.RPush("sfewiqughrwughr")
	fl.LPush("frewqgutbtr")
	fl.RPush("fiqruwght")
	
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
	fl.Clear()

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(fl.Values()[val])
	}
	fmt.Println(fl.Size())
}
