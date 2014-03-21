package singleton

// 構造体の名前を小文字にすることでパッケージ外へのエクスポートを行わない
type singleton struct {
}

// インスタンスを保持する変数も小文字にすることでエクスポートを行わない
var instance *singleton

// インスタンス取得用の関数のみエクスポートしておき、ここでインスタンスが
// 一意であることを保証する
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
