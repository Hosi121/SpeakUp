package controllers

import (
	"log/slog"
	"net/http"

	supabaseAPI "github.com/Hosi121/SpeakUp/supaseAPI"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/auth-go/types"
)

// SignUp ハンドラ関数
func SignUp(c *gin.Context) {
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

	// 成功レスポンス
	slog.Info("User signed in successfully", slog.String("email", request.Email))
	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully", "accessToken": resp.AccessToken})
}
