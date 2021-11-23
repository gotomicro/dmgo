/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gotomicro/dmgo/i18n"
)

// 发版标记
//var version = "8.1.2.38"
//var build_date = "2021.07.14"
//var svn = "7050"

var globalDmDriver = newDmDriver()

func init() {
	sql.Register("dm", globalDmDriver)
}

func driverInit() {
	switch Locale {
	case 0:
		i18n.InitConfig(i18n.Messages_zh_CN)
	case 1:
		i18n.InitConfig(i18n.Messages_en_US)
	case 2:
		i18n.InitConfig(i18n.Messages_zh_TW)
	}
}

type Driver struct {
	filterable
	readPropMutex sync.Mutex
}

func newDmDriver() *Driver {
	d := new(Driver)
	d.idGenerator = dmDriverIDGenerator
	return d
}

func (d *Driver) Open(dsn string) (driver.Conn, error) {
	return d.open(dsn)
}

func (d *Driver) OpenConnector(dsn string) (driver.Connector, error) {
	return d.openConnector(dsn)
}

func (d *Driver) open(dsn string) (*Connection, error) {
	c, err := d.openConnector(dsn)
	if err != nil {
		return nil, err
	}
	return c.connect(context.Background())
}

func (d *Driver) openConnector(dsn string) (*Connector, error) {
	connector := new(Connector).init()
	connector.url = dsn
	connector.dmDriver = d
	d.readPropMutex.Lock()
	err := connector.mergeConfigs(dsn)
	d.readPropMutex.Unlock()
	if err != nil {
		return nil, err
	}
	connector.createFilterChain(connector, nil)
	return connector, nil
}
