// Code generated by "stringer -type=Verbosity"; DO NOT EDIT.

package resp

import "fmt"

const _Verbosity_name = "NormalHiddenInteresting"

var _Verbosity_index = [...]uint8{0, 6, 12, 23}

func (i Verbosity) String() string {
	if i < 0 || i >= Verbosity(len(_Verbosity_index)-1) {
		return fmt.Sprintf("Verbosity(%d)", i)
	}
	return _Verbosity_name[_Verbosity_index[i]:_Verbosity_index[i+1]]
}
