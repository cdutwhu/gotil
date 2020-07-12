package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln             = fmt.Println
	fPt              = fmt.Print
	TryInvoke        = rflx.TryInvoke
	MustInvokeWithMW = rflx.TryInvokeWithMW
	TryInvokeWithMW  = rflx.TryInvokeWithMW
	InvokeRst        = rflx.InvokeRst
)

// Iperson :
type Iperson interface {
	ShowName(s1, s2 string) (string, string)
}

type Person struct {
	Name string
	Age  int
	Fn   func()
}

func (p *Person) ShowName(s1, s2 string) (string, string) {
	return s1 + " P " + s2 + " P " + p.Name, "GOOD JOB"
}

func (p *Person) ShowAge(added int) int {
	return p.Age + added
}

type Student struct {
	Person
	score int
	MW    map[string]map[string][]interface{}
}

// func (s *Student) ShowName(s1, s2 string) string {
// 	return s1 + " S " + s2 + " S " + s.Name
// }

func (s *Student) ShowScore(str string) {
	fPt("MW: " + str + "   ")
	fPln(s.score)
}

func (s *Student) AddScore(added int) {
	fPln(s.score + added)
}

// Show :
func Show(ip Iperson) {
	fPln(ip.ShowName("hello", "world"))
}

func main() {
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

	fn.SetLog("./a.log")

	fPln(" ------------------------------------------- ")
	ret, ok := TryInvoke(s, "ShowName", "1", "Yanlimeng", "more")
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
}
