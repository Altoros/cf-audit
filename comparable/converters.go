package comparable

import "fmt"

func (source ComparableNamedItems) toNamedItems() NamedItems {
	targetList := make(NamedItems, len(source))
	for i, item := range source {
		converted, ok := item.(ItemWithName)
		if !ok {
			fmt.Println("can not perform convertion")
		}
		targetList[i] = converted
	}
	return targetList
}

func (source NamedItems) toComparableNamedItems() ComparableNamedItems {
	targetList := make(ComparableNamedItems, len(source))
	for i, item := range source {
		converted, ok := item.(ComparableItemWithName)
		if !ok {
			fmt.Println("can not perform convertion")
		}
		targetList[i] = converted
	}
	return targetList
}

func (source ComparableItems) toComparableNamedItems() ComparableNamedItems {
	targetList := make(ComparableNamedItems, len(source))
	for i, item := range source {
		converted, ok := item.(ComparableItemWithName)
		if !ok {
			fmt.Println("can not perform convertion")
		}
		targetList[i] = converted
	}
	return targetList
}
