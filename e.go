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

type dm_build_598 struct{}

var Dm_build_599 = &dm_build_598{}

func (Dm_build_601 *dm_build_598) Dm_build_600(dm_build_602 []byte, dm_build_603 int, dm_build_604 byte) int {
	dm_build_602[dm_build_603] = dm_build_604
	return 1
}

func (Dm_build_606 *dm_build_598) Dm_build_605(dm_build_607 []byte, dm_build_608 int, dm_build_609 int8) int {
	dm_build_607[dm_build_608] = byte(dm_build_609)
	return 1
}

func (Dm_build_611 *dm_build_598) Dm_build_610(dm_build_612 []byte, dm_build_613 int, dm_build_614 int16) int {
	dm_build_612[dm_build_613] = byte(dm_build_614)
	dm_build_613++
	dm_build_612[dm_build_613] = byte(dm_build_614 >> 8)
	return 2
}

func (Dm_build_616 *dm_build_598) Dm_build_615(dm_build_617 []byte, dm_build_618 int, dm_build_619 int32) int {
	dm_build_617[dm_build_618] = byte(dm_build_619)
	dm_build_618++
	dm_build_617[dm_build_618] = byte(dm_build_619 >> 8)
	dm_build_618++
	dm_build_617[dm_build_618] = byte(dm_build_619 >> 16)
	dm_build_618++
	dm_build_617[dm_build_618] = byte(dm_build_619 >> 24)
	dm_build_618++
	return 4
}

func (Dm_build_621 *dm_build_598) Dm_build_620(dm_build_622 []byte, dm_build_623 int, dm_build_624 int64) int {
	dm_build_622[dm_build_623] = byte(dm_build_624)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 8)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 16)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 24)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 32)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 40)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 48)
	dm_build_623++
	dm_build_622[dm_build_623] = byte(dm_build_624 >> 56)
	return 8
}

func (Dm_build_626 *dm_build_598) Dm_build_625(dm_build_627 []byte, dm_build_628 int, dm_build_629 float32) int {
	return Dm_build_626.Dm_build_645(dm_build_627, dm_build_628, math.Float32bits(dm_build_629))
}

func (Dm_build_631 *dm_build_598) Dm_build_630(dm_build_632 []byte, dm_build_633 int, dm_build_634 float64) int {
	return Dm_build_631.Dm_build_650(dm_build_632, dm_build_633, math.Float64bits(dm_build_634))
}

func (Dm_build_636 *dm_build_598) Dm_build_635(dm_build_637 []byte, dm_build_638 int, dm_build_639 uint8) int {
	dm_build_637[dm_build_638] = byte(dm_build_639)
	return 1
}

func (Dm_build_641 *dm_build_598) Dm_build_640(dm_build_642 []byte, dm_build_643 int, dm_build_644 uint16) int {
	dm_build_642[dm_build_643] = byte(dm_build_644)
	dm_build_643++
	dm_build_642[dm_build_643] = byte(dm_build_644 >> 8)
	return 2
}

func (Dm_build_646 *dm_build_598) Dm_build_645(dm_build_647 []byte, dm_build_648 int, dm_build_649 uint32) int {
	dm_build_647[dm_build_648] = byte(dm_build_649)
	dm_build_648++
	dm_build_647[dm_build_648] = byte(dm_build_649 >> 8)
	dm_build_648++
	dm_build_647[dm_build_648] = byte(dm_build_649 >> 16)
	dm_build_648++
	dm_build_647[dm_build_648] = byte(dm_build_649 >> 24)
	return 3
}

func (Dm_build_651 *dm_build_598) Dm_build_650(dm_build_652 []byte, dm_build_653 int, dm_build_654 uint64) int {
	dm_build_652[dm_build_653] = byte(dm_build_654)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 8)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 16)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 24)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 32)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 40)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 48)
	dm_build_653++
	dm_build_652[dm_build_653] = byte(dm_build_654 >> 56)
	return 3
}

func (Dm_build_656 *dm_build_598) Dm_build_655(dm_build_657 []byte, dm_build_658 int, dm_build_659 []byte, dm_build_660 int, dm_build_661 int) int {
	copy(dm_build_657[dm_build_658:dm_build_658+dm_build_661], dm_build_659[dm_build_660:dm_build_660+dm_build_661])
	return dm_build_661
}

