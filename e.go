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

type dm_build_648 struct{}

var Dm_build_649 = &dm_build_648{}

func (Dm_build_651 *dm_build_648) Dm_build_650(dm_build_652 []byte, dm_build_653 int, dm_build_654 byte) int {
	dm_build_652[dm_build_653] = dm_build_654
	return 1
}

func (Dm_build_656 *dm_build_648) Dm_build_655(dm_build_657 []byte, dm_build_658 int, dm_build_659 int8) int {
	dm_build_657[dm_build_658] = byte(dm_build_659)
	return 1
}

func (Dm_build_661 *dm_build_648) Dm_build_660(dm_build_662 []byte, dm_build_663 int, dm_build_664 int16) int {
	dm_build_662[dm_build_663] = byte(dm_build_664)
	dm_build_663++
	dm_build_662[dm_build_663] = byte(dm_build_664 >> 8)
	return 2
}

func (Dm_build_666 *dm_build_648) Dm_build_665(dm_build_667 []byte, dm_build_668 int, dm_build_669 int32) int {
	dm_build_667[dm_build_668] = byte(dm_build_669)
	dm_build_668++
	dm_build_667[dm_build_668] = byte(dm_build_669 >> 8)
	dm_build_668++
	dm_build_667[dm_build_668] = byte(dm_build_669 >> 16)
	dm_build_668++
	dm_build_667[dm_build_668] = byte(dm_build_669 >> 24)
	dm_build_668++
	return 4
}

func (Dm_build_671 *dm_build_648) Dm_build_670(dm_build_672 []byte, dm_build_673 int, dm_build_674 int64) int {
	dm_build_672[dm_build_673] = byte(dm_build_674)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 8)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 16)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 24)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 32)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 40)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 48)
	dm_build_673++
	dm_build_672[dm_build_673] = byte(dm_build_674 >> 56)
	return 8
}

func (Dm_build_676 *dm_build_648) Dm_build_675(dm_build_677 []byte, dm_build_678 int, dm_build_679 float32) int {
	return Dm_build_676.Dm_build_695(dm_build_677, dm_build_678, math.Float32bits(dm_build_679))
}

func (Dm_build_681 *dm_build_648) Dm_build_680(dm_build_682 []byte, dm_build_683 int, dm_build_684 float64) int {
	return Dm_build_681.Dm_build_700(dm_build_682, dm_build_683, math.Float64bits(dm_build_684))
}

func (Dm_build_686 *dm_build_648) Dm_build_685(dm_build_687 []byte, dm_build_688 int, dm_build_689 uint8) int {
	dm_build_687[dm_build_688] = byte(dm_build_689)
	return 1
}

func (Dm_build_691 *dm_build_648) Dm_build_690(dm_build_692 []byte, dm_build_693 int, dm_build_694 uint16) int {
	dm_build_692[dm_build_693] = byte(dm_build_694)
	dm_build_693++
	dm_build_692[dm_build_693] = byte(dm_build_694 >> 8)
	return 2
}

func (Dm_build_696 *dm_build_648) Dm_build_695(dm_build_697 []byte, dm_build_698 int, dm_build_699 uint32) int {
	dm_build_697[dm_build_698] = byte(dm_build_699)
	dm_build_698++
	dm_build_697[dm_build_698] = byte(dm_build_699 >> 8)
	dm_build_698++
	dm_build_697[dm_build_698] = byte(dm_build_699 >> 16)
	dm_build_698++
	dm_build_697[dm_build_698] = byte(dm_build_699 >> 24)
	return 3
}

func (Dm_build_701 *dm_build_648) Dm_build_700(dm_build_702 []byte, dm_build_703 int, dm_build_704 uint64) int {
	dm_build_702[dm_build_703] = byte(dm_build_704)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 8)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 16)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 24)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 32)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 40)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 48)
	dm_build_703++
	dm_build_702[dm_build_703] = byte(dm_build_704 >> 56)
	return 3
}

func (Dm_build_706 *dm_build_648) Dm_build_705(dm_build_707 []byte, dm_build_708 int, dm_build_709 []byte, dm_build_710 int, dm_build_711 int) int {
	copy(dm_build_707[dm_build_708:dm_build_708+dm_build_711], dm_build_709[dm_build_710:dm_build_710+dm_build_711])
	return dm_build_711
}

func (Dm_build_713 *dm_build_648) Dm_build_712(dm_build_714 []byte, dm_build_715 int, dm_build_716 []byte, dm_build_717 int, dm_build_718 int) int {
	dm_build_715 += Dm_build_713.Dm_build_695(dm_build_714, dm_build_715, uint32(dm_build_718))
	return 4 + Dm_build_713.Dm_build_705(dm_build_714, dm_build_715, dm_build_716, dm_build_717, dm_build_718)
}

