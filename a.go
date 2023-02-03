/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"gitee.com/chunanyong/dm/security"
)

const (
	Dm_build_695 = 8192
	Dm_build_696 = 2 * time.Second
)

type dm_build_697 struct {
	dm_build_698 *net.TCPConn
	dm_build_699 *tls.Conn
	dm_build_700 *Dm_build_361
	dm_build_701 *DmConnection
	dm_build_702 security.Cipher
	dm_build_703 bool
	dm_build_704 bool
	dm_build_705 *security.DhKey

	dm_build_706 bool
	dm_build_707 string
	dm_build_708 bool
}

func dm_build_709(dm_build_710 *DmConnection) (*dm_build_697, error) {
	dm_build_711, dm_build_712 := dm_build_714(dm_build_710.dmConnector.host+":"+strconv.Itoa(int(dm_build_710.dmConnector.port)), time.Duration(dm_build_710.dmConnector.socketTimeout)*time.Second)
	if dm_build_712 != nil {
		return nil, dm_build_712
	}

	dm_build_713 := dm_build_697{}
	dm_build_713.dm_build_698 = dm_build_711
	dm_build_713.dm_build_700 = Dm_build_364(Dm_build_976)
	dm_build_713.dm_build_701 = dm_build_710
	dm_build_713.dm_build_703 = false
	dm_build_713.dm_build_704 = false
	dm_build_713.dm_build_706 = false
	dm_build_713.dm_build_707 = ""
	dm_build_713.dm_build_708 = false
	dm_build_710.Access = &dm_build_713

	return &dm_build_713, nil
}

func dm_build_714(dm_build_715 string, dm_build_716 time.Duration) (*net.TCPConn, error) {
	dm_build_717, dm_build_718 := net.DialTimeout("tcp", dm_build_715, dm_build_716)
	if dm_build_718 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_715).throw()
	}

	if tcpConn, ok := dm_build_717.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_696)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_720 *dm_build_697) dm_build_719(dm_build_721 dm_build_1097) bool {
	var dm_build_722 = dm_build_720.dm_build_701.dmConnector.compress
	if dm_build_721.dm_build_1112() == Dm_build_1004 || dm_build_722 == Dm_build_1053 {
		return false
	}

	if dm_build_722 == Dm_build_1051 {
		return true
	} else if dm_build_722 == Dm_build_1052 {
		return !dm_build_720.dm_build_701.Local && dm_build_721.dm_build_1110() > Dm_build_1050
	}

	return false
}

func (dm_build_724 *dm_build_697) dm_build_723(dm_build_725 dm_build_1097) bool {
	var dm_build_726 = dm_build_724.dm_build_701.dmConnector.compress
	if dm_build_725.dm_build_1112() == Dm_build_1004 || dm_build_726 == Dm_build_1053 {
		return false
	}

	if dm_build_726 == Dm_build_1051 {
		return true
	} else if dm_build_726 == Dm_build_1052 {
		return dm_build_724.dm_build_700.Dm_build_628(Dm_build_1012) == 1
	}

	return false
}

