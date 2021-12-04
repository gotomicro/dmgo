/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_980 struct {
	dm_build_981 []byte
	dm_build_982 int
}

func Dm_build_983(dm_build_984 int) *Dm_build_980 {
	return &Dm_build_980{make([]byte, 0, dm_build_984), 0}
}

func Dm_build_985(dm_build_986 []byte) *Dm_build_980 {
	return &Dm_build_980{dm_build_986, 0}
}

func (dm_build_988 *Dm_build_980) dm_build_987(dm_build_989 int) *Dm_build_980 {

	dm_build_990 := len(dm_build_988.dm_build_981)
	dm_build_991 := cap(dm_build_988.dm_build_981)

	if dm_build_990+dm_build_989 <= dm_build_991 {
		dm_build_988.dm_build_981 = dm_build_988.dm_build_981[:dm_build_990+dm_build_989]
	} else {
		remain := dm_build_989 + dm_build_990 - dm_build_991
		nbuf := make([]byte, dm_build_989+dm_build_990, 2*dm_build_991+remain)
		copy(nbuf, dm_build_988.dm_build_981)
		dm_build_988.dm_build_981 = nbuf
	}

	return dm_build_988
}

func (dm_build_993 *Dm_build_980) Dm_build_992() int {
	return len(dm_build_993.dm_build_981)
}

func (dm_build_995 *Dm_build_980) Dm_build_994(dm_build_996 int) *Dm_build_980 {
	for i := dm_build_996; i < len(dm_build_995.dm_build_981); i++ {
		dm_build_995.dm_build_981[i] = 0
	}
	dm_build_995.dm_build_981 = dm_build_995.dm_build_981[:dm_build_996]
	return dm_build_995
}

func (dm_build_998 *Dm_build_980) Dm_build_997(dm_build_999 int) *Dm_build_980 {
	dm_build_998.dm_build_982 = dm_build_999
	return dm_build_998
}

func (dm_build_1001 *Dm_build_980) Dm_build_1000() int {
	return dm_build_1001.dm_build_982
}

func (dm_build_1003 *Dm_build_980) Dm_build_1002(dm_build_1004 bool) int {
	return len(dm_build_1003.dm_build_981) - dm_build_1003.dm_build_982
}

func (dm_build_1006 *Dm_build_980) Dm_build_1005(dm_build_1007 int, dm_build_1008 bool, dm_build_1009 bool) *Dm_build_980 {

	if dm_build_1008 {
		if dm_build_1009 {
			dm_build_1006.dm_build_987(dm_build_1007)
		} else {
			dm_build_1006.dm_build_981 = dm_build_1006.dm_build_981[:len(dm_build_1006.dm_build_981)-dm_build_1007]
		}
	} else {
		if dm_build_1009 {
			dm_build_1006.dm_build_982 += dm_build_1007
		} else {
			dm_build_1006.dm_build_982 -= dm_build_1007
		}
	}

	return dm_build_1006
}

