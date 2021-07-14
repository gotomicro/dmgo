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
	Dm_build_688 = 8192
	Dm_build_689 = 2 * time.Second
)

type dm_build_690 struct {
	dm_build_691 *net.TCPConn
	dm_build_692 *tls.Conn
	dm_build_693 *Dm_build_358
	dm_build_694 *DmConnection
	dm_build_695 security.Cipher
	dm_build_696 bool
	dm_build_697 bool
	dm_build_698 *security.DhKey

	dm_build_699 bool
	dm_build_700 string
	dm_build_701 bool
}

func dm_build_702(dm_build_703 *DmConnection) (*dm_build_690, error) {
	dm_build_704, dm_build_705 := dm_build_707(dm_build_703.dmConnector.host+":"+strconv.Itoa(int(dm_build_703.dmConnector.port)), time.Duration(dm_build_703.dmConnector.socketTimeout)*time.Second)
	if dm_build_705 != nil {
		return nil, dm_build_705
	}

	dm_build_706 := dm_build_690{}
	dm_build_706.dm_build_691 = dm_build_704
	dm_build_706.dm_build_693 = Dm_build_361(Dm_build_961)
	dm_build_706.dm_build_694 = dm_build_703
	dm_build_706.dm_build_696 = false
	dm_build_706.dm_build_697 = false
	dm_build_706.dm_build_699 = false
	dm_build_706.dm_build_700 = ""
	dm_build_706.dm_build_701 = false
	dm_build_703.Access = &dm_build_706

	return &dm_build_706, nil
}

func dm_build_707(dm_build_708 string, dm_build_709 time.Duration) (*net.TCPConn, error) {
	dm_build_710, dm_build_711 := net.DialTimeout("tcp", dm_build_708, dm_build_709)
	if dm_build_711 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_708).throw()
	}

	if tcpConn, ok := dm_build_710.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_689)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_713 *dm_build_690) dm_build_712(dm_build_714 dm_build_1081) bool {
	var dm_build_715 = dm_build_713.dm_build_694.dmConnector.compress
	if dm_build_714.dm_build_1096() == Dm_build_988 || dm_build_715 == Dm_build_1037 {
		return false
	}

	if dm_build_715 == Dm_build_1035 {
		return true
	} else if dm_build_715 == Dm_build_1036 {
		return !dm_build_713.dm_build_694.Local && dm_build_714.dm_build_1094() > Dm_build_1034
	}

	return false
}

func (dm_build_717 *dm_build_690) dm_build_716(dm_build_718 dm_build_1081) bool {
	var dm_build_719 = dm_build_717.dm_build_694.dmConnector.compress
	if dm_build_718.dm_build_1096() == Dm_build_988 || dm_build_719 == Dm_build_1037 {
		return false
	}

	if dm_build_719 == Dm_build_1035 {
		return true
	} else if dm_build_719 == Dm_build_1036 {
		return dm_build_717.dm_build_693.Dm_build_621(Dm_build_996) == 1
	}

	return false
}

