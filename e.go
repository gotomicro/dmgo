/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1256 struct{}

var Dm_build_1257 = &dm_build_1256{}

func (Dm_build_1259 *dm_build_1256) Dm_build_1258(dm_build_1260 []byte, dm_build_1261 int, dm_build_1262 byte) int {
	dm_build_1260[dm_build_1261] = dm_build_1262
	return 1
}

func (Dm_build_1264 *dm_build_1256) Dm_build_1263(dm_build_1265 []byte, dm_build_1266 int, dm_build_1267 int8) int {
	dm_build_1265[dm_build_1266] = byte(dm_build_1267)
	return 1
}

func (Dm_build_1269 *dm_build_1256) Dm_build_1268(dm_build_1270 []byte, dm_build_1271 int, dm_build_1272 int16) int {
	dm_build_1270[dm_build_1271] = byte(dm_build_1272)
	dm_build_1271++
	dm_build_1270[dm_build_1271] = byte(dm_build_1272 >> 8)
	return 2
}

func (Dm_build_1274 *dm_build_1256) Dm_build_1273(dm_build_1275 []byte, dm_build_1276 int, dm_build_1277 int32) int {
	dm_build_1275[dm_build_1276] = byte(dm_build_1277)
	dm_build_1276++
	dm_build_1275[dm_build_1276] = byte(dm_build_1277 >> 8)
	dm_build_1276++
	dm_build_1275[dm_build_1276] = byte(dm_build_1277 >> 16)
	dm_build_1276++
	dm_build_1275[dm_build_1276] = byte(dm_build_1277 >> 24)
	dm_build_1276++
	return 4
}

func (Dm_build_1279 *dm_build_1256) Dm_build_1278(dm_build_1280 []byte, dm_build_1281 int, dm_build_1282 int64) int {
	dm_build_1280[dm_build_1281] = byte(dm_build_1282)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 8)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 16)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 24)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 32)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 40)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 48)
	dm_build_1281++
	dm_build_1280[dm_build_1281] = byte(dm_build_1282 >> 56)
	return 8
}

func (Dm_build_1284 *dm_build_1256) Dm_build_1283(dm_build_1285 []byte, dm_build_1286 int, dm_build_1287 float32) int {
	return Dm_build_1284.Dm_build_1303(dm_build_1285, dm_build_1286, math.Float32bits(dm_build_1287))
}

func (Dm_build_1289 *dm_build_1256) Dm_build_1288(dm_build_1290 []byte, dm_build_1291 int, dm_build_1292 float64) int {
	return Dm_build_1289.Dm_build_1308(dm_build_1290, dm_build_1291, math.Float64bits(dm_build_1292))
}

func (Dm_build_1294 *dm_build_1256) Dm_build_1293(dm_build_1295 []byte, dm_build_1296 int, dm_build_1297 uint8) int {
	dm_build_1295[dm_build_1296] = byte(dm_build_1297)
	return 1
}

func (Dm_build_1299 *dm_build_1256) Dm_build_1298(dm_build_1300 []byte, dm_build_1301 int, dm_build_1302 uint16) int {
	dm_build_1300[dm_build_1301] = byte(dm_build_1302)
	dm_build_1301++
	dm_build_1300[dm_build_1301] = byte(dm_build_1302 >> 8)
	return 2
}

func (Dm_build_1304 *dm_build_1256) Dm_build_1303(dm_build_1305 []byte, dm_build_1306 int, dm_build_1307 uint32) int {
	dm_build_1305[dm_build_1306] = byte(dm_build_1307)
	dm_build_1306++
	dm_build_1305[dm_build_1306] = byte(dm_build_1307 >> 8)
	dm_build_1306++
	dm_build_1305[dm_build_1306] = byte(dm_build_1307 >> 16)
	dm_build_1306++
	dm_build_1305[dm_build_1306] = byte(dm_build_1307 >> 24)
	return 3
}