func (Dm_build_720 *dm_build_648) Dm_build_719(dm_build_721 []byte, dm_build_722 int, dm_build_723 []byte, dm_build_724 int, dm_build_725 int) int {
	dm_build_722 += Dm_build_720.Dm_build_690(dm_build_721, dm_build_722, uint16(dm_build_725))
	return 2 + Dm_build_720.Dm_build_705(dm_build_721, dm_build_722, dm_build_723, dm_build_724, dm_build_725)
}

func (Dm_build_727 *dm_build_648) Dm_build_726(dm_build_728 []byte, dm_build_729 int, dm_build_730 string, dm_build_731 string, dm_build_732 *DmConnection) int {
	dm_build_733 := Dm_build_727.Dm_build_865(dm_build_730, dm_build_731, dm_build_732)
	dm_build_729 += Dm_build_727.Dm_build_695(dm_build_728, dm_build_729, uint32(len(dm_build_733)))
	return 4 + Dm_build_727.Dm_build_705(dm_build_728, dm_build_729, dm_build_733, 0, len(dm_build_733))
}

func (Dm_build_735 *dm_build_648) Dm_build_734(dm_build_736 []byte, dm_build_737 int, dm_build_738 string, dm_build_739 string, dm_build_740 *DmConnection) int {
	dm_build_741 := Dm_build_735.Dm_build_865(dm_build_738, dm_build_739, dm_build_740)

	dm_build_737 += Dm_build_735.Dm_build_690(dm_build_736, dm_build_737, uint16(len(dm_build_741)))
	return 2 + Dm_build_735.Dm_build_705(dm_build_736, dm_build_737, dm_build_741, 0, len(dm_build_741))
}

func (Dm_build_743 *dm_build_648) Dm_build_742(dm_build_744 []byte, dm_build_745 int) byte {
	return dm_build_744[dm_build_745]
}

func (Dm_build_747 *dm_build_648) Dm_build_746(dm_build_748 []byte, dm_build_749 int) int16 {
	var dm_build_750 int16
	dm_build_750 = int16(dm_build_748[dm_build_749] & 0xff)
	dm_build_749++
	dm_build_750 |= int16(dm_build_748[dm_build_749]&0xff) << 8
	return dm_build_750
}

func (Dm_build_752 *dm_build_648) Dm_build_751(dm_build_753 []byte, dm_build_754 int) int32 {
	var dm_build_755 int32
	dm_build_755 = int32(dm_build_753[dm_build_754] & 0xff)
	dm_build_754++
	dm_build_755 |= int32(dm_build_753[dm_build_754]&0xff) << 8
	dm_build_754++
	dm_build_755 |= int32(dm_build_753[dm_build_754]&0xff) << 16
	dm_build_754++
	dm_build_755 |= int32(dm_build_753[dm_build_754]&0xff) << 24
	return dm_build_755
}

func (Dm_build_757 *dm_build_648) Dm_build_756(dm_build_758 []byte, dm_build_759 int) int64 {
	var dm_build_760 int64
	dm_build_760 = int64(dm_build_758[dm_build_759] & 0xff)
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 8
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 16
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 24
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 32
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 40
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 48
	dm_build_759++
	dm_build_760 |= int64(dm_build_758[dm_build_759]&0xff) << 56
	return dm_build_760
}

func (Dm_build_762 *dm_build_648) Dm_build_761(dm_build_763 []byte, dm_build_764 int) float32 {
	return math.Float32frombits(Dm_build_762.Dm_build_778(dm_build_763, dm_build_764))
}

func (Dm_build_766 *dm_build_648) Dm_build_765(dm_build_767 []byte, dm_build_768 int) float64 {
	return math.Float64frombits(Dm_build_766.Dm_build_783(dm_build_767, dm_build_768))
}

func (Dm_build_770 *dm_build_648) Dm_build_769(dm_build_771 []byte, dm_build_772 int) uint8 {
	return uint8(dm_build_771[dm_build_772] & 0xff)
}

func (Dm_build_774 *dm_build_648) Dm_build_773(dm_build_775 []byte, dm_build_776 int) uint16 {
	var dm_build_777 uint16
	dm_build_777 = uint16(dm_build_775[dm_build_776] & 0xff)
	dm_build_776++
	dm_build_777 |= uint16(dm_build_775[dm_build_776]&0xff) << 8
	return dm_build_777
}

