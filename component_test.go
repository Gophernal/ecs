package ecs

type TestAComponent struct {
	IncrementBy int
	number      int
}

func (c *TestAComponent) A() {
	c.number += c.IncrementBy
}

func (c *TestAComponent) GetTestAComponent() *TestAComponent { return c }

func (c *TestAComponent) Number() int { return c.number }

type TestBComponent struct {
	ExponentBy int
	number     int
}

func (c *TestBComponent) B() {
	c.number *= c.ExponentBy
}

func (c *TestBComponent) GetTestBComponent() *TestBComponent { return c }

type TestNotAComponent struct{}

func (c *TestNotAComponent) UnA() {}

func (c *TestNotAComponent) GetNotTestAComponent() *TestNotAComponent { return c }

type TestNotBComponent struct{}

func (c *TestNotBComponent) UnB() {}

func (c *TestNotBComponent) GetNotTestBComponent() *TestNotBComponent { return c }
