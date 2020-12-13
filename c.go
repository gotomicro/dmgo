/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_953 struct {
	dm_build_954 []byte
	dm_build_955 int
}

func Dm_build_956(dm_build_957 int) *Dm_build_953 {
	return &Dm_build_953{make([]byte, 0, dm_build_957), 0}
}

func Dm_build_958(dm_build_959 []byte) *Dm_build_953 {
	return &Dm_build_953{dm_build_959, 0}
}

func (dm_build_961 *Dm_build_953) dm_build_960(dm_build_962 int) *Dm_build_953 {

	dm_build_963 := len(dm_build_961.dm_build_954)
	dm_build_964 := cap(dm_build_961.dm_build_954)

	if dm_build_963+dm_build_962 <= dm_build_964 {
		dm_build_961.dm_build_954 = dm_build_961.dm_build_954[:dm_build_963+dm_build_962]
	} else {
		remain := dm_build_962 + dm_build_963 - dm_build_964
		nbuf := make([]byte, dm_build_962+dm_build_963, 2*dm_build_964+remain)
		copy(nbuf, dm_build_961.dm_build_954)
		dm_build_961.dm_build_954 = nbuf
	}

	return dm_build_961
}

func (dm_build_966 *Dm_build_953) Dm_build_965() int {
	return len(dm_build_966.dm_build_954)
}

func (dm_build_968 *Dm_build_953) Dm_build_967(dm_build_969 int) *Dm_build_953 {
	dm_build_968.dm_build_954 = dm_build_968.dm_build_954[:dm_build_969]
	return dm_build_968
}

func (dm_build_971 *Dm_build_953) Dm_build_970(dm_build_972 int) *Dm_build_953 {
	dm_build_971.dm_build_955 = dm_build_972
	return dm_build_971
}

func (dm_build_974 *Dm_build_953) Dm_build_973() int {
	return dm_build_974.dm_build_955
}

func (dm_build_976 *Dm_build_953) Dm_build_975(dm_build_977 bool) int {
	return len(dm_build_976.dm_build_954) - dm_build_976.dm_build_955
}

func (dm_build_979 *Dm_build_953) Dm_build_978(dm_build_980 int, dm_build_981 bool, dm_build_982 bool) *Dm_build_953 {

	if dm_build_981 {
		if dm_build_982 {
			dm_build_979.dm_build_960(dm_build_980)
		} else {
			dm_build_979.dm_build_954 = dm_build_979.dm_build_954[:len(dm_build_979.dm_build_954)-dm_build_980]
		}
	} else {
		if dm_build_982 {
			dm_build_979.dm_build_955 += dm_build_980
		} else {
			dm_build_979.dm_build_955 -= dm_build_980
		}
	}

	return dm_build_979
}

