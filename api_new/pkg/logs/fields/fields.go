package fields

import (
	"context"
	"fmt"
	"time"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

// Error return key:"error" field
func Error(val error) logs.Field {
	return logs.NewStringField("error", val.Error())
}

// Stack return key:"stack" field
func Stack(val error) logs.Field {
	return logs.NewStringField("stack", fmt.Sprintf("%+v", val))
}

// GoroutineID return key:"goroutine_id" field if it existed
func GoroutineID(ctx context.Context) logs.Field {
	if id := goroutine.IDValue(ctx); id != 0 {
		return logs.NewUint64Field("goroutine_id", id)
	}
	return nil
}

// Duration return duration fields
func Duration(key string, val time.Duration) logs.Field {
	return logs.NewDurationField(key, val)
}

// HTTPMethod Http Method
func HTTPMethod(val string) logs.Field {
	return logs.NewStringField("http_method", val)
}

// HTTPPath HTTP URL Path
func HTTPPath(val string) logs.Field {
	return logs.NewStringField("http_path", val)
}

// HTTPHandlerPath ルーティング情報にマッチしたHTTPハンドラのパス.
// ハンドラを特定するのに使える
func HTTPHandlerPath(val string) logs.Field {
	return logs.NewStringField("http_handler_path", val)
}

// HTTPHeader HTTP Request/ResponseのHeader
func HTTPHeader(val string) logs.Field {
	return logs.NewStringField("http_header", val)
}

// HTTPQuery HTTP URLのQueryパラメータ
func HTTPQuery(val string) logs.Field {
	return logs.NewStringField("http_query", val)
}

// HTTPBody HTTP Request/ResponseのBody
func HTTPBody(val string) logs.Field {
	return logs.NewStringField("http_body", val)
}

// HTTPStatus HTTP Responseのステータスコード
func HTTPStatus(val int) logs.Field {
	return logs.NewInt64Field("http_status", int64(val))
}

// MethodName return key:"method" field
func MethodName(val string) logs.Field {
	return logs.NewStringField("method", val)
}

// MethodParams return key:"method_params" field
func MethodParams(val interface{}) logs.Field {
	return logs.NewStringField("method_params", fmt.Sprintf("%+v", val))
}

// Query returns key:"query" field
func Query(val string) logs.Field {
	return logs.NewStringField("query", val)
}

// Schema returns key:"schema" field
func Schema(val string) logs.Field {
	return logs.NewStringField("schema", val)
}

// QueryParams return key:"query_param" field
func QueryParams(val interface{}) logs.Field {
	return logs.NewStringField("query_param", fmt.Sprintf("%v", val))
}

// ExecutingDurationQuery return key:"exec_dur" field through Duration function
func ExecutingDurationQuery(val time.Duration) logs.Field {
	return Duration("exec_dur", val)
}

// ExecutingDuration return key:"exec_dur" field through Duration function
func ExecutingDuration(val time.Duration) logs.Field {
	return Duration("exec_dur", val)
}

// NumberOfRows return key:"num_rows" field
func NumberOfRows(val uint64) logs.Field {
	return logs.NewUint64Field("num_rows", val)
}