func (dm_build_1011 *Dm_build_980) Dm_build_1010(dm_build_1012 io.Reader, dm_build_1013 int) int {
	dm_build_1014 := len(dm_build_1011.dm_build_981)
	dm_build_1011.dm_build_987(dm_build_1013)
	dm_build_1015 := 0
	for dm_build_1013 > 0 {
		n, err := dm_build_1012.Read(dm_build_1011.dm_build_981[dm_build_1014+dm_build_1015:])
		if n > 0 && err == io.EOF {
			dm_build_1015 += n
			dm_build_1011.dm_build_981 = dm_build_1011.dm_build_981[:dm_build_1014+dm_build_1015]
			return dm_build_1015
		} else if n > 0 && err == nil {
			dm_build_1013 -= n
			dm_build_1015 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_1015
}

func (dm_build_1017 *Dm_build_980) Dm_build_1016(dm_build_1018 io.Writer) *Dm_build_980 {
	dm_build_1018.Write(dm_build_1017.dm_build_981)
	return dm_build_1017
}

func (dm_build_1020 *Dm_build_980) Dm_build_1019(dm_build_1021 bool) int {
	dm_build_1022 := len(dm_build_1020.dm_build_981)
	dm_build_1020.dm_build_987(1)

	if dm_build_1021 {
		return copy(dm_build_1020.dm_build_981[dm_build_1022:], []byte{1})
	} else {
		return copy(dm_build_1020.dm_build_981[dm_build_1022:], []byte{0})
	}
}

func (dm_build_1024 *Dm_build_980) Dm_build_1023(dm_build_1025 byte) int {
	dm_build_1026 := len(dm_build_1024.dm_build_981)
	dm_build_1024.dm_build_987(1)

	return copy(dm_build_1024.dm_build_981[dm_build_1026:], Dm_build_623.Dm_build_801(dm_build_1025))
}

func (dm_build_1028 *Dm_build_980) Dm_build_1027(dm_build_1029 int16) int {
	dm_build_1030 := len(dm_build_1028.dm_build_981)
	dm_build_1028.dm_build_987(2)

	return copy(dm_build_1028.dm_build_981[dm_build_1030:], Dm_build_623.Dm_build_804(dm_build_1029))
}

func (dm_build_1032 *Dm_build_980) Dm_build_1031(dm_build_1033 int32) int {
	dm_build_1034 := len(dm_build_1032.dm_build_981)
	dm_build_1032.dm_build_987(4)

	return copy(dm_build_1032.dm_build_981[dm_build_1034:], Dm_build_623.Dm_build_807(dm_build_1033))
}

func (dm_build_1036 *Dm_build_980) Dm_build_1035(dm_build_1037 uint8) int {
	dm_build_1038 := len(dm_build_1036.dm_build_981)
	dm_build_1036.dm_build_987(1)

	return copy(dm_build_1036.dm_build_981[dm_build_1038:], Dm_build_623.Dm_build_819(dm_build_1037))
}

func (dm_build_1040 *Dm_build_980) Dm_build_1039(dm_build_1041 uint16) int {
	dm_build_1042 := len(dm_build_1040.dm_build_981)
	dm_build_1040.dm_build_987(2)

	return copy(dm_build_1040.dm_build_981[dm_build_1042:], Dm_build_623.Dm_build_822(dm_build_1041))
}

func (dm_build_1044 *Dm_build_980) Dm_build_1043(dm_build_1045 uint32) int {
	dm_build_1046 := len(dm_build_1044.dm_build_981)
	dm_build_1044.dm_build_987(4)

	return copy(dm_build_1044.dm_build_981[dm_build_1046:], Dm_build_623.Dm_build_825(dm_build_1045))
}

func (dm_build_1048 *Dm_build_980) Dm_build_1047(dm_build_1049 uint64) int {
	dm_build_1050 := len(dm_build_1048.dm_build_981)
	dm_build_1048.dm_build_987(8)

	return copy(dm_build_1048.dm_build_981[dm_build_1050:], Dm_build_623.Dm_build_828(dm_build_1049))
}

func (dm_build_1052 *Dm_build_980) Dm_build_1051(dm_build_1053 float32) int {
	dm_build_1054 := len(dm_build_1052.dm_build_981)
	dm_build_1052.dm_build_987(4)

	return copy(dm_build_1052.dm_build_981[dm_build_1054:], Dm_build_623.Dm_build_825(math.Float32bits(dm_build_1053)))
}

func (dm_build_1056 *Dm_build_980) Dm_build_1055(dm_build_1057 float64) int {
	dm_build_1058 := len(dm_build_1056.dm_build_981)
	dm_build_1056.dm_build_987(8)

	return copy(dm_build_1056.dm_build_981[dm_build_1058:], Dm_build_623.Dm_build_828(math.Float64bits(dm_build_1057)))
}

func (dm_build_1060 *Dm_build_980) Dm_build_1059(dm_build_1061 []byte) int {
	dm_build_1062 := len(dm_build_1060.dm_build_981)
	dm_build_1060.dm_build_987(len(dm_build_1061))
	return copy(dm_build_1060.dm_build_981[dm_build_1062:], dm_build_1061)
}

func (dm_build_1064 *Dm_build_980) Dm_build_1063(dm_build_1065 []byte) int {
	return dm_build_1064.Dm_build_1031(int32(len(dm_build_1065))) + dm_build_1064.Dm_build_1059(dm_build_1065)
}

func (dm_build_1067 *Dm_build_980) Dm_build_1066(dm_build_1068 []byte) int {
	return dm_build_1067.Dm_build_1035(uint8(len(dm_build_1068))) + dm_build_1067.Dm_build_1059(dm_build_1068)
}

func (dm_build_1070 *Dm_build_980) Dm_build_1069(dm_build_1071 []byte) int {
	return dm_build_1070.Dm_build_1039(uint16(len(dm_build_1071))) + dm_build_1070.Dm_build_1059(dm_build_1071)
}

func (dm_build_1073 *Dm_build_980) Dm_build_1072(dm_build_1074 []byte) int {
	return dm_build_1073.Dm_build_1059(dm_build_1074) + dm_build_1073.Dm_build_1023(0)
}

func (dm_build_1076 *Dm_build_980) Dm_build_1075(dm_build_1077 string, dm_build_1078 string, dm_build_1079 *DmConnection) int {
	dm_build_1080 := Dm_build_623.Dm_build_836(dm_build_1077, dm_build_1078, dm_build_1079)
	return dm_build_1076.Dm_build_1063(dm_build_1080)
}

func (dm_build_1082 *Dm_build_980) Dm_build_1081(dm_build_1083 string, dm_build_1084 string, dm_build_1085 *DmConnection) int {
	dm_build_1086 := Dm_build_623.Dm_build_836(dm_build_1083, dm_build_1084, dm_build_1085)
	return dm_build_1082.Dm_build_1066(dm_build_1086)
}

func (dm_build_1088 *Dm_build_980) Dm_build_1087(dm_build_1089 string, dm_build_1090 string, dm_build_1091 *DmConnection) int {
	dm_build_1092 := Dm_build_623.Dm_build_836(dm_build_1089, dm_build_1090, dm_build_1091)
	return dm_build_1088.Dm_build_1069(dm_build_1092)
}

func (dm_build_1094 *Dm_build_980) Dm_build_1093(dm_build_1095 string, dm_build_1096 string, dm_build_1097 *DmConnection) int {
	dm_build_1098 := Dm_build_623.Dm_build_836(dm_build_1095, dm_build_1096, dm_build_1097)
	return dm_build_1094.Dm_build_1072(dm_build_1098)
}

func (dm_build_1100 *Dm_build_980) Dm_build_1099() byte {
	dm_build_1101 := Dm_build_623.Dm_build_716(dm_build_1100.dm_build_981, dm_build_1100.dm_build_982)
	dm_build_1100.dm_build_982++
	return dm_build_1101
}

func (dm_build_1103 *Dm_build_980) Dm_build_1102() int16 {
	dm_build_1104 := Dm_build_623.Dm_build_720(dm_build_1103.dm_build_981, dm_build_1103.dm_build_982)
	dm_build_1103.dm_build_982 += 2
	return dm_build_1104
}

func (dm_build_1106 *Dm_build_980) Dm_build_1105() int32 {
	dm_build_1107 := Dm_build_623.Dm_build_725(dm_build_1106.dm_build_981, dm_build_1106.dm_build_982)
	dm_build_1106.dm_build_982 += 4
	return dm_build_1107
}

func (dm_build_1109 *Dm_build_980) Dm_build_1108() int64 {
	dm_build_1110 := Dm_build_623.Dm_build_730(dm_build_1109.dm_build_981, dm_build_1109.dm_build_982)
	dm_build_1109.dm_build_982 += 8
	return dm_build_1110
}

func (dm_build_1112 *Dm_build_980) Dm_build_1111() float32 {
	dm_build_1113 := Dm_build_623.Dm_build_735(dm_build_1112.dm_build_981, dm_build_1112.dm_build_982)
	dm_build_1112.dm_build_982 += 4
	return dm_build_1113
}

func (dm_build_1115 *Dm_build_980) Dm_build_1114() float64 {
	dm_build_1116 := Dm_build_623.Dm_build_739(dm_build_1115.dm_build_981, dm_build_1115.dm_build_982)
	dm_build_1115.dm_build_982 += 8
	return dm_build_1116
}

func (dm_build_1118 *Dm_build_980) Dm_build_1117() uint8 {
	dm_build_1119 := Dm_build_623.Dm_build_743(dm_build_1118.dm_build_981, dm_build_1118.dm_build_982)
	dm_build_1118.dm_build_982 += 1
	return dm_build_1119
}

func (dm_build_1121 *Dm_build_980) Dm_build_1120() uint16 {
	dm_build_1122 := Dm_build_623.Dm_build_747(dm_build_1121.dm_build_981, dm_build_1121.dm_build_982)
	dm_build_1121.dm_build_982 += 2
	return dm_build_1122
}

func (dm_build_1124 *Dm_build_980) Dm_build_1123() uint32 {
	dm_build_1125 := Dm_build_623.Dm_build_752(dm_build_1124.dm_build_981, dm_build_1124.dm_build_982)
	dm_build_1124.dm_build_982 += 4
	return dm_build_1125
}

func (dm_build_1127 *Dm_build_980) Dm_build_1126(dm_build_1128 int) []byte {
	dm_build_1129 := Dm_build_623.Dm_build_774(dm_build_1127.dm_build_981, dm_build_1127.dm_build_982, dm_build_1128)
	dm_build_1127.dm_build_982 += dm_build_1128
	return dm_build_1129
}

func (dm_build_1131 *Dm_build_980) Dm_build_1130() []byte {
	return dm_build_1131.Dm_build_1126(int(dm_build_1131.Dm_build_1105()))
}

func (dm_build_1133 *Dm_build_980) Dm_build_1132() []byte {
	return dm_build_1133.Dm_build_1126(int(dm_build_1133.Dm_build_1099()))
}

func (dm_build_1135 *Dm_build_980) Dm_build_1134() []byte {
	return dm_build_1135.Dm_build_1126(int(dm_build_1135.Dm_build_1102()))
}

func (dm_build_1137 *Dm_build_980) Dm_build_1136(dm_build_1138 int) []byte {
	return dm_build_1137.Dm_build_1126(dm_build_1138)
}

func (dm_build_1140 *Dm_build_980) Dm_build_1139() []byte {
	dm_build_1141 := 0
	for dm_build_1140.Dm_build_1099() != 0 {
		dm_build_1141++
	}
	dm_build_1140.Dm_build_1005(dm_build_1141, false, false)
	return dm_build_1140.Dm_build_1126(dm_build_1141)
}

func (dm_build_1143 *Dm_build_980) Dm_build_1142(dm_build_1144 int, dm_build_1145 string, dm_build_1146 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1143.Dm_build_1126(dm_build_1144), dm_build_1145, dm_build_1146)
}

func (dm_build_1148 *Dm_build_980) Dm_build_1147(dm_build_1149 string, dm_build_1150 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1148.Dm_build_1130(), dm_build_1149, dm_build_1150)
}

