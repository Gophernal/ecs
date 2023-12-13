package ecs

type BasicComponent struct {
	id uint64

	parent   Entity
	children []Entity
}

func (c *BasicComponent) ID() uint64 {
	return c.id
}

func (c *BasicComponent) AppendChild(child Entity) {
	appendChild(c, child)
}

func appendChild[E Entity](c *BasicComponent, child E) {
	child.SetParent(c)
	c.children = append(c.children, child)
}

func (c *BasicComponent) RemoveChild(child Entity) {
	d := -1
	for i, v := range c.children {
		if v.ID() == child.ID() {
			d = i
			break
		}
	}
	if d >= 0 {
		c.children = append(c.children[:d], c.children[d+1:]...)
	}
}

func (c *BasicComponent) Children() []Entity {
	return c.children
}

func (c *BasicComponent) Descendents() []Entity {
	return descendents([]Entity{}, c, c)
}

func descendents(in []Entity, this, top Entity) []Entity {
	for _, child := range this.Children() {
		in = descendents(in, child, top)
	}
	if this.ID() == top.ID() {
		return in
	}
	return append(in, this)
}

func (c *BasicComponent) SetParent(parent Entity) {
	c.parent = parent
}

func (c *BasicComponent) GetBasicComponent() *BasicComponent { return c }
