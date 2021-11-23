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

type dm_build_1218 struct{}

var Dm_build_1219 = &dm_build_1218{}

func (Dm_build_1221 *dm_build_1218) Dm_build_1220(dm_build_1222 []byte, dm_build_1223 int, dm_build_1224 byte) int {
	dm_build_1222[dm_build_1223] = dm_build_1224
	return 1
}

func (Dm_build_1226 *dm_build_1218) Dm_build_1225(dm_build_1227 []byte, dm_build_1228 int, dm_build_1229 int8) int {
	dm_build_1227[dm_build_1228] = byte(dm_build_1229)
	return 1
}

func (Dm_build_1231 *dm_build_1218) Dm_build_1230(dm_build_1232 []byte, dm_build_1233 int, dm_build_1234 int16) int {
	dm_build_1232[dm_build_1233] = byte(dm_build_1234)
	dm_build_1233++
	dm_build_1232[dm_build_1233] = byte(dm_build_1234 >> 8)
	return 2
}

func (Dm_build_1236 *dm_build_1218) Dm_build_1235(dm_build_1237 []byte, dm_build_1238 int, dm_build_1239 int32) int {
	dm_build_1237[dm_build_1238] = byte(dm_build_1239)
	dm_build_1238++
	dm_build_1237[dm_build_1238] = byte(dm_build_1239 >> 8)
	dm_build_1238++
	dm_build_1237[dm_build_1238] = byte(dm_build_1239 >> 16)
	dm_build_1238++
	dm_build_1237[dm_build_1238] = byte(dm_build_1239 >> 24)
	dm_build_1238++
	return 4
}

func (Dm_build_1241 *dm_build_1218) Dm_build_1240(dm_build_1242 []byte, dm_build_1243 int, dm_build_1244 int64) int {
	dm_build_1242[dm_build_1243] = byte(dm_build_1244)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 8)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 16)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 24)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 32)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 40)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 48)
	dm_build_1243++
	dm_build_1242[dm_build_1243] = byte(dm_build_1244 >> 56)
	return 8
}

func (Dm_build_1246 *dm_build_1218) Dm_build_1245(dm_build_1247 []byte, dm_build_1248 int, dm_build_1249 float32) int {
	return Dm_build_1246.Dm_build_1265(dm_build_1247, dm_build_1248, math.Float32bits(dm_build_1249))
}

func (Dm_build_1251 *dm_build_1218) Dm_build_1250(dm_build_1252 []byte, dm_build_1253 int, dm_build_1254 float64) int {
	return Dm_build_1251.Dm_build_1270(dm_build_1252, dm_build_1253, math.Float64bits(dm_build_1254))
}

func (Dm_build_1256 *dm_build_1218) Dm_build_1255(dm_build_1257 []byte, dm_build_1258 int, dm_build_1259 uint8) int {
	dm_build_1257[dm_build_1258] = byte(dm_build_1259)
	return 1
}

func (Dm_build_1261 *dm_build_1218) Dm_build_1260(dm_build_1262 []byte, dm_build_1263 int, dm_build_1264 uint16) int {
	dm_build_1262[dm_build_1263] = byte(dm_build_1264)
	dm_build_1263++
	dm_build_1262[dm_build_1263] = byte(dm_build_1264 >> 8)
	return 2
}

func (Dm_build_1266 *dm_build_1218) Dm_build_1265(dm_build_1267 []byte, dm_build_1268 int, dm_build_1269 uint32) int {
	dm_build_1267[dm_build_1268] = byte(dm_build_1269)
	dm_build_1268++
	dm_build_1267[dm_build_1268] = byte(dm_build_1269 >> 8)
	dm_build_1268++
	dm_build_1267[dm_build_1268] = byte(dm_build_1269 >> 16)
	dm_build_1268++
	dm_build_1267[dm_build_1268] = byte(dm_build_1269 >> 24)
	return 3
}

func (Dm_build_1271 *dm_build_1218) Dm_build_1270(dm_build_1272 []byte, dm_build_1273 int, dm_build_1274 uint64) int {
	dm_build_1272[dm_build_1273] = byte(dm_build_1274)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 8)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 16)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 24)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 32)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 40)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 48)
	dm_build_1273++
	dm_build_1272[dm_build_1273] = byte(dm_build_1274 >> 56)
	return 3
}

