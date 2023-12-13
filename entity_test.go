package ecs

// TestEntityA has a Basic and an AComponent
type TestEntityA struct {
	*BasicComponent
	*TestAComponent
}

// TestEntityB has a Basic and a BComponent
type TestEntityB struct {
	*BasicComponent
	*TestBComponent
}

// TestEntityAB has a Basic, an AComponent, and a BComponent
type TestEntityAB struct {
	*BasicComponent
	*TestAComponent
	*TestBComponent
}

// TestEntityABnotB has a Basic, an AComponent, a BComponent and a NotBComponent
type TestEntityABnotB struct {
	*BasicComponent
	*TestAComponent
	*TestBComponent
	*TestNotBComponent
}

type TestEntityAnotA struct {
	*BasicComponent
	*TestAComponent
	*TestNotAComponent
}
