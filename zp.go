/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_0 = "7.6.0.0"

	Dm_build_1 = "7.0.0.9"

	Dm_build_2 = "8.0.0.73"

	Dm_build_3 = "7.1.2.128"

	Dm_build_4 = "7.1.5.144"

	Dm_build_5 = "7.1.6.123"

	Dm_build_6 = 32768 - 128

	Dm_build_7 int16 = 1

	Dm_build_8 int16 = 2

	Dm_build_9 int16 = 3

	Dm_build_10 int16 = 4

	Dm_build_11 int16 = 5

	Dm_build_12 int16 = 6

	Dm_build_13 int16 = 7

	Dm_build_14 int16 = 8

	Dm_build_15 int16 = 9

	Dm_build_16 int16 = 13

	Dm_build_17 int16 = 14

	Dm_build_18 int16 = 15

	Dm_build_19 int16 = 17

	Dm_build_20 int16 = 21

	Dm_build_21 int16 = 24

	Dm_build_22 int16 = 27

	Dm_build_23 int16 = 29

	Dm_build_24 int16 = 30

	Dm_build_25 int16 = 31

	Dm_build_26 int16 = 32

	Dm_build_27 int16 = 44

	Dm_build_28 int16 = 52

	Dm_build_29 int16 = 60

	Dm_build_30 int16 = 71

	Dm_build_31 int16 = 90

	Dm_build_32 int16 = 91

	Dm_build_33 int16 = 200

	Dm_build_34 = 64

	Dm_build_35 = 20

	Dm_build_36 = 0

	Dm_build_37 = 4

	Dm_build_38 = 6

	Dm_build_39 = 10

	Dm_build_40 = 14

	Dm_build_41 = 18

	Dm_build_42 = 19

	Dm_build_43 = 128

	Dm_build_44 = 256

	Dm_build_45 = 0xffff

	Dm_build_46 int32 = 2

	Dm_build_47 = -1

	Dm_build_48 uint16 = 0xFFFE

	Dm_build_49 uint16 = uint16(Dm_build_48 - 3)

	Dm_build_50 uint16 = Dm_build_48

	Dm_build_51 int32 = 0xFFFF

	Dm_build_52 int32 = 0x80

	Dm_build_53 byte = 0x60

	Dm_build_54 uint16 = uint16(Dm_build_50)

	Dm_build_55 uint16 = uint16(Dm_build_51)

	Dm_build_56 int16 = 0x00

	Dm_build_57 int16 = 0x03

	Dm_build_58 int32 = 0x80

	Dm_build_59 byte = 0

	Dm_build_60 byte = 1

	Dm_build_61 byte = 2

	Dm_build_62 byte = 3

	Dm_build_63 byte = 4

	Dm_build_64 byte = Dm_build_59

	Dm_build_65 int = 10

	Dm_build_66 int32 = 32

	Dm_build_67 int32 = 65536

	Dm_build_68 byte = 0

	Dm_build_69 byte = 1

	Dm_build_70 int32 = 0x00000000

	Dm_build_71 int32 = 0x00000020

	Dm_build_72 int32 = 0x00000040

	Dm_build_73 int32 = 0x00000FFF

	Dm_build_74 int32 = 0

	Dm_build_75 int32 = 1

	Dm_build_76 int32 = 2

	Dm_build_77 int32 = 3

	Dm_build_78 = 8192

	Dm_build_79 = 1

	Dm_build_80 = 2

	Dm_build_81 = 0

	Dm_build_82 = 0

	Dm_build_83 = 1

	Dm_build_84 = -1

	Dm_build_85 int16 = 0

	Dm_build_86 int16 = 1

	Dm_build_87 int16 = 2

	Dm_build_88 int16 = 3

	Dm_build_89 int16 = 4

	Dm_build_90 int16 = 127

	Dm_build_91 int16 = Dm_build_90 + 20

	Dm_build_92 int16 = Dm_build_90 + 21

	Dm_build_93 int16 = Dm_build_90 + 22

	Dm_build_94 int16 = Dm_build_90 + 24

	Dm_build_95 int16 = Dm_build_90 + 25

	Dm_build_96 int16 = Dm_build_90 + 26

	Dm_build_97 int16 = Dm_build_90 + 30

	Dm_build_98 int16 = Dm_build_90 + 31

	Dm_build_99 int16 = Dm_build_90 + 32

	Dm_build_100 int16 = Dm_build_90 + 33

	Dm_build_101 int16 = Dm_build_90 + 35

	Dm_build_102 int16 = Dm_build_90 + 38

	Dm_build_103 int16 = Dm_build_90 + 39

	Dm_build_104 int16 = Dm_build_90 + 51

	Dm_build_105 int16 = Dm_build_90 + 71

	Dm_build_106 int16 = Dm_build_90 + 124

	Dm_build_107 int16 = Dm_build_90 + 125

	Dm_build_108 int16 = Dm_build_90 + 126

	Dm_build_109 int16 = Dm_build_90 + 127

	Dm_build_110 int16 = Dm_build_90 + 128

	Dm_build_111 int16 = Dm_build_90 + 129

	Dm_build_112 byte = 0

	Dm_build_113 byte = 2

	Dm_build_114 = 2048

	Dm_build_115 = -1

	Dm_build_116 = 0

	Dm_build_117 = 16000

	Dm_build_118 = 32000

	Dm_build_119 = 0x00000000

	Dm_build_120 = 0x00000020

	Dm_build_121 = 0x00000040

	Dm_build_122 = 0x00000FFF
)

type dm_build_123 interface {
	dm_build_124()
	dm_build_125() error
	dm_build_126()
	dm_build_127(imsg dm_build_123) error
	dm_build_128() error
	dm_build_129() (interface{}, error)
	dm_build_130()
	dm_build_131(imsg dm_build_123) (interface{}, error)
	dm_build_132()
	dm_build_133() error
	dm_build_134() byte
	dm_build_135() int32
	dm_build_136(length int32)
	dm_build_137() int16
}

type dm_build_138 struct {
	dm_build_139 *dm_build_1285

	dm_build_140 int16

	dm_build_141 int32

	dm_build_142 *DmStatement
}

func (dm_build_144 *dm_build_138) dm_build_143(dm_build_145 *dm_build_1285, dm_build_146 int16) *dm_build_138 {
	dm_build_144.dm_build_139 = dm_build_145
	dm_build_144.dm_build_140 = dm_build_146
	return dm_build_144
}

func (dm_build_148 *dm_build_138) dm_build_147(dm_build_149 *dm_build_1285, dm_build_150 int16, dm_build_151 *DmStatement) *dm_build_138 {
	dm_build_148.dm_build_143(dm_build_149, dm_build_150).dm_build_142 = dm_build_151
	return dm_build_148
}

func dm_build_152(dm_build_153 *dm_build_1285, dm_build_154 int16) *dm_build_138 {
	return new(dm_build_138).dm_build_143(dm_build_153, dm_build_154)
}

func dm_build_155(dm_build_156 *dm_build_1285, dm_build_157 int16, dm_build_158 *DmStatement) *dm_build_138 {
	return new(dm_build_138).dm_build_147(dm_build_156, dm_build_157, dm_build_158)
}

func (dm_build_160 *dm_build_138) dm_build_124() {
	dm_build_160.dm_build_139.dm_build_1288.Dm_build_967(0)
	dm_build_160.dm_build_139.dm_build_1288.Dm_build_978(Dm_build_34, true, true)
}

func (dm_build_162 *dm_build_138) dm_build_125() error {
	return nil
}

func (dm_build_164 *dm_build_138) dm_build_126() {
	if dm_build_164.dm_build_142 == nil {
		dm_build_164.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_36, 0)
	} else {
		dm_build_164.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_36, dm_build_164.dm_build_142.id)
	}

	dm_build_164.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_37, dm_build_164.dm_build_140)
	dm_build_164.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_38, int32(dm_build_164.dm_build_139.dm_build_1288.Dm_build_965()-Dm_build_34))
}

func (dm_build_166 *dm_build_138) dm_build_128() error {
	dm_build_166.dm_build_139.dm_build_1288.Dm_build_970(0)
	dm_build_166.dm_build_139.dm_build_1288.Dm_build_978(Dm_build_34, false, true)
	return dm_build_166.dm_build_171()
}

func (dm_build_168 *dm_build_138) dm_build_129() (interface{}, error) {
	return nil, nil
}

func (dm_build_170 *dm_build_138) dm_build_130() {
}

func (dm_build_172 *dm_build_138) dm_build_171() error {
	dm_build_172.dm_build_141 = dm_build_172.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_39)
	if dm_build_172.dm_build_141 < 0 && dm_build_172.dm_build_141 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_172.dm_build_141, dm_build_172.dm_build_173(), nil, ""}).throw()
	} else if dm_build_172.dm_build_141 > 0 {

	} else if dm_build_172.dm_build_140 == Dm_build_33 || dm_build_172.dm_build_140 == Dm_build_7 {
		dm_build_172.dm_build_173()
	}

	return nil
}

