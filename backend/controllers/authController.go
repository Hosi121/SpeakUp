package controllers

import (
	"fmt"

	supabaseAPI "github.com/Hosi121/SpeakUp/supaseAPI"
	"github.com/supabase-community/auth-go/types"
)

func SignUp(email, password string) {
	supabaseAPI.InitSupabase()
	client := supabaseAPI.SupabaseClient
	signupReq := types.SignupRequest{
		Email:    email,
		Password: password,
	}
	resp, err := client.Signup(signupReq)
	if err != nil {
		fmt.Println("Error signed up: %v", err.Error())
	}
	fmt.Println("User signed up successfully: %+v\n", resp)
}

func SignIn(email, password string) {
	supabaseAPI.InitSupabase()
	client := supabaseAPI.SupabaseClient
	signinReq := types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	}
	resp, err := client.Token(signinReq)
	if err != nil {
		fmt.Println("Error signed in: %v", err.Error())
	}
	fmt.Println("User signed in successfully: %+v\n", resp.AccessToken)
}
