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

type dm_build_622 struct{}

var Dm_build_623 = &dm_build_622{}

func (Dm_build_625 *dm_build_622) Dm_build_624(dm_build_626 []byte, dm_build_627 int, dm_build_628 byte) int {
	dm_build_626[dm_build_627] = dm_build_628
	return 1
}

func (Dm_build_630 *dm_build_622) Dm_build_629(dm_build_631 []byte, dm_build_632 int, dm_build_633 int8) int {
	dm_build_631[dm_build_632] = byte(dm_build_633)
	return 1
}

func (Dm_build_635 *dm_build_622) Dm_build_634(dm_build_636 []byte, dm_build_637 int, dm_build_638 int16) int {
	dm_build_636[dm_build_637] = byte(dm_build_638)
	dm_build_637++
	dm_build_636[dm_build_637] = byte(dm_build_638 >> 8)
	return 2
}

func (Dm_build_640 *dm_build_622) Dm_build_639(dm_build_641 []byte, dm_build_642 int, dm_build_643 int32) int {
	dm_build_641[dm_build_642] = byte(dm_build_643)
	dm_build_642++
	dm_build_641[dm_build_642] = byte(dm_build_643 >> 8)
	dm_build_642++
	dm_build_641[dm_build_642] = byte(dm_build_643 >> 16)
	dm_build_642++
	dm_build_641[dm_build_642] = byte(dm_build_643 >> 24)
	dm_build_642++
	return 4
}

func (Dm_build_645 *dm_build_622) Dm_build_644(dm_build_646 []byte, dm_build_647 int, dm_build_648 int64) int {
	dm_build_646[dm_build_647] = byte(dm_build_648)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 8)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 16)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 24)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 32)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 40)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 48)
	dm_build_647++
	dm_build_646[dm_build_647] = byte(dm_build_648 >> 56)
	return 8
}

func (Dm_build_650 *dm_build_622) Dm_build_649(dm_build_651 []byte, dm_build_652 int, dm_build_653 float32) int {
	return Dm_build_650.Dm_build_669(dm_build_651, dm_build_652, math.Float32bits(dm_build_653))
}

func (Dm_build_655 *dm_build_622) Dm_build_654(dm_build_656 []byte, dm_build_657 int, dm_build_658 float64) int {
	return Dm_build_655.Dm_build_674(dm_build_656, dm_build_657, math.Float64bits(dm_build_658))
}

func (Dm_build_660 *dm_build_622) Dm_build_659(dm_build_661 []byte, dm_build_662 int, dm_build_663 uint8) int {
	dm_build_661[dm_build_662] = byte(dm_build_663)
	return 1
}

func (Dm_build_665 *dm_build_622) Dm_build_664(dm_build_666 []byte, dm_build_667 int, dm_build_668 uint16) int {
	dm_build_666[dm_build_667] = byte(dm_build_668)
	dm_build_667++
	dm_build_666[dm_build_667] = byte(dm_build_668 >> 8)
	return 2
}

func (Dm_build_670 *dm_build_622) Dm_build_669(dm_build_671 []byte, dm_build_672 int, dm_build_673 uint32) int {
	dm_build_671[dm_build_672] = byte(dm_build_673)
	dm_build_672++
	dm_build_671[dm_build_672] = byte(dm_build_673 >> 8)
	dm_build_672++
	dm_build_671[dm_build_672] = byte(dm_build_673 >> 16)
	dm_build_672++
	dm_build_671[dm_build_672] = byte(dm_build_673 >> 24)
	return 3
}

func (Dm_build_675 *dm_build_622) Dm_build_674(dm_build_676 []byte, dm_build_677 int, dm_build_678 uint64) int {
	dm_build_676[dm_build_677] = byte(dm_build_678)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 8)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 16)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 24)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 32)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 40)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 48)
	dm_build_677++
	dm_build_676[dm_build_677] = byte(dm_build_678 >> 56)
	return 3
}

func (Dm_build_680 *dm_build_622) Dm_build_679(dm_build_681 []byte, dm_build_682 int, dm_build_683 []byte, dm_build_684 int, dm_build_685 int) int {
	copy(dm_build_681[dm_build_682:dm_build_682+dm_build_685], dm_build_683[dm_build_684:dm_build_684+dm_build_685])
	return dm_build_685
}

func (Dm_build_687 *dm_build_622) Dm_build_686(dm_build_688 []byte, dm_build_689 int, dm_build_690 []byte, dm_build_691 int, dm_build_692 int) int {
	dm_build_689 += Dm_build_687.Dm_build_669(dm_build_688, dm_build_689, uint32(dm_build_692))
	return 4 + Dm_build_687.Dm_build_679(dm_build_688, dm_build_689, dm_build_690, dm_build_691, dm_build_692)
}

