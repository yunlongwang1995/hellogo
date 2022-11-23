package __container

import (
	"fmt"
	"testing"
)

/**
s[begin:end:cap_end]
*/
func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := s1[1:3:5]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}

type School struct {
	room    []int
	teacher *Teacher
}

type Teacher struct {
	Age int
}

func (s School) ChangeName() {
	if s.teacher != nil {
		s.teacher.Age = 18
	}
}

func (s School) SetRoom1(rooms []int) {
	fmt.Printf("%p\n", s.room)
	rooms = append(rooms, 3)
	s.room = rooms
	s.room = append(s.room, 99)
}

func (s School) GetRoom1() []int {
	if s.room != nil {
		return s.room
	}
	return nil
}

func (s *School) SetRoom2(rooms []int) {
	fmt.Printf("%p\n", s)
	rooms = append(rooms, 3)
	s.room = rooms
}

func (s *School) GetRoom2() []int {
	if s.room != nil {
		return s.room
	}
	return nil
}

func TestStructPtr(t *testing.T) {
	s := School{teacher: &Teacher{Age: 12}}
	s.ChangeName()
	fmt.Println(s.teacher.Age)
}

func TestInputSlice1(t *testing.T) {
	s := School{room: []int{8}}
	rooms := []int{1, 2}
	fmt.Printf("%p\n", s.room)
	s.SetRoom1(rooms)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom1())
}

func TestInputSlice1_1(t *testing.T) {
	s := &School{}
	rooms := []int{1, 2}
	fmt.Printf("%p\n", s)
	s.SetRoom1(rooms)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom1())
}

func TestInputSlice2(t *testing.T) {
	s := School{}
	rooms := []int{1, 2}
	fmt.Printf("%p\n", &s)
	s.SetRoom2(rooms)
	rooms = append(rooms, 4)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom2())

	fmt.Printf("%p\n", &s)
	s.SetRoom2(rooms)

	fmt.Println(rooms)
	fmt.Println(s.GetRoom2())
}

func TestInputSlice2_1(t *testing.T) {
	s := &School{}
	rooms := []int{1, 2}
	fmt.Printf("%p\n", s)
	s.SetRoom2(rooms)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom2())
}

func TestReturnSlice1(t *testing.T) {
	s := School{}
	s.room = append(s.room, 1)
	fmt.Println(s.GetRoom1())
	rooms := s.GetRoom1()
	rooms = append(rooms, 2)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom1())
}

func TestReturnSlice2(t *testing.T) {
	s := School{}
	s.room = make([]int, 0, 10)
	s.room = append(s.room, 1)
	fmt.Println(s.GetRoom2())
	rooms := s.GetRoom2()
	rooms = append(rooms, 2)
	fmt.Printf("%p\n", s.room)
	fmt.Printf("%p\n", rooms)
	fmt.Println(rooms)
	fmt.Println(s.GetRoom2())

	fmt.Println(cap(s.GetRoom2()))
	fmt.Println(cap(rooms))
}

func TestSliceCopy(t *testing.T) {
	s1 := make([]int, 10)
	s2 := s1
	s1[0] = 1
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	s1 = append(s1, 13)
	s1 = append(s1, 14)
	s1 = append(s1, 15)

	fmt.Println(s2[0])
	fmt.Println(len(s2))
}
