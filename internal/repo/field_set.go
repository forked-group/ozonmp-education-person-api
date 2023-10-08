package repo

import "math"

const AllFields = math.MaxUint64 // usage: CustomType(AllFields)

type FieldSet[F ~uint64] uint64

func (m FieldSet[F]) Includes(field F) bool {
	return F(m)&field != 0
}
