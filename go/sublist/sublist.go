package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	var len1 = len(l1)
	var len2 = len(l2)
	if len1 == len2 {
		return IsEqual(l1, l2)
	} else if len1 < len2 {
		if IsSub(l1, l2) {
			return RelationSublist
		} else {
			return RelationUnequal
		}
	} else {
		if IsSub(l2, l1) {
			return RelationSuperlist
		} else {
			return RelationUnequal
		}
	}
}

func IsEqual(l1, l2 []int) Relation {
	var res = RelationEqual
	for k, v := range l1 {
		if v != l2[k] {
			res = RelationUnequal
			break
		}
	}
	return res
}

func IsSub(sub, all []int) bool {
	sub_l, all_l := len(sub), len(all)

	for i := 0; i <= all_l-sub_l; i++ {
		match := true
		for j := 0; j < sub_l; j++ {
			if all[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}