func (dm_build_174 *dm_build_138) dm_build_173() string {

	dm_build_175 := dm_build_174.dm_build_139.dm_build_1289.getServerEncoding()

	if dm_build_175 != "" && dm_build_175 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_175 = ENCODING_GB18030
	}

	dm_build_174.dm_build_139.dm_build_1288.Dm_build_978(int(dm_build_174.dm_build_139.dm_build_1288.Dm_build_1078()), false, true)

	dm_build_174.dm_build_139.dm_build_1288.Dm_build_978(int(dm_build_174.dm_build_139.dm_build_1288.Dm_build_1078()), false, true)

	dm_build_174.dm_build_139.dm_build_1288.Dm_build_978(int(dm_build_174.dm_build_139.dm_build_1288.Dm_build_1078()), false, true)

	return dm_build_174.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_175, dm_build_174.dm_build_139.dm_build_1289)
}

func (dm_build_177 *dm_build_138) dm_build_127(dm_build_178 dm_build_123) (dm_build_179 error) {
	dm_build_178.dm_build_124()
	if dm_build_179 = dm_build_178.dm_build_125(); dm_build_179 != nil {
		return dm_build_179
	}
	dm_build_178.dm_build_126()
	return nil
}

func (dm_build_181 *dm_build_138) dm_build_131(dm_build_182 dm_build_123) (dm_build_183 interface{}, dm_build_184 error) {
	dm_build_184 = dm_build_182.dm_build_128()
	if dm_build_184 != nil {
		return nil, dm_build_184
	}
	dm_build_183, dm_build_184 = dm_build_182.dm_build_129()
	if dm_build_184 != nil {
		return nil, dm_build_184
	}
	dm_build_182.dm_build_130()
	return dm_build_183, nil
}

func (dm_build_186 *dm_build_138) dm_build_132() {
	dm_build_186.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_42, dm_build_186.dm_build_134())
}

func (dm_build_188 *dm_build_138) dm_build_133() error {
	dm_build_189 := dm_build_188.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_42)
	dm_build_190 := dm_build_188.dm_build_134()
	if dm_build_189 != dm_build_190 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_192 *dm_build_138) dm_build_134() byte {
	dm_build_193 := dm_build_192.dm_build_139.dm_build_1288.Dm_build_1216(0)

	for i := 1; i < Dm_build_42; i++ {
		dm_build_193 ^= dm_build_192.dm_build_139.dm_build_1288.Dm_build_1216(i)
	}

	return dm_build_193
}

func (dm_build_195 *dm_build_138) dm_build_135() int32 {
	return dm_build_195.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_38)
}

func (dm_build_197 *dm_build_138) dm_build_136(dm_build_198 int32) {
	dm_build_197.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_38, dm_build_198)
}

func (dm_build_200 *dm_build_138) dm_build_137() int16 {
	return dm_build_200.dm_build_140
}

type dm_build_201 struct {
	dm_build_138
}

func dm_build_202(dm_build_203 *dm_build_1285) *dm_build_201 {
	dm_build_204 := new(dm_build_201)
	dm_build_204.dm_build_143(dm_build_203, Dm_build_14)
	return dm_build_204
}

type dm_build_205 struct {
	dm_build_138
	dm_build_206 string
}

func dm_build_207(dm_build_208 *dm_build_1285, dm_build_209 *DmStatement, dm_build_210 string) *dm_build_205 {
	dm_build_211 := new(dm_build_205)
	dm_build_211.dm_build_147(dm_build_208, Dm_build_22, dm_build_209)
	dm_build_211.dm_build_206 = dm_build_210
	dm_build_211.dm_build_142.cursorName = dm_build_210
	return dm_build_211
}

func (dm_build_213 *dm_build_205) dm_build_125() error {
	dm_build_213.dm_build_139.dm_build_1288.Dm_build_1066(dm_build_213.dm_build_206, dm_build_213.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_213.dm_build_139.dm_build_1289)
	dm_build_213.dm_build_139.dm_build_1288.Dm_build_1004(1)
	return nil
}

type Dm_build_214 struct {
	dm_build_230
	dm_build_215 []OptParameter
}

func dm_build_216(dm_build_217 *dm_build_1285, dm_build_218 *DmStatement, dm_build_219 []OptParameter) *Dm_build_214 {
	dm_build_220 := new(Dm_build_214)
	dm_build_220.dm_build_147(dm_build_217, Dm_build_32, dm_build_218)
	dm_build_220.dm_build_215 = dm_build_219
	return dm_build_220
}

func (dm_build_222 *Dm_build_214) dm_build_125() error {
	dm_build_223 := len(dm_build_222.dm_build_215)

	dm_build_222.dm_build_244(int16(dm_build_223), 1)

	dm_build_222.dm_build_139.dm_build_1288.Dm_build_1066(dm_build_222.dm_build_142.nativeSql, dm_build_222.dm_build_142.dmConn.getServerEncoding(), dm_build_222.dm_build_142.dmConn)

	for _, param := range dm_build_222.dm_build_215 {
		dm_build_222.dm_build_139.dm_build_1288.Dm_build_996(param.ioType)
		dm_build_222.dm_build_139.dm_build_1288.Dm_build_1004(int32(param.tp))
		dm_build_222.dm_build_139.dm_build_1288.Dm_build_1004(int32(param.prec))
		dm_build_222.dm_build_139.dm_build_1288.Dm_build_1004(int32(param.scale))
	}

	for _, param := range dm_build_222.dm_build_215 {
		if param.bytes == nil {
			dm_build_222.dm_build_139.dm_build_1288.Dm_build_1012(Dm_build_50)
		} else {
			dm_build_222.dm_build_139.dm_build_1288.Dm_build_1042(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_225 *Dm_build_214) dm_build_129() (interface{}, error) {
	return dm_build_225.dm_build_230.dm_build_129()
}

const (
	Dm_build_226 int = 0x01

	Dm_build_227 int = 0x02

	Dm_build_228 int = 0x04

	Dm_build_229 int = 0x08
)

type dm_build_230 struct {
	dm_build_138
	dm_build_231 [][]interface{}
	dm_build_232 []parameter
	dm_build_233 bool
}

func dm_build_234(dm_build_235 *dm_build_1285, dm_build_236 int16, dm_build_237 *DmStatement) *dm_build_230 {
	dm_build_238 := new(dm_build_230)
	dm_build_238.dm_build_147(dm_build_235, dm_build_236, dm_build_237)
	dm_build_238.dm_build_233 = true
	return dm_build_238
}

func dm_build_239(dm_build_240 *dm_build_1285, dm_build_241 *DmStatement, dm_build_242 [][]interface{}) *dm_build_230 {
	dm_build_243 := new(dm_build_230)

	if dm_build_240.dm_build_1289.Execute2 {
		dm_build_243.dm_build_147(dm_build_240, Dm_build_16, dm_build_241)
	} else {
		dm_build_243.dm_build_147(dm_build_240, Dm_build_12, dm_build_241)
	}

	dm_build_243.dm_build_232 = dm_build_241.params
	dm_build_243.dm_build_231 = dm_build_242
	dm_build_243.dm_build_233 = true
	return dm_build_243
}

func (dm_build_245 *dm_build_230) dm_build_244(dm_build_246 int16, dm_build_247 int64) {

	dm_build_248 := Dm_build_35

	if dm_build_245.dm_build_139.dm_build_1289.autoCommit {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 1)
	} else {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 0)
	}

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1140(dm_build_248, dm_build_246)

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 1)

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_248, dm_build_247)

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_248, dm_build_245.dm_build_142.cursorUpdateRow)

	if dm_build_245.dm_build_142.maxRows <= 0 || dm_build_245.dm_build_142.dmConn.dmConnector.enRsCache {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_248, INT64_MAX)
	} else {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_248, dm_build_245.dm_build_142.maxRows)
	}

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 1)

	if dm_build_245.dm_build_139.dm_build_1289.dmConnector.continueBatchOnError {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 1)
	} else {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 0)
	}

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 0)

	dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_248, 0)

	if dm_build_245.dm_build_142.queryTimeout == 0 {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1144(dm_build_248, -1)
	} else {
		dm_build_248 += dm_build_245.dm_build_139.dm_build_1288.Dm_build_1144(dm_build_248, dm_build_245.dm_build_142.queryTimeout)
	}
}