func (dm_build_984 *Dm_build_953) Dm_build_983(dm_build_985 io.Reader, dm_build_986 int) int {
	dm_build_987 := len(dm_build_984.dm_build_954)
	dm_build_984.dm_build_960(dm_build_986)
	dm_build_988 := 0
	for dm_build_986 > 0 {
		n, err := dm_build_985.Read(dm_build_984.dm_build_954[dm_build_987+dm_build_988:])
		if n > 0 && err == io.EOF {
			dm_build_988 += n
			dm_build_984.dm_build_954 = dm_build_984.dm_build_954[:dm_build_987+dm_build_988]
			return dm_build_988
		} else if n > 0 && err == nil {
			dm_build_986 -= n
			dm_build_988 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_988
}

func (dm_build_990 *Dm_build_953) Dm_build_989(dm_build_991 io.Writer) *Dm_build_953 {
	dm_build_991.Write(dm_build_990.dm_build_954)
	return dm_build_990
}

func (dm_build_993 *Dm_build_953) Dm_build_992(dm_build_994 bool) int {
	dm_build_995 := len(dm_build_993.dm_build_954)
	dm_build_993.dm_build_960(1)

	if dm_build_994 {
		return copy(dm_build_993.dm_build_954[dm_build_995:], []byte{1})
	} else {
		return copy(dm_build_993.dm_build_954[dm_build_995:], []byte{0})
	}
}

func (dm_build_997 *Dm_build_953) Dm_build_996(dm_build_998 byte) int {
	dm_build_999 := len(dm_build_997.dm_build_954)
	dm_build_997.dm_build_960(1)

	return copy(dm_build_997.dm_build_954[dm_build_999:], Dm_build_599.Dm_build_774(dm_build_998))
}

func (dm_build_1001 *Dm_build_953) Dm_build_1000(dm_build_1002 int16) int {
	dm_build_1003 := len(dm_build_1001.dm_build_954)
	dm_build_1001.dm_build_960(2)

	return copy(dm_build_1001.dm_build_954[dm_build_1003:], Dm_build_599.Dm_build_777(dm_build_1002))
}

func (dm_build_1005 *Dm_build_953) Dm_build_1004(dm_build_1006 int32) int {
	dm_build_1007 := len(dm_build_1005.dm_build_954)
	dm_build_1005.dm_build_960(4)

	return copy(dm_build_1005.dm_build_954[dm_build_1007:], Dm_build_599.Dm_build_780(dm_build_1006))
}

func (dm_build_1009 *Dm_build_953) Dm_build_1008(dm_build_1010 uint8) int {
	dm_build_1011 := len(dm_build_1009.dm_build_954)
	dm_build_1009.dm_build_960(1)

	return copy(dm_build_1009.dm_build_954[dm_build_1011:], Dm_build_599.Dm_build_792(dm_build_1010))
}

func (dm_build_1013 *Dm_build_953) Dm_build_1012(dm_build_1014 uint16) int {
	dm_build_1015 := len(dm_build_1013.dm_build_954)
	dm_build_1013.dm_build_960(2)

	return copy(dm_build_1013.dm_build_954[dm_build_1015:], Dm_build_599.Dm_build_795(dm_build_1014))
}

func (dm_build_1017 *Dm_build_953) Dm_build_1016(dm_build_1018 uint32) int {
	dm_build_1019 := len(dm_build_1017.dm_build_954)
	dm_build_1017.dm_build_960(4)

	return copy(dm_build_1017.dm_build_954[dm_build_1019:], Dm_build_599.Dm_build_798(dm_build_1018))
}

func (dm_build_1021 *Dm_build_953) Dm_build_1020(dm_build_1022 uint64) int {
	dm_build_1023 := len(dm_build_1021.dm_build_954)
	dm_build_1021.dm_build_960(8)

	return copy(dm_build_1021.dm_build_954[dm_build_1023:], Dm_build_599.Dm_build_801(dm_build_1022))
}

func (dm_build_1025 *Dm_build_953) Dm_build_1024(dm_build_1026 float32) int {
	dm_build_1027 := len(dm_build_1025.dm_build_954)
	dm_build_1025.dm_build_960(4)

	return copy(dm_build_1025.dm_build_954[dm_build_1027:], Dm_build_599.Dm_build_798(math.Float32bits(dm_build_1026)))
}

func (dm_build_1029 *Dm_build_953) Dm_build_1028(dm_build_1030 float64) int {
	dm_build_1031 := len(dm_build_1029.dm_build_954)
	dm_build_1029.dm_build_960(8)

	return copy(dm_build_1029.dm_build_954[dm_build_1031:], Dm_build_599.Dm_build_801(math.Float64bits(dm_build_1030)))
}

func (dm_build_1033 *Dm_build_953) Dm_build_1032(dm_build_1034 []byte) int {
	dm_build_1035 := len(dm_build_1033.dm_build_954)
	dm_build_1033.dm_build_960(len(dm_build_1034))
	return copy(dm_build_1033.dm_build_954[dm_build_1035:], dm_build_1034)
}

func (dm_build_1037 *Dm_build_953) Dm_build_1036(dm_build_1038 []byte) int {
	return dm_build_1037.Dm_build_1004(int32(len(dm_build_1038))) + dm_build_1037.Dm_build_1032(dm_build_1038)
}

func (dm_build_1040 *Dm_build_953) Dm_build_1039(dm_build_1041 []byte) int {
	return dm_build_1040.Dm_build_1008(uint8(len(dm_build_1041))) + dm_build_1040.Dm_build_1032(dm_build_1041)
}

func (dm_build_1043 *Dm_build_953) Dm_build_1042(dm_build_1044 []byte) int {
	return dm_build_1043.Dm_build_1012(uint16(len(dm_build_1044))) + dm_build_1043.Dm_build_1032(dm_build_1044)
}

func (dm_build_1046 *Dm_build_953) Dm_build_1045(dm_build_1047 []byte) int {
	return dm_build_1046.Dm_build_1032(dm_build_1047) + dm_build_1046.Dm_build_996(0)
}

func (dm_build_1049 *Dm_build_953) Dm_build_1048(dm_build_1050 string, dm_build_1051 string, dm_build_1052 *DmConnection) int {
	dm_build_1053 := Dm_build_599.Dm_build_809(dm_build_1050, dm_build_1051, dm_build_1052)
	return dm_build_1049.Dm_build_1036(dm_build_1053)
}

func (dm_build_1055 *Dm_build_953) Dm_build_1054(dm_build_1056 string, dm_build_1057 string, dm_build_1058 *DmConnection) int {
	dm_build_1059 := Dm_build_599.Dm_build_809(dm_build_1056, dm_build_1057, dm_build_1058)
	return dm_build_1055.Dm_build_1039(dm_build_1059)
}

func (dm_build_1061 *Dm_build_953) Dm_build_1060(dm_build_1062 string, dm_build_1063 string, dm_build_1064 *DmConnection) int {
	dm_build_1065 := Dm_build_599.Dm_build_809(dm_build_1062, dm_build_1063, dm_build_1064)
	return dm_build_1061.Dm_build_1042(dm_build_1065)
}

func (dm_build_1067 *Dm_build_953) Dm_build_1066(dm_build_1068 string, dm_build_1069 string, dm_build_1070 *DmConnection) int {
	dm_build_1071 := Dm_build_599.Dm_build_809(dm_build_1068, dm_build_1069, dm_build_1070)
	return dm_build_1067.Dm_build_1045(dm_build_1071)
}

func (dm_build_1073 *Dm_build_953) Dm_build_1072() byte {
	dm_build_1074 := Dm_build_599.Dm_build_692(dm_build_1073.dm_build_954, dm_build_1073.dm_build_955)
	dm_build_1073.dm_build_955++
	return dm_build_1074
}

func (dm_build_1076 *Dm_build_953) Dm_build_1075() int16 {
	dm_build_1077 := Dm_build_599.Dm_build_696(dm_build_1076.dm_build_954, dm_build_1076.dm_build_955)
	dm_build_1076.dm_build_955 += 2
	return dm_build_1077
}

func (dm_build_1079 *Dm_build_953) Dm_build_1078() int32 {
	dm_build_1080 := Dm_build_599.Dm_build_701(dm_build_1079.dm_build_954, dm_build_1079.dm_build_955)
	dm_build_1079.dm_build_955 += 4
	return dm_build_1080
}

func (dm_build_1082 *Dm_build_953) Dm_build_1081() int64 {
	dm_build_1083 := Dm_build_599.Dm_build_706(dm_build_1082.dm_build_954, dm_build_1082.dm_build_955)
	dm_build_1082.dm_build_955 += 8
	return dm_build_1083
}

func (dm_build_1085 *Dm_build_953) Dm_build_1084() float32 {
	dm_build_1086 := Dm_build_599.Dm_build_711(dm_build_1085.dm_build_954, dm_build_1085.dm_build_955)
	dm_build_1085.dm_build_955 += 4
	return dm_build_1086
}

func (dm_build_1088 *Dm_build_953) Dm_build_1087() float64 {
	dm_build_1089 := Dm_build_599.Dm_build_715(dm_build_1088.dm_build_954, dm_build_1088.dm_build_955)
	dm_build_1088.dm_build_955 += 8
	return dm_build_1089
}

func (dm_build_1091 *Dm_build_953) Dm_build_1090() uint8 {
	dm_build_1092 := Dm_build_599.Dm_build_719(dm_build_1091.dm_build_954, dm_build_1091.dm_build_955)
	dm_build_1091.dm_build_955 += 1
	return dm_build_1092
}

func (dm_build_1094 *Dm_build_953) Dm_build_1093() uint16 {
	dm_build_1095 := Dm_build_599.Dm_build_723(dm_build_1094.dm_build_954, dm_build_1094.dm_build_955)
	dm_build_1094.dm_build_955 += 2
	return dm_build_1095
}

func (dm_build_1097 *Dm_build_953) Dm_build_1096() uint32 {
	dm_build_1098 := Dm_build_599.Dm_build_728(dm_build_1097.dm_build_954, dm_build_1097.dm_build_955)
	dm_build_1097.dm_build_955 += 4
	return dm_build_1098
}

func (dm_build_1100 *Dm_build_953) Dm_build_1099(dm_build_1101 int) []byte {
	dm_build_1102 := Dm_build_599.Dm_build_748(dm_build_1100.dm_build_954, dm_build_1100.dm_build_955, dm_build_1101)
	dm_build_1100.dm_build_955 += dm_build_1101
	return dm_build_1102
}

func (dm_build_1104 *Dm_build_953) Dm_build_1103() []byte {
	return dm_build_1104.Dm_build_1099(int(dm_build_1104.Dm_build_1078()))
}

func (dm_build_1106 *Dm_build_953) Dm_build_1105() []byte {
	return dm_build_1106.Dm_build_1099(int(dm_build_1106.Dm_build_1072()))
}

func (dm_build_1108 *Dm_build_953) Dm_build_1107() []byte {
	return dm_build_1108.Dm_build_1099(int(dm_build_1108.Dm_build_1075()))
}

func (dm_build_1110 *Dm_build_953) Dm_build_1109(dm_build_1111 int) []byte {
	return dm_build_1110.Dm_build_1099(dm_build_1111)
}

func (dm_build_1113 *Dm_build_953) Dm_build_1112() []byte {
	dm_build_1114 := 0
	for dm_build_1113.Dm_build_1072() != 0 {
		dm_build_1114++
	}
	dm_build_1113.Dm_build_978(dm_build_1114, false, false)
	return dm_build_1113.Dm_build_1099(dm_build_1114)
}

func (dm_build_1116 *Dm_build_953) Dm_build_1115(dm_build_1117 int, dm_build_1118 string, dm_build_1119 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1116.Dm_build_1099(dm_build_1117), dm_build_1118, dm_build_1119)
}

func (dm_build_1121 *Dm_build_953) Dm_build_1120(dm_build_1122 string, dm_build_1123 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1121.Dm_build_1103(), dm_build_1122, dm_build_1123)
}

