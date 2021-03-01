package goland_basic
/**
	饿汉加载
	1、饿汉模式 是指全局的单例实例在类加载时才会被构建
	2、饿汉模式 的缺点在于单例实例初始化时可能会比较耗时，因此加载时间会延长
 */
type SingleTon struct {

}

var singleton  *Singleton = &Singleton{}

func GetInstance *Singleton{
	return singleton
}
