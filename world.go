package ecs

import (
	"reflect"
	"sort"
	"sync/atomic"
)

// World keeps track of the systems and passes the ones that implement the right interfaces to
// all the systems when added.
type World struct {
	entID   uint64
	systems systems
}

// NewWorld creates a new World
func NewWorld() *World {
	return &World{}
}

// NewBasic creates a new BasicComponent for use in entities. Every entity must have a BasicComponent.
// It implements the Entity interface!
func (w *World) NewBasic() *BasicComponent {
	return &BasicComponent{id: atomic.AddUint64(&w.entID, 1)}
}

// AddEntity adds an entity to the World.
func (w *World) AddEntity(e Entity) {
	for _, desc := range e.Descendents() {
		for _, sys := range w.systems {
			sys.Add(desc)
		}
	}
	for _, sys := range w.systems {
		sys.Add(e)
	}
}

func (w *World) implements(e Entity, i []interface{}) bool {
	for _, in := range i {
		if reflect.TypeOf(e).Implements(reflect.TypeOf(in)) {
			return true
		}
	}
	return false
}

// RemoveEntity removes and entity from the World.
func (w *World) RemoveEntity(e Entity) {
	for _, desc := range e.Descendents() {
		for _, sys := range w.systems {
			sys.Remove(desc)
		}
	}
	for _, sys := range w.systems {
		sys.Remove(e)
	}
}

// AddSystem adds a system to the World.
func (w *World) AddSystem(s System) {
	s.New(w)
	w.systems = append(w.systems, s)
	sort.Sort(w.systems)
}

// Update passes the systems a list of matching entities to act upon.
func (w *World) Update(dt float32) error {
	for _, sys := range w.systems {
		err := sys.Update(dt)
		if err != nil {
			return err
		}
	}
	return nil
}
