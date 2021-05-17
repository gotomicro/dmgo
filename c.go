/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1239 struct {
	dm_build_1240 []byte
	dm_build_1241 int
}

func Dm_build_1242(dm_build_1243 int) *Dm_build_1239 {
	return &Dm_build_1239{make([]byte, 0, dm_build_1243), 0}
}

func Dm_build_1244(dm_build_1245 []byte) *Dm_build_1239 {
	return &Dm_build_1239{dm_build_1245, 0}
}

func (dm_build_1247 *Dm_build_1239) dm_build_1246(dm_build_1248 int) *Dm_build_1239 {

	dm_build_1249 := len(dm_build_1247.dm_build_1240)
	dm_build_1250 := cap(dm_build_1247.dm_build_1240)

	if dm_build_1249+dm_build_1248 <= dm_build_1250 {
		dm_build_1247.dm_build_1240 = dm_build_1247.dm_build_1240[:dm_build_1249+dm_build_1248]
	} else {
		remain := dm_build_1248 + dm_build_1249 - dm_build_1250
		nbuf := make([]byte, dm_build_1248+dm_build_1249, 2*dm_build_1250+remain)
		copy(nbuf, dm_build_1247.dm_build_1240)
		dm_build_1247.dm_build_1240 = nbuf
	}

	return dm_build_1247
}

func (dm_build_1252 *Dm_build_1239) Dm_build_1251() int {
	return len(dm_build_1252.dm_build_1240)
}

func (dm_build_1254 *Dm_build_1239) Dm_build_1253(dm_build_1255 int) *Dm_build_1239 {
	for i := dm_build_1255; i < len(dm_build_1254.dm_build_1240); i++ {
		dm_build_1254.dm_build_1240[i] = 0
	}
	dm_build_1254.dm_build_1240 = dm_build_1254.dm_build_1240[:dm_build_1255]
	return dm_build_1254
}

func (dm_build_1257 *Dm_build_1239) Dm_build_1256(dm_build_1258 int) *Dm_build_1239 {
	dm_build_1257.dm_build_1241 = dm_build_1258
	return dm_build_1257
}

func (dm_build_1260 *Dm_build_1239) Dm_build_1259() int {
	return dm_build_1260.dm_build_1241
}

func (dm_build_1262 *Dm_build_1239) Dm_build_1261(dm_build_1263 bool) int {
	return len(dm_build_1262.dm_build_1240) - dm_build_1262.dm_build_1241
}

func (dm_build_1265 *Dm_build_1239) Dm_build_1264(dm_build_1266 int, dm_build_1267 bool, dm_build_1268 bool) *Dm_build_1239 {

	if dm_build_1267 {
		if dm_build_1268 {
			dm_build_1265.dm_build_1246(dm_build_1266)
		} else {
			dm_build_1265.dm_build_1240 = dm_build_1265.dm_build_1240[:len(dm_build_1265.dm_build_1240)-dm_build_1266]
		}
	} else {
		if dm_build_1268 {
			dm_build_1265.dm_build_1241 += dm_build_1266
		} else {
			dm_build_1265.dm_build_1241 -= dm_build_1266
		}
	}

	return dm_build_1265
}