func (Dm_build_694 *dm_build_622) Dm_build_693(dm_build_695 []byte, dm_build_696 int, dm_build_697 []byte, dm_build_698 int, dm_build_699 int) int {
	dm_build_696 += Dm_build_694.Dm_build_664(dm_build_695, dm_build_696, uint16(dm_build_699))
	return 2 + Dm_build_694.Dm_build_679(dm_build_695, dm_build_696, dm_build_697, dm_build_698, dm_build_699)
}

func (Dm_build_701 *dm_build_622) Dm_build_700(dm_build_702 []byte, dm_build_703 int, dm_build_704 string, dm_build_705 string, dm_build_706 *DmConnection) int {
	dm_build_707 := Dm_build_701.Dm_build_836(dm_build_704, dm_build_705, dm_build_706)
	dm_build_703 += Dm_build_701.Dm_build_669(dm_build_702, dm_build_703, uint32(len(dm_build_707)))
	return 4 + Dm_build_701.Dm_build_679(dm_build_702, dm_build_703, dm_build_707, 0, len(dm_build_707))
}

func (Dm_build_709 *dm_build_622) Dm_build_708(dm_build_710 []byte, dm_build_711 int, dm_build_712 string, dm_build_713 string, dm_build_714 *DmConnection) int {
	dm_build_715 := Dm_build_709.Dm_build_836(dm_build_712, dm_build_713, dm_build_714)

	dm_build_711 += Dm_build_709.Dm_build_664(dm_build_710, dm_build_711, uint16(len(dm_build_715)))
	return 2 + Dm_build_709.Dm_build_679(dm_build_710, dm_build_711, dm_build_715, 0, len(dm_build_715))
}

func (Dm_build_717 *dm_build_622) Dm_build_716(dm_build_718 []byte, dm_build_719 int) byte {
	return dm_build_718[dm_build_719]
}

func (Dm_build_721 *dm_build_622) Dm_build_720(dm_build_722 []byte, dm_build_723 int) int16 {
	var dm_build_724 int16
	dm_build_724 = int16(dm_build_722[dm_build_723] & 0xff)
	dm_build_723++
	dm_build_724 |= int16(dm_build_722[dm_build_723]&0xff) << 8
	return dm_build_724
}

func (Dm_build_726 *dm_build_622) Dm_build_725(dm_build_727 []byte, dm_build_728 int) int32 {
	var dm_build_729 int32
	dm_build_729 = int32(dm_build_727[dm_build_728] & 0xff)
	dm_build_728++
	dm_build_729 |= int32(dm_build_727[dm_build_728]&0xff) << 8
	dm_build_728++
	dm_build_729 |= int32(dm_build_727[dm_build_728]&0xff) << 16
	dm_build_728++
	dm_build_729 |= int32(dm_build_727[dm_build_728]&0xff) << 24
	return dm_build_729
}

func (Dm_build_731 *dm_build_622) Dm_build_730(dm_build_732 []byte, dm_build_733 int) int64 {
	var dm_build_734 int64
	dm_build_734 = int64(dm_build_732[dm_build_733] & 0xff)
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 8
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 16
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 24
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 32
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 40
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 48
	dm_build_733++
	dm_build_734 |= int64(dm_build_732[dm_build_733]&0xff) << 56
	return dm_build_734
}

func (Dm_build_736 *dm_build_622) Dm_build_735(dm_build_737 []byte, dm_build_738 int) float32 {
	return math.Float32frombits(Dm_build_736.Dm_build_752(dm_build_737, dm_build_738))
}

func (Dm_build_740 *dm_build_622) Dm_build_739(dm_build_741 []byte, dm_build_742 int) float64 {
	return math.Float64frombits(Dm_build_740.Dm_build_757(dm_build_741, dm_build_742))
}

func (Dm_build_744 *dm_build_622) Dm_build_743(dm_build_745 []byte, dm_build_746 int) uint8 {
	return uint8(dm_build_745[dm_build_746] & 0xff)
}

func (Dm_build_748 *dm_build_622) Dm_build_747(dm_build_749 []byte, dm_build_750 int) uint16 {
	var dm_build_751 uint16
	dm_build_751 = uint16(dm_build_749[dm_build_750] & 0xff)
	dm_build_750++
	dm_build_751 |= uint16(dm_build_749[dm_build_750]&0xff) << 8
	return dm_build_751
}

