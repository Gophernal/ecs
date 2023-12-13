package ecs

type Entity interface {
	ID() uint64
	AppendChild(child Entity)
	RemoveChild(child Entity)
	Children() []Entity
	Descendents() []Entity
	SetParent(parent Entity)
}

func RemoveEntityFromSlice[E Entity](remove Entity, s []E) []E {
	d := -1
	for i, e := range s {
		if e.ID() == remove.ID() {
			d = i
		}
	}
	if d >= 0 {
		return append(s[:d], s[d+1:]...)
	}
	return s
}
