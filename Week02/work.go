package Week02

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/astaxie/beego/orm"
	"time"
)
type Blog struct {
	title string
	addtime  time.Time
}
//dao

func (b *Blog) Find(blogId int) (*Blog, error) {
	o := orm.NewOrm()
	err := o.QueryTable(b.TableNameWithPrefix()).Filter("blog_id", blogId).One(b)
	if err != nil {
		return errors.Wrap(sql.ErrNoRows, "set user failed")
	}

	return b.Link()
}

func IsErrNoRows(err error) bool {
	if errors.Cause(err) == sql.ErrNoRows {
		return true
	}
	return false
}

//service
func (c *BlogController) Download() {
	blog, err := models.NewBlog().Find(blogId)
	if err != nil {
		s.Equal(IsErrNoRows(err), true)
	}
}