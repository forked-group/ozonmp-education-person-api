// Code generated by "stringer -type=Sex"; DO NOT EDIT.

package education

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Female-1]
	_ = x[Male-2]
}

const _Sex_name = "FemaleMale"

var _Sex_index = [...]uint8{0, 6, 10}

func (i Sex) String() string {
	i -= 1
	if i < 0 || i >= Sex(len(_Sex_index)-1) {
		return "Sex(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Sex_name[_Sex_index[i]:_Sex_index[i+1]]
}