func (dm_build_1125 *Dm_build_953) Dm_build_1124(dm_build_1126 string, dm_build_1127 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1125.Dm_build_1105(), dm_build_1126, dm_build_1127)
}

func (dm_build_1129 *Dm_build_953) Dm_build_1128(dm_build_1130 string, dm_build_1131 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1129.Dm_build_1107(), dm_build_1130, dm_build_1131)
}

func (dm_build_1133 *Dm_build_953) Dm_build_1132(dm_build_1134 string, dm_build_1135 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1133.Dm_build_1112(), dm_build_1134, dm_build_1135)
}

func (dm_build_1137 *Dm_build_953) Dm_build_1136(dm_build_1138 int, dm_build_1139 byte) int {
	return dm_build_1137.Dm_build_1172(dm_build_1138, Dm_build_599.Dm_build_774(dm_build_1139))
}

func (dm_build_1141 *Dm_build_953) Dm_build_1140(dm_build_1142 int, dm_build_1143 int16) int {
	return dm_build_1141.Dm_build_1172(dm_build_1142, Dm_build_599.Dm_build_777(dm_build_1143))
}

func (dm_build_1145 *Dm_build_953) Dm_build_1144(dm_build_1146 int, dm_build_1147 int32) int {
	return dm_build_1145.Dm_build_1172(dm_build_1146, Dm_build_599.Dm_build_780(dm_build_1147))
}

