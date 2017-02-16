package comparable

type NamedItems []ItemWithName

type NamedItemsPairs []NamedItemsPair

type NamedItemsPair struct {
	First, Second ItemWithName
}

type ItemWithName interface {
	GetName() string
}

func CompareByName(list1 NamedItems, list2 NamedItems) (common NamedItemsPairs, left1 NamedItems, left2 NamedItems) {
	common = NamedItemsPairs{}
	hasCommon1 := make([]bool, len(list1))
	hasCommon2 := make([]bool, len(list2))

	for i1, item1 := range list1 {
		for i2, item2 := range list2 {
			if item1.GetName() == item2.GetName() {
				common = append(common, NamedItemsPair{First: item1, Second: item2})
				hasCommon1[i1] = true
				hasCommon2[i2] = true
				break
			}
		}
	}
	for i, hasCommon := range hasCommon1 {
		if !hasCommon {
			left1 = append(left1, list1[i])
		}
	}
	for i, hasCommon := range hasCommon2 {
		if !hasCommon {
			left2 = append(left2, list2[i])
		}
	}
	return common, left1, left2
}
