package e

import "fmt"

//列表数据显示条数
const (
	LimitUserOrders  = 3 //充值记录
	LimitUserAddr    = 3 //我的地址
	LimitUserDevices = 3 //用户设备
	LimitSysMsg      = 3 //系统消息
)

/**
	翻页偏移量
	pageNum 页索引
	limitNums 显示数量
 */
func Page(pageNum int, limitNums int) (int, int) {
	if pageNum <= 1 {
		return limitNums, 0
	} else {
		return limitNums, (pageNum - 1) * limitNums
	}
}

/**
	where翻页
	maxId 最大筛选主键
	pkName 主键名称
 */
func WPage(minPk int, pkName string) string {
	if minPk <= 0 {
		return ""
	}
	return fmt.Sprintf("%s<%d", pkName, minPk)
}