func (Dm_build_1276 *dm_build_1218) Dm_build_1275(dm_build_1277 []byte, dm_build_1278 int, dm_build_1279 []byte, dm_build_1280 int, dm_build_1281 int) int {
	copy(dm_build_1277[dm_build_1278:dm_build_1278+dm_build_1281], dm_build_1279[dm_build_1280:dm_build_1280+dm_build_1281])
	return dm_build_1281
}

func (Dm_build_1283 *dm_build_1218) Dm_build_1282(dm_build_1284 []byte, dm_build_1285 int, dm_build_1286 []byte, dm_build_1287 int, dm_build_1288 int) int {
	dm_build_1285 += Dm_build_1283.Dm_build_1265(dm_build_1284, dm_build_1285, uint32(dm_build_1288))
	return 4 + Dm_build_1283.Dm_build_1275(dm_build_1284, dm_build_1285, dm_build_1286, dm_build_1287, dm_build_1288)
}

func (Dm_build_1290 *dm_build_1218) Dm_build_1289(dm_build_1291 []byte, dm_build_1292 int, dm_build_1293 []byte, dm_build_1294 int, dm_build_1295 int) int {
	dm_build_1292 += Dm_build_1290.Dm_build_1260(dm_build_1291, dm_build_1292, uint16(dm_build_1295))
	return 2 + Dm_build_1290.Dm_build_1275(dm_build_1291, dm_build_1292, dm_build_1293, dm_build_1294, dm_build_1295)
}

func (Dm_build_1297 *dm_build_1218) Dm_build_1296(dm_build_1298 []byte, dm_build_1299 int, dm_build_1300 string, dm_build_1301 string, dm_build_1302 *Connection) int {
	dm_build_1303 := Dm_build_1297.Dm_build_1432(dm_build_1300, dm_build_1301, dm_build_1302)
	dm_build_1299 += Dm_build_1297.Dm_build_1265(dm_build_1298, dm_build_1299, uint32(len(dm_build_1303)))
	return 4 + Dm_build_1297.Dm_build_1275(dm_build_1298, dm_build_1299, dm_build_1303, 0, len(dm_build_1303))
}

func (Dm_build_1305 *dm_build_1218) Dm_build_1304(dm_build_1306 []byte, dm_build_1307 int, dm_build_1308 string, dm_build_1309 string, dm_build_1310 *Connection) int {
	dm_build_1311 := Dm_build_1305.Dm_build_1432(dm_build_1308, dm_build_1309, dm_build_1310)

	dm_build_1307 += Dm_build_1305.Dm_build_1260(dm_build_1306, dm_build_1307, uint16(len(dm_build_1311)))
	return 2 + Dm_build_1305.Dm_build_1275(dm_build_1306, dm_build_1307, dm_build_1311, 0, len(dm_build_1311))
}

func (Dm_build_1313 *dm_build_1218) Dm_build_1312(dm_build_1314 []byte, dm_build_1315 int) byte {
	return dm_build_1314[dm_build_1315]
}

func (Dm_build_1317 *dm_build_1218) Dm_build_1316(dm_build_1318 []byte, dm_build_1319 int) int16 {
	var dm_build_1320 int16
	dm_build_1320 = int16(dm_build_1318[dm_build_1319] & 0xff)
	dm_build_1319++
	dm_build_1320 |= int16(dm_build_1318[dm_build_1319]&0xff) << 8
	return dm_build_1320
}

func (Dm_build_1322 *dm_build_1218) Dm_build_1321(dm_build_1323 []byte, dm_build_1324 int) int32 {
	var dm_build_1325 int32
	dm_build_1325 = int32(dm_build_1323[dm_build_1324] & 0xff)
	dm_build_1324++
	dm_build_1325 |= int32(dm_build_1323[dm_build_1324]&0xff) << 8
	dm_build_1324++
	dm_build_1325 |= int32(dm_build_1323[dm_build_1324]&0xff) << 16
	dm_build_1324++
	dm_build_1325 |= int32(dm_build_1323[dm_build_1324]&0xff) << 24
	return dm_build_1325
}