func (dm_build_250 *dm_build_230) dm_build_125() error {
	var dm_build_251 int16
	var dm_build_252 int64

	if dm_build_250.dm_build_232 != nil {
		dm_build_251 = int16(len(dm_build_250.dm_build_232))
	} else {
		dm_build_251 = 0
	}

	if dm_build_250.dm_build_231 != nil {
		dm_build_252 = int64(len(dm_build_250.dm_build_231))
	} else {
		dm_build_252 = 0
	}

	dm_build_250.dm_build_244(dm_build_251, dm_build_252)

	if dm_build_251 > 0 {
		err := dm_build_250.dm_build_253(dm_build_250.dm_build_232)
		if err != nil {
			return err
		}
		if dm_build_250.dm_build_231 != nil && len(dm_build_250.dm_build_231) > 0 {
			for _, paramObject := range dm_build_250.dm_build_231 {
				if err := dm_build_250.dm_build_256(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_254 *dm_build_230) dm_build_253(dm_build_255 []parameter) error {
	for _, param := range dm_build_255 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_254.dm_build_139.dm_build_1288.Dm_build_996(IO_TYPE_INOUT)
		} else {
			dm_build_254.dm_build_139.dm_build_1288.Dm_build_996(param.ioType)
		}

		dm_build_254.dm_build_139.dm_build_1288.Dm_build_1004(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_254.dm_build_139.dm_build_1288.Dm_build_1004(lprec)

		dm_build_254.dm_build_139.dm_build_1288.Dm_build_1004(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_254.dm_build_139.dm_build_1288)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_254.dm_build_139.dm_build_1288)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_254.dm_build_139.dm_build_1288)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_257 *dm_build_230) dm_build_256(dm_build_258 []interface{}) error {
	for i := 0; i < len(dm_build_257.dm_build_232); i++ {

		if dm_build_257.dm_build_232[i].colType == CURSOR {
			dm_build_257.dm_build_139.dm_build_1288.Dm_build_1000(ULINT_SIZE)
			dm_build_257.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_257.dm_build_232[i].cursorStmt.id)
			continue
		}

		if dm_build_257.dm_build_232[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_258[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_258[i].([]byte); ok {
				if len(dataBytes) > Dm_build_45 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_257.dm_build_139.dm_build_1288.Dm_build_1042(dataBytes)
			}
		case int:
			if dm_build_258[i] == ParamDataEnum_Null {
				dm_build_257.dm_build_139.dm_build_1288.Dm_build_1012(Dm_build_50)
			} else if dm_build_258[i] == ParamDataEnum_OFF_ROW {
				dm_build_257.dm_build_139.dm_build_1288.Dm_build_1000(0)
			}
		case lobCtl:
			dm_build_257.dm_build_139.dm_build_1288.Dm_build_1012(uint16(Dm_build_49))
			dm_build_257.dm_build_139.dm_build_1288.Dm_build_1032(dm_build_258[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_260 *dm_build_230) dm_build_129() (interface{}, error) {
	dm_build_261 := execInfo{}
	dm_build_262 := dm_build_260.dm_build_142.dmConn

	dm_build_263 := Dm_build_35

	dm_build_261.retSqlType = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_263)
	dm_build_263 += USINT_SIZE

	dm_build_264 := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_263)
	dm_build_263 += USINT_SIZE

	dm_build_261.updateCount = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1225(dm_build_263)
	dm_build_263 += DDWORD_SIZE

	dm_build_265 := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_263)
	dm_build_263 += USINT_SIZE

	dm_build_261.rsUpdatable = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1216(dm_build_263) != 0
	dm_build_263 += BYTE_SIZE

	dm_build_266 := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_263)
	dm_build_263 += ULINT_SIZE

	dm_build_261.printLen = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1222(dm_build_263)
	dm_build_263 += ULINT_SIZE

	var dm_build_267 int16 = -1
	if dm_build_261.retSqlType == Dm_build_100 || dm_build_261.retSqlType == Dm_build_101 {
		dm_build_261.rowid = 0

		dm_build_261.rsBdta = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1216(dm_build_263) == Dm_build_113
		dm_build_263 += BYTE_SIZE

		dm_build_267 = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_263)
		dm_build_263 += USINT_SIZE
		dm_build_263 += 5
	} else {
		dm_build_261.rowid = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1225(dm_build_263)
		dm_build_263 += DDWORD_SIZE
	}

	dm_build_261.execId = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1222(dm_build_263)
	dm_build_263 += ULINT_SIZE

	dm_build_261.rsCacheOffset = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1222(dm_build_263)
	dm_build_263 += ULINT_SIZE

	dm_build_268 := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1216(dm_build_263)
	dm_build_263 += BYTE_SIZE
	dm_build_269 := (dm_build_268 & 0x01) == 0x01
	dm_build_270 := (dm_build_268 & 0x02) == 0x02

	dm_build_262.TrxStatus = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1222(dm_build_263)
	dm_build_262.setTrxFinish(dm_build_262.TrxStatus)
	dm_build_263 += ULINT_SIZE

	if dm_build_261.printLen > 0 {
		bytes := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1099(int(dm_build_261.printLen))
		dm_build_261.printMsg = Dm_build_599.Dm_build_753(bytes, 0, len(bytes), dm_build_262.getServerEncoding(), dm_build_262)
	}

	if dm_build_265 > 0 {
		dm_build_261.outParamDatas = dm_build_260.dm_build_271(int(dm_build_265))
	}

	switch dm_build_261.retSqlType {
	case Dm_build_102:
		dm_build_262.dmConnector.localTimezone = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1075()
	case Dm_build_100:
		dm_build_261.hasResultSet = true
		if dm_build_264 > 0 {
			dm_build_260.dm_build_142.columns = dm_build_260.dm_build_280(int(dm_build_264), dm_build_261.rsBdta)
		}
		dm_build_260.dm_build_290(&dm_build_261, len(dm_build_260.dm_build_142.columns), int(dm_build_266), int(dm_build_267))
	case Dm_build_101:
		if dm_build_264 > 0 || dm_build_266 > 0 {
			dm_build_261.hasResultSet = true
		}
		if dm_build_264 > 0 {
			dm_build_260.dm_build_142.columns = dm_build_260.dm_build_280(int(dm_build_264), dm_build_261.rsBdta)
		}
		dm_build_260.dm_build_290(&dm_build_261, len(dm_build_260.dm_build_142.columns), int(dm_build_266), int(dm_build_267))
	case Dm_build_103:
		dm_build_262.IsoLevel = int32(dm_build_260.dm_build_139.dm_build_1288.Dm_build_1075())
		dm_build_262.ReadOnly = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1072() == 1
	case Dm_build_96:
		dm_build_262.Schema = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_93:
		dm_build_261.explain = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_262.getServerEncoding(), dm_build_262)

	case Dm_build_97, Dm_build_99, Dm_build_98:
		if dm_build_269 {

			counts := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1078()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1081()
			}
			dm_build_261.updateCounts = rowCounts
		}

		if dm_build_270 {
			rows := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1078()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1081()
			}
			dm_build_261.lastInsertId = lastInsertId

		} else if dm_build_261.updateCount == 1 {
			dm_build_261.lastInsertId = dm_build_261.rowid
		}

		if dm_build_260.dm_build_141 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_260.dm_build_296(dm_build_261.updateCounts)
		}
	case Dm_build_106:
		len := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
		dm_build_262.OracleDateFormat = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1115(int(len), dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_108:

		len := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
		dm_build_262.OracleTimestampFormat = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1115(int(len), dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_109:

		len := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
		dm_build_262.OracleTimestampTZFormat = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1115(int(len), dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_107:
		len := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
		dm_build_262.OracleTimeFormat = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1115(int(len), dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_110:
		len := dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
		dm_build_262.OracleTimeTZFormat = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1115(int(len), dm_build_262.getServerEncoding(), dm_build_262)
	case Dm_build_111:
		dm_build_262.OracleDateLanguage = dm_build_260.dm_build_139.dm_build_1288.Dm_build_1090()
	}

	return &dm_build_261, nil
}

func (dm_build_272 *dm_build_230) dm_build_271(dm_build_273 int) [][]byte {
	dm_build_274 := make([]int, dm_build_273)

	dm_build_275 := 0
	for i := 0; i < len(dm_build_272.dm_build_232); i++ {
		if dm_build_272.dm_build_232[i].ioType == IO_TYPE_INOUT || dm_build_272.dm_build_232[i].ioType == IO_TYPE_OUT {
			dm_build_274[dm_build_275] = i
			dm_build_275++
		}
	}

	dm_build_276 := make([][]byte, len(dm_build_272.dm_build_232))
	var dm_build_277 int32
	var dm_build_278 bool
	var dm_build_279 []byte = nil
	for i := 0; i < dm_build_273; i++ {
		dm_build_278 = false
		dm_build_277 = int32(dm_build_272.dm_build_139.dm_build_1288.Dm_build_1093())

		if dm_build_277 == int32(Dm_build_50) {
			dm_build_277 = 0
			dm_build_278 = true
		} else if dm_build_277 == int32(Dm_build_51) {
			dm_build_277 = dm_build_272.dm_build_139.dm_build_1288.Dm_build_1078()
		}

		if dm_build_278 {
			dm_build_276[dm_build_274[i]] = nil
		} else {
			dm_build_279 = dm_build_272.dm_build_139.dm_build_1288.Dm_build_1099(int(dm_build_277))
			dm_build_276[dm_build_274[i]] = dm_build_279
		}
	}

	return dm_build_276
}

func (dm_build_281 *dm_build_230) dm_build_280(dm_build_282 int, dm_build_283 bool) []column {
	dm_build_284 := dm_build_281.dm_build_139.dm_build_1289.getServerEncoding()
	var dm_build_285, dm_build_286, dm_build_287, dm_build_288 int16
	dm_build_289 := make([]column, dm_build_282)
	for i := 0; i < dm_build_282; i++ {
		dm_build_289[i].InitColumn()

		dm_build_289[i].colType = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_289[i].prec = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_289[i].scale = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_289[i].nullable = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1078() != 0

		itemFlag := dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()
		dm_build_289[i].lob = int(itemFlag)&Dm_build_227 != 0
		dm_build_289[i].identity = int(itemFlag)&Dm_build_226 != 0
		dm_build_289[i].readonly = int(itemFlag)&Dm_build_228 != 0

		dm_build_281.dm_build_139.dm_build_1288.Dm_build_978(4, false, true)

		dm_build_281.dm_build_139.dm_build_1288.Dm_build_978(2, false, true)

		dm_build_285 = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_286 = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_287 = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_288 = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()
		dm_build_289[i].name = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_285), dm_build_284, dm_build_281.dm_build_139.dm_build_1289)
		dm_build_289[i].typeName = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_286), dm_build_284, dm_build_281.dm_build_139.dm_build_1289)
		dm_build_289[i].tableName = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_287), dm_build_284, dm_build_281.dm_build_139.dm_build_1289)
		dm_build_289[i].schemaName = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_288), dm_build_284, dm_build_281.dm_build_139.dm_build_1289)

		if dm_build_281.dm_build_142.readBaseColName {
			dm_build_289[i].baseName = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1128(dm_build_284, dm_build_281.dm_build_139.dm_build_1289)
		}

		if dm_build_289[i].lob {
			dm_build_289[i].lobTabId = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1078()
			dm_build_289[i].lobColId = dm_build_281.dm_build_139.dm_build_1288.Dm_build_1075()
		}

	}

	for i := 0; i < dm_build_282; i++ {

		if isComplexType(int(dm_build_289[i].colType), int(dm_build_289[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_281.dm_build_139.dm_build_1289)
			strDesc.unpack(dm_build_281.dm_build_139.dm_build_1288)
			dm_build_289[i].typeDescriptor = strDesc
		}
	}

	return dm_build_289
}

