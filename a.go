/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"gitee.com/chunanyong/dm/security"
)

const (
	Dm_build_1343 = 8192
	Dm_build_1344 = 2 * time.Second
)

type dm_build_1345 struct {
	dm_build_1346 *net.TCPConn
	dm_build_1347 *tls.Conn
	dm_build_1348 *Dm_build_1009
	dm_build_1349 *DmConnection
	dm_build_1350 security.Cipher
	dm_build_1351 bool
	dm_build_1352 bool
	dm_build_1353 *security.DhKey

	dm_build_1354 bool
	dm_build_1355 string
	dm_build_1356 bool
}

func dm_build_1357(dm_build_1358 *DmConnection) (*dm_build_1345, error) {
	dm_build_1359, dm_build_1360 := dm_build_1362(dm_build_1358.dmConnector.host+":"+strconv.Itoa(int(dm_build_1358.dmConnector.port)), time.Duration(dm_build_1358.dmConnector.socketTimeout)*time.Second)
	if dm_build_1360 != nil {
		return nil, dm_build_1360
	}

	dm_build_1361 := dm_build_1345{}
	dm_build_1361.dm_build_1346 = dm_build_1359
	dm_build_1361.dm_build_1348 = Dm_build_1012(Dm_build_14)
	dm_build_1361.dm_build_1349 = dm_build_1358
	dm_build_1361.dm_build_1351 = false
	dm_build_1361.dm_build_1352 = false
	dm_build_1361.dm_build_1354 = false
	dm_build_1361.dm_build_1355 = ""
	dm_build_1361.dm_build_1356 = false
	dm_build_1358.Access = &dm_build_1361

	return &dm_build_1361, nil
}

func dm_build_1362(dm_build_1363 string, dm_build_1364 time.Duration) (*net.TCPConn, error) {
	dm_build_1365, dm_build_1366 := net.DialTimeout("tcp", dm_build_1363, dm_build_1364)
	if dm_build_1366 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_1363).throw()
	}

	if tcpConn, ok := dm_build_1365.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1344)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_1368 *dm_build_1345) dm_build_1367(dm_build_1369 dm_build_135) bool {
	var dm_build_1370 = dm_build_1368.dm_build_1349.dmConnector.compress
	if dm_build_1369.dm_build_150() == Dm_build_42 || dm_build_1370 == Dm_build_91 {
		return false
	}

	if dm_build_1370 == Dm_build_89 {
		return true
	} else if dm_build_1370 == Dm_build_90 {
		return !dm_build_1368.dm_build_1349.Local && dm_build_1369.dm_build_148() > Dm_build_88
	}

	return false
}

func (dm_build_1372 *dm_build_1345) dm_build_1371(dm_build_1373 dm_build_135) bool {
	var dm_build_1374 = dm_build_1372.dm_build_1349.dmConnector.compress
	if dm_build_1373.dm_build_150() == Dm_build_42 || dm_build_1374 == Dm_build_91 {
		return false
	}

	if dm_build_1374 == Dm_build_89 {
		return true
	} else if dm_build_1374 == Dm_build_90 {
		return dm_build_1372.dm_build_1348.Dm_build_1276(Dm_build_50) == 1
	}

	return false
}

