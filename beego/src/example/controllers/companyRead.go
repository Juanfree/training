package controllers

import (
	"example/company"
	"github.com/astaxie/beego"
)

type CompanyRead struct {
	beego.Controller
}

// @router /api/v1/company [get]
func (this *CompanyRead) GetAllCompanies(id int) {

	elasticReadModel := new(company.ElasticFindAllCompanies)

	companyDto, total := elasticReadModel.All()
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total

	this.Data["json"] = m
	this.ServeJSON()
}

// @router /api/v1/company/id/:id [get]
func (this *CompanyRead) GetByInvoiceId(id int) {

	//this.dummyReturn(id)

	this.elasticReturn(id)
}

func (this *CompanyRead) dummyReturn(id int) {
	var r company.FindByIdRequest
	request := r.NewFindByIdRequest(id)
	var dummyReadModel company.DummyFindByCompanyId
	var s company.FindByCompanyIdService
	applicationService := s.New(dummyReadModel)
	companyDto, total := applicationService.Execute(request.Id())
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total
	this.Data["json"] = m
	this.ServeJSON()
}

func (this *CompanyRead) elasticReturn(id int) {
	var elasticReadModel company.ElasticFindByCompanyId
	companyDto, total := elasticReadModel.ByCompanyId(id)
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total
	this.Data["json"] = m
	this.ServeJSON()
}