func (Dm_build_1327 *dm_build_1218) Dm_build_1326(dm_build_1328 []byte, dm_build_1329 int) int64 {
	var dm_build_1330 int64
	dm_build_1330 = int64(dm_build_1328[dm_build_1329] & 0xff)
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 8
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 16
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 24
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 32
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 40
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 48
	dm_build_1329++
	dm_build_1330 |= int64(dm_build_1328[dm_build_1329]&0xff) << 56
	return dm_build_1330
}

func (Dm_build_1332 *dm_build_1218) Dm_build_1331(dm_build_1333 []byte, dm_build_1334 int) float32 {
	return math.Float32frombits(Dm_build_1332.Dm_build_1348(dm_build_1333, dm_build_1334))
}

func (Dm_build_1336 *dm_build_1218) Dm_build_1335(dm_build_1337 []byte, dm_build_1338 int) float64 {
	return math.Float64frombits(Dm_build_1336.Dm_build_1353(dm_build_1337, dm_build_1338))
}

func (Dm_build_1340 *dm_build_1218) Dm_build_1339(dm_build_1341 []byte, dm_build_1342 int) uint8 {
	return uint8(dm_build_1341[dm_build_1342] & 0xff)
}

func (Dm_build_1344 *dm_build_1218) Dm_build_1343(dm_build_1345 []byte, dm_build_1346 int) uint16 {
	var dm_build_1347 uint16
	dm_build_1347 = uint16(dm_build_1345[dm_build_1346] & 0xff)
	dm_build_1346++
	dm_build_1347 |= uint16(dm_build_1345[dm_build_1346]&0xff) << 8
	return dm_build_1347
}

func (Dm_build_1349 *dm_build_1218) Dm_build_1348(dm_build_1350 []byte, dm_build_1351 int) uint32 {
	var dm_build_1352 uint32
	dm_build_1352 = uint32(dm_build_1350[dm_build_1351] & 0xff)
	dm_build_1351++
	dm_build_1352 |= uint32(dm_build_1350[dm_build_1351]&0xff) << 8
	dm_build_1351++
	dm_build_1352 |= uint32(dm_build_1350[dm_build_1351]&0xff) << 16
	dm_build_1351++
	dm_build_1352 |= uint32(dm_build_1350[dm_build_1351]&0xff) << 24
	return dm_build_1352
}

func (Dm_build_1354 *dm_build_1218) Dm_build_1353(dm_build_1355 []byte, dm_build_1356 int) uint64 {
	var dm_build_1357 uint64
	dm_build_1357 = uint64(dm_build_1355[dm_build_1356] & 0xff)
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 8
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 16
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 24
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 32
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 40
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 48
	dm_build_1356++
	dm_build_1357 |= uint64(dm_build_1355[dm_build_1356]&0xff) << 56
	return dm_build_1357
}

func (Dm_build_1359 *dm_build_1218) Dm_build_1358(dm_build_1360 []byte, dm_build_1361 int) []byte {
	dm_build_1362 := Dm_build_1359.Dm_build_1348(dm_build_1360, dm_build_1361)

	dm_build_1363 := make([]byte, dm_build_1362)
	copy(dm_build_1363[:int(dm_build_1362)], dm_build_1360[dm_build_1361+4:dm_build_1361+4+int(dm_build_1362)])
	return dm_build_1363
}

func (Dm_build_1365 *dm_build_1218) Dm_build_1364(dm_build_1366 []byte, dm_build_1367 int) []byte {
	dm_build_1368 := Dm_build_1365.Dm_build_1343(dm_build_1366, dm_build_1367)

	dm_build_1369 := make([]byte, dm_build_1368)
	copy(dm_build_1369[:int(dm_build_1368)], dm_build_1366[dm_build_1367+2:dm_build_1367+2+int(dm_build_1368)])
	return dm_build_1369
}

