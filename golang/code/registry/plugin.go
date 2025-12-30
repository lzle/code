package registry

import (
	"context"
	"errors"
	"sync"
)

type PluginMgr struct {
	plugins map[string]Registry
	mutex   sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// 插件注册
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	_, ok := p.plugins[plugin.Name()]
	if ok {
		return errors.New("registry plugin exist")
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return  pluginMgr.initRegistry(ctx, name, opts...)
}



func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	plugin,ok := p.plugins[name]
	if !ok {
		err = errors.New("plugin %s not exist")
		return
	}
	registry = plugin
	err = registry.Init(ctx, opts...)
	return
}



