func (Dm_build_1309 *dm_build_1256) Dm_build_1308(dm_build_1310 []byte, dm_build_1311 int, dm_build_1312 uint64) int {
	dm_build_1310[dm_build_1311] = byte(dm_build_1312)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 8)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 16)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 24)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 32)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 40)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 48)
	dm_build_1311++
	dm_build_1310[dm_build_1311] = byte(dm_build_1312 >> 56)
	return 3
}

func (Dm_build_1314 *dm_build_1256) Dm_build_1313(dm_build_1315 []byte, dm_build_1316 int, dm_build_1317 []byte, dm_build_1318 int, dm_build_1319 int) int {
	copy(dm_build_1315[dm_build_1316:dm_build_1316+dm_build_1319], dm_build_1317[dm_build_1318:dm_build_1318+dm_build_1319])
	return dm_build_1319
}

func (Dm_build_1321 *dm_build_1256) Dm_build_1320(dm_build_1322 []byte, dm_build_1323 int, dm_build_1324 []byte, dm_build_1325 int, dm_build_1326 int) int {
	dm_build_1323 += Dm_build_1321.Dm_build_1303(dm_build_1322, dm_build_1323, uint32(dm_build_1326))
	return 4 + Dm_build_1321.Dm_build_1313(dm_build_1322, dm_build_1323, dm_build_1324, dm_build_1325, dm_build_1326)
}

func (Dm_build_1328 *dm_build_1256) Dm_build_1327(dm_build_1329 []byte, dm_build_1330 int, dm_build_1331 []byte, dm_build_1332 int, dm_build_1333 int) int {
	dm_build_1330 += Dm_build_1328.Dm_build_1298(dm_build_1329, dm_build_1330, uint16(dm_build_1333))
	return 2 + Dm_build_1328.Dm_build_1313(dm_build_1329, dm_build_1330, dm_build_1331, dm_build_1332, dm_build_1333)
}

func (Dm_build_1335 *dm_build_1256) Dm_build_1334(dm_build_1336 []byte, dm_build_1337 int, dm_build_1338 string, dm_build_1339 string, dm_build_1340 *DmConnection) int {
	dm_build_1341 := Dm_build_1335.Dm_build_1473(dm_build_1338, dm_build_1339, dm_build_1340)
	dm_build_1337 += Dm_build_1335.Dm_build_1303(dm_build_1336, dm_build_1337, uint32(len(dm_build_1341)))
	return 4 + Dm_build_1335.Dm_build_1313(dm_build_1336, dm_build_1337, dm_build_1341, 0, len(dm_build_1341))
}

func (Dm_build_1343 *dm_build_1256) Dm_build_1342(dm_build_1344 []byte, dm_build_1345 int, dm_build_1346 string, dm_build_1347 string, dm_build_1348 *DmConnection) int {
	dm_build_1349 := Dm_build_1343.Dm_build_1473(dm_build_1346, dm_build_1347, dm_build_1348)

	dm_build_1345 += Dm_build_1343.Dm_build_1298(dm_build_1344, dm_build_1345, uint16(len(dm_build_1349)))
	return 2 + Dm_build_1343.Dm_build_1313(dm_build_1344, dm_build_1345, dm_build_1349, 0, len(dm_build_1349))
}

func (Dm_build_1351 *dm_build_1256) Dm_build_1350(dm_build_1352 []byte, dm_build_1353 int) byte {
	return dm_build_1352[dm_build_1353]
}

func (Dm_build_1355 *dm_build_1256) Dm_build_1354(dm_build_1356 []byte, dm_build_1357 int) int16 {
	var dm_build_1358 int16
	dm_build_1358 = int16(dm_build_1356[dm_build_1357] & 0xff)
	dm_build_1357++
	dm_build_1358 |= int16(dm_build_1356[dm_build_1357]&0xff) << 8
	return dm_build_1358
}

func (Dm_build_1360 *dm_build_1256) Dm_build_1359(dm_build_1361 []byte, dm_build_1362 int) int32 {
	var dm_build_1363 int32
	dm_build_1363 = int32(dm_build_1361[dm_build_1362] & 0xff)
	dm_build_1362++
	dm_build_1363 |= int32(dm_build_1361[dm_build_1362]&0xff) << 8
	dm_build_1362++
	dm_build_1363 |= int32(dm_build_1361[dm_build_1362]&0xff) << 16
	dm_build_1362++
	dm_build_1363 |= int32(dm_build_1361[dm_build_1362]&0xff) << 24
	return dm_build_1363
}

