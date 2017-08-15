package event_driver

import (
	"testing"
)

type Student struct {
	Age int
}

func (s *Student) IncrAge(add int) int {
	s.Age += add
	return s.Age
}

func (s *Student) DecrAge(add int) int {
	s.Age -= add
	return s.Age
}

func BenchmarkCases(t *testing.B) {
	handler := NewEventHandler()
	handler.AddEvent(EVENT_TEST_ADD, func(x, y int) {
		//t.Logf("add event case 1: %d\n", x+y)
	})
	handler.AddEvent(EVENT_TEST_ADD, func(x, y int) {
		//t.Logf("add event case 2: %d\n", x*y)
	})
	handler.TriggerEvent(EVENT_TEST_ADD, 2, 4)

	handler.AddEvent(EVENT_TEST_STRUCT, func(x int, st *Student) {
		st.IncrAge(x)
		//t.Logf("add event case 3: %d\n", st.Age)
	})
	stu := &Student{Age: 10}
	for i := 1; i < t.N; i++ {
		handler.TriggerEvent(EVENT_TEST_STRUCT, 2, stu)
	} // cost less than 0.08s

	for i := 1; i < t.N; i++ {
		stu.IncrAge(2) // cost less than 0.005
	}
}

func BenchmarkCases2(t *testing.B) {
	handler := NewEventHandler2()
	handler.AddEvent(EVENT_TEST_ADD, func(param []interface{}) {
		//t.Logf("add event case 1: %d\n", param[0].(int)+param[1].(int))
	})
	handler.AddEvent(EVENT_TEST_ADD, func(param []interface{}) {
		//t.Logf("add event case 2: %d\n", param[0].(int)*param[1].(int))
	})
	handler.TriggerEvent(EVENT_TEST_ADD, 2, 4)

	handler.AddEvent(EVENT_TEST_STRUCT, func(param []interface{}) {
		x := param[0].(int)
		st := param[1].(*Student)
		st.IncrAge(x)
		//t.Logf("add event case 3: %d\n", st.Age)
	})
	stu := &Student{Age: 10}
	for i := 1; i < t.N; i++ {
		handler.TriggerEvent(EVENT_TEST_STRUCT, 2, stu)
	} // cost less than 0.08s

	for i := 1; i < t.N; i++ {
		stu.IncrAge(2) // cost less than 0.005
	}
}