func (dm_build_728 *dm_build_697) dm_build_727(dm_build_729 dm_build_1097) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_731 := dm_build_729.dm_build_1110()

	if dm_build_731 > 0 {

		if dm_build_728.dm_build_719(dm_build_729) {
			var retBytes, err = Compress(dm_build_728.dm_build_700, Dm_build_1005, int(dm_build_731), int(dm_build_728.dm_build_701.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_728.dm_build_700.Dm_build_375(Dm_build_1005)

			dm_build_728.dm_build_700.Dm_build_416(dm_build_731)

			dm_build_728.dm_build_700.Dm_build_444(retBytes)

			dm_build_729.dm_build_1111(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_728.dm_build_700.Dm_build_548(Dm_build_1012, 1)
		}

		if dm_build_728.dm_build_704 {
			dm_build_731 = dm_build_729.dm_build_1110()
			var retBytes = dm_build_728.dm_build_702.Encrypt(dm_build_728.dm_build_700.Dm_build_655(Dm_build_1005, int(dm_build_731)), true)

			dm_build_728.dm_build_700.Dm_build_375(Dm_build_1005)

			dm_build_728.dm_build_700.Dm_build_444(retBytes)

			dm_build_729.dm_build_1111(int32(len(retBytes)))
		}
	}

	if dm_build_728.dm_build_700.Dm_build_373() > Dm_build_977 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_729.dm_build_1106()
	if dm_build_728.dm_build_959(dm_build_729) {
		if dm_build_728.dm_build_699 != nil {
			dm_build_728.dm_build_700.Dm_build_378(0)
			if _, err := dm_build_728.dm_build_700.Dm_build_397(dm_build_728.dm_build_699); err != nil {
				return err
			}
		}
	} else {
		dm_build_728.dm_build_700.Dm_build_378(0)
		if _, err := dm_build_728.dm_build_700.Dm_build_397(dm_build_728.dm_build_698); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_733 *dm_build_697) dm_build_732(dm_build_734 dm_build_1097) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_736 := int32(0)
	if dm_build_733.dm_build_959(dm_build_734) {
		if dm_build_733.dm_build_699 != nil {
			dm_build_733.dm_build_700.Dm_build_375(0)
			if _, err := dm_build_733.dm_build_700.Dm_build_391(dm_build_733.dm_build_699, Dm_build_1005); err != nil {
				return err
			}

			dm_build_736 = dm_build_734.dm_build_1110()
			if dm_build_736 > 0 {
				if _, err := dm_build_733.dm_build_700.Dm_build_391(dm_build_733.dm_build_699, int(dm_build_736)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_733.dm_build_700.Dm_build_375(0)
		if _, err := dm_build_733.dm_build_700.Dm_build_391(dm_build_733.dm_build_698, Dm_build_1005); err != nil {
			return err
		}
		dm_build_736 = dm_build_734.dm_build_1110()

		if dm_build_736 > 0 {
			if _, err := dm_build_733.dm_build_700.Dm_build_391(dm_build_733.dm_build_698, int(dm_build_736)); err != nil {
				return err
			}
		}
	}

	dm_build_734.dm_build_1107()

	dm_build_736 = dm_build_734.dm_build_1110()
	if dm_build_736 <= 0 {
		return nil
	}

	if dm_build_733.dm_build_704 {
		ebytes := dm_build_733.dm_build_700.Dm_build_655(Dm_build_1005, int(dm_build_736))
		bytes, err := dm_build_733.dm_build_702.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_733.dm_build_700.Dm_build_375(Dm_build_1005)
		dm_build_733.dm_build_700.Dm_build_444(bytes)
		dm_build_734.dm_build_1111(int32(len(bytes)))
	}

	if dm_build_733.dm_build_723(dm_build_734) {

		dm_build_736 = dm_build_734.dm_build_1110()
		cbytes := dm_build_733.dm_build_700.Dm_build_655(Dm_build_1005+ULINT_SIZE, int(dm_build_736-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_733.dm_build_701.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_733.dm_build_700.Dm_build_375(Dm_build_1005)
		dm_build_733.dm_build_700.Dm_build_444(bytes)
		dm_build_734.dm_build_1111(int32(len(bytes)))
	}
	return nil
}

func (dm_build_738 *dm_build_697) dm_build_737(dm_build_739 dm_build_1097) (dm_build_740 interface{}, dm_build_741 error) {
	dm_build_741 = dm_build_739.dm_build_1101(dm_build_739)
	if dm_build_741 != nil {
		return nil, dm_build_741
	}

	dm_build_741 = dm_build_738.dm_build_727(dm_build_739)
	if dm_build_741 != nil {
		return nil, dm_build_741
	}

	dm_build_741 = dm_build_738.dm_build_732(dm_build_739)
	if dm_build_741 != nil {
		return nil, dm_build_741
	}

	return dm_build_739.dm_build_1105(dm_build_739)
}

func (dm_build_743 *dm_build_697) dm_build_742() (*dm_build_1551, error) {

	Dm_build_744 := dm_build_1557(dm_build_743)
	_, dm_build_745 := dm_build_743.dm_build_737(Dm_build_744)
	if dm_build_745 != nil {
		return nil, dm_build_745
	}

	return Dm_build_744, nil
}

func (dm_build_747 *dm_build_697) dm_build_746() error {

	dm_build_748 := dm_build_1418(dm_build_747)
	_, dm_build_749 := dm_build_747.dm_build_737(dm_build_748)
	if dm_build_749 != nil {
		return dm_build_749
	}

	return nil
}

func (dm_build_751 *dm_build_697) dm_build_750() error {

	var dm_build_752 *dm_build_1551
	var err error
	if dm_build_752, err = dm_build_751.dm_build_742(); err != nil {
		return err
	}

	if dm_build_751.dm_build_701.sslEncrypt == 2 {
		if err = dm_build_751.dm_build_955(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_751.dm_build_701.sslEncrypt == 1 {
		if err = dm_build_751.dm_build_955(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_751.dm_build_704 || dm_build_751.dm_build_703 {
		k, err := dm_build_751.dm_build_945()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_752.Dm_build_1555)
		encryptType := dm_build_752.dm_build_1553
		hashType := int(dm_build_752.Dm_build_1554)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_751.dm_build_948(encryptType, sessionKey, dm_build_751.dm_build_701.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_751.dm_build_746(); err != nil {
		return err
	}
	return nil
}

func (dm_build_755 *dm_build_697) Dm_build_754(dm_build_756 *DmStatement) error {
	dm_build_757 := dm_build_1580(dm_build_755, dm_build_756)
	_, dm_build_758 := dm_build_755.dm_build_737(dm_build_757)
	if dm_build_758 != nil {
		return dm_build_758
	}

	return nil
}

func (dm_build_760 *dm_build_697) Dm_build_759(dm_build_761 int32) error {
	dm_build_762 := dm_build_1590(dm_build_760, dm_build_761)
	_, dm_build_763 := dm_build_760.dm_build_737(dm_build_762)
	if dm_build_763 != nil {
		return dm_build_763
	}

	return nil
}

func (dm_build_765 *dm_build_697) Dm_build_764(dm_build_766 *DmStatement, dm_build_767 bool, dm_build_768 int16) (*execRetInfo, error) {
	dm_build_769 := dm_build_1457(dm_build_765, dm_build_766, dm_build_767, dm_build_768)
	dm_build_770, dm_build_771 := dm_build_765.dm_build_737(dm_build_769)
	if dm_build_771 != nil {
		return nil, dm_build_771
	}
	return dm_build_770.(*execRetInfo), nil
}

func (dm_build_773 *dm_build_697) Dm_build_772(dm_build_774 *DmStatement, dm_build_775 int16) (*execRetInfo, error) {
	return dm_build_773.Dm_build_764(dm_build_774, false, Dm_build_1057)
}

func (dm_build_777 *dm_build_697) Dm_build_776(dm_build_778 *DmStatement, dm_build_779 []OptParameter) (*execRetInfo, error) {
	dm_build_780, dm_build_781 := dm_build_777.dm_build_737(dm_build_1200(dm_build_777, dm_build_778, dm_build_779))
	if dm_build_781 != nil {
		return nil, dm_build_781
	}

	return dm_build_780.(*execRetInfo), nil
}

func (dm_build_783 *dm_build_697) Dm_build_782(dm_build_784 *DmStatement, dm_build_785 int16) (*execRetInfo, error) {
	return dm_build_783.Dm_build_764(dm_build_784, true, dm_build_785)
}

func (dm_build_787 *dm_build_697) Dm_build_786(dm_build_788 *DmStatement, dm_build_789 [][]interface{}) (*execRetInfo, error) {
	dm_build_790 := dm_build_1232(dm_build_787, dm_build_788, dm_build_789)
	dm_build_791, dm_build_792 := dm_build_787.dm_build_737(dm_build_790)
	if dm_build_792 != nil {
		return nil, dm_build_792
	}
	return dm_build_791.(*execRetInfo), nil
}

func (dm_build_794 *dm_build_697) Dm_build_793(dm_build_795 *DmStatement, dm_build_796 [][]interface{}, dm_build_797 bool) (*execRetInfo, error) {
	var dm_build_798, dm_build_799 = 0, 0
	var dm_build_800 = len(dm_build_796)
	var dm_build_801 [][]interface{}
	var dm_build_802 = NewExceInfo()
	dm_build_802.updateCounts = make([]int64, dm_build_800)
	var dm_build_803 = false
	for dm_build_798 < dm_build_800 {
		for dm_build_799 = dm_build_798; dm_build_799 < dm_build_800; dm_build_799++ {
			paramData := dm_build_796[dm_build_799]
			bindData := make([]interface{}, dm_build_795.paramCount)
			dm_build_803 = false
			for icol := 0; icol < int(dm_build_795.paramCount); icol++ {
				if dm_build_795.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_794.dm_build_928(bindData, paramData, icol) {
					dm_build_803 = true
					break
				}
			}

			if dm_build_803 {
				break
			}
			dm_build_801 = append(dm_build_801, bindData)
		}

		if dm_build_799 != dm_build_798 {
			tmpExecInfo, err := dm_build_794.Dm_build_786(dm_build_795, dm_build_801)
			if err != nil {
				return nil, err
			}
			dm_build_801 = dm_build_801[0:0]
			dm_build_802.union(tmpExecInfo, dm_build_798, dm_build_799-dm_build_798)
		}

		if dm_build_799 < dm_build_800 {
			tmpExecInfo, err := dm_build_794.Dm_build_804(dm_build_795, dm_build_796[dm_build_799], dm_build_797)
			if err != nil {
				return nil, err
			}

			dm_build_797 = true
			dm_build_802.union(tmpExecInfo, dm_build_799, 1)
		}

		dm_build_798 = dm_build_799 + 1
	}
	for _, i := range dm_build_802.updateCounts {
		if i > 0 {
			dm_build_802.updateCount += i
		}
	}
	return dm_build_802, nil
}

func (dm_build_805 *dm_build_697) Dm_build_804(dm_build_806 *DmStatement, dm_build_807 []interface{}, dm_build_808 bool) (*execRetInfo, error) {

	var dm_build_809 = make([]interface{}, dm_build_806.paramCount)
	for icol := 0; icol < int(dm_build_806.paramCount); icol++ {
		if dm_build_806.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_805.dm_build_928(dm_build_809, dm_build_807, icol) {

			if !dm_build_808 {
				preExecute := dm_build_1446(dm_build_805, dm_build_806, dm_build_806.bindParams)
				dm_build_805.dm_build_737(preExecute)
				dm_build_808 = true
			}

			dm_build_805.dm_build_934(dm_build_806, dm_build_806.bindParams[icol], icol, dm_build_807[icol].(iOffRowBinder))
			dm_build_809[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_810 = make([][]interface{}, 1, 1)
	dm_build_810[0] = dm_build_809

	dm_build_811 := dm_build_1232(dm_build_805, dm_build_806, dm_build_810)
	dm_build_812, dm_build_813 := dm_build_805.dm_build_737(dm_build_811)
	if dm_build_813 != nil {
		return nil, dm_build_813
	}
	return dm_build_812.(*execRetInfo), nil
}

func (dm_build_815 *dm_build_697) Dm_build_814(dm_build_816 *DmStatement, dm_build_817 int16) (*execRetInfo, error) {
	dm_build_818 := dm_build_1433(dm_build_815, dm_build_816, dm_build_817)

	dm_build_819, dm_build_820 := dm_build_815.dm_build_737(dm_build_818)
	if dm_build_820 != nil {
		return nil, dm_build_820
	}
	return dm_build_819.(*execRetInfo), nil
}

func (dm_build_822 *dm_build_697) Dm_build_821(dm_build_823 *innerRows, dm_build_824 int64) (*execRetInfo, error) {
	dm_build_825 := dm_build_1338(dm_build_822, dm_build_823, dm_build_824, INT64_MAX)
	dm_build_826, dm_build_827 := dm_build_822.dm_build_737(dm_build_825)
	if dm_build_827 != nil {
		return nil, dm_build_827
	}
	return dm_build_826.(*execRetInfo), nil
}

func (dm_build_829 *dm_build_697) Commit() error {
	dm_build_830 := dm_build_1185(dm_build_829)
	_, dm_build_831 := dm_build_829.dm_build_737(dm_build_830)
	if dm_build_831 != nil {
		return dm_build_831
	}

	return nil
}

func (dm_build_833 *dm_build_697) Rollback() error {
	dm_build_834 := dm_build_1495(dm_build_833)
	_, dm_build_835 := dm_build_833.dm_build_737(dm_build_834)
	if dm_build_835 != nil {
		return dm_build_835
	}

	return nil
}

func (dm_build_837 *dm_build_697) Dm_build_836(dm_build_838 *DmConnection) error {
	dm_build_839 := dm_build_1500(dm_build_837, dm_build_838.IsoLevel)
	_, dm_build_840 := dm_build_837.dm_build_737(dm_build_839)
	if dm_build_840 != nil {
		return dm_build_840
	}

	return nil
}

func (dm_build_842 *dm_build_697) Dm_build_841(dm_build_843 *DmStatement, dm_build_844 string) error {
	dm_build_845 := dm_build_1190(dm_build_842, dm_build_843, dm_build_844)
	_, dm_build_846 := dm_build_842.dm_build_737(dm_build_845)
	if dm_build_846 != nil {
		return dm_build_846
	}

	return nil
}

func (dm_build_848 *dm_build_697) Dm_build_847(dm_build_849 []uint32) ([]int64, error) {
	dm_build_850 := dm_build_1598(dm_build_848, dm_build_849)
	dm_build_851, dm_build_852 := dm_build_848.dm_build_737(dm_build_850)
	if dm_build_852 != nil {
		return nil, dm_build_852
	}
	return dm_build_851.([]int64), nil
}

func (dm_build_854 *dm_build_697) Close() error {
	if dm_build_854.dm_build_708 {
		return nil
	}

	dm_build_855 := dm_build_854.dm_build_698.Close()
	if dm_build_855 != nil {
		return dm_build_855
	}

	dm_build_854.dm_build_701 = nil
	dm_build_854.dm_build_708 = true
	return nil
}

func (dm_build_857 *dm_build_697) dm_build_856(dm_build_858 *lob) (int64, error) {
	dm_build_859 := dm_build_1369(dm_build_857, dm_build_858)
	dm_build_860, dm_build_861 := dm_build_857.dm_build_737(dm_build_859)
	if dm_build_861 != nil {
		return 0, dm_build_861
	}
	return dm_build_860.(int64), nil
}

func (dm_build_863 *dm_build_697) dm_build_862(dm_build_864 *lob, dm_build_865 int32, dm_build_866 int32) ([]byte, error) {
	dm_build_867 := dm_build_1356(dm_build_863, dm_build_864, int(dm_build_865), int(dm_build_866))
	dm_build_868, dm_build_869 := dm_build_863.dm_build_737(dm_build_867)
	if dm_build_869 != nil {
		return nil, dm_build_869
	}
	return dm_build_868.([]byte), nil
}

func (dm_build_871 *dm_build_697) dm_build_870(dm_build_872 *DmBlob, dm_build_873 int32, dm_build_874 int32) ([]byte, error) {
	var dm_build_875 = make([]byte, dm_build_874)
	var dm_build_876 int32 = 0
	var dm_build_877 int32 = 0
	var dm_build_878 []byte
	var dm_build_879 error
	for dm_build_876 < dm_build_874 {
		dm_build_877 = dm_build_874 - dm_build_876
		if dm_build_877 > Dm_build_1090 {
			dm_build_877 = Dm_build_1090
		}
		dm_build_878, dm_build_879 = dm_build_871.dm_build_862(&dm_build_872.lob, dm_build_873, dm_build_877)
		if dm_build_879 != nil {
			return nil, dm_build_879
		}
		if dm_build_878 == nil || len(dm_build_878) == 0 {
			break
		}
		Dm_build_1.Dm_build_57(dm_build_875, int(dm_build_876), dm_build_878, 0, len(dm_build_878))
		dm_build_876 += int32(len(dm_build_878))
		dm_build_873 += int32(len(dm_build_878))
		if dm_build_872.readOver {
			break
		}
	}
	return dm_build_875, nil
}

func (dm_build_881 *dm_build_697) dm_build_880(dm_build_882 *DmClob, dm_build_883 int32, dm_build_884 int32) (string, error) {
	var dm_build_885 bytes.Buffer
	var dm_build_886 int32 = 0
	var dm_build_887 int32 = 0
	var dm_build_888 []byte
	var dm_build_889 string
	var dm_build_890 error
	for dm_build_886 < dm_build_884 {
		dm_build_887 = dm_build_884 - dm_build_886
		if dm_build_887 > Dm_build_1090/2 {
			dm_build_887 = Dm_build_1090 / 2
		}
		dm_build_888, dm_build_890 = dm_build_881.dm_build_862(&dm_build_882.lob, dm_build_883, dm_build_887)
		if dm_build_890 != nil {
			return "", dm_build_890
		}
		if dm_build_888 == nil || len(dm_build_888) == 0 {
			break
		}
		dm_build_889 = Dm_build_1.Dm_build_158(dm_build_888, 0, len(dm_build_888), dm_build_882.serverEncoding, dm_build_881.dm_build_701)

		dm_build_885.WriteString(dm_build_889)
		strLen := utf8.RuneCountInString(dm_build_889)
		dm_build_886 += int32(strLen)
		dm_build_883 += int32(strLen)
		if dm_build_882.readOver {
			break
		}
	}
	return dm_build_885.String(), nil
}

func (dm_build_892 *dm_build_697) dm_build_891(dm_build_893 *DmClob, dm_build_894 int, dm_build_895 string, dm_build_896 string) (int, error) {
	var dm_build_897 = Dm_build_1.Dm_build_217(dm_build_895, dm_build_896, dm_build_892.dm_build_701)
	var dm_build_898 = 0
	var dm_build_899 = len(dm_build_897)
	var dm_build_900 = 0
	var dm_build_901 = 0
	var dm_build_902 = 0
	var dm_build_903 = dm_build_899/Dm_build_1089 + 1
	var dm_build_904 byte = 0
	var dm_build_905 byte = 0x01
	var dm_build_906 byte = 0x02
	for i := 0; i < dm_build_903; i++ {
		dm_build_904 = 0
		if i == 0 {
			dm_build_904 |= dm_build_905
		}
		if i == dm_build_903-1 {
			dm_build_904 |= dm_build_906
		}
		dm_build_902 = dm_build_899 - dm_build_901
		if dm_build_902 > Dm_build_1089 {
			dm_build_902 = Dm_build_1089
		}

		setLobData := dm_build_1514(dm_build_892, &dm_build_893.lob, dm_build_904, dm_build_894, dm_build_897, dm_build_898, dm_build_902)
		ret, err := dm_build_892.dm_build_737(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_900, nil
		} else {
			dm_build_894 += int(tmp)
			dm_build_900 += int(tmp)
			dm_build_901 += dm_build_902
			dm_build_898 += dm_build_902
		}
	}
	return dm_build_900, nil
}

func (dm_build_908 *dm_build_697) dm_build_907(dm_build_909 *DmBlob, dm_build_910 int, dm_build_911 []byte) (int, error) {
	var dm_build_912 = 0
	var dm_build_913 = len(dm_build_911)
	var dm_build_914 = 0
	var dm_build_915 = 0
	var dm_build_916 = 0
	var dm_build_917 = dm_build_913/Dm_build_1089 + 1
	var dm_build_918 byte = 0
	var dm_build_919 byte = 0x01
	var dm_build_920 byte = 0x02
	for i := 0; i < dm_build_917; i++ {
		dm_build_918 = 0
		if i == 0 {
			dm_build_918 |= dm_build_919
		}
		if i == dm_build_917-1 {
			dm_build_918 |= dm_build_920
		}
		dm_build_916 = dm_build_913 - dm_build_915
		if dm_build_916 > Dm_build_1089 {
			dm_build_916 = Dm_build_1089
		}

		setLobData := dm_build_1514(dm_build_908, &dm_build_909.lob, dm_build_918, dm_build_910, dm_build_911, dm_build_912, dm_build_916)
		ret, err := dm_build_908.dm_build_737(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_914, nil
		} else {
			dm_build_910 += int(tmp)
			dm_build_914 += int(tmp)
			dm_build_915 += dm_build_916
			dm_build_912 += dm_build_916
		}
	}
	return dm_build_914, nil
}

func (dm_build_922 *dm_build_697) dm_build_921(dm_build_923 *lob, dm_build_924 int) (int64, error) {
	dm_build_925 := dm_build_1380(dm_build_922, dm_build_923, dm_build_924)
	dm_build_926, dm_build_927 := dm_build_922.dm_build_737(dm_build_925)
	if dm_build_927 != nil {
		return dm_build_923.length, dm_build_927
	}
	return dm_build_926.(int64), nil
}

func (dm_build_929 *dm_build_697) dm_build_928(dm_build_930 []interface{}, dm_build_931 []interface{}, dm_build_932 int) bool {
	var dm_build_933 = false
	dm_build_930[dm_build_932] = dm_build_931[dm_build_932]

	if binder, ok := dm_build_931[dm_build_932].(iOffRowBinder); ok {
		dm_build_933 = true
		dm_build_930[dm_build_932] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_929.dm_build_701) {
			dm_build_930[dm_build_932] = &lobCtl{lob.buildCtlData()}
			dm_build_933 = false
		}
	} else {
		dm_build_930[dm_build_932] = dm_build_931[dm_build_932]
	}
	return dm_build_933
}

func (dm_build_935 *dm_build_697) dm_build_934(dm_build_936 *DmStatement, dm_build_937 parameter, dm_build_938 int, dm_build_939 iOffRowBinder) error {
	var dm_build_940 = Dm_build_287()
	dm_build_939.read(dm_build_940)
	var dm_build_941 = 0
	for !dm_build_939.isReadOver() || dm_build_940.Dm_build_288() > 0 {
		if !dm_build_939.isReadOver() && dm_build_940.Dm_build_288() < Dm_build_1089 {
			dm_build_939.read(dm_build_940)
		}
		if dm_build_940.Dm_build_288() > Dm_build_1089 {
			dm_build_941 = Dm_build_1089
		} else {
			dm_build_941 = dm_build_940.Dm_build_288()
		}

		putData := dm_build_1485(dm_build_935, dm_build_936, int16(dm_build_938), dm_build_940, int32(dm_build_941))
		_, err := dm_build_935.dm_build_737(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_943 *dm_build_697) dm_build_942() ([]byte, error) {
	var dm_build_944 error
	if dm_build_943.dm_build_705 == nil {
		if dm_build_943.dm_build_705, dm_build_944 = security.NewClientKeyPair(); dm_build_944 != nil {
			return nil, dm_build_944
		}
	}
	return security.Bn2Bytes(dm_build_943.dm_build_705.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_946 *dm_build_697) dm_build_945() (*security.DhKey, error) {
	var dm_build_947 error
	if dm_build_946.dm_build_705 == nil {
		if dm_build_946.dm_build_705, dm_build_947 = security.NewClientKeyPair(); dm_build_947 != nil {
			return nil, dm_build_947
		}
	}
	return dm_build_946.dm_build_705, nil
}

func (dm_build_949 *dm_build_697) dm_build_948(dm_build_950 int, dm_build_951 []byte, dm_build_952 string, dm_build_953 int) (dm_build_954 error) {
	if dm_build_950 > 0 && dm_build_950 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_951 != nil {
		dm_build_949.dm_build_702, dm_build_954 = security.NewSymmCipher(dm_build_950, dm_build_951)
	} else if dm_build_950 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_949.dm_build_702, dm_build_954 = security.NewThirdPartCipher(dm_build_950, dm_build_951, dm_build_952, dm_build_953); dm_build_954 != nil {
			dm_build_954 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_954.Error()).throw()
		}
	}
	return
}

func (dm_build_956 *dm_build_697) dm_build_955(dm_build_957 bool) (dm_build_958 error) {
	if dm_build_956.dm_build_699, dm_build_958 = security.NewTLSFromTCP(dm_build_956.dm_build_698, dm_build_956.dm_build_701.dmConnector.sslCertPath, dm_build_956.dm_build_701.dmConnector.sslKeyPath, dm_build_956.dm_build_701.dmConnector.user); dm_build_958 != nil {
		return
	}
	if !dm_build_957 {
		dm_build_956.dm_build_699 = nil
	}
	return
}

func (dm_build_960 *dm_build_697) dm_build_959(dm_build_961 dm_build_1097) bool {
	return dm_build_961.dm_build_1112() != Dm_build_1004 && dm_build_960.dm_build_701.sslEncrypt == 1
}
