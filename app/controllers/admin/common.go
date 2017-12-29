package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/midoks/webcron/app/libs"
	"github.com/midoks/webcron/app/models"
	"strconv"
	"strings"
	"time"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type CommonController struct {
	beego.Controller
	controllerName string
	actionName     string
	pageSize       int
	user           *models.SysUser

	// xsrf data
	_xsrfToken string
	XSRFExpire int
	EnableXSRF bool
}

func (this *CommonController) uLog(behavior string) {
	models.LogAdd(this.user.Id, 1, behavior)
}

func (this *CommonController) D(args ...string) {
	if beego.AppConfig.String("runmode") == "dev" {
		for i := 0; i < len(args); i++ {
			this.Ctx.WriteString(args[i])
		}
		//this.StopRun()
	}
}

func (this *CommonController) initXSRF() {
	this.EnableXSRF = true
	this._xsrfToken = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	this.XSRFExpire = 3600 //过期时间，默认1小时
}

func (this *CommonController) initData() {

	this.Data["pageStartTime"] = time.Now()
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	//println(this.controllerName, this.actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName

}

func (this *CommonController) initMenuData() {

	//菜单导航
	menuNav, curMenuName, curMenuFuncName := models.FuncGetNav(this.controllerName, this.actionName)
	this.Data["menuNav"] = menuNav
	this.Data["curMenuName"] = curMenuName
	this.Data["curMenuFuncName"] = curMenuFuncName
}

//登录状态验证
func (this *CommonController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")

	if len(arr) == 2 {

		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.UserGetById(userId)
			if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
				this.user = user
				this.Data["user"] = user
			}
		}
	}

	// fmt.Println(this.user)

	//跳到登录页
	if (this.user == nil || this.user.Id == 0) && this.controllerName != "login" && (this.actionName != "out") {
		this.redirect(beego.URLFor("LoginController.Index"))
	}

	//跳到首页
	if (this.user != nil) && (this.controllerName == "login" && this.actionName == "index") {
		this.redirect(beego.URLFor("IndexController.Index"))
	}

}

// 是否POST提交
func (this *CommonController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 重定向
func (this *CommonController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//获取用户IP地址
func (this *CommonController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Layout = "layout/index.html"
	this.TplName = tplname
}

// 输出json
func (this *BaseController) retJson(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) retResult(code int, msg interface{}, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	out["data"] = data

	this.retJson(out)
}