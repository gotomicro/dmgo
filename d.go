/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1161 struct {
	dm_build_1162 *list.List
	dm_build_1163 *dm_build_1215
	dm_build_1164 int
}

func Dm_build_1165() *Dm_build_1161 {
	return &Dm_build_1161{
		dm_build_1162: list.New(),
		dm_build_1164: 0,
	}
}

func (dm_build_1167 *Dm_build_1161) Dm_build_1166() int {
	return dm_build_1167.dm_build_1164
}

func (dm_build_1169 *Dm_build_1161) Dm_build_1168(dm_build_1170 *Dm_build_1239, dm_build_1171 int) int {
	var dm_build_1172 = 0
	var dm_build_1173 = 0
	for dm_build_1172 < dm_build_1171 && dm_build_1169.dm_build_1163 != nil {
		dm_build_1173 = dm_build_1169.dm_build_1163.dm_build_1223(dm_build_1170, dm_build_1171-dm_build_1172)
		if dm_build_1169.dm_build_1163.dm_build_1218 == 0 {
			dm_build_1169.dm_build_1205()
		}
		dm_build_1172 += dm_build_1173
		dm_build_1169.dm_build_1164 -= dm_build_1173
	}
	return dm_build_1172
}

func (dm_build_1175 *Dm_build_1161) Dm_build_1174(dm_build_1176 []byte, dm_build_1177 int, dm_build_1178 int) int {
	var dm_build_1179 = 0
	var dm_build_1180 = 0
	for dm_build_1179 < dm_build_1178 && dm_build_1175.dm_build_1163 != nil {
		dm_build_1180 = dm_build_1175.dm_build_1163.dm_build_1227(dm_build_1176, dm_build_1177, dm_build_1178-dm_build_1179)
		if dm_build_1175.dm_build_1163.dm_build_1218 == 0 {
			dm_build_1175.dm_build_1205()
		}
		dm_build_1179 += dm_build_1180
		dm_build_1175.dm_build_1164 -= dm_build_1180
		dm_build_1177 += dm_build_1180
	}
	return dm_build_1179
}

func (dm_build_1182 *Dm_build_1161) Dm_build_1181(dm_build_1183 io.Writer, dm_build_1184 int) int {
	var dm_build_1185 = 0
	var dm_build_1186 = 0
	for dm_build_1185 < dm_build_1184 && dm_build_1182.dm_build_1163 != nil {
		dm_build_1186 = dm_build_1182.dm_build_1163.dm_build_1232(dm_build_1183, dm_build_1184-dm_build_1185)
		if dm_build_1182.dm_build_1163.dm_build_1218 == 0 {
			dm_build_1182.dm_build_1205()
		}
		dm_build_1185 += dm_build_1186
		dm_build_1182.dm_build_1164 -= dm_build_1186
	}
	return dm_build_1185
}

func (dm_build_1188 *Dm_build_1161) Dm_build_1187(dm_build_1189 []byte, dm_build_1190 int, dm_build_1191 int) {
	if dm_build_1191 == 0 {
		return
	}
	var dm_build_1192 = dm_build_1219(dm_build_1189, dm_build_1190, dm_build_1191)
	if dm_build_1188.dm_build_1163 == nil {
		dm_build_1188.dm_build_1163 = dm_build_1192
	} else {
		dm_build_1188.dm_build_1162.PushBack(dm_build_1192)
	}
	dm_build_1188.dm_build_1164 += dm_build_1191
}

func (dm_build_1194 *Dm_build_1161) dm_build_1193(dm_build_1195 int) byte {
	var dm_build_1196 = dm_build_1195
	var dm_build_1197 = dm_build_1194.dm_build_1163
	for dm_build_1196 > 0 && dm_build_1197 != nil {
		if dm_build_1197.dm_build_1218 == 0 {
			continue
		}
		if dm_build_1196 > dm_build_1197.dm_build_1218-1 {
			dm_build_1196 -= dm_build_1197.dm_build_1218
			dm_build_1197 = dm_build_1194.dm_build_1162.Front().Value.(*dm_build_1215)
		} else {
			break
		}
	}
	return dm_build_1197.dm_build_1236(dm_build_1196)
}
func (dm_build_1199 *Dm_build_1161) Dm_build_1198(dm_build_1200 *Dm_build_1161) {
	if dm_build_1200.dm_build_1164 == 0 {
		return
	}
	var dm_build_1201 = dm_build_1200.dm_build_1163
	for dm_build_1201 != nil {
		dm_build_1199.dm_build_1202(dm_build_1201)
		dm_build_1200.dm_build_1205()
		dm_build_1201 = dm_build_1200.dm_build_1163
	}
	dm_build_1200.dm_build_1164 = 0
}
func (dm_build_1203 *Dm_build_1161) dm_build_1202(dm_build_1204 *dm_build_1215) {
	if dm_build_1204.dm_build_1218 == 0 {
		return
	}
	if dm_build_1203.dm_build_1163 == nil {
		dm_build_1203.dm_build_1163 = dm_build_1204
	} else {
		dm_build_1203.dm_build_1162.PushBack(dm_build_1204)
	}
	dm_build_1203.dm_build_1164 += dm_build_1204.dm_build_1218
}

