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
	Dm_build_0 = 8192
	Dm_build_1 = 2 * time.Second
)

type dm_build_2 struct {
	dm_build_3  *net.TCPConn
	dm_build_4  *tls.Conn
	dm_build_5  *Dm_build_1239
	dm_build_6  *DmConnection
	dm_build_7  security.Cipher
	dm_build_8  bool
	dm_build_9  bool
	dm_build_10 *security.DhKey

	dm_build_11 bool
	dm_build_12 string
	dm_build_13 bool
}

func dm_build_14(dm_build_15 *DmConnection) (*dm_build_2, error) {
	dm_build_16, dm_build_17 := dm_build_19(dm_build_15.dmConnector.host+":"+strconv.Itoa(int(dm_build_15.dmConnector.port)), time.Duration(dm_build_15.dmConnector.socketTimeout)*time.Second)
	if dm_build_17 != nil {
		return nil, dm_build_17
	}

	dm_build_18 := dm_build_2{}
	dm_build_18.dm_build_3 = dm_build_16
	dm_build_18.dm_build_5 = Dm_build_1242(Dm_build_273)
	dm_build_18.dm_build_6 = dm_build_15
	dm_build_18.dm_build_8 = false
	dm_build_18.dm_build_9 = false
	dm_build_18.dm_build_11 = false
	dm_build_18.dm_build_12 = ""
	dm_build_18.dm_build_13 = false
	dm_build_15.Access = &dm_build_18

	return &dm_build_18, nil
}

func dm_build_19(dm_build_20 string, dm_build_21 time.Duration) (*net.TCPConn, error) {
	dm_build_22, dm_build_23 := net.DialTimeout("tcp", dm_build_20, dm_build_21)
	if dm_build_23 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_20).throw()
	}

	if tcpConn, ok := dm_build_22.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_25 *dm_build_2) dm_build_24(dm_build_26 dm_build_393) bool {
	var dm_build_27 = dm_build_25.dm_build_6.dmConnector.compress
	if dm_build_26.dm_build_408() == Dm_build_300 || dm_build_27 == Dm_build_349 {
		return false
	}

	if dm_build_27 == Dm_build_347 {
		return true
	} else if dm_build_27 == Dm_build_348 {
		return !dm_build_25.dm_build_6.Local && dm_build_26.dm_build_406() > Dm_build_346
	}

	return false
}

func (dm_build_29 *dm_build_2) dm_build_28(dm_build_30 dm_build_393) bool {
	var dm_build_31 = dm_build_29.dm_build_6.dmConnector.compress
	if dm_build_30.dm_build_408() == Dm_build_300 || dm_build_31 == Dm_build_349 {
		return false
	}

	if dm_build_31 == Dm_build_347 {
		return true
	} else if dm_build_31 == Dm_build_348 {
		return dm_build_29.dm_build_5.Dm_build_1502(Dm_build_308) == 1
	}

	return false
}

