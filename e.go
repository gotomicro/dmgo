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

type dm_build_884 struct{}

var Dm_build_885 = &dm_build_884{}

func (Dm_build_887 *dm_build_884) Dm_build_886(dm_build_888 []byte, dm_build_889 int, dm_build_890 byte) int {
	dm_build_888[dm_build_889] = dm_build_890
	return 1
}

func (Dm_build_892 *dm_build_884) Dm_build_891(dm_build_893 []byte, dm_build_894 int, dm_build_895 int8) int {
	dm_build_893[dm_build_894] = byte(dm_build_895)
	return 1
}

func (Dm_build_897 *dm_build_884) Dm_build_896(dm_build_898 []byte, dm_build_899 int, dm_build_900 int16) int {
	dm_build_898[dm_build_899] = byte(dm_build_900)
	dm_build_899++
	dm_build_898[dm_build_899] = byte(dm_build_900 >> 8)
	return 2
}

func (Dm_build_902 *dm_build_884) Dm_build_901(dm_build_903 []byte, dm_build_904 int, dm_build_905 int32) int {
	dm_build_903[dm_build_904] = byte(dm_build_905)
	dm_build_904++
	dm_build_903[dm_build_904] = byte(dm_build_905 >> 8)
	dm_build_904++
	dm_build_903[dm_build_904] = byte(dm_build_905 >> 16)
	dm_build_904++
	dm_build_903[dm_build_904] = byte(dm_build_905 >> 24)
	dm_build_904++
	return 4
}

func (Dm_build_907 *dm_build_884) Dm_build_906(dm_build_908 []byte, dm_build_909 int, dm_build_910 int64) int {
	dm_build_908[dm_build_909] = byte(dm_build_910)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 8)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 16)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 24)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 32)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 40)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 48)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 56)
	return 8
}

func (Dm_build_912 *dm_build_884) Dm_build_911(dm_build_913 []byte, dm_build_914 int, dm_build_915 float32) int {
	return Dm_build_912.Dm_build_931(dm_build_913, dm_build_914, math.Float32bits(dm_build_915))
}

func (Dm_build_917 *dm_build_884) Dm_build_916(dm_build_918 []byte, dm_build_919 int, dm_build_920 float64) int {
	return Dm_build_917.Dm_build_936(dm_build_918, dm_build_919, math.Float64bits(dm_build_920))
}

func (Dm_build_922 *dm_build_884) Dm_build_921(dm_build_923 []byte, dm_build_924 int, dm_build_925 uint8) int {
	dm_build_923[dm_build_924] = byte(dm_build_925)
	return 1
}

func (Dm_build_927 *dm_build_884) Dm_build_926(dm_build_928 []byte, dm_build_929 int, dm_build_930 uint16) int {
	dm_build_928[dm_build_929] = byte(dm_build_930)
	dm_build_929++
	dm_build_928[dm_build_929] = byte(dm_build_930 >> 8)
	return 2
}

func (Dm_build_932 *dm_build_884) Dm_build_931(dm_build_933 []byte, dm_build_934 int, dm_build_935 uint32) int {
	dm_build_933[dm_build_934] = byte(dm_build_935)
	dm_build_934++
	dm_build_933[dm_build_934] = byte(dm_build_935 >> 8)
	dm_build_934++
	dm_build_933[dm_build_934] = byte(dm_build_935 >> 16)
	dm_build_934++
	dm_build_933[dm_build_934] = byte(dm_build_935 >> 24)
	return 3
}

func (Dm_build_937 *dm_build_884) Dm_build_936(dm_build_938 []byte, dm_build_939 int, dm_build_940 uint64) int {
	dm_build_938[dm_build_939] = byte(dm_build_940)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 8)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 16)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 24)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 32)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 40)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 48)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 56)
	return 3
}