func (dm_build_1270 *Dm_build_1239) Dm_build_1269(dm_build_1271 io.Reader, dm_build_1272 int) int {
	dm_build_1273 := len(dm_build_1270.dm_build_1240)
	dm_build_1270.dm_build_1246(dm_build_1272)
	dm_build_1274 := 0
	for dm_build_1272 > 0 {
		n, err := dm_build_1271.Read(dm_build_1270.dm_build_1240[dm_build_1273+dm_build_1274:])
		if n > 0 && err == io.EOF {
			dm_build_1274 += n
			dm_build_1270.dm_build_1240 = dm_build_1270.dm_build_1240[:dm_build_1273+dm_build_1274]
			return dm_build_1274
		} else if n > 0 && err == nil {
			dm_build_1272 -= n
			dm_build_1274 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_1274
}

func (dm_build_1276 *Dm_build_1239) Dm_build_1275(dm_build_1277 io.Writer) *Dm_build_1239 {
	dm_build_1277.Write(dm_build_1276.dm_build_1240)
	return dm_build_1276
}

func (dm_build_1279 *Dm_build_1239) Dm_build_1278(dm_build_1280 bool) int {
	dm_build_1281 := len(dm_build_1279.dm_build_1240)
	dm_build_1279.dm_build_1246(1)

	if dm_build_1280 {
		return copy(dm_build_1279.dm_build_1240[dm_build_1281:], []byte{1})
	} else {
		return copy(dm_build_1279.dm_build_1240[dm_build_1281:], []byte{0})
	}
}

func (dm_build_1283 *Dm_build_1239) Dm_build_1282(dm_build_1284 byte) int {
	dm_build_1285 := len(dm_build_1283.dm_build_1240)
	dm_build_1283.dm_build_1246(1)

	return copy(dm_build_1283.dm_build_1240[dm_build_1285:], Dm_build_885.Dm_build_1060(dm_build_1284))
}

func (dm_build_1287 *Dm_build_1239) Dm_build_1286(dm_build_1288 int16) int {
	dm_build_1289 := len(dm_build_1287.dm_build_1240)
	dm_build_1287.dm_build_1246(2)

	return copy(dm_build_1287.dm_build_1240[dm_build_1289:], Dm_build_885.Dm_build_1063(dm_build_1288))
}

func (dm_build_1291 *Dm_build_1239) Dm_build_1290(dm_build_1292 int32) int {
	dm_build_1293 := len(dm_build_1291.dm_build_1240)
	dm_build_1291.dm_build_1246(4)

	return copy(dm_build_1291.dm_build_1240[dm_build_1293:], Dm_build_885.Dm_build_1066(dm_build_1292))
}

func (dm_build_1295 *Dm_build_1239) Dm_build_1294(dm_build_1296 uint8) int {
	dm_build_1297 := len(dm_build_1295.dm_build_1240)
	dm_build_1295.dm_build_1246(1)

	return copy(dm_build_1295.dm_build_1240[dm_build_1297:], Dm_build_885.Dm_build_1078(dm_build_1296))
}

func (dm_build_1299 *Dm_build_1239) Dm_build_1298(dm_build_1300 uint16) int {
	dm_build_1301 := len(dm_build_1299.dm_build_1240)
	dm_build_1299.dm_build_1246(2)

	return copy(dm_build_1299.dm_build_1240[dm_build_1301:], Dm_build_885.Dm_build_1081(dm_build_1300))
}

func (dm_build_1303 *Dm_build_1239) Dm_build_1302(dm_build_1304 uint32) int {
	dm_build_1305 := len(dm_build_1303.dm_build_1240)
	dm_build_1303.dm_build_1246(4)

	return copy(dm_build_1303.dm_build_1240[dm_build_1305:], Dm_build_885.Dm_build_1084(dm_build_1304))
}

func (dm_build_1307 *Dm_build_1239) Dm_build_1306(dm_build_1308 uint64) int {
	dm_build_1309 := len(dm_build_1307.dm_build_1240)
	dm_build_1307.dm_build_1246(8)

	return copy(dm_build_1307.dm_build_1240[dm_build_1309:], Dm_build_885.Dm_build_1087(dm_build_1308))
}

func (dm_build_1311 *Dm_build_1239) Dm_build_1310(dm_build_1312 float32) int {
	dm_build_1313 := len(dm_build_1311.dm_build_1240)
	dm_build_1311.dm_build_1246(4)

	return copy(dm_build_1311.dm_build_1240[dm_build_1313:], Dm_build_885.Dm_build_1084(math.Float32bits(dm_build_1312)))
}

func (dm_build_1315 *Dm_build_1239) Dm_build_1314(dm_build_1316 float64) int {
	dm_build_1317 := len(dm_build_1315.dm_build_1240)
	dm_build_1315.dm_build_1246(8)

	return copy(dm_build_1315.dm_build_1240[dm_build_1317:], Dm_build_885.Dm_build_1087(math.Float64bits(dm_build_1316)))
}

func (dm_build_1319 *Dm_build_1239) Dm_build_1318(dm_build_1320 []byte) int {
	dm_build_1321 := len(dm_build_1319.dm_build_1240)
	dm_build_1319.dm_build_1246(len(dm_build_1320))
	return copy(dm_build_1319.dm_build_1240[dm_build_1321:], dm_build_1320)
}

func (dm_build_1323 *Dm_build_1239) Dm_build_1322(dm_build_1324 []byte) int {
	return dm_build_1323.Dm_build_1290(int32(len(dm_build_1324))) + dm_build_1323.Dm_build_1318(dm_build_1324)
}

func (dm_build_1326 *Dm_build_1239) Dm_build_1325(dm_build_1327 []byte) int {
	return dm_build_1326.Dm_build_1294(uint8(len(dm_build_1327))) + dm_build_1326.Dm_build_1318(dm_build_1327)
}

func (dm_build_1329 *Dm_build_1239) Dm_build_1328(dm_build_1330 []byte) int {
	return dm_build_1329.Dm_build_1298(uint16(len(dm_build_1330))) + dm_build_1329.Dm_build_1318(dm_build_1330)
}

func (dm_build_1332 *Dm_build_1239) Dm_build_1331(dm_build_1333 []byte) int {
	return dm_build_1332.Dm_build_1318(dm_build_1333) + dm_build_1332.Dm_build_1282(0)
}

func (dm_build_1335 *Dm_build_1239) Dm_build_1334(dm_build_1336 string, dm_build_1337 string, dm_build_1338 *DmConnection) int {
	dm_build_1339 := Dm_build_885.Dm_build_1095(dm_build_1336, dm_build_1337, dm_build_1338)
	return dm_build_1335.Dm_build_1322(dm_build_1339)
}

func (dm_build_1341 *Dm_build_1239) Dm_build_1340(dm_build_1342 string, dm_build_1343 string, dm_build_1344 *DmConnection) int {
	dm_build_1345 := Dm_build_885.Dm_build_1095(dm_build_1342, dm_build_1343, dm_build_1344)
	return dm_build_1341.Dm_build_1325(dm_build_1345)
}

func (dm_build_1347 *Dm_build_1239) Dm_build_1346(dm_build_1348 string, dm_build_1349 string, dm_build_1350 *DmConnection) int {
	dm_build_1351 := Dm_build_885.Dm_build_1095(dm_build_1348, dm_build_1349, dm_build_1350)
	return dm_build_1347.Dm_build_1328(dm_build_1351)
}

func (dm_build_1353 *Dm_build_1239) Dm_build_1352(dm_build_1354 string, dm_build_1355 string, dm_build_1356 *DmConnection) int {
	dm_build_1357 := Dm_build_885.Dm_build_1095(dm_build_1354, dm_build_1355, dm_build_1356)
	return dm_build_1353.Dm_build_1331(dm_build_1357)
}

func (dm_build_1359 *Dm_build_1239) Dm_build_1358() byte {
	dm_build_1360 := Dm_build_885.Dm_build_978(dm_build_1359.dm_build_1240, dm_build_1359.dm_build_1241)
	dm_build_1359.dm_build_1241++
	return dm_build_1360
}

func (dm_build_1362 *Dm_build_1239) Dm_build_1361() int16 {
	dm_build_1363 := Dm_build_885.Dm_build_982(dm_build_1362.dm_build_1240, dm_build_1362.dm_build_1241)
	dm_build_1362.dm_build_1241 += 2
	return dm_build_1363
}

func (dm_build_1365 *Dm_build_1239) Dm_build_1364() int32 {
	dm_build_1366 := Dm_build_885.Dm_build_987(dm_build_1365.dm_build_1240, dm_build_1365.dm_build_1241)
	dm_build_1365.dm_build_1241 += 4
	return dm_build_1366
}

func (dm_build_1368 *Dm_build_1239) Dm_build_1367() int64 {
	dm_build_1369 := Dm_build_885.Dm_build_992(dm_build_1368.dm_build_1240, dm_build_1368.dm_build_1241)
	dm_build_1368.dm_build_1241 += 8
	return dm_build_1369
}

func (dm_build_1371 *Dm_build_1239) Dm_build_1370() float32 {
	dm_build_1372 := Dm_build_885.Dm_build_997(dm_build_1371.dm_build_1240, dm_build_1371.dm_build_1241)
	dm_build_1371.dm_build_1241 += 4
	return dm_build_1372
}

func (dm_build_1374 *Dm_build_1239) Dm_build_1373() float64 {
	dm_build_1375 := Dm_build_885.Dm_build_1001(dm_build_1374.dm_build_1240, dm_build_1374.dm_build_1241)
	dm_build_1374.dm_build_1241 += 8
	return dm_build_1375
}

func (dm_build_1377 *Dm_build_1239) Dm_build_1376() uint8 {
	dm_build_1378 := Dm_build_885.Dm_build_1005(dm_build_1377.dm_build_1240, dm_build_1377.dm_build_1241)
	dm_build_1377.dm_build_1241 += 1
	return dm_build_1378
}

func (dm_build_1380 *Dm_build_1239) Dm_build_1379() uint16 {
	dm_build_1381 := Dm_build_885.Dm_build_1009(dm_build_1380.dm_build_1240, dm_build_1380.dm_build_1241)
	dm_build_1380.dm_build_1241 += 2
	return dm_build_1381
}

func (dm_build_1383 *Dm_build_1239) Dm_build_1382() uint32 {
	dm_build_1384 := Dm_build_885.Dm_build_1014(dm_build_1383.dm_build_1240, dm_build_1383.dm_build_1241)
	dm_build_1383.dm_build_1241 += 4
	return dm_build_1384
}

func (dm_build_1386 *Dm_build_1239) Dm_build_1385(dm_build_1387 int) []byte {
	dm_build_1388 := Dm_build_885.Dm_build_1034(dm_build_1386.dm_build_1240, dm_build_1386.dm_build_1241, dm_build_1387)
	dm_build_1386.dm_build_1241 += dm_build_1387
	return dm_build_1388
}

func (dm_build_1390 *Dm_build_1239) Dm_build_1389() []byte {
	return dm_build_1390.Dm_build_1385(int(dm_build_1390.Dm_build_1364()))
}

func (dm_build_1392 *Dm_build_1239) Dm_build_1391() []byte {
	return dm_build_1392.Dm_build_1385(int(dm_build_1392.Dm_build_1358()))
}

func (dm_build_1394 *Dm_build_1239) Dm_build_1393() []byte {
	return dm_build_1394.Dm_build_1385(int(dm_build_1394.Dm_build_1361()))
}

func (dm_build_1396 *Dm_build_1239) Dm_build_1395(dm_build_1397 int) []byte {
	return dm_build_1396.Dm_build_1385(dm_build_1397)
}

func (dm_build_1399 *Dm_build_1239) Dm_build_1398() []byte {
	dm_build_1400 := 0
	for dm_build_1399.Dm_build_1358() != 0 {
		dm_build_1400++
	}
	dm_build_1399.Dm_build_1264(dm_build_1400, false, false)
	return dm_build_1399.Dm_build_1385(dm_build_1400)
}

func (dm_build_1402 *Dm_build_1239) Dm_build_1401(dm_build_1403 int, dm_build_1404 string, dm_build_1405 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1402.Dm_build_1385(dm_build_1403), dm_build_1404, dm_build_1405)
}

func (dm_build_1407 *Dm_build_1239) Dm_build_1406(dm_build_1408 string, dm_build_1409 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1407.Dm_build_1389(), dm_build_1408, dm_build_1409)
}