func (Dm_build_1365 *dm_build_1256) Dm_build_1364(dm_build_1366 []byte, dm_build_1367 int) int64 {
	var dm_build_1368 int64
	dm_build_1368 = int64(dm_build_1366[dm_build_1367] & 0xff)
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 8
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 16
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 24
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 32
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 40
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 48
	dm_build_1367++
	dm_build_1368 |= int64(dm_build_1366[dm_build_1367]&0xff) << 56
	return dm_build_1368
}

func (Dm_build_1370 *dm_build_1256) Dm_build_1369(dm_build_1371 []byte, dm_build_1372 int) float32 {
	return math.Float32frombits(Dm_build_1370.Dm_build_1386(dm_build_1371, dm_build_1372))
}

func (Dm_build_1374 *dm_build_1256) Dm_build_1373(dm_build_1375 []byte, dm_build_1376 int) float64 {
	return math.Float64frombits(Dm_build_1374.Dm_build_1391(dm_build_1375, dm_build_1376))
}

func (Dm_build_1378 *dm_build_1256) Dm_build_1377(dm_build_1379 []byte, dm_build_1380 int) uint8 {
	return uint8(dm_build_1379[dm_build_1380] & 0xff)
}

func (Dm_build_1382 *dm_build_1256) Dm_build_1381(dm_build_1383 []byte, dm_build_1384 int) uint16 {
	var dm_build_1385 uint16
	dm_build_1385 = uint16(dm_build_1383[dm_build_1384] & 0xff)
	dm_build_1384++
	dm_build_1385 |= uint16(dm_build_1383[dm_build_1384]&0xff) << 8
	return dm_build_1385
}

func (Dm_build_1387 *dm_build_1256) Dm_build_1386(dm_build_1388 []byte, dm_build_1389 int) uint32 {
	var dm_build_1390 uint32
	dm_build_1390 = uint32(dm_build_1388[dm_build_1389] & 0xff)
	dm_build_1389++
	dm_build_1390 |= uint32(dm_build_1388[dm_build_1389]&0xff) << 8
	dm_build_1389++
	dm_build_1390 |= uint32(dm_build_1388[dm_build_1389]&0xff) << 16
	dm_build_1389++
	dm_build_1390 |= uint32(dm_build_1388[dm_build_1389]&0xff) << 24
	return dm_build_1390
}

func (Dm_build_1392 *dm_build_1256) Dm_build_1391(dm_build_1393 []byte, dm_build_1394 int) uint64 {
	var dm_build_1395 uint64
	dm_build_1395 = uint64(dm_build_1393[dm_build_1394] & 0xff)
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 8
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 16
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 24
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 32
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 40
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 48
	dm_build_1394++
	dm_build_1395 |= uint64(dm_build_1393[dm_build_1394]&0xff) << 56
	return dm_build_1395
}

func (Dm_build_1397 *dm_build_1256) Dm_build_1396(dm_build_1398 []byte, dm_build_1399 int) []byte {
	dm_build_1400 := Dm_build_1397.Dm_build_1386(dm_build_1398, dm_build_1399)

	dm_build_1401 := make([]byte, dm_build_1400)
	copy(dm_build_1401[:int(dm_build_1400)], dm_build_1398[dm_build_1399+4:dm_build_1399+4+int(dm_build_1400)])
	return dm_build_1401
}

func (Dm_build_1403 *dm_build_1256) Dm_build_1402(dm_build_1404 []byte, dm_build_1405 int) []byte {
	dm_build_1406 := Dm_build_1403.Dm_build_1381(dm_build_1404, dm_build_1405)

	dm_build_1407 := make([]byte, dm_build_1406)
	copy(dm_build_1407[:int(dm_build_1406)], dm_build_1404[dm_build_1405+2:dm_build_1405+2+int(dm_build_1406)])
	return dm_build_1407
}

