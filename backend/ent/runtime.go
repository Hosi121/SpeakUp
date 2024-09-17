// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Hosi121/SpeakUp/ent/achievements"
	"github.com/Hosi121/SpeakUp/ent/ai_themes"
	"github.com/Hosi121/SpeakUp/ent/calls"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/Hosi121/SpeakUp/ent/friends"
	"github.com/Hosi121/SpeakUp/ent/memos"
	"github.com/Hosi121/SpeakUp/ent/progress"
	"github.com/Hosi121/SpeakUp/ent/schema"
	"github.com/Hosi121/SpeakUp/ent/sessions"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	achievementsFields := schema.ACHIEVEMENTS{}.Fields()
	_ = achievementsFields
	// achievementsDescCreatedAt is the schema descriptor for created_at field.
	achievementsDescCreatedAt := achievementsFields[2].Descriptor()
	// achievements.DefaultCreatedAt holds the default value on creation for the created_at field.
	achievements.DefaultCreatedAt = achievementsDescCreatedAt.Default.(func() time.Time)
	ai_themesFields := schema.AI_THEMES{}.Fields()
	_ = ai_themesFields
	// ai_themesDescCreatedAt is the schema descriptor for created_at field.
	ai_themesDescCreatedAt := ai_themesFields[1].Descriptor()
	// ai_themes.DefaultCreatedAt holds the default value on creation for the created_at field.
	ai_themes.DefaultCreatedAt = ai_themesDescCreatedAt.Default.(time.Time)
	callsFields := schema.CALLS{}.Fields()
	_ = callsFields
	// callsDescCallStart is the schema descriptor for call_start field.
	callsDescCallStart := callsFields[1].Descriptor()
	// calls.DefaultCallStart holds the default value on creation for the call_start field.
	calls.DefaultCallStart = callsDescCallStart.Default.(func() time.Time)
	// callsDescCallEnd is the schema descriptor for call_end field.
	callsDescCallEnd := callsFields[2].Descriptor()
	// calls.DefaultCallEnd holds the default value on creation for the call_end field.
	calls.DefaultCallEnd = callsDescCallEnd.Default.(func() time.Time)
	// callsDescCreatedAt is the schema descriptor for created_at field.
	callsDescCreatedAt := callsFields[4].Descriptor()
	// calls.DefaultCreatedAt holds the default value on creation for the created_at field.
	calls.DefaultCreatedAt = callsDescCreatedAt.Default.(func() time.Time)
	eventsFields := schema.EVENTS{}.Fields()
	_ = eventsFields
	// eventsDescCreatedAt is the schema descriptor for created_at field.
	eventsDescCreatedAt := eventsFields[3].Descriptor()
	// events.DefaultCreatedAt holds the default value on creation for the created_at field.
	events.DefaultCreatedAt = eventsDescCreatedAt.Default.(time.Time)
	event_recordsFields := schema.EVENT_RECORDS{}.Fields()
	_ = event_recordsFields
	// event_recordsDescRecords is the schema descriptor for records field.
	event_recordsDescRecords := event_recordsFields[2].Descriptor()
	// event_records.DefaultRecords holds the default value on creation for the records field.
	event_records.DefaultRecords = event_recordsDescRecords.Default.(string)
	friendsFields := schema.FRIENDS{}.Fields()
	_ = friendsFields
	// friendsDescCreatedAt is the schema descriptor for created_at field.
	friendsDescCreatedAt := friendsFields[3].Descriptor()
	// friends.DefaultCreatedAt holds the default value on creation for the created_at field.
	friends.DefaultCreatedAt = friendsDescCreatedAt.Default.(func() time.Time)
	memosFields := schema.MEMOS{}.Fields()
	_ = memosFields
	// memosDescMemo1 is the schema descriptor for memo1 field.
	memosDescMemo1 := memosFields[1].Descriptor()
	// memos.DefaultMemo1 holds the default value on creation for the memo1 field.
	memos.DefaultMemo1 = memosDescMemo1.Default.(string)
	// memos.Memo1Validator is a validator for the "memo1" field. It is called by the builders before save.
	memos.Memo1Validator = memosDescMemo1.Validators[0].(func(string) error)
	// memosDescMemo2 is the schema descriptor for memo2 field.
	memosDescMemo2 := memosFields[2].Descriptor()
	// memos.DefaultMemo2 holds the default value on creation for the memo2 field.
	memos.DefaultMemo2 = memosDescMemo2.Default.(string)
	// memos.Memo2Validator is a validator for the "memo2" field. It is called by the builders before save.
	memos.Memo2Validator = memosDescMemo2.Validators[0].(func(string) error)
	progressFields := schema.PROGRESS{}.Fields()
	_ = progressFields
	// progressDescLoginDays is the schema descriptor for login_days field.
	progressDescLoginDays := progressFields[1].Descriptor()
	// progress.DefaultLoginDays holds the default value on creation for the login_days field.
	progress.DefaultLoginDays = progressDescLoginDays.Default.(int)
	// progressDescConsecutiveParticipants is the schema descriptor for consecutive_participants field.
	progressDescConsecutiveParticipants := progressFields[2].Descriptor()
	// progress.DefaultConsecutiveParticipants holds the default value on creation for the consecutive_participants field.
	progress.DefaultConsecutiveParticipants = progressDescConsecutiveParticipants.Default.(int)
	// progressDescConsecutiveLoginDays is the schema descriptor for consecutive_login_days field.
	progressDescConsecutiveLoginDays := progressFields[3].Descriptor()
	// progress.DefaultConsecutiveLoginDays holds the default value on creation for the consecutive_login_days field.
	progress.DefaultConsecutiveLoginDays = progressDescConsecutiveLoginDays.Default.(int)
	sessionsFields := schema.SESSIONS{}.Fields()
	_ = sessionsFields
	// sessionsDescMatchedAt is the schema descriptor for matched_at field.
	sessionsDescMatchedAt := sessionsFields[3].Descriptor()
	// sessions.DefaultMatchedAt holds the default value on creation for the matched_at field.
	sessions.DefaultMatchedAt = sessionsDescMatchedAt.Default.(time.Time)
	usersFields := schema.USERS{}.Fields()
	_ = usersFields
	// usersDescUsername is the schema descriptor for username field.
	usersDescUsername := usersFields[0].Descriptor()
	// users.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	users.UsernameValidator = func() func(string) error {
		validators := usersDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[3].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(func() time.Time)
	// usersDescIsDeleted is the schema descriptor for is_deleted field.
	usersDescIsDeleted := usersFields[4].Descriptor()
	// users.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	users.DefaultIsDeleted = usersDescIsDeleted.Default.(bool)
	// usersDescUpdatedAt is the schema descriptor for updated_at field.
	usersDescUpdatedAt := usersFields[5].Descriptor()
	// users.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	users.DefaultUpdatedAt = usersDescUpdatedAt.Default.(func() time.Time)
}
