package zero

// 插件化服务
// Todo 完善插件接口,添加更多钩子
type IPlugin interface {
	GetPluginInfo() PluginInfo
	Init()
	Start()
}

// 插件相关信息
type PluginInfo struct {
	PluginName string `json:"plugin_name"`
	Author     string `json:"author"`
	Version    string `json:"version"`
	Details    string `json:"details"`
}
