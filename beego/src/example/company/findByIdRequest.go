package company

type FindByIdRequest struct {
	id int
}

func (this FindByIdRequest) NewFindByIdRequest(id int) *FindByIdRequest {
	request := new(FindByIdRequest)
	request.id = id
	return request
}

func (this FindByIdRequest) Id() int {
	return this.id
}
