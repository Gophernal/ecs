package ecs

type TestSystemAAble interface {
	BasicFace
	AFace
}

type NotTestSystemAAble interface {
	NotAFace
}

type TestSystemBAble interface {
	BasicFace
	BFace
}

type NotTestSystemBAble interface {
	NotBFace
}

type TestSystemABAble interface {
	BasicFace
	AFace
	BFace
}

type NotTestSystemABAble interface {
	NotAFace
	NotBFace
}

type AFace interface {
	GetTestAComponent() *TestAComponent
}

type BFace interface {
	GetTestBComponent() *TestBComponent
}

type NotAFace interface {
	GetNotTestAComponent() *TestNotAComponent
}

type NotBFace interface {
	GetNotTestBComponent() *TestNotBComponent
}

// TestSystem is a system that is testable.
type TestSystem interface {
	System
	EntityCount() int
}

// TestComponent is a component that is testable.
type TestEntity interface {
	Number() int
}
