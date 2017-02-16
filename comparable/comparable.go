package comparable

import "errors"

// "errors"

type ComparableList []Comparable

type Comparable interface {
	// Equal(Comparable) (bool, error)
	Mismatches(Comparable) (Mismatches, error)
	Children() (ComparableItems, error)
	Scope() (CollisionScope, error)
	// CanCompare() (bool, error)
}

type ComparableNamedItems []ComparableItemWithName
type ComparableItems []Comparable
type ComparableItemWithName interface {
	Comparable
	ItemWithName
}

type Collision struct {
	Elements   ComparableNamedItems
	Parent     Comparable
	Mismatches Mismatches
	Scope      CollisionScope
}
type Collisions []Collision
type Mismatches []string
type CollisionScope interface{}

func FindCollisions(c1 ComparableItemWithName, c2 ComparableItemWithName) (collisions Collisions, err error) {
	if !canCompare(c1, c2) {
		return nil, errors.New("can not comare")
	}

	mismatches, err := c1.Mismatches(c2)
	if err != nil {
		return nil, err
	}

	if len(mismatches) > 0 {
		collision, err := NewCollision(ComparableNamedItems{c1, c2}, mismatches)
		if err != nil {
			return nil, err
		}
		collisions = append(collisions, collision)
	}

	children1, err := c1.Children()
	if err != nil {
		return nil, err
	}

	children2, err := c2.Children()
	if err != nil {
		return nil, err
	}

	childrenCollisions, err := FindCollisionsInItems(children1.toComparableNamedItems(), children2.toComparableNamedItems())
	if err != nil {
		return nil, err
	}

	collisions = append(collisions, childrenCollisions...)

	return collisions, nil
}

func canCompare(c1 ComparableItemWithName, c2 ComparableItemWithName) bool {
	// return c1.GetType() != c2.GetType()
	return true
}

func FindCollisionsInItems(list1 ComparableNamedItems, list2 ComparableNamedItems) (Collisions, error) {
	// commonPairs, CompareByName()

	collisions := Collisions{}

	commonPairs, left1, left2 := CompareByName(list1.toNamedItems(), list2.toNamedItems())

	for _, pair := range commonPairs {
		first, ok := pair.First.(ComparableItemWithName)
		if !ok {
			return nil, errors.New("wrong conversion")
		}
		second, ok := pair.Second.(ComparableItemWithName)
		if !ok {
			return nil, errors.New("wrong conversion")
		}
		mismatches, err := first.Mismatches(second)
		if len(mismatches) > 0 {
			collision, err := NewCollision(ComparableNamedItems{first, second}, mismatches)
			if err != nil {
				return nil, err
			}
			collisions = append(collisions, collision)
		}
		childrenCollisions, err := FindCollisions(first, second)
		if err != nil {
			return nil, err
		}
		collisions = append(collisions, childrenCollisions...)
	}

	for _, item := range left1.toComparableNamedItems() {
		collision, err := NewCollision(ComparableNamedItems{item}, []string{"missing resource"})
		if err != nil {
			return nil, err
		}
		collisions = append(collisions, collision)
	}
	for _, item := range left2.toComparableNamedItems() {
		collision, err := NewCollision(ComparableNamedItems{item}, []string{"missing resource"})
		if err != nil {
			return nil, err
		}
		collisions = append(collisions, collision)
	}

	return collisions, nil
}

func NewCollision(items ComparableNamedItems, mismatches Mismatches) (Collision, error) {
	scope, err := items[0].Scope()
	if err != nil {
		return Collision{}, err
	}
	return Collision{
		Elements:   items,
		Mismatches: mismatches,
		Scope:      scope,
	}, nil
}