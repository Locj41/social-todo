package common

type Paging struct {
	Page  int   `form:"page"`
	Limit int   `form:"limit"`
	Total int64 `form:"-"`
}


func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}