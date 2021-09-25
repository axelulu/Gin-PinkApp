package models

type Follow struct {
	UserId   int64 `json:"user_id" db:"user_id"`
	FollowId int64 `json:"follow_id" db:"follow_id"`
}

type FollowId struct {
	FollowId string `json:"follow_id" db:"follow_id"`
}

type Like struct {
	UserId int64 `json:"user_id" db:"user_id"`
	PostId int64 `json:"post_id" db:"post_id"`
	Type   int64 `json:"type" db:"type"`
}

type LikeId struct {
	PostId string `json:"post_id" db:"post_id"`
}

type Star struct {
	UserId int64 `json:"user_id" db:"user_id"`
	PostId int64 `json:"post_id" db:"post_id"`
}

type StarId struct {
	PostId string `json:"post_id" db:"post_id"`
}

type Coin struct {
	UserId int64 `json:"user_id" db:"user_id"`
	PostId int64 `json:"post_id" db:"post_id"`
	Num    int64 `json:"num" db:"num"`
}

type CoinId struct {
	PostId string `json:"post_id" db:"post_id"`
	Coin   string `json:"coin" db:"coin" validate:"oneof=coin 1 2"`
}
