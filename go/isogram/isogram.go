package isogram

import "strings"

func IsIsogram(word string) bool {
    m := make(map[int]int)
	for _ ,  l:= range strings.ToUpper(word) {
        li:=int(l)
        if li>=65 && li <=90{
            m[li]+=1
            if m[li]>1{
                return false
            }
        }
    }

    return true
}
