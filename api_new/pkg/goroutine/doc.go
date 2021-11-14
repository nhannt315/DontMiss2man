/*
Package goroutine は、uno向けのgoルーチンを操作する関数を提供する。

unoとしては、goルーチンIDとinterceptorを追加的に実装している。

GoルーチンID

goルーチンIDは、実行/生成されたgoルーチンの識別に使用する。

goルーチンIDをログへ出力することで、1つのgoルーチンに着目してログを確認することができる。
( goルーチンIDがない場合、複数goルーチンのログが混ざってしまうため、ログ確認が難しくなる。 )

goルーチンIDは、Contextへ設定されるため、必要なfuncへはContextを引き継ぐ必要がある。

Interceptor

interceptorは、goroutineの処理前後に追加動作を実施できるようにしている。
*/
package goroutine
