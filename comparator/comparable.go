package comparator

import (
// "errors"
)

type ComparableList []Comparable

type Comparable interface {
	// Equal(*Comparable) (bool, error)
	// Children() ([]Comparable, error)
	// Parent() (*Comparable, error)
	FindCollisions() (Collisions, error)
}

type Collisions []Collision

type Collision struct {
	Elements []Comparable
	Parent   Comparable
}

// func FindCollisions(elements []Comparable) (Collisions, error) {
// 	if len(elements) < 2 {
// 		return nil, errors.New("Provide more than 1 element to compare")
// 	}
// 	firstElement := elements[0]
// 	secondElement := elements[1]
// 	collisions, err := firstElement.Collisions(secondElement)
// 	return collisions, err

// 	// firstElement := elements[0]
// 	// collisions := []Collision{}
// 	// for _, element := range elements[:1] {
// 	// 	equal, err := firstElement.Equal(element)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	if !equal {
// 	// 		parent, err := element.Parent()
// 	// 		if err != nil {
// 	// 			return nil, err
// 	// 		}
// 	// 		collisions = append(collisions, Collision{Elements: elements, Parent: parent})
// 	// 	}
// 	// }
// 	// return collisions, nil
// }
