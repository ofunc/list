package list_test

import (
	"testing"

	"github.com/ofunc/list"
)

func TestNew(t *testing.T) {
	l := list.New(0, nil)
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail() != nil {
		t.FailNow()
	}
}

func TestLen(t *testing.T) {
	if list.Series(0, 1).Take(3).Len() != 3 {
		t.FailNow()
	}
}

func TestForce(t *testing.T) {
	var tailf func() *list.List
	x := 0
	tailf = func() *list.List {
		x++
		return list.New(x, tailf)
	}
	a := list.New(x, tailf).Take(3)
	b := a.Force().Force()
	if a != b {
		t.FailNow()
	}
	for ; a != nil && b != nil; a, b = a.Tail(), b.Tail() {
		if a.Head() != b.Head() {
			t.FailNow()
		}
	}
}

func TestEach(t *testing.T) {
	list.Repeat(0).Take(8).Each(func(x interface{}) {
		if x != 0 {
			t.FailNow()
		}
	})
}

func TestAll(t *testing.T) {
	ok := list.Repeat(0).Take(8).All(func(x interface{}) bool {
		return x == 0
	})
	if !ok {
		t.FailNow()
	}

	ok = list.Repeat(0).Take(8).All(func(x interface{}) bool {
		return x != 0
	})
	if ok {
		t.FailNow()
	}
}

func TestAny(t *testing.T) {
	ok := list.Series(0, 1).Take(8).Any(func(x interface{}) bool {
		return x == 3
	})
	if !ok {
		t.FailNow()
	}

	ok = list.Series(0, 1).Take(8).Any(func(x interface{}) bool {
		return x.(int) < 0
	})
	if ok {
		t.FailNow()
	}
}

func TestCons(t *testing.T) {
	a := new(list.List)
	b := a.Cons(0)
	if b.Head() != 0 {
		t.FailNow()
	}
	if b.Tail() != a {
		t.FailNow()
	}
}

func TestMap(t *testing.T) {
	l := list.Series(0, 1).Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	})
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 2 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 4 {
		t.FailNow()
	}

	l = (*list.List)(nil).Map(func(x interface{}) interface{} {
		return 1
	})
	if l != nil {
		t.FailNow()
	}
}

func TestFilter(t *testing.T) {
	l := list.Series(0, 1).Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 2 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 4 {
		t.FailNow()
	}

	l = list.Repeat(0).Take(8).Filter(func(x interface{}) bool {
		return false
	})
	if l != nil {
		t.FailNow()
	}
}

func TestFold(t *testing.T) {
	r := list.Series(0, 1).Take(8).Fold(0, func(r, x interface{}) interface{} {
		return r.(int) + x.(int)
	})
	if r != 28 {
		t.FailNow()
	}
}

func TestTake(t *testing.T) {
	l := list.Series(0, 1).Take(3)
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}
}

func TestDrop(t *testing.T) {
	l := list.Series(0, 1).Drop(3)
	if l.Head() != 3 {
		t.FailNow()
	}
}

func TestCut(t *testing.T) {
	l := list.Series(0, 1).Take(6).Cut(3)
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}

	a := list.Repeat(0)
	b := a.Cut(0)
	if a != b {
		t.FailNow()
	}
}

func TestTakeWhile(t *testing.T) {
	l := list.Series(0, 1).TakeWhile(func(x interface{}) bool {
		return x.(int) < 3
	})
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}

	l = (*list.List)(nil).TakeWhile(func(x interface{}) bool {
		return true
	})
	if l != nil {
		t.FailNow()
	}
}

func TestDropWhile(t *testing.T) {
	l := list.Series(0, 1).DropWhile(func(x interface{}) bool {
		return x.(int) < 3
	})
	if l.Head() != 3 {
		t.FailNow()
	}
}

func TestCutWhile(t *testing.T) {
	l := list.Series(0, 1).Take(6).CutWhile(func(x interface{}) bool {
		return x.(int) >= 3
	})
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}
}

func TestMake(t *testing.T) {
	l := list.Make(0, 1, 2)
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}
}

func TestRepeat(t *testing.T) {
	l := list.Repeat(0).Take(3)
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 0 {
		t.FailNow()
	}
}

func TestSeries(t *testing.T) {
	l := list.Series(0, 1).Take(3)
	if l.Len() != 3 {
		t.FailNow()
	}
	if l.Head() != 0 {
		t.FailNow()
	}
	if l.Tail().Head() != 1 {
		t.FailNow()
	}
	if l.Tail().Tail().Head() != 2 {
		t.FailNow()
	}
}

func TestConcat(t *testing.T) {
	a := list.Series(0, 1).Take(3)
	b := list.Repeat(3)
	c := list.Concat(a, b)
	if c.Head() != 0 {
		t.FailNow()
	}
	if c.Tail().Head() != 1 {
		t.FailNow()
	}
	if c.Tail().Tail().Head() != 2 {
		t.FailNow()
	}
	if c.Tail().Tail().Tail().Head() != 3 {
		t.FailNow()
	}
	if c.Tail().Tail().Tail().Tail().Head() != 3 {
		t.FailNow()
	}
}

func TestZip(t *testing.T) {
	var fibs *list.List
	fibs = list.New(1, func() *list.List {
		return list.Zip(func(xs ...interface{}) interface{} {
			return xs[0].(int) + xs[1].(int)
		}, fibs, fibs.Tail())
	}).Cons(0)
	if fibs.Head() != 0 {
		t.FailNow()
	}
	if fibs.Tail().Head() != 1 {
		t.FailNow()
	}
	if fibs.Tail().Tail().Head() != 1 {
		t.FailNow()
	}
	if fibs.Tail().Tail().Tail().Head() != 2 {
		t.FailNow()
	}
	if fibs.Tail().Tail().Tail().Tail().Head() != 3 {
		t.FailNow()
	}
	if fibs.Tail().Tail().Tail().Tail().Tail().Head() != 5 {
		t.FailNow()
	}

	l := list.Zip(nil, fibs, nil)
	if l != nil {
		t.FailNow()
	}
}
