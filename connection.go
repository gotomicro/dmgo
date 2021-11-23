/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync/atomic"

	"github.com/gotomicro/dmgo/parser"
	"golang.org/x/text/encoding"
)

type Connection struct {
	filterable

	dmConnector        *Connector
	Access             *Access
	stmtMap            map[int32]*DmStatement
	stmtPool           []stmtPoolInfo
	lastExecInfo       *execRetInfo
	lexer              *parser.Lexer
	encode             encoding.Encoding
	encodeBuffer       *bytes.Buffer
	transformReaderDst []byte
	transformReaderSrc []byte

	serverEncoding     string
	GlobalServerSeries int
	ServerVersion      string
	Malini2            bool
	Execute2           bool
	LobEmptyCompOrcl   bool
	IsoLevel           int32
	ReadOnly           bool
	NewLobFlag         bool
	sslEncrypt         int
	MaxRowSize         int32
	DDLAutoCommit      bool
	BackslashEscape    bool
	SvrStat            int32
	SvrMode            int32
	ConstParaOpt       bool
	DbTimezone         int16
	LifeTimeRemainder  int16
	InstanceName       string
	Schema             string
	LastLoginIP        string
	LastLoginTime      string
	FailedAttempts     int32
	LoginWarningID     int32
	GraceTimeRemainder int32
	Guid               string
	DbName             string
	StandbyHost        string
	StandbyPort        int32
	StandbyCount       int32
	SessionID          int64
	OracleDateLanguage byte
	FormatDate         string
	FormatTimestamp    string
	FormatTimestampTZ  string
	FormatTime         string
	FormatTimeTZ       string
	Local              bool
	MsgVersion         int32
	TrxStatus          int32
	dscControl         bool
	trxFinish          bool
	sessionID          int64
	autoCommit         bool
	isBatch            bool

	watching bool
	watcher  chan<- context.Context
	closech  chan struct{}
	finished chan<- struct{}
	canceled atomicError
	closed   atomicBool
}

func (dmConn *Connection) setTrxFinish(status int32) {
	switch status & Dm_build_721 {
	case Dm_build_718, Dm_build_719, Dm_build_720:
		dmConn.trxFinish = true
	default:
		dmConn.trxFinish = false
	}
}

func (dmConn *Connection) init() {
	if dmConn.dmConnector.stmtPoolMaxSize > 0 {
		dmConn.stmtPool = make([]stmtPoolInfo, 0, dmConn.dmConnector.stmtPoolMaxSize)
	}

	dmConn.stmtMap = make(map[int32]*DmStatement)
	dmConn.DbTimezone = 0
	dmConn.GlobalServerSeries = 0
	dmConn.MaxRowSize = 0
	dmConn.LobEmptyCompOrcl = false
	dmConn.ReadOnly = false
	dmConn.DDLAutoCommit = false
	dmConn.ConstParaOpt = false
	dmConn.IsoLevel = -1
	dmConn.sessionID = -1
	dmConn.Malini2 = true
	dmConn.NewLobFlag = true
	dmConn.Execute2 = true
	dmConn.serverEncoding = ENCODING_GB18030
	dmConn.TrxStatus = Dm_build_669
	dmConn.OracleDateLanguage = byte(Locale)
	dmConn.lastExecInfo = NewExceInfo()
	dmConn.MsgVersion = Dm_build_603

	dmConn.idGenerator = dmConnIDGenerator
}

func (dmConn *Connection) reset() {
	dmConn.DbTimezone = 0
	dmConn.GlobalServerSeries = 0
	dmConn.MaxRowSize = 0
	dmConn.LobEmptyCompOrcl = false
	dmConn.ReadOnly = false
	dmConn.DDLAutoCommit = false
	dmConn.ConstParaOpt = false
	dmConn.IsoLevel = -1
	dmConn.sessionID = -1
	dmConn.Malini2 = true
	dmConn.NewLobFlag = true
	dmConn.Execute2 = true
	dmConn.serverEncoding = ENCODING_GB18030
	dmConn.TrxStatus = Dm_build_669
}

