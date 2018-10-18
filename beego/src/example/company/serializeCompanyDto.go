package company

type SerializeCompanyDto interface {
	Serialize(dto CompanyDto) []string
}
