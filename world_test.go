package ecs

import (
	"reflect"
	"testing"
)

// TestWorldIntegration tests creating a new world and using it to create a few entities and update.
func TestWorldCreation(t *testing.T) {
	w := NewWorld()

	sysA := &TestSystemA{}
	w.AddSystem(sysA)
	sysB := &TestSystemB{}
	w.AddSystem(sysB)
	sysAB := &TestSystemAB{}
	w.AddSystem(sysAB)
	sysANotB := &TestSystemANotB{}
	w.AddSystem(sysANotB)

	a := &TestEntityA{BasicComponent: w.NewBasic()}
	a.TestAComponent = &TestAComponent{
		IncrementBy: 1,
		number:      20,
	}
	w.AddEntity(a)

	checkSystemEntities(t, a, sysA, 1)
	checkSystemEntities(t, a, sysB, 0)
	checkSystemEntities(t, a, sysAB, 0)
	checkSystemEntities(t, a, sysANotB, 1)

	b := &TestEntityB{BasicComponent: w.NewBasic()}
	b.TestBComponent = &TestBComponent{
		ExponentBy: 2,
		number:     1,
	}
	w.AddEntity(b)

	checkSystemEntities(t, b, sysA, 1)
	checkSystemEntities(t, b, sysB, 1)
	checkSystemEntities(t, b, sysAB, 0)
	checkSystemEntities(t, b, sysANotB, 1)

	ab := &TestEntityAB{BasicComponent: w.NewBasic()}
	ab.TestAComponent = &TestAComponent{
		IncrementBy: 5,
		number:      0,
	}
	ab.TestBComponent = &TestBComponent{
		ExponentBy: 5,
		number:     10,
	}
	w.AddEntity(ab)

	checkSystemEntities(t, ab, sysA, 2)
	checkSystemEntities(t, ab, sysB, 2)
	checkSystemEntities(t, ab, sysAB, 1)
	checkSystemEntities(t, ab, sysANotB, 1)

	abNotb := &TestEntityABnotB{BasicComponent: w.NewBasic()}
	abNotb.TestAComponent = &TestAComponent{IncrementBy: 4}
	abNotb.TestBComponent = &TestBComponent{}
	abNotb.TestNotBComponent = &TestNotBComponent{}
	w.AddEntity(abNotb)

	checkSystemEntities(t, a, sysA, 3)
	checkSystemEntities(t, a, sysB, 2)
	checkSystemEntities(t, a, sysAB, 2)
	checkSystemEntities(t, a, sysANotB, 1)

	aNota := &TestEntityAnotA{BasicComponent: w.NewBasic()}
	aNota.TestAComponent = &TestAComponent{IncrementBy: 6}
	aNota.TestNotAComponent = &TestNotAComponent{}
	w.AddEntity(aNota)

	checkSystemEntities(t, a, sysA, 3)
	checkSystemEntities(t, a, sysB, 2)
	checkSystemEntities(t, a, sysAB, 2)
	checkSystemEntities(t, a, sysANotB, 2)

	checkUpdateEntities(t, []TestEntity{a, ab, abNotb, aNota}, []int{20, 0, 0, 0})
	w.Update(2)
	checkUpdateEntities(t, []TestEntity{a, ab, abNotb, aNota}, []int{22, 10, 8, 6})
	w.Update(42)
	checkUpdateEntities(t, []TestEntity{a, ab, abNotb, aNota}, []int{24, 20, 16, 12})
	w.Update(0.000003)
	checkUpdateEntities(t, []TestEntity{a, ab, abNotb, aNota}, []int{26, 30, 24, 18})
}

func checkSystemEntities(t *testing.T, e Entity, system TestSystem, expected int) {
	if system.EntityCount() == expected {
		t.Logf("entity [[ %v ]] was successfully added to system [[ %v ]]", reflect.TypeOf(e), reflect.TypeOf(system))
	} else {
		t.Fatalf("entity [[ %v ]] was not added to system [[ %v ]] properly", reflect.TypeOf(e), reflect.TypeOf(system))
	}
}

func checkUpdateEntities(t *testing.T, entities []TestEntity, expected []int) {
	for i, e := range entities {
		if e.Number() == expected[i] {
			t.Logf("entity [[ %v ]] matched expected value [[ %v ]]", reflect.TypeOf(e), expected[i])
		} else {
			t.Fatalf("entity [[ %v ]] did not match expected value.\nWanted: [[ %v ]]\nGot: [[ %v ]]", reflect.TypeOf(e), expected[i], e.Number())
		}
	}
}

func BenchmarkWorld10Systems10Entities(b *testing.B) { benchHelper(10, 10, false, false, b) }

func BenchmarkWorld100Systems10Entities(b *testing.B) { benchHelper(100, 10, false, false, b) }

func BenchmarkWorld10Systems100Entities(b *testing.B) { benchHelper(10, 100, false, false, b) }

func BenchmarkWorld100Systems100Entities(b *testing.B) { benchHelper(100, 100, false, false, b) }

func BenchmarkWorld10Systems1000Entities(b *testing.B) { benchHelper(10, 1000, false, false, b) }

func BenchmarkWorld100Systems1000Entities(b *testing.B) { benchHelper(100, 1000, false, false, b) }

func BenchmarkWorld10Systems10000Entities(b *testing.B) { benchHelper(10, 10000, false, false, b) }

func BenchmarkWorld100Systems10000Entities(b *testing.B) { benchHelper(100, 10000, false, false, b) }

func BenchmarkWorldAddDuringUpdate(b *testing.B) { benchHelper(100, 10000, true, false, b) }

func BenchmarkWorldRemoveDuringUpdate(b *testing.B) { benchHelper(100, 10000, false, true, b) }

func BenchmarkWorldAddRemoveDuringUpdate(b *testing.B) { benchHelper(100, 10000, true, true, b) }

func benchHelper(sysCount, entCount int, add, remove bool, b *testing.B) {
	w := NewWorld()

	for i := 0; i < sysCount; i++ {
		w.AddSystem(&TestSystemA{})
	}

	for i := 0; i < entCount; i++ {
		ent := &TestEntityA{BasicComponent: w.NewBasic()}
		ent.TestAComponent = &TestAComponent{IncrementBy: 1}
		w.AddEntity(ent)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(1)
		ent := &TestEntityA{BasicComponent: w.NewBasic()}
		ent.TestAComponent = &TestAComponent{IncrementBy: 1}
		if add {
			w.AddEntity(ent)
		}
		if remove {
			w.RemoveEntity(ent)
		}
	}
}