func (Dm_build_1409 *dm_build_1256) Dm_build_1408(dm_build_1410 []byte, dm_build_1411 int, dm_build_1412 int) []byte {

	dm_build_1413 := make([]byte, dm_build_1412)
	copy(dm_build_1413[:dm_build_1412], dm_build_1410[dm_build_1411:dm_build_1411+dm_build_1412])
	return dm_build_1413
}

func (Dm_build_1415 *dm_build_1256) Dm_build_1414(dm_build_1416 []byte, dm_build_1417 int, dm_build_1418 int, dm_build_1419 string, dm_build_1420 *DmConnection) string {
	return Dm_build_1415.Dm_build_1509(dm_build_1416[dm_build_1417:dm_build_1417+dm_build_1418], dm_build_1419, dm_build_1420)
}

func (Dm_build_1422 *dm_build_1256) Dm_build_1421(dm_build_1423 []byte, dm_build_1424 int, dm_build_1425 string, dm_build_1426 *DmConnection) string {
	dm_build_1427 := Dm_build_1422.Dm_build_1386(dm_build_1423, dm_build_1424)
	dm_build_1424 += 4
	return Dm_build_1422.Dm_build_1414(dm_build_1423, dm_build_1424, int(dm_build_1427), dm_build_1425, dm_build_1426)
}

func (Dm_build_1429 *dm_build_1256) Dm_build_1428(dm_build_1430 []byte, dm_build_1431 int, dm_build_1432 string, dm_build_1433 *DmConnection) string {
	dm_build_1434 := Dm_build_1429.Dm_build_1381(dm_build_1430, dm_build_1431)
	dm_build_1431 += 2
	return Dm_build_1429.Dm_build_1414(dm_build_1430, dm_build_1431, int(dm_build_1434), dm_build_1432, dm_build_1433)
}

func (Dm_build_1436 *dm_build_1256) Dm_build_1435(dm_build_1437 byte) []byte {
	return []byte{dm_build_1437}
}

func (Dm_build_1439 *dm_build_1256) Dm_build_1438(dm_build_1440 int8) []byte {
	return []byte{byte(dm_build_1440)}
}

func (Dm_build_1442 *dm_build_1256) Dm_build_1441(dm_build_1443 int16) []byte {
	return []byte{byte(dm_build_1443), byte(dm_build_1443 >> 8)}
}

func (Dm_build_1445 *dm_build_1256) Dm_build_1444(dm_build_1446 int32) []byte {
	return []byte{byte(dm_build_1446), byte(dm_build_1446 >> 8), byte(dm_build_1446 >> 16), byte(dm_build_1446 >> 24)}
}

func (Dm_build_1448 *dm_build_1256) Dm_build_1447(dm_build_1449 int64) []byte {
	return []byte{byte(dm_build_1449), byte(dm_build_1449 >> 8), byte(dm_build_1449 >> 16), byte(dm_build_1449 >> 24), byte(dm_build_1449 >> 32),
		byte(dm_build_1449 >> 40), byte(dm_build_1449 >> 48), byte(dm_build_1449 >> 56)}
}

func (Dm_build_1451 *dm_build_1256) Dm_build_1450(dm_build_1452 float32) []byte {
	return Dm_build_1451.Dm_build_1462(math.Float32bits(dm_build_1452))
}

func (Dm_build_1454 *dm_build_1256) Dm_build_1453(dm_build_1455 float64) []byte {
	return Dm_build_1454.Dm_build_1465(math.Float64bits(dm_build_1455))
}

func (Dm_build_1457 *dm_build_1256) Dm_build_1456(dm_build_1458 uint8) []byte {
	return []byte{byte(dm_build_1458)}
}

func (Dm_build_1460 *dm_build_1256) Dm_build_1459(dm_build_1461 uint16) []byte {
	return []byte{byte(dm_build_1461), byte(dm_build_1461 >> 8)}
}