func (Dm_build_942 *dm_build_884) Dm_build_941(dm_build_943 []byte, dm_build_944 int, dm_build_945 []byte, dm_build_946 int, dm_build_947 int) int {
	copy(dm_build_943[dm_build_944:dm_build_944+dm_build_947], dm_build_945[dm_build_946:dm_build_946+dm_build_947])
	return dm_build_947
}

func (Dm_build_949 *dm_build_884) Dm_build_948(dm_build_950 []byte, dm_build_951 int, dm_build_952 []byte, dm_build_953 int, dm_build_954 int) int {
	dm_build_951 += Dm_build_949.Dm_build_931(dm_build_950, dm_build_951, uint32(dm_build_954))
	return 4 + Dm_build_949.Dm_build_941(dm_build_950, dm_build_951, dm_build_952, dm_build_953, dm_build_954)
}

func (Dm_build_956 *dm_build_884) Dm_build_955(dm_build_957 []byte, dm_build_958 int, dm_build_959 []byte, dm_build_960 int, dm_build_961 int) int {
	dm_build_958 += Dm_build_956.Dm_build_926(dm_build_957, dm_build_958, uint16(dm_build_961))
	return 2 + Dm_build_956.Dm_build_941(dm_build_957, dm_build_958, dm_build_959, dm_build_960, dm_build_961)
}

func (Dm_build_963 *dm_build_884) Dm_build_962(dm_build_964 []byte, dm_build_965 int, dm_build_966 string, dm_build_967 string, dm_build_968 *DmConnection) int {
	dm_build_969 := Dm_build_963.Dm_build_1095(dm_build_966, dm_build_967, dm_build_968)
	dm_build_965 += Dm_build_963.Dm_build_931(dm_build_964, dm_build_965, uint32(len(dm_build_969)))
	return 4 + Dm_build_963.Dm_build_941(dm_build_964, dm_build_965, dm_build_969, 0, len(dm_build_969))
}

func (Dm_build_971 *dm_build_884) Dm_build_970(dm_build_972 []byte, dm_build_973 int, dm_build_974 string, dm_build_975 string, dm_build_976 *DmConnection) int {
	dm_build_977 := Dm_build_971.Dm_build_1095(dm_build_974, dm_build_975, dm_build_976)

	dm_build_973 += Dm_build_971.Dm_build_926(dm_build_972, dm_build_973, uint16(len(dm_build_977)))
	return 2 + Dm_build_971.Dm_build_941(dm_build_972, dm_build_973, dm_build_977, 0, len(dm_build_977))
}

func (Dm_build_979 *dm_build_884) Dm_build_978(dm_build_980 []byte, dm_build_981 int) byte {
	return dm_build_980[dm_build_981]
}

func (Dm_build_983 *dm_build_884) Dm_build_982(dm_build_984 []byte, dm_build_985 int) int16 {
	var dm_build_986 int16
	dm_build_986 = int16(dm_build_984[dm_build_985] & 0xff)
	dm_build_985++
	dm_build_986 |= int16(dm_build_984[dm_build_985]&0xff) << 8
	return dm_build_986
}

func (Dm_build_988 *dm_build_884) Dm_build_987(dm_build_989 []byte, dm_build_990 int) int32 {
	var dm_build_991 int32
	dm_build_991 = int32(dm_build_989[dm_build_990] & 0xff)
	dm_build_990++
	dm_build_991 |= int32(dm_build_989[dm_build_990]&0xff) << 8
	dm_build_990++
	dm_build_991 |= int32(dm_build_989[dm_build_990]&0xff) << 16
	dm_build_990++
	dm_build_991 |= int32(dm_build_989[dm_build_990]&0xff) << 24
	return dm_build_991
}

func (Dm_build_993 *dm_build_884) Dm_build_992(dm_build_994 []byte, dm_build_995 int) int64 {
	var dm_build_996 int64
	dm_build_996 = int64(dm_build_994[dm_build_995] & 0xff)
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 8
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 16
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 24
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 32
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 40
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 48
	dm_build_995++
	dm_build_996 |= int64(dm_build_994[dm_build_995]&0xff) << 56
	return dm_build_996
}

