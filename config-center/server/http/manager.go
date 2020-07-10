package http

func CreateEP(c *gin.Context) {
	var params struct {
		EpName   string `json:"epName" binding:"required"`
		RealName string `json:"realName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.Set(api.ReqError, result.NewError(result.ParamsError).SetChildErr("work", nil, err.Error()))
		return
	}
}