func (Dm_build_1463 *dm_build_1256) Dm_build_1462(dm_build_1464 uint32) []byte {
	return []byte{byte(dm_build_1464), byte(dm_build_1464 >> 8), byte(dm_build_1464 >> 16), byte(dm_build_1464 >> 24)}
}

func (Dm_build_1466 *dm_build_1256) Dm_build_1465(dm_build_1467 uint64) []byte {
	return []byte{byte(dm_build_1467), byte(dm_build_1467 >> 8), byte(dm_build_1467 >> 16), byte(dm_build_1467 >> 24), byte(dm_build_1467 >> 32), byte(dm_build_1467 >> 40), byte(dm_build_1467 >> 48), byte(dm_build_1467 >> 56)}
}

func (Dm_build_1469 *dm_build_1256) Dm_build_1468(dm_build_1470 []byte, dm_build_1471 string, dm_build_1472 *DmConnection) []byte {
	if dm_build_1471 == "UTF-8" {
		return dm_build_1470
	}

	if dm_build_1472 == nil {
		if e := dm_build_1514(dm_build_1471); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1470), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1472.encodeBuffer == nil {
		dm_build_1472.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1472.encode = dm_build_1514(dm_build_1472.getServerEncoding())
		dm_build_1472.transformReaderDst = make([]byte, 4096)
		dm_build_1472.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1472.encode; e != nil {

		dm_build_1472.encodeBuffer.Reset()

		n, err := dm_build_1472.encodeBuffer.ReadFrom(
			Dm_build_1528(bytes.NewReader(dm_build_1470), e.NewEncoder(), dm_build_1472.transformReaderDst, dm_build_1472.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1472.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1474 *dm_build_1256) Dm_build_1473(dm_build_1475 string, dm_build_1476 string, dm_build_1477 *DmConnection) []byte {
	return Dm_build_1474.Dm_build_1468([]byte(dm_build_1475), dm_build_1476, dm_build_1477)
}

func (Dm_build_1479 *dm_build_1256) Dm_build_1478(dm_build_1480 []byte) byte {
	return Dm_build_1479.Dm_build_1350(dm_build_1480, 0)
}

func (Dm_build_1482 *dm_build_1256) Dm_build_1481(dm_build_1483 []byte) int16 {
	return Dm_build_1482.Dm_build_1354(dm_build_1483, 0)
}

func (Dm_build_1485 *dm_build_1256) Dm_build_1484(dm_build_1486 []byte) int32 {
	return Dm_build_1485.Dm_build_1359(dm_build_1486, 0)
}

func (Dm_build_1488 *dm_build_1256) Dm_build_1487(dm_build_1489 []byte) int64 {
	return Dm_build_1488.Dm_build_1364(dm_build_1489, 0)
}

func (Dm_build_1491 *dm_build_1256) Dm_build_1490(dm_build_1492 []byte) float32 {
	return Dm_build_1491.Dm_build_1369(dm_build_1492, 0)
}

func (Dm_build_1494 *dm_build_1256) Dm_build_1493(dm_build_1495 []byte) float64 {
	return Dm_build_1494.Dm_build_1373(dm_build_1495, 0)
}

func (Dm_build_1497 *dm_build_1256) Dm_build_1496(dm_build_1498 []byte) uint8 {
	return Dm_build_1497.Dm_build_1377(dm_build_1498, 0)
}

func (Dm_build_1500 *dm_build_1256) Dm_build_1499(dm_build_1501 []byte) uint16 {
	return Dm_build_1500.Dm_build_1381(dm_build_1501, 0)
}

func (Dm_build_1503 *dm_build_1256) Dm_build_1502(dm_build_1504 []byte) uint32 {
	return Dm_build_1503.Dm_build_1386(dm_build_1504, 0)
}

func (Dm_build_1506 *dm_build_1256) Dm_build_1505(dm_build_1507 []byte, dm_build_1508 string) []byte {
	if dm_build_1508 == "UTF-8" {
		return dm_build_1507
	}

	if e := dm_build_1514(dm_build_1508); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1507), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1510 *dm_build_1256) Dm_build_1509(dm_build_1511 []byte, dm_build_1512 string, dm_build_1513 *DmConnection) string {
	return string(Dm_build_1510.Dm_build_1505(dm_build_1511, dm_build_1512))
}

func dm_build_1514(dm_build_1515 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1515); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1516 struct {
	dm_build_1517 io.Reader
	dm_build_1518 transform.Transformer
	dm_build_1519 error

	dm_build_1520                []byte
	dm_build_1521, dm_build_1522 int

	dm_build_1523                []byte
	dm_build_1524, dm_build_1525 int

	dm_build_1526 bool
}

const dm_build_1527 = 4096

func Dm_build_1528(dm_build_1529 io.Reader, dm_build_1530 transform.Transformer, dm_build_1531 []byte, dm_build_1532 []byte) *Dm_build_1516 {
	dm_build_1530.Reset()
	return &Dm_build_1516{
		dm_build_1517: dm_build_1529,
		dm_build_1518: dm_build_1530,
		dm_build_1520: dm_build_1531,
		dm_build_1523: dm_build_1532,
	}
}

func (dm_build_1534 *Dm_build_1516) Read(dm_build_1535 []byte) (int, error) {
	dm_build_1536, dm_build_1537 := 0, error(nil)
	for {

		if dm_build_1534.dm_build_1521 != dm_build_1534.dm_build_1522 {
			dm_build_1536 = copy(dm_build_1535, dm_build_1534.dm_build_1520[dm_build_1534.dm_build_1521:dm_build_1534.dm_build_1522])
			dm_build_1534.dm_build_1521 += dm_build_1536
			if dm_build_1534.dm_build_1521 == dm_build_1534.dm_build_1522 && dm_build_1534.dm_build_1526 {
				return dm_build_1536, dm_build_1534.dm_build_1519
			}
			return dm_build_1536, nil
		} else if dm_build_1534.dm_build_1526 {
			return 0, dm_build_1534.dm_build_1519
		}

		if dm_build_1534.dm_build_1524 != dm_build_1534.dm_build_1525 || dm_build_1534.dm_build_1519 != nil {
			dm_build_1534.dm_build_1521 = 0
			dm_build_1534.dm_build_1522, dm_build_1536, dm_build_1537 = dm_build_1534.dm_build_1518.Transform(dm_build_1534.dm_build_1520, dm_build_1534.dm_build_1523[dm_build_1534.dm_build_1524:dm_build_1534.dm_build_1525], dm_build_1534.dm_build_1519 == io.EOF)
			dm_build_1534.dm_build_1524 += dm_build_1536

			switch {
			case dm_build_1537 == nil:
				if dm_build_1534.dm_build_1524 != dm_build_1534.dm_build_1525 {
					dm_build_1534.dm_build_1519 = nil
				}

				dm_build_1534.dm_build_1526 = dm_build_1534.dm_build_1519 != nil
				continue
			case dm_build_1537 == transform.ErrShortDst && (dm_build_1534.dm_build_1522 != 0 || dm_build_1536 != 0):

				continue
			case dm_build_1537 == transform.ErrShortSrc && dm_build_1534.dm_build_1525-dm_build_1534.dm_build_1524 != len(dm_build_1534.dm_build_1523) && dm_build_1534.dm_build_1519 == nil:

			default:
				dm_build_1534.dm_build_1526 = true

				if dm_build_1534.dm_build_1519 == nil || dm_build_1534.dm_build_1519 == io.EOF {
					dm_build_1534.dm_build_1519 = dm_build_1537
				}
				continue
			}
		}

		if dm_build_1534.dm_build_1524 != 0 {
			dm_build_1534.dm_build_1524, dm_build_1534.dm_build_1525 = 0, copy(dm_build_1534.dm_build_1523, dm_build_1534.dm_build_1523[dm_build_1534.dm_build_1524:dm_build_1534.dm_build_1525])
		}
		dm_build_1536, dm_build_1534.dm_build_1519 = dm_build_1534.dm_build_1517.Read(dm_build_1534.dm_build_1523[dm_build_1534.dm_build_1525:])
		dm_build_1534.dm_build_1525 += dm_build_1536
	}
}
