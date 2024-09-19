package controllers

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/users"
	supabaseAPI "github.com/Hosi121/SpeakUp/supaseAPI"
	"github.com/Hosi121/SpeakUp/utils"
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
func SignIn(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Supabase authentication
		supabaseAPI.InitSupabase()
		supabaseClient := supabaseAPI.SupabaseClient

		signinReq := types.TokenRequest{
			GrantType: "password",
			Email:     request.Email,
			Password:  request.Password,
		}
		_, err := supabaseClient.Token(signinReq)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Use request context
		ctx := c.Request.Context()

		// Fetch user from database
		user, err := client.USERS.
			Query().
			Where(users.EmailEQ(request.Email)).
			First(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Update updated_at field
		err = client.USERS.
			UpdateOneID(user.ID).
			SetUpdatedAt(time.Now()).
			Exec(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Generate JWT token
		jwtToken, err := utils.GenerateJWT(strconv.Itoa(user.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{
			"message":     "User signed in successfully",
			"accessToken": jwtToken,
		})
	}
}