func (dm_build_291 *dm_build_230) dm_build_290(dm_build_292 *execInfo, dm_build_293 int, dm_build_294 int, dm_build_295 int) {
	if dm_build_294 > 0 {
		startOffset := dm_build_291.dm_build_139.dm_build_1288.Dm_build_973()
		if dm_build_292.rsBdta {
			dm_build_292.rsDatas = dm_build_291.dm_build_309(dm_build_291.dm_build_142.columns, dm_build_295)
		} else {
			datas := make([][][]byte, dm_build_294)

			for i := 0; i < dm_build_294; i++ {

				datas[i] = make([][]byte, dm_build_293+1)

				dm_build_291.dm_build_139.dm_build_1288.Dm_build_978(2, false, true)

				datas[i][0] = dm_build_291.dm_build_139.dm_build_1288.Dm_build_1099(LINT64_SIZE)

				dm_build_291.dm_build_139.dm_build_1288.Dm_build_978(2*dm_build_293, false, true)

				for j := 1; j < dm_build_293+1; j++ {

					colLen := dm_build_291.dm_build_139.dm_build_1288.Dm_build_1093()
					if colLen == Dm_build_54 {
						datas[i][j] = nil
					} else if colLen != Dm_build_55 {
						datas[i][j] = dm_build_291.dm_build_139.dm_build_1288.Dm_build_1099(int(colLen))
					} else {
						datas[i][j] = dm_build_291.dm_build_139.dm_build_1288.Dm_build_1103()
					}
				}
			}

			dm_build_292.rsDatas = datas
		}
		dm_build_292.rsSizeof = dm_build_291.dm_build_139.dm_build_1288.Dm_build_973() - startOffset
	}

	if dm_build_292.rsCacheOffset > 0 {
		tbCount := dm_build_291.dm_build_139.dm_build_1288.Dm_build_1075()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_291.dm_build_139.dm_build_1288.Dm_build_1078()
			tss[i] = dm_build_291.dm_build_139.dm_build_1288.Dm_build_1081()
		}

		dm_build_292.tbIds = ids[:]
		dm_build_292.tbTss = tss[:]
	}
}

func (dm_build_297 *dm_build_230) dm_build_296(dm_build_298 []int64) error {

	dm_build_297.dm_build_139.dm_build_1288.Dm_build_978(4, false, true)

	dm_build_299 := dm_build_297.dm_build_139.dm_build_1288.Dm_build_1078()

	dm_build_300 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_299); i++ {
		irow := dm_build_297.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_298[irow] = -3

		code := dm_build_297.dm_build_139.dm_build_1288.Dm_build_1078()

		errStr := dm_build_297.dm_build_139.dm_build_1288.Dm_build_1128(dm_build_297.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_297.dm_build_139.dm_build_1289)

		dm_build_300 = append(dm_build_300, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_300) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_300 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_301 = 0

	Dm_build_302 = Dm_build_301 + ULINT_SIZE

	Dm_build_303 = Dm_build_302 + USINT_SIZE

	Dm_build_304 = Dm_build_303 + ULINT_SIZE

	Dm_build_305 = Dm_build_304 + ULINT_SIZE

	Dm_build_306 = Dm_build_305 + BYTE_SIZE

	Dm_build_307 = -2

	Dm_build_308 = -3
)

func (dm_build_310 *dm_build_230) dm_build_309(dm_build_311 []column, dm_build_312 int) [][][]byte {

	dm_build_313 := dm_build_310.dm_build_139.dm_build_1288.Dm_build_1096()

	dm_build_314 := dm_build_310.dm_build_139.dm_build_1288.Dm_build_1093()

	var dm_build_315 bool

	if dm_build_312 >= 0 && int(dm_build_314) == len(dm_build_311)+1 {
		dm_build_315 = true
	} else {
		dm_build_315 = false
	}

	dm_build_310.dm_build_139.dm_build_1288.Dm_build_978(ULINT_SIZE, false, true)

	dm_build_310.dm_build_139.dm_build_1288.Dm_build_978(ULINT_SIZE, false, true)

	dm_build_310.dm_build_139.dm_build_1288.Dm_build_978(BYTE_SIZE, false, true)

	dm_build_316 := make([]uint16, dm_build_314)
	for icol := 0; icol < int(dm_build_314); icol++ {
		dm_build_316[icol] = dm_build_310.dm_build_139.dm_build_1288.Dm_build_1093()
	}

	dm_build_317 := make([]uint32, dm_build_314)
	dm_build_318 := make([][][]byte, dm_build_313)

	for i := uint32(0); i < dm_build_313; i++ {
		dm_build_318[i] = make([][]byte, len(dm_build_311)+1)
	}

	for icol := 0; icol < int(dm_build_314); icol++ {
		dm_build_317[icol] = dm_build_310.dm_build_139.dm_build_1288.Dm_build_1096()
	}

	for icol := 0; icol < int(dm_build_314); icol++ {

		dataCol := icol + 1
		if dm_build_315 && icol == dm_build_312 {
			dataCol = 0
		} else if dm_build_315 && icol > dm_build_312 {
			dataCol = icol
		}

		allNotNull := dm_build_310.dm_build_139.dm_build_1288.Dm_build_1078() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_313)
			for irow := uint32(0); irow < dm_build_313; irow++ {
				isNull[irow] = dm_build_310.dm_build_139.dm_build_1288.Dm_build_1072() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_313; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_318[irow][dataCol] = dm_build_310.dm_build_319(int(dm_build_316[icol]))
			}
		}
	}

	if !dm_build_315 && dm_build_312 >= 0 {
		for irow := uint32(0); irow < dm_build_313; irow++ {
			dm_build_318[irow][0] = dm_build_318[irow][dm_build_312+1]
		}
	}

	return dm_build_318
}

func (dm_build_320 *dm_build_230) dm_build_319(dm_build_321 int) []byte {

	dm_build_322 := dm_build_320.dm_build_325(dm_build_321)

	dm_build_323 := int32(0)
	if dm_build_322 == Dm_build_307 {
		dm_build_323 = dm_build_320.dm_build_139.dm_build_1288.Dm_build_1078()
		dm_build_322 = int(dm_build_320.dm_build_139.dm_build_1288.Dm_build_1078())
	} else if dm_build_322 == Dm_build_308 {
		dm_build_322 = int(dm_build_320.dm_build_139.dm_build_1288.Dm_build_1078())
	}

	dm_build_324 := dm_build_320.dm_build_139.dm_build_1288.Dm_build_1099(dm_build_322 + int(dm_build_323))
	if dm_build_323 == 0 {
		return dm_build_324
	}

	for i := dm_build_322; i < len(dm_build_324); i++ {
		dm_build_324[i] = ' '
	}
	return dm_build_324
}

func (dm_build_326 *dm_build_230) dm_build_325(dm_build_327 int) int {

	dm_build_328 := 0
	switch dm_build_327 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_328 = 4

	case BIGINT:

		dm_build_328 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_328 = Dm_build_307

	case DECIMAL:
		dm_build_328 = Dm_build_308

	case REAL:
		dm_build_328 = 4

	case DOUBLE:
		dm_build_328 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_328 = 12

	case INTERVAL_YM:
		dm_build_328 = 12

	case INTERVAL_DT:
		dm_build_328 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_328
}

const (
	Dm_build_329 = Dm_build_35

	Dm_build_330 = Dm_build_329 + DDWORD_SIZE

	Dm_build_331 = Dm_build_330 + LINT64_SIZE

	Dm_build_332 = Dm_build_331 + USINT_SIZE

	Dm_build_333 = Dm_build_35

	Dm_build_334 = Dm_build_333 + DDWORD_SIZE
)

type dm_build_335 struct {
	dm_build_230
	dm_build_336 *innerRows
	dm_build_337 int64
	dm_build_338 int64
}

func dm_build_339(dm_build_340 *dm_build_1285, dm_build_341 *innerRows, dm_build_342 int64, dm_build_343 int64) *dm_build_335 {
	dm_build_344 := new(dm_build_335)
	dm_build_344.dm_build_147(dm_build_340, Dm_build_13, dm_build_341.dmStmt)
	dm_build_344.dm_build_336 = dm_build_341
	dm_build_344.dm_build_337 = dm_build_342
	dm_build_344.dm_build_338 = dm_build_343
	return dm_build_344
}