func (dmConn *Connection) checkClosed() error {
	if dmConn.closed.IsSet() {
		return driver.ErrBadConn
	}

	return nil
}

func (dmConn *Connection) executeInner(query string, execType int16) (interface{}, error) {

	stmt, err := NewDmStmt(dmConn, query)

	if err != nil {
		return nil, err
	}

	if execType == Dm_build_686 {
		defer stmt.close()
	}

	stmt.innerUsed = true
	if stmt.dmConn.dmConnector.escapeProcess {
		stmt.nativeSql, err = stmt.dmConn.escape(stmt.nativeSql, stmt.dmConn.dmConnector.keyWords)
		if err != nil {
			stmt.close()
			return nil, err
		}
	}

	var optParamList []OptParameter

	if stmt.dmConn.ConstParaOpt {
		optParamList = make([]OptParameter, 0)
		stmt.nativeSql, optParamList, err = stmt.dmConn.execOpt(stmt.nativeSql, optParamList, stmt.dmConn.getServerEncoding())
		if err != nil {
			stmt.close()
			optParamList = nil
		}
	}

	if execType == Dm_build_685 && dmConn.dmConnector.enRsCache {
		rpv, err := rp.get(stmt, query)
		if err != nil {
			return nil, err
		}

		if rpv != nil {
			stmt.execInfo = rpv.execInfo
			dmConn.lastExecInfo = rpv.execInfo
			return newDmRows(rpv.getResultSet(stmt)), nil
		}
	}

	var info *execRetInfo

	if optParamList != nil && len(optParamList) > 0 {
		info, err = dmConn.Access.Dm_build_411(stmt, optParamList)
		if err != nil {
			stmt.nativeSql = query
			info, err = dmConn.Access.Dm_build_417(stmt, execType)
		}
	} else {
		info, err = dmConn.Access.Dm_build_417(stmt, execType)
	}

	if err != nil {
		stmt.close()
		return nil, err
	}
	dmConn.lastExecInfo = info

	if info.hasResultSet {
		return newDmRows(newInnerRows(0, stmt, info)), nil
	} else {
		return newDmResult(stmt, info), nil
	}
}

func g2dbIsoLevel(isoLevel int32) int32 {
	switch isoLevel {
	case 1:
		return Dm_build_673
	case 2:
		return Dm_build_674
	case 4:
		return Dm_build_675
	case 6:
		return Dm_build_676
	default:
		return -1
	}
}

func (dmConn *Connection) Begin() (driver.Tx, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.begin()
	} else {
		return dmConn.filterChain.reset().DmConnectionBegin(dmConn)
	}
}

func (dmConn *Connection) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.beginTx(ctx, opts)
	}
	return dmConn.filterChain.reset().DmConnectionBeginTx(dmConn, ctx, opts)
}

func (dmConn *Connection) Commit() error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.commit()
	} else {
		return dmConn.filterChain.reset().DmConnectionCommit(dmConn)
	}
}

func (dmConn *Connection) Rollback() error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.rollback()
	} else {
		return dmConn.filterChain.reset().DmConnectionRollback(dmConn)
	}
}

func (dmConn *Connection) Close() error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.close()
	} else {
		return dmConn.filterChain.reset().DmConnectionClose(dmConn)
	}
}

func (dmConn *Connection) Ping(ctx context.Context) error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.ping(ctx)
	} else {
		return dmConn.filterChain.reset().DmConnectionPing(dmConn, ctx)
	}
}

func (dmConn *Connection) Exec(query string, args []driver.Value) (driver.Result, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.exec(query, args)
	}
	return dmConn.filterChain.reset().DmConnectionExec(dmConn, query, args)
}

func (dmConn *Connection) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.execContext(ctx, query, args)
	}
	return dmConn.filterChain.reset().DmConnectionExecContext(dmConn, ctx, query, args)
}

func (dmConn *Connection) Query(query string, args []driver.Value) (driver.Rows, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.query(query, args)
	}
	return dmConn.filterChain.reset().DmConnectionQuery(dmConn, query, args)
}

