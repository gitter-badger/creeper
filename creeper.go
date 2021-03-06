package creeper

import "io/ioutil"

func Open(path string) *Creeper {
	buf, _ := ioutil.ReadFile(path)
	raw := string(buf)
	return New(raw)
}

func New(raw string) *Creeper {
	f := Formatting(raw)
	return NewByFormatted(f)
}

func NewByFormatted(f *Formatted) *Creeper {
	c := new(Creeper)

	c.Nodes = f.Nodes
	for _, n := range c.Nodes {
		n.Creeper = c
	}

	c.Towns = f.Towns
	for _, t := range c.Towns {
		t.Creeper = c
	}

	cache := map[string]string{}
	c.Cache_Get = func(k string) (string, bool) {
		v, e := cache[k]
		return v, e
	}
	c.Cache_Set = func(k string, v string) {
		cache[k] = v
	}

	return c
}

type Creeper struct {
	Nodes []*Node
	Towns []*Town

	Cache_Get func(string) (string, bool)
	Cache_Set func(string, string)

	Node *Node
}

func (c *Creeper) Array(key string) *Creeper {
	if c.Node == nil {
		c.Node = c.Nodes[0]
	}
	c.Node = c.Node.SearchFlatScope(key)
	return c
}

func (c *Creeper) MString(key string) string {
	v, _ := c.String(key)
	return v
}

func (c *Creeper) String(key string) (string, error) {
	return c.Node.FirstChildNode.SearchFlatScope(key).Value()
}

func (c *Creeper) Each(cle func(*Creeper)) {
	for {
		cle(c)
		c.Next()
	}
}

func (c *Creeper) Next() *Creeper {
	c.Node.Inc()
	return c
}
