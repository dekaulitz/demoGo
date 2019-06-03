package repository

import (
	"demoGo/apps/handler/exception"
	"demoGo/apps/repository/paging"
	"demoGo/configuration"
	"time"
)

type UsersInterface interface {
	Index() ([]*UsersEntity, *exception.ErrorException)
	Store(user *UsersEntity) (*UsersEntity, *exception.ErrorException)
	Show(id string) (*UsersEntity, *exception.ErrorException)
	Delete(id string) *exception.ErrorException
	Update(id string, user *UsersEntity) *exception.ErrorException
	Paging(paging *paging.Pagination) (*paging.Pagination, *exception.ErrorException)
}

type UsersEntity struct {
	ID        int64     `xorm:"'id' pk autoincr" json:"id"`
	Name      string    `xorm:"'name'" json:"name"`
	Email     string    `xorm:"'email'" json:"email"`
	Password  string    `xorm:"'password'" json:"password"`
	CreatedAt time.Time `xorm:"'created_at' created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"'updated_at' updated" json:"updatedAt"`
}

const (
	tableUsers = "users"
)

var (
	paginationRule = []string{"id", "name", "email", "created_at", "updated_at"}
)

func init() {
	configuration.GetEngine().Sync(new(UsersEntity))
	configuration.GetEngine().ShowSQL(true)
}

type User struct{}

func GetUserRepository() UsersInterface {
	userRespo := &User{}
	return userRespo
}

func (u User) Show(id string) (*UsersEntity, *exception.ErrorException) {
	var users UsersEntity
	sess := configuration.GetConnection()
	defer sess.Close()
	sess.Table(tableUsers)
	sess.ID(id)
	isExist, err := sess.Get(&users)
	if err != nil {
		return nil, exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	if !isExist {
		return nil, exception.NewException(exception.RECORD_NOT_FOUND).Throw(" doesnt exist")
	}
	return &users, nil
}

func (u User) Index() ([]*UsersEntity, *exception.ErrorException) {
	var users []*UsersEntity
	sess := configuration.GetSession()
	sess.Table(tableUsers)
	err := sess.Find(&users)
	if err != nil {
		return nil, exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	return users, nil
}

func (u User) Store(user *UsersEntity) (*UsersEntity, *exception.ErrorException) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	sess := configuration.GetSession()
	sess.Table(tableUsers)
	isSuccess, err := sess.Insert(user)
	if err != nil {
		return nil, exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	if isSuccess == 0 {
		return nil, exception.NewException(exception.FAIL_TO_SAVE).Throw("Data not successfully saved")
	}
	return user, nil
}
func (u User) Delete(id string) *exception.ErrorException {
	var user UsersEntity
	sess := configuration.GetSession()
	sess.Table(tableUsers)
	isSuccess, err := sess.ID(id).Delete(&user)
	if err != nil {
		return exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	if isSuccess == 0 {
		return exception.NewException(exception.RECORD_NOT_FOUND).Throw("Data not successfully deleted")
	}
	return nil
}

func (u User) Update(id string, user *UsersEntity) *exception.ErrorException {
	sess := configuration.GetSession()
	sess.Table(tableUsers)
	sess.ID(id)
	_, err := sess.Update(user)
	if err != nil {
		return exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	return nil
}

func (u User) Paging(property *paging.Pagination) (*paging.Pagination, *exception.ErrorException) {
	users := []UsersEntity{}
	counter := configuration.GetSession()
	counter.Table(tableUsers)
	sess := configuration.GetSession()
	sess.Table(tableUsers)

	/*
		splitting session with 2 session and injecting current session with pagination builder
		remove it if you dont want to use paging feature
	*/
	paging.GetPage().SetQueryBuilder(property, counter, sess, paginationRule)

	err := sess.Find(&users)
	if err != nil {
		return nil, exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	/**
	setup the property of paging
	*/
	property.Rows = users
	property.RowCount, err = counter.Count(&UsersEntity{})
	property.PageCount = paging.GetPage().GetRowCount(property)

	if err != nil {
		return nil, exception.NewException(exception.ERROR_DATABASE_ERROR).Throw(err.Error())
	}
	return property, nil
}
