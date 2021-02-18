package boot

import (
	"context"

	"github.com/go-spring/spring-core/bean"
	"github.com/go-spring/spring-core/core"
)

func ApplicationContext() core.ApplicationContext {
	return gApp.ApplicationContext()
}

// GetProfile 返回运行环境
func GetProfile() string {
	return ApplicationContext().GetProfile()
}

func GetProperty(key string) interface{} {
	return ApplicationContext().GetProperty(key)
}

func WireBean(i interface{}) {
	ApplicationContext().WireBean(i)
}

// Beans 获取所有 Bean 的定义，不能保证解析和注入，请谨慎使用该函数!
func Beans() []*core.BeanInstance {
	return ApplicationContext().Beans()
}

func GetBean(i interface{}, selector ...bean.Selector) bool {
	return ApplicationContext().GetBean(i, selector...)
}

func FindBean(selector bean.Selector) (bean.Instance, bool) {
	return ApplicationContext().FindBean(selector)
}

func CollectBeans(i interface{}, selectors ...bean.Selector) bool {
	return ApplicationContext().CollectBeans(i, selectors...)
}

func Invoke(fn interface{}, args ...core.Arg) error {
	return ApplicationContext().Invoke(fn, args...)
}

type GoFuncWithContext func(context.Context)

func Go(fn GoFuncWithContext) {
	appCtx := ApplicationContext()
	appCtx.Go(func() { fn(appCtx.Context()) })
}
