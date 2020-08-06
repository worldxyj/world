package core

import (
	"fmt"
	"strconv"
	"strings"
)

//把url路径中的path请求，变成key val，path参数值形式：/user/list/p/1，转成map[string]string
func Path2Map(path string) map[string]string {
	var data = make(map[string]string)
	slice := strings.Split(strings.Trim(path, "/"), "/")
	cnt := len(slice)
	if cnt%2 == 1 {
		cnt = cnt - 1
	}
	if cnt > 0 {
		for i := 0; i < cnt; {
			data[slice[i]] = slice[(i + 1)]
			i = i + 2
		}
	}
	return data
}

//分页函数
//rollPage:展示分页的个数
//totalRows：总记录
//currentPage:每页显示记录数
//urlPrefix:url链接前缀
//urlParams:url键值对参数
func Paginations(rollPage, totalRows, listRows, currentPage int, urlPrefix string, urlParams ...interface{}) string {
	var (
		htmlPage, path string
		pages          []int
		params         []string
	)
	htmlPage = "<ul class=\"pagination\" role=\"navigation\"><li><span>" + strconv.Itoa(totalRows) + " 条记录</span></li>"
	if listRows <= 0 {
		listRows = 10
	}
	//总页数
	totalPage := totalRows / listRows
	if totalRows%listRows > 0 {
		totalPage += 1
	}
	//只有1页的时候，不分页
	if totalPage < 2 {
		return ""
	}
	params_len := len(urlParams)
	if params_len > 0 {
		if params_len%2 > 0 {
			params_len = params_len - 1
		}
		for i := 0; i < params_len; {
			key := strings.TrimSpace(fmt.Sprintf("%v", urlParams[i]))
			val := strings.TrimSpace(fmt.Sprintf("%v", urlParams[i+1]))
			//键存在，同时值不为0也不为空
			if len(key) > 0 && len(val) > 0 && val != "0" {
				params = append(params, key, val)
			}
			i = i + 2
		}
	}

	path = strings.Trim(urlPrefix, "/")
	if len(params) > 0 {
		path = path + "/" + strings.Trim(strings.Join(params, "/"), "/")
	}
	//最后再处理一次“/”，是为了防止urlPrifix参数为空时，出现多余的“/”
	path = "/" + strings.Trim(path, "/") + "/p/"

	if currentPage > totalPage {
		currentPage = totalPage
	}
	if currentPage < 1 {
		currentPage = 1
	}
	index := 0
	rp := rollPage * 2
	for i := rp; i > 0; i-- {
		p := currentPage + rollPage - i
		if p > 0 && p <= totalPage {

			pages = append(pages, p)
		}
	}
	for k, v := range pages {
		if v == currentPage {
			index = k
		}
	}
	pages_len := len(pages)
	if currentPage > 4 {
		htmlPage += `<li><a class="num" href="` + path + `1">1</a></li><li class="disabled" aria-disabled="true"><span>...</span></li>`
	}
	if pages_len <= rollPage {
		for _, v := range pages {
			if v == currentPage {
				htmlPage += fmt.Sprintf(`<li class="active"><a href="javascript:void(0);">%d</a></li>`, v)
			} else {
				htmlPage += fmt.Sprintf(`<li><a class="num" href="`+path+`%d">%d</a></li>`, v, v)
			}
		}

	} else {
		index_min := index - rollPage/2
		index_max := index + rollPage/2
		page_slice := make([]int, 0)
		if index_min > 0 && index_max < pages_len { //切片索引未越界
			page_slice = pages[index_min:index_max]
		} else {
			if index_min < 0 {
				page_slice = pages[0:rollPage]
			} else if index_max > pages_len {
				page_slice = pages[(pages_len - rollPage):pages_len]
			} else {
				page_slice = pages[index_min:index_max]
			}

		}

		for _, v := range page_slice {
			if v == currentPage {
				htmlPage += fmt.Sprintf(`<li class="active"><a href="javascript:void(0);">%d</a></li>`, v)
			} else {
				htmlPage += fmt.Sprintf(`<li><a class="num" href="`+path+`%d">%d</a></li>`, v, v)
			}
		}

	}
	if currentPage < totalPage-2 {
		htmlPage += fmt.Sprintf(`<li class="disabled" aria-disabled="true"><span>...</span></li><li><a class="num" href="`+path+`%d">%d</a></li>`, totalPage, totalPage)
	}
	htmlPage += "</ul>"
	return htmlPage
}
