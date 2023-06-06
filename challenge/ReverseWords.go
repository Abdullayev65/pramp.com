package challenge

type span struct {
	start int
	end   int
	space bool
}

func ReverseWords(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	spans := make([]span, 0)
	temp := span{start: 0, space: arr[0] == " "}

	for i, v := range arr {
		if (temp.space && v != " ") || (!temp.space && v == " ") {
			temp.end = i
			spans = append(spans, temp)
			temp = span{start: i, space: v == " "}
		}
	}
	temp.end = len(arr)
	spans = append(spans, temp)

	res := make([]string, 0, len(arr))
	for i := len(spans) - 1; i >= 0; i-- {
		res = append(res, arr[spans[i].start:spans[i].end]...)
	}

	return res
}
