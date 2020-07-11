package rflx

import "testing"

// Iperson :
type Iperson interface {
	ShowName(s1, s2 string) string
}

type Person struct {
	Name string
	Age  int
	Fn   func()
}

func (p *Person) ShowName(s1, s2 string) string {
	return s1 + " P " + s2 + " P " + p.Name
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
	fPt(str + "   ")
	fPln(s.score)
}

func (s *Student) AddScore(added int) {
	fPln(s.score + added)
}

// Show :
func Show(ip Iperson) {
	fPln(ip.ShowName("hello", "world"))
}

func TestTryInvoke(t *testing.T) {
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

	ret, ok := TryInvoke(s, "ShowName", "1", "Yanlimeng")
	fPln(ret, ok)

	// Show(s)

	// fPln(MustInvokeWithMW(s, "ShowName", "Great", "haohaidong"))

	// results, ok, err := TryInvokeWithMW(s, "ShowName", "Great")
	// if FailOnErr("%v", err); ok {
	// 	Iname, err := InvRst(results, 0)
	// 	FailOnErr("%v", err)
	// 	name := Iname.(string)
	// 	fPln(name)
	// }
}
