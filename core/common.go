package core

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type CommonService struct {
	Engine *xorm.Engine
}

type InitData struct {
	Banner []model.Article `json:"banner"`
	Up     []model.Article `json:"up"`
	List   []model.Article `json:"list"`
}

func (cs CommonService) GetInitData(ctx *gin.Context) {
	ad := &model.ArticleData{
		Engine: cs.Engine,
	}
	banner, _ := ad.GetList(true, false, "", -1, model.Official, 1, 20)
	up, _ := ad.GetList(false, true, "", -1, model.Official, 1, 20)
	list, _ := ad.GetList(true, true, "", -1, model.Official, 1, 20)
	data := &InitData{
		Banner: banner,
		Up:     up,
		List:   list,
	}
	util.ReturnData(ctx, true, data)
}

func (cs CommonService) GetUserInfo(ctx *gin.Context) {
	check := util.GetString("github.bind")
	if check != "" {
		var checkUser User
		err := json.Unmarshal([]byte(check), &checkUser)
		if err != nil {
			util.ReturnMessage(ctx, false, "获取博主信息失败")
		} else {
			checkUser.Id = ""
			util.ReturnData(ctx, true, checkUser)
		}
	} else {
		util.ReturnMessage(ctx, false, "系统尚未初始化")
	}
}