func (Dm_build_1371 *dm_build_1218) Dm_build_1370(dm_build_1372 []byte, dm_build_1373 int, dm_build_1374 int) []byte {

	dm_build_1375 := make([]byte, dm_build_1374)
	copy(dm_build_1375[:dm_build_1374], dm_build_1372[dm_build_1373:dm_build_1373+dm_build_1374])
	return dm_build_1375
}

func (Dm_build_1377 *dm_build_1218) Dm_build_1376(dm_build_1378 []byte, dm_build_1379 int, dm_build_1380 int, dm_build_1381 string, dm_build_1382 *Connection) string {
	return Dm_build_1377.Dm_build_1469(dm_build_1378[dm_build_1379:dm_build_1379+dm_build_1380], dm_build_1381, dm_build_1382)
}

func (Dm_build_1384 *dm_build_1218) Dm_build_1383(dm_build_1385 []byte, dm_build_1386 int, dm_build_1387 string, dm_build_1388 *Connection) string {
	dm_build_1389 := Dm_build_1384.Dm_build_1348(dm_build_1385, dm_build_1386)
	dm_build_1386 += 4
	return Dm_build_1384.Dm_build_1376(dm_build_1385, dm_build_1386, int(dm_build_1389), dm_build_1387, dm_build_1388)
}

func (Dm_build_1391 *dm_build_1218) Dm_build_1390(dm_build_1392 []byte, dm_build_1393 int, dm_build_1394 string, dm_build_1395 *Connection) string {
	dm_build_1396 := Dm_build_1391.Dm_build_1343(dm_build_1392, dm_build_1393)
	dm_build_1393 += 2
	return Dm_build_1391.Dm_build_1376(dm_build_1392, dm_build_1393, int(dm_build_1396), dm_build_1394, dm_build_1395)
}

func (Dm_build_1398 *dm_build_1218) Dm_build_1397(dm_build_1399 byte) []byte {
	return []byte{dm_build_1399}
}

func (Dm_build_1401 *dm_build_1218) Dm_build_1400(dm_build_1402 int16) []byte {
	return []byte{byte(dm_build_1402), byte(dm_build_1402 >> 8)}
}

func (Dm_build_1404 *dm_build_1218) Dm_build_1403(dm_build_1405 int32) []byte {
	return []byte{byte(dm_build_1405), byte(dm_build_1405 >> 8), byte(dm_build_1405 >> 16), byte(dm_build_1405 >> 24)}
}

func (Dm_build_1407 *dm_build_1218) Dm_build_1406(dm_build_1408 int64) []byte {
	return []byte{byte(dm_build_1408), byte(dm_build_1408 >> 8), byte(dm_build_1408 >> 16), byte(dm_build_1408 >> 24), byte(dm_build_1408 >> 32),
		byte(dm_build_1408 >> 40), byte(dm_build_1408 >> 48), byte(dm_build_1408 >> 56)}
}

func (Dm_build_1410 *dm_build_1218) Dm_build_1409(dm_build_1411 float32) []byte {
	return Dm_build_1410.Dm_build_1421(math.Float32bits(dm_build_1411))
}

func (Dm_build_1413 *dm_build_1218) Dm_build_1412(dm_build_1414 float64) []byte {
	return Dm_build_1413.Dm_build_1424(math.Float64bits(dm_build_1414))
}

func (Dm_build_1416 *dm_build_1218) Dm_build_1415(dm_build_1417 uint8) []byte {
	return []byte{byte(dm_build_1417)}
}

func (Dm_build_1419 *dm_build_1218) Dm_build_1418(dm_build_1420 uint16) []byte {
	return []byte{byte(dm_build_1420), byte(dm_build_1420 >> 8)}
}

func (Dm_build_1422 *dm_build_1218) Dm_build_1421(dm_build_1423 uint32) []byte {
	return []byte{byte(dm_build_1423), byte(dm_build_1423 >> 8), byte(dm_build_1423 >> 16), byte(dm_build_1423 >> 24)}
}

func (Dm_build_1425 *dm_build_1218) Dm_build_1424(dm_build_1426 uint64) []byte {
	return []byte{byte(dm_build_1426), byte(dm_build_1426 >> 8), byte(dm_build_1426 >> 16), byte(dm_build_1426 >> 24), byte(dm_build_1426 >> 32), byte(dm_build_1426 >> 40), byte(dm_build_1426 >> 48), byte(dm_build_1426 >> 56)}
}