func (dm_build_346 *dm_build_335) dm_build_125() error {

	dm_build_346.dm_build_139.dm_build_1288.Dm_build_1148(Dm_build_329, dm_build_346.dm_build_337)

	dm_build_346.dm_build_139.dm_build_1288.Dm_build_1148(Dm_build_330, dm_build_346.dm_build_338)

	dm_build_346.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_331, dm_build_346.dm_build_336.id)

	dm_build_347 := dm_build_346.dm_build_336.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_348 int32
	if dm_build_346.dm_build_336.sizeOfRow != 0 && dm_build_346.dm_build_336.fetchSize != 0 {
		if dm_build_346.dm_build_336.sizeOfRow*dm_build_346.dm_build_336.fetchSize > int(INT32_MAX) {
			dm_build_348 = INT32_MAX
		} else {
			dm_build_348 = int32(dm_build_346.dm_build_336.sizeOfRow * dm_build_346.dm_build_336.fetchSize)
		}

		if dm_build_348 < Dm_build_66 {
			dm_build_347 = int(Dm_build_66)
		} else if dm_build_348 > Dm_build_67 {
			dm_build_347 = int(Dm_build_67)
		} else {
			dm_build_347 = int(dm_build_348)
		}

		dm_build_346.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_332, int32(dm_build_347))
	}

	return nil
}

func (dm_build_350 *dm_build_335) dm_build_129() (interface{}, error) {
	dm_build_351 := execInfo{}
	dm_build_351.rsBdta = dm_build_350.dm_build_336.isBdta

	dm_build_351.updateCount = dm_build_350.dm_build_139.dm_build_1288.Dm_build_1225(Dm_build_333)
	dm_build_352 := dm_build_350.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_334)

	dm_build_350.dm_build_290(&dm_build_351, len(dm_build_350.dm_build_336.columns), int(dm_build_352), -1)

	return &dm_build_351, nil
}

type dm_build_353 struct {
	dm_build_138
	dm_build_354 *lob
	dm_build_355 int
	dm_build_356 int
}

func dm_build_357(dm_build_358 *dm_build_1285, dm_build_359 *lob, dm_build_360 int, dm_build_361 int) *dm_build_353 {
	dm_build_362 := new(dm_build_353)
	dm_build_362.dm_build_143(dm_build_358, Dm_build_26)
	dm_build_362.dm_build_354 = dm_build_359
	dm_build_362.dm_build_355 = dm_build_360
	dm_build_362.dm_build_356 = dm_build_361
	return dm_build_362
}

func (dm_build_364 *dm_build_353) dm_build_125() error {

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_996(byte(dm_build_364.dm_build_354.lobFlag))

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_364.dm_build_354.tabId)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.colId)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_364.dm_build_354.blobId))

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.groupId)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.fileId)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_364.dm_build_354.pageNo)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.curFileId)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_364.dm_build_354.curPageNo)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_364.dm_build_354.totalOffset)

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(int32(dm_build_364.dm_build_355))

	dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(int32(dm_build_364.dm_build_356))

	if dm_build_364.dm_build_139.dm_build_1289.NewLobFlag {
		dm_build_364.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_364.dm_build_354.rowId))
		dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.exGroupId)
		dm_build_364.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_364.dm_build_354.exFileId)
		dm_build_364.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_364.dm_build_354.exPageNo)
	}

	return nil
}

func (dm_build_366 *dm_build_353) dm_build_129() (interface{}, error) {

	dm_build_366.dm_build_354.readOver = dm_build_366.dm_build_139.dm_build_1288.Dm_build_1072() == 1
	var dm_build_367 = dm_build_366.dm_build_139.dm_build_1288.Dm_build_1096()
	if dm_build_367 <= 0 {
		return []byte(nil), nil
	}
	dm_build_366.dm_build_354.curFileId = dm_build_366.dm_build_139.dm_build_1288.Dm_build_1075()
	dm_build_366.dm_build_354.curPageNo = dm_build_366.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_366.dm_build_354.totalOffset = dm_build_366.dm_build_139.dm_build_1288.Dm_build_1078()

	return dm_build_366.dm_build_139.dm_build_1288.Dm_build_1109(int(dm_build_367)), nil
}

type dm_build_368 struct {
	dm_build_138
	dm_build_369 *lob
}

func dm_build_370(dm_build_371 *dm_build_1285, dm_build_372 *lob) *dm_build_368 {
	dm_build_373 := new(dm_build_368)
	dm_build_373.dm_build_143(dm_build_371, Dm_build_23)
	dm_build_373.dm_build_369 = dm_build_372
	return dm_build_373
}

func (dm_build_375 *dm_build_368) dm_build_125() error {

	dm_build_375.dm_build_139.dm_build_1288.Dm_build_996(byte(dm_build_375.dm_build_369.lobFlag))

	dm_build_375.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_375.dm_build_369.blobId))

	dm_build_375.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_375.dm_build_369.groupId)

	dm_build_375.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_375.dm_build_369.fileId)

	dm_build_375.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_375.dm_build_369.pageNo)

	if dm_build_375.dm_build_139.dm_build_1289.NewLobFlag {
		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_375.dm_build_369.tabId)
		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_375.dm_build_369.colId)
		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_375.dm_build_369.rowId))

		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_375.dm_build_369.exGroupId)
		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_375.dm_build_369.exFileId)
		dm_build_375.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_375.dm_build_369.exPageNo)
	}

	return nil
}

func (dm_build_377 *dm_build_368) dm_build_129() (interface{}, error) {

	return int64(dm_build_377.dm_build_139.dm_build_1288.Dm_build_1078()), nil
}

type dm_build_378 struct {
	dm_build_138
	dm_build_379 *lob
	dm_build_380 int
}

func dm_build_381(dm_build_382 *dm_build_1285, dm_build_383 *lob, dm_build_384 int) *dm_build_378 {
	dm_build_385 := new(dm_build_378)
	dm_build_385.dm_build_143(dm_build_382, Dm_build_25)
	dm_build_385.dm_build_379 = dm_build_383
	dm_build_385.dm_build_380 = dm_build_384
	return dm_build_385
}

func (dm_build_387 *dm_build_378) dm_build_125() error {

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_996(byte(dm_build_387.dm_build_379.lobFlag))

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_387.dm_build_379.blobId))

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_387.dm_build_379.groupId)

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_387.dm_build_379.fileId)

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_387.dm_build_379.pageNo)

	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_387.dm_build_379.tabId)
	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_387.dm_build_379.colId)
	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_387.dm_build_379.rowId))
	dm_build_387.dm_build_139.dm_build_1288.Dm_build_1032(Dm_build_599.Dm_build_798(uint32(dm_build_387.dm_build_380)))

	if dm_build_387.dm_build_139.dm_build_1289.NewLobFlag {
		dm_build_387.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_387.dm_build_379.exGroupId)
		dm_build_387.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_387.dm_build_379.exFileId)
		dm_build_387.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_387.dm_build_379.exPageNo)
	}
	return nil
}

func (dm_build_389 *dm_build_378) dm_build_129() (interface{}, error) {

	dm_build_390 := dm_build_389.dm_build_139.dm_build_1288.Dm_build_1096()
	dm_build_389.dm_build_379.blobId = dm_build_389.dm_build_139.dm_build_1288.Dm_build_1081()
	dm_build_389.dm_build_379.resetCurrentInfo()
	return int64(dm_build_390), nil
}

const (
	Dm_build_391 = Dm_build_35

	Dm_build_392 = Dm_build_391 + ULINT_SIZE

	Dm_build_393 = Dm_build_392 + ULINT_SIZE

	Dm_build_394 = Dm_build_393 + ULINT_SIZE

	Dm_build_395 = Dm_build_394 + BYTE_SIZE

	Dm_build_396 = Dm_build_395 + USINT_SIZE

	Dm_build_397 = Dm_build_396 + ULINT_SIZE

	Dm_build_398 = Dm_build_397 + BYTE_SIZE

	Dm_build_399 = Dm_build_398 + BYTE_SIZE

	Dm_build_400 = Dm_build_399 + BYTE_SIZE

	Dm_build_401 = Dm_build_35

	Dm_build_402 = Dm_build_401 + ULINT_SIZE

	Dm_build_403 = Dm_build_402 + ULINT_SIZE

	Dm_build_404 = Dm_build_403 + BYTE_SIZE

	Dm_build_405 = Dm_build_404 + ULINT_SIZE

	Dm_build_406 = Dm_build_405 + BYTE_SIZE

	Dm_build_407 = Dm_build_406 + BYTE_SIZE

	Dm_build_408 = Dm_build_407 + USINT_SIZE

	Dm_build_409 = Dm_build_408 + USINT_SIZE

	Dm_build_410 = Dm_build_409 + BYTE_SIZE

	Dm_build_411 = Dm_build_410 + USINT_SIZE

	Dm_build_412 = Dm_build_411 + BYTE_SIZE

	Dm_build_413 = Dm_build_412 + BYTE_SIZE

	Dm_build_414 = Dm_build_413 + ULINT_SIZE
)

type dm_build_415 struct {
	dm_build_138

	dm_build_416 *DmConnection

	dm_build_417 bool
}

