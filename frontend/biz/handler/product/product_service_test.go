package product

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetProduct(t *testing.T) {
	// 设置测试服务器
	h := server.Default()
	h.GET("/product", GetProduct)

	// 发起请求，获取产品ID为123的信息
	path := "/product?id=1"
	w := ut.PerformRequest(h.Engine, "GET", path,
		&ut.Body{Body: bytes.NewBufferString(""), Len: 0},
		ut.Header{})

	// 获取响应
	resp := w.Result()

	// 打印完整的响应内容，这样你可以看到所有返回的数据
	t.Logf("Response Status: %d", resp.StatusCode())
	t.Logf("Response Body: %s", string(resp.Body()))
}

func TestSearchProducts(t *testing.T) {
	h := server.Default()
	h.GET("/search", SearchProducts)
	path := "/search"                                         // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
