package util

import (
	"encoding/base64"
	"time"

	"github.com/skye-z/betax-blog/model"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func InitDB() *xorm.Engine {
	OutLogf("Data", "load engine")
	engine, err := xorm.NewEngine("sqlite", "./local.store")
	if err != nil {
		panic(err)
	}
	return engine
}

func InitDBTable(engine *xorm.Engine) {
	OutLogf("Data", "load table")
	err := engine.Sync2(new(model.Class))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.TagConnect))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Tag))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Article))
	if err != nil {
		panic(err)
	}
	// 初始化分类
	cd := &model.ClassData{
		Engine: engine,
	}
	classList, err := cd.GetList()
	if err == nil {
		if len(classList) == 0 {
			cd.Add(&model.Class{
				Id:    1,
				Name:  "随手记",
				Super: 0,
			})
		}
	}
	// 初始化文章
	ad := &model.ArticleData{
		Engine: engine,
	}
	number, err := ad.GetNumber("", -1)
	if err == nil {
		if number == 0 {
			content, _ := base64.StdEncoding.DecodeString("5qyi6L+O5L2/55SoYEJldGFYIEJsb2dgfgoK6L+Z5piv5LiA5Liq5bCP5ben57K+6Ie044CB5Lqn54mp5Y2V5LiA5LiU5byA6ZSA5p6B5L2O55qE5Yqo5oCB5Y2a5a6i56iL5bqPLgoKIyMg5Li76KaB5Yqf6IO9CgoxLiBNYXJrZG93biDmlK/mjIEKMi4g5Zu+54mH5LiK5LygCjMuIEFJIOaRmOimgQo0LiDpu5HmmpfmqKHlvI8KNS4gR2l0aHViIOaOpeWFpQoK6Zmk5q2k5LmL5aSW5bCx5piv5LiA5Liq5Y2a5a6i56iL5bqP5bqU5b2T5pyJ55qE5Yqf6IO95ZWmLCDmr5TlpoLmlofnq6Dnva7pobbjgIHova7mkq3jgIHmkJzntKLjgIHpmo/mnLrliJfooajlkozmoIfnrb7jgIHliIbnsbvjgIHmlofku7bnrqHnkIbnrYnlip/og70uCgojIyDliJ3lp4vljJborr7nva4KCuWuieijheWlvWBCZXRhWCBCbG9nYOWQjuiuv+mXruW8gOaUvuerr+WPoywg5L2g5Lya55yL5Yiw5LiA5LiqYOWIneWni+WMlumFjee9rmDnmoTpobXpnaIsIOWcqOi/meS4qumhtemdouS4reS9oOacgOWwkemcgOimgeWhq+WGmWBgR2l0aHViIOeZu+W9lWBg55qE5LiJ5Liq6YWN572u6aG5LgoK5pys5Y2a5a6i55u05o6l5L2/55So5LqGR2l0aHVi5o6I5p2D55qE5pa55byP6L+b6KGM55m75b2VLgoK5YW25Lit5Zue6LCD5Zyw5Z2A5Lya5qC55o2u5L2g6K6/6Zeu5Y2a5a6i55qE5Zyw5Z2A5p2l6Ieq5Yqo55Sf5oiQLCDop4TliJnkuLpgaHR0cFtzXTovL+Wfn+WQjTrnq6/lj6Mvb2F1dGgyL2NhbGxiYWNrYC4KCuWPpuWkluS4pOWPguaVsOmcgOimgeS9oFvlnKggR2l0aHViIOS4reWIm+W7uiBPQXV0aDIg5bqU55SoXShgaHR0cHM6Ly9naXRodWIuY29tL3NldHRpbmdzL2RldmVsb3BlcnNgKSwg54S25ZCO5aGr5YWlR2l0aHViIOe7meWHuueahGBDbGllbnQgSWRg5ZKMYFNlY3JldGAuCgojIyDnrqHnkIblkI7lj7AKCuaDs+W/heS9oOW3sue7j+WujOaIkOWIneWni+WMluiuvue9ruS6hiwg5q2j5Zyo55aR5oOR6KaB5aaC5L2V6L+b5YWl5ZCO5Y+w5ZGiPwoK5L2g55yL6aG16Z2i5bqV6YOoYEJ1aWx0IG9uIEJldGFYIEJsb2dg5ZCO6Z2i5pyJ5Liq5Zu+5qCHLCDngrnlh7vlroPlsLHog73nmbvlvZXnrqHnkIblkI7lj7DllaYuCgoK5pyA5ZCOLCDpnZ7luLjmhJ/osKLkvaDog73pgInmi6lgQmV0YVggQmxvZ2AsIOW4jOacm+S9oOiDveS/neaMgemaj+aJi+iusOW9leeahOWlveS5oOaDrywg5Zyo5ZCO5Y+w5Yig6Zmk5pys5paH5ZCO5byA5aeL5L2g55qE56ys5LiA56+H5paH56ug5ZCnIQo=")
			ad.Add(&model.Article{
				Id:           1,
				Weight:       0,
				IsBanner:     false,
				IsUp:         false,
				Type:         model.Markdown,
				Title:        "使用前必看!",
				Abstract:     "欢迎使用 BetaX Blog, 在开始使用前不妨看看本文, 本文介绍了它的主要功能、初始化设置要点和管理后台进入方式",
				Class:        1,
				Content:      string(content),
				State:        model.Official,
				CreationTime: time.Now().Unix() * 1000,
				ReleaseTime:  time.Now().Unix() * 1000,
			})
		}
	}
}