func (Dm_build_753 *dm_build_622) Dm_build_752(dm_build_754 []byte, dm_build_755 int) uint32 {
	var dm_build_756 uint32
	dm_build_756 = uint32(dm_build_754[dm_build_755] & 0xff)
	dm_build_755++
	dm_build_756 |= uint32(dm_build_754[dm_build_755]&0xff) << 8
	dm_build_755++
	dm_build_756 |= uint32(dm_build_754[dm_build_755]&0xff) << 16
	dm_build_755++
	dm_build_756 |= uint32(dm_build_754[dm_build_755]&0xff) << 24
	return dm_build_756
}

func (Dm_build_758 *dm_build_622) Dm_build_757(dm_build_759 []byte, dm_build_760 int) uint64 {
	var dm_build_761 uint64
	dm_build_761 = uint64(dm_build_759[dm_build_760] & 0xff)
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 8
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 16
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 24
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 32
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 40
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 48
	dm_build_760++
	dm_build_761 |= uint64(dm_build_759[dm_build_760]&0xff) << 56
	return dm_build_761
}

func (Dm_build_763 *dm_build_622) Dm_build_762(dm_build_764 []byte, dm_build_765 int) []byte {
	dm_build_766 := Dm_build_763.Dm_build_752(dm_build_764, dm_build_765)

	dm_build_767 := make([]byte, dm_build_766)
	copy(dm_build_767[:int(dm_build_766)], dm_build_764[dm_build_765+4:dm_build_765+4+int(dm_build_766)])
	return dm_build_767
}

func (Dm_build_769 *dm_build_622) Dm_build_768(dm_build_770 []byte, dm_build_771 int) []byte {
	dm_build_772 := Dm_build_769.Dm_build_747(dm_build_770, dm_build_771)

	dm_build_773 := make([]byte, dm_build_772)
	copy(dm_build_773[:int(dm_build_772)], dm_build_770[dm_build_771+2:dm_build_771+2+int(dm_build_772)])
	return dm_build_773
}

func (Dm_build_775 *dm_build_622) Dm_build_774(dm_build_776 []byte, dm_build_777 int, dm_build_778 int) []byte {

	dm_build_779 := make([]byte, dm_build_778)
	copy(dm_build_779[:dm_build_778], dm_build_776[dm_build_777:dm_build_777+dm_build_778])
	return dm_build_779
}

func (Dm_build_781 *dm_build_622) Dm_build_780(dm_build_782 []byte, dm_build_783 int, dm_build_784 int, dm_build_785 string, dm_build_786 *DmConnection) string {
	return Dm_build_781.Dm_build_873(dm_build_782[dm_build_783:dm_build_783+dm_build_784], dm_build_785, dm_build_786)
}

func (Dm_build_788 *dm_build_622) Dm_build_787(dm_build_789 []byte, dm_build_790 int, dm_build_791 string, dm_build_792 *DmConnection) string {
	dm_build_793 := Dm_build_788.Dm_build_752(dm_build_789, dm_build_790)
	dm_build_790 += 4
	return Dm_build_788.Dm_build_780(dm_build_789, dm_build_790, int(dm_build_793), dm_build_791, dm_build_792)
}

func (Dm_build_795 *dm_build_622) Dm_build_794(dm_build_796 []byte, dm_build_797 int, dm_build_798 string, dm_build_799 *DmConnection) string {
	dm_build_800 := Dm_build_795.Dm_build_747(dm_build_796, dm_build_797)
	dm_build_797 += 2
	return Dm_build_795.Dm_build_780(dm_build_796, dm_build_797, int(dm_build_800), dm_build_798, dm_build_799)
}

func (Dm_build_802 *dm_build_622) Dm_build_801(dm_build_803 byte) []byte {
	return []byte{dm_build_803}
}

func (Dm_build_805 *dm_build_622) Dm_build_804(dm_build_806 int16) []byte {
	return []byte{byte(dm_build_806), byte(dm_build_806 >> 8)}
}

func (Dm_build_808 *dm_build_622) Dm_build_807(dm_build_809 int32) []byte {
	return []byte{byte(dm_build_809), byte(dm_build_809 >> 8), byte(dm_build_809 >> 16), byte(dm_build_809 >> 24)}
}

func (Dm_build_811 *dm_build_622) Dm_build_810(dm_build_812 int64) []byte {
	return []byte{byte(dm_build_812), byte(dm_build_812 >> 8), byte(dm_build_812 >> 16), byte(dm_build_812 >> 24), byte(dm_build_812 >> 32),
		byte(dm_build_812 >> 40), byte(dm_build_812 >> 48), byte(dm_build_812 >> 56)}
}

