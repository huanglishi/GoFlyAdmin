package controller

/**
* app路由引入口《引入模块控制器》
*
* 请把您使用包用 _ "gofly/app/home/XX"导入您编写的包 自动生成路由
* 不需要使用的模块则注释掉 例如home模块暂时用不到就注释掉，这样不占用资源，使用是取消注释即可。
* 路由规则：包路径“home/article” + 包中结构体“Cate”转小写+方法名(首字母转小写_ "gofly/app/business/datacenter"
 */
import (
	_ "gofly/app/admin"
	_ "gofly/app/business"
	_ "gofly/app/common"
	_ "gofly/app/wxapp"
	_ "gofly/app/wxoffi"
)
