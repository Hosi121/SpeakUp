package controllers

import (
	"context"
	"log/slog"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/friends"
)

func SendMessage(senderID, targetID, message string) error {
	// DBの用意
	dsn := config.GetDSN()
	db_client, err := ent.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open connection to database: %v", err)
		return err
	}
	defer db_client.Close()

	ctx := context.Background()
	// DBからfriend_idを取得
	friend, err := db_client.FRIENDS.
		Query().
		Where(
			friends.UserID(senderID),
			friends.TargetUserID(targetID),
		)
	if err != nil {
		slog.Error("Not found this email: %v", err)
		return err
	}

	// CHATSにメッセージを登録
	chat, err := db_client.CHATS.
		Create().
		SetFriendID(friend.ID).
		SetMessage(message).
		Save(ctx)
	if err != nil {
		slog.Error("Failed creating chat: %v", err)
	}
	slog.Info("Created chat: %v", chat)
}
