package metrics

import "testing"

func BenchmarkRegistry(b *testing.B) {
	r := NewCollectry()
	r.Collector("foo", NewCounter())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Each(func(string, interface{}) {})
	}
}

func TestRegistry(t *testing.T) {
	r := NewCollectry()
	r.Collector("foo", NewCounter())
	i := 0
	r.Each(func(name string, iface interface{}) {
		i++
		if "foo" != name {
			t.Fatal(name)
		}
		if _, ok := iface.(Counter); !ok {
			t.Fatal(iface)
		}
	})
	if 1 != i {
		t.Fatal(i)
	}
	r.Uncollector("foo")
	i = 0
	r.Each(func(string, interface{}) { i++ })
	if 0 != i {
		t.Fatal(i)
	}
}

func TestRegistryDuplicate(t *testing.T) {
	r := NewCollectry()
	if err := r.Collector("foo", NewCounter()); nil != err {
		t.Fatal(err)
	}
	if err := r.Collector("foo", NewGauge()); nil == err {
		t.Fatal(err)
	}
	i := 0
	r.Each(func(name string, iface interface{}) {
		i++
		if _, ok := iface.(Counter); !ok {
			t.Fatal(iface)
		}
	})
	if 1 != i {
		t.Fatal(i)
	}
}

func TestRegistryGet(t *testing.T) {
	r := NewCollectry()
	r.Collector("foo", NewCounter())
	if count := r.Get("foo").(Counter).Count(); 0 != count {
		t.Fatal(count)
	}
	r.Get("foo").(Counter).Inc(1)
	if count := r.Get("foo").(Counter).Count(); 1 != count {
		t.Fatal(count)
	}
}

func TestRegistryGetOrCollector(t *testing.T) {
	r := NewCollectry()

	// First metric wins with GetOrCollector
	_ = r.GetOrCollector("foo", NewCounter())
	m := r.GetOrCollector("foo", NewGauge())
	if _, ok := m.(Counter); !ok {
		t.Fatal(m)
	}

	i := 0
	r.Each(func(name string, iface interface{}) {
		i++
		if name != "foo" {
			t.Fatal(name)
		}
		if _, ok := iface.(Counter); !ok {
			t.Fatal(iface)
		}
	})
	if i != 1 {
		t.Fatal(i)
	}
}

func TestRegistryGetOrCollectorWithLazyInstantiation(t *testing.T) {
	r := NewCollectry()

	// First metric wins with GetOrCollector
	_ = r.GetOrCollector("foo", NewCounter)
	m := r.GetOrCollector("foo", NewGauge)
	if _, ok := m.(Counter); !ok {
		t.Fatal(m)
	}

	i := 0
	r.Each(func(name string, iface interface{}) {
		i++
		if name != "foo" {
			t.Fatal(name)
		}
		if _, ok := iface.(Counter); !ok {
			t.Fatal(iface)
		}
	})
	if i != 1 {
		t.Fatal(i)
	}
}

func TestPrefixedChildRegistryGetOrCollector(t *testing.T) {
	r := NewCollectry()
	pr := NewPrefixedChildCollectry(r, "prefix.")

	_ = pr.GetOrCollector("foo", NewCounter)

	r.Each(func(name string, m interface{}) {
		if name != "prefix.foo" {
			t.Fatal(name)
		}
	})
}

func TestPrefixedRegistryGetOrCollector(t *testing.T) {
	r := NewPrefixedCollectry("prefix.")

	_ = r.GetOrCollector("foo", NewCounter)

	r.Each(func(name string, m interface{}) {
		if name != "prefix.foo" {
			t.Fatal(name)
		}
	})
}

func TestPrefixedRegistryCollector(t *testing.T) {
	r := NewPrefixedCollectry("prefix.")

	_ = r.Collector("foo", NewCounter)

	r.Each(func(name string, m interface{}) {
		if name != "prefix.foo" {
			t.Fatal(name)
		}
	})
}

func TestPrefixedRegistryUnCollector(t *testing.T) {
	r := NewPrefixedCollectry("prefix.")

	_ = r.Collector("foo", NewCounter)

	r.Each(func(name string, m interface{}) {
		if name != "prefix.foo" {
			t.Fatal(name)
		}
	})

	r.Uncollector("foo")

	i := 0
	r.Each(func(name string, m interface{}) {
		i++
	})

	if i != 0 {
		t.Fatal(i)
	}
}