func (Dm_build_779 *dm_build_648) Dm_build_778(dm_build_780 []byte, dm_build_781 int) uint32 {
	var dm_build_782 uint32
	dm_build_782 = uint32(dm_build_780[dm_build_781] & 0xff)
	dm_build_781++
	dm_build_782 |= uint32(dm_build_780[dm_build_781]&0xff) << 8
	dm_build_781++
	dm_build_782 |= uint32(dm_build_780[dm_build_781]&0xff) << 16
	dm_build_781++
	dm_build_782 |= uint32(dm_build_780[dm_build_781]&0xff) << 24
	return dm_build_782
}

func (Dm_build_784 *dm_build_648) Dm_build_783(dm_build_785 []byte, dm_build_786 int) uint64 {
	var dm_build_787 uint64
	dm_build_787 = uint64(dm_build_785[dm_build_786] & 0xff)
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 8
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 16
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 24
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 32
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 40
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 48
	dm_build_786++
	dm_build_787 |= uint64(dm_build_785[dm_build_786]&0xff) << 56
	return dm_build_787
}

func (Dm_build_789 *dm_build_648) Dm_build_788(dm_build_790 []byte, dm_build_791 int) []byte {
	dm_build_792 := Dm_build_789.Dm_build_778(dm_build_790, dm_build_791)

	dm_build_793 := make([]byte, dm_build_792)
	copy(dm_build_793[:int(dm_build_792)], dm_build_790[dm_build_791+4:dm_build_791+4+int(dm_build_792)])
	return dm_build_793
}

func (Dm_build_795 *dm_build_648) Dm_build_794(dm_build_796 []byte, dm_build_797 int) []byte {
	dm_build_798 := Dm_build_795.Dm_build_773(dm_build_796, dm_build_797)

	dm_build_799 := make([]byte, dm_build_798)
	copy(dm_build_799[:int(dm_build_798)], dm_build_796[dm_build_797+2:dm_build_797+2+int(dm_build_798)])
	return dm_build_799
}

func (Dm_build_801 *dm_build_648) Dm_build_800(dm_build_802 []byte, dm_build_803 int, dm_build_804 int) []byte {

	dm_build_805 := make([]byte, dm_build_804)
	copy(dm_build_805[:dm_build_804], dm_build_802[dm_build_803:dm_build_803+dm_build_804])
	return dm_build_805
}

func (Dm_build_807 *dm_build_648) Dm_build_806(dm_build_808 []byte, dm_build_809 int, dm_build_810 int, dm_build_811 string, dm_build_812 *DmConnection) string {
	return Dm_build_807.Dm_build_902(dm_build_808[dm_build_809:dm_build_809+dm_build_810], dm_build_811, dm_build_812)
}

func (Dm_build_814 *dm_build_648) Dm_build_813(dm_build_815 []byte, dm_build_816 int, dm_build_817 string, dm_build_818 *DmConnection) string {
	dm_build_819 := Dm_build_814.Dm_build_778(dm_build_815, dm_build_816)
	dm_build_816 += 4
	return Dm_build_814.Dm_build_806(dm_build_815, dm_build_816, int(dm_build_819), dm_build_817, dm_build_818)
}

func (Dm_build_821 *dm_build_648) Dm_build_820(dm_build_822 []byte, dm_build_823 int, dm_build_824 string, dm_build_825 *DmConnection) string {
	dm_build_826 := Dm_build_821.Dm_build_773(dm_build_822, dm_build_823)
	dm_build_823 += 2
	return Dm_build_821.Dm_build_806(dm_build_822, dm_build_823, int(dm_build_826), dm_build_824, dm_build_825)
}

func (Dm_build_828 *dm_build_648) Dm_build_827(dm_build_829 byte) []byte {
	return []byte{dm_build_829}
}

func (Dm_build_831 *dm_build_648) Dm_build_830(dm_build_832 int8) []byte {
	return []byte{byte(dm_build_832)}
}

func (Dm_build_834 *dm_build_648) Dm_build_833(dm_build_835 int16) []byte {
	return []byte{byte(dm_build_835), byte(dm_build_835 >> 8)}
}

func (Dm_build_837 *dm_build_648) Dm_build_836(dm_build_838 int32) []byte {
	return []byte{byte(dm_build_838), byte(dm_build_838 >> 8), byte(dm_build_838 >> 16), byte(dm_build_838 >> 24)}
}