func (Dm_build_663 *dm_build_598) Dm_build_662(dm_build_664 []byte, dm_build_665 int, dm_build_666 []byte, dm_build_667 int, dm_build_668 int) int {
	dm_build_665 += Dm_build_663.Dm_build_645(dm_build_664, dm_build_665, uint32(dm_build_668))
	return 4 + Dm_build_663.Dm_build_655(dm_build_664, dm_build_665, dm_build_666, dm_build_667, dm_build_668)
}

func (Dm_build_670 *dm_build_598) Dm_build_669(dm_build_671 []byte, dm_build_672 int, dm_build_673 []byte, dm_build_674 int, dm_build_675 int) int {
	dm_build_672 += Dm_build_670.Dm_build_640(dm_build_671, dm_build_672, uint16(dm_build_675))
	return 2 + Dm_build_670.Dm_build_655(dm_build_671, dm_build_672, dm_build_673, dm_build_674, dm_build_675)
}

func (Dm_build_677 *dm_build_598) Dm_build_676(dm_build_678 []byte, dm_build_679 int, dm_build_680 string, dm_build_681 string, dm_build_682 *DmConnection) int {
	dm_build_683 := Dm_build_677.Dm_build_809(dm_build_680, dm_build_681, dm_build_682)
	dm_build_679 += Dm_build_677.Dm_build_645(dm_build_678, dm_build_679, uint32(len(dm_build_683)))
	return 4 + Dm_build_677.Dm_build_655(dm_build_678, dm_build_679, dm_build_683, 0, len(dm_build_683))
}

func (Dm_build_685 *dm_build_598) Dm_build_684(dm_build_686 []byte, dm_build_687 int, dm_build_688 string, dm_build_689 string, dm_build_690 *DmConnection) int {
	dm_build_691 := Dm_build_685.Dm_build_809(dm_build_688, dm_build_689, dm_build_690)

	dm_build_687 += Dm_build_685.Dm_build_640(dm_build_686, dm_build_687, uint16(len(dm_build_691)))
	return 2 + Dm_build_685.Dm_build_655(dm_build_686, dm_build_687, dm_build_691, 0, len(dm_build_691))
}

func (Dm_build_693 *dm_build_598) Dm_build_692(dm_build_694 []byte, dm_build_695 int) byte {
	return dm_build_694[dm_build_695]
}

func (Dm_build_697 *dm_build_598) Dm_build_696(dm_build_698 []byte, dm_build_699 int) int16 {
	var dm_build_700 int16
	dm_build_700 = int16(dm_build_698[dm_build_699] & 0xff)
	dm_build_699++
	dm_build_700 |= int16(dm_build_698[dm_build_699]&0xff) << 8
	return dm_build_700
}

func (Dm_build_702 *dm_build_598) Dm_build_701(dm_build_703 []byte, dm_build_704 int) int32 {
	var dm_build_705 int32
	dm_build_705 = int32(dm_build_703[dm_build_704] & 0xff)
	dm_build_704++
	dm_build_705 |= int32(dm_build_703[dm_build_704]&0xff) << 8
	dm_build_704++
	dm_build_705 |= int32(dm_build_703[dm_build_704]&0xff) << 16
	dm_build_704++
	dm_build_705 |= int32(dm_build_703[dm_build_704]&0xff) << 24
	return dm_build_705
}

func (Dm_build_707 *dm_build_598) Dm_build_706(dm_build_708 []byte, dm_build_709 int) int64 {
	var dm_build_710 int64
	dm_build_710 = int64(dm_build_708[dm_build_709] & 0xff)
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 8
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 16
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 24
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 32
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 40
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 48
	dm_build_709++
	dm_build_710 |= int64(dm_build_708[dm_build_709]&0xff) << 56
	return dm_build_710
}

func (Dm_build_712 *dm_build_598) Dm_build_711(dm_build_713 []byte, dm_build_714 int) float32 {
	return math.Float32frombits(Dm_build_712.Dm_build_728(dm_build_713, dm_build_714))
}

func (Dm_build_716 *dm_build_598) Dm_build_715(dm_build_717 []byte, dm_build_718 int) float64 {
	return math.Float64frombits(Dm_build_716.Dm_build_733(dm_build_717, dm_build_718))
}

func (Dm_build_720 *dm_build_598) Dm_build_719(dm_build_721 []byte, dm_build_722 int) uint8 {
	return uint8(dm_build_721[dm_build_722] & 0xff)
}

