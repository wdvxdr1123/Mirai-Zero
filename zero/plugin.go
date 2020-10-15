package zero

// 插件化服务
// Todo 完善插件接口,添加更多钩子
type IPlugin interface {
	// 获取插件信息
	GetPluginInfo() PluginInfo
	// 初始化插件
	Init(*Accessory)
	// 登录成功钩子
	OnLogin()
	// 启用插件
	Enable()
	// 停用插件
	Stop()
}

// 插件相关信息
type PluginInfo struct {
	PluginName string           `json:"plugin_name"`
	Author     string           `json:"author"`
	Version    string           `json:"version"`
	Details    string           `json:"details"`
	Permission PluginPermission `json:"permission"` // 先写着挖个坑，不一定会填
}

type PluginPermission int

const (
	Normal PluginPermission = iota
	Manage
	Root
)
