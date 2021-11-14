/*
Package appinfo は、アプリケーション(プログラム/binary)の情報を提供する。

ここで提供されるバージョンなどの静的な情報は、 go build時に設定される想定。
 ( -ldflags "-X github.com/moneyforward/fullback/pkg/appinfor.xxx=yyy' " など)
*/
package appinfo
