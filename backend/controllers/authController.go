package controllers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/users"
	jwt_auth "github.com/Hosi121/SpeakUp/jwt_auth"
	supabaseAPI "github.com/Hosi121/SpeakUp/supaseAPI"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/auth-go/types"
)

// SignUp ハンドラ関数
func SignUp(c *gin.Context) {
	// リクエストボディからemailとpasswordを取得
	var request struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Supabaseクライアントの初期化
	supabaseAPI.InitSupabase()
	client := supabaseAPI.SupabaseClient

	// サインアップリクエストの作成
	signupReq := types.SignupRequest{
		Email:    request.Email,
		Password: request.Password,
	}
	resp, err := client.Signup(signupReq)
	if err != nil {
		// エラーハンドリングとレスポンス
		slog.Error("Error signing up", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 成功レスポンス
	slog.Info("User signed up successfully", slog.String("email", request.Email))
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully", "user": resp})

	// DBの用意
	dsn := config.GetDSN()
	db_client, err := ent.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open connection to database: %v", err)
	}
	defer db_client.Close()

	ctx := context.Background()
	// DBのUSERテーブルに登録
	user, err := db_client.USERS.
		Create().
		SetUsername(request.UserName).
		SetEmail(request.Email).
		SetAvatarURL("https://example.com/avatar.png").
		SetRole("USER").
		Save(ctx)
	if err != nil {
		slog.Error("Failed creating user: %v", err)
	}
	slog.Info("Created user: %v", user)
}

// SignIn ハンドラ関数
func SignIn(c *gin.Context) {
	// リクエストボディからemailとpasswordを取得
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Supabaseクライアントの初期化
	supabaseAPI.InitSupabase()
	client := supabaseAPI.SupabaseClient

	// サインインリクエストの作成
	signinReq := types.TokenRequest{
		GrantType: "password",
		Email:     request.Email,
		Password:  request.Password,
	}
	resp, err := client.Token(signinReq)
	if err != nil {
		// エラーハンドリングとレスポンス
		slog.Error("Error signing in", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// DBの用意
	dsn := config.GetDSN()
	db_client, err := ent.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open connection to database: %v", err)
		return
	}
	defer db_client.Close()

	ctx := context.Background()
	// emailが一致するユーザを取得
	user, err := db_client.USERS.
		Query().
		Where(
			users.EmailEQ(request.Email), // Emailが一致する行を検索
		).
		Order(ent.Desc(users.FieldCreatedAt)). // created_atで結果を降順ソート
		First(ctx)                             // 最初の1件を取得
	if err != nil {
		slog.Error("Not found this email: %v", err)
		// fmt.Println("Not found this email: %v", err)
		return
	}

	// アクセストークンを登録
	err = db_client.USERS.
		UpdateOneID(user.ID).
		SetAccessToken(resp.AccessToken).
		Exec(ctx)
	if err != nil {
		slog.Error("Failed to update access token: %v", err)
		// fmt.Println("Failed to update access token: %v", err)
		return
	}

	// updated_atを更新
	err = db_client.USERS.
		UpdateOneID(user.ID).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
	if err != nil {
		slog.Error("Failed to update created_at: %v", err)
		// fmt.Println("Failed to update created_at: %v", err)
		return
	}

	// JWTトークンを生成
	jwt_token, err := jwt_auth.GenerateJWT(user.ID)
	if err != nil {
		slog.Error("Failed to generate JWT: %v", err)
	}

	// 成功レスポンス
	slog.Info("User signed in successfully", slog.String("email", request.Email))
	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully", "accessToken": jwt_token})

}
