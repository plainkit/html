package blox

// Non-global content helpers (unified for all tags that accept children)
type TxtOpt struct {
	s string
}

func Text(s string) TxtOpt {
	return TxtOpt{s}
}

// T is an alias for Text to reduce verbosity
func T(s string) TxtOpt { return Text(s) }

type UnsafeTxtOpt struct {
	s string
}

func UnsafeText(s string) UnsafeTxtOpt {
	return UnsafeTxtOpt{s}
}

type ChildOpt struct {
	c Component
}

func Child(c Component) ChildOpt {
	return ChildOpt{c}
}

// C is an alias for Child to reduce verbosity
func C(c Component) ChildOpt { return Child(c) }