func dm_build_418(dm_build_419 *dm_build_1285) *dm_build_415 {
	dm_build_420 := new(dm_build_415)
	dm_build_420.dm_build_143(dm_build_419, Dm_build_7)
	dm_build_420.dm_build_416 = dm_build_419.dm_build_1289
	return dm_build_420
}

func (dm_build_422 *dm_build_415) dm_build_125() error {

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_391, Dm_build_46)

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_392, g2dbIsoLevel(dm_build_422.dm_build_416.IsoLevel))
	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_393, int32(Locale))
	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_395, dm_build_422.dm_build_416.dmConnector.localTimezone)

	if dm_build_422.dm_build_416.ReadOnly {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_394, Dm_build_69)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_394, Dm_build_68)
	}

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_396, int32(dm_build_422.dm_build_416.dmConnector.sessionTimeout))

	if dm_build_422.dm_build_416.dmConnector.mppLocal {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_397, 1)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_397, 0)
	}

	if dm_build_422.dm_build_416.dmConnector.rwSeparate {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_398, 1)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_398, 0)
	}

	if dm_build_422.dm_build_416.NewLobFlag {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_399, 1)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_399, 0)
	}

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_400, dm_build_422.dm_build_416.dmConnector.osAuthType)

	dm_build_423 := dm_build_422.dm_build_416.getServerEncoding()

	if dm_build_422.dm_build_139.dm_build_1294 != "" {

	}

	dm_build_424 := Dm_build_599.Dm_build_809(dm_build_422.dm_build_416.dmConnector.user, dm_build_423, dm_build_422.dm_build_139.dm_build_1289)
	dm_build_425 := Dm_build_599.Dm_build_809(dm_build_422.dm_build_416.dmConnector.password, dm_build_423, dm_build_422.dm_build_139.dm_build_1289)
	if len(dm_build_424) > Dm_build_43 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_425) > Dm_build_43 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_422.dm_build_139.dm_build_1291 && dm_build_422.dm_build_416.dmConnector.loginCertificate != "" {

	} else if dm_build_422.dm_build_139.dm_build_1291 {
		dm_build_424 = dm_build_422.dm_build_139.dm_build_1290.Encrypt(dm_build_424, false)
		dm_build_425 = dm_build_422.dm_build_139.dm_build_1290.Encrypt(dm_build_425, false)
	}

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1036(dm_build_424)
	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1036(dm_build_425)

	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1048(dm_build_422.dm_build_416.dmConnector.appName, dm_build_423, dm_build_422.dm_build_139.dm_build_1289)
	dm_build_422.dm_build_139.dm_build_1288.Dm_build_1048(dm_build_422.dm_build_416.dmConnector.osName, dm_build_423, dm_build_422.dm_build_139.dm_build_1289)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1048(hostName, dm_build_423, dm_build_422.dm_build_139.dm_build_1289)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_1048("", dm_build_423, dm_build_422.dm_build_139.dm_build_1289)
	}

	if dm_build_422.dm_build_416.dmConnector.rwStandby {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_996(1)
	} else {
		dm_build_422.dm_build_139.dm_build_1288.Dm_build_996(0)
	}

	return nil
}

func (dm_build_427 *dm_build_415) dm_build_129() (interface{}, error) {

	dm_build_427.dm_build_416.MaxRowSize = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_401)
	dm_build_427.dm_build_416.DDLAutoCommit = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_403) == 1
	dm_build_427.dm_build_416.IsoLevel = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_404)
	dm_build_427.dm_build_416.dmConnector.caseSensitive = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_405) == 1
	dm_build_427.dm_build_416.BackslashEscape = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_406) == 1
	dm_build_427.dm_build_416.SvrStat = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1219(Dm_build_408)
	dm_build_427.dm_build_416.SvrMode = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1219(Dm_build_407)
	dm_build_427.dm_build_416.ConstParaOpt = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_409) == 1
	dm_build_427.dm_build_416.DbTimezone = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1219(Dm_build_410)
	dm_build_427.dm_build_416.NewLobFlag = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_412) == 1

	if dm_build_427.dm_build_416.dmConnector.bufPrefetch == 0 {
		dm_build_427.dm_build_416.dmConnector.bufPrefetch = int(dm_build_427.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_413))
	}

	dm_build_427.dm_build_416.LifeTimeRemainder = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1219(Dm_build_414)

	dm_build_428 := dm_build_427.dm_build_416.getServerEncoding()

	dm_build_427.dm_build_416.InstanceName = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
	dm_build_427.dm_build_416.Schema = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
	dm_build_427.dm_build_416.LastLoginIP = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
	dm_build_427.dm_build_416.LastLoginTime = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
	dm_build_427.dm_build_416.FailedAttempts = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_427.dm_build_416.LoginWarningID = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_427.dm_build_416.GraceTimeRemainder = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_427.dm_build_416.Guid = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
	dm_build_427.dm_build_416.DbName = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)

	if dm_build_427.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_411) == 1 {
		dm_build_427.dm_build_416.StandbyHost = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_428, dm_build_427.dm_build_139.dm_build_1289)
		dm_build_427.dm_build_416.StandbyPort = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1078()
		dm_build_427.dm_build_416.StandbyCount = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1093()
	}

	if dm_build_427.dm_build_139.dm_build_1288.Dm_build_975(false) > 0 {
		dm_build_427.dm_build_416.SessionID = dm_build_427.dm_build_139.dm_build_1288.Dm_build_1081()
	}

	if dm_build_427.dm_build_139.dm_build_1288.Dm_build_975(false) > 0 {
		if dm_build_427.dm_build_139.dm_build_1288.Dm_build_1072() == 1 {

			dm_build_427.dm_build_416.OracleDateFormat = "DD-MON-YY"

			dm_build_427.dm_build_416.OracleTimestampFormat = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_427.dm_build_416.OracleTimestampTZFormat = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_427.dm_build_416.OracleTimeFormat = "HH12.MI.SS.FF6 AM"

			dm_build_427.dm_build_416.OracleTimeTZFormat = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	return nil, nil
}

const (
	Dm_build_429 = Dm_build_35
)

type dm_build_430 struct {
	dm_build_230
	dm_build_431 int16
}

func dm_build_432(dm_build_433 *dm_build_1285, dm_build_434 *DmStatement, dm_build_435 int16) *dm_build_430 {
	dm_build_436 := new(dm_build_430)
	dm_build_436.dm_build_147(dm_build_433, Dm_build_27, dm_build_434)
	dm_build_436.dm_build_431 = dm_build_435
	return dm_build_436
}

func (dm_build_438 *dm_build_430) dm_build_125() error {
	dm_build_438.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_429, dm_build_438.dm_build_431)
	return nil
}

func (dm_build_440 *dm_build_430) dm_build_129() (interface{}, error) {
	return dm_build_440.dm_build_230.dm_build_129()
}

const (
	Dm_build_441 = Dm_build_35
)

type dm_build_442 struct {
	dm_build_230
	dm_build_443 []parameter
}

func dm_build_444(dm_build_445 *dm_build_1285, dm_build_446 *DmStatement, dm_build_447 []parameter) *dm_build_442 {
	dm_build_448 := new(dm_build_442)
	dm_build_448.dm_build_147(dm_build_445, Dm_build_31, dm_build_446)
	dm_build_448.dm_build_443 = dm_build_447
	return dm_build_448
}

func (dm_build_450 *dm_build_442) dm_build_125() error {

	if dm_build_450.dm_build_443 == nil {
		dm_build_450.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_441, 0)
	} else {
		dm_build_450.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_441, int16(len(dm_build_450.dm_build_443)))
	}

	return dm_build_450.dm_build_253(dm_build_450.dm_build_443)
}

type dm_build_451 struct {
	dm_build_230
	dm_build_452 bool
	dm_build_453 int16
}

func dm_build_454(dm_build_455 *dm_build_1285, dm_build_456 *DmStatement, dm_build_457 bool, dm_build_458 int16) *dm_build_451 {
	dm_build_459 := new(dm_build_451)
	dm_build_459.dm_build_147(dm_build_455, Dm_build_11, dm_build_456)
	dm_build_459.dm_build_452 = dm_build_457
	dm_build_459.dm_build_453 = dm_build_458
	return dm_build_459
}

func (dm_build_461 *dm_build_451) dm_build_125() error {

	dm_build_462 := Dm_build_35

	if dm_build_461.dm_build_139.dm_build_1289.autoCommit {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 1)
	} else {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)
	}

	if dm_build_461.dm_build_452 {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 1)
	} else {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)
	}

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 1)

	if dm_build_461.dm_build_139.dm_build_1289.CompatibleOracle() {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)
	} else {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 1)
	}

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1140(dm_build_462, dm_build_461.dm_build_453)

	if dm_build_461.dm_build_142.maxRows <= 0 || dm_build_461.dm_build_139.dm_build_1289.dmConnector.enRsCache {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_462, INT64_MAX)
	} else {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1148(dm_build_462, dm_build_461.dm_build_142.maxRows)
	}

	if dm_build_461.dm_build_139.dm_build_1289.dmConnector.isBdtaRS {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, Dm_build_113)
	} else {
		dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, Dm_build_112)
	}

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1140(dm_build_462, 0)

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 1)

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1136(dm_build_462, 0)

	dm_build_462 += dm_build_461.dm_build_139.dm_build_1288.Dm_build_1144(dm_build_462, dm_build_461.dm_build_142.queryTimeout)

	dm_build_461.dm_build_139.dm_build_1288.Dm_build_1066(dm_build_461.dm_build_142.nativeSql, dm_build_461.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_461.dm_build_139.dm_build_1289)

	return nil
}

