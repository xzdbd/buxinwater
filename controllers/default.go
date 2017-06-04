package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/xzdbd/squeak-fuxinwater/models"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	var loginStatus bool
	user := models.Userinfo{}
	if err := c.ParseForm(&user); err != nil {
		beego.Error(err.Error())
	}

	if user.Username == "admin" && user.Password == "admin" {
		loginStatus = true
	}

	if loginStatus {
		c.Redirect("/index", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
}

func (c *MainController) Post() {
	year := c.GetString("year")
	month := c.GetString("month")
	stationIds := c.GetString("stationIds")
	updateRain := c.GetString("updateRain")
	beego.Trace("stationIds", stationIds)
	if year != "" && month != "" {
		rainInfos, err := models.GetRainInfoByMonth(year, month, stationIds)
		if err != nil {
			beego.Trace("error:", err)
			c.Data["json"] = err.Error()
		}
		if len(rainInfos) > 0 {
			var totalRain float64
			gridHtml := "<tbody>"
			gridHtml += "<tr>" + `<th class="success">站名</th>` + `<th class="success">雨量</th>` + "</tr>"
			for i := 0; i < len(rainInfos); i++ {
				totalRain += rainInfos[i].Rain
				gridHtml += `<tr>` +
					`<td>` + rainInfos[i].Name + `</td>` +
					`<td>` + strconv.FormatFloat(rainInfos[i].Rain, 'f', 2, 64) + `</td>` +
					`</tr>`
			}
			gridHtml += "</tbody>"
			gridHtml += `<tr><td colspan="2" style="text-align:right;">总计雨量:` + strconv.FormatFloat(totalRain, 'f', 2, 64) + `</td></tr>`
			c.Data["json"] = gridHtml
		}
	} else if year != "" && month == "" && updateRain != "true" {
		rainInfos, err := models.GetRainInfoByYear(year, stationIds)
		if err != nil {
			beego.Trace("error:", err)
			c.Data["json"] = err.Error()
		}
		if len(rainInfos) > 0 {
			var totalRain float64
			gridHtml := "<tbody>"
			gridHtml += "<tr>" + `<th class="success">站名</th>` + `<th class="success">雨量</th>` + "</tr>"
			for i := 0; i < len(rainInfos); i++ {
				totalRain += rainInfos[i].Rain
				gridHtml += `<tr>` +
					`<td>` + rainInfos[i].Name + `</td>` +
					`<td>` + strconv.FormatFloat(rainInfos[i].Rain, 'f', 2, 64) + `</td>` +
					`</tr>`
			}
			gridHtml += `<tr><td colspan="2" style="text-align:right;">总计雨量:` + strconv.FormatFloat(totalRain, 'f', 2, 64) + `</td></tr>`
			gridHtml += "</tbody>"
			c.Data["json"] = gridHtml
		}
	} else if updateRain == "true" && year != "" {
		updateSuccess := models.UpdateRainInfoByYear(year)
		if updateSuccess {
			c.Data["json"] = "success"
		} else {
			c.Data["json"] = "fail"
		}
	}
	c.ServeJSON()
}
