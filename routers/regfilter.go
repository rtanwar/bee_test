package routers

/*
CMS 过滤器
*/
import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

/*
过滤器
*/
var sessionName = beego.AppConfig.String("SessionName")
var FilterUser = func(ctx *context.Context) {
	sess := ctx.Input.Session(sessionName)

	// UserName := ctx.Input.Session(sessionName)
	// fmt.Print("Session %s", v)
	// beego.Info("Session %s", v)
	var lowerUrl string = strings.ToLower(ctx.Request.RequestURI)
	// fmt.Println(lowerUrl)
	// if strings.Contains(lowerUrl, "/login")
	if sess == nil && !noCheckUrl(lowerUrl) {
		ctx.Redirect(302, "/login")
	}
	// m := sess.(map[string]interface{})
	// UserName := m["user"]
	// if userId == 0 && strings.Contains(lowerUrl, "/admin") {
	// 	if !noCheckUrl(lowerUrl) {
	// 		ctx.Redirect(301, "/admin/loginout")
	// 	}
	// }
}

/*排除的URL */
func noCheckUrl(lowerUrl string) bool {
	var flag bool = false
	// flag = flag || strings.Contains(lowerUrl, "/admin/center")
	// flag = flag || strings.Contains(lowerUrl, "/static")
	// flag = flag || strings.Contains(lowerUrl, "/admin/loginout")
	// flag = flag || lowerUrl == "/admin"

	flag = lowerUrl == "/login"
	return flag
}
