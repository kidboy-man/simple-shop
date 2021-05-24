package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/kidboy-man/simple-shop/conf"
	"github.com/kidboy-man/simple-shop/datatransfers"
	usecase "github.com/kidboy-man/simple-shop/usecases"
)

type MetadataPublicController struct {
	beego.Controller
	metadataUcase usecase.MetadataUsecase
}

func (c *MetadataPublicController) Prepare() {
	c.metadataUcase = usecase.NewMetadataUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Metadata
// @Description get all metadata
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *MetadataPublicController) GetAll(limit, page int) interface{} {
	metadata, _, err := c.metadataUcase.GetAll(&datatransfers.ListQueryParams{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		c.Data["error"] = err.Error()
		return c.Data
	}

	json, err := json.Marshal(metadata)
	if err != nil {
		c.Data["error"] = err.Error()
		return c.Data
	}

	return string(json)
}
