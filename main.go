package main

import (
	"fmt"
	"unsafe"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln             = fmt.Println
	fPt              = fmt.Print
	TryInvoke        = rflx.TryInvoke
	MustInvokeWithMW = rflx.MustInvokeWithMW
	TryInvokeWithMW  = rflx.TryInvokeWithMW
	InvokeRst        = rflx.InvokeRst
)

// IPerson :
type IPerson interface {
	ShowName(s1, s2 string) (string, string)
}

// Person :
type Person struct {
	Name string
	Age  int
	Fn   func()
}

// ShowName :
func (p *Person) ShowName(s1, s2 string) (string, string) {
	return s1 + " P " + s2 + " P " + p.Name, "GOOD JOB"
}

// ShowAge :
func (p *Person) ShowAge(added int) int {
	return p.Age + added
}

// Student :
type Student struct {
	Person
	score int
	MW    map[string]map[string][]interface{}
}

// func (s *Student) ShowName(s1, s2 string) string {
// 	return s1 + " S " + s2 + " S " + s.Name
// }

// ShowScore :
func (s *Student) ShowScore(str string) {
	fPt("MW: " + str + "   ")
	fPln(s.score)
}

// AddScore :
func (s *Student) AddScore(added int) {
	fPln(s.score + added)
}

// Show :
func Show(ip IPerson) {
	fPln(ip.ShowName("hello", "world"))
}

func main() {

	fPln(net.LocalIP())

	// return

	x := struct {
		a bool
		b int16
		c []int
	}{false, 2, []int{1, 2, 3}}
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x.a)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"

	arr := []int{1, 2, 3, 4, 5, 6}
	pa := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + unsafe.Sizeof(arr[0])))
	*pa = 33
	fmt.Println(arr)
	
	// return

	// ----------------------------------------------------- //

	s := &Student{
		Person: Person{
			Name: "HAOHAIDONG",
			Age:  22,
		},
		score: 100,
		MW: map[string]map[string][]interface{}{
			"ShowScore": {
				"*":        {"$1"},
				"ShowName": {"$@"},
			},
			// "AddScore": {
			// 	"$@":       {1000},
			// 	"ShowName": {500},
			// },
		},
	}

	fn.EnableLog2F(true, "./a.log")

	fPln(" ------------------------------------------- ")
	ret, ok := TryInvoke(s, "ShowName", "1", "Yanlimeng")
	fPln(ret, ok)
	fPln(" ------------------------------------------- ")
	Show(s)
	fPln(" ------------------------------------------- ")
	fPln(MustInvokeWithMW(s, "ShowName", "Great", "haohaidong"))
	fPln(" ------------------------------------------- ")
	results, ok := TryInvokeWithMW(s, "ShowName", "Great", "YANLIMENG")
	if ok {
		name := InvokeRst(results, 0).(string)
		msg := InvokeRst(results, 1).(string)
		fPln(name)
		fPln(msg)
	}

	fn.EnableLog2F(false, "")
}