func (dm_build_1152 *Dm_build_980) Dm_build_1151(dm_build_1153 string, dm_build_1154 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1152.Dm_build_1132(), dm_build_1153, dm_build_1154)
}

func (dm_build_1156 *Dm_build_980) Dm_build_1155(dm_build_1157 string, dm_build_1158 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1156.Dm_build_1134(), dm_build_1157, dm_build_1158)
}

func (dm_build_1160 *Dm_build_980) Dm_build_1159(dm_build_1161 string, dm_build_1162 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1160.Dm_build_1139(), dm_build_1161, dm_build_1162)
}

func (dm_build_1164 *Dm_build_980) Dm_build_1163(dm_build_1165 int, dm_build_1166 byte) int {
	return dm_build_1164.Dm_build_1199(dm_build_1165, Dm_build_623.Dm_build_801(dm_build_1166))
}

func (dm_build_1168 *Dm_build_980) Dm_build_1167(dm_build_1169 int, dm_build_1170 int16) int {
	return dm_build_1168.Dm_build_1199(dm_build_1169, Dm_build_623.Dm_build_804(dm_build_1170))
}

func (dm_build_1172 *Dm_build_980) Dm_build_1171(dm_build_1173 int, dm_build_1174 int32) int {
	return dm_build_1172.Dm_build_1199(dm_build_1173, Dm_build_623.Dm_build_807(dm_build_1174))
}