func (dm_build_1149 *Dm_build_953) Dm_build_1148(dm_build_1150 int, dm_build_1151 int64) int {
	return dm_build_1149.Dm_build_1172(dm_build_1150, Dm_build_599.Dm_build_783(dm_build_1151))
}

func (dm_build_1153 *Dm_build_953) Dm_build_1152(dm_build_1154 int, dm_build_1155 float32) int {
	return dm_build_1153.Dm_build_1172(dm_build_1154, Dm_build_599.Dm_build_786(dm_build_1155))
}

func (dm_build_1157 *Dm_build_953) Dm_build_1156(dm_build_1158 int, dm_build_1159 float64) int {
	return dm_build_1157.Dm_build_1172(dm_build_1158, Dm_build_599.Dm_build_789(dm_build_1159))
}

func (dm_build_1161 *Dm_build_953) Dm_build_1160(dm_build_1162 int, dm_build_1163 uint8) int {
	return dm_build_1161.Dm_build_1172(dm_build_1162, Dm_build_599.Dm_build_792(dm_build_1163))
}

func (dm_build_1165 *Dm_build_953) Dm_build_1164(dm_build_1166 int, dm_build_1167 uint16) int {
	return dm_build_1165.Dm_build_1172(dm_build_1166, Dm_build_599.Dm_build_795(dm_build_1167))
}

