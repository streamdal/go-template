package badger

import (
	"context"

	dbadger "github.com/dgraph-io/badger"

	"github.com/sirupsen/logrus"

	"github.com/batchcorp/go-template/types"
)

type IBadger interface {
	Get(key string) ([]byte, error)
	Add(key string, data []byte) error
	Delete(key string) error
}

type Badger struct {
	DB  *dbadger.DB
	ctx context.Context
	log *logrus.Entry
}

func New(badgerDir string, ctx context.Context) (*Badger, error) {
	db, err := dbadger.Open(dbadger.DefaultOptions(badgerDir))
	if err != nil {
		return nil, err
	}

	b := &Badger{
		DB:  db,
		ctx: ctx,
		log: logrus.WithField("pkg", "badger"),
	}

	if b.ctx == nil {
		b.ctx = context.Background()
	}

	return b, nil
}

func (b *Badger) Get(key string) ([]byte, error) {
	var val []byte
	err := b.DB.View(func(txn *dbadger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == dbadger.ErrKeyNotFound {
				return types.NotFoundErr
			}
		}

		// fuck this thing
		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (b *Badger) Add(key string, data []byte) error {
	return b.DB.Update(func(txn *dbadger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

func (b *Badger) Delete(key string) error {
	return b.DB.Update(func(txn *dbadger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
