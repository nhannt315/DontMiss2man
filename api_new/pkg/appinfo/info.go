package appinfo

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	defaultVersion  = "v0.0.0"
	defaultCommitID = "unknown"
)

// version バージョン(gitのtagとか)
var version = defaultVersion

// commitID ソースのコミットID(gitの commit SHA とか)
var commitID = defaultCommitID

var packageRootPath = ""

// pkgGet パッケージ取得用の定義
type pkgGet struct {
}

func init() {

	pkg := reflect.TypeOf(pkgGet{}).PkgPath()
	pkgWords := strings.Split(pkg, "/")
	if packageRootPath == "" {
		switch len(pkgWords) {
		case 1:
			packageRootPath = pkgWords[0]
		case 2:
			packageRootPath = strings.Join(pkgWords[0:2], "/")
		default:
			packageRootPath = strings.Join(pkgWords[0:3], "/")
		}
	}
}

// SimpleVersion アプリケーション識別するバージョン
//	バージョンが設定されていれば v1.0.0 などのバージョン
//	バージョンが設定されていなければ コミットID
func SimpleVersion() string {
	if version != defaultVersion {
		return version
	}
	return commitID // バージョンないのでコミットID
}

// FullVersion アプリケーション識別するバージョン(+コミットID)
func FullVersion() string {
	return fmt.Sprintf("%s:%s",
		version, commitID)
}

// PackageRootPath アプリケーションのパッケージrootパス(github.com/moneyforward/XXXX)
func PackageRootPath() string {
	return packageRootPath
}
