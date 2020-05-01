package presuffix

import (
	"fmt"
	"strconv"
	"time"
)

// TimeWork will invoke all the test i'm running
// func TimeWork() (t time.Time) {
// 	return
// }
func TimeWork() (t time.Time) {
	t = time.Now()
	shard, _ := t.MarshalJSON()
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC)) // 27 Mar 2020 16:16
	fmt.Println(string(shard))
	return
}
func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}
