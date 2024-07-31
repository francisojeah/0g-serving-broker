package db

import (
	"strings"

	"github.com/0glabs/0g-serving-agent/provider/model"
	"github.com/pkg/errors"
)

func (d *DB) GetUserAccount(userAddress string) (model.User, error) {
	account := model.User{}
	ret := d.db.Where(&model.User{User: userAddress}).First(&account)
	return account, ret.Error
}

func (d *DB) CreateUserAccounts(accounts []model.User) error {
	ret := d.db.Create(&accounts)
	return ret.Error
}

func (d *DB) ListUserAccount(opt *model.UserListOptions) ([]model.User, error) {
	tx := d.db.Model(model.User{})

	if opt != nil {
		if opt.MinUnsettledFee != nil {
			tx = tx.Where("unsettled_fee > ?", *opt.MinUnsettledFee)
		}
		if opt.MaxLastBalanceCheckTime != nil {
			tx = tx.Where("last_balance_check_time < ?", *opt.MaxLastBalanceCheckTime)
		}
	}
	list := []model.User{}
	ret := tx.Order("created_at DESC").Find(&list)
	return list, ret.Error
}

func (d *DB) DeleteUserAccounts(userAddresses []string) error {
	if len(userAddresses) == 0 {
		return nil
	}
	return d.db.Where("name IN (?)", userAddresses).Delete(&model.User{}).Error
}

func (d *DB) UpdateUserAccount(userAddress string, new model.User) error {
	old := model.User{}
	ret := d.db.Where(&model.User{User: userAddress}).First(&old)
	if ret.Error != nil {
		errors.Wrap(ret.Error, "get account from db")
	}
	if err := model.ValidateUpdateUser(old, new); err != nil {
		return err
	}
	if new.LastBalanceCheckTime != nil {
		old.LastBalanceCheckTime = new.LastBalanceCheckTime
	}
	if new.LastRequestNonce != nil {
		old.LastRequestNonce = new.LastRequestNonce
	}
	if new.LastResponseTokenCount != nil {
		old.LastResponseTokenCount = new.LastResponseTokenCount
	}
	if new.LockBalance != nil {
		old.LockBalance = new.LockBalance
	}
	if new.UnsettledFee != nil {
		old.UnsettledFee = new.UnsettledFee
	}

	ret = d.db.Where(&model.User{User: old.User}).Updates(old)
	return ret.Error
}

func (d *DB) BatchUpdateUserAccount(news []model.User) error {
	olds, err := d.ListUserAccount(nil)
	if err != nil {
		return err
	}
	oldAccountMap := make(map[string]bool, len(olds))
	for _, old := range olds {
		oldAccountMap[strings.ToLower(old.User)] = true
	}

	var toAdd, toUpdate []model.User
	var toRemove []string
	for i, new := range news {
		key := strings.ToLower(new.User)
		if oldAccountMap[key] {
			delete(oldAccountMap, key)
			// BatchUpdateUserAccount is currently used to synchronize accounts from the contract to the database.
			// All new data should be updated in the database as each record has a new LastBalanceCheckTime.
			toUpdate = append(toUpdate, news[i])
			continue
		}
		toAdd = append(toAdd, news[i])
	}
	for k := range oldAccountMap {
		toRemove = append(toRemove, k)
	}

	// TODO: add Redis RW lock
	if err := d.CreateUserAccounts(toAdd); err != nil {
		return err
	}
	for i := range toUpdate {
		if ret := d.db.Where(&model.User{User: toUpdate[i].User}).Updates(toUpdate[i]); ret.Error != nil {
			return ret.Error
		}
	}
	return d.DeleteUserAccounts(toRemove)
}
