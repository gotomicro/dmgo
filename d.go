/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1498 struct {
	dm_build_1499 *list.List
	dm_build_1500 *dm_build_1552
	dm_build_1501 int
}

func Dm_build_1502() *Dm_build_1498 {
	return &Dm_build_1498{
		dm_build_1499: list.New(),
		dm_build_1501: 0,
	}
}

func (dm_build_1504 *Dm_build_1498) Dm_build_1503() int {
	return dm_build_1504.dm_build_1501
}

func (dm_build_1506 *Dm_build_1498) Dm_build_1505(dm_build_1507 *Dm_build_0, dm_build_1508 int) int {
	var dm_build_1509 = 0
	var dm_build_1510 = 0
	for dm_build_1509 < dm_build_1508 && dm_build_1506.dm_build_1500 != nil {
		dm_build_1510 = dm_build_1506.dm_build_1500.dm_build_1560(dm_build_1507, dm_build_1508-dm_build_1509)
		if dm_build_1506.dm_build_1500.dm_build_1555 == 0 {
			dm_build_1506.dm_build_1542()
		}
		dm_build_1509 += dm_build_1510
		dm_build_1506.dm_build_1501 -= dm_build_1510
	}
	return dm_build_1509
}

func (dm_build_1512 *Dm_build_1498) Dm_build_1511(dm_build_1513 []byte, dm_build_1514 int, dm_build_1515 int) int {
	var dm_build_1516 = 0
	var dm_build_1517 = 0
	for dm_build_1516 < dm_build_1515 && dm_build_1512.dm_build_1500 != nil {
		dm_build_1517 = dm_build_1512.dm_build_1500.dm_build_1564(dm_build_1513, dm_build_1514, dm_build_1515-dm_build_1516)
		if dm_build_1512.dm_build_1500.dm_build_1555 == 0 {
			dm_build_1512.dm_build_1542()
		}
		dm_build_1516 += dm_build_1517
		dm_build_1512.dm_build_1501 -= dm_build_1517
		dm_build_1514 += dm_build_1517
	}
	return dm_build_1516
}

func (dm_build_1519 *Dm_build_1498) Dm_build_1518(dm_build_1520 io.Writer, dm_build_1521 int) int {
	var dm_build_1522 = 0
	var dm_build_1523 = 0
	for dm_build_1522 < dm_build_1521 && dm_build_1519.dm_build_1500 != nil {
		dm_build_1523 = dm_build_1519.dm_build_1500.dm_build_1569(dm_build_1520, dm_build_1521-dm_build_1522)
		if dm_build_1519.dm_build_1500.dm_build_1555 == 0 {
			dm_build_1519.dm_build_1542()
		}
		dm_build_1522 += dm_build_1523
		dm_build_1519.dm_build_1501 -= dm_build_1523
	}
	return dm_build_1522
}

func (dm_build_1525 *Dm_build_1498) Dm_build_1524(dm_build_1526 []byte, dm_build_1527 int, dm_build_1528 int) {
	if dm_build_1528 == 0 {
		return
	}
	var dm_build_1529 = dm_build_1556(dm_build_1526, dm_build_1527, dm_build_1528)
	if dm_build_1525.dm_build_1500 == nil {
		dm_build_1525.dm_build_1500 = dm_build_1529
	} else {
		dm_build_1525.dm_build_1499.PushBack(dm_build_1529)
	}
	dm_build_1525.dm_build_1501 += dm_build_1528
}

func (dm_build_1531 *Dm_build_1498) dm_build_1530(dm_build_1532 int) byte {
	var dm_build_1533 = dm_build_1532
	var dm_build_1534 = dm_build_1531.dm_build_1500
	for dm_build_1533 > 0 && dm_build_1534 != nil {
		if dm_build_1534.dm_build_1555 == 0 {
			continue
		}
		if dm_build_1533 > dm_build_1534.dm_build_1555-1 {
			dm_build_1533 -= dm_build_1534.dm_build_1555
			dm_build_1534 = dm_build_1531.dm_build_1499.Front().Value.(*dm_build_1552)
		} else {
			break
		}
	}
	return dm_build_1534.dm_build_1573(dm_build_1533)
}
func (dm_build_1536 *Dm_build_1498) Dm_build_1535(dm_build_1537 *Dm_build_1498) {
	if dm_build_1537.dm_build_1501 == 0 {
		return
	}
	var dm_build_1538 = dm_build_1537.dm_build_1500
	for dm_build_1538 != nil {
		dm_build_1536.dm_build_1539(dm_build_1538)
		dm_build_1537.dm_build_1542()
		dm_build_1538 = dm_build_1537.dm_build_1500
	}
	dm_build_1537.dm_build_1501 = 0
}
func (dm_build_1540 *Dm_build_1498) dm_build_1539(dm_build_1541 *dm_build_1552) {
	if dm_build_1541.dm_build_1555 == 0 {
		return
	}
	if dm_build_1540.dm_build_1500 == nil {
		dm_build_1540.dm_build_1500 = dm_build_1541
	} else {
		dm_build_1540.dm_build_1499.PushBack(dm_build_1541)
	}
	dm_build_1540.dm_build_1501 += dm_build_1541.dm_build_1555
}

