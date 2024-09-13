package supabaseAPI

import (
	"github.com/Hosi121/SpeakUp/config"
	"github.com/supabase-community/auth-go"
)

var SupabaseClient *auth.Client

func InitSupabase() {
	supabaseRef, supabaseKey := config.GetSupabaseAPI()
	client := auth.New(supabaseRef, supabaseKey)
	SupabaseClient = &client
}
