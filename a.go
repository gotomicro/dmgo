/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gotomicro/dmgo/security"
)

const (
	Dm_build_334 = 8192
	Dm_build_335 = 2 * time.Second
)

type dm_build_336 struct {
	dm_build_337 net.Conn
	dm_build_338 *tls.Conn
	dm_build_339 *Dm_build_0
	dm_build_340 *DmConnection
	dm_build_341 security.Cipher
	dm_build_342 bool
	dm_build_343 bool
	dm_build_344 *security.DhKey

	dm_build_345 bool
	dm_build_346 string
	dm_build_347 bool
}

func dm_build_348(dm_build_349 context.Context, dm_build_350 *DmConnection) (*dm_build_336, error) {
	var dm_build_351 net.Conn
	var dm_build_352 error

	dialsLock.RLock()
	dm_build_353, dm_build_354 := dials[dm_build_350.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_354 {
		dm_build_351, dm_build_352 = dm_build_353(dm_build_349, dm_build_350.dmConnector.host+":"+strconv.Itoa(int(dm_build_350.dmConnector.port)))
	} else {
		dm_build_351, dm_build_352 = dm_build_356(dm_build_350.dmConnector.host+":"+strconv.Itoa(int(dm_build_350.dmConnector.port)), time.Duration(dm_build_350.dmConnector.socketTimeout)*time.Second)
	}
	if dm_build_352 != nil {
		return nil, dm_build_352
	}

	dm_build_355 := dm_build_336{}
	dm_build_355.dm_build_337 = dm_build_351
	dm_build_355.dm_build_339 = Dm_build_3(Dm_build_621)
	dm_build_355.dm_build_340 = dm_build_350
	dm_build_355.dm_build_342 = false
	dm_build_355.dm_build_343 = false
	dm_build_355.dm_build_345 = false
	dm_build_355.dm_build_346 = ""
	dm_build_355.dm_build_347 = false
	dm_build_350.Access = &dm_build_355

	return &dm_build_355, nil
}

func dm_build_356(dm_build_357 string, dm_build_358 time.Duration) (net.Conn, error) {
	dm_build_359, dm_build_360 := net.DialTimeout("tcp", dm_build_357, dm_build_358)
	if dm_build_360 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_357).throw()
	}

	if tcpConn, ok := dm_build_359.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_335)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_359, nil
}

func (dm_build_362 *dm_build_336) dm_build_361(dm_build_363 dm_build_742) bool {
	var dm_build_364 = dm_build_362.dm_build_340.dmConnector.compress
	if dm_build_363.dm_build_757() == Dm_build_649 || dm_build_364 == Dm_build_698 {
		return false
	}

	if dm_build_364 == Dm_build_696 {
		return true
	} else if dm_build_364 == Dm_build_697 {
		return !dm_build_362.dm_build_340.Local && dm_build_363.dm_build_755() > Dm_build_695
	}

	return false
}

func (dm_build_366 *dm_build_336) dm_build_365(dm_build_367 dm_build_742) bool {
	var dm_build_368 = dm_build_366.dm_build_340.dmConnector.compress
	if dm_build_367.dm_build_757() == Dm_build_649 || dm_build_368 == Dm_build_698 {
		return false
	}

	if dm_build_368 == Dm_build_696 {
		return true
	} else if dm_build_368 == Dm_build_697 {
		return dm_build_366.dm_build_339.Dm_build_267(Dm_build_657) == 1
	}

	return false
}

