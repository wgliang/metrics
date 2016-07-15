package poetry

import (
	"testing"
	"fmt"
)

func cToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}

func TestNewflow(t *testing.T) {
	fl := NewFlow(5)

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}

func TestRPush(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush(([]byte)("htwiu4bhnyimn"))
	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.RPush(([]byte)("fiqruwght"))

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}

func TestLPop(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush(([]byte)("htwiu4bhnyimn"))
	fl.LPush(([]byte)("fewnignr"))
	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("frewqgutbtr"))
	fl.RPush(([]byte)("fiqruwght"))
	fl.LPop()
	fl.LPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}

func TestLPush(t *testing.T) {
	fl := NewFlow(5)
	fl.LPush(([]byte)("htwiu4bhnyimn"))
	fl.LPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("fiqruwght"))

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}

func TestRPop(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush(([]byte)("htwiu4bhnyimn"))
	fl.LPush(([]byte)("fewnignr"))
	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("frewqgutbtr"))
	fl.RPush(([]byte)("fiqruwght"))
	fl.RPop()
	fl.RPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}

func TestSize(t *testing.T) {
	fl := NewFlow(5)
	fmt.Println(fl.Size())
	fl.RPush(([]byte)("htwiu4bhnyimn"))

	fmt.Println(fl.Size())
	fl.LPush(([]byte)("fewnignr"))

	fmt.Println(fl.Size())
	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("frewqgutbtr"))
	fl.RPush(([]byte)("fiqruwght"))
	fl.RPop()
	fl.RPop()
	fmt.Println(fl.Size())
}

func TestValues(t *testing.T) {
	fl := NewFlow(5)
	fmt.Println(fl.Values())
	fl.RPush(([]byte)("htwiu4bhnyimn"))
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}

	fl.LPush(([]byte)("fewnignr"))
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}

	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("frewqgutbtr"))
	fl.RPush(([]byte)("fiqruwght"))
	fl.RPop()
	
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fl.RPop()
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
}

func TestClear(t *testing.T) {
	fl := NewFlow(5)
	fl.RPush(([]byte)("htwiu4bhnyimn"))
	fl.LPush(([]byte)("fewnignr"))
	fl.RPush(([]byte)("sfewiqughrwughr"))
	fl.LPush(([]byte)("frewqgutbtr"))
	fl.RPush(([]byte)("fiqruwght"))
	
	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
	fl.Clear()

	for val := 0; val < len(fl.Values()); val++ {
		// n := bytes.Index(, []byte{0})
		// s := string((fl.Values()[val])[:n])
		fmt.Println(cToGoString(fl.Values()[val]))
	}
	fmt.Println(fl.Size())
}
