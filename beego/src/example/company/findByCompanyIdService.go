package company

type FindByCompanyIdService struct {
	findByCompanyId FindByCompanyId
}

func (this FindByCompanyIdService) New(readModel FindByCompanyId) *FindByCompanyIdService {
	s := new(FindByCompanyIdService)
	s.findByCompanyId = readModel
	return s
}

func (this FindByCompanyIdService) Execute(r *FindByIdRequest) ([]*CompanyDto, int64) {
	return this.findByCompanyId.ByCompanyId(r.Id())
}

func (this FindByCompanyIdService) ThreadExecute(r *FindByIdRequest) ([]*CompanyDto, int64) {
	id := make(chan int)
	companyDtos := make(chan CompanyDtos)

	go this.byCompanyId(id, companyDtos)

	id <- r.Id()

	g := <-companyDtos

	return g.CompanyDto, g.Total
}

func (this FindByCompanyIdService) byCompanyId(id <-chan int, companyDtos chan<- CompanyDtos) {

	companyId := <-id

	g, t := this.findByCompanyId.ByCompanyId(companyId)
	var result CompanyDtos
	result.CompanyDto = g
	result.Total = t

	companyDtos <- result
}
