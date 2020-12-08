package queue

import "testing"

func TestQueue(t *testing.T) {
	var q4s Queue
	q4s.Enqueue("abc")
	q4s.Enqueue("def")
	q4s.Enqueue("ghi")
	q4s.Enqueue("jkl")

	tmp := q4s.Copy()

	fPln(q4s.Sprint("-"))
	fPln("-------------")

	val, ok := q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	val, ok = q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	q4s.Enqueue("***")
	fPln(q4s, q4s.Sprint("-"))

	val, ok = q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	val, ok = q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	val, ok = q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	val, ok = q4s.Dequeue()
	fPln(val, ok, q4s, q4s.Sprint("-"))

	fPln(tmp.Sprint("/"))
}
