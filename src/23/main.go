package main

import (
	_ "embed"
	"math"
	"regexp"
)

const extension = "DCBADBAC"

type apod rune
type room []apod
type rooms [4]room
type hall [11]apod
type state struct {
	rooms  rooms
	hall   hall
	energy int
	size   int
}

//go:embed input.txt
var input string

// Somehow there are some inputs that don't work, but part 1 & 2 work for my input.
// For the example input only part 1 gives the correct answer.
func main() {
	s1 := state{parseRooms(input, false), hall{}, 0, 2}
	s2 := state{parseRooms(input, true), hall{}, 0, 4}

	println(s1.organize().energy)
	println(s2.organize().energy)
}

func (s state) organize() state {
	if s.settled() {
		return s
	}

	ms := state{s.rooms, s.hall, math.MaxInt, s.size}

	for i, a := range s.hall {
		if a == 0 {
			continue
		}

		t := a.home()
		tr := s.rooms.get(t)

		if tr.only(a) && s.hall.reachable(i, toHallPos(t), false) {
			return s.slot(i, t).organize()
		}
	}

	for i, r := range s.rooms {
		if r.only(toApod(i)) {
			continue
		}

		a := r.peek()
		f, t := toHallPos(i), a.home()

		if s.hall.reachable(f, toHallPos(t), true) && s.rooms.get(t).only(a) {
			return s.moveOut(i, f).slot(f, t).organize()
		}
	}

	for i, r := range s.rooms {
		if r.empty() || r.only(toApod(i)) {
			continue
		}

		for _, t := range s.hall.available(i) {
			ns := s.moveOut(i, t).organize()

			if ns.energy < ms.energy {
				ms = ns
			}
		}
	}

	return ms
}

func (s state) moveOut(from, to int) state {
	rs, a := s.rooms.pop(from)
	d := abs(toHallPos(from) - to)
	e := a.cost() * (d + s.size - rs.get(from).len())

	return state{rs, s.hall.set(to, a), s.energy + e, s.size}
}

func (s state) slot(from, to int) state {
	a := s.hall[from]
	rs := s.rooms.add(to, a)
	d := abs(from - toHallPos(to))
	e := a.cost() * (d + 1 + s.size - rs.get(to).len())

	return state{rs, s.hall.remove(from), s.energy + e, s.size}
}

func (s state) settled() bool {
	for i, r := range s.rooms {
		if r.len() < s.size || !r.only(toApod(i)) {
			return false
		}
	}

	return true
}

func (rs rooms) get(i int) room { return rs[i] }
func (rs rooms) set(i int, r room) rooms {
	rs[i] = r
	return rs
}

func (rs rooms) pop(i int) (rooms, apod) {
	r, a := rs.get(i).pop()
	return rs.set(i, r), a
}

func (rs rooms) add(i int, a apod) rooms {
	return rs.set(i, rs.get(i).add(a))
}

func (r room) pop() (room, apod) { return r[0 : r.len()-1], r[r.len()-1] }
func (r room) peek() apod        { return r[r.len()-1] }
func (r room) len() int          { return len(r) }
func (r room) empty() bool       { return r.len() == 0 }
func (r room) add(a apod) room   { return append(r.copy(), a) }
func (r room) fill(a apod) room  { return append(room{a}, r...) }
func (r room) get(i int) apod    { return r[i] }

func (r room) copy() room {
	n := make(room, len(r))

	for i, a := range r {
		n[i] = a
	}

	return n
}

func (r room) only(a apod) bool {
	for _, o := range r {
		if o != a {
			return false
		}
	}

	return true
}

func (h hall) occupied(i int) bool   { return h[i] > 0 }
func (h hall) get(i int) apod        { return h[i] }
func (h hall) isRoomExit(i int) bool { return i > 0 && i < len(h)-1 && i%2 == 0 }

func (h hall) set(i int, a apod) hall {
	h[i] = a
	return h
}

func (h hall) remove(i int) hall {
	h[i] = 0
	return h
}

func (h hall) reachable(from int, to int, inc bool) bool {
	s, e := from+1, to

	if inc {
		s -= 1
	}

	if from > to {
		s, e = to, from-1

		if inc {
			s += 1
		}
	}

	for i := s; i <= e; i++ {
		if h.occupied(i) {
			return false
		}
	}

	return true
}

func (h hall) available(from int) []int {
	a := make([]int, 0)

	for i := from - 1; i >= 0; i-- {
		if h.occupied(i) {
			break
		}

		if !h.isRoomExit(i) {
			a = append(a, i)
		}
	}

	for i := from + 1; i < len(h); i++ {
		if h.occupied(i) {
			break
		}

		if !h.isRoomExit(i) {
			a = append(a, i)
		}

	}

	return a
}

func (a apod) home() int { return int(a - 'A') }
func (a apod) cost() int { return int(math.Pow(10, float64(a.home()))) }

func toHallPos(ri int) int { return 2 + (ri * 2) }
func toApod(ri int) apod   { return apod(ri + 'A') }
func abs(i int) int        { return int(math.Abs(float64(i))) }

func parseRooms(input string, extended bool) rooms {
	var rs rooms

	reg := regexp.MustCompile(`[ABCD]`)
	res := reg.FindAllString(input, 8)

	if extended {
		ext := reg.FindAllString(extension, 8)
		res = append(append(res[:4], ext...), res[4:]...)
	}

	for i, as := range res {
		ri := i % 4
		rs = rs.set(ri, rs[ri].fill(apod(as[0])))
	}

	return rs
}
