/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm

import (
	"os"

	"github.com/gotomicro/dmgo/util"
)

var LogDirDef, _ = os.Getwd()

var StatDirDef, _ = os.Getwd()

const (
	DEFAULT_PORT int32 = 5236

	//log level
	LOG_OFF int = 0

	LOG_ERROR int = 1

	LOG_WARN int = 2

	LOG_SQL int = 3

	LOG_INFO int = 4

	LOG_DEBUG int = 5

	LOG_ALL int = 9

	//stat
	STAT_SQL_REMOVE_LATEST int = 0

	STAT_SQL_REMOVE_OLDEST int = 1

	// 编码字符集
	ENCODING_UTF8 string = "UTF-8"

	ENCODING_EUCKR string = "EUC-KR"

	ENCODING_GB18030 string = "GB18030"

	DbAliveCheckFreqDef = 0

	LocaleDef = 0

	// log
	LogLevelDef = LOG_OFF // 日志级别：off, error, warn, sql, info, all

	LogFlushFreqDef = 10 // 日志刷盘时间s (>=0)

	LogFlushQueueSizeDef = 100 //日志队列大小

	LogBufferSizeDef = 32 * 1024 // 日志缓冲区大小 (>0)

	// stat
	StatEnableDef = false //

	StatFlushFreqDef = 3 // 日志刷盘时间s (>=0)

	StatSlowSqlCountDef = 100 // 慢sql top行数，(0-1000)

	StatHighFreqSqlCountDef = 100 // 高频sql top行数， (0-1000)

	StatSqlMaxCountDef = 100000 // sql 统计最大值(0-100000)

	StatSqlRemoveModeDef = STAT_SQL_REMOVE_LATEST // 记录sql数超过最大值时，sql淘汰方式
)

var (
	DbAliveCheckFreq = DbAliveCheckFreqDef

	Locale = LocaleDef // 0:简体中文 1：英文 2:繁体中文

	// log
	LogLevel = LogLevelDef // 日志级别：off, error, warn, sql, info, all

	LogDir = LogDirDef

	LogFlushFreq = LogFlushFreqDef // 日志刷盘时间s (>=0)

	LogFlushQueueSize = LogFlushQueueSizeDef

	LogBufferSize = LogBufferSizeDef // 日志缓冲区大小 (>0)

	// stat
	StatEnable = StatEnableDef //

	StatDir = StatDirDef // jdbc工作目录,所有生成的文件都在该目录下

	StatFlushFreq = StatFlushFreqDef // 日志刷盘时间s (>=0)

	StatSlowSqlCount = StatSlowSqlCountDef // 慢sql top行数，(0-1000)

	StatHighFreqSqlCount = StatHighFreqSqlCountDef // 高频sql top行数， (0-1000)

	StatSqlMaxCount = StatSqlMaxCountDef // sql 统计最大值(0-100000)

	StatSqlRemoveMode = StatSqlRemoveModeDef // 记录sql数超过最大值时，sql淘汰方式

	/*---------------------------------------------------------------*/
	ServerGroupMap = make(map[string]*epGroup)
)

func ParseLogLevel(props *Properties) int {
	logLevel := LOG_OFF
	value := props.GetString(LogLevelKey, "")
	if value != "" && !util.StringUtil.IsDigit(value) {
		if util.StringUtil.EqualsIgnoreCase("debug", value) {
			logLevel = LOG_DEBUG
		} else if util.StringUtil.EqualsIgnoreCase("info", value) {
			logLevel = LOG_INFO
		} else if util.StringUtil.EqualsIgnoreCase("sql", value) {
			logLevel = LOG_SQL
		} else if util.StringUtil.EqualsIgnoreCase("warn", value) {
			logLevel = LOG_WARN
		} else if util.StringUtil.EqualsIgnoreCase("error", value) {
			logLevel = LOG_ERROR
		} else if util.StringUtil.EqualsIgnoreCase("off", value) {
			logLevel = LOG_OFF
		} else if util.StringUtil.EqualsIgnoreCase("all", value) {
			logLevel = LOG_ALL
		}
	} else {
		logLevel = props.GetInt(LogLevelKey, logLevel, LOG_OFF, LOG_INFO)
	}

	return logLevel
}
