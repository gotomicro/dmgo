/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_902 struct {
	dm_build_903 *list.List
	dm_build_904 *dm_build_956
	dm_build_905 int
}

func Dm_build_906() *Dm_build_902 {
	return &Dm_build_902{
		dm_build_903: list.New(),
		dm_build_905: 0,
	}
}

func (dm_build_908 *Dm_build_902) Dm_build_907() int {
	return dm_build_908.dm_build_905
}

func (dm_build_910 *Dm_build_902) Dm_build_909(dm_build_911 *Dm_build_980, dm_build_912 int) int {
	var dm_build_913 = 0
	var dm_build_914 = 0
	for dm_build_913 < dm_build_912 && dm_build_910.dm_build_904 != nil {
		dm_build_914 = dm_build_910.dm_build_904.dm_build_964(dm_build_911, dm_build_912-dm_build_913)
		if dm_build_910.dm_build_904.dm_build_959 == 0 {
			dm_build_910.dm_build_946()
		}
		dm_build_913 += dm_build_914
		dm_build_910.dm_build_905 -= dm_build_914
	}
	return dm_build_913
}

func (dm_build_916 *Dm_build_902) Dm_build_915(dm_build_917 []byte, dm_build_918 int, dm_build_919 int) int {
	var dm_build_920 = 0
	var dm_build_921 = 0
	for dm_build_920 < dm_build_919 && dm_build_916.dm_build_904 != nil {
		dm_build_921 = dm_build_916.dm_build_904.dm_build_968(dm_build_917, dm_build_918, dm_build_919-dm_build_920)
		if dm_build_916.dm_build_904.dm_build_959 == 0 {
			dm_build_916.dm_build_946()
		}
		dm_build_920 += dm_build_921
		dm_build_916.dm_build_905 -= dm_build_921
		dm_build_918 += dm_build_921
	}
	return dm_build_920
}

func (dm_build_923 *Dm_build_902) Dm_build_922(dm_build_924 io.Writer, dm_build_925 int) int {
	var dm_build_926 = 0
	var dm_build_927 = 0
	for dm_build_926 < dm_build_925 && dm_build_923.dm_build_904 != nil {
		dm_build_927 = dm_build_923.dm_build_904.dm_build_973(dm_build_924, dm_build_925-dm_build_926)
		if dm_build_923.dm_build_904.dm_build_959 == 0 {
			dm_build_923.dm_build_946()
		}
		dm_build_926 += dm_build_927
		dm_build_923.dm_build_905 -= dm_build_927
	}
	return dm_build_926
}

func (dm_build_929 *Dm_build_902) Dm_build_928(dm_build_930 []byte, dm_build_931 int, dm_build_932 int) {
	if dm_build_932 == 0 {
		return
	}
	var dm_build_933 = dm_build_960(dm_build_930, dm_build_931, dm_build_932)
	if dm_build_929.dm_build_904 == nil {
		dm_build_929.dm_build_904 = dm_build_933
	} else {
		dm_build_929.dm_build_903.PushBack(dm_build_933)
	}
	dm_build_929.dm_build_905 += dm_build_932
}

func (dm_build_935 *Dm_build_902) dm_build_934(dm_build_936 int) byte {
	var dm_build_937 = dm_build_936
	var dm_build_938 = dm_build_935.dm_build_904
	for dm_build_937 > 0 && dm_build_938 != nil {
		if dm_build_938.dm_build_959 == 0 {
			continue
		}
		if dm_build_937 > dm_build_938.dm_build_959-1 {
			dm_build_937 -= dm_build_938.dm_build_959
			dm_build_938 = dm_build_935.dm_build_903.Front().Value.(*dm_build_956)
		} else {
			break
		}
	}
	return dm_build_938.dm_build_977(dm_build_937)
}
func (dm_build_940 *Dm_build_902) Dm_build_939(dm_build_941 *Dm_build_902) {
	if dm_build_941.dm_build_905 == 0 {
		return
	}
	var dm_build_942 = dm_build_941.dm_build_904
	for dm_build_942 != nil {
		dm_build_940.dm_build_943(dm_build_942)
		dm_build_941.dm_build_946()
		dm_build_942 = dm_build_941.dm_build_904
	}
	dm_build_941.dm_build_905 = 0
}
func (dm_build_944 *Dm_build_902) dm_build_943(dm_build_945 *dm_build_956) {
	if dm_build_945.dm_build_959 == 0 {
		return
	}
	if dm_build_944.dm_build_904 == nil {
		dm_build_944.dm_build_904 = dm_build_945
	} else {
		dm_build_944.dm_build_903.PushBack(dm_build_945)
	}
	dm_build_944.dm_build_905 += dm_build_945.dm_build_959
}