func (Dm_build_814 *dm_build_622) Dm_build_813(dm_build_815 float32) []byte {
	return Dm_build_814.Dm_build_825(math.Float32bits(dm_build_815))
}

func (Dm_build_817 *dm_build_622) Dm_build_816(dm_build_818 float64) []byte {
	return Dm_build_817.Dm_build_828(math.Float64bits(dm_build_818))
}

func (Dm_build_820 *dm_build_622) Dm_build_819(dm_build_821 uint8) []byte {
	return []byte{byte(dm_build_821)}
}

func (Dm_build_823 *dm_build_622) Dm_build_822(dm_build_824 uint16) []byte {
	return []byte{byte(dm_build_824), byte(dm_build_824 >> 8)}
}

func (Dm_build_826 *dm_build_622) Dm_build_825(dm_build_827 uint32) []byte {
	return []byte{byte(dm_build_827), byte(dm_build_827 >> 8), byte(dm_build_827 >> 16), byte(dm_build_827 >> 24)}
}

func (Dm_build_829 *dm_build_622) Dm_build_828(dm_build_830 uint64) []byte {
	return []byte{byte(dm_build_830), byte(dm_build_830 >> 8), byte(dm_build_830 >> 16), byte(dm_build_830 >> 24), byte(dm_build_830 >> 32), byte(dm_build_830 >> 40), byte(dm_build_830 >> 48), byte(dm_build_830 >> 56)}
}

