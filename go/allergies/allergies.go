package allergies

var list = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(allergies uint) (res []string) {
	for idx, item := range list {
		if allergies&(1<<idx) == (1 << idx) {
			res = append(res, item)
		}
	}

	return
}

func AllergicTo(allergies uint, allergen string) bool {
	var idx int
	for i, item := range list {
		if item == allergen {
			idx = i
			break
		}
	}
	return allergies&(1<<idx) == (1 << idx)
}
