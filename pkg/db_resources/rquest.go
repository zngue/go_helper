package db_resources

import "gorm.io/gorm"

type DBFn func(db *gorm.DB) *gorm.DB

type Common struct {
	Where  map[string]any
	Order  []string
	Select any
	Fn     DBFn
}

// Request 查询多条数据
type Request struct {
	*Common
	*Page
	IsSingle bool
}

type RequestFn func(request *Request) *Request

// RequestWhere  RequestFn 设置Where条件
func RequestWhere(where map[string]any) RequestFn {
	return func(request *Request) *Request {
		request.Common.Where = where
		return request
	}
}

// RequestOrder  RequestFn 设置Order条件
func RequestOrder(order []string) RequestFn {
	return func(request *Request) *Request {
		request.Common.Order = order
		return request
	}
}

// RequestSelect  RequestFn 设置Select条件
func RequestSelect(selects any) RequestFn {
	return func(request *Request) *Request {
		request.Common.Select = selects
		return request
	}
}

// RequestDBFn  DB设置Fn条件
func RequestDBFn(fn DBFn) RequestFn {
	return func(request *Request) *Request {
		request.Common.Fn = fn
		return request
	}
}

// RequestPage  RequestFn 设置分页条件
func RequestPage(page, PageSize int) RequestFn {
	return func(request *Request) *Request {
		pageOp := new(Page)
		pageOp.Page = page
		pageOp.PageSize = PageSize
		request.Page = pageOp
		return request
	}
}

// RequestIsSingle  RequestFn 设置单条数据
func RequestIsSingle(isSingle bool) RequestFn {
	return func(request *Request) *Request {
		request.IsSingle = isSingle
		return request
	}
}

// NewRequest 实例化Request
func NewRequest(fns ...RequestFn) *Request {
	common := new(Common)
	request := new(Request)
	request.Common = common
	if fns == nil {
		return request
	}
	for _, fn := range fns {
		request = fn(request)
	}
	return request
}
