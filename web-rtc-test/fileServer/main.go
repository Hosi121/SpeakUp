package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// セキュアなファイルサーバーハンドラ
func secureFileServer(root string) http.HandlerFunc {
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir(root))

	return func(w http.ResponseWriter, r *http.Request) {
		// パスのクリーニング
		path := filepath.Clean(r.URL.Path)

		// 隠しファイル/ディレクトリへのアクセスを防ぐ
		if strings.Contains(path, "/.") {
			http.Error(w, "不正なパス", http.StatusForbidden)
			return
		}

		// パスが root ディレクトリの外を参照していないか確認
		if requestedPath, err := filepath.Abs(filepath.Join(root, path)); err != nil || !strings.HasPrefix(requestedPath, absoluteRoot) {
			http.Error(w, "不正なパス", http.StatusForbidden)
			return
		}

		// Content-Type スニッフィングを制限
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// クリックジャッキング対策
		w.Header().Set("X-Frame-Options", "DENY")

		// XSS保護
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		fs.ServeHTTP(w, r)
	}
}

func main() {
	port := flag.String("p", "8080", "port number")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()

	// 絶対パスの取得
	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	// サーバー起動前の権限チェック
	if _, err := filepath.EvalSymlinks(absDir); err != nil {
		log.Fatal("指定されたディレクトリにアクセスできません:", err)
	}

	log.Printf("Serving %s on HTTP port: %s\n", absDir, *port)

	// セキュアなファイルサーバーを設定
	http.Handle("/", secureFileServer(*dir))

	// TLS/SSL設定を追加する場合
	// log.Fatal(http.ListenAndServeTLS(":"+*port, "cert.pem", "key.pem", nil))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