func (Dm_build_998 *dm_build_884) Dm_build_997(dm_build_999 []byte, dm_build_1000 int) float32 {
	return math.Float32frombits(Dm_build_998.Dm_build_1014(dm_build_999, dm_build_1000))
}

func (Dm_build_1002 *dm_build_884) Dm_build_1001(dm_build_1003 []byte, dm_build_1004 int) float64 {
	return math.Float64frombits(Dm_build_1002.Dm_build_1019(dm_build_1003, dm_build_1004))
}

func (Dm_build_1006 *dm_build_884) Dm_build_1005(dm_build_1007 []byte, dm_build_1008 int) uint8 {
	return uint8(dm_build_1007[dm_build_1008] & 0xff)
}

func (Dm_build_1010 *dm_build_884) Dm_build_1009(dm_build_1011 []byte, dm_build_1012 int) uint16 {
	var dm_build_1013 uint16
	dm_build_1013 = uint16(dm_build_1011[dm_build_1012] & 0xff)
	dm_build_1012++
	dm_build_1013 |= uint16(dm_build_1011[dm_build_1012]&0xff) << 8
	return dm_build_1013
}

func (Dm_build_1015 *dm_build_884) Dm_build_1014(dm_build_1016 []byte, dm_build_1017 int) uint32 {
	var dm_build_1018 uint32
	dm_build_1018 = uint32(dm_build_1016[dm_build_1017] & 0xff)
	dm_build_1017++
	dm_build_1018 |= uint32(dm_build_1016[dm_build_1017]&0xff) << 8
	dm_build_1017++
	dm_build_1018 |= uint32(dm_build_1016[dm_build_1017]&0xff) << 16
	dm_build_1017++
	dm_build_1018 |= uint32(dm_build_1016[dm_build_1017]&0xff) << 24
	return dm_build_1018
}

func (Dm_build_1020 *dm_build_884) Dm_build_1019(dm_build_1021 []byte, dm_build_1022 int) uint64 {
	var dm_build_1023 uint64
	dm_build_1023 = uint64(dm_build_1021[dm_build_1022] & 0xff)
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 8
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 16
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 24
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 32
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 40
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 48
	dm_build_1022++
	dm_build_1023 |= uint64(dm_build_1021[dm_build_1022]&0xff) << 56
	return dm_build_1023
}

func (Dm_build_1025 *dm_build_884) Dm_build_1024(dm_build_1026 []byte, dm_build_1027 int) []byte {
	dm_build_1028 := Dm_build_1025.Dm_build_1014(dm_build_1026, dm_build_1027)
	return dm_build_1026[dm_build_1027+4 : dm_build_1027+4+int(dm_build_1028)]

}

func (Dm_build_1030 *dm_build_884) Dm_build_1029(dm_build_1031 []byte, dm_build_1032 int) []byte {
	dm_build_1033 := Dm_build_1030.Dm_build_1009(dm_build_1031, dm_build_1032)
	return dm_build_1031[dm_build_1032+2 : dm_build_1032+2+int(dm_build_1033)]

}

func (Dm_build_1035 *dm_build_884) Dm_build_1034(dm_build_1036 []byte, dm_build_1037 int, dm_build_1038 int) []byte {
	return dm_build_1036[dm_build_1037 : dm_build_1037+dm_build_1038]

}

func (Dm_build_1040 *dm_build_884) Dm_build_1039(dm_build_1041 []byte, dm_build_1042 int, dm_build_1043 int, dm_build_1044 string, dm_build_1045 *DmConnection) string {
	return Dm_build_1040.Dm_build_1132(dm_build_1041[dm_build_1042:dm_build_1042+dm_build_1043], dm_build_1044, dm_build_1045)
}

