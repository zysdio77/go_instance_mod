package variable

import (
	"time"
)

type UserDataJson struct {
	UID          int    `json:"uid"`
	Lv           int    `json:"lv"`
	CoinNum      string `json:"coinNum"`
	FreeCoin     string `json:"freeCoin"`
	TotalBet     string `json:"totalBet"`
	TotalWin     string `json:"totalWin"`
	Vip          int    `json:"vip"`
	UserGambling string `json:"userGambling"`
	WildStamp    string `json:"wildStamp"`
	UserValue    string `json:"userValue"`
	LogoutTime   string `json:"logoutTime"`
	Version      string `json:"version"`
	GuideF       string `json:"guideF"`
	GuideNF      string `json:"guideNF"`
	Sig          string `json:"sig"` //验证位
	//Token        string `json:"token"`
	Platform string `json:"platform"`
}

type UserDataDb struct {
	Id           int64     `json:"id" form:"id" db:"id" gorm:"id"`
	Channel      string    `json:"channel"`
	Platform     string    `json:"platform"`
	DevModel     string    `json:"dev_model"`
	BindFb       int       `json:"bind_fb"`
	BindApple    int       `json:"bind_apple"`
	Lv           int       `json:"lv"`
	CoinNum      string    `json:"coin_num"`
	PayNum       int64     `json:"pay_num"`
	BuyCoin      string    `json:"buy_coin"`
	FreeCoin     string    `json:"free_coin"`
	TotalBet     string    `json:"total_bet"`
	TotalWin     string    `json:"total_win"`
	Vip          int       `json:"vip"`
	UserGambling string    `json:"user_gambling"`
	WildStamp    string    `json:"wild_stamp"`
	UserValue    string    `json:"user_value"`
	CreateTime   time.Time `json:"create_time"`
	UpdateRime   time.Time `json:"update_time"`
	LogoutRime   time.Time `json:"logout_time"`
	GuideF       string    `json:"guide_f"`
	GuideNf      string    `json:"guide_nf"`
	InboxDetail  string    `json:"inbox_detail"`
	Extend       string    `json:"extend"`
}
type UserData struct {
	Id           int64     `db:"id"`
	Channel      string    `db:"channel"`
	Platform     string    `db:"platform"`
	DevModel     string    `db:"dev_model"`
	BindFb       int       `db:"bind_fb"`
	BindApple    int       `db:"bind_apple"`
	Lv           int       `db:"lv"`
	CoinNum      string    `db:"coin_num"`
	PayNum       int64     `db:"pay_num"`
	BuyCoin      string    `db:"buy_coin"`
	FreeCoin     string    `db:"free_coin"`
	TotalBet     string    `db:"total_bet"`
	TotalWin     string    `db:"total_win"`
	Vip          int       `db:"vip"`
	UserGambling string    `db:"user_gambling"`
	WildStamp    string    `db:"wild_stamp"`
	UserValue    string    `db:"user_value"`
	LogoutTime   string `db:"logout_time"`
	GuideF       string    `db:"guide_f"`
	GuideNf      string    `db:"guide_nf"`
	InboxDetail  string    `db:"inbox_detail"`
	Extend       string    `db:"extend"`
}

type FreeCoinDetail struct {
}
type c struct {
}
type d struct {
}
type e struct {
}
type f struct {
}
type g struct {
}
