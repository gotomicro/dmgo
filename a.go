/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"crypto/tls"
	"net"
	"strconv"
	"time"

	"gitee.com/chunanyong/dm/security"
)

const (
	Dm_build_1310 = 8192
	Dm_build_1311 = 2 * time.Second
)

type dm_build_1312 struct {
	dm_build_1313 *net.TCPConn
	dm_build_1314 *tls.Conn
	dm_build_1315 *Dm_build_980
	dm_build_1316 *DmConnection
	dm_build_1317 security.Cipher
	dm_build_1318 bool
	dm_build_1319 bool
	dm_build_1320 *security.DhKey

	dm_build_1321 bool
	dm_build_1322 string
	dm_build_1323 bool
}

func dm_build_1324(dm_build_1325 *DmConnection) (*dm_build_1312, error) {
	dm_build_1326, dm_build_1327 := dm_build_1329(dm_build_1325.dmConnector.host+":"+strconv.Itoa(int(dm_build_1325.dmConnector.port)), time.Duration(dm_build_1325.dmConnector.socketTimeout)*time.Second)
	if dm_build_1327 != nil {
		return nil, dm_build_1327
	}

	dm_build_1328 := dm_build_1312{}
	dm_build_1328.dm_build_1313 = dm_build_1326
	dm_build_1328.dm_build_1315 = Dm_build_983(Dm_build_7)
	dm_build_1328.dm_build_1316 = dm_build_1325
	dm_build_1328.dm_build_1318 = false
	dm_build_1328.dm_build_1319 = false
	dm_build_1328.dm_build_1321 = false
	dm_build_1328.dm_build_1322 = ""
	dm_build_1328.dm_build_1323 = false
	dm_build_1325.Access = &dm_build_1328

	return &dm_build_1328, nil
}

func dm_build_1329(dm_build_1330 string, dm_build_1331 time.Duration) (*net.TCPConn, error) {
	dm_build_1332, dm_build_1333 := net.DialTimeout("tcp", dm_build_1330, dm_build_1331)
	if dm_build_1333 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_1330).throw()
	}

	if tcpConn, ok := dm_build_1332.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1311)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_1335 *dm_build_1312) dm_build_1334(dm_build_1336 dm_build_128) bool {
	var dm_build_1337 = dm_build_1335.dm_build_1316.dmConnector.compress
	if dm_build_1336.dm_build_143() == Dm_build_35 || dm_build_1337 == Dm_build_84 {
		return false
	}

	if dm_build_1337 == Dm_build_82 {
		return true
	} else if dm_build_1337 == Dm_build_83 {
		return !dm_build_1335.dm_build_1316.Local && dm_build_1336.dm_build_141() > Dm_build_81
	}

	return false
}

func (dm_build_1339 *dm_build_1312) dm_build_1338(dm_build_1340 dm_build_128) bool {
	var dm_build_1341 = dm_build_1339.dm_build_1316.dmConnector.compress
	if dm_build_1340.dm_build_143() == Dm_build_35 || dm_build_1341 == Dm_build_84 {
		return false
	}

	if dm_build_1341 == Dm_build_82 {
		return true
	} else if dm_build_1341 == Dm_build_83 {
		return dm_build_1339.dm_build_1315.Dm_build_1243(Dm_build_43) == 1
	}

	return false
}

