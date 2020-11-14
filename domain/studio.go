package domain

type Studio struct {
	Name string
}

func (self *Studio) Equals(studio Studio) bool {
	return self.Name == studio.Name
}
