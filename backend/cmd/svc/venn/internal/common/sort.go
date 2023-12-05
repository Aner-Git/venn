package common

type SortBy struct {
	ID   string
	Desc bool
}

func (sb *SortBy) ByName() string {
	return sb.ID
}

func (sb *SortBy) GetOrder() string {
	if sb.Desc {
		return "DESC"
	}
	return "ASC"
}