func (dm_build_1169 *Dm_build_953) Dm_build_1168(dm_build_1170 int, dm_build_1171 uint32) int {
	return dm_build_1169.Dm_build_1172(dm_build_1170, Dm_build_599.Dm_build_798(dm_build_1171))
}

func (dm_build_1173 *Dm_build_953) Dm_build_1172(dm_build_1174 int, dm_build_1175 []byte) int {
	return copy(dm_build_1173.dm_build_954[dm_build_1174:], dm_build_1175)
}

func (dm_build_1177 *Dm_build_953) Dm_build_1176(dm_build_1178 int, dm_build_1179 []byte) int {
	return dm_build_1177.Dm_build_1144(dm_build_1178, int32(len(dm_build_1179))) + dm_build_1177.Dm_build_1172(dm_build_1178+4, dm_build_1179)
}

func (dm_build_1181 *Dm_build_953) Dm_build_1180(dm_build_1182 int, dm_build_1183 []byte) int {
	return dm_build_1181.Dm_build_1136(dm_build_1182, byte(len(dm_build_1183))) + dm_build_1181.Dm_build_1172(dm_build_1182+1, dm_build_1183)
}

func (dm_build_1185 *Dm_build_953) Dm_build_1184(dm_build_1186 int, dm_build_1187 []byte) int {
	return dm_build_1185.Dm_build_1140(dm_build_1186, int16(len(dm_build_1187))) + dm_build_1185.Dm_build_1172(dm_build_1186+2, dm_build_1187)
}

func (dm_build_1189 *Dm_build_953) Dm_build_1188(dm_build_1190 int, dm_build_1191 []byte) int {
	return dm_build_1189.Dm_build_1172(dm_build_1190, dm_build_1191) + dm_build_1189.Dm_build_1136(dm_build_1190+len(dm_build_1191), 0)
}

func (dm_build_1193 *Dm_build_953) Dm_build_1192(dm_build_1194 int, dm_build_1195 string, dm_build_1196 string, dm_build_1197 *DmConnection) int {
	return dm_build_1193.Dm_build_1176(dm_build_1194, Dm_build_599.Dm_build_809(dm_build_1195, dm_build_1196, dm_build_1197))
}

func (dm_build_1199 *Dm_build_953) Dm_build_1198(dm_build_1200 int, dm_build_1201 string, dm_build_1202 string, dm_build_1203 *DmConnection) int {
	return dm_build_1199.Dm_build_1180(dm_build_1200, Dm_build_599.Dm_build_809(dm_build_1201, dm_build_1202, dm_build_1203))
}

func (dm_build_1205 *Dm_build_953) Dm_build_1204(dm_build_1206 int, dm_build_1207 string, dm_build_1208 string, dm_build_1209 *DmConnection) int {
	return dm_build_1205.Dm_build_1184(dm_build_1206, Dm_build_599.Dm_build_809(dm_build_1207, dm_build_1208, dm_build_1209))
}

func (dm_build_1211 *Dm_build_953) Dm_build_1210(dm_build_1212 int, dm_build_1213 string, dm_build_1214 string, dm_build_1215 *DmConnection) int {
	return dm_build_1211.Dm_build_1188(dm_build_1212, Dm_build_599.Dm_build_809(dm_build_1213, dm_build_1214, dm_build_1215))
}

func (dm_build_1217 *Dm_build_953) Dm_build_1216(dm_build_1218 int) byte {
	return Dm_build_599.Dm_build_814(dm_build_1217.Dm_build_1243(dm_build_1218, 1))
}

func (dm_build_1220 *Dm_build_953) Dm_build_1219(dm_build_1221 int) int16 {
	return Dm_build_599.Dm_build_817(dm_build_1220.Dm_build_1243(dm_build_1221, 2))
}