func (Dm_build_724 *dm_build_598) Dm_build_723(dm_build_725 []byte, dm_build_726 int) uint16 {
	var dm_build_727 uint16
	dm_build_727 = uint16(dm_build_725[dm_build_726] & 0xff)
	dm_build_726++
	dm_build_727 |= uint16(dm_build_725[dm_build_726]&0xff) << 8
	return dm_build_727
}

func (Dm_build_729 *dm_build_598) Dm_build_728(dm_build_730 []byte, dm_build_731 int) uint32 {
	var dm_build_732 uint32
	dm_build_732 = uint32(dm_build_730[dm_build_731] & 0xff)
	dm_build_731++
	dm_build_732 |= uint32(dm_build_730[dm_build_731]&0xff) << 8
	dm_build_731++
	dm_build_732 |= uint32(dm_build_730[dm_build_731]&0xff) << 16
	dm_build_731++
	dm_build_732 |= uint32(dm_build_730[dm_build_731]&0xff) << 24
	return dm_build_732
}

func (Dm_build_734 *dm_build_598) Dm_build_733(dm_build_735 []byte, dm_build_736 int) uint64 {
	var dm_build_737 uint64
	dm_build_737 = uint64(dm_build_735[dm_build_736] & 0xff)
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 8
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 16
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 24
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 32
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 40
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 48
	dm_build_736++
	dm_build_737 |= uint64(dm_build_735[dm_build_736]&0xff) << 56
	return dm_build_737
}

func (Dm_build_739 *dm_build_598) Dm_build_738(dm_build_740 []byte, dm_build_741 int) []byte {
	dm_build_742 := Dm_build_739.Dm_build_728(dm_build_740, dm_build_741)
	return dm_build_740[dm_build_741+4 : dm_build_741+4+int(dm_build_742)]

}

func (Dm_build_744 *dm_build_598) Dm_build_743(dm_build_745 []byte, dm_build_746 int) []byte {
	dm_build_747 := Dm_build_744.Dm_build_723(dm_build_745, dm_build_746)
	return dm_build_745[dm_build_746+2 : dm_build_746+2+int(dm_build_747)]

}

func (Dm_build_749 *dm_build_598) Dm_build_748(dm_build_750 []byte, dm_build_751 int, dm_build_752 int) []byte {
	return dm_build_750[dm_build_751 : dm_build_751+dm_build_752]

}

func (Dm_build_754 *dm_build_598) Dm_build_753(dm_build_755 []byte, dm_build_756 int, dm_build_757 int, dm_build_758 string, dm_build_759 *DmConnection) string {
	return Dm_build_754.Dm_build_846(dm_build_755[dm_build_756:dm_build_756+dm_build_757], dm_build_758, dm_build_759)
}

func (Dm_build_761 *dm_build_598) Dm_build_760(dm_build_762 []byte, dm_build_763 int, dm_build_764 string, dm_build_765 *DmConnection) string {
	dm_build_766 := Dm_build_761.Dm_build_728(dm_build_762, dm_build_763)
	dm_build_763 += 4
	return Dm_build_761.Dm_build_753(dm_build_762, dm_build_763, int(dm_build_766), dm_build_764, dm_build_765)
}

func (Dm_build_768 *dm_build_598) Dm_build_767(dm_build_769 []byte, dm_build_770 int, dm_build_771 string, dm_build_772 *DmConnection) string {
	dm_build_773 := Dm_build_768.Dm_build_723(dm_build_769, dm_build_770)
	dm_build_770 += 2
	return Dm_build_768.Dm_build_753(dm_build_769, dm_build_770, int(dm_build_773), dm_build_771, dm_build_772)
}

func (Dm_build_775 *dm_build_598) Dm_build_774(dm_build_776 byte) []byte {
	return []byte{dm_build_776}
}

func (Dm_build_778 *dm_build_598) Dm_build_777(dm_build_779 int16) []byte {
	return []byte{byte(dm_build_779), byte(dm_build_779 >> 8)}
}

func (Dm_build_781 *dm_build_598) Dm_build_780(dm_build_782 int32) []byte {
	return []byte{byte(dm_build_782), byte(dm_build_782 >> 8), byte(dm_build_782 >> 16), byte(dm_build_782 >> 24)}
}

func (Dm_build_784 *dm_build_598) Dm_build_783(dm_build_785 int64) []byte {
	return []byte{byte(dm_build_785), byte(dm_build_785 >> 8), byte(dm_build_785 >> 16), byte(dm_build_785 >> 24), byte(dm_build_785 >> 32),
		byte(dm_build_785 >> 40), byte(dm_build_785 >> 48), byte(dm_build_785 >> 56)}
}