func (dm_build_1376 *dm_build_1345) dm_build_1375(dm_build_1377 dm_build_135) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_1379 := dm_build_1377.dm_build_148()

	if dm_build_1379 > 0 {

		if dm_build_1376.dm_build_1367(dm_build_1377) {
			var retBytes, err = Compress(dm_build_1376.dm_build_1348, Dm_build_43, int(dm_build_1379), int(dm_build_1376.dm_build_1349.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_1376.dm_build_1348.Dm_build_1023(Dm_build_43)

			dm_build_1376.dm_build_1348.Dm_build_1064(dm_build_1379)

			dm_build_1376.dm_build_1348.Dm_build_1092(retBytes)

			dm_build_1377.dm_build_149(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_1376.dm_build_1348.Dm_build_1196(Dm_build_50, 1)
		}

		if dm_build_1376.dm_build_1352 {
			dm_build_1379 = dm_build_1377.dm_build_148()
			var retBytes = dm_build_1376.dm_build_1350.Encrypt(dm_build_1376.dm_build_1348.Dm_build_1303(Dm_build_43, int(dm_build_1379)), true)

			dm_build_1376.dm_build_1348.Dm_build_1023(Dm_build_43)

			dm_build_1376.dm_build_1348.Dm_build_1092(retBytes)

			dm_build_1377.dm_build_149(int32(len(retBytes)))
		}
	}

	if dm_build_1376.dm_build_1348.Dm_build_1021() > Dm_build_15 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_1377.dm_build_144()
	if dm_build_1376.dm_build_1610(dm_build_1377) {
		if dm_build_1376.dm_build_1347 != nil {
			dm_build_1376.dm_build_1348.Dm_build_1026(0)
			if _, err := dm_build_1376.dm_build_1348.Dm_build_1045(dm_build_1376.dm_build_1347); err != nil {
				return err
			}
		}
	} else {
		dm_build_1376.dm_build_1348.Dm_build_1026(0)
		if _, err := dm_build_1376.dm_build_1348.Dm_build_1045(dm_build_1376.dm_build_1346); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1381 *dm_build_1345) dm_build_1380(dm_build_1382 dm_build_135) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_1384 := int32(0)
	if dm_build_1381.dm_build_1610(dm_build_1382) {
		if dm_build_1381.dm_build_1347 != nil {
			dm_build_1381.dm_build_1348.Dm_build_1023(0)
			if _, err := dm_build_1381.dm_build_1348.Dm_build_1039(dm_build_1381.dm_build_1347, Dm_build_43); err != nil {
				return err
			}

			dm_build_1384 = dm_build_1382.dm_build_148()
			if dm_build_1384 > 0 {
				if _, err := dm_build_1381.dm_build_1348.Dm_build_1039(dm_build_1381.dm_build_1347, int(dm_build_1384)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_1381.dm_build_1348.Dm_build_1023(0)
		if _, err := dm_build_1381.dm_build_1348.Dm_build_1039(dm_build_1381.dm_build_1346, Dm_build_43); err != nil {
			return err
		}
		dm_build_1384 = dm_build_1382.dm_build_148()

		if dm_build_1384 > 0 {
			if _, err := dm_build_1381.dm_build_1348.Dm_build_1039(dm_build_1381.dm_build_1346, int(dm_build_1384)); err != nil {
				return err
			}
		}
	}

	dm_build_1382.dm_build_145()

	dm_build_1384 = dm_build_1382.dm_build_148()
	if dm_build_1384 <= 0 {
		return nil
	}

	if dm_build_1381.dm_build_1352 {
		ebytes := dm_build_1381.dm_build_1348.Dm_build_1303(Dm_build_43, int(dm_build_1384))
		bytes, err := dm_build_1381.dm_build_1350.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_1381.dm_build_1348.Dm_build_1023(Dm_build_43)
		dm_build_1381.dm_build_1348.Dm_build_1092(bytes)
		dm_build_1382.dm_build_149(int32(len(bytes)))
	}

	if dm_build_1381.dm_build_1371(dm_build_1382) {

		dm_build_1384 = dm_build_1382.dm_build_148()
		cbytes := dm_build_1381.dm_build_1348.Dm_build_1303(Dm_build_43+ULINT_SIZE, int(dm_build_1384-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_1381.dm_build_1349.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_1381.dm_build_1348.Dm_build_1023(Dm_build_43)
		dm_build_1381.dm_build_1348.Dm_build_1092(bytes)
		dm_build_1382.dm_build_149(int32(len(bytes)))
	}
	return nil
}

func (dm_build_1386 *dm_build_1345) dm_build_1385(dm_build_1387 dm_build_135) (dm_build_1388 interface{}, dm_build_1389 error) {
	if dm_build_1386.dm_build_1356 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_1390 := dm_build_1386.dm_build_1349
	dm_build_1390.mu.Lock()
	defer dm_build_1390.mu.Unlock()
	dm_build_1389 = dm_build_1387.dm_build_139(dm_build_1387)
	if dm_build_1389 != nil {
		return nil, dm_build_1389
	}

	dm_build_1389 = dm_build_1386.dm_build_1375(dm_build_1387)
	if dm_build_1389 != nil {
		return nil, dm_build_1389
	}

	dm_build_1389 = dm_build_1386.dm_build_1380(dm_build_1387)
	if dm_build_1389 != nil {
		return nil, dm_build_1389
	}

	return dm_build_1387.dm_build_143(dm_build_1387)
}

func (dm_build_1392 *dm_build_1345) dm_build_1391() (*dm_build_591, error) {

	Dm_build_1393 := dm_build_597(dm_build_1392)
	_, dm_build_1394 := dm_build_1392.dm_build_1385(Dm_build_1393)
	if dm_build_1394 != nil {
		return nil, dm_build_1394
	}

	return Dm_build_1393, nil
}

func (dm_build_1396 *dm_build_1345) dm_build_1395() error {

	dm_build_1397 := dm_build_458(dm_build_1396)
	_, dm_build_1398 := dm_build_1396.dm_build_1385(dm_build_1397)
	if dm_build_1398 != nil {
		return dm_build_1398
	}

	return nil
}

func (dm_build_1400 *dm_build_1345) dm_build_1399() error {

	var dm_build_1401 *dm_build_591
	var err error
	if dm_build_1401, err = dm_build_1400.dm_build_1391(); err != nil {
		return err
	}

	if dm_build_1400.dm_build_1349.sslEncrypt == 2 {
		if err = dm_build_1400.dm_build_1606(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_1400.dm_build_1349.sslEncrypt == 1 {
		if err = dm_build_1400.dm_build_1606(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_1400.dm_build_1352 || dm_build_1400.dm_build_1351 {
		k, err := dm_build_1400.dm_build_1596()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_1401.Dm_build_595)
		encryptType := dm_build_1401.dm_build_593
		hashType := int(dm_build_1401.Dm_build_594)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_1400.dm_build_1599(encryptType, sessionKey, dm_build_1400.dm_build_1349.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_1400.dm_build_1395(); err != nil {
		return err
	}
	return nil
}

func (dm_build_1404 *dm_build_1345) Dm_build_1403(dm_build_1405 *DmStatement) error {
	dm_build_1406 := dm_build_620(dm_build_1404, dm_build_1405)
	_, dm_build_1407 := dm_build_1404.dm_build_1385(dm_build_1406)
	if dm_build_1407 != nil {
		return dm_build_1407
	}

	return nil
}

func (dm_build_1409 *dm_build_1345) Dm_build_1408(dm_build_1410 int32) error {
	dm_build_1411 := dm_build_630(dm_build_1409, dm_build_1410)
	_, dm_build_1412 := dm_build_1409.dm_build_1385(dm_build_1411)
	if dm_build_1412 != nil {
		return dm_build_1412
	}

	return nil
}

func (dm_build_1414 *dm_build_1345) Dm_build_1413(dm_build_1415 *DmStatement, dm_build_1416 bool, dm_build_1417 int16) (*execRetInfo, error) {
	dm_build_1418 := dm_build_497(dm_build_1414, dm_build_1415, dm_build_1416, dm_build_1417)
	dm_build_1419, dm_build_1420 := dm_build_1414.dm_build_1385(dm_build_1418)
	if dm_build_1420 != nil {
		return nil, dm_build_1420
	}
	return dm_build_1419.(*execRetInfo), nil
}

func (dm_build_1422 *dm_build_1345) Dm_build_1421(dm_build_1423 *DmStatement, dm_build_1424 int16) (*execRetInfo, error) {
	return dm_build_1422.Dm_build_1413(dm_build_1423, false, Dm_build_95)
}

func (dm_build_1426 *dm_build_1345) Dm_build_1425(dm_build_1427 *DmStatement, dm_build_1428 []OptParameter) (*execRetInfo, error) {
	dm_build_1429, dm_build_1430 := dm_build_1426.dm_build_1385(dm_build_238(dm_build_1426, dm_build_1427, dm_build_1428))
	if dm_build_1430 != nil {
		return nil, dm_build_1430
	}

	return dm_build_1429.(*execRetInfo), nil
}

func (dm_build_1432 *dm_build_1345) Dm_build_1431(dm_build_1433 *DmStatement, dm_build_1434 int16) (*execRetInfo, error) {
	return dm_build_1432.Dm_build_1413(dm_build_1433, true, dm_build_1434)
}

func (dm_build_1436 *dm_build_1345) Dm_build_1435(dm_build_1437 *DmStatement, dm_build_1438 [][]interface{}) (*execRetInfo, error) {
	dm_build_1439 := dm_build_270(dm_build_1436, dm_build_1437, dm_build_1438)
	dm_build_1440, dm_build_1441 := dm_build_1436.dm_build_1385(dm_build_1439)
	if dm_build_1441 != nil {
		return nil, dm_build_1441
	}
	return dm_build_1440.(*execRetInfo), nil
}

func (dm_build_1443 *dm_build_1345) Dm_build_1442(dm_build_1444 *DmStatement, dm_build_1445 [][]interface{}, dm_build_1446 bool) (*execRetInfo, error) {
	var dm_build_1447, dm_build_1448 = 0, 0
	var dm_build_1449 = len(dm_build_1445)
	var dm_build_1450 [][]interface{}
	var dm_build_1451 = NewExceInfo()
	dm_build_1451.updateCounts = make([]int64, dm_build_1449)
	var dm_build_1452 = false
	for dm_build_1447 < dm_build_1449 {
		for dm_build_1448 = dm_build_1447; dm_build_1448 < dm_build_1449; dm_build_1448++ {
			paramData := dm_build_1445[dm_build_1448]
			bindData := make([]interface{}, dm_build_1444.paramCount)
			dm_build_1452 = false
			for icol := 0; icol < int(dm_build_1444.paramCount); icol++ {
				if dm_build_1444.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_1443.dm_build_1579(bindData, paramData, icol) {
					dm_build_1452 = true
					break
				}
			}

			if dm_build_1452 {
				break
			}
			dm_build_1450 = append(dm_build_1450, bindData)
		}

		if dm_build_1448 != dm_build_1447 {
			tmpExecInfo, err := dm_build_1443.Dm_build_1435(dm_build_1444, dm_build_1450)
			if err != nil {
				return nil, err
			}
			dm_build_1450 = dm_build_1450[0:0]
			dm_build_1451.union(tmpExecInfo, dm_build_1447, dm_build_1448-dm_build_1447)
		}

		if dm_build_1448 < dm_build_1449 {
			tmpExecInfo, err := dm_build_1443.Dm_build_1453(dm_build_1444, dm_build_1445[dm_build_1448], dm_build_1446)
			if err != nil {
				return nil, err
			}

			dm_build_1446 = true
			dm_build_1451.union(tmpExecInfo, dm_build_1448, 1)
		}

		dm_build_1447 = dm_build_1448 + 1
	}
	for _, i := range dm_build_1451.updateCounts {
		if i > 0 {
			dm_build_1451.updateCount += i
		}
	}
	return dm_build_1451, nil
}

func (dm_build_1454 *dm_build_1345) Dm_build_1453(dm_build_1455 *DmStatement, dm_build_1456 []interface{}, dm_build_1457 bool) (*execRetInfo, error) {

	var dm_build_1458 = make([]interface{}, dm_build_1455.paramCount)
	for icol := 0; icol < int(dm_build_1455.paramCount); icol++ {
		if dm_build_1455.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_1454.dm_build_1579(dm_build_1458, dm_build_1456, icol) {

			if !dm_build_1457 {
				preExecute := dm_build_486(dm_build_1454, dm_build_1455, dm_build_1455.bindParams)
				dm_build_1454.dm_build_1385(preExecute)
				dm_build_1457 = true
			}

			dm_build_1454.dm_build_1585(dm_build_1455, dm_build_1455.bindParams[icol], icol, dm_build_1456[icol].(iOffRowBinder))
			dm_build_1458[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_1459 = make([][]interface{}, 1, 1)
	dm_build_1459[0] = dm_build_1458

	dm_build_1460 := dm_build_270(dm_build_1454, dm_build_1455, dm_build_1459)
	dm_build_1461, dm_build_1462 := dm_build_1454.dm_build_1385(dm_build_1460)
	if dm_build_1462 != nil {
		return nil, dm_build_1462
	}
	return dm_build_1461.(*execRetInfo), nil
}

func (dm_build_1464 *dm_build_1345) Dm_build_1463(dm_build_1465 *DmStatement, dm_build_1466 int16) (*execRetInfo, error) {
	dm_build_1467 := dm_build_473(dm_build_1464, dm_build_1465, dm_build_1466)

	dm_build_1468, dm_build_1469 := dm_build_1464.dm_build_1385(dm_build_1467)
	if dm_build_1469 != nil {
		return nil, dm_build_1469
	}
	return dm_build_1468.(*execRetInfo), nil
}

func (dm_build_1471 *dm_build_1345) Dm_build_1470(dm_build_1472 *innerRows, dm_build_1473 int64) (*execRetInfo, error) {
	dm_build_1474 := dm_build_376(dm_build_1471, dm_build_1472, dm_build_1473, INT64_MAX)
	dm_build_1475, dm_build_1476 := dm_build_1471.dm_build_1385(dm_build_1474)
	if dm_build_1476 != nil {
		return nil, dm_build_1476
	}
	return dm_build_1475.(*execRetInfo), nil
}

func (dm_build_1478 *dm_build_1345) Commit() error {
	dm_build_1479 := dm_build_223(dm_build_1478)
	_, dm_build_1480 := dm_build_1478.dm_build_1385(dm_build_1479)
	if dm_build_1480 != nil {
		return dm_build_1480
	}

	return nil
}

func (dm_build_1482 *dm_build_1345) Rollback() error {
	dm_build_1483 := dm_build_535(dm_build_1482)
	_, dm_build_1484 := dm_build_1482.dm_build_1385(dm_build_1483)
	if dm_build_1484 != nil {
		return dm_build_1484
	}

	return nil
}

func (dm_build_1486 *dm_build_1345) Dm_build_1485(dm_build_1487 *DmConnection) error {
	dm_build_1488 := dm_build_540(dm_build_1486, dm_build_1487.IsoLevel)
	_, dm_build_1489 := dm_build_1486.dm_build_1385(dm_build_1488)
	if dm_build_1489 != nil {
		return dm_build_1489
	}

	return nil
}

func (dm_build_1491 *dm_build_1345) Dm_build_1490(dm_build_1492 *DmStatement, dm_build_1493 string) error {
	dm_build_1494 := dm_build_228(dm_build_1491, dm_build_1492, dm_build_1493)
	_, dm_build_1495 := dm_build_1491.dm_build_1385(dm_build_1494)
	if dm_build_1495 != nil {
		return dm_build_1495
	}

	return nil
}

func (dm_build_1497 *dm_build_1345) Dm_build_1496(dm_build_1498 []uint32) ([]int64, error) {
	dm_build_1499 := dm_build_638(dm_build_1497, dm_build_1498)
	dm_build_1500, dm_build_1501 := dm_build_1497.dm_build_1385(dm_build_1499)
	if dm_build_1501 != nil {
		return nil, dm_build_1501
	}
	return dm_build_1500.([]int64), nil
}

func (dm_build_1503 *dm_build_1345) Close() error {
	if dm_build_1503.dm_build_1356 {
		return nil
	}

	dm_build_1504 := dm_build_1503.dm_build_1346.Close()
	if dm_build_1504 != nil {
		return dm_build_1504
	}

	dm_build_1503.dm_build_1349 = nil
	dm_build_1503.dm_build_1356 = true
	return nil
}

func (dm_build_1506 *dm_build_1345) dm_build_1505(dm_build_1507 *lob) (int64, error) {
	dm_build_1508 := dm_build_409(dm_build_1506, dm_build_1507)
	dm_build_1509, dm_build_1510 := dm_build_1506.dm_build_1385(dm_build_1508)
	if dm_build_1510 != nil {
		return 0, dm_build_1510
	}
	return dm_build_1509.(int64), nil
}

func (dm_build_1512 *dm_build_1345) dm_build_1511(dm_build_1513 *lob, dm_build_1514 int32, dm_build_1515 int32) (*lobRetInfo, error) {
	dm_build_1516 := dm_build_394(dm_build_1512, dm_build_1513, int(dm_build_1514), int(dm_build_1515))
	dm_build_1517, dm_build_1518 := dm_build_1512.dm_build_1385(dm_build_1516)
	if dm_build_1518 != nil {
		return nil, dm_build_1518
	}
	return dm_build_1517.(*lobRetInfo), nil
}

func (dm_build_1520 *dm_build_1345) dm_build_1519(dm_build_1521 *DmBlob, dm_build_1522 int32, dm_build_1523 int32) ([]byte, error) {
	var dm_build_1524 = make([]byte, dm_build_1523)
	var dm_build_1525 int32 = 0
	var dm_build_1526 int32 = 0
	var dm_build_1527 *lobRetInfo
	var dm_build_1528 []byte
	var dm_build_1529 error
	for dm_build_1525 < dm_build_1523 {
		dm_build_1526 = dm_build_1523 - dm_build_1525
		if dm_build_1526 > Dm_build_128 {
			dm_build_1526 = Dm_build_128
		}
		dm_build_1527, dm_build_1529 = dm_build_1520.dm_build_1511(&dm_build_1521.lob, dm_build_1522+dm_build_1525, dm_build_1526)
		if dm_build_1529 != nil {
			return nil, dm_build_1529
		}
		dm_build_1528 = dm_build_1527.data
		if dm_build_1528 == nil || len(dm_build_1528) == 0 {
			break
		}
		Dm_build_649.Dm_build_705(dm_build_1524, int(dm_build_1525), dm_build_1528, 0, len(dm_build_1528))
		dm_build_1525 += int32(len(dm_build_1528))
		if dm_build_1521.readOver {
			break
		}
	}
	return dm_build_1524, nil
}

func (dm_build_1531 *dm_build_1345) dm_build_1530(dm_build_1532 *DmClob, dm_build_1533 int32, dm_build_1534 int32) (string, error) {
	var dm_build_1535 bytes.Buffer
	var dm_build_1536 int32 = 0
	var dm_build_1537 int32 = 0
	var dm_build_1538 *lobRetInfo
	var dm_build_1539 []byte
	var dm_build_1540 string
	var dm_build_1541 error
	for dm_build_1536 < dm_build_1534 {
		dm_build_1537 = dm_build_1534 - dm_build_1536
		if dm_build_1537 > Dm_build_128/2 {
			dm_build_1537 = Dm_build_128 / 2
		}
		dm_build_1538, dm_build_1541 = dm_build_1531.dm_build_1511(&dm_build_1532.lob, dm_build_1533+dm_build_1536, dm_build_1537)
		if dm_build_1541 != nil {
			return "", dm_build_1541
		}
		dm_build_1539 = dm_build_1538.data
		if dm_build_1539 == nil || len(dm_build_1539) == 0 {
			break
		}
		dm_build_1540 = Dm_build_649.Dm_build_806(dm_build_1539, 0, len(dm_build_1539), dm_build_1532.serverEncoding, dm_build_1531.dm_build_1349)

		dm_build_1535.WriteString(dm_build_1540)
		var strLen = dm_build_1538.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_1540))
		}
		dm_build_1536 += int32(strLen)
		if dm_build_1532.readOver {
			break
		}
	}
	return dm_build_1535.String(), nil
}

func (dm_build_1543 *dm_build_1345) dm_build_1542(dm_build_1544 *DmClob, dm_build_1545 int, dm_build_1546 string, dm_build_1547 string) (int, error) {
	var dm_build_1548 = Dm_build_649.Dm_build_865(dm_build_1546, dm_build_1547, dm_build_1543.dm_build_1349)
	var dm_build_1549 = 0
	var dm_build_1550 = len(dm_build_1548)
	var dm_build_1551 = 0
	var dm_build_1552 = 0
	var dm_build_1553 = 0
	var dm_build_1554 = dm_build_1550/Dm_build_127 + 1
	var dm_build_1555 byte = 0
	var dm_build_1556 byte = 0x01
	var dm_build_1557 byte = 0x02
	for i := 0; i < dm_build_1554; i++ {
		dm_build_1555 = 0
		if i == 0 {
			dm_build_1555 |= dm_build_1556
		}
		if i == dm_build_1554-1 {
			dm_build_1555 |= dm_build_1557
		}
		dm_build_1553 = dm_build_1550 - dm_build_1552
		if dm_build_1553 > Dm_build_127 {
			dm_build_1553 = Dm_build_127
		}

		setLobData := dm_build_554(dm_build_1543, &dm_build_1544.lob, dm_build_1555, dm_build_1545, dm_build_1548, dm_build_1549, dm_build_1553)
		ret, err := dm_build_1543.dm_build_1385(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_1551, nil
		} else {
			dm_build_1545 += int(tmp)
			dm_build_1551 += int(tmp)
			dm_build_1552 += dm_build_1553
			dm_build_1549 += dm_build_1553
		}
	}
	return dm_build_1551, nil
}

func (dm_build_1559 *dm_build_1345) dm_build_1558(dm_build_1560 *DmBlob, dm_build_1561 int, dm_build_1562 []byte) (int, error) {
	var dm_build_1563 = 0
	var dm_build_1564 = len(dm_build_1562)
	var dm_build_1565 = 0
	var dm_build_1566 = 0
	var dm_build_1567 = 0
	var dm_build_1568 = dm_build_1564/Dm_build_127 + 1
	var dm_build_1569 byte = 0
	var dm_build_1570 byte = 0x01
	var dm_build_1571 byte = 0x02
	for i := 0; i < dm_build_1568; i++ {
		dm_build_1569 = 0
		if i == 0 {
			dm_build_1569 |= dm_build_1570
		}
		if i == dm_build_1568-1 {
			dm_build_1569 |= dm_build_1571
		}
		dm_build_1567 = dm_build_1564 - dm_build_1566
		if dm_build_1567 > Dm_build_127 {
			dm_build_1567 = Dm_build_127
		}

		setLobData := dm_build_554(dm_build_1559, &dm_build_1560.lob, dm_build_1569, dm_build_1561, dm_build_1562, dm_build_1563, dm_build_1567)
		ret, err := dm_build_1559.dm_build_1385(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_1565, nil
		} else {
			dm_build_1561 += int(tmp)
			dm_build_1565 += int(tmp)
			dm_build_1566 += dm_build_1567
			dm_build_1563 += dm_build_1567
		}
	}
	return dm_build_1565, nil
}

func (dm_build_1573 *dm_build_1345) dm_build_1572(dm_build_1574 *lob, dm_build_1575 int) (int64, error) {
	dm_build_1576 := dm_build_420(dm_build_1573, dm_build_1574, dm_build_1575)
	dm_build_1577, dm_build_1578 := dm_build_1573.dm_build_1385(dm_build_1576)
	if dm_build_1578 != nil {
		return dm_build_1574.length, dm_build_1578
	}
	return dm_build_1577.(int64), nil
}

func (dm_build_1580 *dm_build_1345) dm_build_1579(dm_build_1581 []interface{}, dm_build_1582 []interface{}, dm_build_1583 int) bool {
	var dm_build_1584 = false
	dm_build_1581[dm_build_1583] = dm_build_1582[dm_build_1583]

	if binder, ok := dm_build_1582[dm_build_1583].(iOffRowBinder); ok {
		dm_build_1584 = true
		dm_build_1581[dm_build_1583] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_1580.dm_build_1349) {
			dm_build_1581[dm_build_1583] = &lobCtl{lob.buildCtlData()}
			dm_build_1584 = false
		}
	} else {
		dm_build_1581[dm_build_1583] = dm_build_1582[dm_build_1583]
	}
	return dm_build_1584
}

func (dm_build_1586 *dm_build_1345) dm_build_1585(dm_build_1587 *DmStatement, dm_build_1588 parameter, dm_build_1589 int, dm_build_1590 iOffRowBinder) error {
	var dm_build_1591 = Dm_build_935()
	dm_build_1590.read(dm_build_1591)
	var dm_build_1592 = 0
	for !dm_build_1590.isReadOver() || dm_build_1591.Dm_build_936() > 0 {
		if !dm_build_1590.isReadOver() && dm_build_1591.Dm_build_936() < Dm_build_127 {
			dm_build_1590.read(dm_build_1591)
		}
		if dm_build_1591.Dm_build_936() > Dm_build_127 {
			dm_build_1592 = Dm_build_127
		} else {
			dm_build_1592 = dm_build_1591.Dm_build_936()
		}

		putData := dm_build_525(dm_build_1586, dm_build_1587, int16(dm_build_1589), dm_build_1591, int32(dm_build_1592))
		_, err := dm_build_1586.dm_build_1385(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1594 *dm_build_1345) dm_build_1593() ([]byte, error) {
	var dm_build_1595 error
	if dm_build_1594.dm_build_1353 == nil {
		if dm_build_1594.dm_build_1353, dm_build_1595 = security.NewClientKeyPair(); dm_build_1595 != nil {
			return nil, dm_build_1595
		}
	}
	return security.Bn2Bytes(dm_build_1594.dm_build_1353.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_1597 *dm_build_1345) dm_build_1596() (*security.DhKey, error) {
	var dm_build_1598 error
	if dm_build_1597.dm_build_1353 == nil {
		if dm_build_1597.dm_build_1353, dm_build_1598 = security.NewClientKeyPair(); dm_build_1598 != nil {
			return nil, dm_build_1598
		}
	}
	return dm_build_1597.dm_build_1353, nil
}

func (dm_build_1600 *dm_build_1345) dm_build_1599(dm_build_1601 int, dm_build_1602 []byte, dm_build_1603 string, dm_build_1604 int) (dm_build_1605 error) {
	if dm_build_1601 > 0 && dm_build_1601 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_1602 != nil {
		dm_build_1600.dm_build_1350, dm_build_1605 = security.NewSymmCipher(dm_build_1601, dm_build_1602)
	} else if dm_build_1601 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_1600.dm_build_1350, dm_build_1605 = security.NewThirdPartCipher(dm_build_1601, dm_build_1602, dm_build_1603, dm_build_1604); dm_build_1605 != nil {
			dm_build_1605 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_1605.Error()).throw()
		}
	}
	return
}

func (dm_build_1607 *dm_build_1345) dm_build_1606(dm_build_1608 bool) (dm_build_1609 error) {
	if dm_build_1607.dm_build_1347, dm_build_1609 = security.NewTLSFromTCP(dm_build_1607.dm_build_1346, dm_build_1607.dm_build_1349.dmConnector.sslCertPath, dm_build_1607.dm_build_1349.dmConnector.sslKeyPath, dm_build_1607.dm_build_1349.dmConnector.user); dm_build_1609 != nil {
		return
	}
	if !dm_build_1608 {
		dm_build_1607.dm_build_1347 = nil
	}
	return
}

func (dm_build_1611 *dm_build_1345) dm_build_1610(dm_build_1612 dm_build_135) bool {
	return dm_build_1612.dm_build_150() != Dm_build_42 && dm_build_1611.dm_build_1349.sslEncrypt == 1
}
