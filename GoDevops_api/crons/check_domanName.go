package crons

import "github.com/astaxie/beego/logs"

type check_domainNameJob func()

var check_domainName *check_domainNameJob

func init() {
	crontab.AddJob("0 */5 * * * ?", check_domainName)
}

func (j *check_domainNameJob) Run() {
	logs.Info("check_domainName")
}
