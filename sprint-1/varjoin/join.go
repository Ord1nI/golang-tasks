//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	var str string
	var el_count int = len(args)
	for q, i := range args {
		str += i
		if q+1 != el_count {
			str += sep
		}
	}
	return str
}