func (dm_build_1411 *Dm_build_1239) Dm_build_1410(dm_build_1412 string, dm_build_1413 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1411.Dm_build_1391(), dm_build_1412, dm_build_1413)
}

func (dm_build_1415 *Dm_build_1239) Dm_build_1414(dm_build_1416 string, dm_build_1417 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1415.Dm_build_1393(), dm_build_1416, dm_build_1417)
}

func (dm_build_1419 *Dm_build_1239) Dm_build_1418(dm_build_1420 string, dm_build_1421 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1419.Dm_build_1398(), dm_build_1420, dm_build_1421)
}

func (dm_build_1423 *Dm_build_1239) Dm_build_1422(dm_build_1424 int, dm_build_1425 byte) int {
	return dm_build_1423.Dm_build_1458(dm_build_1424, Dm_build_885.Dm_build_1060(dm_build_1425))
}

func (dm_build_1427 *Dm_build_1239) Dm_build_1426(dm_build_1428 int, dm_build_1429 int16) int {
	return dm_build_1427.Dm_build_1458(dm_build_1428, Dm_build_885.Dm_build_1063(dm_build_1429))
}

func (dm_build_1431 *Dm_build_1239) Dm_build_1430(dm_build_1432 int, dm_build_1433 int32) int {
	return dm_build_1431.Dm_build_1458(dm_build_1432, Dm_build_885.Dm_build_1066(dm_build_1433))
}

