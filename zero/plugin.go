package zero

// 插件化服务
// Todo 完善插件接口,添加更多钩子
type IPlugin interface {
	// 获取插件信息
	GetPluginInfo() PluginInfo
	// 初始化插件
	Init()
	// 登录成功钩子
	OnLogin()
	// 输出日志信息
	Log() <-chan LogEvent
}

// 插件相关信息
type PluginInfo struct {
	PluginName string `json:"plugin_name"`
	Author     string `json:"author"`
	Version    string `json:"version"`
	Details    string `json:"details"`
}
