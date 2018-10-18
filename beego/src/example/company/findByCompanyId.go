package company

type FindByCompanyId interface {
	ByCompanyId(id int) ([]*CompanyDto, int64)
}
