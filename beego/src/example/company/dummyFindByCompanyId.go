package company

type DummyFindByCompanyId struct {
}

func (e DummyFindByCompanyId) ByCompanyId(id int) ([]*CompanyDto, int64) {

	c1 := new(CompanyDto)
	c1.CompanyId = 1
	c1.CompanyName = "test1"
	c1.Location = "location1"
	c1.NumberOfTransactions = 11
	c1.LastUpdated = 11

	c2 := new(CompanyDto)
	c2.CompanyId = 2
	c2.CompanyName = "test2"
	c2.Location = "location2"
	c2.NumberOfTransactions = 22
	c2.LastUpdated = 2

	c3 := new(CompanyDto)
	c3.CompanyId = 3
	c3.CompanyName = "test3"
	c3.Location = "location3"
	c3.NumberOfTransactions = 33
	c3.LastUpdated = 33

	var s []*CompanyDto
	s = append(s, c1, c2, c3)

	return s, int64(len(s))
}