func (dm_build_1435 *Dm_build_1239) Dm_build_1434(dm_build_1436 int, dm_build_1437 int64) int {
	return dm_build_1435.Dm_build_1458(dm_build_1436, Dm_build_885.Dm_build_1069(dm_build_1437))
}

func (dm_build_1439 *Dm_build_1239) Dm_build_1438(dm_build_1440 int, dm_build_1441 float32) int {
	return dm_build_1439.Dm_build_1458(dm_build_1440, Dm_build_885.Dm_build_1072(dm_build_1441))
}

func (dm_build_1443 *Dm_build_1239) Dm_build_1442(dm_build_1444 int, dm_build_1445 float64) int {
	return dm_build_1443.Dm_build_1458(dm_build_1444, Dm_build_885.Dm_build_1075(dm_build_1445))
}

func (dm_build_1447 *Dm_build_1239) Dm_build_1446(dm_build_1448 int, dm_build_1449 uint8) int {
	return dm_build_1447.Dm_build_1458(dm_build_1448, Dm_build_885.Dm_build_1078(dm_build_1449))
}

func (dm_build_1451 *Dm_build_1239) Dm_build_1450(dm_build_1452 int, dm_build_1453 uint16) int {
	return dm_build_1451.Dm_build_1458(dm_build_1452, Dm_build_885.Dm_build_1081(dm_build_1453))
}