func (dm_build_721 *dm_build_690) dm_build_720(dm_build_722 dm_build_1081) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_724 := dm_build_722.dm_build_1094()

	if dm_build_724 > 0 {

		if dm_build_721.dm_build_712(dm_build_722) {
			var retBytes, err = Compress(dm_build_721.dm_build_693, Dm_build_989, int(dm_build_724), int(dm_build_721.dm_build_694.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_721.dm_build_693.Dm_build_372(Dm_build_989)

			dm_build_721.dm_build_693.Dm_build_409(dm_build_724)

			dm_build_721.dm_build_693.Dm_build_437(retBytes)

			dm_build_722.dm_build_1095(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_721.dm_build_693.Dm_build_541(Dm_build_996, 1)
		}

		if dm_build_721.dm_build_697 {
			dm_build_724 = dm_build_722.dm_build_1094()
			var retBytes = dm_build_721.dm_build_695.Encrypt(dm_build_721.dm_build_693.Dm_build_648(Dm_build_989, int(dm_build_724)), true)

			dm_build_721.dm_build_693.Dm_build_372(Dm_build_989)

			dm_build_721.dm_build_693.Dm_build_437(retBytes)

			dm_build_722.dm_build_1095(int32(len(retBytes)))
		}
	}

	dm_build_722.dm_build_1090()
	if dm_build_721.dm_build_952(dm_build_722) {
		if dm_build_721.dm_build_692 != nil {
			dm_build_721.dm_build_693.Dm_build_375(0)
			dm_build_721.dm_build_693.Dm_build_394(dm_build_721.dm_build_692)
		}
	} else {
		dm_build_721.dm_build_693.Dm_build_375(0)
		dm_build_721.dm_build_693.Dm_build_394(dm_build_721.dm_build_691)
	}
	return nil
}

func (dm_build_726 *dm_build_690) dm_build_725(dm_build_727 dm_build_1081) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_729 := int32(0)
	if dm_build_726.dm_build_952(dm_build_727) {
		if dm_build_726.dm_build_692 != nil {
			dm_build_726.dm_build_693.Dm_build_372(0)
			dm_build_726.dm_build_693.Dm_build_388(dm_build_726.dm_build_692, Dm_build_989)
			dm_build_729 = dm_build_727.dm_build_1094()
			if dm_build_729 > 0 {
				dm_build_726.dm_build_693.Dm_build_388(dm_build_726.dm_build_692, int(dm_build_729))
			}
		}
	} else {

		dm_build_726.dm_build_693.Dm_build_372(0)
		dm_build_726.dm_build_693.Dm_build_388(dm_build_726.dm_build_691, Dm_build_989)
		dm_build_729 = dm_build_727.dm_build_1094()

		if dm_build_729 > 0 {
			dm_build_726.dm_build_693.Dm_build_388(dm_build_726.dm_build_691, int(dm_build_729))
		}
	}

	dm_build_727.dm_build_1091()

	dm_build_729 = dm_build_727.dm_build_1094()
	if dm_build_729 <= 0 {
		return nil
	}

	if dm_build_726.dm_build_697 {
		ebytes := dm_build_726.dm_build_693.Dm_build_648(Dm_build_989, int(dm_build_729))
		bytes, err := dm_build_726.dm_build_695.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_726.dm_build_693.Dm_build_372(Dm_build_989)
		dm_build_726.dm_build_693.Dm_build_437(bytes)
		dm_build_727.dm_build_1095(int32(len(bytes)))
	}

	if dm_build_726.dm_build_716(dm_build_727) {

		dm_build_729 = dm_build_727.dm_build_1094()
		cbytes := dm_build_726.dm_build_693.Dm_build_648(Dm_build_989+ULINT_SIZE, int(dm_build_729-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_726.dm_build_694.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_726.dm_build_693.Dm_build_372(Dm_build_989)
		dm_build_726.dm_build_693.Dm_build_437(bytes)
		dm_build_727.dm_build_1095(int32(len(bytes)))
	}
	return nil
}

func (dm_build_731 *dm_build_690) dm_build_730(dm_build_732 dm_build_1081) (dm_build_733 interface{}, dm_build_734 error) {
	dm_build_734 = dm_build_732.dm_build_1085(dm_build_732)
	if dm_build_734 != nil {
		return nil, dm_build_734
	}

	dm_build_734 = dm_build_731.dm_build_720(dm_build_732)
	if dm_build_734 != nil {
		return nil, dm_build_734
	}

	dm_build_734 = dm_build_731.dm_build_725(dm_build_732)
	if dm_build_734 != nil {
		return nil, dm_build_734
	}

	return dm_build_732.dm_build_1089(dm_build_732)
}

func (dm_build_736 *dm_build_690) dm_build_735() (*dm_build_1518, error) {

	Dm_build_737 := dm_build_1524(dm_build_736)
	_, dm_build_738 := dm_build_736.dm_build_730(Dm_build_737)
	if dm_build_738 != nil {
		return nil, dm_build_738
	}

	return Dm_build_737, nil
}

func (dm_build_740 *dm_build_690) dm_build_739() error {

	dm_build_741 := dm_build_1386(dm_build_740)
	_, dm_build_742 := dm_build_740.dm_build_730(dm_build_741)
	if dm_build_742 != nil {
		return dm_build_742
	}

	return nil
}

func (dm_build_744 *dm_build_690) dm_build_743() error {

	var dm_build_745 *dm_build_1518
	var err error
	if dm_build_745, err = dm_build_744.dm_build_735(); err != nil {
		return err
	}

	if dm_build_744.dm_build_694.sslEncrypt == 2 {
		if err = dm_build_744.dm_build_948(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_744.dm_build_694.sslEncrypt == 1 {
		if err = dm_build_744.dm_build_948(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_744.dm_build_697 || dm_build_744.dm_build_696 {
		k, err := dm_build_744.dm_build_938()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_745.Dm_build_1522)
		encryptType := dm_build_745.dm_build_1520
		hashType := int(dm_build_745.Dm_build_1521)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_744.dm_build_941(encryptType, sessionKey, dm_build_744.dm_build_694.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_744.dm_build_739(); err != nil {
		return err
	}
	return nil
}

func (dm_build_748 *dm_build_690) Dm_build_747(dm_build_749 *DmStatement) error {
	dm_build_750 := dm_build_1547(dm_build_748, dm_build_749)
	_, dm_build_751 := dm_build_748.dm_build_730(dm_build_750)
	if dm_build_751 != nil {
		return dm_build_751
	}

	return nil
}

func (dm_build_753 *dm_build_690) Dm_build_752(dm_build_754 int32) error {
	dm_build_755 := dm_build_1557(dm_build_753, dm_build_754)
	_, dm_build_756 := dm_build_753.dm_build_730(dm_build_755)
	if dm_build_756 != nil {
		return dm_build_756
	}

	return nil
}

func (dm_build_758 *dm_build_690) Dm_build_757(dm_build_759 *DmStatement, dm_build_760 bool, dm_build_761 int16) (*execRetInfo, error) {
	dm_build_762 := dm_build_1424(dm_build_758, dm_build_759, dm_build_760, dm_build_761)
	dm_build_763, dm_build_764 := dm_build_758.dm_build_730(dm_build_762)
	if dm_build_764 != nil {
		return nil, dm_build_764
	}
	return dm_build_763.(*execRetInfo), nil
}

func (dm_build_766 *dm_build_690) Dm_build_765(dm_build_767 *DmStatement, dm_build_768 int16) (*execRetInfo, error) {
	return dm_build_766.Dm_build_757(dm_build_767, false, Dm_build_1041)
}

func (dm_build_770 *dm_build_690) Dm_build_769(dm_build_771 *DmStatement, dm_build_772 []OptParameter) (*execRetInfo, error) {
	dm_build_773, dm_build_774 := dm_build_770.dm_build_730(dm_build_1183(dm_build_770, dm_build_771, dm_build_772))
	if dm_build_774 != nil {
		return nil, dm_build_774
	}

	return dm_build_773.(*execRetInfo), nil
}

func (dm_build_776 *dm_build_690) Dm_build_775(dm_build_777 *DmStatement, dm_build_778 int16) (*execRetInfo, error) {
	return dm_build_776.Dm_build_757(dm_build_777, true, dm_build_778)
}

func (dm_build_780 *dm_build_690) Dm_build_779(dm_build_781 *DmStatement, dm_build_782 [][]interface{}) (*execRetInfo, error) {
	dm_build_783 := dm_build_1206(dm_build_780, dm_build_781, dm_build_782)
	dm_build_784, dm_build_785 := dm_build_780.dm_build_730(dm_build_783)
	if dm_build_785 != nil {
		return nil, dm_build_785
	}
	return dm_build_784.(*execRetInfo), nil
}

func (dm_build_787 *dm_build_690) Dm_build_786(dm_build_788 *DmStatement, dm_build_789 [][]interface{}, dm_build_790 bool) (*execRetInfo, error) {
	var dm_build_791, dm_build_792 = 0, 0
	var dm_build_793 = len(dm_build_789)
	var dm_build_794 [][]interface{}
	var dm_build_795 = NewExceInfo()
	dm_build_795.updateCounts = make([]int64, dm_build_793)
	var dm_build_796 = false
	for dm_build_791 < dm_build_793 {
		for dm_build_792 = dm_build_791; dm_build_792 < dm_build_793; dm_build_792++ {
			paramData := dm_build_789[dm_build_792]
			bindData := make([]interface{}, dm_build_788.paramCount)
			dm_build_796 = false
			for icol := 0; icol < int(dm_build_788.paramCount); icol++ {
				if dm_build_788.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_787.dm_build_921(bindData, paramData, icol) {
					dm_build_796 = true
					break
				}
			}

			if dm_build_796 {
				break
			}
			dm_build_794 = append(dm_build_794, bindData)
		}

		if dm_build_792 != dm_build_791 {
			tmpExecInfo, err := dm_build_787.Dm_build_779(dm_build_788, dm_build_794)
			if err != nil {
				return nil, err
			}
			dm_build_794 = dm_build_794[0:0]
			dm_build_795.union(tmpExecInfo, dm_build_791, dm_build_792-dm_build_791)
		}

		if dm_build_792 < dm_build_793 {
			tmpExecInfo, err := dm_build_787.Dm_build_797(dm_build_788, dm_build_789[dm_build_792], dm_build_790)
			if err != nil {
				return nil, err
			}

			dm_build_790 = true
			dm_build_795.union(tmpExecInfo, dm_build_792, 1)
		}

		dm_build_791 = dm_build_792 + 1
	}
	for _, i := range dm_build_795.updateCounts {
		if i > 0 {
			dm_build_795.updateCount += i
		}
	}
	return dm_build_795, nil
}

func (dm_build_798 *dm_build_690) Dm_build_797(dm_build_799 *DmStatement, dm_build_800 []interface{}, dm_build_801 bool) (*execRetInfo, error) {

	var dm_build_802 = make([]interface{}, dm_build_799.paramCount)
	for icol := 0; icol < int(dm_build_799.paramCount); icol++ {
		if dm_build_799.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_798.dm_build_921(dm_build_802, dm_build_800, icol) {

			if !dm_build_801 {
				preExecute := dm_build_1414(dm_build_798, dm_build_799, dm_build_799.params)
				dm_build_798.dm_build_730(preExecute)
				dm_build_801 = true
			}

			dm_build_798.dm_build_927(dm_build_799, dm_build_799.params[icol], icol, dm_build_800[icol].(iOffRowBinder))
			dm_build_802[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_803 = make([][]interface{}, 1, 1)
	dm_build_803[0] = dm_build_802

	dm_build_804 := dm_build_1206(dm_build_798, dm_build_799, dm_build_803)
	dm_build_805, dm_build_806 := dm_build_798.dm_build_730(dm_build_804)
	if dm_build_806 != nil {
		return nil, dm_build_806
	}
	return dm_build_805.(*execRetInfo), nil
}

func (dm_build_808 *dm_build_690) Dm_build_807(dm_build_809 *DmStatement, dm_build_810 int16) (*execRetInfo, error) {
	dm_build_811 := dm_build_1401(dm_build_808, dm_build_809, dm_build_810)

	dm_build_812, dm_build_813 := dm_build_808.dm_build_730(dm_build_811)
	if dm_build_813 != nil {
		return nil, dm_build_813
	}
	return dm_build_812.(*execRetInfo), nil
}

func (dm_build_815 *dm_build_690) Dm_build_814(dm_build_816 *innerRows, dm_build_817 int64) (*execRetInfo, error) {
	dm_build_818 := dm_build_1306(dm_build_815, dm_build_816, dm_build_817, INT64_MAX)
	dm_build_819, dm_build_820 := dm_build_815.dm_build_730(dm_build_818)
	if dm_build_820 != nil {
		return nil, dm_build_820
	}
	return dm_build_819.(*execRetInfo), nil
}

func (dm_build_822 *dm_build_690) Commit() error {
	dm_build_823 := dm_build_1169(dm_build_822)
	_, dm_build_824 := dm_build_822.dm_build_730(dm_build_823)
	if dm_build_824 != nil {
		return dm_build_824
	}

	return nil
}

func (dm_build_826 *dm_build_690) Rollback() error {
	dm_build_827 := dm_build_1462(dm_build_826)
	_, dm_build_828 := dm_build_826.dm_build_730(dm_build_827)
	if dm_build_828 != nil {
		return dm_build_828
	}

	return nil
}

func (dm_build_830 *dm_build_690) Dm_build_829(dm_build_831 *DmConnection) error {
	dm_build_832 := dm_build_1467(dm_build_830, dm_build_831.IsoLevel)
	_, dm_build_833 := dm_build_830.dm_build_730(dm_build_832)
	if dm_build_833 != nil {
		return dm_build_833
	}

	return nil
}

func (dm_build_835 *dm_build_690) Dm_build_834(dm_build_836 *DmStatement, dm_build_837 string) error {
	dm_build_838 := dm_build_1174(dm_build_835, dm_build_836, dm_build_837)
	_, dm_build_839 := dm_build_835.dm_build_730(dm_build_838)
	if dm_build_839 != nil {
		return dm_build_839
	}

	return nil
}

func (dm_build_841 *dm_build_690) Dm_build_840(dm_build_842 []uint32) ([]int64, error) {
	dm_build_843 := dm_build_1565(dm_build_841, dm_build_842)
	dm_build_844, dm_build_845 := dm_build_841.dm_build_730(dm_build_843)
	if dm_build_845 != nil {
		return nil, dm_build_845
	}
	return dm_build_844.([]int64), nil
}

func (dm_build_847 *dm_build_690) Close() error {
	if dm_build_847.dm_build_701 {
		return nil
	}

	dm_build_848 := dm_build_847.dm_build_691.Close()
	if dm_build_848 != nil {
		return dm_build_848
	}

	dm_build_847.dm_build_694 = nil
	dm_build_847.dm_build_701 = true
	return nil
}

func (dm_build_850 *dm_build_690) dm_build_849(dm_build_851 *lob) (int64, error) {
	dm_build_852 := dm_build_1337(dm_build_850, dm_build_851)
	dm_build_853, dm_build_854 := dm_build_850.dm_build_730(dm_build_852)
	if dm_build_854 != nil {
		return 0, dm_build_854
	}
	return dm_build_853.(int64), nil
}

func (dm_build_856 *dm_build_690) dm_build_855(dm_build_857 *lob, dm_build_858 int32, dm_build_859 int32) ([]byte, error) {
	dm_build_860 := dm_build_1324(dm_build_856, dm_build_857, int(dm_build_858), int(dm_build_859))
	dm_build_861, dm_build_862 := dm_build_856.dm_build_730(dm_build_860)
	if dm_build_862 != nil {
		return nil, dm_build_862
	}
	return dm_build_861.([]byte), nil
}

func (dm_build_864 *dm_build_690) dm_build_863(dm_build_865 *DmBlob, dm_build_866 int32, dm_build_867 int32) ([]byte, error) {
	var dm_build_868 = make([]byte, dm_build_867)
	var dm_build_869 int32 = 0
	var dm_build_870 int32 = 0
	var dm_build_871 []byte
	var dm_build_872 error
	for dm_build_869 < dm_build_867 {
		dm_build_870 = dm_build_867 - dm_build_869
		if dm_build_870 > Dm_build_1074 {
			dm_build_870 = Dm_build_1074
		}
		dm_build_871, dm_build_872 = dm_build_864.dm_build_855(&dm_build_865.lob, dm_build_866, dm_build_870)
		if dm_build_872 != nil {
			return nil, dm_build_872
		}
		if dm_build_871 == nil || len(dm_build_871) == 0 {
			break
		}
		Dm_build_1.Dm_build_57(dm_build_868, int(dm_build_869), dm_build_871, 0, len(dm_build_871))
		dm_build_869 += int32(len(dm_build_871))
		dm_build_866 += int32(len(dm_build_871))
		if dm_build_865.readOver {
			break
		}
	}
	return dm_build_868, nil
}

func (dm_build_874 *dm_build_690) dm_build_873(dm_build_875 *DmClob, dm_build_876 int32, dm_build_877 int32) (string, error) {
	var dm_build_878 bytes.Buffer
	var dm_build_879 int32 = 0
	var dm_build_880 int32 = 0
	var dm_build_881 []byte
	var dm_build_882 string
	var dm_build_883 error
	for dm_build_879 < dm_build_877 {
		dm_build_880 = dm_build_877 - dm_build_879
		if dm_build_880 > Dm_build_1074/2 {
			dm_build_880 = Dm_build_1074 / 2
		}
		dm_build_881, dm_build_883 = dm_build_874.dm_build_855(&dm_build_875.lob, dm_build_876, dm_build_880)
		if dm_build_883 != nil {
			return "", dm_build_883
		}
		if dm_build_881 == nil || len(dm_build_881) == 0 {
			break
		}
		dm_build_882 = Dm_build_1.Dm_build_158(dm_build_881, 0, len(dm_build_881), dm_build_875.serverEncoding, dm_build_874.dm_build_694)

		dm_build_878.WriteString(dm_build_882)
		dm_build_879 += int32(len(dm_build_882))
		dm_build_876 += int32(len(dm_build_882))
		if dm_build_875.readOver {
			break
		}
	}
	return dm_build_878.String(), nil
}

func (dm_build_885 *dm_build_690) dm_build_884(dm_build_886 *DmClob, dm_build_887 int, dm_build_888 string, dm_build_889 string) (int, error) {
	var dm_build_890 = Dm_build_1.Dm_build_214(dm_build_888, dm_build_889, dm_build_885.dm_build_694)
	var dm_build_891 = 0
	var dm_build_892 = len(dm_build_890)
	var dm_build_893 = 0
	var dm_build_894 = 0
	var dm_build_895 = 0
	var dm_build_896 = dm_build_892/Dm_build_1073 + 1
	var dm_build_897 byte = 0
	var dm_build_898 byte = 0x01
	var dm_build_899 byte = 0x02
	for i := 0; i < dm_build_896; i++ {
		dm_build_897 = 0
		if i == 0 {
			dm_build_897 |= dm_build_898
		}
		if i == dm_build_896-1 {
			dm_build_897 |= dm_build_899
		}
		dm_build_895 = dm_build_892 - dm_build_894
		if dm_build_895 > Dm_build_1073 {
			dm_build_895 = Dm_build_1073
		}

		setLobData := dm_build_1481(dm_build_885, &dm_build_886.lob, dm_build_897, dm_build_887, dm_build_890, dm_build_891, dm_build_895)
		ret, err := dm_build_885.dm_build_730(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_893, nil
		} else {
			dm_build_887 += int(tmp)
			dm_build_893 += int(tmp)
			dm_build_894 += dm_build_895
			dm_build_891 += dm_build_895
		}
	}
	return dm_build_893, nil
}

func (dm_build_901 *dm_build_690) dm_build_900(dm_build_902 *DmBlob, dm_build_903 int, dm_build_904 []byte) (int, error) {
	var dm_build_905 = 0
	var dm_build_906 = len(dm_build_904)
	var dm_build_907 = 0
	var dm_build_908 = 0
	var dm_build_909 = 0
	var dm_build_910 = dm_build_906/Dm_build_1073 + 1
	var dm_build_911 byte = 0
	var dm_build_912 byte = 0x01
	var dm_build_913 byte = 0x02
	for i := 0; i < dm_build_910; i++ {
		dm_build_911 = 0
		if i == 0 {
			dm_build_911 |= dm_build_912
		}
		if i == dm_build_910-1 {
			dm_build_911 |= dm_build_913
		}
		dm_build_909 = dm_build_906 - dm_build_908
		if dm_build_909 > Dm_build_1073 {
			dm_build_909 = Dm_build_1073
		}

		setLobData := dm_build_1481(dm_build_901, &dm_build_902.lob, dm_build_911, dm_build_903, dm_build_904, dm_build_905, dm_build_909)
		ret, err := dm_build_901.dm_build_730(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_907, nil
		} else {
			dm_build_903 += int(tmp)
			dm_build_907 += int(tmp)
			dm_build_908 += dm_build_909
			dm_build_905 += dm_build_909
		}
	}
	return dm_build_907, nil
}

func (dm_build_915 *dm_build_690) dm_build_914(dm_build_916 *lob, dm_build_917 int) (int64, error) {
	dm_build_918 := dm_build_1348(dm_build_915, dm_build_916, dm_build_917)
	dm_build_919, dm_build_920 := dm_build_915.dm_build_730(dm_build_918)
	if dm_build_920 != nil {
		return dm_build_916.length, dm_build_920
	}
	return dm_build_919.(int64), nil
}

func (dm_build_922 *dm_build_690) dm_build_921(dm_build_923 []interface{}, dm_build_924 []interface{}, dm_build_925 int) bool {
	var dm_build_926 = false
	if dm_build_925 >= len(dm_build_924) || dm_build_924[dm_build_925] == nil {
		dm_build_923[dm_build_925] = ParamDataEnum_Null
	} else if binder, ok := dm_build_924[dm_build_925].(iOffRowBinder); ok {
		dm_build_926 = true
		dm_build_923[dm_build_925] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_922.dm_build_694) {
			dm_build_923[dm_build_925] = &lobCtl{lob.buildCtlData()}
			dm_build_926 = false
		}
	} else {
		dm_build_923[dm_build_925] = dm_build_924[dm_build_925]
	}
	return dm_build_926
}

func (dm_build_928 *dm_build_690) dm_build_927(dm_build_929 *DmStatement, dm_build_930 parameter, dm_build_931 int, dm_build_932 iOffRowBinder) error {
	var dm_build_933 = Dm_build_284()
	dm_build_932.read(dm_build_933)
	var dm_build_934 = 0
	for !dm_build_932.isReadOver() || dm_build_933.Dm_build_285() > 0 {
		if !dm_build_932.isReadOver() && dm_build_933.Dm_build_285() < Dm_build_1073 {
			dm_build_932.read(dm_build_933)
		}
		if dm_build_933.Dm_build_285() > Dm_build_1073 {
			dm_build_934 = Dm_build_1073
		} else {
			dm_build_934 = dm_build_933.Dm_build_285()
		}

		putData := dm_build_1452(dm_build_928, dm_build_929, int16(dm_build_931), dm_build_933, int32(dm_build_934))
		_, err := dm_build_928.dm_build_730(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_936 *dm_build_690) dm_build_935() ([]byte, error) {
	var dm_build_937 error
	if dm_build_936.dm_build_698 == nil {
		if dm_build_936.dm_build_698, dm_build_937 = security.NewClientKeyPair(); dm_build_937 != nil {
			return nil, dm_build_937
		}
	}
	return security.Bn2Bytes(dm_build_936.dm_build_698.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_939 *dm_build_690) dm_build_938() (*security.DhKey, error) {
	var dm_build_940 error
	if dm_build_939.dm_build_698 == nil {
		if dm_build_939.dm_build_698, dm_build_940 = security.NewClientKeyPair(); dm_build_940 != nil {
			return nil, dm_build_940
		}
	}
	return dm_build_939.dm_build_698, nil
}

func (dm_build_942 *dm_build_690) dm_build_941(dm_build_943 int, dm_build_944 []byte, dm_build_945 string, dm_build_946 int) (dm_build_947 error) {
	if dm_build_943 > 0 && dm_build_943 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_944 != nil {
		dm_build_942.dm_build_695, dm_build_947 = security.NewSymmCipher(dm_build_943, dm_build_944)
	} else if dm_build_943 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_942.dm_build_695, dm_build_947 = security.NewThirdPartCipher(dm_build_943, dm_build_944, dm_build_945, dm_build_946); dm_build_947 != nil {
			dm_build_947 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_947.Error()).throw()
		}
	}
	return
}

func (dm_build_949 *dm_build_690) dm_build_948(dm_build_950 bool) (dm_build_951 error) {
	if dm_build_949.dm_build_692, dm_build_951 = security.NewTLSFromTCP(dm_build_949.dm_build_691, dm_build_949.dm_build_694.dmConnector.sslCertPath, dm_build_949.dm_build_694.dmConnector.sslKeyPath, dm_build_949.dm_build_694.dmConnector.user); dm_build_951 != nil {
		return
	}
	if !dm_build_950 {
		dm_build_949.dm_build_692 = nil
	}
	return
}

func (dm_build_953 *dm_build_690) dm_build_952(dm_build_954 dm_build_1081) bool {
	return dm_build_954.dm_build_1096() != Dm_build_988 && dm_build_953.dm_build_694.sslEncrypt == 1
}
