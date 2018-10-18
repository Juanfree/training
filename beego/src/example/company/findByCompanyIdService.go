package company

type FindByCompanyIdService struct {
	findByCompanyId FindByCompanyId
}

func (this FindByCompanyIdService) New(readModel FindByCompanyId) *FindByCompanyIdService {
	s := new(FindByCompanyIdService)
	s.findByCompanyId = readModel
	return s
}

func (this FindByCompanyIdService) Execute(id int) ([]*CompanyDto, int64) {
	return this.findByCompanyId.ByCompanyId(id)
}