func (Dm_build_787 *dm_build_598) Dm_build_786(dm_build_788 float32) []byte {
	return Dm_build_787.Dm_build_798(math.Float32bits(dm_build_788))
}

func (Dm_build_790 *dm_build_598) Dm_build_789(dm_build_791 float64) []byte {
	return Dm_build_790.Dm_build_801(math.Float64bits(dm_build_791))
}

func (Dm_build_793 *dm_build_598) Dm_build_792(dm_build_794 uint8) []byte {
	return []byte{byte(dm_build_794)}
}

func (Dm_build_796 *dm_build_598) Dm_build_795(dm_build_797 uint16) []byte {
	return []byte{byte(dm_build_797), byte(dm_build_797 >> 8)}
}

func (Dm_build_799 *dm_build_598) Dm_build_798(dm_build_800 uint32) []byte {
	return []byte{byte(dm_build_800), byte(dm_build_800 >> 8), byte(dm_build_800 >> 16), byte(dm_build_800 >> 24)}
}

func (Dm_build_802 *dm_build_598) Dm_build_801(dm_build_803 uint64) []byte {
	return []byte{byte(dm_build_803), byte(dm_build_803 >> 8), byte(dm_build_803 >> 16), byte(dm_build_803 >> 24), byte(dm_build_803 >> 32), byte(dm_build_803 >> 40), byte(dm_build_803 >> 48), byte(dm_build_803 >> 56)}
}

