package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func secureFileServer(root string) http.HandlerFunc {
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir(root))

	return func(w http.ResponseWriter, r *http.Request) {
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

		log.Println(path)
		if path == "/" || !strings.Contains(path, ".") {
			r.URL.Path = "/"
		}

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		fs.ServeHTTP(w, r)
	}
}

func main() {
	port := flag.String("p", "8080", "port number")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()

	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := filepath.EvalSymlinks(absDir); err != nil {
		log.Fatal("指定されたディレクトリにアクセスできません:", err)
	}

	log.Printf("Serving %s on HTTP port: %s\n", absDir, *port)

	http.Handle("/", secureFileServer(*dir))

	// TLS/SSL設定を追加する場合
	// log.Fatal(http.ListenAndServeTLS(":"+*port, "cert.pem", "key.pem", nil))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