func (dm_build_1343 *dm_build_1312) dm_build_1342(dm_build_1344 dm_build_128) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_1346 := dm_build_1344.dm_build_141()

	if dm_build_1346 > 0 {

		if dm_build_1343.dm_build_1334(dm_build_1344) {
			var retBytes, err = Compress(dm_build_1343.dm_build_1315, Dm_build_36, int(dm_build_1346), int(dm_build_1343.dm_build_1316.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_1343.dm_build_1315.Dm_build_994(Dm_build_36)

			dm_build_1343.dm_build_1315.Dm_build_1031(dm_build_1346)

			dm_build_1343.dm_build_1315.Dm_build_1059(retBytes)

			dm_build_1344.dm_build_142(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_1343.dm_build_1315.Dm_build_1163(Dm_build_43, 1)
		}

		if dm_build_1343.dm_build_1319 {
			dm_build_1346 = dm_build_1344.dm_build_141()
			var retBytes = dm_build_1343.dm_build_1317.Encrypt(dm_build_1343.dm_build_1315.Dm_build_1270(Dm_build_36, int(dm_build_1346)), true)

			dm_build_1343.dm_build_1315.Dm_build_994(Dm_build_36)

			dm_build_1343.dm_build_1315.Dm_build_1059(retBytes)

			dm_build_1344.dm_build_142(int32(len(retBytes)))
		}
	}

	if dm_build_1343.dm_build_1315.Dm_build_992() > Dm_build_8 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_1344.dm_build_137()
	if dm_build_1343.dm_build_1574(dm_build_1344) {
		if dm_build_1343.dm_build_1314 != nil {
			dm_build_1343.dm_build_1315.Dm_build_997(0)
			dm_build_1343.dm_build_1315.Dm_build_1016(dm_build_1343.dm_build_1314)
		}
	} else {
		dm_build_1343.dm_build_1315.Dm_build_997(0)
		dm_build_1343.dm_build_1315.Dm_build_1016(dm_build_1343.dm_build_1313)
	}
	return nil
}

func (dm_build_1348 *dm_build_1312) dm_build_1347(dm_build_1349 dm_build_128) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_1351 := int32(0)
	if dm_build_1348.dm_build_1574(dm_build_1349) {
		if dm_build_1348.dm_build_1314 != nil {
			dm_build_1348.dm_build_1315.Dm_build_994(0)
			dm_build_1348.dm_build_1315.Dm_build_1010(dm_build_1348.dm_build_1314, Dm_build_36)
			dm_build_1351 = dm_build_1349.dm_build_141()
			if dm_build_1351 > 0 {
				dm_build_1348.dm_build_1315.Dm_build_1010(dm_build_1348.dm_build_1314, int(dm_build_1351))
			}
		}
	} else {

		dm_build_1348.dm_build_1315.Dm_build_994(0)
		dm_build_1348.dm_build_1315.Dm_build_1010(dm_build_1348.dm_build_1313, Dm_build_36)
		dm_build_1351 = dm_build_1349.dm_build_141()

		if dm_build_1351 > 0 {
			dm_build_1348.dm_build_1315.Dm_build_1010(dm_build_1348.dm_build_1313, int(dm_build_1351))
		}
	}

	dm_build_1349.dm_build_138()

	dm_build_1351 = dm_build_1349.dm_build_141()
	if dm_build_1351 <= 0 {
		return nil
	}

	if dm_build_1348.dm_build_1319 {
		ebytes := dm_build_1348.dm_build_1315.Dm_build_1270(Dm_build_36, int(dm_build_1351))
		bytes, err := dm_build_1348.dm_build_1317.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_1348.dm_build_1315.Dm_build_994(Dm_build_36)
		dm_build_1348.dm_build_1315.Dm_build_1059(bytes)
		dm_build_1349.dm_build_142(int32(len(bytes)))
	}

	if dm_build_1348.dm_build_1338(dm_build_1349) {

		dm_build_1351 = dm_build_1349.dm_build_141()
		cbytes := dm_build_1348.dm_build_1315.Dm_build_1270(Dm_build_36+ULINT_SIZE, int(dm_build_1351-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_1348.dm_build_1316.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_1348.dm_build_1315.Dm_build_994(Dm_build_36)
		dm_build_1348.dm_build_1315.Dm_build_1059(bytes)
		dm_build_1349.dm_build_142(int32(len(bytes)))
	}
	return nil
}

func (dm_build_1353 *dm_build_1312) dm_build_1352(dm_build_1354 dm_build_128) (dm_build_1355 interface{}, dm_build_1356 error) {
	dm_build_1356 = dm_build_1354.dm_build_132(dm_build_1354)
	if dm_build_1356 != nil {
		return nil, dm_build_1356
	}

	dm_build_1356 = dm_build_1353.dm_build_1342(dm_build_1354)
	if dm_build_1356 != nil {
		return nil, dm_build_1356
	}

	dm_build_1356 = dm_build_1353.dm_build_1347(dm_build_1354)
	if dm_build_1356 != nil {
		return nil, dm_build_1356
	}

	return dm_build_1354.dm_build_136(dm_build_1354)
}

func (dm_build_1358 *dm_build_1312) dm_build_1357() (*dm_build_565, error) {

	Dm_build_1359 := dm_build_571(dm_build_1358)
	_, dm_build_1360 := dm_build_1358.dm_build_1352(Dm_build_1359)
	if dm_build_1360 != nil {
		return nil, dm_build_1360
	}

	return Dm_build_1359, nil
}

func (dm_build_1362 *dm_build_1312) dm_build_1361() error {

	dm_build_1363 := dm_build_433(dm_build_1362)
	_, dm_build_1364 := dm_build_1362.dm_build_1352(dm_build_1363)
	if dm_build_1364 != nil {
		return dm_build_1364
	}

	return nil
}

func (dm_build_1366 *dm_build_1312) dm_build_1365() error {

	var dm_build_1367 *dm_build_565
	var err error
	if dm_build_1367, err = dm_build_1366.dm_build_1357(); err != nil {
		return err
	}

	if dm_build_1366.dm_build_1316.sslEncrypt == 2 {
		if err = dm_build_1366.dm_build_1570(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_1366.dm_build_1316.sslEncrypt == 1 {
		if err = dm_build_1366.dm_build_1570(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_1366.dm_build_1319 || dm_build_1366.dm_build_1318 {
		k, err := dm_build_1366.dm_build_1560()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_1367.Dm_build_569)
		encryptType := dm_build_1367.dm_build_567
		hashType := int(dm_build_1367.Dm_build_568)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_1366.dm_build_1563(encryptType, sessionKey, dm_build_1366.dm_build_1316.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_1366.dm_build_1361(); err != nil {
		return err
	}
	return nil
}

func (dm_build_1370 *dm_build_1312) Dm_build_1369(dm_build_1371 *DmStatement) error {
	dm_build_1372 := dm_build_594(dm_build_1370, dm_build_1371)
	_, dm_build_1373 := dm_build_1370.dm_build_1352(dm_build_1372)
	if dm_build_1373 != nil {
		return dm_build_1373
	}

	return nil
}

func (dm_build_1375 *dm_build_1312) Dm_build_1374(dm_build_1376 int32) error {
	dm_build_1377 := dm_build_604(dm_build_1375, dm_build_1376)
	_, dm_build_1378 := dm_build_1375.dm_build_1352(dm_build_1377)
	if dm_build_1378 != nil {
		return dm_build_1378
	}

	return nil
}

func (dm_build_1380 *dm_build_1312) Dm_build_1379(dm_build_1381 *DmStatement, dm_build_1382 bool, dm_build_1383 int16) (*execRetInfo, error) {
	dm_build_1384 := dm_build_471(dm_build_1380, dm_build_1381, dm_build_1382, dm_build_1383)
	dm_build_1385, dm_build_1386 := dm_build_1380.dm_build_1352(dm_build_1384)
	if dm_build_1386 != nil {
		return nil, dm_build_1386
	}
	return dm_build_1385.(*execRetInfo), nil
}

func (dm_build_1388 *dm_build_1312) Dm_build_1387(dm_build_1389 *DmStatement, dm_build_1390 int16) (*execRetInfo, error) {
	return dm_build_1388.Dm_build_1379(dm_build_1389, false, Dm_build_88)
}

func (dm_build_1392 *dm_build_1312) Dm_build_1391(dm_build_1393 *DmStatement, dm_build_1394 []OptParameter) (*execRetInfo, error) {
	dm_build_1395, dm_build_1396 := dm_build_1392.dm_build_1352(dm_build_230(dm_build_1392, dm_build_1393, dm_build_1394))
	if dm_build_1396 != nil {
		return nil, dm_build_1396
	}

	return dm_build_1395.(*execRetInfo), nil
}

func (dm_build_1398 *dm_build_1312) Dm_build_1397(dm_build_1399 *DmStatement, dm_build_1400 int16) (*execRetInfo, error) {
	return dm_build_1398.Dm_build_1379(dm_build_1399, true, dm_build_1400)
}

func (dm_build_1402 *dm_build_1312) Dm_build_1401(dm_build_1403 *DmStatement, dm_build_1404 [][]interface{}) (*execRetInfo, error) {
	dm_build_1405 := dm_build_253(dm_build_1402, dm_build_1403, dm_build_1404)
	dm_build_1406, dm_build_1407 := dm_build_1402.dm_build_1352(dm_build_1405)
	if dm_build_1407 != nil {
		return nil, dm_build_1407
	}
	return dm_build_1406.(*execRetInfo), nil
}

func (dm_build_1409 *dm_build_1312) Dm_build_1408(dm_build_1410 *DmStatement, dm_build_1411 [][]interface{}, dm_build_1412 bool) (*execRetInfo, error) {
	var dm_build_1413, dm_build_1414 = 0, 0
	var dm_build_1415 = len(dm_build_1411)
	var dm_build_1416 [][]interface{}
	var dm_build_1417 = NewExceInfo()
	dm_build_1417.updateCounts = make([]int64, dm_build_1415)
	var dm_build_1418 = false
	for dm_build_1413 < dm_build_1415 {
		for dm_build_1414 = dm_build_1413; dm_build_1414 < dm_build_1415; dm_build_1414++ {
			paramData := dm_build_1411[dm_build_1414]
			bindData := make([]interface{}, dm_build_1410.paramCount)
			dm_build_1418 = false
			for icol := 0; icol < int(dm_build_1410.paramCount); icol++ {
				if dm_build_1410.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_1409.dm_build_1543(bindData, paramData, icol) {
					dm_build_1418 = true
					break
				}
			}

			if dm_build_1418 {
				break
			}
			dm_build_1416 = append(dm_build_1416, bindData)
		}

		if dm_build_1414 != dm_build_1413 {
			tmpExecInfo, err := dm_build_1409.Dm_build_1401(dm_build_1410, dm_build_1416)
			if err != nil {
				return nil, err
			}
			dm_build_1416 = dm_build_1416[0:0]
			dm_build_1417.union(tmpExecInfo, dm_build_1413, dm_build_1414-dm_build_1413)
		}

		if dm_build_1414 < dm_build_1415 {
			tmpExecInfo, err := dm_build_1409.Dm_build_1419(dm_build_1410, dm_build_1411[dm_build_1414], dm_build_1412)
			if err != nil {
				return nil, err
			}

			dm_build_1412 = true
			dm_build_1417.union(tmpExecInfo, dm_build_1414, 1)
		}

		dm_build_1413 = dm_build_1414 + 1
	}
	for _, i := range dm_build_1417.updateCounts {
		if i > 0 {
			dm_build_1417.updateCount += i
		}
	}
	return dm_build_1417, nil
}

func (dm_build_1420 *dm_build_1312) Dm_build_1419(dm_build_1421 *DmStatement, dm_build_1422 []interface{}, dm_build_1423 bool) (*execRetInfo, error) {

	var dm_build_1424 = make([]interface{}, dm_build_1421.paramCount)
	for icol := 0; icol < int(dm_build_1421.paramCount); icol++ {
		if dm_build_1421.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_1420.dm_build_1543(dm_build_1424, dm_build_1422, icol) {

			if !dm_build_1423 {
				preExecute := dm_build_461(dm_build_1420, dm_build_1421, dm_build_1421.params)
				dm_build_1420.dm_build_1352(preExecute)
				dm_build_1423 = true
			}

			dm_build_1420.dm_build_1549(dm_build_1421, dm_build_1421.params[icol], icol, dm_build_1422[icol].(iOffRowBinder))
			dm_build_1424[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_1425 = make([][]interface{}, 1, 1)
	dm_build_1425[0] = dm_build_1424

	dm_build_1426 := dm_build_253(dm_build_1420, dm_build_1421, dm_build_1425)
	dm_build_1427, dm_build_1428 := dm_build_1420.dm_build_1352(dm_build_1426)
	if dm_build_1428 != nil {
		return nil, dm_build_1428
	}
	return dm_build_1427.(*execRetInfo), nil
}

func (dm_build_1430 *dm_build_1312) Dm_build_1429(dm_build_1431 *DmStatement, dm_build_1432 int16) (*execRetInfo, error) {
	dm_build_1433 := dm_build_448(dm_build_1430, dm_build_1431, dm_build_1432)

	dm_build_1434, dm_build_1435 := dm_build_1430.dm_build_1352(dm_build_1433)
	if dm_build_1435 != nil {
		return nil, dm_build_1435
	}
	return dm_build_1434.(*execRetInfo), nil
}

func (dm_build_1437 *dm_build_1312) Dm_build_1436(dm_build_1438 *innerRows, dm_build_1439 int64) (*execRetInfo, error) {
	dm_build_1440 := dm_build_353(dm_build_1437, dm_build_1438, dm_build_1439, INT64_MAX)
	dm_build_1441, dm_build_1442 := dm_build_1437.dm_build_1352(dm_build_1440)
	if dm_build_1442 != nil {
		return nil, dm_build_1442
	}
	return dm_build_1441.(*execRetInfo), nil
}

func (dm_build_1444 *dm_build_1312) Commit() error {
	dm_build_1445 := dm_build_216(dm_build_1444)
	_, dm_build_1446 := dm_build_1444.dm_build_1352(dm_build_1445)
	if dm_build_1446 != nil {
		return dm_build_1446
	}

	return nil
}

func (dm_build_1448 *dm_build_1312) Rollback() error {
	dm_build_1449 := dm_build_509(dm_build_1448)
	_, dm_build_1450 := dm_build_1448.dm_build_1352(dm_build_1449)
	if dm_build_1450 != nil {
		return dm_build_1450
	}

	return nil
}

func (dm_build_1452 *dm_build_1312) Dm_build_1451(dm_build_1453 *DmConnection) error {
	dm_build_1454 := dm_build_514(dm_build_1452, dm_build_1453.IsoLevel)
	_, dm_build_1455 := dm_build_1452.dm_build_1352(dm_build_1454)
	if dm_build_1455 != nil {
		return dm_build_1455
	}

	return nil
}

func (dm_build_1457 *dm_build_1312) Dm_build_1456(dm_build_1458 *DmStatement, dm_build_1459 string) error {
	dm_build_1460 := dm_build_221(dm_build_1457, dm_build_1458, dm_build_1459)
	_, dm_build_1461 := dm_build_1457.dm_build_1352(dm_build_1460)
	if dm_build_1461 != nil {
		return dm_build_1461
	}

	return nil
}

func (dm_build_1463 *dm_build_1312) Dm_build_1462(dm_build_1464 []uint32) ([]int64, error) {
	dm_build_1465 := dm_build_612(dm_build_1463, dm_build_1464)
	dm_build_1466, dm_build_1467 := dm_build_1463.dm_build_1352(dm_build_1465)
	if dm_build_1467 != nil {
		return nil, dm_build_1467
	}
	return dm_build_1466.([]int64), nil
}

func (dm_build_1469 *dm_build_1312) Close() error {
	if dm_build_1469.dm_build_1323 {
		return nil
	}

	dm_build_1470 := dm_build_1469.dm_build_1313.Close()
	if dm_build_1470 != nil {
		return dm_build_1470
	}

	dm_build_1469.dm_build_1316 = nil
	dm_build_1469.dm_build_1323 = true
	return nil
}

func (dm_build_1472 *dm_build_1312) dm_build_1471(dm_build_1473 *lob) (int64, error) {
	dm_build_1474 := dm_build_384(dm_build_1472, dm_build_1473)
	dm_build_1475, dm_build_1476 := dm_build_1472.dm_build_1352(dm_build_1474)
	if dm_build_1476 != nil {
		return 0, dm_build_1476
	}
	return dm_build_1475.(int64), nil
}

func (dm_build_1478 *dm_build_1312) dm_build_1477(dm_build_1479 *lob, dm_build_1480 int32, dm_build_1481 int32) ([]byte, error) {
	dm_build_1482 := dm_build_371(dm_build_1478, dm_build_1479, int(dm_build_1480), int(dm_build_1481))
	dm_build_1483, dm_build_1484 := dm_build_1478.dm_build_1352(dm_build_1482)
	if dm_build_1484 != nil {
		return nil, dm_build_1484
	}
	return dm_build_1483.([]byte), nil
}

func (dm_build_1486 *dm_build_1312) dm_build_1485(dm_build_1487 *DmBlob, dm_build_1488 int32, dm_build_1489 int32) ([]byte, error) {
	var dm_build_1490 = make([]byte, dm_build_1489)
	var dm_build_1491 int32 = 0
	var dm_build_1492 int32 = 0
	var dm_build_1493 []byte
	var dm_build_1494 error
	for dm_build_1491 < dm_build_1489 {
		dm_build_1492 = dm_build_1489 - dm_build_1491
		if dm_build_1492 > Dm_build_121 {
			dm_build_1492 = Dm_build_121
		}
		dm_build_1493, dm_build_1494 = dm_build_1486.dm_build_1477(&dm_build_1487.lob, dm_build_1488, dm_build_1492)
		if dm_build_1494 != nil {
			return nil, dm_build_1494
		}
		if dm_build_1493 == nil || len(dm_build_1493) == 0 {
			break
		}
		Dm_build_623.Dm_build_679(dm_build_1490, int(dm_build_1491), dm_build_1493, 0, len(dm_build_1493))
		dm_build_1491 += int32(len(dm_build_1493))
		dm_build_1488 += int32(len(dm_build_1493))
		if dm_build_1487.readOver {
			break
		}
	}
	return dm_build_1490, nil
}

func (dm_build_1496 *dm_build_1312) dm_build_1495(dm_build_1497 *DmClob, dm_build_1498 int32, dm_build_1499 int32) (string, error) {
	var dm_build_1500 bytes.Buffer
	var dm_build_1501 int32 = 0
	var dm_build_1502 int32 = 0
	var dm_build_1503 []byte
	var dm_build_1504 string
	var dm_build_1505 error
	for dm_build_1501 < dm_build_1499 {
		dm_build_1502 = dm_build_1499 - dm_build_1501
		if dm_build_1502 > Dm_build_121/2 {
			dm_build_1502 = Dm_build_121 / 2
		}
		dm_build_1503, dm_build_1505 = dm_build_1496.dm_build_1477(&dm_build_1497.lob, dm_build_1498, dm_build_1502)
		if dm_build_1505 != nil {
			return "", dm_build_1505
		}
		if dm_build_1503 == nil || len(dm_build_1503) == 0 {
			break
		}
		dm_build_1504 = Dm_build_623.Dm_build_780(dm_build_1503, 0, len(dm_build_1503), dm_build_1497.serverEncoding, dm_build_1496.dm_build_1316)

		dm_build_1500.WriteString(dm_build_1504)
		dm_build_1501 += int32(len(dm_build_1504))
		dm_build_1498 += int32(len(dm_build_1504))
		if dm_build_1497.readOver {
			break
		}
	}
	return dm_build_1500.String(), nil
}

func (dm_build_1507 *dm_build_1312) dm_build_1506(dm_build_1508 *DmClob, dm_build_1509 int, dm_build_1510 string, dm_build_1511 string) (int, error) {
	var dm_build_1512 = Dm_build_623.Dm_build_836(dm_build_1510, dm_build_1511, dm_build_1507.dm_build_1316)
	var dm_build_1513 = 0
	var dm_build_1514 = len(dm_build_1512)
	var dm_build_1515 = 0
	var dm_build_1516 = 0
	var dm_build_1517 = 0
	var dm_build_1518 = dm_build_1514/Dm_build_120 + 1
	var dm_build_1519 byte = 0
	var dm_build_1520 byte = 0x01
	var dm_build_1521 byte = 0x02
	for i := 0; i < dm_build_1518; i++ {
		dm_build_1519 = 0
		if i == 0 {
			dm_build_1519 |= dm_build_1520
		}
		if i == dm_build_1518-1 {
			dm_build_1519 |= dm_build_1521
		}
		dm_build_1517 = dm_build_1514 - dm_build_1516
		if dm_build_1517 > Dm_build_120 {
			dm_build_1517 = Dm_build_120
		}

		setLobData := dm_build_528(dm_build_1507, &dm_build_1508.lob, dm_build_1519, dm_build_1509, dm_build_1512, dm_build_1513, dm_build_1517)
		ret, err := dm_build_1507.dm_build_1352(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_1515, nil
		} else {
			dm_build_1509 += int(tmp)
			dm_build_1515 += int(tmp)
			dm_build_1516 += dm_build_1517
			dm_build_1513 += dm_build_1517
		}
	}
	return dm_build_1515, nil
}

func (dm_build_1523 *dm_build_1312) dm_build_1522(dm_build_1524 *DmBlob, dm_build_1525 int, dm_build_1526 []byte) (int, error) {
	var dm_build_1527 = 0
	var dm_build_1528 = len(dm_build_1526)
	var dm_build_1529 = 0
	var dm_build_1530 = 0
	var dm_build_1531 = 0
	var dm_build_1532 = dm_build_1528/Dm_build_120 + 1
	var dm_build_1533 byte = 0
	var dm_build_1534 byte = 0x01
	var dm_build_1535 byte = 0x02
	for i := 0; i < dm_build_1532; i++ {
		dm_build_1533 = 0
		if i == 0 {
			dm_build_1533 |= dm_build_1534
		}
		if i == dm_build_1532-1 {
			dm_build_1533 |= dm_build_1535
		}
		dm_build_1531 = dm_build_1528 - dm_build_1530
		if dm_build_1531 > Dm_build_120 {
			dm_build_1531 = Dm_build_120
		}

		setLobData := dm_build_528(dm_build_1523, &dm_build_1524.lob, dm_build_1533, dm_build_1525, dm_build_1526, dm_build_1527, dm_build_1531)
		ret, err := dm_build_1523.dm_build_1352(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_1529, nil
		} else {
			dm_build_1525 += int(tmp)
			dm_build_1529 += int(tmp)
			dm_build_1530 += dm_build_1531
			dm_build_1527 += dm_build_1531
		}
	}
	return dm_build_1529, nil
}

func (dm_build_1537 *dm_build_1312) dm_build_1536(dm_build_1538 *lob, dm_build_1539 int) (int64, error) {
	dm_build_1540 := dm_build_395(dm_build_1537, dm_build_1538, dm_build_1539)
	dm_build_1541, dm_build_1542 := dm_build_1537.dm_build_1352(dm_build_1540)
	if dm_build_1542 != nil {
		return dm_build_1538.length, dm_build_1542
	}
	return dm_build_1541.(int64), nil
}

func (dm_build_1544 *dm_build_1312) dm_build_1543(dm_build_1545 []interface{}, dm_build_1546 []interface{}, dm_build_1547 int) bool {
	var dm_build_1548 = false
	if dm_build_1547 >= len(dm_build_1546) || dm_build_1546[dm_build_1547] == nil {
		dm_build_1545[dm_build_1547] = ParamDataEnum_Null
	} else if binder, ok := dm_build_1546[dm_build_1547].(iOffRowBinder); ok {
		dm_build_1548 = true
		dm_build_1545[dm_build_1547] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_1544.dm_build_1316) {
			dm_build_1545[dm_build_1547] = &lobCtl{lob.buildCtlData()}
			dm_build_1548 = false
		}
	} else {
		dm_build_1545[dm_build_1547] = dm_build_1546[dm_build_1547]
	}
	return dm_build_1548
}

func (dm_build_1550 *dm_build_1312) dm_build_1549(dm_build_1551 *DmStatement, dm_build_1552 parameter, dm_build_1553 int, dm_build_1554 iOffRowBinder) error {
	var dm_build_1555 = Dm_build_906()
	dm_build_1554.read(dm_build_1555)
	var dm_build_1556 = 0
	for !dm_build_1554.isReadOver() || dm_build_1555.Dm_build_907() > 0 {
		if !dm_build_1554.isReadOver() && dm_build_1555.Dm_build_907() < Dm_build_120 {
			dm_build_1554.read(dm_build_1555)
		}
		if dm_build_1555.Dm_build_907() > Dm_build_120 {
			dm_build_1556 = Dm_build_120
		} else {
			dm_build_1556 = dm_build_1555.Dm_build_907()
		}

		putData := dm_build_499(dm_build_1550, dm_build_1551, int16(dm_build_1553), dm_build_1555, int32(dm_build_1556))
		_, err := dm_build_1550.dm_build_1352(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1558 *dm_build_1312) dm_build_1557() ([]byte, error) {
	var dm_build_1559 error
	if dm_build_1558.dm_build_1320 == nil {
		if dm_build_1558.dm_build_1320, dm_build_1559 = security.NewClientKeyPair(); dm_build_1559 != nil {
			return nil, dm_build_1559
		}
	}
	return security.Bn2Bytes(dm_build_1558.dm_build_1320.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_1561 *dm_build_1312) dm_build_1560() (*security.DhKey, error) {
	var dm_build_1562 error
	if dm_build_1561.dm_build_1320 == nil {
		if dm_build_1561.dm_build_1320, dm_build_1562 = security.NewClientKeyPair(); dm_build_1562 != nil {
			return nil, dm_build_1562
		}
	}
	return dm_build_1561.dm_build_1320, nil
}

func (dm_build_1564 *dm_build_1312) dm_build_1563(dm_build_1565 int, dm_build_1566 []byte, dm_build_1567 string, dm_build_1568 int) (dm_build_1569 error) {
	if dm_build_1565 > 0 && dm_build_1565 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_1566 != nil {
		dm_build_1564.dm_build_1317, dm_build_1569 = security.NewSymmCipher(dm_build_1565, dm_build_1566)
	} else if dm_build_1565 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_1564.dm_build_1317, dm_build_1569 = security.NewThirdPartCipher(dm_build_1565, dm_build_1566, dm_build_1567, dm_build_1568); dm_build_1569 != nil {
			dm_build_1569 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_1569.Error()).throw()
		}
	}
	return
}

func (dm_build_1571 *dm_build_1312) dm_build_1570(dm_build_1572 bool) (dm_build_1573 error) {
	if dm_build_1571.dm_build_1314, dm_build_1573 = security.NewTLSFromTCP(dm_build_1571.dm_build_1313, dm_build_1571.dm_build_1316.dmConnector.sslCertPath, dm_build_1571.dm_build_1316.dmConnector.sslKeyPath, dm_build_1571.dm_build_1316.dmConnector.user); dm_build_1573 != nil {
		return
	}
	if !dm_build_1572 {
		dm_build_1571.dm_build_1314 = nil
	}
	return
}

func (dm_build_1575 *dm_build_1312) dm_build_1574(dm_build_1576 dm_build_128) bool {
	return dm_build_1576.dm_build_143() != Dm_build_35 && dm_build_1575.dm_build_1316.sslEncrypt == 1
}
