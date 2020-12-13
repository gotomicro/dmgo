/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"crypto/tls"
	"dm/security"
	"net"
	"strconv"
	"time"
)

const (
	Dm_build_1283 = 8192
	Dm_build_1284 = 2 * time.Second
)

type dm_build_1285 struct {
	dm_build_1286 *net.TCPConn
	dm_build_1287 *tls.Conn
	dm_build_1288 *Dm_build_953
	dm_build_1289 *DmConnection
	dm_build_1290 security.Cipher
	dm_build_1291 bool
	dm_build_1292 bool
	dm_build_1293 *security.DhKey
	dm_build_1294 string
	dm_build_1295 bool
}

func dm_build_1296(dm_build_1297 *DmConnection) (*dm_build_1285, error) {
	dm_build_1298, dm_build_1299 := dm_build_1301(dm_build_1297.dmConnector.host+":"+strconv.Itoa(dm_build_1297.dmConnector.port), time.Duration(dm_build_1297.dmConnector.socketTimeout)*time.Second)
	if dm_build_1299 != nil {
		return nil, dm_build_1299
	}

	dm_build_1300 := dm_build_1285{}
	dm_build_1300.dm_build_1286 = dm_build_1298
	dm_build_1300.dm_build_1288 = Dm_build_956(Dm_build_6)
	dm_build_1300.dm_build_1289 = dm_build_1297
	dm_build_1300.dm_build_1291 = false
	dm_build_1300.dm_build_1292 = false
	dm_build_1300.dm_build_1294 = ""
	dm_build_1300.dm_build_1295 = false
	dm_build_1297.Access = &dm_build_1300

	return &dm_build_1300, nil
}

func dm_build_1301(dm_build_1302 string, dm_build_1303 time.Duration) (*net.TCPConn, error) {
	dm_build_1304, dm_build_1305 := net.DialTimeout("tcp", dm_build_1302, dm_build_1303)
	if dm_build_1305 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_1302).throw()
	}

	if tcpConn, ok := dm_build_1304.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1284)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_1307 *dm_build_1285) dm_build_1306(dm_build_1308 dm_build_123) bool {
	var dm_build_1309 = dm_build_1307.dm_build_1289.dmConnector.compress
	if dm_build_1308.dm_build_137() == Dm_build_33 || dm_build_1309 == Dm_build_81 {
		return false
	}

	if dm_build_1309 == Dm_build_79 {
		return true
	} else if dm_build_1309 == Dm_build_80 {
		return !dm_build_1307.dm_build_1289.Local && dm_build_1308.dm_build_135() > Dm_build_78
	}

	return false
}

func (dm_build_1311 *dm_build_1285) dm_build_1310(dm_build_1312 dm_build_123) bool {
	var dm_build_1313 = dm_build_1311.dm_build_1289.dmConnector.compress
	if dm_build_1312.dm_build_137() == Dm_build_33 || dm_build_1313 == Dm_build_81 {
		return false
	}

	if dm_build_1313 == Dm_build_79 {
		return true
	} else if dm_build_1313 == Dm_build_80 {
		return dm_build_1311.dm_build_1288.Dm_build_1216(Dm_build_41) == 1
	}

	return false
}

