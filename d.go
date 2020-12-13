/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_875 struct {
	dm_build_876 *list.List
	dm_build_877 *dm_build_929
	dm_build_878 int
}

func Dm_build_879() *Dm_build_875 {
	return &Dm_build_875{
		dm_build_876: list.New(),
		dm_build_878: 0,
	}
}

func (dm_build_881 *Dm_build_875) Dm_build_880() int {
	return dm_build_881.dm_build_878
}

func (dm_build_883 *Dm_build_875) Dm_build_882(dm_build_884 *Dm_build_953, dm_build_885 int) int {
	var dm_build_886 = 0
	var dm_build_887 = 0
	for dm_build_886 < dm_build_885 && dm_build_883.dm_build_877 != nil {
		dm_build_887 = dm_build_883.dm_build_877.dm_build_937(dm_build_884, dm_build_885-dm_build_886)
		if dm_build_883.dm_build_877.dm_build_932 == 0 {
			dm_build_883.dm_build_919()
		}
		dm_build_886 += dm_build_887
		dm_build_883.dm_build_878 -= dm_build_887
	}
	return dm_build_886
}

func (dm_build_889 *Dm_build_875) Dm_build_888(dm_build_890 []byte, dm_build_891 int, dm_build_892 int) int {
	var dm_build_893 = 0
	var dm_build_894 = 0
	for dm_build_893 < dm_build_892 && dm_build_889.dm_build_877 != nil {
		dm_build_894 = dm_build_889.dm_build_877.dm_build_941(dm_build_890, dm_build_891, dm_build_892-dm_build_893)
		if dm_build_889.dm_build_877.dm_build_932 == 0 {
			dm_build_889.dm_build_919()
		}
		dm_build_893 += dm_build_894
		dm_build_889.dm_build_878 -= dm_build_894
		dm_build_891 += dm_build_894
	}
	return dm_build_893
}

func (dm_build_896 *Dm_build_875) Dm_build_895(dm_build_897 io.Writer, dm_build_898 int) int {
	var dm_build_899 = 0
	var dm_build_900 = 0
	for dm_build_899 < dm_build_898 && dm_build_896.dm_build_877 != nil {
		dm_build_900 = dm_build_896.dm_build_877.dm_build_946(dm_build_897, dm_build_898-dm_build_899)
		if dm_build_896.dm_build_877.dm_build_932 == 0 {
			dm_build_896.dm_build_919()
		}
		dm_build_899 += dm_build_900
		dm_build_896.dm_build_878 -= dm_build_900
	}
	return dm_build_899
}

func (dm_build_902 *Dm_build_875) Dm_build_901(dm_build_903 []byte, dm_build_904 int, dm_build_905 int) {
	if dm_build_905 == 0 {
		return
	}
	var dm_build_906 = dm_build_933(dm_build_903, dm_build_904, dm_build_905)
	if dm_build_902.dm_build_877 == nil {
		dm_build_902.dm_build_877 = dm_build_906
	} else {
		dm_build_902.dm_build_876.PushBack(dm_build_906)
	}
	dm_build_902.dm_build_878 += dm_build_905
}

func (dm_build_908 *Dm_build_875) dm_build_907(dm_build_909 int) byte {
	var dm_build_910 = dm_build_909
	var dm_build_911 = dm_build_908.dm_build_877
	for dm_build_910 > 0 && dm_build_911 != nil {
		if dm_build_911.dm_build_932 == 0 {
			continue
		}
		if dm_build_910 > dm_build_911.dm_build_932-1 {
			dm_build_910 -= dm_build_911.dm_build_932
			dm_build_911 = dm_build_908.dm_build_876.Front().Value.(*dm_build_929)
		} else {
			break
		}
	}
	return dm_build_911.dm_build_950(dm_build_910)
}
func (dm_build_913 *Dm_build_875) Dm_build_912(dm_build_914 *Dm_build_875) {
	if dm_build_914.dm_build_878 == 0 {
		return
	}
	var dm_build_915 = dm_build_914.dm_build_877
	for dm_build_915 != nil {
		dm_build_913.dm_build_916(dm_build_915)
		dm_build_914.dm_build_919()
		dm_build_915 = dm_build_914.dm_build_877
	}
	dm_build_914.dm_build_878 = 0
}
func (dm_build_917 *Dm_build_875) dm_build_916(dm_build_918 *dm_build_929) {
	if dm_build_918.dm_build_932 == 0 {
		return
	}
	if dm_build_917.dm_build_877 == nil {
		dm_build_917.dm_build_877 = dm_build_918
	} else {
		dm_build_917.dm_build_876.PushBack(dm_build_918)
	}
	dm_build_917.dm_build_878 += dm_build_918.dm_build_932
}