func (dmConn *Connection) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.queryContext(ctx, query, args)
	}
	return dmConn.filterChain.reset().DmConnectionQueryContext(dmConn, ctx, query, args)
}

func (dmConn *Connection) Prepare(query string) (driver.Stmt, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.prepare(query)
	}
	return dmConn.filterChain.reset().DmConnectionPrepare(dmConn, query)
}

func (dmConn *Connection) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.prepareContext(ctx, query)
	}
	return dmConn.filterChain.reset().DmConnectionPrepareContext(dmConn, ctx, query)
}

func (dmConn *Connection) ResetSession(ctx context.Context) error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.resetSession(ctx)
	}
	return dmConn.filterChain.reset().DmConnectionResetSession(dmConn, ctx)
}

func (dmConn *Connection) CheckNamedValue(nv *driver.NamedValue) error {
	if len(dmConn.filterChain.filters) == 0 {
		return dmConn.checkNamedValue(nv)
	}
	return dmConn.filterChain.reset().DmConnectionCheckNamedValue(dmConn, nv)
}

func (dmConn *Connection) begin() (*Connection, error) {
	return dmConn.beginTx(context.Background(), driver.TxOptions{driver.IsolationLevel(sql.LevelDefault), false})
}

func (dmConn *Connection) beginTx(ctx context.Context, opts driver.TxOptions) (*Connection, error) {
	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if err := dmConn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dmConn.finish()

	if sql.IsolationLevel(opts.Isolation) == sql.LevelDefault {
		opts.Isolation = driver.IsolationLevel(sql.LevelReadCommitted)
	}

	dmConn.ReadOnly = opts.ReadOnly

	if dmConn.IsoLevel == int32(opts.Isolation) {
		return dmConn, nil
	}

	switch sql.IsolationLevel(opts.Isolation) {
	case sql.LevelDefault:
		return dmConn, nil
	case sql.LevelReadUncommitted, sql.LevelReadCommitted, sql.LevelSerializable:
		dmConn.IsoLevel = int32(opts.Isolation)
	case sql.LevelRepeatableRead:
		if dmConn.CompatibleMysql() {
			dmConn.IsoLevel = int32(sql.LevelReadCommitted)
		} else {
			return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
		}
	default:
		return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
	}

	err = dmConn.Access.Dm_build_471(dmConn)
	if err != nil {
		return nil, err
	}
	return dmConn, nil
}

