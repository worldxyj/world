package core

func init() {
	// 初始化日志
	logsInit()
	// 初始化缓存
	cacheInit()
	// 初始化数据库
	databaseInit()
	//初始化模板函数
	viewFunctions()
	//自定义错误页面
	errorCustom()
}
