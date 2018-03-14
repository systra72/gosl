// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func TestQueue01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Queue01")

	s2q := func(val string) Qmember { return Qmember(val) }
	q2s := func(val interface{}) string { return val.(string) }

	guessedMaxSize := 20
	qu := NewQueue(guessedMaxSize)
	member := qu.Out()
	if member != nil {
		tst.Errorf("returned member should be nil in an empty Queue\n")
		return
	}
	qu.Debug = true

	// add
	io.PfYel("In(l)\n")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 0)
	qu.In(s2q("l"))
	chk.String(tst, "[l]", io.Sf("%v", qu.ring))
	chk.String(tst, "[l]", qu.String())
	chk.String(tst, q2s(qu.Front()), "l")
	chk.String(tst, q2s(qu.Back()), "l")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nIn(o)\n")
	qu.In(s2q("o"))
	chk.String(tst, "[l o]", io.Sf("%v", qu.ring))
	chk.String(tst, "[l o]", qu.String())
	chk.String(tst, q2s(qu.Front()), "l")
	chk.String(tst, q2s(qu.Back()), "o")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)

	io.PfYel("\nIn(v)\n")
	qu.In(s2q("v"))
	chk.String(tst, "[l o v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[l o v]", qu.String())
	chk.String(tst, q2s(qu.Front()), "l")
	chk.String(tst, q2s(qu.Back()), "v")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 3)

	// remove
	io.PfYel("\nOut(l)\n")
	res := q2s(qu.Out())
	chk.String(tst, "[l o v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[o v]", qu.String())
	chk.String(tst, res, "l")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)

	io.PfYel("\nOut(o)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[l o v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[v]", qu.String())
	chk.String(tst, res, "o")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nOut(v)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[l o v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[]", qu.String())
	chk.String(tst, res, "v")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 0)

	// try to remove more in empty queue
	io.PfYel("\nOut(nothing)\n")
	member = qu.Out()
	if member != nil {
		tst.Errorf("returned member should be nil in an empty Queue\n")
		return
	}
	chk.String(tst, "[]", qu.String())
	chk.Int(tst, "len(queue)", qu.Nmembers(), 0)

	// add again
	io.PfYel("\nIn(a)\n")
	qu.In(s2q("a"))
	chk.String(tst, "[a o v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[a]", qu.String())
	chk.String(tst, q2s(qu.Front()), "a")
	chk.String(tst, q2s(qu.Back()), "a")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nIn(b)\n")
	qu.In(s2q("b"))
	chk.String(tst, "[a b v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[a b]", qu.String())
	chk.String(tst, q2s(qu.Front()), "a")
	chk.String(tst, q2s(qu.Back()), "b")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)

	io.PfYel("\nOut(a)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[a b v]", io.Sf("%v", qu.ring))
	chk.String(tst, "[b]", qu.String())
	chk.String(tst, q2s(qu.Front()), "b")
	chk.String(tst, q2s(qu.Back()), "b")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nIn(a) again\n")
	qu.In(s2q("a"))
	chk.String(tst, "[a b a]", io.Sf("%v", qu.ring))
	chk.String(tst, "[b a]", qu.String())
	chk.String(tst, q2s(qu.Front()), "b")
	chk.String(tst, q2s(qu.Back()), "a")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)

	io.PfYel("\nIn(c)\n")
	qu.In(s2q("c"))
	chk.String(tst, "[c b a]", io.Sf("%v", qu.ring))
	chk.String(tst, "[b a c]", qu.String())
	chk.String(tst, q2s(qu.Front()), "b")
	chk.String(tst, q2s(qu.Back()), "c")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 3)

	io.PfYel("\nIn(x)\n")
	qu.In(s2q("x"))
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[b a c x]", qu.String())
	chk.String(tst, q2s(qu.Front()), "b")
	chk.String(tst, q2s(qu.Back()), "x")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 4)

	io.PfYel("\nOut(b)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[a c x]", qu.String())
	chk.String(tst, q2s(qu.Front()), "a")
	chk.String(tst, q2s(qu.Back()), "x")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 3)

	io.PfYel("\nOut(a)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[c x]", qu.String())
	chk.String(tst, q2s(qu.Front()), "c")
	chk.String(tst, q2s(qu.Back()), "x")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)

	io.PfYel("\nOut(c)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[x]", qu.String())
	chk.String(tst, q2s(qu.Front()), "x")
	chk.String(tst, q2s(qu.Back()), "x")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nOut(x)\n")
	res = q2s(qu.Out())
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[]", qu.String())
	chk.Int(tst, "len(queue)", qu.Nmembers(), 0)

	io.PfYel("\nOut(nothing)\n")
	chk.String(tst, "[b a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[]", qu.String())
	chk.Int(tst, "len(queue)", qu.Nmembers(), 0)

	io.PfYel("\nIn(i)\n")
	qu.In(s2q("i"))
	chk.String(tst, "[i a c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[i]", qu.String())
	chk.String(tst, q2s(qu.Front()), "i")
	chk.String(tst, q2s(qu.Back()), "i")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 1)

	io.PfYel("\nIn(j)\n")
	qu.In(s2q("j"))
	chk.String(tst, "[i j c x]", io.Sf("%v", qu.ring))
	chk.String(tst, "[i j]", qu.String())
	chk.String(tst, q2s(qu.Front()), "i")
	chk.String(tst, q2s(qu.Back()), "j")
	chk.Int(tst, "len(queue)", qu.Nmembers(), 2)
}