func (dm_build_370 *dm_build_336) dm_build_369(dm_build_371 dm_build_742) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_373 := dm_build_371.dm_build_755()

	if dm_build_373 > 0 {

		if dm_build_370.dm_build_361(dm_build_371) {
			var retBytes, err = Compress(dm_build_370.dm_build_339, Dm_build_650, int(dm_build_373), int(dm_build_370.dm_build_340.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_370.dm_build_339.Dm_build_14(Dm_build_650)

			dm_build_370.dm_build_339.Dm_build_55(dm_build_373)

			dm_build_370.dm_build_339.Dm_build_83(retBytes)

			dm_build_371.dm_build_756(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_370.dm_build_339.Dm_build_187(Dm_build_657, 1)
		}

		if dm_build_370.dm_build_343 {
			dm_build_373 = dm_build_371.dm_build_755()
			var retBytes = dm_build_370.dm_build_341.Encrypt(dm_build_370.dm_build_339.Dm_build_294(Dm_build_650, int(dm_build_373)), true)

			dm_build_370.dm_build_339.Dm_build_14(Dm_build_650)

			dm_build_370.dm_build_339.Dm_build_83(retBytes)

			dm_build_371.dm_build_756(int32(len(retBytes)))
		}
	}

	if dm_build_370.dm_build_339.Dm_build_12() > Dm_build_622 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_371.dm_build_751()
	if dm_build_370.dm_build_604(dm_build_371) {
		if dm_build_370.dm_build_338 != nil {
			dm_build_370.dm_build_339.Dm_build_17(0)
			if _, err := dm_build_370.dm_build_339.Dm_build_36(dm_build_370.dm_build_338); err != nil {
				return err
			}
		}
	} else {
		dm_build_370.dm_build_339.Dm_build_17(0)
		if _, err := dm_build_370.dm_build_339.Dm_build_36(dm_build_370.dm_build_337); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_375 *dm_build_336) dm_build_374(dm_build_376 dm_build_742) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_378 := int32(0)
	if dm_build_375.dm_build_604(dm_build_376) {
		if dm_build_375.dm_build_338 != nil {
			dm_build_375.dm_build_339.Dm_build_14(0)
			if _, err := dm_build_375.dm_build_339.Dm_build_30(dm_build_375.dm_build_338, Dm_build_650); err != nil {
				return err
			}

			dm_build_378 = dm_build_376.dm_build_755()
			if dm_build_378 > 0 {
				if _, err := dm_build_375.dm_build_339.Dm_build_30(dm_build_375.dm_build_338, int(dm_build_378)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_375.dm_build_339.Dm_build_14(0)
		if _, err := dm_build_375.dm_build_339.Dm_build_30(dm_build_375.dm_build_337, Dm_build_650); err != nil {
			return err
		}
		dm_build_378 = dm_build_376.dm_build_755()

		if dm_build_378 > 0 {
			if _, err := dm_build_375.dm_build_339.Dm_build_30(dm_build_375.dm_build_337, int(dm_build_378)); err != nil {
				return err
			}
		}
	}

	dm_build_376.dm_build_752()

	dm_build_378 = dm_build_376.dm_build_755()
	if dm_build_378 <= 0 {
		return nil
	}

	if dm_build_375.dm_build_343 {
		ebytes := dm_build_375.dm_build_339.Dm_build_294(Dm_build_650, int(dm_build_378))
		bytes, err := dm_build_375.dm_build_341.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_375.dm_build_339.Dm_build_14(Dm_build_650)
		dm_build_375.dm_build_339.Dm_build_83(bytes)
		dm_build_376.dm_build_756(int32(len(bytes)))
	}

	if dm_build_375.dm_build_365(dm_build_376) {

		dm_build_378 = dm_build_376.dm_build_755()
		cbytes := dm_build_375.dm_build_339.Dm_build_294(Dm_build_650+ULINT_SIZE, int(dm_build_378-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_375.dm_build_340.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_375.dm_build_339.Dm_build_14(Dm_build_650)
		dm_build_375.dm_build_339.Dm_build_83(bytes)
		dm_build_376.dm_build_756(int32(len(bytes)))
	}
	return nil
}

func (dm_build_380 *dm_build_336) dm_build_379(dm_build_381 dm_build_742) (dm_build_382 interface{}, dm_build_383 error) {
	if dm_build_380.dm_build_347 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_384 := dm_build_380.dm_build_340
	dm_build_384.mu.Lock()
	defer dm_build_384.mu.Unlock()
	dm_build_383 = dm_build_381.dm_build_746(dm_build_381)
	if dm_build_383 != nil {
		return nil, dm_build_383
	}

	dm_build_383 = dm_build_380.dm_build_369(dm_build_381)
	if dm_build_383 != nil {
		return nil, dm_build_383
	}

	dm_build_383 = dm_build_380.dm_build_374(dm_build_381)
	if dm_build_383 != nil {
		return nil, dm_build_383
	}

	return dm_build_381.dm_build_750(dm_build_381)
}

func (dm_build_386 *dm_build_336) dm_build_385() (*dm_build_1199, error) {

	Dm_build_387 := dm_build_1205(dm_build_386)
	_, dm_build_388 := dm_build_386.dm_build_379(Dm_build_387)
	if dm_build_388 != nil {
		return nil, dm_build_388
	}

	return Dm_build_387, nil
}

func (dm_build_390 *dm_build_336) dm_build_389() error {

	dm_build_391 := dm_build_1066(dm_build_390)
	_, dm_build_392 := dm_build_390.dm_build_379(dm_build_391)
	if dm_build_392 != nil {
		return dm_build_392
	}

	return nil
}

func (dm_build_394 *dm_build_336) dm_build_393() error {

	var dm_build_395 *dm_build_1199
	var err error
	if dm_build_395, err = dm_build_394.dm_build_385(); err != nil {
		return err
	}

	if dm_build_394.dm_build_340.sslEncrypt == 2 {
		if err = dm_build_394.dm_build_600(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_394.dm_build_340.sslEncrypt == 1 {
		if err = dm_build_394.dm_build_600(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_394.dm_build_343 || dm_build_394.dm_build_342 {
		k, err := dm_build_394.dm_build_590()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_395.Dm_build_1203)
		encryptType := dm_build_395.dm_build_1201
		hashType := int(dm_build_395.Dm_build_1202)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_394.dm_build_593(encryptType, sessionKey, dm_build_394.dm_build_340.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_394.dm_build_389(); err != nil {
		return err
	}
	return nil
}

func (dm_build_398 *dm_build_336) Dm_build_397(dm_build_399 *DmStatement) error {
	dm_build_400 := dm_build_1228(dm_build_398, dm_build_399)
	_, dm_build_401 := dm_build_398.dm_build_379(dm_build_400)
	if dm_build_401 != nil {
		return dm_build_401
	}

	return nil
}

func (dm_build_403 *dm_build_336) Dm_build_402(dm_build_404 int32) error {
	dm_build_405 := dm_build_1238(dm_build_403, dm_build_404)
	_, dm_build_406 := dm_build_403.dm_build_379(dm_build_405)
	if dm_build_406 != nil {
		return dm_build_406
	}

	return nil
}

func (dm_build_408 *dm_build_336) Dm_build_407(dm_build_409 *DmStatement, dm_build_410 bool, dm_build_411 int16) (*execRetInfo, error) {
	dm_build_412 := dm_build_1105(dm_build_408, dm_build_409, dm_build_410, dm_build_411)
	dm_build_413, dm_build_414 := dm_build_408.dm_build_379(dm_build_412)
	if dm_build_414 != nil {
		return nil, dm_build_414
	}
	return dm_build_413.(*execRetInfo), nil
}

func (dm_build_416 *dm_build_336) Dm_build_415(dm_build_417 *DmStatement, dm_build_418 int16) (*execRetInfo, error) {
	return dm_build_416.Dm_build_407(dm_build_417, false, Dm_build_702)
}

func (dm_build_420 *dm_build_336) Dm_build_419(dm_build_421 *DmStatement, dm_build_422 []OptParameter) (*execRetInfo, error) {
	dm_build_423, dm_build_424 := dm_build_420.dm_build_379(dm_build_845(dm_build_420, dm_build_421, dm_build_422))
	if dm_build_424 != nil {
		return nil, dm_build_424
	}

	return dm_build_423.(*execRetInfo), nil
}

func (dm_build_426 *dm_build_336) Dm_build_425(dm_build_427 *DmStatement, dm_build_428 int16) (*execRetInfo, error) {
	return dm_build_426.Dm_build_407(dm_build_427, true, dm_build_428)
}

func (dm_build_430 *dm_build_336) Dm_build_429(dm_build_431 *DmStatement, dm_build_432 [][]interface{}) (*execRetInfo, error) {
	dm_build_433 := dm_build_877(dm_build_430, dm_build_431, dm_build_432)
	dm_build_434, dm_build_435 := dm_build_430.dm_build_379(dm_build_433)
	if dm_build_435 != nil {
		return nil, dm_build_435
	}
	return dm_build_434.(*execRetInfo), nil
}

func (dm_build_437 *dm_build_336) Dm_build_436(dm_build_438 *DmStatement, dm_build_439 [][]interface{}, dm_build_440 bool) (*execRetInfo, error) {
	var dm_build_441, dm_build_442 = 0, 0
	var dm_build_443 = len(dm_build_439)
	var dm_build_444 [][]interface{}
	var dm_build_445 = NewExceInfo()
	dm_build_445.updateCounts = make([]int64, dm_build_443)
	var dm_build_446 = false
	for dm_build_441 < dm_build_443 {
		for dm_build_442 = dm_build_441; dm_build_442 < dm_build_443; dm_build_442++ {
			paramData := dm_build_439[dm_build_442]
			bindData := make([]interface{}, dm_build_438.paramCount)
			dm_build_446 = false
			for icol := 0; icol < int(dm_build_438.paramCount); icol++ {
				if dm_build_438.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_437.dm_build_573(bindData, paramData, icol) {
					dm_build_446 = true
					break
				}
			}

			if dm_build_446 {
				break
			}
			dm_build_444 = append(dm_build_444, bindData)
		}

		if dm_build_442 != dm_build_441 {
			tmpExecInfo, err := dm_build_437.Dm_build_429(dm_build_438, dm_build_444)
			if err != nil {
				return nil, err
			}
			dm_build_444 = dm_build_444[0:0]
			dm_build_445.union(tmpExecInfo, dm_build_441, dm_build_442-dm_build_441)
		}

		if dm_build_442 < dm_build_443 {
			tmpExecInfo, err := dm_build_437.Dm_build_447(dm_build_438, dm_build_439[dm_build_442], dm_build_440)
			if err != nil {
				return nil, err
			}

			dm_build_440 = true
			dm_build_445.union(tmpExecInfo, dm_build_442, 1)
		}

		dm_build_441 = dm_build_442 + 1
	}
	for _, i := range dm_build_445.updateCounts {
		if i > 0 {
			dm_build_445.updateCount += i
		}
	}
	return dm_build_445, nil
}

func (dm_build_448 *dm_build_336) Dm_build_447(dm_build_449 *DmStatement, dm_build_450 []interface{}, dm_build_451 bool) (*execRetInfo, error) {

	var dm_build_452 = make([]interface{}, dm_build_449.paramCount)
	for icol := 0; icol < int(dm_build_449.paramCount); icol++ {
		if dm_build_449.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_448.dm_build_573(dm_build_452, dm_build_450, icol) {

			if !dm_build_451 {
				preExecute := dm_build_1094(dm_build_448, dm_build_449, dm_build_449.bindParams)
				dm_build_448.dm_build_379(preExecute)
				dm_build_451 = true
			}

			dm_build_448.dm_build_579(dm_build_449, dm_build_449.bindParams[icol], icol, dm_build_450[icol].(iOffRowBinder))
			dm_build_452[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_453 = make([][]interface{}, 1, 1)
	dm_build_453[0] = dm_build_452

	dm_build_454 := dm_build_877(dm_build_448, dm_build_449, dm_build_453)
	dm_build_455, dm_build_456 := dm_build_448.dm_build_379(dm_build_454)
	if dm_build_456 != nil {
		return nil, dm_build_456
	}
	return dm_build_455.(*execRetInfo), nil
}

func (dm_build_458 *dm_build_336) Dm_build_457(dm_build_459 *DmStatement, dm_build_460 int16) (*execRetInfo, error) {
	dm_build_461 := dm_build_1081(dm_build_458, dm_build_459, dm_build_460)

	dm_build_462, dm_build_463 := dm_build_458.dm_build_379(dm_build_461)
	if dm_build_463 != nil {
		return nil, dm_build_463
	}
	return dm_build_462.(*execRetInfo), nil
}

func (dm_build_465 *dm_build_336) Dm_build_464(dm_build_466 *innerRows, dm_build_467 int64) (*execRetInfo, error) {
	dm_build_468 := dm_build_984(dm_build_465, dm_build_466, dm_build_467, INT64_MAX)
	dm_build_469, dm_build_470 := dm_build_465.dm_build_379(dm_build_468)
	if dm_build_470 != nil {
		return nil, dm_build_470
	}
	return dm_build_469.(*execRetInfo), nil
}

func (dm_build_472 *dm_build_336) Commit() error {
	dm_build_473 := dm_build_830(dm_build_472)
	_, dm_build_474 := dm_build_472.dm_build_379(dm_build_473)
	if dm_build_474 != nil {
		return dm_build_474
	}

	return nil
}

func (dm_build_476 *dm_build_336) Rollback() error {
	dm_build_477 := dm_build_1143(dm_build_476)
	_, dm_build_478 := dm_build_476.dm_build_379(dm_build_477)
	if dm_build_478 != nil {
		return dm_build_478
	}

	return nil
}

func (dm_build_480 *dm_build_336) Dm_build_479(dm_build_481 *DmConnection) error {
	dm_build_482 := dm_build_1148(dm_build_480, dm_build_481.IsoLevel)
	_, dm_build_483 := dm_build_480.dm_build_379(dm_build_482)
	if dm_build_483 != nil {
		return dm_build_483
	}

	return nil
}

func (dm_build_485 *dm_build_336) Dm_build_484(dm_build_486 *DmStatement, dm_build_487 string) error {
	dm_build_488 := dm_build_835(dm_build_485, dm_build_486, dm_build_487)
	_, dm_build_489 := dm_build_485.dm_build_379(dm_build_488)
	if dm_build_489 != nil {
		return dm_build_489
	}

	return nil
}

func (dm_build_491 *dm_build_336) Dm_build_490(dm_build_492 []uint32) ([]int64, error) {
	dm_build_493 := dm_build_1246(dm_build_491, dm_build_492)
	dm_build_494, dm_build_495 := dm_build_491.dm_build_379(dm_build_493)
	if dm_build_495 != nil {
		return nil, dm_build_495
	}
	return dm_build_494.([]int64), nil
}

func (dm_build_497 *dm_build_336) Close() error {
	if dm_build_497.dm_build_347 {
		return nil
	}

	dm_build_498 := dm_build_497.dm_build_337.Close()
	if dm_build_498 != nil {
		return dm_build_498
	}

	dm_build_497.dm_build_340 = nil
	dm_build_497.dm_build_347 = true
	return nil
}

func (dm_build_500 *dm_build_336) dm_build_499(dm_build_501 *lob) (int64, error) {
	dm_build_502 := dm_build_1017(dm_build_500, dm_build_501)
	dm_build_503, dm_build_504 := dm_build_500.dm_build_379(dm_build_502)
	if dm_build_504 != nil {
		return 0, dm_build_504
	}
	return dm_build_503.(int64), nil
}

func (dm_build_506 *dm_build_336) dm_build_505(dm_build_507 *lob, dm_build_508 int32, dm_build_509 int32) (*lobRetInfo, error) {
	dm_build_510 := dm_build_1002(dm_build_506, dm_build_507, int(dm_build_508), int(dm_build_509))
	dm_build_511, dm_build_512 := dm_build_506.dm_build_379(dm_build_510)
	if dm_build_512 != nil {
		return nil, dm_build_512
	}
	return dm_build_511.(*lobRetInfo), nil
}

func (dm_build_514 *dm_build_336) dm_build_513(dm_build_515 *DmBlob, dm_build_516 int32, dm_build_517 int32) ([]byte, error) {
	var dm_build_518 = make([]byte, dm_build_517)
	var dm_build_519 int32 = 0
	var dm_build_520 int32 = 0
	var dm_build_521 *lobRetInfo
	var dm_build_522 []byte
	var dm_build_523 error
	for dm_build_519 < dm_build_517 {
		dm_build_520 = dm_build_517 - dm_build_519
		if dm_build_520 > Dm_build_735 {
			dm_build_520 = Dm_build_735
		}
		dm_build_521, dm_build_523 = dm_build_514.dm_build_505(&dm_build_515.lob, dm_build_516+dm_build_519, dm_build_520)
		if dm_build_523 != nil {
			return nil, dm_build_523
		}
		dm_build_522 = dm_build_521.data
		if dm_build_522 == nil || len(dm_build_522) == 0 {
			break
		}
		Dm_build_1257.Dm_build_1313(dm_build_518, int(dm_build_519), dm_build_522, 0, len(dm_build_522))
		dm_build_519 += int32(len(dm_build_522))
		if dm_build_515.readOver {
			break
		}
	}
	return dm_build_518, nil
}

func (dm_build_525 *dm_build_336) dm_build_524(dm_build_526 *DmClob, dm_build_527 int32, dm_build_528 int32) (string, error) {
	var dm_build_529 bytes.Buffer
	var dm_build_530 int32 = 0
	var dm_build_531 int32 = 0
	var dm_build_532 *lobRetInfo
	var dm_build_533 []byte
	var dm_build_534 string
	var dm_build_535 error
	for dm_build_530 < dm_build_528 {
		dm_build_531 = dm_build_528 - dm_build_530
		if dm_build_531 > Dm_build_735/2 {
			dm_build_531 = Dm_build_735 / 2
		}
		dm_build_532, dm_build_535 = dm_build_525.dm_build_505(&dm_build_526.lob, dm_build_527+dm_build_530, dm_build_531)
		if dm_build_535 != nil {
			return "", dm_build_535
		}
		dm_build_533 = dm_build_532.data
		if dm_build_533 == nil || len(dm_build_533) == 0 {
			break
		}
		dm_build_534 = Dm_build_1257.Dm_build_1414(dm_build_533, 0, len(dm_build_533), dm_build_526.serverEncoding, dm_build_525.dm_build_340)

		dm_build_529.WriteString(dm_build_534)
		var strLen = dm_build_532.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_534))
		}
		dm_build_530 += int32(strLen)
		if dm_build_526.readOver {
			break
		}
	}
	return dm_build_529.String(), nil
}

func (dm_build_537 *dm_build_336) dm_build_536(dm_build_538 *DmClob, dm_build_539 int, dm_build_540 string, dm_build_541 string) (int, error) {
	var dm_build_542 = Dm_build_1257.Dm_build_1473(dm_build_540, dm_build_541, dm_build_537.dm_build_340)
	var dm_build_543 = 0
	var dm_build_544 = len(dm_build_542)
	var dm_build_545 = 0
	var dm_build_546 = 0
	var dm_build_547 = 0
	var dm_build_548 = dm_build_544/Dm_build_734 + 1
	var dm_build_549 byte = 0
	var dm_build_550 byte = 0x01
	var dm_build_551 byte = 0x02
	for i := 0; i < dm_build_548; i++ {
		dm_build_549 = 0
		if i == 0 {
			dm_build_549 |= dm_build_550
		}
		if i == dm_build_548-1 {
			dm_build_549 |= dm_build_551
		}
		dm_build_547 = dm_build_544 - dm_build_546
		if dm_build_547 > Dm_build_734 {
			dm_build_547 = Dm_build_734
		}

		setLobData := dm_build_1162(dm_build_537, &dm_build_538.lob, dm_build_549, dm_build_539, dm_build_542, dm_build_543, dm_build_547)
		ret, err := dm_build_537.dm_build_379(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_545, nil
		} else {
			dm_build_539 += int(tmp)
			dm_build_545 += int(tmp)
			dm_build_546 += dm_build_547
			dm_build_543 += dm_build_547
		}
	}
	return dm_build_545, nil
}

func (dm_build_553 *dm_build_336) dm_build_552(dm_build_554 *DmBlob, dm_build_555 int, dm_build_556 []byte) (int, error) {
	var dm_build_557 = 0
	var dm_build_558 = len(dm_build_556)
	var dm_build_559 = 0
	var dm_build_560 = 0
	var dm_build_561 = 0
	var dm_build_562 = dm_build_558/Dm_build_734 + 1
	var dm_build_563 byte = 0
	var dm_build_564 byte = 0x01
	var dm_build_565 byte = 0x02
	for i := 0; i < dm_build_562; i++ {
		dm_build_563 = 0
		if i == 0 {
			dm_build_563 |= dm_build_564
		}
		if i == dm_build_562-1 {
			dm_build_563 |= dm_build_565
		}
		dm_build_561 = dm_build_558 - dm_build_560
		if dm_build_561 > Dm_build_734 {
			dm_build_561 = Dm_build_734
		}

		setLobData := dm_build_1162(dm_build_553, &dm_build_554.lob, dm_build_563, dm_build_555, dm_build_556, dm_build_557, dm_build_561)
		ret, err := dm_build_553.dm_build_379(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_559, nil
		} else {
			dm_build_555 += int(tmp)
			dm_build_559 += int(tmp)
			dm_build_560 += dm_build_561
			dm_build_557 += dm_build_561
		}
	}
	return dm_build_559, nil
}

func (dm_build_567 *dm_build_336) dm_build_566(dm_build_568 *lob, dm_build_569 int) (int64, error) {
	dm_build_570 := dm_build_1028(dm_build_567, dm_build_568, dm_build_569)
	dm_build_571, dm_build_572 := dm_build_567.dm_build_379(dm_build_570)
	if dm_build_572 != nil {
		return dm_build_568.length, dm_build_572
	}
	return dm_build_571.(int64), nil
}

func (dm_build_574 *dm_build_336) dm_build_573(dm_build_575 []interface{}, dm_build_576 []interface{}, dm_build_577 int) bool {
	var dm_build_578 = false
	dm_build_575[dm_build_577] = dm_build_576[dm_build_577]

	if binder, ok := dm_build_576[dm_build_577].(iOffRowBinder); ok {
		dm_build_578 = true
		dm_build_575[dm_build_577] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_574.dm_build_340) {
			dm_build_575[dm_build_577] = &lobCtl{lob.buildCtlData()}
			dm_build_578 = false
		}
	} else {
		dm_build_575[dm_build_577] = dm_build_576[dm_build_577]
	}
	return dm_build_578
}

func (dm_build_580 *dm_build_336) dm_build_579(dm_build_581 *DmStatement, dm_build_582 parameter, dm_build_583 int, dm_build_584 iOffRowBinder) error {
	var dm_build_585 = Dm_build_1542()
	dm_build_584.read(dm_build_585)
	var dm_build_586 = 0
	for !dm_build_584.isReadOver() || dm_build_585.Dm_build_1543() > 0 {
		if !dm_build_584.isReadOver() && dm_build_585.Dm_build_1543() < Dm_build_734 {
			dm_build_584.read(dm_build_585)
		}
		if dm_build_585.Dm_build_1543() > Dm_build_734 {
			dm_build_586 = Dm_build_734
		} else {
			dm_build_586 = dm_build_585.Dm_build_1543()
		}

		putData := dm_build_1133(dm_build_580, dm_build_581, int16(dm_build_583), dm_build_585, int32(dm_build_586))
		_, err := dm_build_580.dm_build_379(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_588 *dm_build_336) dm_build_587() ([]byte, error) {
	var dm_build_589 error
	if dm_build_588.dm_build_344 == nil {
		if dm_build_588.dm_build_344, dm_build_589 = security.NewClientKeyPair(); dm_build_589 != nil {
			return nil, dm_build_589
		}
	}
	return security.Bn2Bytes(dm_build_588.dm_build_344.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_591 *dm_build_336) dm_build_590() (*security.DhKey, error) {
	var dm_build_592 error
	if dm_build_591.dm_build_344 == nil {
		if dm_build_591.dm_build_344, dm_build_592 = security.NewClientKeyPair(); dm_build_592 != nil {
			return nil, dm_build_592
		}
	}
	return dm_build_591.dm_build_344, nil
}

func (dm_build_594 *dm_build_336) dm_build_593(dm_build_595 int, dm_build_596 []byte, dm_build_597 string, dm_build_598 int) (dm_build_599 error) {
	if dm_build_595 > 0 && dm_build_595 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_596 != nil {
		dm_build_594.dm_build_341, dm_build_599 = security.NewSymmCipher(dm_build_595, dm_build_596)
	} else if dm_build_595 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_594.dm_build_341, dm_build_599 = security.NewThirdPartCipher(dm_build_595, dm_build_596, dm_build_597, dm_build_598); dm_build_599 != nil {
			dm_build_599 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_599.Error()).throw()
		}
	}
	return
}

func (dm_build_601 *dm_build_336) dm_build_600(dm_build_602 bool) (dm_build_603 error) {
	if dm_build_601.dm_build_338, dm_build_603 = security.NewTLSFromTCP(dm_build_601.dm_build_337, dm_build_601.dm_build_340.dmConnector.sslCertPath, dm_build_601.dm_build_340.dmConnector.sslKeyPath, dm_build_601.dm_build_340.dmConnector.user); dm_build_603 != nil {
		return
	}
	if !dm_build_602 {
		dm_build_601.dm_build_338 = nil
	}
	return
}

func (dm_build_605 *dm_build_336) dm_build_604(dm_build_606 dm_build_742) bool {
	return dm_build_606.dm_build_757() != Dm_build_649 && dm_build_605.dm_build_340.sslEncrypt == 1
}
