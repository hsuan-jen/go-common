package dao

import (
	"flag"
	"os"
	"testing"

	"go-common/app/job/main/account-recovery/conf"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "dev" {
		flag.Set("app_id", "main.account.account-recovery-job")
		flag.Set("conf_token", "cf691261b587f3735c1ceec1f27f8177")
		flag.Set("tree_id", "55366")
		flag.Set("conf_version", "docker-1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	} else {
		flag.Set("conf", "../cmd/account-recovery-job.toml")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	m.Run()
	os.Exit(0)
}