func (dm_build_464 *dm_build_451) dm_build_129() (interface{}, error) {

	if dm_build_464.dm_build_452 {
		return dm_build_464.dm_build_230.dm_build_129()
	}

	dm_build_465 := NewExceInfo()
	dm_build_466 := Dm_build_35

	dm_build_465.retSqlType = dm_build_464.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_466)
	dm_build_466 += USINT_SIZE

	dm_build_467 := dm_build_464.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_466)
	dm_build_466 += USINT_SIZE

	dm_build_468 := dm_build_464.dm_build_139.dm_build_1288.Dm_build_1219(dm_build_466)
	dm_build_466 += USINT_SIZE

	dm_build_464.dm_build_139.dm_build_1288.Dm_build_1225(dm_build_466)
	dm_build_466 += DDWORD_SIZE

	dm_build_464.dm_build_139.dm_build_1289.TrxStatus = dm_build_464.dm_build_139.dm_build_1288.Dm_build_1222(dm_build_466)
	dm_build_466 += ULINT_SIZE

	if dm_build_467 > 0 {
		dm_build_464.dm_build_142.params = dm_build_464.dm_build_469(int(dm_build_467))
		dm_build_464.dm_build_142.paramCount = dm_build_467
	} else {
		dm_build_464.dm_build_142.params = make([]parameter, 0)
		dm_build_464.dm_build_142.paramCount = 0
	}

	if dm_build_468 > 0 {
		dm_build_464.dm_build_142.columns = dm_build_464.dm_build_280(int(dm_build_468), dm_build_465.rsBdta)
	} else {
		dm_build_464.dm_build_142.columns = make([]column, 0)
	}

	return dm_build_465, nil
}

func (dm_build_470 *dm_build_451) dm_build_469(dm_build_471 int) []parameter {

	var dm_build_472, dm_build_473, dm_build_474, dm_build_475 int16

	dm_build_476 := make([]parameter, dm_build_471)
	for i := 0; i < dm_build_471; i++ {

		dm_build_476[i].InitParameter()

		dm_build_476[i].colType = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_476[i].prec = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_476[i].scale = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_476[i].nullable = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078() != 0

		itemFlag := dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()

		if int(itemFlag)&Dm_build_229 != 0 {
			dm_build_476[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_476[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_476[i].lob = int(itemFlag)&Dm_build_227 != 0

		dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078()

		dm_build_476[i].ioType = byte(dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075())

		dm_build_472 = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_473 = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_474 = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()

		dm_build_475 = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()
		dm_build_476[i].name = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_472), dm_build_470.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_470.dm_build_139.dm_build_1289)
		dm_build_476[i].typeName = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_473), dm_build_470.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_470.dm_build_139.dm_build_1289)
		dm_build_476[i].tableName = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_474), dm_build_470.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_470.dm_build_139.dm_build_1289)
		dm_build_476[i].schemaName = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1115(int(dm_build_475), dm_build_470.dm_build_139.dm_build_1289.getServerEncoding(), dm_build_470.dm_build_139.dm_build_1289)

		if dm_build_476[i].lob {
			dm_build_476[i].lobTabId = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1078()
			dm_build_476[i].lobColId = dm_build_470.dm_build_139.dm_build_1288.Dm_build_1075()
		}
	}

	for i := 0; i < dm_build_471; i++ {

		if isComplexType(int(dm_build_476[i].colType), int(dm_build_476[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_470.dm_build_139.dm_build_1289)
			strDesc.unpack(dm_build_470.dm_build_139.dm_build_1288)
			dm_build_476[i].typeDescriptor = strDesc
		}
	}

	return dm_build_476
}

const (
	Dm_build_477 = Dm_build_35
)

type dm_build_478 struct {
	dm_build_138
	dm_build_479 int16
	dm_build_480 *Dm_build_875
	dm_build_481 int32
}

func dm_build_482(dm_build_483 *dm_build_1285, dm_build_484 *DmStatement, dm_build_485 int16, dm_build_486 *Dm_build_875, dm_build_487 int32) *dm_build_478 {
	dm_build_488 := new(dm_build_478)
	dm_build_488.dm_build_147(dm_build_483, Dm_build_17, dm_build_484)
	dm_build_488.dm_build_479 = dm_build_485
	dm_build_488.dm_build_480 = dm_build_486
	dm_build_488.dm_build_481 = dm_build_487
	return dm_build_488
}

func (dm_build_490 *dm_build_478) dm_build_125() error {
	dm_build_490.dm_build_139.dm_build_1288.Dm_build_1140(Dm_build_477, dm_build_490.dm_build_479)

	dm_build_490.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_490.dm_build_481)

	if dm_build_490.dm_build_139.dm_build_1289.NewLobFlag {
		dm_build_490.dm_build_139.dm_build_1288.Dm_build_1004(-1)
	}
	dm_build_490.dm_build_480.Dm_build_882(dm_build_490.dm_build_139.dm_build_1288, int(dm_build_490.dm_build_481))
	return nil
}

type dm_build_491 struct {
	dm_build_138
}

func dm_build_492(dm_build_493 *dm_build_1285) *dm_build_491 {
	dm_build_494 := new(dm_build_491)
	dm_build_494.dm_build_143(dm_build_493, Dm_build_15)
	return dm_build_494
}

type dm_build_495 struct {
	dm_build_138
	dm_build_496 int32
}

func dm_build_497(dm_build_498 *dm_build_1285, dm_build_499 int32) *dm_build_495 {
	dm_build_500 := new(dm_build_495)
	dm_build_500.dm_build_143(dm_build_498, Dm_build_28)
	dm_build_500.dm_build_496 = dm_build_499
	return dm_build_500
}

func (dm_build_502 *dm_build_495) dm_build_125() error {

	dm_build_503 := Dm_build_35
	dm_build_503 += dm_build_502.dm_build_139.dm_build_1288.Dm_build_1144(dm_build_503, g2dbIsoLevel(dm_build_502.dm_build_496))
	return nil
}

type dm_build_504 struct {
	dm_build_138
	dm_build_505 *lob
	dm_build_506 byte
	dm_build_507 int
	dm_build_508 []byte
	dm_build_509 int
	dm_build_510 int
}

func dm_build_511(dm_build_512 *dm_build_1285, dm_build_513 *lob, dm_build_514 byte, dm_build_515 int, dm_build_516 []byte,
	dm_build_517 int, dm_build_518 int) *dm_build_504 {
	dm_build_519 := new(dm_build_504)
	dm_build_519.dm_build_143(dm_build_512, Dm_build_24)
	dm_build_519.dm_build_505 = dm_build_513
	dm_build_519.dm_build_506 = dm_build_514
	dm_build_519.dm_build_507 = dm_build_515
	dm_build_519.dm_build_508 = dm_build_516
	dm_build_519.dm_build_509 = dm_build_517
	dm_build_519.dm_build_510 = dm_build_518
	return dm_build_519
}

func (dm_build_521 *dm_build_504) dm_build_125() error {

	dm_build_521.dm_build_139.dm_build_1288.Dm_build_996(byte(dm_build_521.dm_build_505.lobFlag))
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_996(dm_build_521.dm_build_506)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_521.dm_build_505.blobId))
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.groupId)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.fileId)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_521.dm_build_505.pageNo)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.curFileId)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_521.dm_build_505.curPageNo)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_521.dm_build_505.totalOffset)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_521.dm_build_505.tabId)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.colId)
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1020(uint64(dm_build_521.dm_build_505.rowId))

	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(int32(dm_build_521.dm_build_507))
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(int32(dm_build_521.dm_build_510))
	dm_build_521.dm_build_139.dm_build_1288.Dm_build_1032(dm_build_521.dm_build_508)

	if dm_build_521.dm_build_139.dm_build_1289.NewLobFlag {
		dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.exGroupId)
		dm_build_521.dm_build_139.dm_build_1288.Dm_build_1000(dm_build_521.dm_build_505.exFileId)
		dm_build_521.dm_build_139.dm_build_1288.Dm_build_1004(dm_build_521.dm_build_505.exPageNo)
	}
	return nil
}

func (dm_build_523 *dm_build_504) dm_build_129() (interface{}, error) {

	var dm_build_524 = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_523.dm_build_505.blobId = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1081()
	dm_build_523.dm_build_505.fileId = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1075()
	dm_build_523.dm_build_505.pageNo = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_523.dm_build_505.curFileId = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1075()
	dm_build_523.dm_build_505.curPageNo = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1078()
	dm_build_523.dm_build_505.totalOffset = dm_build_523.dm_build_139.dm_build_1288.Dm_build_1078()
	return dm_build_524, nil
}

