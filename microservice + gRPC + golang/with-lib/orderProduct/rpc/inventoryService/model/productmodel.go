package model

import (
	"context"
	"orderProduct/rpc/inventoryService/types"

	log "orderProduct/utils"

	"github.com/globalsign/mgo/bson"

	cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongo"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefixProductCacheKey = "cache:Product:"

type ProductModel interface {
	Insert(ctx context.Context, data *Product) error
	FindOneById(ctx context.Context, id string) (*Product, error)
	FindOneProduct(ctx context.Context, product types.Product) (*Product, error)
	FindManyProductByID(ctx context.Context, productID []string) (*[]Product, error)
	UpdateProduct(ctx context.Context, product types.Product, selector interface{}) error
	UpdateMultiProduct(ctx context.Context, product types.Product, selector interface{}) error
	Delete(ctx context.Context, id string) error
}

type defaultProductModel struct {
	*mongoc.Model
}

func NewProductModel(url, collection string, c cachec.CacheConf) ProductModel {
	return &defaultProductModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultProductModel) FindOneById(ctx context.Context, id string) (*Product, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Product
	key := prefixProductCacheKey + id
	err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) FindManyProductByID(ctx context.Context, productID []string) (*[]Product, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)

	var data []Product
	var bsonID = make([]bson.ObjectId, len(productID))

	for index, id := range productID {
		if !bson.IsObjectIdHex(id) {
			return nil, ErrInvalidObjectId
		}
		bsonID[index] = bson.ObjectIdHex(id)
	}

	query := bson.D{{"_id", bson.D{{"$in", bsonID}}}}

	err = m.GetCollection(session).FindAllNoCache(&data, query)
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) FindOneProduct(ctx context.Context, product types.Product) (*Product, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)

	var data Product

	product.Data["name"] = product.Name
	doc, err := getBSONObject(product.Data)
	log.Info("MONGO DB ", doc)

	err = m.GetCollection(session).FindOneNoCache(&data, doc)
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) UpdateProduct(ctx context.Context, product types.Product, selector interface{}) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	product.Data["name"] = product.Name
	doc, err := getBSONObject(product.Data)
	update := bson.D{{"$inc", bson.D{{"quantity", -product.Quantity}}}}

	defer m.PutSession(session)
	return m.GetCollection(session).UpdateNoCache(doc, update)
}

func (m *defaultProductModel) UpdateMultiProduct(ctx context.Context, product types.Product, selector interface{}) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	// mongo.BulkInserter

	bulkInserter := mongo.NewBulkInserter(session, "inventory", func() string { return "products" })

	bulkInserter.Insert(bson.D{{"$inc", bson.D{{"quantity", -product.Quantity}}}})

	// bulkInserter.SetResultHandler()

	product.Data["name"] = product.Name
	doc, err := getBSONObject(product.Data)
	update := bson.D{{"$inc", bson.D{{"quantity", -product.Quantity}}}}

	defer m.PutSession(session)
	return m.GetCollection(session).UpdateNoCache(doc, update)
}

func (m *defaultProductModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixProductCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}
