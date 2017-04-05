// Code generated by immutableGen. DO NOT EDIT.

package _testFiles

//immutableVet:skipFile

import (
	"github.com/myitcv/immutable"
)

//
// intS is an immutable type and has the following template:
//
// 	[]int
//
type intS struct {
	theSlice []int
	mutable  bool
	__tmpl   _Imm_intS
}

var _ immutable.Immutable = new(intS)
var _ = new(intS).__tmpl

func newIntS(s ...int) *intS {
	c := make([]int, len(s))
	copy(c, s)

	return &intS{
		theSlice: c,
	}
}

func newIntSLen(l int) *intS {
	c := make([]int, l)

	return &intS{
		theSlice: c,
	}
}

func (m *intS) Mutable() bool {
	return m.mutable
}

func (m *intS) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *intS) Get(i int) int {
	return m.theSlice[i]
}

func (m *intS) AsMutable() *intS {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *intS) dup() *intS {
	resSlice := make([]int, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &intS{
		theSlice: resSlice,
	}

	return res
}

func (m *intS) AsImmutable(v *intS) *intS {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *intS) Range() []int {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *intS) WithMutable(f func(mi *intS)) *intS {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *intS) WithImmutable(f func(mi *intS)) *intS {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *intS) Set(i int, v int) *intS {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *intS) Append(v ...int) *intS {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *intS) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}

//
// Dummy is an immutable type and has the following template:
//
// 	struct {
// 		Name string
// 	}
//
type Dummy struct {
	_Name string

	mutable bool
	__tmpl  _Imm_Dummy
}

var _ immutable.Immutable = new(Dummy)
var _ = new(Dummy).__tmpl

func (s *Dummy) AsMutable() *Dummy {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *Dummy) AsImmutable(v *Dummy) *Dummy {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *Dummy) Mutable() bool {
	return s.mutable
}

func (s *Dummy) WithMutable(f func(si *Dummy)) *Dummy {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *Dummy) WithImmutable(f func(si *Dummy)) *Dummy {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}
func (s *Dummy) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *Dummy) Name() string {
	return s._Name
}

// SetName is the setter for Name()
func (s *Dummy) SetName(n string) *Dummy {
	if s.mutable {
		s._Name = n
		return s
	}

	res := *s
	res._Name = n
	return &res
}

//
// Dummy2 is an immutable type and has the following template:
//
// 	struct {
// 		name	[]byte
// 		other	*Dummy3
// 		mine	MyIntf
// 		another	MyType
// 	}
//
type Dummy2 struct {
	_name    []byte
	_other   *Dummy3
	_mine    MyIntf
	_another MyType

	mutable bool
	__tmpl  _Imm_Dummy2
}

var _ immutable.Immutable = new(Dummy2)
var _ = new(Dummy2).__tmpl

func (s *Dummy2) AsMutable() *Dummy2 {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *Dummy2) AsImmutable(v *Dummy2) *Dummy2 {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *Dummy2) Mutable() bool {
	return s.mutable
}

func (s *Dummy2) WithMutable(f func(si *Dummy2)) *Dummy2 {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *Dummy2) WithImmutable(f func(si *Dummy2)) *Dummy2 {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}
func (s *Dummy2) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	{
		v := s._other

		if v != nil && !v.IsDeeplyNonMutable(seen) {
			return false
		}
	}
	return true
}
func (s *Dummy2) name() []byte {
	return s._name
}

// setName is the setter for Name()
func (s *Dummy2) setName(n []byte) *Dummy2 {
	if s.mutable {
		s._name = n
		return s
	}

	res := *s
	res._name = n
	return &res
}
func (s *Dummy2) other() *Dummy3 {
	return s._other
}

// setOther is the setter for Other()
func (s *Dummy2) setOther(n *Dummy3) *Dummy2 {
	if s.mutable {
		s._other = n
		return s
	}

	res := *s
	res._other = n
	return &res
}
func (s *Dummy2) mine() MyIntf {
	return s._mine
}

// setMine is the setter for Mine()
func (s *Dummy2) setMine(n MyIntf) *Dummy2 {
	if s.mutable {
		s._mine = n
		return s
	}

	res := *s
	res._mine = n
	return &res
}
func (s *Dummy2) another() MyType {
	return s._another
}

// setAnother is the setter for Another()
func (s *Dummy2) setAnother(n MyType) *Dummy2 {
	if s.mutable {
		s._another = n
		return s
	}

	res := *s
	res._another = n
	return &res
}

//
// Dummy3 is an immutable type and has the following template:
//
// 	struct {
// 		other *Dummy2
// 	}
//
type Dummy3 struct {
	_other *Dummy2

	mutable bool
	__tmpl  _Imm_Dummy3
}

var _ immutable.Immutable = new(Dummy3)
var _ = new(Dummy3).__tmpl

func (s *Dummy3) AsMutable() *Dummy3 {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *Dummy3) AsImmutable(v *Dummy3) *Dummy3 {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *Dummy3) Mutable() bool {
	return s.mutable
}

func (s *Dummy3) WithMutable(f func(si *Dummy3)) *Dummy3 {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *Dummy3) WithImmutable(f func(si *Dummy3)) *Dummy3 {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}
func (s *Dummy3) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	{
		v := s._other

		if v != nil && !v.IsDeeplyNonMutable(seen) {
			return false
		}
	}
	return true
}
func (s *Dummy3) other() *Dummy2 {
	return s._other
}

// setOther is the setter for Other()
func (s *Dummy3) setOther(n *Dummy2) *Dummy3 {
	if s.mutable {
		s._other = n
		return s
	}

	res := *s
	res._other = n
	return &res
}
