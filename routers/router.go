// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"wordbook/modules/users"
	"wordbook/modules/words"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&users.UserController{},
			),
		),
		beego.NSNamespace("/words",
			beego.NSInclude(
				&words.WordController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
