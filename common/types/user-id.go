package types

type UserID string

func (i UserID) GetUserID() UserID {
	return i
}
