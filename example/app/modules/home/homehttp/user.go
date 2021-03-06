package homehttp

import (
	"math/rand"
	"net/http"
	"time"

	"xorm.io/xorm"

	"github.com/hulklab/yago/coms/orm"

	"github.com/hulklab/yago/example/app/modules/home/homedao"

	"github.com/hulklab/yago/base/basemodel"

	"github.com/hulklab/yago/example/app/modules/home/homehttp/homemiddleware"

	"github.com/hulklab/yago"
	"github.com/hulklab/yago/base/basehttp"
	"github.com/hulklab/yago/example/app/g"
	"github.com/hulklab/yago/example/app/modules/home/homemodel"
)

type UserHttp struct {
	basehttp.BaseHttp
}

type HttpMetadata struct {
	Label string `json:"label"`
}

func init() {
	userHttp := new(UserHttp)

	// simple route, not recommend
	yago.AddHttpRouter("/user/hello", http.MethodGet, userHttp.HelloAction)
	yago.AddHttpRouter("/user/add", http.MethodPost, userHttp.AddAction)
	yago.AddHttpRouter("/user/delete", http.MethodPost, userHttp.DeleteAction)
	yago.AddHttpRouter("/user/detail", http.MethodGet, userHttp.DetailAction)
	yago.AddHttpRouter("/user/update", http.MethodPost, userHttp.UpdateAction)
	yago.AddHttpRouter("/user/list", http.MethodPost, userHttp.ListAction)
	yago.AddHttpRouter("/user/base-list", http.MethodPost, userHttp.BaseListAction)
	yago.AddHttpRouter("/user/upload", http.MethodPost, userHttp.UploadAction)
	yago.AddHttpRouter("/user/hello/:name", http.MethodGet, userHttp.Hello2Action)
	yago.AddHttpRouter("/user/cookie", http.MethodGet, userHttp.CookieAction)
	yago.AddHttpRouter("/user/metadata", http.MethodGet, userHttp.MetadataAction, HttpMetadata{
		Label: "自定义HTTP名称",
	})

	// routing groups are recommended
	memberGroup := yago.NewHttpGroupRouter("/user/member")
	memberGroup.Use(homemiddleware.CheckUserName)
	{
		memberGroup.Post("/:name", userHttp.UserSetAction)
		memberGroup.Get("/:name", userHttp.UserGetAction)
		memberGroup.Put("/:name", userHttp.UserUpdateAction)
		memberGroup.Delete("/:name", userHttp.UserDeleteAction)

		consumeSubGroup := memberGroup.Group("/consume")
		consumeSubGroup.Use(homemiddleware.ComputeConsume)
		consumeSubGroup.Patch("/sleep/:name", userHttp.ConsumeSleepAction)
	}

	yago.SetHttpNoRouter(userHttp.NoRouterAction)
}

func (h *UserHttp) NoRouterAction(c *yago.Ctx) {
	c.JSON(http.StatusNotFound, g.Hash{
		"error": "404, page not exists",
	})
}

// curl -X GET 'http://127.0.0.1:8080/user/hello?username=zhangsan'
func (h *UserHttp) HelloAction(c *yago.Ctx) {
	var p struct {
		Username string `json:"username" validate:"omitempty,max=20" form:"username" label:"姓名"`
	}

	err := c.ShouldBind(&p)
	if err != nil {
		c.SetError(err)
		return
	}

	data := "hello " + p.Username

	c.SetData(data)
}

// curl 'http://127.0.0.1:8080/user/add' -H "Content-type:application/x-www-form-urlencoded" -XPOST -d "username=lisi&phone=13090001112"
func (h *UserHttp) AddAction(c *yago.Ctx) {
	var p struct {
		Username string `json:"username" validate:"required,max=20" form:"username" label:"姓名"`
		Phone    string `json:"phone" validate:"required,phone" form:"phone" label:"手机号"`
	}

	err := c.ShouldBind(&p)
	if err != nil {
		c.SetError(err)
		return
	}

	err = orm.Ins().Transactional(func(session *xorm.Session) error {
		_, err := homemodel.NewUserModel(basemodel.WithSession(session)).Add(p.Username, p.Phone, nil)
		if err != nil {
			return err
		}
		// other model (create update delete) method in the same transaction ......
		return nil
	}, orm.WithContext(c))

	c.SetDataOrErr("OK", err)
}

var p struct {
	Id int64 `json:"id" validate:"required" form:"id" label:"Id"`
}

// curl 'http://127.0.0.1:8080/user/delete' -H "Content-type:application/json" -XPOST -d '{"id":1}'
func (h *UserHttp) DeleteAction(c *yago.Ctx) {

	err := c.ShouldBind(&p)
	if err != nil {
		c.SetError(err)
		return
	}

	model := homemodel.NewUserModel()

	data, err := model.DeleteById(p.Id)

	c.SetDataOrErr(data, err)
}

