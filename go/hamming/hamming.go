package hamming
import "errors"

func Distance(a, b string) (int, error) {
    if len(a)!=len(b){
		return -1,errors.New("different lengths")        
    }

	var distance int =0;
    for i:=0;i<len(a);i++ {
        if a[i]!=b[i]{
            distance++
        }
    }
	return distance, nil
}
