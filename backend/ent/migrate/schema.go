// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AithemeSsColumns holds the columns for the "aitheme_ss" table.
	AithemeSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "theme_id", Type: field.TypeInt, Unique: true},
		{Name: "theme_text", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AithemeSsTable holds the schema information for the "aitheme_ss" table.
	AithemeSsTable = &schema.Table{
		Name:       "aitheme_ss",
		Columns:    AithemeSsColumns,
		PrimaryKey: []*schema.Column{AithemeSsColumns[0]},
	}
	// CallSsColumns holds the columns for the "call_ss" table.
	CallSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "call_id", Type: field.TypeInt, Unique: true},
		{Name: "session_id", Type: field.TypeInt},
		{Name: "call_start", Type: field.TypeTime},
		{Name: "call_end", Type: field.TypeTime},
		{Name: "rating", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "matchings_makes", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CallSsTable holds the schema information for the "call_ss" table.
	CallSsTable = &schema.Table{
		Name:       "call_ss",
		Columns:    CallSsColumns,
		PrimaryKey: []*schema.Column{CallSsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "call_ss_matching_ss_makes",
				Columns:    []*schema.Column{CallSsColumns[7]},
				RefColumns: []*schema.Column{MatchingSsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FriendSsColumns holds the columns for the "friend_ss" table.
	FriendSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "friend_id", Type: field.TypeInt, Unique: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "target_user_id", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "FRIEND", "BLOCKED"}},
		{Name: "created_at", Type: field.TypeTime},
	}
	// FriendSsTable holds the schema information for the "friend_ss" table.
	FriendSsTable = &schema.Table{
		Name:       "friend_ss",
		Columns:    FriendSsColumns,
		PrimaryKey: []*schema.Column{FriendSsColumns[0]},
	}
	// MatchingSsColumns holds the columns for the "matching_ss" table.
	MatchingSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "match_id", Type: field.TypeInt, Unique: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "matched_user_id", Type: field.TypeInt},
		{Name: "session_id", Type: field.TypeInt},
		{Name: "matched_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"MATCHED", "PROCCESSING", "FINISHED"}},
		{Name: "sessions_has", Type: field.TypeInt, Nullable: true},
	}
	// MatchingSsTable holds the schema information for the "matching_ss" table.
	MatchingSsTable = &schema.Table{
		Name:       "matching_ss",
		Columns:    MatchingSsColumns,
		PrimaryKey: []*schema.Column{MatchingSsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matching_ss_session_ss_has",
				Columns:    []*schema.Column{MatchingSsColumns[7]},
				RefColumns: []*schema.Column{SessionSsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MemoSsColumns holds the columns for the "memo_ss" table.
	MemoSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "memo_id", Type: field.TypeInt, Unique: true},
		{Name: "user_id", Type: field.TypeInt, Unique: true},
		{Name: "memo1", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "memo2", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "users_prepares", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// MemoSsTable holds the schema information for the "memo_ss" table.
	MemoSsTable = &schema.Table{
		Name:       "memo_ss",
		Columns:    MemoSsColumns,
		PrimaryKey: []*schema.Column{MemoSsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "memo_ss_user_ss_prepares",
				Columns:    []*schema.Column{MemoSsColumns[5]},
				RefColumns: []*schema.Column{UserSsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionSsColumns holds the columns for the "session_ss" table.
	SessionSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "session_id", Type: field.TypeInt, Unique: true},
		{Name: "session_start", Type: field.TypeTime},
		{Name: "session_end", Type: field.TypeTime},
		{Name: "theme_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "sessions_uses", Type: field.TypeInt, Nullable: true},
	}
	// SessionSsTable holds the schema information for the "session_ss" table.
	SessionSsTable = &schema.Table{
		Name:       "session_ss",
		Columns:    SessionSsColumns,
		PrimaryKey: []*schema.Column{SessionSsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "session_ss_aitheme_ss_uses",
				Columns:    []*schema.Column{SessionSsColumns[6]},
				RefColumns: []*schema.Column{AithemeSsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserSsColumns holds the columns for the "user_ss" table.
	UserSsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt, Unique: true},
		{Name: "username", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SUPERUSER", "ADMIN", "USER"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "access_token", Type: field.TypeString, Size: 2147483647, Default: ""},
	}
	// UserSsTable holds the schema information for the "user_ss" table.
	UserSsTable = &schema.Table{
		Name:       "user_ss",
		Columns:    UserSsColumns,
		PrimaryKey: []*schema.Column{UserSsColumns[0]},
	}
	// UsersConnectsColumns holds the columns for the "users_connects" table.
	UsersConnectsColumns = []*schema.Column{
		{Name: "users_id", Type: field.TypeInt},
		{Name: "friends_id", Type: field.TypeInt},
	}
	// UsersConnectsTable holds the schema information for the "users_connects" table.
	UsersConnectsTable = &schema.Table{
		Name:       "users_connects",
		Columns:    UsersConnectsColumns,
		PrimaryKey: []*schema.Column{UsersConnectsColumns[0], UsersConnectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_connects_users_id",
				Columns:    []*schema.Column{UsersConnectsColumns[0]},
				RefColumns: []*schema.Column{UserSsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_connects_friends_id",
				Columns:    []*schema.Column{UsersConnectsColumns[1]},
				RefColumns: []*schema.Column{FriendSsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersParticipatesColumns holds the columns for the "users_participates" table.
	UsersParticipatesColumns = []*schema.Column{
		{Name: "users_id", Type: field.TypeInt},
		{Name: "matchings_id", Type: field.TypeInt},
	}
	// UsersParticipatesTable holds the schema information for the "users_participates" table.
	UsersParticipatesTable = &schema.Table{
		Name:       "users_participates",
		Columns:    UsersParticipatesColumns,
		PrimaryKey: []*schema.Column{UsersParticipatesColumns[0], UsersParticipatesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_participates_users_id",
				Columns:    []*schema.Column{UsersParticipatesColumns[0]},
				RefColumns: []*schema.Column{UserSsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_participates_matchings_id",
				Columns:    []*schema.Column{UsersParticipatesColumns[1]},
				RefColumns: []*schema.Column{MatchingSsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AithemeSsTable,
		CallSsTable,
		FriendSsTable,
		MatchingSsTable,
		MemoSsTable,
		SessionSsTable,
		UserSsTable,
		UsersConnectsTable,
		UsersParticipatesTable,
	}
)

func init() {
	CallSsTable.ForeignKeys[0].RefTable = MatchingSsTable
	MatchingSsTable.ForeignKeys[0].RefTable = SessionSsTable
	MemoSsTable.ForeignKeys[0].RefTable = UserSsTable
	SessionSsTable.ForeignKeys[0].RefTable = AithemeSsTable
	UsersConnectsTable.ForeignKeys[0].RefTable = UserSsTable
	UsersConnectsTable.ForeignKeys[1].RefTable = FriendSsTable
	UsersParticipatesTable.ForeignKeys[0].RefTable = UserSsTable
	UsersParticipatesTable.ForeignKeys[1].RefTable = MatchingSsTable
}
