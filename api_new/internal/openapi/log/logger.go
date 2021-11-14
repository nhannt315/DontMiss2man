package log

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/fields"
	"github.com/nhannt315/real_estate_api/pkg/set"
	pkgstrings "github.com/nhannt315/real_estate_api/pkg/strings"
)

// Option Loggerオプション
type Option func(l *Logger)

// MaskHeaderKeys 値をマスクしたいHTTPヘッダーキーを指定する
func MaskHeaderKeys(keys ...string) Option {
	return func(l *Logger) {
		for _, k := range keys {
			k = strings.ToLower(k)
			l.maskHeaderKeys.add(k)
		}
	}
}

// MaskRequestBodyKeys 値をマスクしたいResponseのFieldを指定する
func MaskRequestBodyKeys(keys ...string) Option {
	return func(l *Logger) {
		l.maskResponseKeys = pkgstrings.GenerateRegexs(keys...)
	}
}

// IgnorePaths ログ出力したくない URL Pathを指定する。今の所完全一致のみ
func IgnorePaths(paths ...string) Option {
	return func(l *Logger) {
		for _, p := range paths {
			l.ignoreHandlerPaths.Add(p)
		}
	}
}

// Logger OpenAPIのリクエスト/レスポンスをログ出力する
type Logger struct {
	l                  logs.Logger
	maskHeaderKeys     *maskHeaderKeys
	ignoreHandlerPaths *set.StringSet
	maskResponseKeys   []*pkgstrings.MaskJSONField
}

// NewLogger OpenAPIのロガーを作成
func NewLogger(logger logs.Logger, opts ...Option) *Logger {
	l := &Logger{
		l:                  logger,
		maskHeaderKeys:     newMaskHeaderKeys(),
		ignoreHandlerPaths: set.NewStringSet(),
	}

	for _, o := range opts {
		o(l)
	}

	return l
}

const logMsg = "openapi"

func (m *Logger) shouldLog(ectx echo.Context) bool {
	return !m.ignoreHandlerPaths.Contains(ectx.Path())
}

// LogResponse レスポンス内容をログ出力
func (m *Logger) LogResponse(ectx echo.Context, respErr error) {
	if !m.shouldLog(ectx) {
		return
	}

	ctx := ectx.Request().Context()
	respFields := m.commonFields(ectx)
	resp := ectx.Response()

	status := http.StatusInternalServerError
	if resp != nil {
		respFields = append(respFields, fields.HTTPHeader(headerString(resp.Header(), m.maskHeaderKeys)))
		respFields = append(respFields, fields.HTTPStatus(resp.Status))
		status = resp.Status
	} else if respErr == nil {
		respErr = errors.New("response is nil")
	}

	switch {
	case respErr == nil:
		m.l.AddFields(respFields...).Info(ctx, logMsg)
	case (status / 100) == 4: // 400系はユーザーエラー
		respFields = append(respFields, fields.Error(respErr))
		m.l.AddFields(respFields...).Info(ctx, logMsg)
	default:
		m.l.AddFields(respFields...).Error(ctx, respErr)
	}
}

// LogRequest リクエスト情報をログ出力
func (m *Logger) LogRequest(ectx echo.Context) error {
	if !m.shouldLog(ectx) {
		return nil
	}

	req := ectx.Request()
	ctx := req.Context()

	reqFields := m.commonFields(ectx)

	reqFields = append(reqFields, fields.HTTPHeader(headerString(req.Header, m.maskHeaderKeys)))

	if shouldLogRequestBody(ectx) {
		// Request Body
		reqBody, err := copyRequestBody(req)
		if err != nil {
			m.l.AddFields(reqFields...).Info(ctx, logMsg)
			return err
		}
		if len(reqBody) > 0 {
			reqFields = append(reqFields, fields.HTTPBody(pkgstrings.MaskJSON(string(reqBody), m.maskResponseKeys)))
		}
	}

	m.l.AddFields(reqFields...).Info(ctx, logMsg)
	return nil
}

func (m *Logger) commonFields(ectx echo.Context) []logs.Field {
	req := ectx.Request()
	var common []logs.Field
	// Method
	common = append(common, fields.HTTPMethod(ectx.Request().Method))
	// Path
	common = append(common, fields.HTTPPath(req.URL.Path))
	common = append(common, fields.HTTPHandlerPath(ectx.Path()))
	// Query
	if q := ectx.QueryString(); q != "" {
		common = append(common, fields.HTTPQuery(ectx.QueryString()))
	}

	return common
}

func headerString(header http.Header, maskKeys *maskHeaderKeys) string {
	var sb strings.Builder
	delim := ""
	for k := range header {
		var v string
		if maskKeys.contains(k) {
			v = pkgstrings.MaskedStringText
		} else {
			v = header.Get(k)
		}
		sb.WriteString(delim)
		sb.WriteString(`'`)
		sb.WriteString(k)
		sb.WriteString(": ")
		sb.WriteString(v)
		sb.WriteString(`'`)
		delim = ", "
	}

	return sb.String()
}

func copyRequestBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}
	// Request Body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read request body")
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

	return reqBody, nil
}

func shouldLogRequestBody(ectx echo.Context) bool {
	req := ectx.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	// content-type: jsonのみ
	return strings.HasPrefix(ctype, echo.MIMEApplicationJSON)
}

type maskHeaderKeys struct {
	s *set.StringSet
}

func newMaskHeaderKeys() *maskHeaderKeys {
	return &maskHeaderKeys{s: set.NewStringSet()}
}

func (h *maskHeaderKeys) add(val string) {
	val = strings.ToLower(val)
	h.s.Add(val)
}

func (h *maskHeaderKeys) contains(val string) bool {
	val = strings.ToLower(val)
	return h.s.Contains(val)
}
