// Code generated by "stringer -type=Education"; DO NOT EDIT.

package education

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Preschool-1]
	_ = x[PrimaryGeneral-2]
	_ = x[BasicGeneral-3]
	_ = x[SecondaryGeneral-4]
	_ = x[SecondaryVocational-5]
	_ = x[Higher1-6]
	_ = x[Higher2-7]
	_ = x[Higher3-8]
}

const _Education_name = "PreschoolPrimaryGeneralBasicGeneralSecondaryGeneralSecondaryVocationalHigher1Higher2Higher3"

var _Education_index = [...]uint8{0, 9, 23, 35, 51, 70, 77, 84, 91}

func (i Education) String() string {
	i -= 1
	if i < 0 || i >= Education(len(_Education_index)-1) {
		return "Education(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Education_name[_Education_index[i]:_Education_index[i+1]]
}