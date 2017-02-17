package fakes

import (
	"errors"

	. "github.com/Altoros/cf-audit/comparable"
)

type FakeComparable struct {
	// ComparableItemWithName
	Name         string
	StringValue  string
	IntValue     int
	ChildrenList []FakeComparable
	Parent       *FakeComparable
}

func (f FakeComparable) GetName() string {
	return f.Name
}

func (f1 FakeComparable) Mismatches(fc2 Comparable) (Mismatches, error) {
	f2, err := fromComparable(fc2)
	if err != nil {
		return nil, err
	}
	mismatches := Mismatches{}
	if f1.StringValue != f2.StringValue {
		mismatches = append(mismatches, "SrtingValue")
	}
	if f1.IntValue != f2.IntValue {
		mismatches = append(mismatches, "IntValue")
	}
	return mismatches, nil
}

func (f FakeComparable) Children() (ComparableItems, error) {
	items := ComparableItems{}
	for _, fakeComparable := range f.ChildrenList {
		items = append(items, fakeComparable)
	}
	return items, nil
}

func (f FakeComparable) Scope() (CollisionScope, error) {
	return f.Name, nil
}

// func NewNode(name string, v1 string, v2 int, children []FakeComparable) FakeComparable {
// 	return NewNode(name, v1, v2, children)
// }
func NewChildren(children ...FakeComparable) []FakeComparable {
	return children
}

func NewNode(name string, v1 string, v2 int, children []FakeComparable) FakeComparable {
	return FakeComparable{
		Name:         name,
		StringValue:  v1,
		IntValue:     v2,
		ChildrenList: children,
	}
}

func fromComparable(c Comparable) (FakeComparable, error) {
	result, ok := c.(FakeComparable)
	if !ok {
		return FakeComparable{}, errors.New("convert error")
	}
	return result, nil
}
