/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1538 struct {
	dm_build_1539 *list.List
	dm_build_1540 *dm_build_1592
	dm_build_1541 int
}

func Dm_build_1542() *Dm_build_1538 {
	return &Dm_build_1538{
		dm_build_1539: list.New(),
		dm_build_1541: 0,
	}
}

func (dm_build_1544 *Dm_build_1538) Dm_build_1543() int {
	return dm_build_1544.dm_build_1541
}

func (dm_build_1546 *Dm_build_1538) Dm_build_1545(dm_build_1547 *Dm_build_0, dm_build_1548 int) int {
	var dm_build_1549 = 0
	var dm_build_1550 = 0
	for dm_build_1549 < dm_build_1548 && dm_build_1546.dm_build_1540 != nil {
		dm_build_1550 = dm_build_1546.dm_build_1540.dm_build_1600(dm_build_1547, dm_build_1548-dm_build_1549)
		if dm_build_1546.dm_build_1540.dm_build_1595 == 0 {
			dm_build_1546.dm_build_1582()
		}
		dm_build_1549 += dm_build_1550
		dm_build_1546.dm_build_1541 -= dm_build_1550
	}
	return dm_build_1549
}

func (dm_build_1552 *Dm_build_1538) Dm_build_1551(dm_build_1553 []byte, dm_build_1554 int, dm_build_1555 int) int {
	var dm_build_1556 = 0
	var dm_build_1557 = 0
	for dm_build_1556 < dm_build_1555 && dm_build_1552.dm_build_1540 != nil {
		dm_build_1557 = dm_build_1552.dm_build_1540.dm_build_1604(dm_build_1553, dm_build_1554, dm_build_1555-dm_build_1556)
		if dm_build_1552.dm_build_1540.dm_build_1595 == 0 {
			dm_build_1552.dm_build_1582()
		}
		dm_build_1556 += dm_build_1557
		dm_build_1552.dm_build_1541 -= dm_build_1557
		dm_build_1554 += dm_build_1557
	}
	return dm_build_1556
}

func (dm_build_1559 *Dm_build_1538) Dm_build_1558(dm_build_1560 io.Writer, dm_build_1561 int) int {
	var dm_build_1562 = 0
	var dm_build_1563 = 0
	for dm_build_1562 < dm_build_1561 && dm_build_1559.dm_build_1540 != nil {
		dm_build_1563 = dm_build_1559.dm_build_1540.dm_build_1609(dm_build_1560, dm_build_1561-dm_build_1562)
		if dm_build_1559.dm_build_1540.dm_build_1595 == 0 {
			dm_build_1559.dm_build_1582()
		}
		dm_build_1562 += dm_build_1563
		dm_build_1559.dm_build_1541 -= dm_build_1563
	}
	return dm_build_1562
}

func (dm_build_1565 *Dm_build_1538) Dm_build_1564(dm_build_1566 []byte, dm_build_1567 int, dm_build_1568 int) {
	if dm_build_1568 == 0 {
		return
	}
	var dm_build_1569 = dm_build_1596(dm_build_1566, dm_build_1567, dm_build_1568)
	if dm_build_1565.dm_build_1540 == nil {
		dm_build_1565.dm_build_1540 = dm_build_1569
	} else {
		dm_build_1565.dm_build_1539.PushBack(dm_build_1569)
	}
	dm_build_1565.dm_build_1541 += dm_build_1568
}

func (dm_build_1571 *Dm_build_1538) dm_build_1570(dm_build_1572 int) byte {
	var dm_build_1573 = dm_build_1572
	var dm_build_1574 = dm_build_1571.dm_build_1540
	for dm_build_1573 > 0 && dm_build_1574 != nil {
		if dm_build_1574.dm_build_1595 == 0 {
			continue
		}
		if dm_build_1573 > dm_build_1574.dm_build_1595-1 {
			dm_build_1573 -= dm_build_1574.dm_build_1595
			dm_build_1574 = dm_build_1571.dm_build_1539.Front().Value.(*dm_build_1592)
		} else {
			break
		}
	}
	return dm_build_1574.dm_build_1613(dm_build_1573)
}
func (dm_build_1576 *Dm_build_1538) Dm_build_1575(dm_build_1577 *Dm_build_1538) {
	if dm_build_1577.dm_build_1541 == 0 {
		return
	}
	var dm_build_1578 = dm_build_1577.dm_build_1540
	for dm_build_1578 != nil {
		dm_build_1576.dm_build_1579(dm_build_1578)
		dm_build_1577.dm_build_1582()
		dm_build_1578 = dm_build_1577.dm_build_1540
	}
	dm_build_1577.dm_build_1541 = 0
}
func (dm_build_1580 *Dm_build_1538) dm_build_1579(dm_build_1581 *dm_build_1592) {
	if dm_build_1581.dm_build_1595 == 0 {
		return
	}
	if dm_build_1580.dm_build_1540 == nil {
		dm_build_1580.dm_build_1540 = dm_build_1581
	} else {
		dm_build_1580.dm_build_1539.PushBack(dm_build_1581)
	}
	dm_build_1580.dm_build_1541 += dm_build_1581.dm_build_1595
}

