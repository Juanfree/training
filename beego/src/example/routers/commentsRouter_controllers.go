package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["example/controllers:CompanyRead"] = append(beego.GlobalControllerRouter["example/controllers:CompanyRead"],
		beego.ControllerComments{
			Method:           "GetAllCompanies",
			Router:           `/api/v1/company`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id"),
			),
			Params: nil})

	beego.GlobalControllerRouter["example/controllers:CompanyRead"] = append(beego.GlobalControllerRouter["example/controllers:CompanyRead"],
		beego.ControllerComments{
			Method:           "GetByInvoiceId",
			Router:           `/api/v1/company/id/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["example/controllers:CompanyRead"] = append(beego.GlobalControllerRouter["example/controllers:CompanyRead"],
		beego.ControllerComments{
			Method:           "GetByInvoiceIdDummy",
			Router:           `/api/v1/company/id/:id/dummy`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["example/controllers:CompanyRead"] = append(beego.GlobalControllerRouter["example/controllers:CompanyRead"],
		beego.ControllerComments{
			Method:           "GetByInvoiceIdWithThreads",
			Router:           `/api/v1/company/id/:id/threads`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.InPath),
			),
			Params: nil})

}
