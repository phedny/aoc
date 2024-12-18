package util

import "github.com/tidwall/btree"

func AStar[T any, V any](init T, cost func(T) int, fn func(T) (V, []T, bool)) (V, bool) {
	var front btree.Map[int, []T]
	front.Set(cost(init), []T{init})
	for {
		h, ts, has := front.DeleteAt(0)
		if !has {
			var zero V
			return zero, false
		}
		if len(ts) > 1 {
			front.Set(h, ts[1:])
		}
		v, newTs, done := fn(ts[0])
		// fmt.Println(h, ts[0], "-->", v, newTs, done)
		if done {
			return v, true
		}
		for _, t := range newTs {
			h := cost(t)
			ts, _ := front.Get(h)
			front.Set(h, append(ts, t))
		}
	}
}
