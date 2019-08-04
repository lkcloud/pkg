package datastore

import "time"

type Setting func(aliasName string)

func WithMaxIdleConns(maxIdleConns int) Setting {
	return func(aliasName string) {
		al := getDbAlias(aliasName)
		al.MaxIdleConns = maxIdleConns
		al.DB.DB().SetMaxIdleConns(maxIdleConns)
	}
}

func WithConnMaxLifetime(d time.Duration) Setting {
	return func(aliasName string) {
		al := getDbAlias(aliasName)
		al.ConnMaxLifetime = d
		al.DB.DB().SetConnMaxLifetime(d)
	}
}

func WithMaxOpenConns(maxOpenConns int) Setting {
	return func(aliasName string) {
		al := getDbAlias(aliasName)
		al.MaxOpenConns = maxOpenConns
		al.DB.DB().SetMaxOpenConns(maxOpenConns)
	}
}
