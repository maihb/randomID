package rid

import (
	"time"
)

func gen(base []byte) string {
	//变化小的放低位，变化大的放高位
	var ms_2 uint64 = 100       //毫秒的十位数和个位数
	var year uint64 = 50 * ms_2 //年
	var mon uint64 = 12 * year  //年
	var day uint64 = 31 * mon   //月
	var hour uint64 = 24 * day  //小时
	var mn uint64 = 60 * hour   //分钟
	var sce uint64 = 60 * mn    //秒
	//var ms_1 uint64 = 1000 * sce //毫秒 的百位数
	t := time.Now()
	millisecond := uint64(t.Nanosecond() / 1e6)
	m2 := millisecond % 100
	y := ms_2 * uint64(t.Year()%50)
	m := year * uint64(t.Month())
	d := mon * uint64(t.Day())
	h := day * uint64(t.Hour())
	min := hour * uint64(t.Minute())
	s := mn * uint64(t.Second())
	ss := sce * (millisecond / 100)

	total := m2 + y + m + d + h + min + s + ss
	var rid []byte
	for {
		if total <= 0 {
			break
		}
		i := total % 62
		rid = append(rid, base[i])
		total = (total - i) / 62
	}
	return string(rid)
}
