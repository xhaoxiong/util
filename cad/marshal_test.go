/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  marshal_test.go
 * @Version: 1.0.0
 * @Date: 4/8/20 7:08 下午
 */
package cad

import (
	"os"
	"testing"
	"time"
)

type Profile struct {
	DbType     string
	IsError    string
	Status     string
	Version    string
	CurTime    string
	ServerIp   string
	ServerPort string
	HostName   string
}

func TestMarshal(t *testing.T) {
	profile := Profile{}
	pcName, _ := os.Hostname()
	profile.DbType = "sqlite"
	profile.IsError = "ok"
	profile.Status = "0"
	profile.Version = "V1.0.0"
	profile.CurTime = time.Now().Local().Format("2006-01-02 15:04:05")
	profile.ServerIp = "127.0.0.1"
	profile.ServerPort = ""
	profile.HostName = pcName
	marshal, err := Marshal(profile)
	if err != nil {
		panic(err)
	}

	t.Log(marshal)
}