func (Dm_build_1047 *dm_build_884) Dm_build_1046(dm_build_1048 []byte, dm_build_1049 int, dm_build_1050 string, dm_build_1051 *DmConnection) string {
	dm_build_1052 := Dm_build_1047.Dm_build_1014(dm_build_1048, dm_build_1049)
	dm_build_1049 += 4
	return Dm_build_1047.Dm_build_1039(dm_build_1048, dm_build_1049, int(dm_build_1052), dm_build_1050, dm_build_1051)
}

func (Dm_build_1054 *dm_build_884) Dm_build_1053(dm_build_1055 []byte, dm_build_1056 int, dm_build_1057 string, dm_build_1058 *DmConnection) string {
	dm_build_1059 := Dm_build_1054.Dm_build_1009(dm_build_1055, dm_build_1056)
	dm_build_1056 += 2
	return Dm_build_1054.Dm_build_1039(dm_build_1055, dm_build_1056, int(dm_build_1059), dm_build_1057, dm_build_1058)
}

func (Dm_build_1061 *dm_build_884) Dm_build_1060(dm_build_1062 byte) []byte {
	return []byte{dm_build_1062}
}

func (Dm_build_1064 *dm_build_884) Dm_build_1063(dm_build_1065 int16) []byte {
	return []byte{byte(dm_build_1065), byte(dm_build_1065 >> 8)}
}

func (Dm_build_1067 *dm_build_884) Dm_build_1066(dm_build_1068 int32) []byte {
	return []byte{byte(dm_build_1068), byte(dm_build_1068 >> 8), byte(dm_build_1068 >> 16), byte(dm_build_1068 >> 24)}
}

func (Dm_build_1070 *dm_build_884) Dm_build_1069(dm_build_1071 int64) []byte {
	return []byte{byte(dm_build_1071), byte(dm_build_1071 >> 8), byte(dm_build_1071 >> 16), byte(dm_build_1071 >> 24), byte(dm_build_1071 >> 32),
		byte(dm_build_1071 >> 40), byte(dm_build_1071 >> 48), byte(dm_build_1071 >> 56)}
}

func (Dm_build_1073 *dm_build_884) Dm_build_1072(dm_build_1074 float32) []byte {
	return Dm_build_1073.Dm_build_1084(math.Float32bits(dm_build_1074))
}

func (Dm_build_1076 *dm_build_884) Dm_build_1075(dm_build_1077 float64) []byte {
	return Dm_build_1076.Dm_build_1087(math.Float64bits(dm_build_1077))
}

func (Dm_build_1079 *dm_build_884) Dm_build_1078(dm_build_1080 uint8) []byte {
	return []byte{byte(dm_build_1080)}
}

func (Dm_build_1082 *dm_build_884) Dm_build_1081(dm_build_1083 uint16) []byte {
	return []byte{byte(dm_build_1083), byte(dm_build_1083 >> 8)}
}

func (Dm_build_1085 *dm_build_884) Dm_build_1084(dm_build_1086 uint32) []byte {
	return []byte{byte(dm_build_1086), byte(dm_build_1086 >> 8), byte(dm_build_1086 >> 16), byte(dm_build_1086 >> 24)}
}

func (Dm_build_1088 *dm_build_884) Dm_build_1087(dm_build_1089 uint64) []byte {
	return []byte{byte(dm_build_1089), byte(dm_build_1089 >> 8), byte(dm_build_1089 >> 16), byte(dm_build_1089 >> 24), byte(dm_build_1089 >> 32), byte(dm_build_1089 >> 40), byte(dm_build_1089 >> 48), byte(dm_build_1089 >> 56)}
}

