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
	ordersFieldNames          = builderx.RawFieldNames(&Orders{})
	ordersRows                = strings.Join(ordersFieldNames, ",")
	ordersRowsExpectAutoSet   = strings.Join(stringx.Remove(ordersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ordersRowsWithPlaceHolder = strings.Join(stringx.Remove(ordersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheOrderOrdersIdPrefix = "cache:order:orders:id:"
)

type (
	OrdersModel interface {
		Insert(token string) (sql.Result, error)
		FindOne(id int64) (*Orders, error)
		Update(data *Orders) error
		Delete(id int64) error
	}

	defaultOrdersModel struct {
		sqlc.CachedConn
		table string
	}

	Orders struct {
		Id    int64  `db:"id"`
		Token string `db:"token"`
	}
)

func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) OrdersModel {
	return &defaultOrdersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`orders`",
	}
}

func (m *defaultOrdersModel) Insert(token string) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, ordersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, token)

	return ret, err
}

func (m *defaultOrdersModel) FindOne(id int64) (*Orders, error) {
	orderOrdersIdKey := fmt.Sprintf("%s%v", cacheOrderOrdersIdPrefix, id)
	var resp Orders
	err := m.QueryRow(&resp, orderOrdersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
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

func (m *defaultOrdersModel) Update(data *Orders) error {
	orderOrdersIdKey := fmt.Sprintf("%s%v", cacheOrderOrdersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ordersRowsWithPlaceHolder)
		return conn.Exec(query, data.Token, data.Id)
	}, orderOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) Delete(id int64) error {

	orderOrdersIdKey := fmt.Sprintf("%s%v", cacheOrderOrdersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, orderOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrderOrdersIdPrefix, primary)
}

func (m *defaultOrdersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