func (dm_build_1176 *Dm_build_980) Dm_build_1175(dm_build_1177 int, dm_build_1178 int64) int {
	return dm_build_1176.Dm_build_1199(dm_build_1177, Dm_build_623.Dm_build_810(dm_build_1178))
}

func (dm_build_1180 *Dm_build_980) Dm_build_1179(dm_build_1181 int, dm_build_1182 float32) int {
	return dm_build_1180.Dm_build_1199(dm_build_1181, Dm_build_623.Dm_build_813(dm_build_1182))
}

func (dm_build_1184 *Dm_build_980) Dm_build_1183(dm_build_1185 int, dm_build_1186 float64) int {
	return dm_build_1184.Dm_build_1199(dm_build_1185, Dm_build_623.Dm_build_816(dm_build_1186))
}

func (dm_build_1188 *Dm_build_980) Dm_build_1187(dm_build_1189 int, dm_build_1190 uint8) int {
	return dm_build_1188.Dm_build_1199(dm_build_1189, Dm_build_623.Dm_build_819(dm_build_1190))
}

func (dm_build_1192 *Dm_build_980) Dm_build_1191(dm_build_1193 int, dm_build_1194 uint16) int {
	return dm_build_1192.Dm_build_1199(dm_build_1193, Dm_build_623.Dm_build_822(dm_build_1194))
}