func (dm_build_1583 *Dm_build_1538) dm_build_1582() {
	var dm_build_1584 = dm_build_1583.dm_build_1539.Front()
	if dm_build_1584 == nil {
		dm_build_1583.dm_build_1540 = nil
	} else {
		dm_build_1583.dm_build_1540 = dm_build_1584.Value.(*dm_build_1592)
		dm_build_1583.dm_build_1539.Remove(dm_build_1584)
	}
}

func (dm_build_1586 *Dm_build_1538) Dm_build_1585() []byte {
	var dm_build_1587 = make([]byte, dm_build_1586.dm_build_1541)
	var dm_build_1588 = dm_build_1586.dm_build_1540
	var dm_build_1589 = 0
	var dm_build_1590 = len(dm_build_1587)
	var dm_build_1591 = 0
	for dm_build_1588 != nil {
		if dm_build_1588.dm_build_1595 > 0 {
			if dm_build_1590 > dm_build_1588.dm_build_1595 {
				dm_build_1591 = dm_build_1588.dm_build_1595
			} else {
				dm_build_1591 = dm_build_1590
			}
			copy(dm_build_1587[dm_build_1589:dm_build_1589+dm_build_1591], dm_build_1588.dm_build_1593[dm_build_1588.dm_build_1594:dm_build_1588.dm_build_1594+dm_build_1591])
			dm_build_1589 += dm_build_1591
			dm_build_1590 -= dm_build_1591
		}
		if dm_build_1586.dm_build_1539.Front() == nil {
			dm_build_1588 = nil
		} else {
			dm_build_1588 = dm_build_1586.dm_build_1539.Front().Value.(*dm_build_1592)
		}
	}
	return dm_build_1587
}

type dm_build_1592 struct {
	dm_build_1593 []byte
	dm_build_1594 int
	dm_build_1595 int
}

func dm_build_1596(dm_build_1597 []byte, dm_build_1598 int, dm_build_1599 int) *dm_build_1592 {
	return &dm_build_1592{
		dm_build_1597,
		dm_build_1598,
		dm_build_1599,
	}
}

func (dm_build_1601 *dm_build_1592) dm_build_1600(dm_build_1602 *Dm_build_0, dm_build_1603 int) int {
	if dm_build_1601.dm_build_1595 <= dm_build_1603 {
		dm_build_1603 = dm_build_1601.dm_build_1595
	}
	dm_build_1602.Dm_build_83(dm_build_1601.dm_build_1593[dm_build_1601.dm_build_1594 : dm_build_1601.dm_build_1594+dm_build_1603])
	dm_build_1601.dm_build_1594 += dm_build_1603
	dm_build_1601.dm_build_1595 -= dm_build_1603
	return dm_build_1603
}

func (dm_build_1605 *dm_build_1592) dm_build_1604(dm_build_1606 []byte, dm_build_1607 int, dm_build_1608 int) int {
	if dm_build_1605.dm_build_1595 <= dm_build_1608 {
		dm_build_1608 = dm_build_1605.dm_build_1595
	}
	copy(dm_build_1606[dm_build_1607:dm_build_1607+dm_build_1608], dm_build_1605.dm_build_1593[dm_build_1605.dm_build_1594:dm_build_1605.dm_build_1594+dm_build_1608])
	dm_build_1605.dm_build_1594 += dm_build_1608
	dm_build_1605.dm_build_1595 -= dm_build_1608
	return dm_build_1608
}

func (dm_build_1610 *dm_build_1592) dm_build_1609(dm_build_1611 io.Writer, dm_build_1612 int) int {
	if dm_build_1610.dm_build_1595 <= dm_build_1612 {
		dm_build_1612 = dm_build_1610.dm_build_1595
	}
	dm_build_1611.Write(dm_build_1610.dm_build_1593[dm_build_1610.dm_build_1594 : dm_build_1610.dm_build_1594+dm_build_1612])
	dm_build_1610.dm_build_1594 += dm_build_1612
	dm_build_1610.dm_build_1595 -= dm_build_1612
	return dm_build_1612
}
func (dm_build_1614 *dm_build_1592) dm_build_1613(dm_build_1615 int) byte {
	return dm_build_1614.dm_build_1593[dm_build_1614.dm_build_1594+dm_build_1615]
}