func (dmConn *Connection) commit() error {
	err := dmConn.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		dmConn.autoCommit = dmConn.dmConnector.autoCommit
	}()

	if !dmConn.autoCommit {
		err = dmConn.Access.Commit()
		if err != nil {
			return err
		}
		dmConn.trxFinish = true
		return nil
	} else if !dmConn.dmConnector.alwayseAllowCommit {
		return ECGO_COMMIT_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (dmConn *Connection) rollback() error {
	err := dmConn.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		dmConn.autoCommit = dmConn.dmConnector.autoCommit
	}()

	if !dmConn.autoCommit {
		err = dmConn.Access.Rollback()
		if err != nil {
			return err
		}
		dmConn.trxFinish = true
		return nil
	} else if !dmConn.dmConnector.alwayseAllowCommit {
		return ECGO_ROLLBACK_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (dmConn *Connection) reconnect() error {
	err := dmConn.Access.Close()
	if err != nil {
		return err
	}

	for _, stmt := range dmConn.stmtMap {
		stmt.closed = true
		for id, _ := range stmt.rsMap {
			delete(stmt.rsMap, id)
		}
	}

	if dmConn.stmtPool != nil {
		dmConn.stmtPool = dmConn.stmtPool[:0]
	}

	dmConn.dmConnector.reConnection = dmConn

	if dmConn.dmConnector.group != nil {
		_, err = dmConn.dmConnector.group.connect(dmConn.dmConnector)
		if err != nil {
			return err
		}
	} else {
		_, err = dmConn.dmConnector.connect(context.Background())
	}

	for _, stmt := range dmConn.stmtMap {
		err = dmConn.Access.Dm_build_389(stmt)
		if err != nil {
			return err
		}

		if stmt.paramCount > 0 {
			err = stmt.prepare()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (dmConn *Connection) close() error {
	if dmConn.closed.IsSet() {
		return nil
	}

	close(dmConn.closech)
	if dmConn.Access == nil {
		return nil
	}

	err := dmConn.rollback()
	if err != nil {
		return err
	}

	for _, stmt := range dmConn.stmtMap {
		err = stmt.free()
		if err != nil {
			return err
		}
	}

	if dmConn.stmtPool != nil {
		for _, spi := range dmConn.stmtPool {
			err = dmConn.Access.Dm_build_394(spi.id)
			if err != nil {
				return err
			}
		}
		dmConn.stmtPool = nil
	}

	err = dmConn.Access.Close()
	if err != nil {
		return err
	}

	dmConn.closed.Set(true)
	return nil
}

func (dmConn *Connection) ping(ctx context.Context) error {
	rows, err := dmConn.query("select 1", nil)
	if err != nil {
		return err
	}
	return rows.close()
}

func (dmConn *Connection) exec(query string, args []driver.Value) (*DmResult, error) {
	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dmConn.prepare(query)
		defer stmt.close()
		if err != nil {
			return nil, err
		}
		dmConn.lastExecInfo = stmt.execInfo

		return stmt.exec(args)
	} else {
		r1, err := dmConn.executeInner(query, Dm_build_686)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmResult); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_EXEC_SQL.throw()
		}
	}
}

func (dmConn *Connection) execContext(ctx context.Context, query string, args []driver.NamedValue) (*DmResult, error) {

	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if err := dmConn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dmConn.finish()

	if args != nil && len(args) > 0 {
		stmt, err := dmConn.prepare(query)
		defer stmt.close()
		if err != nil {
			return nil, err
		}
		dmConn.lastExecInfo = stmt.execInfo

		return stmt.execContext(ctx, args)
	} else {
		r1, err := dmConn.executeInner(query, Dm_build_686)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmResult); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_EXEC_SQL.throw()
		}
	}
}

func (dmConn *Connection) query(query string, args []driver.Value) (*DmRows, error) {

	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dmConn.prepare(query)
		if err != nil {
			stmt.close()
			return nil, err
		}
		dmConn.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		return stmt.query(args)

	} else {
		r1, err := dmConn.executeInner(query, Dm_build_685)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmRows); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_QUERY_SQL.throw()
		}
	}
}

func (dmConn *Connection) queryContext(ctx context.Context, query string, args []driver.NamedValue) (*DmRows, error) {

	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if err := dmConn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dmConn.finish()

	if args != nil && len(args) > 0 {
		stmt, err := dmConn.prepare(query)
		if err != nil {
			stmt.close()
			return nil, err
		}
		dmConn.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		return stmt.queryContext(ctx, args)

	} else {
		r1, err := dmConn.executeInner(query, Dm_build_685)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmRows); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_QUERY_SQL.throw()
		}
	}

}

func (dmConn *Connection) prepare(query string) (*DmStatement, error) {
	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	stmt, err := NewDmStmt(dmConn, query)
	if err != nil {
		return nil, err
	}

	err = stmt.prepare()
	return stmt, err
}