const (
	Dm_build_525 = Dm_build_35

	Dm_build_526 = Dm_build_525 + ULINT_SIZE

	Dm_build_527 = Dm_build_526 + ULINT_SIZE

	Dm_build_528 = Dm_build_527 + BYTE_SIZE

	Dm_build_529 = Dm_build_528 + BYTE_SIZE

	Dm_build_530 = Dm_build_529 + BYTE_SIZE

	Dm_build_531 = Dm_build_530 + BYTE_SIZE

	Dm_build_532 = Dm_build_35

	Dm_build_533 = Dm_build_532 + ULINT_SIZE

	Dm_build_534 = Dm_build_533 + ULINT_SIZE

	Dm_build_535 = Dm_build_534 + ULINT_SIZE

	Dm_build_536 = Dm_build_535 + ULINT_SIZE

	Dm_build_537 = Dm_build_536 + ULINT_SIZE

	Dm_build_538 = Dm_build_537 + BYTE_SIZE

	Dm_build_539 = Dm_build_538 + BYTE_SIZE

	Dm_build_540 = Dm_build_539 + BYTE_SIZE
)

type dm_build_541 struct {
	dm_build_138
	dm_build_542 *DmConnection
	dm_build_543 int
	Dm_build_544 int32
	Dm_build_545 []byte
	dm_build_546 byte
}

func dm_build_547(dm_build_548 *dm_build_1285) *dm_build_541 {
	dm_build_549 := new(dm_build_541)
	dm_build_549.dm_build_143(dm_build_548, Dm_build_33)
	dm_build_549.dm_build_542 = dm_build_548.dm_build_1289
	return dm_build_549
}

func dm_build_550(dm_build_551 string, dm_build_552 string) int {
	dm_build_553 := strings.Split(dm_build_551, ".")
	dm_build_554 := strings.Split(dm_build_552, ".")

	for i, serStr := range dm_build_553 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_554[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_556 *dm_build_541) dm_build_125() error {

	dm_build_556.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_525, int32(0))
	dm_build_556.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_526, int32(dm_build_556.dm_build_542.dmConnector.compress))

	if dm_build_556.dm_build_542.dmConnector.loginEncrypt {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_528, 2)
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_527, 1)
	} else {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_528, 0)
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_527, 0)
	}

	if dm_build_556.dm_build_542.dmConnector.isBdtaRS {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_529, Dm_build_113)
	} else {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_529, Dm_build_112)
	}

	dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_530, byte(dm_build_556.dm_build_542.dmConnector.compressID))

	if dm_build_556.dm_build_542.dmConnector.loginCertificate != "" {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_531, 1)
	} else {
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_531, 0)
	}

	dm_build_557 := dm_build_556.dm_build_542.getServerEncoding()
	dm_build_556.dm_build_139.dm_build_1288.Dm_build_1048(Dm_build_0, dm_build_557, dm_build_556.dm_build_139.dm_build_1289)

	var dm_build_558 byte
	if dm_build_556.dm_build_542.dmConnector.uKeyName != "" {
		dm_build_558 = 1
	} else {
		dm_build_558 = 0
	}

	dm_build_556.dm_build_139.dm_build_1288.Dm_build_996(0)

	if dm_build_558 == 1 {

	}

	if dm_build_556.dm_build_542.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_556.dm_build_139.dm_build_1529()
		if err != nil {
			return err
		}
		dm_build_556.dm_build_139.dm_build_1288.Dm_build_1036(clientPubKey)
	}

	return nil
}

func (dm_build_560 *dm_build_541) dm_build_128() error {
	dm_build_560.dm_build_139.dm_build_1288.Dm_build_970(0)
	dm_build_560.dm_build_139.dm_build_1288.Dm_build_978(Dm_build_34, false, true)
	return nil
}

func (dm_build_562 *dm_build_541) dm_build_129() (interface{}, error) {

	dm_build_562.dm_build_542.sslEncrypt = int(dm_build_562.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_532))
	dm_build_562.dm_build_542.GlobalServerSeries = int(dm_build_562.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_533))

	switch dm_build_562.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_534) {
	case 1:
		dm_build_562.dm_build_542.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_562.dm_build_542.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_562.dm_build_542.serverEncoding = ENCODING_GB18030
	}

	dm_build_562.dm_build_542.dmConnector.compress = int(dm_build_562.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_535))
	dm_build_563 := dm_build_562.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_537)
	dm_build_564 := dm_build_562.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_538)
	dm_build_562.dm_build_542.dmConnector.isBdtaRS = dm_build_562.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_539) > 0
	dm_build_562.dm_build_542.dmConnector.compressID = int8(dm_build_562.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_540))

	dm_build_565 := dm_build_562.dm_build_171()
	if dm_build_565 != nil {
		return nil, dm_build_565
	}

	dm_build_566 := dm_build_562.dm_build_139.dm_build_1288.Dm_build_1120(dm_build_562.dm_build_542.getServerEncoding(), dm_build_562.dm_build_139.dm_build_1289)
	if dm_build_550(dm_build_566, Dm_build_1) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_562.dm_build_542.ServerVersion = dm_build_566
	dm_build_562.dm_build_542.Malini2 = dm_build_550(dm_build_566, Dm_build_2) > 0
	dm_build_562.dm_build_542.Execute2 = dm_build_550(dm_build_566, Dm_build_3) > 0
	dm_build_562.dm_build_542.LobEmptyCompOrcl = dm_build_550(dm_build_566, Dm_build_4) > 0

	if dm_build_562.dm_build_139.dm_build_1289.dmConnector.uKeyName != "" {
		dm_build_562.dm_build_546 = 1
	} else {
		dm_build_562.dm_build_546 = 0
	}

	if dm_build_562.dm_build_546 == 1 {
		dm_build_562.dm_build_139.dm_build_1294 = dm_build_562.dm_build_139.dm_build_1288.Dm_build_1115(16, dm_build_562.dm_build_542.getServerEncoding(), dm_build_562.dm_build_139.dm_build_1289)
	}

	dm_build_562.dm_build_543 = -1
	dm_build_567 := false
	dm_build_568 := false
	dm_build_562.Dm_build_544 = -1
	if dm_build_564 > 0 {
		dm_build_562.dm_build_543 = int(dm_build_562.dm_build_139.dm_build_1288.Dm_build_1078())
	}

	if dm_build_563 > 0 {

		if dm_build_562.dm_build_543 == -1 {
			dm_build_567 = true
		} else {
			dm_build_568 = true
		}

		dm_build_562.Dm_build_545 = dm_build_562.dm_build_139.dm_build_1288.Dm_build_1103()
	}

	if dm_build_564 == 2 {
		dm_build_562.Dm_build_544 = dm_build_562.dm_build_139.dm_build_1288.Dm_build_1078()
	}
	dm_build_562.dm_build_139.dm_build_1291 = dm_build_567
	dm_build_562.dm_build_139.dm_build_1292 = dm_build_568

	return nil, nil
}

type dm_build_569 struct {
	dm_build_138
}

func dm_build_570(dm_build_571 *dm_build_1285, dm_build_572 *DmStatement) *dm_build_569 {
	dm_build_573 := new(dm_build_569)
	dm_build_573.dm_build_147(dm_build_571, Dm_build_9, dm_build_572)
	return dm_build_573
}

func (dm_build_575 *dm_build_569) dm_build_125() error {

	dm_build_575.dm_build_139.dm_build_1288.Dm_build_1136(Dm_build_35, 1)
	return nil
}

func (dm_build_577 *dm_build_569) dm_build_129() (interface{}, error) {

	dm_build_577.dm_build_142.id = dm_build_577.dm_build_139.dm_build_1288.Dm_build_1222(Dm_build_36)

	dm_build_577.dm_build_142.readBaseColName = dm_build_577.dm_build_139.dm_build_1288.Dm_build_1216(Dm_build_35) == 1
	return nil, nil
}

type dm_build_578 struct {
	dm_build_138
	dm_build_579 int32
}

func dm_build_580(dm_build_581 *dm_build_1285, dm_build_582 int32) *dm_build_578 {
	dm_build_583 := new(dm_build_578)
	dm_build_583.dm_build_143(dm_build_581, Dm_build_10)
	dm_build_583.dm_build_579 = dm_build_582
	return dm_build_583
}

func (dm_build_585 *dm_build_578) dm_build_126() {
	dm_build_585.dm_build_138.dm_build_126()
	dm_build_585.dm_build_139.dm_build_1288.Dm_build_1144(Dm_build_36, dm_build_585.dm_build_579)
}

type dm_build_586 struct {
	dm_build_138
	dm_build_587 []uint32
}

func dm_build_588(dm_build_589 *dm_build_1285, dm_build_590 []uint32) *dm_build_586 {
	dm_build_591 := new(dm_build_586)
	dm_build_591.dm_build_143(dm_build_589, Dm_build_30)
	dm_build_591.dm_build_587 = dm_build_590
	return dm_build_591
}

func (dm_build_593 *dm_build_586) dm_build_125() error {

	dm_build_593.dm_build_139.dm_build_1288.Dm_build_1164(Dm_build_35, uint16(len(dm_build_593.dm_build_587)))

	for _, tableID := range dm_build_593.dm_build_587 {
		dm_build_593.dm_build_139.dm_build_1288.Dm_build_1016(uint32(tableID))
	}

	return nil
}

func (dm_build_595 *dm_build_586) dm_build_129() (interface{}, error) {
	dm_build_596 := dm_build_595.dm_build_139.dm_build_1288.Dm_build_1237(Dm_build_35)
	if dm_build_596 <= 0 {
		return nil, nil
	}

	dm_build_597 := make([]int64, dm_build_596)
	for i := 0; i < int(dm_build_596); i++ {
		dm_build_597[i] = dm_build_595.dm_build_139.dm_build_1288.Dm_build_1081()
	}

	return dm_build_597, nil
}