func (dm_build_920 *Dm_build_875) dm_build_919() {
	var dm_build_921 = dm_build_920.dm_build_876.Front()
	if dm_build_921 == nil {
		dm_build_920.dm_build_877 = nil
	} else {
		dm_build_920.dm_build_877 = dm_build_921.Value.(*dm_build_929)
		dm_build_920.dm_build_876.Remove(dm_build_921)
	}
}

func (dm_build_923 *Dm_build_875) Dm_build_922() []byte {
	var dm_build_924 = make([]byte, dm_build_923.dm_build_878)
	var dm_build_925 = dm_build_923.dm_build_877
	var dm_build_926 = 0
	var dm_build_927 = len(dm_build_924)
	var dm_build_928 = 0
	for dm_build_925 != nil {
		if dm_build_925.dm_build_932 > 0 {
			if dm_build_927 > dm_build_925.dm_build_932 {
				dm_build_928 = dm_build_925.dm_build_932
			} else {
				dm_build_928 = dm_build_927
			}
			copy(dm_build_924[dm_build_926:dm_build_926+dm_build_928], dm_build_925.dm_build_930[dm_build_925.dm_build_931:dm_build_925.dm_build_931+dm_build_928])
			dm_build_926 += dm_build_928
			dm_build_927 -= dm_build_928
		}
		if dm_build_923.dm_build_876.Front() == nil {
			dm_build_925 = nil
		} else {
			dm_build_925 = dm_build_923.dm_build_876.Front().Value.(*dm_build_929)
		}
	}
	return dm_build_924
}

type dm_build_929 struct {
	dm_build_930 []byte
	dm_build_931 int
	dm_build_932 int
}

func dm_build_933(dm_build_934 []byte, dm_build_935 int, dm_build_936 int) *dm_build_929 {
	return &dm_build_929{
		dm_build_934,
		dm_build_935,
		dm_build_936,
	}
}

func (dm_build_938 *dm_build_929) dm_build_937(dm_build_939 *Dm_build_953, dm_build_940 int) int {
	if dm_build_938.dm_build_932 <= dm_build_940 {
		dm_build_940 = dm_build_938.dm_build_932
	}
	dm_build_939.Dm_build_1032(dm_build_938.dm_build_930[dm_build_938.dm_build_931 : dm_build_938.dm_build_931+dm_build_940])
	dm_build_938.dm_build_931 += dm_build_940
	dm_build_938.dm_build_932 -= dm_build_940
	return dm_build_940
}

func (dm_build_942 *dm_build_929) dm_build_941(dm_build_943 []byte, dm_build_944 int, dm_build_945 int) int {
	if dm_build_942.dm_build_932 <= dm_build_945 {
		dm_build_945 = dm_build_942.dm_build_932
	}
	copy(dm_build_943[dm_build_944:dm_build_944+dm_build_945], dm_build_942.dm_build_930[dm_build_942.dm_build_931:dm_build_942.dm_build_931+dm_build_945])
	dm_build_942.dm_build_931 += dm_build_945
	dm_build_942.dm_build_932 -= dm_build_945
	return dm_build_945
}

func (dm_build_947 *dm_build_929) dm_build_946(dm_build_948 io.Writer, dm_build_949 int) int {
	if dm_build_947.dm_build_932 <= dm_build_949 {
		dm_build_949 = dm_build_947.dm_build_932
	}
	dm_build_948.Write(dm_build_947.dm_build_930[dm_build_947.dm_build_931 : dm_build_947.dm_build_931+dm_build_949])
	dm_build_947.dm_build_931 += dm_build_949
	dm_build_947.dm_build_932 -= dm_build_949
	return dm_build_949
}
func (dm_build_951 *dm_build_929) dm_build_950(dm_build_952 int) byte {
	return dm_build_951.dm_build_930[dm_build_951.dm_build_931+dm_build_952]
}
