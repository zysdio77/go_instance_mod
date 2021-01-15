package modle

type ReadUserData struct {
	Id           int    `gorm:"id"`
	Channel      string `gorm:"channel"`
	Platform     string `gorm:"platform"`
	DevModel     string `gorm:"dev_model"`
	BindFb       int    `gorm:"bind_fb"`
	BindApple    int    `gorm:"bind_apple"`
	Lv           int    `gorm:"lv"`
	CoinNum      string `gorm:"coin_num"`
	PayNum       int64  `gorm:"pay_num"`
	BuyCoin      string `gorm:"buy_coin"`
	FreeCoin     string `gorm:"free_coin"`
	TotalBet     string `gorm:"total_bet"`
	TotalWin     string `gorm:"total_win"`
	Vip          int    `gorm:"vip"`
	UserGambling string `gorm:"user_gambling"`
	WildStamp    string `gorm:"wild_stamp"`
	UserValue    string `gorm:"user_value"`
	CreateTime   string `gorm:"create_time"`
	UpdateTime   string `gorm:"update_time"`
	LogoutTime   string `gorm:"logout_time"`
	GuideF       string `gorm:"guide_f"`
	GuideNf      string `gorm:"guide_nf"`
	InboxDetail  string `gorm:"inbox_detail"`
	Extend       string `gorm:"extend"`
}
type WriteUserData struct {
	Id           int    `gorm:"id"`
	Channel      string `gorm:"channel"`
	Platform     string `gorm:"platform"`
	DevModel     string `gorm:"dev_model"`
	BindFb       int    `gorm:"bind_fb"`
	BindApple    int    `gorm:"bind_apple"`
	Lv           int    `gorm:"lv"`
	CoinNum      string `gorm:"coin_num"`
	PayNum       int64  `gorm:"pay_num"`
	BuyCoin      string `gorm:"buy_coin"`
	FreeCoin     string `gorm:"free_coin"`
	TotalBet     string `gorm:"total_bet"`
	TotalWin     string `gorm:"total_win"`
	Vip          int    `gorm:"vip"`
	UserGambling string `gorm:"user_gambling"`
	WildStamp    string `gorm:"wild_stamp"`
	UserValue    string `gorm:"user_value"`
	LogoutTime   string `gorm:"logout_time"` //logout_time
	GuideF       string `gorm:"guide_f"`
	GuideNf      string `gorm:"guide_nf"`
	InboxDetail  string `gorm:"inbox_detail"`
	Extend       string `gorm:"extend"`
}

type ReadFreeData struct {
	Id             int    `gorm:"id"`
	Uid            int    `gorm:"uid"`
	Channel        string `gorm:"channel"`
	Platform       string `gorm:"platform"`
	FreeCoinDetail string `gorm:"free_coin_detail"`
	DateStr        string `gorm:"date_str"`
	Version        string `gorm:"version"`
	CreateTime     string `gorm:"create_time"`
	Extend         string `gorm:"extend"`
}

type WriteFreeData struct {
	Uid            int    `gorm:"uid"`
	Channel        string `gorm:"channel"`
	Platform       string `gorm:"platform"`
	FreeCoinDetail string `gorm:"free_coin_detail"`
	DateStr        string `gorm:"date_str"`
	Version        string `gorm:"version"`
}