func (dm_build_1455 *Dm_build_1239) Dm_build_1454(dm_build_1456 int, dm_build_1457 uint32) int {
	return dm_build_1455.Dm_build_1458(dm_build_1456, Dm_build_885.Dm_build_1084(dm_build_1457))
}

func (dm_build_1459 *Dm_build_1239) Dm_build_1458(dm_build_1460 int, dm_build_1461 []byte) int {
	return copy(dm_build_1459.dm_build_1240[dm_build_1460:], dm_build_1461)
}

func (dm_build_1463 *Dm_build_1239) Dm_build_1462(dm_build_1464 int, dm_build_1465 []byte) int {
	return dm_build_1463.Dm_build_1430(dm_build_1464, int32(len(dm_build_1465))) + dm_build_1463.Dm_build_1458(dm_build_1464+4, dm_build_1465)
}

func (dm_build_1467 *Dm_build_1239) Dm_build_1466(dm_build_1468 int, dm_build_1469 []byte) int {
	return dm_build_1467.Dm_build_1422(dm_build_1468, byte(len(dm_build_1469))) + dm_build_1467.Dm_build_1458(dm_build_1468+1, dm_build_1469)
}

func (dm_build_1471 *Dm_build_1239) Dm_build_1470(dm_build_1472 int, dm_build_1473 []byte) int {
	return dm_build_1471.Dm_build_1426(dm_build_1472, int16(len(dm_build_1473))) + dm_build_1471.Dm_build_1458(dm_build_1472+2, dm_build_1473)
}

func (dm_build_1475 *Dm_build_1239) Dm_build_1474(dm_build_1476 int, dm_build_1477 []byte) int {
	return dm_build_1475.Dm_build_1458(dm_build_1476, dm_build_1477) + dm_build_1475.Dm_build_1422(dm_build_1476+len(dm_build_1477), 0)
}

func (dm_build_1479 *Dm_build_1239) Dm_build_1478(dm_build_1480 int, dm_build_1481 string, dm_build_1482 string, dm_build_1483 *DmConnection) int {
	return dm_build_1479.Dm_build_1462(dm_build_1480, Dm_build_885.Dm_build_1095(dm_build_1481, dm_build_1482, dm_build_1483))
}

func (dm_build_1485 *Dm_build_1239) Dm_build_1484(dm_build_1486 int, dm_build_1487 string, dm_build_1488 string, dm_build_1489 *DmConnection) int {
	return dm_build_1485.Dm_build_1466(dm_build_1486, Dm_build_885.Dm_build_1095(dm_build_1487, dm_build_1488, dm_build_1489))
}

func (dm_build_1491 *Dm_build_1239) Dm_build_1490(dm_build_1492 int, dm_build_1493 string, dm_build_1494 string, dm_build_1495 *DmConnection) int {
	return dm_build_1491.Dm_build_1470(dm_build_1492, Dm_build_885.Dm_build_1095(dm_build_1493, dm_build_1494, dm_build_1495))
}

func (dm_build_1497 *Dm_build_1239) Dm_build_1496(dm_build_1498 int, dm_build_1499 string, dm_build_1500 string, dm_build_1501 *DmConnection) int {
	return dm_build_1497.Dm_build_1474(dm_build_1498, Dm_build_885.Dm_build_1095(dm_build_1499, dm_build_1500, dm_build_1501))
}

func (dm_build_1503 *Dm_build_1239) Dm_build_1502(dm_build_1504 int) byte {
	return Dm_build_885.Dm_build_1100(dm_build_1503.Dm_build_1529(dm_build_1504, 1))
}