func (Dm_build_805 *dm_build_598) Dm_build_804(dm_build_806 []byte, dm_build_807 string, dm_build_808 *DmConnection) []byte {
	if dm_build_807 == "UTF-8" {
		return dm_build_806
	}

	if dm_build_808 == nil {
		if e := dm_build_851(dm_build_807); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_806), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_808.encodeBuffer == nil {
		dm_build_808.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_808.encode = dm_build_851(dm_build_808.getServerEncoding())
		dm_build_808.transformReaderDst = make([]byte, 4096)
		dm_build_808.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_808.encode; e != nil {

		dm_build_808.encodeBuffer.Reset()

		n, err := dm_build_808.encodeBuffer.ReadFrom(
			Dm_build_865(bytes.NewReader(dm_build_806), e.NewEncoder(), dm_build_808.transformReaderDst, dm_build_808.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_808.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_810 *dm_build_598) Dm_build_809(dm_build_811 string, dm_build_812 string, dm_build_813 *DmConnection) []byte {
	return Dm_build_810.Dm_build_804([]byte(dm_build_811), dm_build_812, dm_build_813)
}

func (Dm_build_815 *dm_build_598) Dm_build_814(dm_build_816 []byte) byte {
	return Dm_build_815.Dm_build_692(dm_build_816, 0)
}

func (Dm_build_818 *dm_build_598) Dm_build_817(dm_build_819 []byte) int16 {
	return Dm_build_818.Dm_build_696(dm_build_819, 0)
}

func (Dm_build_821 *dm_build_598) Dm_build_820(dm_build_822 []byte) int32 {
	return Dm_build_821.Dm_build_701(dm_build_822, 0)
}

func (Dm_build_824 *dm_build_598) Dm_build_823(dm_build_825 []byte) int64 {
	return Dm_build_824.Dm_build_706(dm_build_825, 0)
}

func (Dm_build_827 *dm_build_598) Dm_build_826(dm_build_828 []byte) float32 {
	return Dm_build_827.Dm_build_711(dm_build_828, 0)
}

func (Dm_build_830 *dm_build_598) Dm_build_829(dm_build_831 []byte) float64 {
	return Dm_build_830.Dm_build_715(dm_build_831, 0)
}

func (Dm_build_833 *dm_build_598) Dm_build_832(dm_build_834 []byte) uint8 {
	return Dm_build_833.Dm_build_719(dm_build_834, 0)
}

func (Dm_build_836 *dm_build_598) Dm_build_835(dm_build_837 []byte) uint16 {
	return Dm_build_836.Dm_build_723(dm_build_837, 0)
}

func (Dm_build_839 *dm_build_598) Dm_build_838(dm_build_840 []byte) uint32 {
	return Dm_build_839.Dm_build_728(dm_build_840, 0)
}

func (Dm_build_842 *dm_build_598) Dm_build_841(dm_build_843 []byte, dm_build_844 string, dm_build_845 *DmConnection) []byte {
	if dm_build_844 == "UTF-8" {
		return dm_build_843
	}

	if dm_build_845 == nil {
		if e := dm_build_851(dm_build_844); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_843), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_845.encodeBuffer == nil {
		dm_build_845.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_845.encode = dm_build_851(dm_build_845.getServerEncoding())
		dm_build_845.transformReaderDst = make([]byte, 4096)
		dm_build_845.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_845.encode; e != nil {

		dm_build_845.encodeBuffer.Reset()

		n, err := dm_build_845.encodeBuffer.ReadFrom(
			Dm_build_865(bytes.NewReader(dm_build_843), e.NewDecoder(), dm_build_845.transformReaderDst, dm_build_845.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_845.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_847 *dm_build_598) Dm_build_846(dm_build_848 []byte, dm_build_849 string, dm_build_850 *DmConnection) string {
	return string(Dm_build_847.Dm_build_841(dm_build_848, dm_build_849, dm_build_850))
}

func dm_build_851(dm_build_852 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_852); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_853 struct {
	dm_build_854 io.Reader
	dm_build_855 transform.Transformer
	dm_build_856 error

	dm_build_857               []byte
	dm_build_858, dm_build_859 int

	dm_build_860               []byte
	dm_build_861, dm_build_862 int

	dm_build_863 bool
}

const dm_build_864 = 4096

func Dm_build_865(dm_build_866 io.Reader, dm_build_867 transform.Transformer, dm_build_868 []byte, dm_build_869 []byte) *Dm_build_853 {
	dm_build_867.Reset()
	return &Dm_build_853{
		dm_build_854: dm_build_866,
		dm_build_855: dm_build_867,
		dm_build_857: dm_build_868,
		dm_build_860: dm_build_869,
	}
}

func (dm_build_871 *Dm_build_853) Read(dm_build_872 []byte) (int, error) {
	dm_build_873, dm_build_874 := 0, error(nil)
	for {

		if dm_build_871.dm_build_858 != dm_build_871.dm_build_859 {
			dm_build_873 = copy(dm_build_872, dm_build_871.dm_build_857[dm_build_871.dm_build_858:dm_build_871.dm_build_859])
			dm_build_871.dm_build_858 += dm_build_873
			if dm_build_871.dm_build_858 == dm_build_871.dm_build_859 && dm_build_871.dm_build_863 {
				return dm_build_873, dm_build_871.dm_build_856
			}
			return dm_build_873, nil
		} else if dm_build_871.dm_build_863 {
			return 0, dm_build_871.dm_build_856
		}

		if dm_build_871.dm_build_861 != dm_build_871.dm_build_862 || dm_build_871.dm_build_856 != nil {
			dm_build_871.dm_build_858 = 0
			dm_build_871.dm_build_859, dm_build_873, dm_build_874 = dm_build_871.dm_build_855.Transform(dm_build_871.dm_build_857, dm_build_871.dm_build_860[dm_build_871.dm_build_861:dm_build_871.dm_build_862], dm_build_871.dm_build_856 == io.EOF)
			dm_build_871.dm_build_861 += dm_build_873

			switch {
			case dm_build_874 == nil:
				if dm_build_871.dm_build_861 != dm_build_871.dm_build_862 {
					dm_build_871.dm_build_856 = nil
				}

				dm_build_871.dm_build_863 = dm_build_871.dm_build_856 != nil
				continue
			case dm_build_874 == transform.ErrShortDst && (dm_build_871.dm_build_859 != 0 || dm_build_873 != 0):

				continue
			case dm_build_874 == transform.ErrShortSrc && dm_build_871.dm_build_862-dm_build_871.dm_build_861 != len(dm_build_871.dm_build_860) && dm_build_871.dm_build_856 == nil:

			default:
				dm_build_871.dm_build_863 = true

				if dm_build_871.dm_build_856 == nil || dm_build_871.dm_build_856 == io.EOF {
					dm_build_871.dm_build_856 = dm_build_874
				}
				continue
			}
		}

		if dm_build_871.dm_build_861 != 0 {
			dm_build_871.dm_build_861, dm_build_871.dm_build_862 = 0, copy(dm_build_871.dm_build_860, dm_build_871.dm_build_860[dm_build_871.dm_build_861:dm_build_871.dm_build_862])
		}
		dm_build_873, dm_build_871.dm_build_856 = dm_build_871.dm_build_854.Read(dm_build_871.dm_build_860[dm_build_871.dm_build_862:])
		dm_build_871.dm_build_862 += dm_build_873
	}
}
