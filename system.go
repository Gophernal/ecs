package ecs

type System interface {
	Update(dt float32) error
	New(w *World) error
	Priority() int
	Add(entity Entity)
	Remove(entity Entity)
}

type systems []System

func (s systems) Len() int { return len(s) }

func (s systems) Less(i, j int) bool { return s[i].Priority() > s[j].Priority() }

func (s systems) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
