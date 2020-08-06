package core

import (
	"encoding/json"
	"github.com/astaxie/beego"
	beeCache "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Cache beeCache.Cache

func cacheInit() {
	cacheConf := make(map[string]interface{})
	var err error
	if beego.AppConfig.String("cache::cachedriver") == "file" {
		cacheConf["CachePath"] = beego.AppConfig.String("cache::cachepath")
		cacheConf["FileSuffix"] = beego.AppConfig.String("cache::filesuffix")
		cacheConf["DirectoryLevel"] = beego.AppConfig.String("cache::directorylevel")
		cacheConf["EmbedExpiry"] = beego.AppConfig.String("cache::embedexpiry")
		cacheConfJson, _ := json.Marshal(cacheConf)
		Cache, err = beeCache.NewCache("file", string(cacheConfJson))
	} else {
		cacheConf["key"] = beego.AppConfig.String("cache::key")
		cacheConf["conn"] = beego.AppConfig.String("cache::conn")
		cacheConf["dbNum"] = beego.AppConfig.String("cache::dbnum")
		cacheConf["password"] = beego.AppConfig.String("cache::password")
		cacheConfJson, _ := json.Marshal(cacheConf)
		Cache, err = beeCache.NewCache("redis", string(cacheConfJson))
	}
	if err != nil {
		panic(err)
	}
}