func (dm_build_1196 *Dm_build_980) Dm_build_1195(dm_build_1197 int, dm_build_1198 uint32) int {
	return dm_build_1196.Dm_build_1199(dm_build_1197, Dm_build_623.Dm_build_825(dm_build_1198))
}

func (dm_build_1200 *Dm_build_980) Dm_build_1199(dm_build_1201 int, dm_build_1202 []byte) int {
	return copy(dm_build_1200.dm_build_981[dm_build_1201:], dm_build_1202)
}

func (dm_build_1204 *Dm_build_980) Dm_build_1203(dm_build_1205 int, dm_build_1206 []byte) int {
	return dm_build_1204.Dm_build_1171(dm_build_1205, int32(len(dm_build_1206))) + dm_build_1204.Dm_build_1199(dm_build_1205+4, dm_build_1206)
}

func (dm_build_1208 *Dm_build_980) Dm_build_1207(dm_build_1209 int, dm_build_1210 []byte) int {
	return dm_build_1208.Dm_build_1163(dm_build_1209, byte(len(dm_build_1210))) + dm_build_1208.Dm_build_1199(dm_build_1209+1, dm_build_1210)
}

func (dm_build_1212 *Dm_build_980) Dm_build_1211(dm_build_1213 int, dm_build_1214 []byte) int {
	return dm_build_1212.Dm_build_1167(dm_build_1213, int16(len(dm_build_1214))) + dm_build_1212.Dm_build_1199(dm_build_1213+2, dm_build_1214)
}

func (dm_build_1216 *Dm_build_980) Dm_build_1215(dm_build_1217 int, dm_build_1218 []byte) int {
	return dm_build_1216.Dm_build_1199(dm_build_1217, dm_build_1218) + dm_build_1216.Dm_build_1163(dm_build_1217+len(dm_build_1218), 0)
}

func (dm_build_1220 *Dm_build_980) Dm_build_1219(dm_build_1221 int, dm_build_1222 string, dm_build_1223 string, dm_build_1224 *DmConnection) int {
	return dm_build_1220.Dm_build_1203(dm_build_1221, Dm_build_623.Dm_build_836(dm_build_1222, dm_build_1223, dm_build_1224))
}

func (dm_build_1226 *Dm_build_980) Dm_build_1225(dm_build_1227 int, dm_build_1228 string, dm_build_1229 string, dm_build_1230 *DmConnection) int {
	return dm_build_1226.Dm_build_1207(dm_build_1227, Dm_build_623.Dm_build_836(dm_build_1228, dm_build_1229, dm_build_1230))
}

func (dm_build_1232 *Dm_build_980) Dm_build_1231(dm_build_1233 int, dm_build_1234 string, dm_build_1235 string, dm_build_1236 *DmConnection) int {
	return dm_build_1232.Dm_build_1211(dm_build_1233, Dm_build_623.Dm_build_836(dm_build_1234, dm_build_1235, dm_build_1236))
}

func (dm_build_1238 *Dm_build_980) Dm_build_1237(dm_build_1239 int, dm_build_1240 string, dm_build_1241 string, dm_build_1242 *DmConnection) int {
	return dm_build_1238.Dm_build_1215(dm_build_1239, Dm_build_623.Dm_build_836(dm_build_1240, dm_build_1241, dm_build_1242))
}

func (dm_build_1244 *Dm_build_980) Dm_build_1243(dm_build_1245 int) byte {
	return Dm_build_623.Dm_build_841(dm_build_1244.Dm_build_1270(dm_build_1245, 1))
}

func (dm_build_1247 *Dm_build_980) Dm_build_1246(dm_build_1248 int) int16 {
	return Dm_build_623.Dm_build_844(dm_build_1247.Dm_build_1270(dm_build_1248, 2))
}

