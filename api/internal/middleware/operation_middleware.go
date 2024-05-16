package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}

type OperationMiddleware struct {
	CoreRpc       coreclient.Core
	ExcludedPaths []string
}

func NewOperationMiddleware(CoreRpc coreclient.Core) *OperationMiddleware {
	return &OperationMiddleware{
		CoreRpc:       CoreRpc,
		ExcludedPaths: []string{}, // 排除路径
	}
}

func (m *OperationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, path := range m.ExcludedPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				// 如果请求路径匹配排除路径，则跳过记录日志
				next.ServeHTTP(w, r)
				return
			}
		}
		//请求时间
		reqTime := time.Now()
		// 	// User's UUID | 用户的UUID
		// Uuid *string `protobuf:"bytes,4,opt,name=uuid,proto3,oneof" json:"uuid,omitempty"`
		uuid := r.Context().Value("userId").(string)
		// // HTTP request method|HTTP请求方法
		// Method *string `protobuf:"bytes,5,opt,name=method,proto3,oneof" json:"method,omitempty"`
		method := r.Method
		// // HTTP request path|HTTP请求路径
		// Path *string `protobuf:"bytes,6,opt,name=path,proto3,oneof" json:"path,omitempty"`
		path := r.URL.Path
		// // HTTP request headers|HTTP请求头部
		// Headers *string `protobuf:"bytes,7,opt,name=headers,proto3,oneof" json:"headers,omitempty"`
		headers, err := json.Marshal(r.Header)
		if err != nil {
			headers = []byte("")
		}
		//超过 20148 则 截取请求头部
		if len(headers) > 2048 {
			headers = headers[:2048]
		}
		// // HTTP request body|HTTP请求体
		// Body *string `protobuf:"bytes,8,opt,name=body,proto3,oneof" json:"body,omitempty"`
		// 读取请求主体
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logx.Errorw("read request body error", logx.Field("detail", err.Error()))
			body = []byte("")
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// 创建一个自定义的 ResponseWriter，用于记录响应
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}
		next(recorder, r)
		// 获取响应时间
		resTime := time.Now()

		// 计算请求耗时
		costTime := resTime.Sub(reqTime).Milliseconds()

		// // // HTTP response status code|HTTP响应状态码
		// // StatusCode *int64 `protobuf:"varint,9,opt,name=status_code,json=statusCode,proto3,oneof" json:"status_code,omitempty"`
		statusCode := recorder.statusCode
		// // // HTTP response headers|HTTP响应头部
		// // ResHeaders *string `protobuf:"bytes,10,opt,name=res_headers,json=resHeaders,proto3,oneof" json:"res_headers,omitempty"`
		// resHeaders, err := json.Marshal(recorder.Header())
		// if err != nil {
		// 	resHeaders = []byte("")
		// }
		// // // HTTP response body|HTTP响应体
		// // ResBody *string `protobuf:"bytes,11,opt,name=res_body,json=resBody,proto3,oneof" json:"res_body,omitempty"`
		// resBody := recorder.body

		// 记录操作日志
		_, err = m.CoreRpc.CreateLogOperation(r.Context(), &core.LogOperationInfo{
			Uuid:       pointy.GetPointer(uuid),
			Method:     pointy.GetPointer(method),
			Path:       pointy.GetPointer(path),
			Headers:    pointy.GetPointer(string(headers)),
			Body:       pointy.GetPointer(string(body)),
			StatusCode: pointy.GetPointer(int64(statusCode)),
			ReqTime:    pointy.GetPointer(reqTime.UnixMilli()),
			ResTime:    pointy.GetPointer(resTime.UnixMilli()),
			CostTime:   pointy.GetPointer(uint64(costTime)),
		})
		if err != nil {
			logx.Errorw("create log operation error", logx.Field("detail", err.Error()))
		}
	}
}
