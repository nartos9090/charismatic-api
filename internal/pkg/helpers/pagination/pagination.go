package helpers_pagination

func (v *Pagination) Parse() {
	if v.Limit < 0 {
		v.Limit = 9999
	} else if v.Limit == 0 {
		v.Limit = 10
	}

	if v.Page <= 0 {
		v.Page = 1
	}

	v.Offset = (v.Page - 1) * v.Limit
}