func (Dm_build_840 *dm_build_648) Dm_build_839(dm_build_841 int64) []byte {
	return []byte{byte(dm_build_841), byte(dm_build_841 >> 8), byte(dm_build_841 >> 16), byte(dm_build_841 >> 24), byte(dm_build_841 >> 32),
		byte(dm_build_841 >> 40), byte(dm_build_841 >> 48), byte(dm_build_841 >> 56)}
}

func (Dm_build_843 *dm_build_648) Dm_build_842(dm_build_844 float32) []byte {
	return Dm_build_843.Dm_build_854(math.Float32bits(dm_build_844))
}

func (Dm_build_846 *dm_build_648) Dm_build_845(dm_build_847 float64) []byte {
	return Dm_build_846.Dm_build_857(math.Float64bits(dm_build_847))
}

func (Dm_build_849 *dm_build_648) Dm_build_848(dm_build_850 uint8) []byte {
	return []byte{byte(dm_build_850)}
}

func (Dm_build_852 *dm_build_648) Dm_build_851(dm_build_853 uint16) []byte {
	return []byte{byte(dm_build_853), byte(dm_build_853 >> 8)}
}

func (Dm_build_855 *dm_build_648) Dm_build_854(dm_build_856 uint32) []byte {
	return []byte{byte(dm_build_856), byte(dm_build_856 >> 8), byte(dm_build_856 >> 16), byte(dm_build_856 >> 24)}
}

func (Dm_build_858 *dm_build_648) Dm_build_857(dm_build_859 uint64) []byte {
	return []byte{byte(dm_build_859), byte(dm_build_859 >> 8), byte(dm_build_859 >> 16), byte(dm_build_859 >> 24), byte(dm_build_859 >> 32), byte(dm_build_859 >> 40), byte(dm_build_859 >> 48), byte(dm_build_859 >> 56)}
}

