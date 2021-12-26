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
	productsFieldNames          = builderx.RawFieldNames(&Products{})
	productsRows                = strings.Join(productsFieldNames, ",")
	productsRowsExpectAutoSet   = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	productsRowsWithPlaceHolder = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheOrderProductsIdPrefix = "cache:order:products:id:"
)

type (
	ProductsModel interface {
		Insert(data *Products) (sql.Result, error)
		FindOne(id int64) (*Products, error)
		Update(data *Products) error
		Delete(id int64) error
	}

	defaultProductsModel struct {
		sqlc.CachedConn
		table string
	}

	Products struct {
		Id       int             `db:"id"`
		OrderId  int             `db:"order_id"`
		Name     string          `db:"name"`
		Cost     sql.NullInt64   `db:"cost"`
		Weight   sql.NullFloat64 `db:"weight"`
		Quantity int             `db:"quantity"`
		Data     sql.NullString  `db:"data"`
	}
)

func NewProductsModel(conn sqlx.SqlConn, c cache.CacheConf) ProductsModel {
	return &defaultProductsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`products`",
	}
}

func (m *defaultProductsModel) Insert(data *Products) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, productsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.OrderId, data.Name, data.Cost, data.Weight, data.Quantity, data.Data)

	return ret, err
}

func (m *defaultProductsModel) FindOne(id int64) (*Products, error) {
	orderProductsIdKey := fmt.Sprintf("%s%v", cacheOrderProductsIdPrefix, id)
	var resp Products
	err := m.QueryRow(&resp, orderProductsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
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

func (m *defaultProductsModel) Update(data *Products) error {
	orderProductsIdKey := fmt.Sprintf("%s%v", cacheOrderProductsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productsRowsWithPlaceHolder)
		return conn.Exec(query, data.OrderId, data.Name, data.Cost, data.Weight, data.Quantity, data.Data, data.Id)
	}, orderProductsIdKey)
	return err
}

func (m *defaultProductsModel) Delete(id int64) error {

	orderProductsIdKey := fmt.Sprintf("%s%v", cacheOrderProductsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, orderProductsIdKey)
	return err
}

func (m *defaultProductsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrderProductsIdPrefix, primary)
}

func (m *defaultProductsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