// curl 'http://127.0.0.1:8080/user/detail?id=2' -H "Content-type:application/json" -XGET
func (h *UserHttp) DetailAction(c *yago.Ctx) {
	err := c.ShouldBind(&p)
	if err != nil {
		c.SetError(err)
		return
	}

	model := homemodel.NewUserModel()

	data, err := model.GetDetail(p.Id)

	c.SetDataOrErr(data, err)
}

// curl 'http://127.0.0.1:8080/user/update' -H "Content-type:application/json" -XPOST -d '{"id":2,"username":"zhangsan"}'
func (h *UserHttp) UpdateAction(c *yago.Ctx) {
	var p struct {
		Id       int64  `json:"id" validate:"required" form:"id" label:"Id"`
		Username string `json:"username" validate:"required" form:"username" label:"姓名"`
	}

	err := c.ShouldBind(&p)
	if err != nil {
		c.SetError(err)
		return
	}

	model := homemodel.NewUserModel()

	var options = make(map[string]interface{})

	if p.Username != "" {
		options["username"] = p.Username
	}

	user, err := model.UpdateById(p.Id, options)

	c.SetDataOrErr(user, err)
}

// curl 'http://127.0.0.1:8080/user/list' -H "Content-type:application/json" -XPOST -d '{"pagesize":1}'
func (h *UserHttp) ListAction(c *yago.Ctx) {
	type p struct {
		Q        string `json:"q" validate:"omitempty" form:"q"`
		Page     int    `json:"page" validate:"omitempty" form:"name" label:"当前页"`
		Pagesize int    `json:"pagesize" validate:"omitempty" form:"pagesize" label:"页大小"`
	}

	pi := p{
		Page:     1,
		Pagesize: 10,
	}

	err := c.ShouldBind(&pi)
	if err != nil {
		c.SetError(err)
		return
	}

	model := homemodel.NewUserModel()
	total, users := model.GetList(pi.Q, pi.Page, pi.Pagesize)
	c.SetData(map[string]interface{}{
		"total": total,
		"list":  users,
	})
}

// curl 'http://127.0.0.1:8080/user/base-list' -H "Content-type:application/json" -XPOST -d '{"page":1}'
func (h *UserHttp) BaseListAction(c *yago.Ctx) {
	type p struct {
		Page    int               `json:"page" validate:"min=1"`
		Size    int               `json:"size" validate:"oneof=10 20 50 100"`
		Q       string            `json:"q" validate:"-"`
		Filters basemodel.Filters `json:"filters" validate:"-"`
		Orders  basemodel.Orders  `json:"orders" validate:"-"`
	}

	pi := p{
		Page: 1,
		Size: 10,
	}

	err := c.ShouldBindJSON(&pi)
	if err != nil {
		c.SetError(err)
		return
	}

	var users []*homedao.UserDao

	model := homemodel.NewUserModel()

	total, err := model.PageList(&basemodel.PageQuery{
		Page: pi.Page,
		Size: pi.Size,
		Q: basemodel.Q{
			pi.Q: {
				"username",
			},
		},
		Orders:  pi.Orders,
		Filters: pi.Filters,
	}, &users)

	c.SetDataOrErr(g.Hash{
		"total": total,
		"list":  users,
	}, err)
}

func (h *UserHttp) UploadAction(c *yago.Ctx) {

	file, _ := c.FormFile("file")

	// Upload the file to specific dst.
	if err := c.SaveUploadedFile(file, "/Users/xxx/Downloads/upload_test.png"); err != nil {
		c.SetError(yago.NewErr(err.Error()))
		return
	}

	c.SetData(file.Filename)
}

func (h *UserHttp) Hello2Action(c *yago.Ctx) {
	name := c.Param("name")

	c.SetData("hello " + name)
}

func (h *UserHttp) UserSetAction(c *yago.Ctx) {
	name := c.Param("name")

	c.SetData("set " + name)
}

func (h *UserHttp) UserGetAction(c *yago.Ctx) {
	name := c.Param("name")

	c.SetData("get " + name)
}

func (h *UserHttp) UserUpdateAction(c *yago.Ctx) {
	name := c.Param("name")

	c.SetData("update " + name)
}

func (h *UserHttp) UserDeleteAction(c *yago.Ctx) {
	name := c.Param("name")

	c.SetData("delete " + name)
}

func (h *UserHttp) CookieAction(c *yago.Ctx) {
	cookie, err := c.Cookie("user")

	if err != nil {
		c.SetError(err)
		return
	}

	c.SetData("hello " + cookie)
}

func (h *UserHttp) ConsumeSleepAction(c *yago.Ctx) {
	c.SetData("I'm sleeping zzz.....")
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}

func (h *UserHttp) MetadataAction(c *yago.Ctx) {
	data := "get label from metadata:"

	for _, router := range yago.GetHttpRouters() {
		if router.Url() == "/user/metadata" {
			v, ok := router.Metadata.([]interface{})
			if ok {
				data = data + v[0].(HttpMetadata).Label
			}
			break
		}
	}

	c.SetData(data)
}