func (Dm_build_861 *dm_build_648) Dm_build_860(dm_build_862 []byte, dm_build_863 string, dm_build_864 *DmConnection) []byte {
	if dm_build_863 == "UTF-8" {
		return dm_build_862
	}

	if dm_build_864 == nil {
		if e := dm_build_907(dm_build_863); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_862), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_864.encodeBuffer == nil {
		dm_build_864.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_864.encode = dm_build_907(dm_build_864.getServerEncoding())
		dm_build_864.transformReaderDst = make([]byte, 4096)
		dm_build_864.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_864.encode; e != nil {

		dm_build_864.encodeBuffer.Reset()

		n, err := dm_build_864.encodeBuffer.ReadFrom(
			Dm_build_921(bytes.NewReader(dm_build_862), e.NewEncoder(), dm_build_864.transformReaderDst, dm_build_864.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_864.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_866 *dm_build_648) Dm_build_865(dm_build_867 string, dm_build_868 string, dm_build_869 *DmConnection) []byte {
	return Dm_build_866.Dm_build_860([]byte(dm_build_867), dm_build_868, dm_build_869)
}

func (Dm_build_871 *dm_build_648) Dm_build_870(dm_build_872 []byte) byte {
	return Dm_build_871.Dm_build_742(dm_build_872, 0)
}

func (Dm_build_874 *dm_build_648) Dm_build_873(dm_build_875 []byte) int16 {
	return Dm_build_874.Dm_build_746(dm_build_875, 0)
}

func (Dm_build_877 *dm_build_648) Dm_build_876(dm_build_878 []byte) int32 {
	return Dm_build_877.Dm_build_751(dm_build_878, 0)
}

func (Dm_build_880 *dm_build_648) Dm_build_879(dm_build_881 []byte) int64 {
	return Dm_build_880.Dm_build_756(dm_build_881, 0)
}

func (Dm_build_883 *dm_build_648) Dm_build_882(dm_build_884 []byte) float32 {
	return Dm_build_883.Dm_build_761(dm_build_884, 0)
}

func (Dm_build_886 *dm_build_648) Dm_build_885(dm_build_887 []byte) float64 {
	return Dm_build_886.Dm_build_765(dm_build_887, 0)
}

func (Dm_build_889 *dm_build_648) Dm_build_888(dm_build_890 []byte) uint8 {
	return Dm_build_889.Dm_build_769(dm_build_890, 0)
}

func (Dm_build_892 *dm_build_648) Dm_build_891(dm_build_893 []byte) uint16 {
	return Dm_build_892.Dm_build_773(dm_build_893, 0)
}

func (Dm_build_895 *dm_build_648) Dm_build_894(dm_build_896 []byte) uint32 {
	return Dm_build_895.Dm_build_778(dm_build_896, 0)
}

func (Dm_build_898 *dm_build_648) Dm_build_897(dm_build_899 []byte, dm_build_900 string, dm_build_901 *DmConnection) []byte {
	if dm_build_900 == "UTF-8" {
		return dm_build_899
	}

	if dm_build_901 == nil {
		if e := dm_build_907(dm_build_900); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_899), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_901.encodeBuffer == nil {
		dm_build_901.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_901.encode = dm_build_907(dm_build_901.getServerEncoding())
		dm_build_901.transformReaderDst = make([]byte, 4096)
		dm_build_901.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_901.encode; e != nil {

		dm_build_901.encodeBuffer.Reset()

		n, err := dm_build_901.encodeBuffer.ReadFrom(
			Dm_build_921(bytes.NewReader(dm_build_899), e.NewDecoder(), dm_build_901.transformReaderDst, dm_build_901.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_901.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_903 *dm_build_648) Dm_build_902(dm_build_904 []byte, dm_build_905 string, dm_build_906 *DmConnection) string {
	return string(Dm_build_903.Dm_build_897(dm_build_904, dm_build_905, dm_build_906))
}

func dm_build_907(dm_build_908 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_908); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_909 struct {
	dm_build_910 io.Reader
	dm_build_911 transform.Transformer
	dm_build_912 error

	dm_build_913               []byte
	dm_build_914, dm_build_915 int

	dm_build_916               []byte
	dm_build_917, dm_build_918 int

	dm_build_919 bool
}

const dm_build_920 = 4096

func Dm_build_921(dm_build_922 io.Reader, dm_build_923 transform.Transformer, dm_build_924 []byte, dm_build_925 []byte) *Dm_build_909 {
	dm_build_923.Reset()
	return &Dm_build_909{
		dm_build_910: dm_build_922,
		dm_build_911: dm_build_923,
		dm_build_913: dm_build_924,
		dm_build_916: dm_build_925,
	}
}

func (dm_build_927 *Dm_build_909) Read(dm_build_928 []byte) (int, error) {
	dm_build_929, dm_build_930 := 0, error(nil)
	for {

		if dm_build_927.dm_build_914 != dm_build_927.dm_build_915 {
			dm_build_929 = copy(dm_build_928, dm_build_927.dm_build_913[dm_build_927.dm_build_914:dm_build_927.dm_build_915])
			dm_build_927.dm_build_914 += dm_build_929
			if dm_build_927.dm_build_914 == dm_build_927.dm_build_915 && dm_build_927.dm_build_919 {
				return dm_build_929, dm_build_927.dm_build_912
			}
			return dm_build_929, nil
		} else if dm_build_927.dm_build_919 {
			return 0, dm_build_927.dm_build_912
		}

		if dm_build_927.dm_build_917 != dm_build_927.dm_build_918 || dm_build_927.dm_build_912 != nil {
			dm_build_927.dm_build_914 = 0
			dm_build_927.dm_build_915, dm_build_929, dm_build_930 = dm_build_927.dm_build_911.Transform(dm_build_927.dm_build_913, dm_build_927.dm_build_916[dm_build_927.dm_build_917:dm_build_927.dm_build_918], dm_build_927.dm_build_912 == io.EOF)
			dm_build_927.dm_build_917 += dm_build_929

			switch {
			case dm_build_930 == nil:
				if dm_build_927.dm_build_917 != dm_build_927.dm_build_918 {
					dm_build_927.dm_build_912 = nil
				}

				dm_build_927.dm_build_919 = dm_build_927.dm_build_912 != nil
				continue
			case dm_build_930 == transform.ErrShortDst && (dm_build_927.dm_build_915 != 0 || dm_build_929 != 0):

				continue
			case dm_build_930 == transform.ErrShortSrc && dm_build_927.dm_build_918-dm_build_927.dm_build_917 != len(dm_build_927.dm_build_916) && dm_build_927.dm_build_912 == nil:

			default:
				dm_build_927.dm_build_919 = true

				if dm_build_927.dm_build_912 == nil || dm_build_927.dm_build_912 == io.EOF {
					dm_build_927.dm_build_912 = dm_build_930
				}
				continue
			}
		}

		if dm_build_927.dm_build_917 != 0 {
			dm_build_927.dm_build_917, dm_build_927.dm_build_918 = 0, copy(dm_build_927.dm_build_916, dm_build_927.dm_build_916[dm_build_927.dm_build_917:dm_build_927.dm_build_918])
		}
		dm_build_929, dm_build_927.dm_build_912 = dm_build_927.dm_build_910.Read(dm_build_927.dm_build_916[dm_build_927.dm_build_918:])
		dm_build_927.dm_build_918 += dm_build_929
	}
}
