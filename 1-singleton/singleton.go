package singleton

import "sync"

// Singleton -------------------------------------------------------------
// 单例模式采用了 饿汉式 和 懒汉式 两种实现，个人其实更倾向于饿汉式的实现，简单，并且可以将问题及早暴露，
// 懒汉式虽然支持延迟加载，但是这只是把冷启动时间放到了第一次使用的时候，并没有本质上解决问题，
// 并且为了实现懒汉式还不可避免的需要加锁

// Singleton 饿汉式单例模式
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}

// -------------------------------------------------------------
// 懒汉式单例模式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 双重检测
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
