// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Hosi121/SpeakUp/ent/aithemes"
	"github.com/Hosi121/SpeakUp/ent/calls"
	"github.com/Hosi121/SpeakUp/ent/friends"
	"github.com/Hosi121/SpeakUp/ent/matchings"
	"github.com/Hosi121/SpeakUp/ent/memos"
	"github.com/Hosi121/SpeakUp/ent/schema"
	"github.com/Hosi121/SpeakUp/ent/sessions"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	aithemesFields := schema.AITHEMES{}.Fields()
	_ = aithemesFields
	// aithemesDescCreatedAt is the schema descriptor for created_at field.
	aithemesDescCreatedAt := aithemesFields[1].Descriptor()
	// aithemes.DefaultCreatedAt holds the default value on creation for the created_at field.
	aithemes.DefaultCreatedAt = aithemesDescCreatedAt.Default.(time.Time)
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
	friendsFields := schema.FRIENDS{}.Fields()
	_ = friendsFields
	// friendsDescCreatedAt is the schema descriptor for created_at field.
	friendsDescCreatedAt := friendsFields[3].Descriptor()
	// friends.DefaultCreatedAt holds the default value on creation for the created_at field.
	friends.DefaultCreatedAt = friendsDescCreatedAt.Default.(func() time.Time)
	matchingsFields := schema.MATCHINGS{}.Fields()
	_ = matchingsFields
	// matchingsDescMatchedAt is the schema descriptor for matched_at field.
	matchingsDescMatchedAt := matchingsFields[3].Descriptor()
	// matchings.DefaultMatchedAt holds the default value on creation for the matched_at field.
	matchings.DefaultMatchedAt = matchingsDescMatchedAt.Default.(time.Time)
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
	sessionsFields := schema.SESSIONS{}.Fields()
	_ = sessionsFields
	// sessionsDescCreatedAt is the schema descriptor for created_at field.
	sessionsDescCreatedAt := sessionsFields[3].Descriptor()
	// sessions.DefaultCreatedAt holds the default value on creation for the created_at field.
	sessions.DefaultCreatedAt = sessionsDescCreatedAt.Default.(time.Time)
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
	// usersDescEmail is the schema descriptor for email field.
	usersDescEmail := usersFields[1].Descriptor()
	// users.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	users.EmailValidator = func() func(string) error {
		validators := usersDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[4].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(func() time.Time)
	// usersDescIsDeleted is the schema descriptor for is_deleted field.
	usersDescIsDeleted := usersFields[5].Descriptor()
	// users.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	users.DefaultIsDeleted = usersDescIsDeleted.Default.(bool)
	// usersDescUpdatedAt is the schema descriptor for updated_at field.
	usersDescUpdatedAt := usersFields[6].Descriptor()
	// users.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	users.DefaultUpdatedAt = usersDescUpdatedAt.Default.(func() time.Time)
	// usersDescAccessToken is the schema descriptor for access_token field.
	usersDescAccessToken := usersFields[7].Descriptor()
	// users.DefaultAccessToken holds the default value on creation for the access_token field.
	users.DefaultAccessToken = usersDescAccessToken.Default.(string)
}