func (dm_build_1250 *Dm_build_980) Dm_build_1249(dm_build_1251 int) int32 {
	return Dm_build_623.Dm_build_847(dm_build_1250.Dm_build_1270(dm_build_1251, 4))
}

func (dm_build_1253 *Dm_build_980) Dm_build_1252(dm_build_1254 int) int64 {
	return Dm_build_623.Dm_build_850(dm_build_1253.Dm_build_1270(dm_build_1254, 8))
}

func (dm_build_1256 *Dm_build_980) Dm_build_1255(dm_build_1257 int) float32 {
	return Dm_build_623.Dm_build_853(dm_build_1256.Dm_build_1270(dm_build_1257, 4))
}

func (dm_build_1259 *Dm_build_980) Dm_build_1258(dm_build_1260 int) float64 {
	return Dm_build_623.Dm_build_856(dm_build_1259.Dm_build_1270(dm_build_1260, 8))
}

func (dm_build_1262 *Dm_build_980) Dm_build_1261(dm_build_1263 int) uint8 {
	return Dm_build_623.Dm_build_859(dm_build_1262.Dm_build_1270(dm_build_1263, 1))
}

func (dm_build_1265 *Dm_build_980) Dm_build_1264(dm_build_1266 int) uint16 {
	return Dm_build_623.Dm_build_862(dm_build_1265.Dm_build_1270(dm_build_1266, 2))
}

func (dm_build_1268 *Dm_build_980) Dm_build_1267(dm_build_1269 int) uint32 {
	return Dm_build_623.Dm_build_865(dm_build_1268.Dm_build_1270(dm_build_1269, 4))
}

func (dm_build_1271 *Dm_build_980) Dm_build_1270(dm_build_1272 int, dm_build_1273 int) []byte {
	return dm_build_1271.dm_build_981[dm_build_1272 : dm_build_1272+dm_build_1273]
}

func (dm_build_1275 *Dm_build_980) Dm_build_1274(dm_build_1276 int) []byte {
	dm_build_1277 := dm_build_1275.Dm_build_1249(dm_build_1276)
	return dm_build_1275.Dm_build_1270(dm_build_1276+4, int(dm_build_1277))
}

func (dm_build_1279 *Dm_build_980) Dm_build_1278(dm_build_1280 int) []byte {
	dm_build_1281 := dm_build_1279.Dm_build_1243(dm_build_1280)
	return dm_build_1279.Dm_build_1270(dm_build_1280+1, int(dm_build_1281))
}

func (dm_build_1283 *Dm_build_980) Dm_build_1282(dm_build_1284 int) []byte {
	dm_build_1285 := dm_build_1283.Dm_build_1246(dm_build_1284)
	return dm_build_1283.Dm_build_1270(dm_build_1284+2, int(dm_build_1285))
}

func (dm_build_1287 *Dm_build_980) Dm_build_1286(dm_build_1288 int) []byte {
	dm_build_1289 := 0
	for dm_build_1287.Dm_build_1243(dm_build_1288) != 0 {
		dm_build_1288++
		dm_build_1289++
	}

	return dm_build_1287.Dm_build_1270(dm_build_1288-dm_build_1289, int(dm_build_1289))
}

func (dm_build_1291 *Dm_build_980) Dm_build_1290(dm_build_1292 int, dm_build_1293 string, dm_build_1294 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1291.Dm_build_1274(dm_build_1292), dm_build_1293, dm_build_1294)
}

func (dm_build_1296 *Dm_build_980) Dm_build_1295(dm_build_1297 int, dm_build_1298 string, dm_build_1299 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1296.Dm_build_1278(dm_build_1297), dm_build_1298, dm_build_1299)
}

func (dm_build_1301 *Dm_build_980) Dm_build_1300(dm_build_1302 int, dm_build_1303 string, dm_build_1304 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1301.Dm_build_1282(dm_build_1302), dm_build_1303, dm_build_1304)
}

func (dm_build_1306 *Dm_build_980) Dm_build_1305(dm_build_1307 int, dm_build_1308 string, dm_build_1309 *DmConnection) string {
	return Dm_build_623.Dm_build_873(dm_build_1306.Dm_build_1286(dm_build_1307), dm_build_1308, dm_build_1309)
}
