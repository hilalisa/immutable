// My favourite license

package main

type myMap struct {
	theMap map[string]int

	mutable bool
}

func newMyMap() myMap {
	return myMap{
		theMap: make(map[string]int),
	}
}

func (m myMap) Len() int {
	return len(m.theMap)
}

func (m myMap) Get(k string) (int, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m myMap) AsMutable() myMap {
	res := m.dup()
	res.mutable = true

	return res
}

func (m myMap) dup() myMap {
	resMap := make(map[string]int, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := myMap{
		theMap: resMap,
	}

	return res
}

func (m myMap) AsImmutable() myMap {
	m.mutable = false

	return m
}

func (m myMap) Range() map[string]int {
	return m.theMap
}

func (m myMap) WithMutations(f func(mi myMap)) myMap {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable()

	// TODO optimise here if the maps are identical?

	return res
}

func (m myMap) Set(k string, v int) myMap {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	// TODO: work out a way of enabling this
	// if n, ok := m.theMap[k]; ok && n == v {
	// 	return m
	// }

	res := m.dup()
	res.theMap[k] = v

	return res
}

type Slice struct {
	theSlice []*string

	mutable bool
}

func NewSlice() Slice {
	return Slice{}
}

func (m Slice) Len() int {
	return len(m.theSlice)
}

func (m Slice) Get(i int) *string {
	return m.theSlice[i]
}

func (m Slice) AsMutable() Slice {
	res := m.dup()
	res.mutable = true

	return res
}

func (m Slice) dup() Slice {
	resSlice := make([]*string, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := Slice{
		theSlice: resSlice,
	}

	return res
}

func (m Slice) AsImmutable() Slice {
	m.mutable = false

	return m
}

func (m Slice) Range() []*string {
	return m.theSlice
}

func (m Slice) WithMutations(f func(mi Slice)) Slice {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable()

	// TODO optimise here if the maps are identical?

	return res
}

func (m Slice) Set(i int, v *string) Slice {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	// TODO: work out a way of enabling this
	// if m.theSlice[i] == v {
	// 	return m
	// }

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m Slice) Append(v ...*string) Slice {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	// TODO: work out a way of enabling this
	// if m.theSlice[i] == v {
	// 	return m
	// }

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}

type myStruct struct {
	_Name, _surname string `tag:"value"`
	_age            int    `tag:"age"`

	mutable bool
}

func newMyStruct() *myStruct {
	return &myStruct{}
}

func (s *myStruct) AsMutable() *myStruct {
	res := *s
	res.mutable = true
	return &res
}

func (s *myStruct) AsImmutable() *myStruct {
	s.mutable = false
	return s
}

func (s *myStruct) WithMutations(f func(si *myStruct)) *myStruct {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable()
	if *res == *s {
		return s
	}

	return res
}

// my field comment
//somethingspecial
/*

	Heelo

*/
func (s *myStruct) Name() string {
	return s._Name
}

func (s *myStruct) SetName(n string) *myStruct {
	// TODO: see if we can make this work
	// if n == s.Name {
	// 	return s
	// }

	if s.mutable {
		s._Name = n
		return s
	}

	res := *s
	res._Name = n
	return &res
}

// my field comment
//somethingspecial
/*

	Heelo

*/
func (s *myStruct) surname() string {
	return s._surname
}

func (s *myStruct) setSurname(n string) *myStruct {
	// TODO: see if we can make this work
	// if n == s.surname {
	// 	return s
	// }

	if s.mutable {
		s._surname = n
		return s
	}

	res := *s
	res._surname = n
	return &res
}

func (s *myStruct) age() int {
	return s._age
}

func (s *myStruct) setAge(n int) *myStruct {
	// TODO: see if we can make this work
	// if n == s.age {
	// 	return s
	// }

	if s.mutable {
		s._age = n
		return s
	}

	res := *s
	res._age = n
	return &res
}
