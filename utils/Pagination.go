package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {
	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}

// Pagination is page util
func Pagination(ctx *gin.Context) (pageStr string, num int, err error) {
	limit := ctx.DefaultQuery("page_size", "8")
	pageNumber := ctx.DefaultQuery("page_number", "1")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return "", 0, err
	}
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		return "", 0, err
	}
	if pageNumberInt != 0 {
		pageNumberInt--
	}
	offsetInt := limitInt * pageNumberInt
	pageStr = fmt.Sprintf(" limit %d offset %d", limitInt+1, offsetInt)
	return pageStr, limitInt, nil
}
