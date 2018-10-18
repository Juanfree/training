package controllers

import (
	"example/company"
	"fmt"
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
	this.elasticReturn(id)
}

func (this *CompanyRead) elasticReturn(id int) {

	elasticReadModel := new(company.ElasticFindByCompanyId)

	var s company.FindByCompanyIdService
	applicationService := s.New(elasticReadModel)

	var r company.FindByIdRequest
	companyDto, total := applicationService.Execute(r.NewFindByIdRequest(id))
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total
	this.Data["json"] = m
	this.ServeJSON()
}

// @router /api/v1/company/id/:id/dummy [get]
func (this *CompanyRead) GetByInvoiceIdDummy(id int) {
	this.dummyReturn(id)
}

func (this *CompanyRead) dummyReturn(id int) {
	var r company.FindByIdRequest
	request := r.NewFindByIdRequest(id)
	var dummyReadModel company.DummyFindByCompanyId
	var s company.FindByCompanyIdService
	applicationService := s.New(dummyReadModel)
	companyDto, total := applicationService.Execute(request)
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total
	this.Data["json"] = m
	this.ServeJSON()
}

// @router /api/v1/company/id/:id/threads [get]
func (this *CompanyRead) GetByInvoiceIdWithThreads(id int) {
	fmt.Print("principio")
	this.elasticReturnWithThreads(id)
}

func (this *CompanyRead) elasticReturnWithThreads(id int) {
	elasticReadModel := new(company.ElasticFindByCompanyId)
	var s company.FindByCompanyIdService
	applicationService := s.New(elasticReadModel)

	var r company.FindByIdRequest
	companyDto, total := applicationService.ThreadExecute(r.NewFindByIdRequest(id))
	m := make(map[string]interface{})
	m["data"] = companyDto
	m["total"] = total
	this.Data["json"] = m
	this.ServeJSON()
}
