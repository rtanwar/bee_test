package controllers

import (
	"github.com/astaxie/beego"
)

// oprations for World
type WorldController struct {
	beego.Controller
}

func (c *WorldController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create World
// @Param	body		body 	models.World	true		"body for World content"
// @Success 200 {int} models.World.Id
// @Failure 403 body is empty
// @router / [post]
func (c *WorldController) Post() {

}

// @Title Get
// @Description get World by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.World
// @Failure 403 :id is empty
// @router /:id [get]
func (c *WorldController) GetOne() {

}

// @Title Get All
// @Description get World
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.World
// @Failure 403
// @router / [get]
func (c *WorldController) GetAll() {

}

// @Title Update
// @Description update the World
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.World	true		"body for World content"
// @Success 200 {object} models.World
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WorldController) Put() {

}

// @Title Delete
// @Description delete the World
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WorldController) Delete() {

}
