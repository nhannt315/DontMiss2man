/*
Package errors は、 error を操作する関数を提供する。

goのデフォルト error とともに、goのエラーで提供しないstack tracingのエラーも表示したいので、両方を出力できるようにしている。

"errors" + "github.com/pkg/errors"(スタックトレースのため) のwrapper。

stack情報をerrorに保存するためには、Wrap, Wrapfで errorにstack traceを含めるようおにwrapして、

出力する側で %+v 形で表示する。
*/
package errors