func (Dm_build_1091 *dm_build_884) Dm_build_1090(dm_build_1092 []byte, dm_build_1093 string, dm_build_1094 *DmConnection) []byte {
	if dm_build_1093 == "UTF-8" {
		return dm_build_1092
	}

	if dm_build_1094 == nil {
		if e := dm_build_1137(dm_build_1093); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1092), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1094.encodeBuffer == nil {
		dm_build_1094.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1094.encode = dm_build_1137(dm_build_1094.getServerEncoding())
		dm_build_1094.transformReaderDst = make([]byte, 4096)
		dm_build_1094.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1094.encode; e != nil {

		dm_build_1094.encodeBuffer.Reset()

		n, err := dm_build_1094.encodeBuffer.ReadFrom(
			Dm_build_1151(bytes.NewReader(dm_build_1092), e.NewEncoder(), dm_build_1094.transformReaderDst, dm_build_1094.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1094.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1096 *dm_build_884) Dm_build_1095(dm_build_1097 string, dm_build_1098 string, dm_build_1099 *DmConnection) []byte {
	return Dm_build_1096.Dm_build_1090([]byte(dm_build_1097), dm_build_1098, dm_build_1099)
}

func (Dm_build_1101 *dm_build_884) Dm_build_1100(dm_build_1102 []byte) byte {
	return Dm_build_1101.Dm_build_978(dm_build_1102, 0)
}

func (Dm_build_1104 *dm_build_884) Dm_build_1103(dm_build_1105 []byte) int16 {
	return Dm_build_1104.Dm_build_982(dm_build_1105, 0)
}

func (Dm_build_1107 *dm_build_884) Dm_build_1106(dm_build_1108 []byte) int32 {
	return Dm_build_1107.Dm_build_987(dm_build_1108, 0)
}

func (Dm_build_1110 *dm_build_884) Dm_build_1109(dm_build_1111 []byte) int64 {
	return Dm_build_1110.Dm_build_992(dm_build_1111, 0)
}

func (Dm_build_1113 *dm_build_884) Dm_build_1112(dm_build_1114 []byte) float32 {
	return Dm_build_1113.Dm_build_997(dm_build_1114, 0)
}

func (Dm_build_1116 *dm_build_884) Dm_build_1115(dm_build_1117 []byte) float64 {
	return Dm_build_1116.Dm_build_1001(dm_build_1117, 0)
}

func (Dm_build_1119 *dm_build_884) Dm_build_1118(dm_build_1120 []byte) uint8 {
	return Dm_build_1119.Dm_build_1005(dm_build_1120, 0)
}

func (Dm_build_1122 *dm_build_884) Dm_build_1121(dm_build_1123 []byte) uint16 {
	return Dm_build_1122.Dm_build_1009(dm_build_1123, 0)
}

func (Dm_build_1125 *dm_build_884) Dm_build_1124(dm_build_1126 []byte) uint32 {
	return Dm_build_1125.Dm_build_1014(dm_build_1126, 0)
}

func (Dm_build_1128 *dm_build_884) Dm_build_1127(dm_build_1129 []byte, dm_build_1130 string, dm_build_1131 *DmConnection) []byte {
	if dm_build_1130 == "UTF-8" {
		return dm_build_1129
	}

	if dm_build_1131 == nil {
		if e := dm_build_1137(dm_build_1130); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1129), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1131.encodeBuffer == nil {
		dm_build_1131.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1131.encode = dm_build_1137(dm_build_1131.getServerEncoding())
		dm_build_1131.transformReaderDst = make([]byte, 4096)
		dm_build_1131.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1131.encode; e != nil {

		dm_build_1131.encodeBuffer.Reset()

		n, err := dm_build_1131.encodeBuffer.ReadFrom(
			Dm_build_1151(bytes.NewReader(dm_build_1129), e.NewDecoder(), dm_build_1131.transformReaderDst, dm_build_1131.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1131.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1133 *dm_build_884) Dm_build_1132(dm_build_1134 []byte, dm_build_1135 string, dm_build_1136 *DmConnection) string {
	return string(Dm_build_1133.Dm_build_1127(dm_build_1134, dm_build_1135, dm_build_1136))
}

func dm_build_1137(dm_build_1138 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1138); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1139 struct {
	dm_build_1140 io.Reader
	dm_build_1141 transform.Transformer
	dm_build_1142 error

	dm_build_1143                []byte
	dm_build_1144, dm_build_1145 int

	dm_build_1146                []byte
	dm_build_1147, dm_build_1148 int

	dm_build_1149 bool
}

const dm_build_1150 = 4096

func Dm_build_1151(dm_build_1152 io.Reader, dm_build_1153 transform.Transformer, dm_build_1154 []byte, dm_build_1155 []byte) *Dm_build_1139 {
	dm_build_1153.Reset()
	return &Dm_build_1139{
		dm_build_1140: dm_build_1152,
		dm_build_1141: dm_build_1153,
		dm_build_1143: dm_build_1154,
		dm_build_1146: dm_build_1155,
	}
}

func (dm_build_1157 *Dm_build_1139) Read(dm_build_1158 []byte) (int, error) {
	dm_build_1159, dm_build_1160 := 0, error(nil)
	for {

		if dm_build_1157.dm_build_1144 != dm_build_1157.dm_build_1145 {
			dm_build_1159 = copy(dm_build_1158, dm_build_1157.dm_build_1143[dm_build_1157.dm_build_1144:dm_build_1157.dm_build_1145])
			dm_build_1157.dm_build_1144 += dm_build_1159
			if dm_build_1157.dm_build_1144 == dm_build_1157.dm_build_1145 && dm_build_1157.dm_build_1149 {
				return dm_build_1159, dm_build_1157.dm_build_1142
			}
			return dm_build_1159, nil
		} else if dm_build_1157.dm_build_1149 {
			return 0, dm_build_1157.dm_build_1142
		}

		if dm_build_1157.dm_build_1147 != dm_build_1157.dm_build_1148 || dm_build_1157.dm_build_1142 != nil {
			dm_build_1157.dm_build_1144 = 0
			dm_build_1157.dm_build_1145, dm_build_1159, dm_build_1160 = dm_build_1157.dm_build_1141.Transform(dm_build_1157.dm_build_1143, dm_build_1157.dm_build_1146[dm_build_1157.dm_build_1147:dm_build_1157.dm_build_1148], dm_build_1157.dm_build_1142 == io.EOF)
			dm_build_1157.dm_build_1147 += dm_build_1159

			switch {
			case dm_build_1160 == nil:
				if dm_build_1157.dm_build_1147 != dm_build_1157.dm_build_1148 {
					dm_build_1157.dm_build_1142 = nil
				}

				dm_build_1157.dm_build_1149 = dm_build_1157.dm_build_1142 != nil
				continue
			case dm_build_1160 == transform.ErrShortDst && (dm_build_1157.dm_build_1145 != 0 || dm_build_1159 != 0):

				continue
			case dm_build_1160 == transform.ErrShortSrc && dm_build_1157.dm_build_1148-dm_build_1157.dm_build_1147 != len(dm_build_1157.dm_build_1146) && dm_build_1157.dm_build_1142 == nil:

			default:
				dm_build_1157.dm_build_1149 = true

				if dm_build_1157.dm_build_1142 == nil || dm_build_1157.dm_build_1142 == io.EOF {
					dm_build_1157.dm_build_1142 = dm_build_1160
				}
				continue
			}
		}

		if dm_build_1157.dm_build_1147 != 0 {
			dm_build_1157.dm_build_1147, dm_build_1157.dm_build_1148 = 0, copy(dm_build_1157.dm_build_1146, dm_build_1157.dm_build_1146[dm_build_1157.dm_build_1147:dm_build_1157.dm_build_1148])
		}
		dm_build_1159, dm_build_1157.dm_build_1142 = dm_build_1157.dm_build_1140.Read(dm_build_1157.dm_build_1146[dm_build_1157.dm_build_1148:])
		dm_build_1157.dm_build_1148 += dm_build_1159
	}
}
