package ecs

// TestSystemA is a system used for testing that takes entities that implement AAble
type TestSystemA struct {
	entities []TestEntityA
	lookup   map[uint64]TestEntityA
}

func (s *TestSystemA) New(w *World) error {
	s.lookup = make(map[uint64]TestEntityA)
	return nil
}

func (*TestSystemA) Priority() int { return 0 }

func (s *TestSystemA) Add(e Entity) {
	if _, ok := s.lookup[e.ID()]; ok {
		return //already has this entity
	}
	entity, ok := e.(TestSystemAAble)
	if !ok {
		return
	}
	_, ok = e.(NotTestSystemAAble)
	if ok {
		return
	}
	ent := TestEntityA{entity.GetBasicComponent(), entity.GetTestAComponent()}
	s.lookup[e.ID()] = ent
	s.entities = append(s.entities, ent)
}

func (s *TestSystemA) Remove(e Entity) {
	if _, ok := s.lookup[e.ID()]; !ok {
		return
	}
	delete(s.lookup, e.ID())
	s.entities = RemoveEntityFromSlice(e, s.entities)
}

func (s *TestSystemA) Update(dt float32) error {
	for _, entity := range s.entities {
		entity.A()
	}
	return nil
}

func (s *TestSystemA) EntityCount() int { return len(s.entities) }

// TestSystemB is a system used for testing that takes entities that implement BAble
type TestSystemB struct {
	entities []TestEntityB
}

func (*TestSystemB) New(w *World) error { return nil }

func (*TestSystemB) Priority() int { return 0 }

func (s *TestSystemB) Add(e Entity) {
	ent, ok := e.(TestSystemBAble)
	if !ok {
		return
	}
	_, ok = e.(NotTestSystemBAble)
	if ok {
		return
	}
	s.entities = append(s.entities, TestEntityB{ent.GetBasicComponent(), ent.GetTestBComponent()})
}

func (s *TestSystemB) Remove(e Entity) {
	s.entities = RemoveEntityFromSlice(e, s.entities)
}

func (s *TestSystemB) Update(dt float32) error {
	for _, entity := range s.entities {
		entity.B()
	}
	return nil
}

func (s *TestSystemB) EntityCount() int { return len(s.entities) }

// // TestSystemAB is a system used for testing that takes entities that implement both AAble and BAble
type TestSystemAB struct {
	entities []TestEntityAB
}

func (*TestSystemAB) New(w *World) error { return nil }

func (*TestSystemAB) Priority() int { return 1 }

func (s *TestSystemAB) Add(e Entity) {
	ent, ok := e.(TestSystemABAble)
	if !ok {
		return
	}
	_, ok = e.(NotTestSystemABAble)
	if ok {
		return
	}
	s.entities = append(s.entities, TestEntityAB{ent.GetBasicComponent(), ent.GetTestAComponent(), ent.GetTestBComponent()})
}

func (s *TestSystemAB) Remove(e Entity) {
	s.entities = RemoveEntityFromSlice(e, s.entities)
}

func (s *TestSystemAB) Update(dt float32) error {
	for _, entity := range s.entities {
		entity.A()
		entity.B()
	}
	return nil
}

func (s *TestSystemAB) EntityCount() int { return len(s.entities) }

// TestSystemANotB is a system used for testing that takes entites that implement AAble but not BAble
type TestSystemANotB struct {
	entities []TestEntityA
}

func (*TestSystemANotB) New(w *World) error { return nil }

func (*TestSystemANotB) Priority() int { return 2 }

func (s *TestSystemANotB) Add(e Entity) {
	ent, ok := e.(TestSystemAAble)
	if !ok {
		return
	}
	_, ok = e.(TestSystemBAble)
	if ok {
		return
	}
	s.entities = append(s.entities, TestEntityA{ent.GetBasicComponent(), ent.GetTestAComponent()})
}

func (s *TestSystemANotB) Remove(e Entity) {
	s.entities = RemoveEntityFromSlice(e, s.entities)
}

func (s *TestSystemANotB) Update(dt float32) error {
	for _, entity := range s.entities {
		entity.A()
	}
	return nil
}

func (s *TestSystemANotB) EntityCount() int { return len(s.entities) }
