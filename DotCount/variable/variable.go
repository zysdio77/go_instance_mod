package variable

type UserDtail struct {
	Id           int64  `json:"id" form:"id" db:"id" gorm:"id"`
	Channel      string `json:"channel"`
	Platform     string `json:"platform"`
	DevModel     string `json:"dev_model"`
	BindFb       int    `json:"bind_fb"`
	BindApple    int    `json:"bind_apple"`
	Lv           int    `json:"lv"`
	CoinNum      string `json:"coin_num"`
	PayNum       int64  `json:"pay_num"`
	BuyCoin      string `json:"buy_coin"`
	FreeCoin     string `json:"free_coin"`
	TotalBet     string `json:"total_bet"`
	TotalWin     string `json:"total_win"`
	Vip          int    `json:"vip"`
	UserGambling string `json:"user_gambling"`
	WildStamp    string `json:"wild_stamp"`
	UserValue    string `json:"user_value"`
	CreateTime   string `json:"create_time"`
	UpdateRime   string `json:"update_time"`
	LogoutRime   string `json:"logout_time"`
	GuideF       string `json:"guide_f"`
	GuideNf      string `json:"guide_nf"`
	InboxDetail  string `json:"inbox_detail"`
	Extend       string `json:"extend"`
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