func (dm_build_1543 *Dm_build_1498) dm_build_1542() {
	var dm_build_1544 = dm_build_1543.dm_build_1499.Front()
	if dm_build_1544 == nil {
		dm_build_1543.dm_build_1500 = nil
	} else {
		dm_build_1543.dm_build_1500 = dm_build_1544.Value.(*dm_build_1552)
		dm_build_1543.dm_build_1499.Remove(dm_build_1544)
	}
}

func (dm_build_1546 *Dm_build_1498) Dm_build_1545() []byte {
	var dm_build_1547 = make([]byte, dm_build_1546.dm_build_1501)
	var dm_build_1548 = dm_build_1546.dm_build_1500
	var dm_build_1549 = 0
	var dm_build_1550 = len(dm_build_1547)
	var dm_build_1551 = 0
	for dm_build_1548 != nil {
		if dm_build_1548.dm_build_1555 > 0 {
			if dm_build_1550 > dm_build_1548.dm_build_1555 {
				dm_build_1551 = dm_build_1548.dm_build_1555
			} else {
				dm_build_1551 = dm_build_1550
			}
			copy(dm_build_1547[dm_build_1549:dm_build_1549+dm_build_1551], dm_build_1548.dm_build_1553[dm_build_1548.dm_build_1554:dm_build_1548.dm_build_1554+dm_build_1551])
			dm_build_1549 += dm_build_1551
			dm_build_1550 -= dm_build_1551
		}
		if dm_build_1546.dm_build_1499.Front() == nil {
			dm_build_1548 = nil
		} else {
			dm_build_1548 = dm_build_1546.dm_build_1499.Front().Value.(*dm_build_1552)
		}
	}
	return dm_build_1547
}

type dm_build_1552 struct {
	dm_build_1553 []byte
	dm_build_1554 int
	dm_build_1555 int
}

func dm_build_1556(dm_build_1557 []byte, dm_build_1558 int, dm_build_1559 int) *dm_build_1552 {
	return &dm_build_1552{
		dm_build_1557,
		dm_build_1558,
		dm_build_1559,
	}
}

func (dm_build_1561 *dm_build_1552) dm_build_1560(dm_build_1562 *Dm_build_0, dm_build_1563 int) int {
	if dm_build_1561.dm_build_1555 <= dm_build_1563 {
		dm_build_1563 = dm_build_1561.dm_build_1555
	}
	dm_build_1562.Dm_build_79(dm_build_1561.dm_build_1553[dm_build_1561.dm_build_1554 : dm_build_1561.dm_build_1554+dm_build_1563])
	dm_build_1561.dm_build_1554 += dm_build_1563
	dm_build_1561.dm_build_1555 -= dm_build_1563
	return dm_build_1563
}

func (dm_build_1565 *dm_build_1552) dm_build_1564(dm_build_1566 []byte, dm_build_1567 int, dm_build_1568 int) int {
	if dm_build_1565.dm_build_1555 <= dm_build_1568 {
		dm_build_1568 = dm_build_1565.dm_build_1555
	}
	copy(dm_build_1566[dm_build_1567:dm_build_1567+dm_build_1568], dm_build_1565.dm_build_1553[dm_build_1565.dm_build_1554:dm_build_1565.dm_build_1554+dm_build_1568])
	dm_build_1565.dm_build_1554 += dm_build_1568
	dm_build_1565.dm_build_1555 -= dm_build_1568
	return dm_build_1568
}

func (dm_build_1570 *dm_build_1552) dm_build_1569(dm_build_1571 io.Writer, dm_build_1572 int) int {
	if dm_build_1570.dm_build_1555 <= dm_build_1572 {
		dm_build_1572 = dm_build_1570.dm_build_1555
	}
	dm_build_1571.Write(dm_build_1570.dm_build_1553[dm_build_1570.dm_build_1554 : dm_build_1570.dm_build_1554+dm_build_1572])
	dm_build_1570.dm_build_1554 += dm_build_1572
	dm_build_1570.dm_build_1555 -= dm_build_1572
	return dm_build_1572
}
func (dm_build_1574 *dm_build_1552) dm_build_1573(dm_build_1575 int) byte {
	return dm_build_1574.dm_build_1553[dm_build_1574.dm_build_1554+dm_build_1575]
}