func (dm_build_1315 *dm_build_1285) dm_build_1314(dm_build_1316 dm_build_123) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_1318 := dm_build_1316.dm_build_135()

	if dm_build_1318 > 0 {

		if dm_build_1315.dm_build_1306(dm_build_1316) {
			var retBytes, err = Compress(dm_build_1315.dm_build_1288, Dm_build_34, int(dm_build_1318), int(dm_build_1315.dm_build_1289.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_1315.dm_build_1288.Dm_build_967(Dm_build_34)

			dm_build_1315.dm_build_1288.Dm_build_1004(dm_build_1318)

			dm_build_1315.dm_build_1288.Dm_build_1032(retBytes)

			dm_build_1316.dm_build_136(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_1315.dm_build_1288.Dm_build_1136(Dm_build_41, 1)
		}

		if dm_build_1315.dm_build_1292 {
			dm_build_1318 = dm_build_1316.dm_build_135()
			var retBytes = dm_build_1315.dm_build_1290.Encrypt(dm_build_1315.dm_build_1288.Dm_build_1243(Dm_build_34, int(dm_build_1318)), true)

			dm_build_1315.dm_build_1288.Dm_build_967(Dm_build_34)

			dm_build_1315.dm_build_1288.Dm_build_1032(retBytes)

			dm_build_1316.dm_build_136(int32(len(retBytes)))
		}
	}

	dm_build_1316.dm_build_132()
	if dm_build_1315.dm_build_1546(dm_build_1316) {
		if dm_build_1315.dm_build_1287 != nil {
			dm_build_1315.dm_build_1288.Dm_build_970(0)
			dm_build_1315.dm_build_1288.Dm_build_989(dm_build_1315.dm_build_1287)
		}
	} else {
		dm_build_1315.dm_build_1288.Dm_build_970(0)
		dm_build_1315.dm_build_1288.Dm_build_989(dm_build_1315.dm_build_1286)
	}
	return nil
}

func (dm_build_1320 *dm_build_1285) dm_build_1319(dm_build_1321 dm_build_123) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_1323 := int32(0)
	if dm_build_1320.dm_build_1546(dm_build_1321) {
		if dm_build_1320.dm_build_1287 != nil {
			dm_build_1320.dm_build_1288.Dm_build_967(0)
			dm_build_1320.dm_build_1288.Dm_build_983(dm_build_1320.dm_build_1287, Dm_build_34)
			dm_build_1323 = dm_build_1321.dm_build_135()
			if dm_build_1323 > 0 {
				dm_build_1320.dm_build_1288.Dm_build_983(dm_build_1320.dm_build_1287, int(dm_build_1323))
			}
		}
	} else {

		dm_build_1320.dm_build_1288.Dm_build_967(0)
		dm_build_1320.dm_build_1288.Dm_build_983(dm_build_1320.dm_build_1286, Dm_build_34)
		dm_build_1323 = dm_build_1321.dm_build_135()

		if dm_build_1323 > 0 {
			dm_build_1320.dm_build_1288.Dm_build_983(dm_build_1320.dm_build_1286, int(dm_build_1323))
		}
	}

	dm_build_1321.dm_build_133()

	if dm_build_1323 <= 0 {
		return nil
	}

	if dm_build_1320.dm_build_1292 {
		dm_build_1323 = dm_build_1321.dm_build_135()
		ebytes := dm_build_1320.dm_build_1288.Dm_build_1243(Dm_build_34, int(dm_build_1323))
		bytes, err := dm_build_1320.dm_build_1290.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_1320.dm_build_1288.Dm_build_967(Dm_build_34)
		dm_build_1320.dm_build_1288.Dm_build_1032(bytes)
		dm_build_1321.dm_build_136(int32(len(bytes)))
	}

	if dm_build_1320.dm_build_1310(dm_build_1321) {

		dm_build_1323 = dm_build_1321.dm_build_135()
		cbytes := dm_build_1320.dm_build_1288.Dm_build_1243(Dm_build_34+ULINT_SIZE, int(dm_build_1323-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_1320.dm_build_1289.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_1320.dm_build_1288.Dm_build_967(Dm_build_34)
		dm_build_1320.dm_build_1288.Dm_build_1032(bytes)
		dm_build_1321.dm_build_136(int32(len(bytes)))
	}
	return nil
}

func (dm_build_1325 *dm_build_1285) dm_build_1324(dm_build_1326 dm_build_123) (dm_build_1327 interface{}, dm_build_1328 error) {
	dm_build_1328 = dm_build_1326.dm_build_127(dm_build_1326)
	if dm_build_1328 != nil {
		return nil, dm_build_1328
	}

	dm_build_1328 = dm_build_1325.dm_build_1314(dm_build_1326)
	if dm_build_1328 != nil {
		return nil, dm_build_1328
	}

	dm_build_1328 = dm_build_1325.dm_build_1319(dm_build_1326)
	if dm_build_1328 != nil {
		return nil, dm_build_1328
	}

	return dm_build_1326.dm_build_131(dm_build_1326)
}

func (dm_build_1330 *dm_build_1285) dm_build_1329() (*dm_build_541, error) {

	Dm_build_1331 := dm_build_547(dm_build_1330)
	_, dm_build_1332 := dm_build_1330.dm_build_1324(Dm_build_1331)
	if dm_build_1332 != nil {
		return nil, dm_build_1332
	}

	return Dm_build_1331, nil
}

func (dm_build_1334 *dm_build_1285) dm_build_1333() error {

	dm_build_1335 := dm_build_418(dm_build_1334)
	_, dm_build_1336 := dm_build_1334.dm_build_1324(dm_build_1335)
	if dm_build_1336 != nil {
		return dm_build_1336
	}

	return nil
}

func (dm_build_1338 *dm_build_1285) dm_build_1337() error {

	var dm_build_1339 *dm_build_541
	var err error
	if dm_build_1339, err = dm_build_1338.dm_build_1329(); err != nil {
		return err
	}

	if dm_build_1338.dm_build_1289.sslEncrypt == 2 {
		if err = dm_build_1338.dm_build_1542(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_1338.dm_build_1289.sslEncrypt == 1 {
		if err = dm_build_1338.dm_build_1542(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_1338.dm_build_1292 || dm_build_1338.dm_build_1291 {
		k, err := dm_build_1338.dm_build_1532()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_1339.Dm_build_545)
		encryptType := dm_build_1339.dm_build_543
		hashType := int(dm_build_1339.Dm_build_544)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_1338.dm_build_1535(encryptType, sessionKey, dm_build_1338.dm_build_1289.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_1338.dm_build_1333(); err != nil {
		return err
	}
	return nil
}

func (dm_build_1342 *dm_build_1285) Dm_build_1341(dm_build_1343 *DmStatement) error {
	dm_build_1344 := dm_build_570(dm_build_1342, dm_build_1343)
	_, dm_build_1345 := dm_build_1342.dm_build_1324(dm_build_1344)
	if dm_build_1345 != nil {
		return dm_build_1345
	}

	return nil
}

func (dm_build_1347 *dm_build_1285) Dm_build_1346(dm_build_1348 int32) error {
	dm_build_1349 := dm_build_580(dm_build_1347, dm_build_1348)
	_, dm_build_1350 := dm_build_1347.dm_build_1324(dm_build_1349)
	if dm_build_1350 != nil {
		return dm_build_1350
	}

	return nil
}

func (dm_build_1352 *dm_build_1285) Dm_build_1351(dm_build_1353 *DmStatement, dm_build_1354 bool, dm_build_1355 int16) (*execInfo, error) {
	dm_build_1356 := dm_build_454(dm_build_1352, dm_build_1353, dm_build_1354, dm_build_1355)
	dm_build_1357, dm_build_1358 := dm_build_1352.dm_build_1324(dm_build_1356)
	if dm_build_1358 != nil {
		return nil, dm_build_1358
	}
	return dm_build_1357.(*execInfo), nil
}

func (dm_build_1360 *dm_build_1285) Dm_build_1359(dm_build_1361 *DmStatement, dm_build_1362 int16) (*execInfo, error) {
	return dm_build_1360.Dm_build_1351(dm_build_1361, false, Dm_build_85)
}

func (dm_build_1364 *dm_build_1285) Dm_build_1363(dm_build_1365 *DmStatement, dm_build_1366 []OptParameter) (*execInfo, error) {
	dm_build_1367, dm_build_1368 := dm_build_1364.dm_build_1324(dm_build_216(dm_build_1364, dm_build_1365, dm_build_1366))
	if dm_build_1368 != nil {
		return nil, dm_build_1368
	}

	return dm_build_1367.(*execInfo), nil
}

func (dm_build_1370 *dm_build_1285) Dm_build_1369(dm_build_1371 *DmStatement, dm_build_1372 int16) (*execInfo, error) {
	return dm_build_1370.Dm_build_1351(dm_build_1371, true, dm_build_1372)
}

func (dm_build_1374 *dm_build_1285) Dm_build_1373(dm_build_1375 *DmStatement, dm_build_1376 [][]interface{}) (*execInfo, error) {
	dm_build_1377 := dm_build_239(dm_build_1374, dm_build_1375, dm_build_1376)
	dm_build_1378, dm_build_1379 := dm_build_1374.dm_build_1324(dm_build_1377)
	if dm_build_1379 != nil {
		return nil, dm_build_1379
	}
	return dm_build_1378.(*execInfo), nil
}

func (dm_build_1381 *dm_build_1285) Dm_build_1380(dm_build_1382 *DmStatement, dm_build_1383 [][]interface{}) (*execInfo, error) {
	var dm_build_1384, dm_build_1385 = 0, 0
	var dm_build_1386 = len(dm_build_1383)
	var dm_build_1387 [][]interface{}
	var dm_build_1388 = NewExceInfo()
	dm_build_1388.updateCounts = make([]int64, dm_build_1386)
	var dm_build_1389 = false
	var dm_build_1390 = false
	for dm_build_1384 < dm_build_1386 {
		for dm_build_1385 = dm_build_1384; dm_build_1385 < dm_build_1386; dm_build_1385++ {
			paramData := dm_build_1383[dm_build_1385]
			bindData := make([]interface{}, dm_build_1382.paramCount)
			dm_build_1389 = false
			for icol := 0; icol < int(dm_build_1382.paramCount); icol++ {
				if dm_build_1382.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_1381.dm_build_1515(bindData, paramData, icol) {
					dm_build_1389 = true
					break
				}
			}

			if dm_build_1389 {
				break
			}
			dm_build_1387 = append(dm_build_1387, bindData)
		}

		if dm_build_1385 != dm_build_1384 {
			tmpExecInfo, err := dm_build_1381.Dm_build_1373(dm_build_1382, dm_build_1387)
			if err != nil {
				return nil, err
			}
			dm_build_1387 = dm_build_1387[0:0]
			if dm_build_1385-dm_build_1384 == 1 {
				dm_build_1388.updateCounts[dm_build_1384] = tmpExecInfo.updateCount
			} else if tmpExecInfo.updateCounts != nil {
				copy(dm_build_1388.updateCounts[dm_build_1384:dm_build_1385], tmpExecInfo.updateCounts[0:dm_build_1385-dm_build_1384])
			}
		}

		if dm_build_1385 < dm_build_1386 {
			tmpExecInfo, err := dm_build_1381.Dm_build_1391(dm_build_1382, dm_build_1383[dm_build_1385], dm_build_1390)
			if err != nil {
				return nil, err
			}

			dm_build_1390 = true
			dm_build_1388.updateCounts[dm_build_1385] = tmpExecInfo.updateCount
		}

		dm_build_1384 = dm_build_1385 + 1
	}
	for _, i := range dm_build_1388.updateCounts {
		dm_build_1388.updateCount += i
	}
	return dm_build_1388, nil
}

func (dm_build_1392 *dm_build_1285) Dm_build_1391(dm_build_1393 *DmStatement, dm_build_1394 []interface{}, dm_build_1395 bool) (*execInfo, error) {

	var dm_build_1396 = make([]interface{}, dm_build_1393.paramCount)
	for icol := 0; icol < int(dm_build_1393.paramCount); icol++ {
		if dm_build_1393.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_1392.dm_build_1515(dm_build_1396, dm_build_1394, icol) {

			if !dm_build_1395 {
				preExecute := dm_build_444(dm_build_1392, dm_build_1393, dm_build_1393.params)
				dm_build_1392.dm_build_1324(preExecute)
				dm_build_1395 = true
			}

			dm_build_1392.dm_build_1521(dm_build_1393, dm_build_1393.params[icol], icol, dm_build_1394[icol].(iOffRowBinder))
			dm_build_1396[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_1397 = make([][]interface{}, 1, 1)
	dm_build_1397[0] = dm_build_1396

	dm_build_1398 := dm_build_239(dm_build_1392, dm_build_1393, dm_build_1397)
	dm_build_1399, dm_build_1400 := dm_build_1392.dm_build_1324(dm_build_1398)
	if dm_build_1400 != nil {
		return nil, dm_build_1400
	}
	return dm_build_1399.(*execInfo), nil
}

func (dm_build_1402 *dm_build_1285) Dm_build_1401(dm_build_1403 *DmStatement, dm_build_1404 int16) (*execInfo, error) {
	dm_build_1405 := dm_build_432(dm_build_1402, dm_build_1403, dm_build_1404)

	dm_build_1406, dm_build_1407 := dm_build_1402.dm_build_1324(dm_build_1405)
	if dm_build_1407 != nil {
		return nil, dm_build_1407
	}
	return dm_build_1406.(*execInfo), nil
}

func (dm_build_1409 *dm_build_1285) Dm_build_1408(dm_build_1410 *innerRows, dm_build_1411 int64) (*execInfo, error) {
	dm_build_1412 := dm_build_339(dm_build_1409, dm_build_1410, dm_build_1411, INT64_MAX)
	dm_build_1413, dm_build_1414 := dm_build_1409.dm_build_1324(dm_build_1412)
	if dm_build_1414 != nil {
		return nil, dm_build_1414
	}
	return dm_build_1413.(*execInfo), nil
}

func (dm_build_1416 *dm_build_1285) Commit() error {
	dm_build_1417 := dm_build_202(dm_build_1416)
	_, dm_build_1418 := dm_build_1416.dm_build_1324(dm_build_1417)
	if dm_build_1418 != nil {
		return dm_build_1418
	}

	return nil
}

func (dm_build_1420 *dm_build_1285) Rollback() error {
	dm_build_1421 := dm_build_492(dm_build_1420)
	_, dm_build_1422 := dm_build_1420.dm_build_1324(dm_build_1421)
	if dm_build_1422 != nil {
		return dm_build_1422
	}

	return nil
}

func (dm_build_1424 *dm_build_1285) Dm_build_1423(dm_build_1425 *DmConnection) error {
	dm_build_1426 := dm_build_497(dm_build_1424, dm_build_1425.IsoLevel)
	_, dm_build_1427 := dm_build_1424.dm_build_1324(dm_build_1426)
	if dm_build_1427 != nil {
		return dm_build_1427
	}

	return nil
}

func (dm_build_1429 *dm_build_1285) Dm_build_1428(dm_build_1430 *DmStatement, dm_build_1431 string) error {
	dm_build_1432 := dm_build_207(dm_build_1429, dm_build_1430, dm_build_1431)
	_, dm_build_1433 := dm_build_1429.dm_build_1324(dm_build_1432)
	if dm_build_1433 != nil {
		return dm_build_1433
	}

	return nil
}

func (dm_build_1435 *dm_build_1285) Dm_build_1434(dm_build_1436 []uint32) ([]int64, error) {
	dm_build_1437 := dm_build_588(dm_build_1435, dm_build_1436)
	dm_build_1438, dm_build_1439 := dm_build_1435.dm_build_1324(dm_build_1437)
	if dm_build_1439 != nil {
		return nil, dm_build_1439
	}
	return dm_build_1438.([]int64), nil
}

func (dm_build_1441 *dm_build_1285) Close() error {
	if dm_build_1441.dm_build_1295 {
		return nil
	}

	dm_build_1442 := dm_build_1441.dm_build_1286.Close()
	if dm_build_1442 != nil {
		return dm_build_1442
	}

	dm_build_1441.dm_build_1289 = nil
	dm_build_1441.dm_build_1295 = true
	return nil
}

func (dm_build_1444 *dm_build_1285) dm_build_1443(dm_build_1445 *lob) (int64, error) {
	dm_build_1446 := dm_build_370(dm_build_1444, dm_build_1445)
	dm_build_1447, dm_build_1448 := dm_build_1444.dm_build_1324(dm_build_1446)
	if dm_build_1448 != nil {
		return 0, dm_build_1448
	}
	return dm_build_1447.(int64), nil
}

func (dm_build_1450 *dm_build_1285) dm_build_1449(dm_build_1451 *lob, dm_build_1452 int32, dm_build_1453 int32) ([]byte, error) {
	dm_build_1454 := dm_build_357(dm_build_1450, dm_build_1451, int(dm_build_1452), int(dm_build_1453))
	dm_build_1455, dm_build_1456 := dm_build_1450.dm_build_1324(dm_build_1454)
	if dm_build_1456 != nil {
		return nil, dm_build_1456
	}
	return dm_build_1455.([]byte), nil
}

func (dm_build_1458 *dm_build_1285) dm_build_1457(dm_build_1459 *DmBlob, dm_build_1460 int32, dm_build_1461 int32) ([]byte, error) {
	var dm_build_1462 = make([]byte, dm_build_1461)
	var dm_build_1463 int32 = 0
	var dm_build_1464 int32 = 0
	var dm_build_1465 []byte
	var dm_build_1466 error
	for dm_build_1463 < dm_build_1461 {
		dm_build_1464 = dm_build_1461 - dm_build_1463
		if dm_build_1464 > Dm_build_118 {
			dm_build_1464 = Dm_build_118
		}
		dm_build_1465, dm_build_1466 = dm_build_1458.dm_build_1449(&dm_build_1459.lob, dm_build_1460, dm_build_1461)
		if dm_build_1466 != nil {
			return nil, dm_build_1466
		}
		if dm_build_1465 == nil || len(dm_build_1465) == 0 {
			break
		}
		Dm_build_599.Dm_build_655(dm_build_1462, int(dm_build_1463), dm_build_1465, 0, len(dm_build_1465))
		dm_build_1463 += int32(len(dm_build_1465))
		dm_build_1460 += int32(len(dm_build_1465))
		if dm_build_1459.readOver {
			break
		}
	}
	return dm_build_1462, nil
}

func (dm_build_1468 *dm_build_1285) dm_build_1467(dm_build_1469 *DmClob, dm_build_1470 int32, dm_build_1471 int32) (string, error) {
	var dm_build_1472 = ""
	var dm_build_1473 int32 = 0
	var dm_build_1474 int32 = 0
	var dm_build_1475 []byte
	var dm_build_1476 string
	var dm_build_1477 error
	for dm_build_1473 < dm_build_1471 {
		dm_build_1474 = dm_build_1471 - dm_build_1473
		if dm_build_1474 > Dm_build_118/2 {
			dm_build_1474 = Dm_build_118 / 2
		}
		dm_build_1475, dm_build_1477 = dm_build_1468.dm_build_1449(&dm_build_1469.lob, dm_build_1470, dm_build_1471)
		if dm_build_1477 != nil {
			return "", dm_build_1477
		}
		if dm_build_1475 == nil || len(dm_build_1475) == 0 {
			break
		}
		dm_build_1476 = Dm_build_599.Dm_build_753(dm_build_1475, 0, len(dm_build_1475), dm_build_1469.serverEncoding, dm_build_1468.dm_build_1289)
		dm_build_1472 += dm_build_1476
		dm_build_1473 += int32(len(dm_build_1476))
		dm_build_1470 += int32(len(dm_build_1475))
		if dm_build_1469.readOver {
			break
		}
	}
	return dm_build_1472, nil
}

func (dm_build_1479 *dm_build_1285) dm_build_1478(dm_build_1480 *DmClob, dm_build_1481 int, dm_build_1482 string, dm_build_1483 string) (int, error) {
	var dm_build_1484 = Dm_build_599.Dm_build_809(dm_build_1482, dm_build_1483, dm_build_1479.dm_build_1289)
	var dm_build_1485 = 0
	var dm_build_1486 = len(dm_build_1484)
	var dm_build_1487 = 0
	var dm_build_1488 = 0
	var dm_build_1489 = 0
	var dm_build_1490 = dm_build_1486/Dm_build_117 + 1
	var dm_build_1491 byte = 0
	var dm_build_1492 byte = 0x01
	var dm_build_1493 byte = 0x02
	for i := 0; i < dm_build_1490; i++ {
		dm_build_1491 = 0
		if i == 0 {
			dm_build_1491 |= dm_build_1492
		}
		if i == dm_build_1490-1 {
			dm_build_1491 |= dm_build_1493
		}
		dm_build_1489 = dm_build_1486 - dm_build_1488
		if dm_build_1489 > Dm_build_117 {
			dm_build_1489 = Dm_build_117
		}

		setLobData := dm_build_511(dm_build_1479, &dm_build_1480.lob, dm_build_1491, dm_build_1481, dm_build_1484, dm_build_1485, dm_build_1489)
		ret, err := dm_build_1479.dm_build_1324(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_1487, nil
		} else {
			dm_build_1481 += int(tmp)
			dm_build_1487 += int(tmp)
			dm_build_1488 += dm_build_1489
			dm_build_1485 += dm_build_1489
		}
	}
	return dm_build_1487, nil
}

func (dm_build_1495 *dm_build_1285) dm_build_1494(dm_build_1496 *DmBlob, dm_build_1497 int, dm_build_1498 []byte) (int, error) {
	var dm_build_1499 = 0
	var dm_build_1500 = len(dm_build_1498)
	var dm_build_1501 = 0
	var dm_build_1502 = 0
	var dm_build_1503 = 0
	var dm_build_1504 = dm_build_1500/Dm_build_117 + 1
	var dm_build_1505 byte = 0
	var dm_build_1506 byte = 0x01
	var dm_build_1507 byte = 0x02
	for i := 0; i < dm_build_1504; i++ {
		dm_build_1505 = 0
		if i == 0 {
			dm_build_1505 |= dm_build_1506
		}
		if i == dm_build_1504-1 {
			dm_build_1505 |= dm_build_1507
		}
		dm_build_1503 = dm_build_1500 - dm_build_1502
		if dm_build_1503 > Dm_build_117 {
			dm_build_1503 = Dm_build_117
		}

		setLobData := dm_build_511(dm_build_1495, &dm_build_1496.lob, dm_build_1505, dm_build_1497, dm_build_1498, dm_build_1499, dm_build_1503)
		ret, err := dm_build_1495.dm_build_1324(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_1501, nil
		} else {
			dm_build_1497 += int(tmp)
			dm_build_1501 += int(tmp)
			dm_build_1502 += dm_build_1503
			dm_build_1499 += dm_build_1503
		}
	}
	return dm_build_1501, nil
}

func (dm_build_1509 *dm_build_1285) dm_build_1508(dm_build_1510 *lob, dm_build_1511 int) (int64, error) {
	dm_build_1512 := dm_build_381(dm_build_1509, dm_build_1510, dm_build_1511)
	dm_build_1513, dm_build_1514 := dm_build_1509.dm_build_1324(dm_build_1512)
	if dm_build_1514 != nil {
		return dm_build_1510.length, dm_build_1514
	}
	return dm_build_1513.(int64), nil
}

func (dm_build_1516 *dm_build_1285) dm_build_1515(dm_build_1517 []interface{}, dm_build_1518 []interface{}, dm_build_1519 int) bool {
	var dm_build_1520 = false
	if dm_build_1519 >= len(dm_build_1518) || dm_build_1518[dm_build_1519] == nil {
		dm_build_1517[dm_build_1519] = ParamDataEnum_Null
	} else if binder, ok := dm_build_1518[dm_build_1519].(iOffRowBinder); ok {
		dm_build_1520 = true
		dm_build_1517[dm_build_1519] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_1516.dm_build_1289) {
			dm_build_1517[dm_build_1519] = &lobCtl{lob.buildCtlData()}
			dm_build_1520 = false
		}
	} else {
		dm_build_1517[dm_build_1519] = dm_build_1518[dm_build_1519]
	}
	return dm_build_1520
}

func (dm_build_1522 *dm_build_1285) dm_build_1521(dm_build_1523 *DmStatement, dm_build_1524 parameter, dm_build_1525 int, dm_build_1526 iOffRowBinder) error {
	var dm_build_1527 = Dm_build_879()
	dm_build_1526.read(dm_build_1527)
	var dm_build_1528 = 0
	for !dm_build_1526.isReadOver() || dm_build_1527.Dm_build_880() > 0 {
		if !dm_build_1526.isReadOver() && dm_build_1527.Dm_build_880() < Dm_build_117 {
			dm_build_1526.read(dm_build_1527)
		}
		if dm_build_1527.Dm_build_880() > Dm_build_117 {
			dm_build_1528 = Dm_build_117
		} else {
			dm_build_1528 = dm_build_1527.Dm_build_880()
		}

		putData := dm_build_482(dm_build_1522, dm_build_1523, int16(dm_build_1525), dm_build_1527, int32(dm_build_1528))
		_, err := dm_build_1522.dm_build_1324(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1530 *dm_build_1285) dm_build_1529() ([]byte, error) {
	var dm_build_1531 error
	if dm_build_1530.dm_build_1293 == nil {
		if dm_build_1530.dm_build_1293, dm_build_1531 = security.NewClientKeyPair(); dm_build_1531 != nil {
			return nil, dm_build_1531
		}
	}
	return security.Bn2Bytes(dm_build_1530.dm_build_1293.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_1533 *dm_build_1285) dm_build_1532() (*security.DhKey, error) {
	var dm_build_1534 error
	if dm_build_1533.dm_build_1293 == nil {
		if dm_build_1533.dm_build_1293, dm_build_1534 = security.NewClientKeyPair(); dm_build_1534 != nil {
			return nil, dm_build_1534
		}
	}
	return dm_build_1533.dm_build_1293, nil
}

func (dm_build_1536 *dm_build_1285) dm_build_1535(dm_build_1537 int, dm_build_1538 []byte, dm_build_1539 string, dm_build_1540 int) (dm_build_1541 error) {
	if dm_build_1537 > 0 && dm_build_1537 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_1538 != nil {
		dm_build_1536.dm_build_1290, dm_build_1541 = security.NewSymmCipher(dm_build_1537, dm_build_1538)
	} else if dm_build_1537 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_1536.dm_build_1290, dm_build_1541 = security.NewThirdPartCipher(dm_build_1537, dm_build_1538, dm_build_1539, dm_build_1540); dm_build_1541 != nil {
			dm_build_1541 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_1541.Error()).throw()
		}
	}
	return
}

func (dm_build_1543 *dm_build_1285) dm_build_1542(dm_build_1544 bool) (dm_build_1545 error) {
	if dm_build_1543.dm_build_1287, dm_build_1545 = security.NewTLSFromTCP(dm_build_1543.dm_build_1286, dm_build_1543.dm_build_1289.dmConnector.sslCertPath, dm_build_1543.dm_build_1289.dmConnector.sslKeyPath, dm_build_1543.dm_build_1289.dmConnector.user); dm_build_1545 != nil {
		return
	}
	if !dm_build_1544 {
		dm_build_1543.dm_build_1287 = nil
	}
	return
}

func (dm_build_1547 *dm_build_1285) dm_build_1546(dm_build_1548 dm_build_123) bool {
	return dm_build_1548.dm_build_137() != Dm_build_33 && dm_build_1547.dm_build_1289.sslEncrypt == 1
}