func (dm_build_947 *Dm_build_902) dm_build_946() {
	var dm_build_948 = dm_build_947.dm_build_903.Front()
	if dm_build_948 == nil {
		dm_build_947.dm_build_904 = nil
	} else {
		dm_build_947.dm_build_904 = dm_build_948.Value.(*dm_build_956)
		dm_build_947.dm_build_903.Remove(dm_build_948)
	}
}

func (dm_build_950 *Dm_build_902) Dm_build_949() []byte {
	var dm_build_951 = make([]byte, dm_build_950.dm_build_905)
	var dm_build_952 = dm_build_950.dm_build_904
	var dm_build_953 = 0
	var dm_build_954 = len(dm_build_951)
	var dm_build_955 = 0
	for dm_build_952 != nil {
		if dm_build_952.dm_build_959 > 0 {
			if dm_build_954 > dm_build_952.dm_build_959 {
				dm_build_955 = dm_build_952.dm_build_959
			} else {
				dm_build_955 = dm_build_954
			}
			copy(dm_build_951[dm_build_953:dm_build_953+dm_build_955], dm_build_952.dm_build_957[dm_build_952.dm_build_958:dm_build_952.dm_build_958+dm_build_955])
			dm_build_953 += dm_build_955
			dm_build_954 -= dm_build_955
		}
		if dm_build_950.dm_build_903.Front() == nil {
			dm_build_952 = nil
		} else {
			dm_build_952 = dm_build_950.dm_build_903.Front().Value.(*dm_build_956)
		}
	}
	return dm_build_951
}

type dm_build_956 struct {
	dm_build_957 []byte
	dm_build_958 int
	dm_build_959 int
}

func dm_build_960(dm_build_961 []byte, dm_build_962 int, dm_build_963 int) *dm_build_956 {
	return &dm_build_956{
		dm_build_961,
		dm_build_962,
		dm_build_963,
	}
}

func (dm_build_965 *dm_build_956) dm_build_964(dm_build_966 *Dm_build_980, dm_build_967 int) int {
	if dm_build_965.dm_build_959 <= dm_build_967 {
		dm_build_967 = dm_build_965.dm_build_959
	}
	dm_build_966.Dm_build_1059(dm_build_965.dm_build_957[dm_build_965.dm_build_958 : dm_build_965.dm_build_958+dm_build_967])
	dm_build_965.dm_build_958 += dm_build_967
	dm_build_965.dm_build_959 -= dm_build_967
	return dm_build_967
}

func (dm_build_969 *dm_build_956) dm_build_968(dm_build_970 []byte, dm_build_971 int, dm_build_972 int) int {
	if dm_build_969.dm_build_959 <= dm_build_972 {
		dm_build_972 = dm_build_969.dm_build_959
	}
	copy(dm_build_970[dm_build_971:dm_build_971+dm_build_972], dm_build_969.dm_build_957[dm_build_969.dm_build_958:dm_build_969.dm_build_958+dm_build_972])
	dm_build_969.dm_build_958 += dm_build_972
	dm_build_969.dm_build_959 -= dm_build_972
	return dm_build_972
}

func (dm_build_974 *dm_build_956) dm_build_973(dm_build_975 io.Writer, dm_build_976 int) int {
	if dm_build_974.dm_build_959 <= dm_build_976 {
		dm_build_976 = dm_build_974.dm_build_959
	}
	dm_build_975.Write(dm_build_974.dm_build_957[dm_build_974.dm_build_958 : dm_build_974.dm_build_958+dm_build_976])
	dm_build_974.dm_build_958 += dm_build_976
	dm_build_974.dm_build_959 -= dm_build_976
	return dm_build_976
}
func (dm_build_978 *dm_build_956) dm_build_977(dm_build_979 int) byte {
	return dm_build_978.dm_build_957[dm_build_978.dm_build_958+dm_build_979]
}
