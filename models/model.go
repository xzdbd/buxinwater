package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Userinfo struct {
	Id       int64  `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
	role     int64
}

type RainInfo struct {
	Name string
	Rain float64
}

func init() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("spatialdbconnection"))
	orm.SetMaxIdleConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
}

func GetRainInfoByMonth(year string, month string, stationIds string) (rainInfos []RainInfo, err error) {
	o := orm.NewOrm()
	num, err := o.Raw("select sum(雨量) as rain, 站名 as name from rains where extract(year from 日期) = ? and extract(month from 日期) = ? and 站号 in ("+stationIds+") group by 站名;", year, month).QueryRows(&rainInfos)
	if err != nil {
		beego.Error(err)
	}
	if num == 0 {
		o.Raw("select distinct 0 as rain, 站名 as name from rains where 站号 in (" + stationIds + ") group by 站名;").QueryRows(&rainInfos)
	}
	return
}

func GetRainInfoByYear(year string, stationIds string) (rainInfos []RainInfo, err error) {
	o := orm.NewOrm()
	num, err := o.Raw("select sum(雨量) as rain, 站名 as name from rains where extract(year from 日期) = ? and 站号 in ("+stationIds+") group by 站名;", year).QueryRows(&rainInfos)
	if err != nil {
		beego.Error(err)
	}
	if num == 0 {
		o.Raw("select distinct 0 as rain, 站名 as name from rains where 站号 in (" + stationIds + ") group by 站名;").QueryRows(&rainInfos)
	}
	return
}

func UpdateRainInfoByYear(year string) bool {
	o := orm.NewOrm()
	var rainInfos []RainInfo
	num, err := o.Raw("select sum(雨量) as rain, 站名 as name from rains where extract(year from 日期) = ? group by 站名;", year).QueryRows(&rainInfos)
	if err != nil {
		beego.Error(err)
		return false
	}
	if num <= 0 {
		return false
	}

	for i := 0; i < len(rainInfos); i++ {
		_, ormerr := o.Raw("UPDATE dataloader.cezhan SET jyl=? WHERE 站名=?", rainInfos[i].Rain, rainInfos[i].Name).Exec()
		if ormerr != nil {
			beego.Error("Error when updating cezhan jyl", ormerr)
			return false
		}
	}
	return true
}