func (dm_build_1506 *Dm_build_1239) Dm_build_1505(dm_build_1507 int) int16 {
	return Dm_build_885.Dm_build_1103(dm_build_1506.Dm_build_1529(dm_build_1507, 2))
}

func (dm_build_1509 *Dm_build_1239) Dm_build_1508(dm_build_1510 int) int32 {
	return Dm_build_885.Dm_build_1106(dm_build_1509.Dm_build_1529(dm_build_1510, 4))
}

func (dm_build_1512 *Dm_build_1239) Dm_build_1511(dm_build_1513 int) int64 {
	return Dm_build_885.Dm_build_1109(dm_build_1512.Dm_build_1529(dm_build_1513, 8))
}

func (dm_build_1515 *Dm_build_1239) Dm_build_1514(dm_build_1516 int) float32 {
	return Dm_build_885.Dm_build_1112(dm_build_1515.Dm_build_1529(dm_build_1516, 4))
}

func (dm_build_1518 *Dm_build_1239) Dm_build_1517(dm_build_1519 int) float64 {
	return Dm_build_885.Dm_build_1115(dm_build_1518.Dm_build_1529(dm_build_1519, 8))
}

func (dm_build_1521 *Dm_build_1239) Dm_build_1520(dm_build_1522 int) uint8 {
	return Dm_build_885.Dm_build_1118(dm_build_1521.Dm_build_1529(dm_build_1522, 1))
}

func (dm_build_1524 *Dm_build_1239) Dm_build_1523(dm_build_1525 int) uint16 {
	return Dm_build_885.Dm_build_1121(dm_build_1524.Dm_build_1529(dm_build_1525, 2))
}

func (dm_build_1527 *Dm_build_1239) Dm_build_1526(dm_build_1528 int) uint32 {
	return Dm_build_885.Dm_build_1124(dm_build_1527.Dm_build_1529(dm_build_1528, 4))
}

func (dm_build_1530 *Dm_build_1239) Dm_build_1529(dm_build_1531 int, dm_build_1532 int) []byte {
	return dm_build_1530.dm_build_1240[dm_build_1531 : dm_build_1531+dm_build_1532]
}

func (dm_build_1534 *Dm_build_1239) Dm_build_1533(dm_build_1535 int) []byte {
	dm_build_1536 := dm_build_1534.Dm_build_1508(dm_build_1535)
	return dm_build_1534.Dm_build_1529(dm_build_1535+4, int(dm_build_1536))
}

func (dm_build_1538 *Dm_build_1239) Dm_build_1537(dm_build_1539 int) []byte {
	dm_build_1540 := dm_build_1538.Dm_build_1502(dm_build_1539)
	return dm_build_1538.Dm_build_1529(dm_build_1539+1, int(dm_build_1540))
}

func (dm_build_1542 *Dm_build_1239) Dm_build_1541(dm_build_1543 int) []byte {
	dm_build_1544 := dm_build_1542.Dm_build_1505(dm_build_1543)
	return dm_build_1542.Dm_build_1529(dm_build_1543+2, int(dm_build_1544))
}

func (dm_build_1546 *Dm_build_1239) Dm_build_1545(dm_build_1547 int) []byte {
	dm_build_1548 := 0
	for dm_build_1546.Dm_build_1502(dm_build_1547) != 0 {
		dm_build_1547++
		dm_build_1548++
	}

	return dm_build_1546.Dm_build_1529(dm_build_1547-dm_build_1548, int(dm_build_1548))
}

func (dm_build_1550 *Dm_build_1239) Dm_build_1549(dm_build_1551 int, dm_build_1552 string, dm_build_1553 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1550.Dm_build_1533(dm_build_1551), dm_build_1552, dm_build_1553)
}

func (dm_build_1555 *Dm_build_1239) Dm_build_1554(dm_build_1556 int, dm_build_1557 string, dm_build_1558 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1555.Dm_build_1537(dm_build_1556), dm_build_1557, dm_build_1558)
}

func (dm_build_1560 *Dm_build_1239) Dm_build_1559(dm_build_1561 int, dm_build_1562 string, dm_build_1563 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1560.Dm_build_1541(dm_build_1561), dm_build_1562, dm_build_1563)
}

func (dm_build_1565 *Dm_build_1239) Dm_build_1564(dm_build_1566 int, dm_build_1567 string, dm_build_1568 *DmConnection) string {
	return Dm_build_885.Dm_build_1132(dm_build_1565.Dm_build_1545(dm_build_1566), dm_build_1567, dm_build_1568)
}
