package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usersFieldNames          = builderx.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserPrefix = "cache#User#user#"
)

type (
	UsersModel interface {
		Insert(data *Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		Update(data *Users) error
		Delete(id int64) error
		FindOneByToken(token string) (*Users, error)
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id    int64  `db:"id"`
		Token string `db:"token"`
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Token)

	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserPrefix, id)
	var resp Users
	err := m.QueryRow(&resp, userUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByToken(token string) (*Users, error) {
	userToken := fmt.Sprintf("%s%v", cacheUserPrefix, token)
	var resp Users
	err := m.QueryRow(&resp, userToken, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `token` = ?", usersRows, m.table)
		return conn.QueryRow(v, query, token)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(data *Users) error {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.Exec(query, data.Token, data.Id)
	}, userUsersIdKey)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {

	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