func (Dm_build_1428 *dm_build_1218) Dm_build_1427(dm_build_1429 []byte, dm_build_1430 string, dm_build_1431 *Connection) []byte {
	if dm_build_1430 == "UTF-8" {
		return dm_build_1429
	}

	if dm_build_1431 == nil {
		if e := dm_build_1474(dm_build_1430); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1429), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1431.encodeBuffer == nil {
		dm_build_1431.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1431.encode = dm_build_1474(dm_build_1431.getServerEncoding())
		dm_build_1431.transformReaderDst = make([]byte, 4096)
		dm_build_1431.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1431.encode; e != nil {

		dm_build_1431.encodeBuffer.Reset()

		n, err := dm_build_1431.encodeBuffer.ReadFrom(
			Dm_build_1488(bytes.NewReader(dm_build_1429), e.NewEncoder(), dm_build_1431.transformReaderDst, dm_build_1431.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1431.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1433 *dm_build_1218) Dm_build_1432(dm_build_1434 string, dm_build_1435 string, dm_build_1436 *Connection) []byte {
	return Dm_build_1433.Dm_build_1427([]byte(dm_build_1434), dm_build_1435, dm_build_1436)
}

func (Dm_build_1438 *dm_build_1218) Dm_build_1437(dm_build_1439 []byte) byte {
	return Dm_build_1438.Dm_build_1312(dm_build_1439, 0)
}

func (Dm_build_1441 *dm_build_1218) Dm_build_1440(dm_build_1442 []byte) int16 {
	return Dm_build_1441.Dm_build_1316(dm_build_1442, 0)
}

func (Dm_build_1444 *dm_build_1218) Dm_build_1443(dm_build_1445 []byte) int32 {
	return Dm_build_1444.Dm_build_1321(dm_build_1445, 0)
}

func (Dm_build_1447 *dm_build_1218) Dm_build_1446(dm_build_1448 []byte) int64 {
	return Dm_build_1447.Dm_build_1326(dm_build_1448, 0)
}

func (Dm_build_1450 *dm_build_1218) Dm_build_1449(dm_build_1451 []byte) float32 {
	return Dm_build_1450.Dm_build_1331(dm_build_1451, 0)
}

func (Dm_build_1453 *dm_build_1218) Dm_build_1452(dm_build_1454 []byte) float64 {
	return Dm_build_1453.Dm_build_1335(dm_build_1454, 0)
}

func (Dm_build_1456 *dm_build_1218) Dm_build_1455(dm_build_1457 []byte) uint8 {
	return Dm_build_1456.Dm_build_1339(dm_build_1457, 0)
}

func (Dm_build_1459 *dm_build_1218) Dm_build_1458(dm_build_1460 []byte) uint16 {
	return Dm_build_1459.Dm_build_1343(dm_build_1460, 0)
}

func (Dm_build_1462 *dm_build_1218) Dm_build_1461(dm_build_1463 []byte) uint32 {
	return Dm_build_1462.Dm_build_1348(dm_build_1463, 0)
}

func (Dm_build_1465 *dm_build_1218) Dm_build_1464(dm_build_1466 []byte, dm_build_1467 string, dm_build_1468 *Connection) []byte {
	if dm_build_1467 == "UTF-8" {
		return dm_build_1466
	}

	if dm_build_1468 == nil {
		if e := dm_build_1474(dm_build_1467); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1466), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1468.encodeBuffer == nil {
		dm_build_1468.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1468.encode = dm_build_1474(dm_build_1468.getServerEncoding())
		dm_build_1468.transformReaderDst = make([]byte, 4096)
		dm_build_1468.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1468.encode; e != nil {

		dm_build_1468.encodeBuffer.Reset()

		n, err := dm_build_1468.encodeBuffer.ReadFrom(
			Dm_build_1488(bytes.NewReader(dm_build_1466), e.NewDecoder(), dm_build_1468.transformReaderDst, dm_build_1468.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1468.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1470 *dm_build_1218) Dm_build_1469(dm_build_1471 []byte, dm_build_1472 string, dm_build_1473 *Connection) string {
	return string(Dm_build_1470.Dm_build_1464(dm_build_1471, dm_build_1472, dm_build_1473))
}

func dm_build_1474(dm_build_1475 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1475); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1476 struct {
	dm_build_1477 io.Reader
	dm_build_1478 transform.Transformer
	dm_build_1479 error

	dm_build_1480                []byte
	dm_build_1481, dm_build_1482 int

	dm_build_1483                []byte
	dm_build_1484, dm_build_1485 int

	dm_build_1486 bool
}

const dm_build_1487 = 4096

func Dm_build_1488(dm_build_1489 io.Reader, dm_build_1490 transform.Transformer, dm_build_1491 []byte, dm_build_1492 []byte) *Dm_build_1476 {
	dm_build_1490.Reset()
	return &Dm_build_1476{
		dm_build_1477: dm_build_1489,
		dm_build_1478: dm_build_1490,
		dm_build_1480: dm_build_1491,
		dm_build_1483: dm_build_1492,
	}
}

func (dm_build_1494 *Dm_build_1476) Read(dm_build_1495 []byte) (int, error) {
	dm_build_1496, dm_build_1497 := 0, error(nil)
	for {

		if dm_build_1494.dm_build_1481 != dm_build_1494.dm_build_1482 {
			dm_build_1496 = copy(dm_build_1495, dm_build_1494.dm_build_1480[dm_build_1494.dm_build_1481:dm_build_1494.dm_build_1482])
			dm_build_1494.dm_build_1481 += dm_build_1496
			if dm_build_1494.dm_build_1481 == dm_build_1494.dm_build_1482 && dm_build_1494.dm_build_1486 {
				return dm_build_1496, dm_build_1494.dm_build_1479
			}
			return dm_build_1496, nil
		} else if dm_build_1494.dm_build_1486 {
			return 0, dm_build_1494.dm_build_1479
		}

		if dm_build_1494.dm_build_1484 != dm_build_1494.dm_build_1485 || dm_build_1494.dm_build_1479 != nil {
			dm_build_1494.dm_build_1481 = 0
			dm_build_1494.dm_build_1482, dm_build_1496, dm_build_1497 = dm_build_1494.dm_build_1478.Transform(dm_build_1494.dm_build_1480, dm_build_1494.dm_build_1483[dm_build_1494.dm_build_1484:dm_build_1494.dm_build_1485], dm_build_1494.dm_build_1479 == io.EOF)
			dm_build_1494.dm_build_1484 += dm_build_1496

			switch {
			case dm_build_1497 == nil:
				if dm_build_1494.dm_build_1484 != dm_build_1494.dm_build_1485 {
					dm_build_1494.dm_build_1479 = nil
				}

				dm_build_1494.dm_build_1486 = dm_build_1494.dm_build_1479 != nil
				continue
			case dm_build_1497 == transform.ErrShortDst && (dm_build_1494.dm_build_1482 != 0 || dm_build_1496 != 0):

				continue
			case dm_build_1497 == transform.ErrShortSrc && dm_build_1494.dm_build_1485-dm_build_1494.dm_build_1484 != len(dm_build_1494.dm_build_1483) && dm_build_1494.dm_build_1479 == nil:

			default:
				dm_build_1494.dm_build_1486 = true

				if dm_build_1494.dm_build_1479 == nil || dm_build_1494.dm_build_1479 == io.EOF {
					dm_build_1494.dm_build_1479 = dm_build_1497
				}
				continue
			}
		}

		if dm_build_1494.dm_build_1484 != 0 {
			dm_build_1494.dm_build_1484, dm_build_1494.dm_build_1485 = 0, copy(dm_build_1494.dm_build_1483, dm_build_1494.dm_build_1483[dm_build_1494.dm_build_1484:dm_build_1494.dm_build_1485])
		}
		dm_build_1496, dm_build_1494.dm_build_1479 = dm_build_1494.dm_build_1477.Read(dm_build_1494.dm_build_1483[dm_build_1494.dm_build_1485:])
		dm_build_1494.dm_build_1485 += dm_build_1496
	}
}
