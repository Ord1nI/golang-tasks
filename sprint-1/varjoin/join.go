//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	var str string
	for _, i := range args {
		str += sep+i
	}
	if len(str) > 0 && sep != "" {
		return str[1:]
	}
	return str
}
