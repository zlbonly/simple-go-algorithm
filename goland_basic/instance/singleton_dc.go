package goland_basic

import "sync"

type Singleton struct {
}

var (
	singleton *Singleton
	locker    sync.Mutex
)

/**
懒汉加载（Lazy Loading）
1、懒汉模式是指全局的单例实例只会在第一次被使用时才会被构建
2、懒汉模式的缺点在于非线程安全，由于在创建时可能存在线程访问，因此有可能出现多个实例。
*/
func GetInstance1() *Singleton {
	if singleton == nil {
		singleton = &Singleton{} // 非线程安全。
	}
	return singleton
}

/**
1、由于懒汉模式是非线程安全，可通过加锁来解决并发时的线程安全，一般使用互斥锁来解决可能出现的数据不一致问题，但每次加锁也需付出性能代价

缺点：每次请求单例时都会加锁和解锁，锁的目的在于解决对象初始化时可能出现的并发问题，当对象被创建之后，实际上加锁已经失去了意义，此时会拖慢速度

*/
func GetInstance2() *Singleton {
	locker.Lock()
	defer locker.Unlock()
	if singleton == nil {
		singleton = &Singleton{}
	}
	return singleton
}

/**
懒汉双重锁（Check-lock-Check）
为解决重复加锁和解锁的问题，可引入双重锁检查机制（Check-lock-Check），又称为DCL(Double Check Lock)。即第一判断时不加锁，第二次判断时加锁以保证线程安全，一旦对象建立则获取对象时就无需加锁了。
为避免懒汉模式每次都需加锁带来的性能损耗，可采用双重锁来避免每次加锁以提高代码效率，即带检查所的单例模式。
*/
func GetInstance() *Singleton {
	if singleton == nil {
		locker.Lock()
		defer locker.Unlock()

		if singleton == nil {
			singleton = &Singleton{}
		}
	}
	return singleton
}

/*
	双重锁的缺陷在于编译器优化不会检查实例的存储状态，同时每次访问都需要检查两次，为解决这个问题，可采用sync/atomic包中原子性操作来自动加载并设置标记。
	参考singleton_atomic.go
*/
func GetInstanceAtom() {

}

/**
sync.Once提供的Do(f func())方法会使用添加锁来进行原子操作来保证回调函数只执行一次，sync.Once内部本质上也是双重检查的方式。
参考 singleton_once.go
*/
func GetInstanceSyncOnce() {

}