func (dm_build_1206 *Dm_build_1161) dm_build_1205() {
	var dm_build_1207 = dm_build_1206.dm_build_1162.Front()
	if dm_build_1207 == nil {
		dm_build_1206.dm_build_1163 = nil
	} else {
		dm_build_1206.dm_build_1163 = dm_build_1207.Value.(*dm_build_1215)
		dm_build_1206.dm_build_1162.Remove(dm_build_1207)
	}
}

func (dm_build_1209 *Dm_build_1161) Dm_build_1208() []byte {
	var dm_build_1210 = make([]byte, dm_build_1209.dm_build_1164)
	var dm_build_1211 = dm_build_1209.dm_build_1163
	var dm_build_1212 = 0
	var dm_build_1213 = len(dm_build_1210)
	var dm_build_1214 = 0
	for dm_build_1211 != nil {
		if dm_build_1211.dm_build_1218 > 0 {
			if dm_build_1213 > dm_build_1211.dm_build_1218 {
				dm_build_1214 = dm_build_1211.dm_build_1218
			} else {
				dm_build_1214 = dm_build_1213
			}
			copy(dm_build_1210[dm_build_1212:dm_build_1212+dm_build_1214], dm_build_1211.dm_build_1216[dm_build_1211.dm_build_1217:dm_build_1211.dm_build_1217+dm_build_1214])
			dm_build_1212 += dm_build_1214
			dm_build_1213 -= dm_build_1214
		}
		if dm_build_1209.dm_build_1162.Front() == nil {
			dm_build_1211 = nil
		} else {
			dm_build_1211 = dm_build_1209.dm_build_1162.Front().Value.(*dm_build_1215)
		}
	}
	return dm_build_1210
}

type dm_build_1215 struct {
	dm_build_1216 []byte
	dm_build_1217 int
	dm_build_1218 int
}

func dm_build_1219(dm_build_1220 []byte, dm_build_1221 int, dm_build_1222 int) *dm_build_1215 {
	return &dm_build_1215{
		dm_build_1220,
		dm_build_1221,
		dm_build_1222,
	}
}

func (dm_build_1224 *dm_build_1215) dm_build_1223(dm_build_1225 *Dm_build_1239, dm_build_1226 int) int {
	if dm_build_1224.dm_build_1218 <= dm_build_1226 {
		dm_build_1226 = dm_build_1224.dm_build_1218
	}
	dm_build_1225.Dm_build_1318(dm_build_1224.dm_build_1216[dm_build_1224.dm_build_1217 : dm_build_1224.dm_build_1217+dm_build_1226])
	dm_build_1224.dm_build_1217 += dm_build_1226
	dm_build_1224.dm_build_1218 -= dm_build_1226
	return dm_build_1226
}

func (dm_build_1228 *dm_build_1215) dm_build_1227(dm_build_1229 []byte, dm_build_1230 int, dm_build_1231 int) int {
	if dm_build_1228.dm_build_1218 <= dm_build_1231 {
		dm_build_1231 = dm_build_1228.dm_build_1218
	}
	copy(dm_build_1229[dm_build_1230:dm_build_1230+dm_build_1231], dm_build_1228.dm_build_1216[dm_build_1228.dm_build_1217:dm_build_1228.dm_build_1217+dm_build_1231])
	dm_build_1228.dm_build_1217 += dm_build_1231
	dm_build_1228.dm_build_1218 -= dm_build_1231
	return dm_build_1231
}

func (dm_build_1233 *dm_build_1215) dm_build_1232(dm_build_1234 io.Writer, dm_build_1235 int) int {
	if dm_build_1233.dm_build_1218 <= dm_build_1235 {
		dm_build_1235 = dm_build_1233.dm_build_1218
	}
	dm_build_1234.Write(dm_build_1233.dm_build_1216[dm_build_1233.dm_build_1217 : dm_build_1233.dm_build_1217+dm_build_1235])
	dm_build_1233.dm_build_1217 += dm_build_1235
	dm_build_1233.dm_build_1218 -= dm_build_1235
	return dm_build_1235
}
func (dm_build_1237 *dm_build_1215) dm_build_1236(dm_build_1238 int) byte {
	return dm_build_1237.dm_build_1216[dm_build_1237.dm_build_1217+dm_build_1238]
}