func (dmConn *Connection) prepareContext(ctx context.Context, query string) (*DmStatement, error) {
	err := dmConn.checkClosed()
	if err != nil {
		return nil, err
	}

	if err := dmConn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dmConn.finish()

	stmt, err := dmConn.prepare(query)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (dmConn *Connection) resetSession(ctx context.Context) error {
	err := dmConn.checkClosed()
	if err != nil {
		return err
	}

	for _, stmt := range dmConn.stmtMap {
		stmt.inUse = false
	}

	return nil
}

func (dmConn *Connection) checkNamedValue(nv *driver.NamedValue) error {
	var err error
	var cvt = converter{dmConn, false}
	nv.Value, err = cvt.ConvertValue(nv.Value)
	dmConn.isBatch = cvt.isBatch
	return err
}

func (dmConn *Connection) driverQuery(query string) (*DmStatement, *DmRows, error) {
	stmt, err := NewDmStmt(dmConn, query)
	if err != nil {
		return nil, nil, err
	}
	stmt.innerUsed = true
	stmt.innerExec = true
	info, err := dmConn.Access.Dm_build_417(stmt, Dm_build_685)
	if err != nil {
		return nil, nil, err
	}
	dmConn.lastExecInfo = info
	stmt.innerExec = false
	return stmt, newDmRows(newInnerRows(0, stmt, info)), nil
}

func (dmConn *Connection) getIndexOnEPGroup() int32 {
	if dmConn.dmConnector.group == nil || dmConn.dmConnector.group.epList == nil {
		return -1
	}
	for i := 0; i < len(dmConn.dmConnector.group.epList); i++ {
		ep := dmConn.dmConnector.group.epList[i]
		if dmConn.dmConnector.host == ep.host && dmConn.dmConnector.port == ep.port {
			return int32(i)
		}
	}
	return -1
}

func (dmConn *Connection) getServerEncoding() string {
	if dmConn.dmConnector.charCode != "" {
		return dmConn.dmConnector.charCode
	}
	return dmConn.serverEncoding
}

func (dmConn *Connection) lobFetchAll() bool {
	return dmConn.dmConnector.lobMode == 2
}

func (dmConn *Connection) CompatibleOracle() bool {
	return dmConn.dmConnector.compatibleMode == COMPATIBLE_MODE_ORACLE
}

func (dmConn *Connection) CompatibleMysql() bool {
	return dmConn.dmConnector.compatibleMode == COMPATIBLE_MODE_MYSQL
}

func (dmConn *Connection) cancel(err error) {
	dmConn.canceled.Set(err)
	fmt.Println(dmConn.close())
}

func (dmConn *Connection) finish() {
	if !dmConn.watching || dmConn.finished == nil {
		return
	}
	select {
	case dmConn.finished <- struct{}{}:
		dmConn.watching = false
	case <-dmConn.closech:
	}
}

func (dmConn *Connection) startWatcher() {
	watcher := make(chan context.Context, 1)
	dmConn.watcher = watcher
	finished := make(chan struct{})
	dmConn.finished = finished
	go func() {
		for {
			var ctx context.Context
			select {
			case ctx = <-watcher:
			case <-dmConn.closech:
				return
			}

			select {
			case <-ctx.Done():
				dmConn.cancel(ctx.Err())
			case <-finished:
			case <-dmConn.closech:
				return
			}
		}
	}()
}

func (dmConn *Connection) watchCancel(ctx context.Context) error {
	if dmConn.watching {

		return dmConn.close()
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	if ctx.Done() == nil {
		return nil
	}

	if dmConn.watcher == nil {
		return nil
	}

	dmConn.watching = true
	dmConn.watcher <- ctx
	return nil
}

type noCopy struct{}

func (*noCopy) Lock() {}

type atomicBool struct {
	_noCopy noCopy
	value   uint32
}

func (ab *atomicBool) IsSet() bool {
	return atomic.LoadUint32(&ab.value) > 0
}

func (ab *atomicBool) Set(value bool) {
	if value {
		atomic.StoreUint32(&ab.value, 1)
	} else {
		atomic.StoreUint32(&ab.value, 0)
	}
}

func (ab *atomicBool) TrySet(value bool) bool {
	if value {
		return atomic.SwapUint32(&ab.value, 1) == 0
	}
	return atomic.SwapUint32(&ab.value, 0) > 0
}

type atomicError struct {
	_noCopy noCopy
	value   atomic.Value
}

func (ae *atomicError) Set(value error) {
	ae.value.Store(value)
}

func (ae *atomicError) Value() error {
	if v := ae.value.Load(); v != nil {

		return v.(error)
	}
	return nil
}