func (Dm_build_832 *dm_build_622) Dm_build_831(dm_build_833 []byte, dm_build_834 string, dm_build_835 *DmConnection) []byte {
	if dm_build_834 == "UTF-8" {
		return dm_build_833
	}

	if dm_build_835 == nil {
		if e := dm_build_878(dm_build_834); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_833), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_835.encodeBuffer == nil {
		dm_build_835.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_835.encode = dm_build_878(dm_build_835.getServerEncoding())
		dm_build_835.transformReaderDst = make([]byte, 4096)
		dm_build_835.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_835.encode; e != nil {

		dm_build_835.encodeBuffer.Reset()

		n, err := dm_build_835.encodeBuffer.ReadFrom(
			Dm_build_892(bytes.NewReader(dm_build_833), e.NewEncoder(), dm_build_835.transformReaderDst, dm_build_835.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_835.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_837 *dm_build_622) Dm_build_836(dm_build_838 string, dm_build_839 string, dm_build_840 *DmConnection) []byte {
	return Dm_build_837.Dm_build_831([]byte(dm_build_838), dm_build_839, dm_build_840)
}

func (Dm_build_842 *dm_build_622) Dm_build_841(dm_build_843 []byte) byte {
	return Dm_build_842.Dm_build_716(dm_build_843, 0)
}

func (Dm_build_845 *dm_build_622) Dm_build_844(dm_build_846 []byte) int16 {
	return Dm_build_845.Dm_build_720(dm_build_846, 0)
}

func (Dm_build_848 *dm_build_622) Dm_build_847(dm_build_849 []byte) int32 {
	return Dm_build_848.Dm_build_725(dm_build_849, 0)
}

func (Dm_build_851 *dm_build_622) Dm_build_850(dm_build_852 []byte) int64 {
	return Dm_build_851.Dm_build_730(dm_build_852, 0)
}

func (Dm_build_854 *dm_build_622) Dm_build_853(dm_build_855 []byte) float32 {
	return Dm_build_854.Dm_build_735(dm_build_855, 0)
}

func (Dm_build_857 *dm_build_622) Dm_build_856(dm_build_858 []byte) float64 {
	return Dm_build_857.Dm_build_739(dm_build_858, 0)
}

func (Dm_build_860 *dm_build_622) Dm_build_859(dm_build_861 []byte) uint8 {
	return Dm_build_860.Dm_build_743(dm_build_861, 0)
}

func (Dm_build_863 *dm_build_622) Dm_build_862(dm_build_864 []byte) uint16 {
	return Dm_build_863.Dm_build_747(dm_build_864, 0)
}

func (Dm_build_866 *dm_build_622) Dm_build_865(dm_build_867 []byte) uint32 {
	return Dm_build_866.Dm_build_752(dm_build_867, 0)
}

func (Dm_build_869 *dm_build_622) Dm_build_868(dm_build_870 []byte, dm_build_871 string, dm_build_872 *DmConnection) []byte {
	if dm_build_871 == "UTF-8" {
		return dm_build_870
	}

	if dm_build_872 == nil {
		if e := dm_build_878(dm_build_871); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_870), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_872.encodeBuffer == nil {
		dm_build_872.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_872.encode = dm_build_878(dm_build_872.getServerEncoding())
		dm_build_872.transformReaderDst = make([]byte, 4096)
		dm_build_872.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_872.encode; e != nil {

		dm_build_872.encodeBuffer.Reset()

		n, err := dm_build_872.encodeBuffer.ReadFrom(
			Dm_build_892(bytes.NewReader(dm_build_870), e.NewDecoder(), dm_build_872.transformReaderDst, dm_build_872.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_872.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_874 *dm_build_622) Dm_build_873(dm_build_875 []byte, dm_build_876 string, dm_build_877 *DmConnection) string {
	return string(Dm_build_874.Dm_build_868(dm_build_875, dm_build_876, dm_build_877))
}

func dm_build_878(dm_build_879 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_879); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_880 struct {
	dm_build_881 io.Reader
	dm_build_882 transform.Transformer
	dm_build_883 error

	dm_build_884               []byte
	dm_build_885, dm_build_886 int

	dm_build_887               []byte
	dm_build_888, dm_build_889 int

	dm_build_890 bool
}

const dm_build_891 = 4096

func Dm_build_892(dm_build_893 io.Reader, dm_build_894 transform.Transformer, dm_build_895 []byte, dm_build_896 []byte) *Dm_build_880 {
	dm_build_894.Reset()
	return &Dm_build_880{
		dm_build_881: dm_build_893,
		dm_build_882: dm_build_894,
		dm_build_884: dm_build_895,
		dm_build_887: dm_build_896,
	}
}

func (dm_build_898 *Dm_build_880) Read(dm_build_899 []byte) (int, error) {
	dm_build_900, dm_build_901 := 0, error(nil)
	for {

		if dm_build_898.dm_build_885 != dm_build_898.dm_build_886 {
			dm_build_900 = copy(dm_build_899, dm_build_898.dm_build_884[dm_build_898.dm_build_885:dm_build_898.dm_build_886])
			dm_build_898.dm_build_885 += dm_build_900
			if dm_build_898.dm_build_885 == dm_build_898.dm_build_886 && dm_build_898.dm_build_890 {
				return dm_build_900, dm_build_898.dm_build_883
			}
			return dm_build_900, nil
		} else if dm_build_898.dm_build_890 {
			return 0, dm_build_898.dm_build_883
		}

		if dm_build_898.dm_build_888 != dm_build_898.dm_build_889 || dm_build_898.dm_build_883 != nil {
			dm_build_898.dm_build_885 = 0
			dm_build_898.dm_build_886, dm_build_900, dm_build_901 = dm_build_898.dm_build_882.Transform(dm_build_898.dm_build_884, dm_build_898.dm_build_887[dm_build_898.dm_build_888:dm_build_898.dm_build_889], dm_build_898.dm_build_883 == io.EOF)
			dm_build_898.dm_build_888 += dm_build_900

			switch {
			case dm_build_901 == nil:
				if dm_build_898.dm_build_888 != dm_build_898.dm_build_889 {
					dm_build_898.dm_build_883 = nil
				}

				dm_build_898.dm_build_890 = dm_build_898.dm_build_883 != nil
				continue
			case dm_build_901 == transform.ErrShortDst && (dm_build_898.dm_build_886 != 0 || dm_build_900 != 0):

				continue
			case dm_build_901 == transform.ErrShortSrc && dm_build_898.dm_build_889-dm_build_898.dm_build_888 != len(dm_build_898.dm_build_887) && dm_build_898.dm_build_883 == nil:

			default:
				dm_build_898.dm_build_890 = true

				if dm_build_898.dm_build_883 == nil || dm_build_898.dm_build_883 == io.EOF {
					dm_build_898.dm_build_883 = dm_build_901
				}
				continue
			}
		}

		if dm_build_898.dm_build_888 != 0 {
			dm_build_898.dm_build_888, dm_build_898.dm_build_889 = 0, copy(dm_build_898.dm_build_887, dm_build_898.dm_build_887[dm_build_898.dm_build_888:dm_build_898.dm_build_889])
		}
		dm_build_900, dm_build_898.dm_build_883 = dm_build_898.dm_build_881.Read(dm_build_898.dm_build_887[dm_build_898.dm_build_889:])
		dm_build_898.dm_build_889 += dm_build_900
	}
}
