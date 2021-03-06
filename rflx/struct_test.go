package rflx

import "testing"

type User struct {
	Name  string
	Happy bool
	Age   int
}

func TestStruct2Env(t *testing.T) {
	// Struct2Env("MyError", "abc")

	user := &User{Name: "./../Frank", Happy: true, Age: 18}
	Struct2Env("MyUser", user)

	user1 := &User{}
	fPf("%+v\n", *user1)
	IUser2 := Env2Struct("MyUser", &User{})
	user2 := IUser2.(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}

func TestEnv2Struct(t *testing.T) {
	user1 := &User{}
	fPf("%+v\n", *user1)
	IUser2 := Env2Struct("MyUser", &User{})
	user2 := IUser2.(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}

func TestStruct2Map(t *testing.T) {
	s := &struct {
		A string
		B int
		c bool
		F func(string) string
	}{A: "aa", B: 22, c: false, F: func(str string) string { return str }}
	m := Struct2Map(s)
	fPln(m)

	// ERROR !!!
	// ks, vs := MapKVs(m)
	// fPln(ks)
	// fPln(vs)

	fPln("--------------------")

	// s1 := &struct {
	// 	A1 string
	// 	B1 string
	// 	C1 string
	// }{A1: "aa", B1: "22", C1: "false"}
	// m, err = Struct2Map(s1)
	// FailOnErr("%v", err)
	// fPln(m)

	// ks, vs, err := MapKVs(m)
	// FailOnErr("%v", err)
	// fPln(ks, vs)
}

func TestStructFields(t *testing.T) {
	s := &struct {
		A  string
		B  int
		c  bool
		Fn func(string) string
	}{A: "aa", B: 22, c: false, Fn: func(str string) string { return str }}
	fPln(StructFields(s))
}