func (dm_build_33 *dm_build_2) dm_build_32(dm_build_34 dm_build_393) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_36 := dm_build_34.dm_build_406()

	if dm_build_36 > 0 {

		if dm_build_33.dm_build_24(dm_build_34) {
			var retBytes, err = Compress(dm_build_33.dm_build_5, Dm_build_301, int(dm_build_36), int(dm_build_33.dm_build_6.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_33.dm_build_5.Dm_build_1253(Dm_build_301)

			dm_build_33.dm_build_5.Dm_build_1290(dm_build_36)

			dm_build_33.dm_build_5.Dm_build_1318(retBytes)

			dm_build_34.dm_build_407(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_33.dm_build_5.Dm_build_1422(Dm_build_308, 1)
		}

		if dm_build_33.dm_build_9 {
			dm_build_36 = dm_build_34.dm_build_406()
			var retBytes = dm_build_33.dm_build_7.Encrypt(dm_build_33.dm_build_5.Dm_build_1529(Dm_build_301, int(dm_build_36)), true)

			dm_build_33.dm_build_5.Dm_build_1253(Dm_build_301)

			dm_build_33.dm_build_5.Dm_build_1318(retBytes)

			dm_build_34.dm_build_407(int32(len(retBytes)))
		}
	}

	dm_build_34.dm_build_402()
	if dm_build_33.dm_build_264(dm_build_34) {
		if dm_build_33.dm_build_4 != nil {
			dm_build_33.dm_build_5.Dm_build_1256(0)
			dm_build_33.dm_build_5.Dm_build_1275(dm_build_33.dm_build_4)
		}
	} else {
		dm_build_33.dm_build_5.Dm_build_1256(0)
		dm_build_33.dm_build_5.Dm_build_1275(dm_build_33.dm_build_3)
	}
	return nil
}

func (dm_build_38 *dm_build_2) dm_build_37(dm_build_39 dm_build_393) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_41 := int32(0)
	if dm_build_38.dm_build_264(dm_build_39) {
		if dm_build_38.dm_build_4 != nil {
			dm_build_38.dm_build_5.Dm_build_1253(0)
			dm_build_38.dm_build_5.Dm_build_1269(dm_build_38.dm_build_4, Dm_build_301)
			dm_build_41 = dm_build_39.dm_build_406()
			if dm_build_41 > 0 {
				dm_build_38.dm_build_5.Dm_build_1269(dm_build_38.dm_build_4, int(dm_build_41))
			}
		}
	} else {

		dm_build_38.dm_build_5.Dm_build_1253(0)
		dm_build_38.dm_build_5.Dm_build_1269(dm_build_38.dm_build_3, Dm_build_301)
		dm_build_41 = dm_build_39.dm_build_406()

		if dm_build_41 > 0 {
			dm_build_38.dm_build_5.Dm_build_1269(dm_build_38.dm_build_3, int(dm_build_41))
		}
	}

	dm_build_39.dm_build_403()

	dm_build_41 = dm_build_39.dm_build_406()
	if dm_build_41 <= 0 {
		return nil
	}

	if dm_build_38.dm_build_9 {
		ebytes := dm_build_38.dm_build_5.Dm_build_1529(Dm_build_301, int(dm_build_41))
		bytes, err := dm_build_38.dm_build_7.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_38.dm_build_5.Dm_build_1253(Dm_build_301)
		dm_build_38.dm_build_5.Dm_build_1318(bytes)
		dm_build_39.dm_build_407(int32(len(bytes)))
	}

	if dm_build_38.dm_build_28(dm_build_39) {

		dm_build_41 = dm_build_39.dm_build_406()
		cbytes := dm_build_38.dm_build_5.Dm_build_1529(Dm_build_301+ULINT_SIZE, int(dm_build_41-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_38.dm_build_6.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_38.dm_build_5.Dm_build_1253(Dm_build_301)
		dm_build_38.dm_build_5.Dm_build_1318(bytes)
		dm_build_39.dm_build_407(int32(len(bytes)))
	}
	return nil
}

func (dm_build_43 *dm_build_2) dm_build_42(dm_build_44 dm_build_393) (dm_build_45 interface{}, dm_build_46 error) {
	dm_build_46 = dm_build_44.dm_build_397(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	dm_build_46 = dm_build_43.dm_build_32(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	dm_build_46 = dm_build_43.dm_build_37(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	return dm_build_44.dm_build_401(dm_build_44)
}

func (dm_build_48 *dm_build_2) dm_build_47() (*dm_build_827, error) {

	Dm_build_49 := dm_build_833(dm_build_48)
	_, dm_build_50 := dm_build_48.dm_build_42(Dm_build_49)
	if dm_build_50 != nil {
		return nil, dm_build_50
	}

	return Dm_build_49, nil
}

func (dm_build_52 *dm_build_2) dm_build_51() error {

	dm_build_53 := dm_build_698(dm_build_52)
	_, dm_build_54 := dm_build_52.dm_build_42(dm_build_53)
	if dm_build_54 != nil {
		return dm_build_54
	}

	return nil
}

func (dm_build_56 *dm_build_2) dm_build_55() error {

	var dm_build_57 *dm_build_827
	var err error
	if dm_build_57, err = dm_build_56.dm_build_47(); err != nil {
		return err
	}

	if dm_build_56.dm_build_6.sslEncrypt == 2 {
		if err = dm_build_56.dm_build_260(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_56.dm_build_6.sslEncrypt == 1 {
		if err = dm_build_56.dm_build_260(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_56.dm_build_9 || dm_build_56.dm_build_8 {
		k, err := dm_build_56.dm_build_250()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_57.Dm_build_831)
		encryptType := dm_build_57.dm_build_829
		hashType := int(dm_build_57.Dm_build_830)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_56.dm_build_253(encryptType, sessionKey, dm_build_56.dm_build_6.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_56.dm_build_51(); err != nil {
		return err
	}
	return nil
}

func (dm_build_60 *dm_build_2) Dm_build_59(dm_build_61 *DmStatement) error {
	dm_build_62 := dm_build_856(dm_build_60, dm_build_61)
	_, dm_build_63 := dm_build_60.dm_build_42(dm_build_62)
	if dm_build_63 != nil {
		return dm_build_63
	}

	return nil
}

func (dm_build_65 *dm_build_2) Dm_build_64(dm_build_66 int32) error {
	dm_build_67 := dm_build_866(dm_build_65, dm_build_66)
	_, dm_build_68 := dm_build_65.dm_build_42(dm_build_67)
	if dm_build_68 != nil {
		return dm_build_68
	}

	return nil
}

func (dm_build_70 *dm_build_2) Dm_build_69(dm_build_71 *DmStatement, dm_build_72 bool, dm_build_73 int16) (*execRetInfo, error) {
	dm_build_74 := dm_build_735(dm_build_70, dm_build_71, dm_build_72, dm_build_73)
	dm_build_75, dm_build_76 := dm_build_70.dm_build_42(dm_build_74)
	if dm_build_76 != nil {
		return nil, dm_build_76
	}
	return dm_build_75.(*execRetInfo), nil
}

func (dm_build_78 *dm_build_2) Dm_build_77(dm_build_79 *DmStatement, dm_build_80 int16) (*execRetInfo, error) {
	return dm_build_78.Dm_build_69(dm_build_79, false, Dm_build_353)
}

func (dm_build_82 *dm_build_2) Dm_build_81(dm_build_83 *DmStatement, dm_build_84 []OptParameter) (*execRetInfo, error) {
	dm_build_85, dm_build_86 := dm_build_82.dm_build_42(dm_build_495(dm_build_82, dm_build_83, dm_build_84))
	if dm_build_86 != nil {
		return nil, dm_build_86
	}

	return dm_build_85.(*execRetInfo), nil
}

func (dm_build_88 *dm_build_2) Dm_build_87(dm_build_89 *DmStatement, dm_build_90 int16) (*execRetInfo, error) {
	return dm_build_88.Dm_build_69(dm_build_89, true, dm_build_90)
}

func (dm_build_92 *dm_build_2) Dm_build_91(dm_build_93 *DmStatement, dm_build_94 [][]interface{}) (*execRetInfo, error) {
	dm_build_95 := dm_build_518(dm_build_92, dm_build_93, dm_build_94)
	dm_build_96, dm_build_97 := dm_build_92.dm_build_42(dm_build_95)
	if dm_build_97 != nil {
		return nil, dm_build_97
	}
	return dm_build_96.(*execRetInfo), nil
}

func (dm_build_99 *dm_build_2) Dm_build_98(dm_build_100 *DmStatement, dm_build_101 [][]interface{}, dm_build_102 bool) (*execRetInfo, error) {
	var dm_build_103, dm_build_104 = 0, 0
	var dm_build_105 = len(dm_build_101)
	var dm_build_106 [][]interface{}
	var dm_build_107 = NewExceInfo()
	dm_build_107.updateCounts = make([]int64, dm_build_105)
	var dm_build_108 = false
	for dm_build_103 < dm_build_105 {
		for dm_build_104 = dm_build_103; dm_build_104 < dm_build_105; dm_build_104++ {
			paramData := dm_build_101[dm_build_104]
			bindData := make([]interface{}, dm_build_100.paramCount)
			dm_build_108 = false
			for icol := 0; icol < int(dm_build_100.paramCount); icol++ {
				if dm_build_100.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_99.dm_build_233(bindData, paramData, icol) {
					dm_build_108 = true
					break
				}
			}

			if dm_build_108 {
				break
			}
			dm_build_106 = append(dm_build_106, bindData)
		}

		if dm_build_104 != dm_build_103 {
			tmpExecInfo, err := dm_build_99.Dm_build_91(dm_build_100, dm_build_106)
			if err != nil {
				return nil, err
			}
			dm_build_106 = dm_build_106[0:0]
			dm_build_107.union(tmpExecInfo, dm_build_103, dm_build_104-dm_build_103)
		}

		if dm_build_104 < dm_build_105 {
			tmpExecInfo, err := dm_build_99.Dm_build_109(dm_build_100, dm_build_101[dm_build_104], dm_build_102)
			if err != nil {
				return nil, err
			}

			dm_build_102 = true
			dm_build_107.union(tmpExecInfo, dm_build_104, 1)
		}

		dm_build_103 = dm_build_104 + 1
	}
	for _, i := range dm_build_107.updateCounts {
		if i > 0 {
			dm_build_107.updateCount += i
		}
	}
	return dm_build_107, nil
}

func (dm_build_110 *dm_build_2) Dm_build_109(dm_build_111 *DmStatement, dm_build_112 []interface{}, dm_build_113 bool) (*execRetInfo, error) {

	var dm_build_114 = make([]interface{}, dm_build_111.paramCount)
	for icol := 0; icol < int(dm_build_111.paramCount); icol++ {
		if dm_build_111.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_110.dm_build_233(dm_build_114, dm_build_112, icol) {

			if !dm_build_113 {
				preExecute := dm_build_725(dm_build_110, dm_build_111, dm_build_111.params)
				dm_build_110.dm_build_42(preExecute)
				dm_build_113 = true
			}

			dm_build_110.dm_build_239(dm_build_111, dm_build_111.params[icol], icol, dm_build_112[icol].(iOffRowBinder))
			dm_build_114[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_115 = make([][]interface{}, 1, 1)
	dm_build_115[0] = dm_build_114

	dm_build_116 := dm_build_518(dm_build_110, dm_build_111, dm_build_115)
	dm_build_117, dm_build_118 := dm_build_110.dm_build_42(dm_build_116)
	if dm_build_118 != nil {
		return nil, dm_build_118
	}
	return dm_build_117.(*execRetInfo), nil
}

func (dm_build_120 *dm_build_2) Dm_build_119(dm_build_121 *DmStatement, dm_build_122 int16) (*execRetInfo, error) {
	dm_build_123 := dm_build_712(dm_build_120, dm_build_121, dm_build_122)

	dm_build_124, dm_build_125 := dm_build_120.dm_build_42(dm_build_123)
	if dm_build_125 != nil {
		return nil, dm_build_125
	}
	return dm_build_124.(*execRetInfo), nil
}

func (dm_build_127 *dm_build_2) Dm_build_126(dm_build_128 *innerRows, dm_build_129 int64) (*execRetInfo, error) {
	dm_build_130 := dm_build_618(dm_build_127, dm_build_128, dm_build_129, INT64_MAX)
	dm_build_131, dm_build_132 := dm_build_127.dm_build_42(dm_build_130)
	if dm_build_132 != nil {
		return nil, dm_build_132
	}
	return dm_build_131.(*execRetInfo), nil
}

func (dm_build_134 *dm_build_2) Commit() error {
	dm_build_135 := dm_build_481(dm_build_134)
	_, dm_build_136 := dm_build_134.dm_build_42(dm_build_135)
	if dm_build_136 != nil {
		return dm_build_136
	}

	return nil
}

func (dm_build_138 *dm_build_2) Rollback() error {
	dm_build_139 := dm_build_773(dm_build_138)
	_, dm_build_140 := dm_build_138.dm_build_42(dm_build_139)
	if dm_build_140 != nil {
		return dm_build_140
	}

	return nil
}

func (dm_build_142 *dm_build_2) Dm_build_141(dm_build_143 *DmConnection) error {
	dm_build_144 := dm_build_778(dm_build_142, dm_build_143.IsoLevel)
	_, dm_build_145 := dm_build_142.dm_build_42(dm_build_144)
	if dm_build_145 != nil {
		return dm_build_145
	}

	return nil
}

func (dm_build_147 *dm_build_2) Dm_build_146(dm_build_148 *DmStatement, dm_build_149 string) error {
	dm_build_150 := dm_build_486(dm_build_147, dm_build_148, dm_build_149)
	_, dm_build_151 := dm_build_147.dm_build_42(dm_build_150)
	if dm_build_151 != nil {
		return dm_build_151
	}

	return nil
}

func (dm_build_153 *dm_build_2) Dm_build_152(dm_build_154 []uint32) ([]int64, error) {
	dm_build_155 := dm_build_874(dm_build_153, dm_build_154)
	dm_build_156, dm_build_157 := dm_build_153.dm_build_42(dm_build_155)
	if dm_build_157 != nil {
		return nil, dm_build_157
	}
	return dm_build_156.([]int64), nil
}

func (dm_build_159 *dm_build_2) Close() error {
	if dm_build_159.dm_build_13 {
		return nil
	}

	dm_build_160 := dm_build_159.dm_build_3.Close()
	if dm_build_160 != nil {
		return dm_build_160
	}

	dm_build_159.dm_build_6 = nil
	dm_build_159.dm_build_13 = true
	return nil
}

func (dm_build_162 *dm_build_2) dm_build_161(dm_build_163 *lob) (int64, error) {
	dm_build_164 := dm_build_649(dm_build_162, dm_build_163)
	dm_build_165, dm_build_166 := dm_build_162.dm_build_42(dm_build_164)
	if dm_build_166 != nil {
		return 0, dm_build_166
	}
	return dm_build_165.(int64), nil
}

func (dm_build_168 *dm_build_2) dm_build_167(dm_build_169 *lob, dm_build_170 int32, dm_build_171 int32) ([]byte, error) {
	dm_build_172 := dm_build_636(dm_build_168, dm_build_169, int(dm_build_170), int(dm_build_171))
	dm_build_173, dm_build_174 := dm_build_168.dm_build_42(dm_build_172)
	if dm_build_174 != nil {
		return nil, dm_build_174
	}
	return dm_build_173.([]byte), nil
}

func (dm_build_176 *dm_build_2) dm_build_175(dm_build_177 *DmBlob, dm_build_178 int32, dm_build_179 int32) ([]byte, error) {
	var dm_build_180 = make([]byte, dm_build_179)
	var dm_build_181 int32 = 0
	var dm_build_182 int32 = 0
	var dm_build_183 []byte
	var dm_build_184 error
	for dm_build_181 < dm_build_179 {
		dm_build_182 = dm_build_179 - dm_build_181
		if dm_build_182 > Dm_build_386 {
			dm_build_182 = Dm_build_386
		}
		dm_build_183, dm_build_184 = dm_build_176.dm_build_167(&dm_build_177.lob, dm_build_178, dm_build_182)
		if dm_build_184 != nil {
			return nil, dm_build_184
		}
		if dm_build_183 == nil || len(dm_build_183) == 0 {
			break
		}
		Dm_build_885.Dm_build_941(dm_build_180, int(dm_build_181), dm_build_183, 0, len(dm_build_183))
		dm_build_181 += int32(len(dm_build_183))
		dm_build_178 += int32(len(dm_build_183))
		if dm_build_177.readOver {
			break
		}
	}
	return dm_build_180, nil
}

func (dm_build_186 *dm_build_2) dm_build_185(dm_build_187 *DmClob, dm_build_188 int32, dm_build_189 int32) (string, error) {
	var dm_build_190 bytes.Buffer
	var dm_build_191 int32 = 0
	var dm_build_192 int32 = 0
	var dm_build_193 []byte
	var dm_build_194 string
	var dm_build_195 error
	for dm_build_191 < dm_build_189 {
		dm_build_192 = dm_build_189 - dm_build_191
		if dm_build_192 > Dm_build_386/2 {
			dm_build_192 = Dm_build_386 / 2
		}
		dm_build_193, dm_build_195 = dm_build_186.dm_build_167(&dm_build_187.lob, dm_build_188, dm_build_192)
		if dm_build_195 != nil {
			return "", dm_build_195
		}
		if dm_build_193 == nil || len(dm_build_193) == 0 {
			break
		}
		dm_build_194 = Dm_build_885.Dm_build_1039(dm_build_193, 0, len(dm_build_193), dm_build_187.serverEncoding, dm_build_186.dm_build_6)

		dm_build_190.WriteString(dm_build_194)
		dm_build_191 += int32(len(dm_build_194))
		dm_build_188 += int32(len(dm_build_194))
		if dm_build_187.readOver {
			break
		}
	}
	return dm_build_190.String(), nil
}

func (dm_build_197 *dm_build_2) dm_build_196(dm_build_198 *DmClob, dm_build_199 int, dm_build_200 string, dm_build_201 string) (int, error) {
	var dm_build_202 = Dm_build_885.Dm_build_1095(dm_build_200, dm_build_201, dm_build_197.dm_build_6)
	var dm_build_203 = 0
	var dm_build_204 = len(dm_build_202)
	var dm_build_205 = 0
	var dm_build_206 = 0
	var dm_build_207 = 0
	var dm_build_208 = dm_build_204/Dm_build_385 + 1
	var dm_build_209 byte = 0
	var dm_build_210 byte = 0x01
	var dm_build_211 byte = 0x02
	for i := 0; i < dm_build_208; i++ {
		dm_build_209 = 0
		if i == 0 {
			dm_build_209 |= dm_build_210
		}
		if i == dm_build_208-1 {
			dm_build_209 |= dm_build_211
		}
		dm_build_207 = dm_build_204 - dm_build_206
		if dm_build_207 > Dm_build_385 {
			dm_build_207 = Dm_build_385
		}

		setLobData := dm_build_792(dm_build_197, &dm_build_198.lob, dm_build_209, dm_build_199, dm_build_202, dm_build_203, dm_build_207)
		ret, err := dm_build_197.dm_build_42(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_205, nil
		} else {
			dm_build_199 += int(tmp)
			dm_build_205 += int(tmp)
			dm_build_206 += dm_build_207
			dm_build_203 += dm_build_207
		}
	}
	return dm_build_205, nil
}

func (dm_build_213 *dm_build_2) dm_build_212(dm_build_214 *DmBlob, dm_build_215 int, dm_build_216 []byte) (int, error) {
	var dm_build_217 = 0
	var dm_build_218 = len(dm_build_216)
	var dm_build_219 = 0
	var dm_build_220 = 0
	var dm_build_221 = 0
	var dm_build_222 = dm_build_218/Dm_build_385 + 1
	var dm_build_223 byte = 0
	var dm_build_224 byte = 0x01
	var dm_build_225 byte = 0x02
	for i := 0; i < dm_build_222; i++ {
		dm_build_223 = 0
		if i == 0 {
			dm_build_223 |= dm_build_224
		}
		if i == dm_build_222-1 {
			dm_build_223 |= dm_build_225
		}
		dm_build_221 = dm_build_218 - dm_build_220
		if dm_build_221 > Dm_build_385 {
			dm_build_221 = Dm_build_385
		}

		setLobData := dm_build_792(dm_build_213, &dm_build_214.lob, dm_build_223, dm_build_215, dm_build_216, dm_build_217, dm_build_221)
		ret, err := dm_build_213.dm_build_42(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_219, nil
		} else {
			dm_build_215 += int(tmp)
			dm_build_219 += int(tmp)
			dm_build_220 += dm_build_221
			dm_build_217 += dm_build_221
		}
	}
	return dm_build_219, nil
}

func (dm_build_227 *dm_build_2) dm_build_226(dm_build_228 *lob, dm_build_229 int) (int64, error) {
	dm_build_230 := dm_build_660(dm_build_227, dm_build_228, dm_build_229)
	dm_build_231, dm_build_232 := dm_build_227.dm_build_42(dm_build_230)
	if dm_build_232 != nil {
		return dm_build_228.length, dm_build_232
	}
	return dm_build_231.(int64), nil
}

func (dm_build_234 *dm_build_2) dm_build_233(dm_build_235 []interface{}, dm_build_236 []interface{}, dm_build_237 int) bool {
	var dm_build_238 = false
	if dm_build_237 >= len(dm_build_236) || dm_build_236[dm_build_237] == nil {
		dm_build_235[dm_build_237] = ParamDataEnum_Null
	} else if binder, ok := dm_build_236[dm_build_237].(iOffRowBinder); ok {
		dm_build_238 = true
		dm_build_235[dm_build_237] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_234.dm_build_6) {
			dm_build_235[dm_build_237] = &lobCtl{lob.buildCtlData()}
			dm_build_238 = false
		}
	} else {
		dm_build_235[dm_build_237] = dm_build_236[dm_build_237]
	}
	return dm_build_238
}

func (dm_build_240 *dm_build_2) dm_build_239(dm_build_241 *DmStatement, dm_build_242 parameter, dm_build_243 int, dm_build_244 iOffRowBinder) error {
	var dm_build_245 = Dm_build_1165()
	dm_build_244.read(dm_build_245)
	var dm_build_246 = 0
	for !dm_build_244.isReadOver() || dm_build_245.Dm_build_1166() > 0 {
		if !dm_build_244.isReadOver() && dm_build_245.Dm_build_1166() < Dm_build_385 {
			dm_build_244.read(dm_build_245)
		}
		if dm_build_245.Dm_build_1166() > Dm_build_385 {
			dm_build_246 = Dm_build_385
		} else {
			dm_build_246 = dm_build_245.Dm_build_1166()
		}

		putData := dm_build_763(dm_build_240, dm_build_241, int16(dm_build_243), dm_build_245, int32(dm_build_246))
		_, err := dm_build_240.dm_build_42(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_248 *dm_build_2) dm_build_247() ([]byte, error) {
	var dm_build_249 error
	if dm_build_248.dm_build_10 == nil {
		if dm_build_248.dm_build_10, dm_build_249 = security.NewClientKeyPair(); dm_build_249 != nil {
			return nil, dm_build_249
		}
	}
	return security.Bn2Bytes(dm_build_248.dm_build_10.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_251 *dm_build_2) dm_build_250() (*security.DhKey, error) {
	var dm_build_252 error
	if dm_build_251.dm_build_10 == nil {
		if dm_build_251.dm_build_10, dm_build_252 = security.NewClientKeyPair(); dm_build_252 != nil {
			return nil, dm_build_252
		}
	}
	return dm_build_251.dm_build_10, nil
}

func (dm_build_254 *dm_build_2) dm_build_253(dm_build_255 int, dm_build_256 []byte, dm_build_257 string, dm_build_258 int) (dm_build_259 error) {
	if dm_build_255 > 0 && dm_build_255 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_256 != nil {
		dm_build_254.dm_build_7, dm_build_259 = security.NewSymmCipher(dm_build_255, dm_build_256)
	} else if dm_build_255 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_254.dm_build_7, dm_build_259 = security.NewThirdPartCipher(dm_build_255, dm_build_256, dm_build_257, dm_build_258); dm_build_259 != nil {
			dm_build_259 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_259.Error()).throw()
		}
	}
	return
}

func (dm_build_261 *dm_build_2) dm_build_260(dm_build_262 bool) (dm_build_263 error) {
	if dm_build_261.dm_build_4, dm_build_263 = security.NewTLSFromTCP(dm_build_261.dm_build_3, dm_build_261.dm_build_6.dmConnector.sslCertPath, dm_build_261.dm_build_6.dmConnector.sslKeyPath, dm_build_261.dm_build_6.dmConnector.user); dm_build_263 != nil {
		return
	}
	if !dm_build_262 {
		dm_build_261.dm_build_4 = nil
	}
	return
}

func (dm_build_265 *dm_build_2) dm_build_264(dm_build_266 dm_build_393) bool {
	return dm_build_266.dm_build_408() != Dm_build_300 && dm_build_265.dm_build_6.sslEncrypt == 1
}