func (dm_build_1223 *Dm_build_953) Dm_build_1222(dm_build_1224 int) int32 {
	return Dm_build_599.Dm_build_820(dm_build_1223.Dm_build_1243(dm_build_1224, 4))
}

func (dm_build_1226 *Dm_build_953) Dm_build_1225(dm_build_1227 int) int64 {
	return Dm_build_599.Dm_build_823(dm_build_1226.Dm_build_1243(dm_build_1227, 8))
}

func (dm_build_1229 *Dm_build_953) Dm_build_1228(dm_build_1230 int) float32 {
	return Dm_build_599.Dm_build_826(dm_build_1229.Dm_build_1243(dm_build_1230, 4))
}

func (dm_build_1232 *Dm_build_953) Dm_build_1231(dm_build_1233 int) float64 {
	return Dm_build_599.Dm_build_829(dm_build_1232.Dm_build_1243(dm_build_1233, 8))
}

func (dm_build_1235 *Dm_build_953) Dm_build_1234(dm_build_1236 int) uint8 {
	return Dm_build_599.Dm_build_832(dm_build_1235.Dm_build_1243(dm_build_1236, 1))
}

func (dm_build_1238 *Dm_build_953) Dm_build_1237(dm_build_1239 int) uint16 {
	return Dm_build_599.Dm_build_835(dm_build_1238.Dm_build_1243(dm_build_1239, 2))
}

func (dm_build_1241 *Dm_build_953) Dm_build_1240(dm_build_1242 int) uint32 {
	return Dm_build_599.Dm_build_838(dm_build_1241.Dm_build_1243(dm_build_1242, 4))
}

func (dm_build_1244 *Dm_build_953) Dm_build_1243(dm_build_1245 int, dm_build_1246 int) []byte {
	return dm_build_1244.dm_build_954[dm_build_1245 : dm_build_1245+dm_build_1246]
}

func (dm_build_1248 *Dm_build_953) Dm_build_1247(dm_build_1249 int) []byte {
	dm_build_1250 := dm_build_1248.Dm_build_1222(dm_build_1249)
	return dm_build_1248.Dm_build_1243(dm_build_1249+4, int(dm_build_1250))
}

func (dm_build_1252 *Dm_build_953) Dm_build_1251(dm_build_1253 int) []byte {
	dm_build_1254 := dm_build_1252.Dm_build_1216(dm_build_1253)
	return dm_build_1252.Dm_build_1243(dm_build_1253+1, int(dm_build_1254))
}

func (dm_build_1256 *Dm_build_953) Dm_build_1255(dm_build_1257 int) []byte {
	dm_build_1258 := dm_build_1256.Dm_build_1219(dm_build_1257)
	return dm_build_1256.Dm_build_1243(dm_build_1257+2, int(dm_build_1258))
}

func (dm_build_1260 *Dm_build_953) Dm_build_1259(dm_build_1261 int) []byte {
	dm_build_1262 := 0
	for dm_build_1260.Dm_build_1216(dm_build_1261) != 0 {
		dm_build_1261++
		dm_build_1262++
	}

	return dm_build_1260.Dm_build_1243(dm_build_1261-dm_build_1262, int(dm_build_1262))
}

func (dm_build_1264 *Dm_build_953) Dm_build_1263(dm_build_1265 int, dm_build_1266 string, dm_build_1267 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1264.Dm_build_1247(dm_build_1265), dm_build_1266, dm_build_1267)
}

func (dm_build_1269 *Dm_build_953) Dm_build_1268(dm_build_1270 int, dm_build_1271 string, dm_build_1272 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1269.Dm_build_1251(dm_build_1270), dm_build_1271, dm_build_1272)
}

func (dm_build_1274 *Dm_build_953) Dm_build_1273(dm_build_1275 int, dm_build_1276 string, dm_build_1277 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1274.Dm_build_1255(dm_build_1275), dm_build_1276, dm_build_1277)
}

func (dm_build_1279 *Dm_build_953) Dm_build_1278(dm_build_1280 int, dm_build_1281 string, dm_build_1282 *DmConnection) string {
	return Dm_build_599.Dm_build_846(dm_build_1279.Dm_build_1259(dm_build_1280), dm_build_1281, dm_build_1282)
}
