package time_test

import (
	"testing"
	"time"
)

// time.After 会在指定的时间段 d 后，向返回值发送当时的时间。
func TestTimeAfter(t *testing.T) {
	var c chan int
	f := func(int) {}
	select {
	case m := <-c:
		f(m)
	case <-time.After(10 * time.Second):
		t.Log("timed out")
	}
}

func TestTimeAfter2(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		time.Sleep(3 * time.Second)
		ch <- 1
	}(ch)

	select {
	case <-ch:
		t.Log(ch)
	case <-time.After(2 * time.Second):
		t.Log("timeout")
	}
}

func TestTimeAfter3(t *testing.T) {
	timeCh := time.After(3 * time.Second)
	t.Logf("Start time: %v.\n", time.Now())
	t.Logf("Receive from timeCh: %v.\n", <-timeCh)
	t.Logf("End time: %v.\n", time.Now())
}

// time.Tick 是对NewTicker更易用的包装，提供了对定时器channel的访问。
// 通常用于没必要停止定时器的客户端。
func TestTimeTick(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	for now := range tick {
		t.Logf("now = %v\n", now)
	}
}

// time.Duration 表示两个时间点之间经过的时间，以纳秒为单位。
// type Duration int64
func TestTimeDuration(t *testing.T) {
	f := func() {}
	t0 := time.Now()
	f()
	t1 := time.Now()
	t.Logf("The call took %v to run.\n", t1.Sub(t0))
}

// time.ParseDuration 解析一个时间段字符串。
func TestTimeParseDuration(t *testing.T) {
	hours, _ := time.ParseDuration("10h")
	hourMinutes, _ := time.ParseDuration("1h30m")
	micro, _ := time.ParseDuration("1µs")
	micro2, _ := time.ParseDuration("1us")

	t.Logf("hours = %v, %T\n", hours, hours)
	t.Logf("There are %.0f seconds in %v.\n", hours.Seconds(), hours)
	t.Logf("hourMinutes = %v\n", hourMinutes)
	t.Logf("micro = %v\n", micro)
	t.Logf("micro2 = %v\n", micro2)
}

// time.Since 返回从 t 到现在经过的时间。
func TestTimeSince(t *testing.T) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	t.Logf("The call took %v to run.\n", time.Since(start))
}

// time.Until 返回从现在到 t 经过的时间。
func TestTimeUntil(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(1 * time.Second)
	t.Logf("t1 = %v, t2 = %v\n", t1, t2)
	//t.Logf("t1 until t2 = %v\n", t1.Until(t2))
}

// time.Abs 返回 d 的绝对时间。
func TestTimeAbs(t *testing.T) {
	d, _ := time.ParseDuration("1h30m")
	t.Logf("d.Abs() = %v\n", d.Abs())
}

// time.Hours 返回时间段内的整数小时部分。
func TestTimeHours(t *testing.T) {
	d, _ := time.ParseDuration("1h30m50s20ms200us70ns")
	t.Logf("d.Hours() = %v\n", d.Hours())
	t.Logf("d.Minutes() = %v\n", d.Minutes())
	t.Logf("d.Seconds() = %v\n", d.Seconds())
	t.Logf("d.Milliseconds() = %v\n", d.Milliseconds())
	t.Logf("d.Microseconds() = %v\n", d.Microseconds())
	t.Logf("d.Nanoseconds() = %v\n", d.Nanoseconds())
}

// d.Round 返回舍入给定时间 d 到最接近给定时间段 m 的倍数的输出
func TestTimeRound(t *testing.T) {
	d, _ := time.ParseDuration("1h15m30.918273645s")

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		t.Logf("d.Round(%v) = %v\n", r, d.Round(r).String())
	}
}

// d.Truncate 返回时间戳 d 与时间粒度 m 的最小整数倍对齐，然后返回对齐之后的时间戳
// 与 d.Round 不同的是，d.Truncate 只会向下舍入。
func TestTimeTruncate(t *testing.T) {
	d, _ := time.ParseDuration("1h15m30.918273645s")

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range trunc {
		t.Logf("d.Truncate(%6s) = %s\n", r, d.Truncate(r).String())
	}
}

// time.FixedZone 返回一个 Location，它的名字和偏移量都是固定的。
func TestTimeFixedZone(t *testing.T) {
	tz := time.FixedZone("UTC-8", -8*60*60)
	t.Logf("tz = %v\n", tz)
	t.Logf("tz.String() = %v\n", tz.String())
}

// time.LoadLocation 从时区数据库中加载名为 name 的 Location。
func TestTimeLoadLocation(t *testing.T) {
	tz, _ := time.LoadLocation("Asia/Shanghai")
	t.Logf("tz = %v\n", tz)
	t.Logf("tz.String() = %v\n", tz.String())

	timeIn := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	t.Logf("time = %v\n", timeIn)
	t.Logf("time = %v\n", timeIn.In(tz))
}

// time.LoadLocationFromTZData 从时区数据库中加载时区数据。
func TestTimeLoadLocationFromTZData(t *testing.T) {
	tz, _ := time.LoadLocationFromTZData("Asia/Shanghai", []byte("xxxx"))
	t.Logf("tz = %v\n", tz)
}

// time.Month 表示一年的某个月份。
// type Month int
func TestTimeMonth(t *testing.T) {
	now := time.Now()
	t.Logf("now = %v\n", now)
	t.Logf("now.Month() = %v\n", now.Month())
	t.Logf("now.Month().String() = %v\n", now.Month().String())
}

// time.NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，
// 并会每隔时间段 d 就向该通道发送当时的时间。
func TestTimeNewTicker(t *testing.T) {
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			t.Logf("Done!!!")
			return
		case curTime := <-ticker.C:
			t.Logf("Current time: %v\n", curTime)
		}
	}
}

// time.Reset 停止一个 Ticker，并重置这个 Ticker 的 duration 为新的 duration
func TestTimeTickerReset(t *testing.T) {
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	ticker.Reset(3 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			t.Logf("Done!!!")
			return
		case curTime := <-ticker.C:
			t.Logf("Current time: %v\n", curTime)
		}
	}
}

// time.Date 返回一个根据参数创建的时间点。
func TestTimeDate(t *testing.T) {
	date := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t.Logf("Go launched at %s\n", date.Local())
}

// time.Now 返回当前本地时间。
func TestTimeNow(t *testing.T) {
	t.Logf("time.Now() = %v\n", time.Now())
}

// time.Parse 解析一个格式化的时间字符串并返回它代表的时间。
func TestTimeParseLayout(t *testing.T) {
	t1, _ := time.Parse(time.DateTime, "2020-01-01 00:00:00")
	t.Logf("t1 = %v\n", t1)
	t.Logf("t1.Format(time.ANSIC) = %v\n", t1.Format(time.ANSIC))
}

// time.ParseInLocation 类似于 Parse，但有两个重要的不同之处。
// 1.在没有时区信息的情况下，Parse 将时间解释为 UTC； ParseInLocation 将时间解释为给定位置的时间。
// 2.当给定区域偏移量或缩写时，Parse 会尝试将其与本地位置进行匹配； ParseInLocation 使用给定的位置。
func TestTimeParseInLocation(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	timeLoc1, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CST)", location)
	t.Logf("timeLoc1 = %v\n", timeLoc1)

	const shortForm = "2006-Jan-02"
	timeLoc2, _ := time.ParseInLocation(shortForm, "2012-Jul-09", location)
	t.Logf("timeLoc2 = %v\n", timeLoc2)
}

// time.Unix 使用秒和纳秒来返回相对January 1, 1970 UTC的时间间隔。
func TestTimeUnix(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.Unix() = %v\n", timeLoc.Unix())

	t1 := time.Unix(timeLoc.Unix(), 0).Unix()
	t.Logf("t1 = %v\n", t1)
}

// time.UnixMicro 返回相对January 1, 1970 UTC的时间间隔，以微秒为单位。
func TestTimeUnixMicro(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.UnixMicro() = %v\n", timeLoc.UnixMicro())

	t1 := time.UnixMicro(timeLoc.UnixMicro()).UnixMicro()
	t.Logf("t1 = %v\n", t1)
}

// time.UnixMilli 返回相对January 1, 1970 UTC的时间间隔，以毫秒为单位。
func TestTimeUnixMilli(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.UnixMilli() = %v\n", timeLoc.UnixMilli())

	t1 := time.UnixMilli(timeLoc.UnixMilli()).UnixMilli()
	t.Logf("t1 = %v\n", t1)
}

// time.AddDate 返回增加了给定年数、月数和天数的时间点。
func TestTimeAddDate(t *testing.T) {
	t1 := time.Now()
	t.Logf("t1 = %v\n", t1)
	t.Logf("t1.AddDate(1, 1, 1) = %v\n", t1.AddDate(1, 1, 1))
}

// time.After 返回 t 是否在 u 之后。
func TestTimeAfter4(t *testing.T) {
	timeLoc := time.Now()
	timeBefore := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Logf("timeLoc is after timeBefore: %v\n", timeLoc.After(timeBefore))
	t.Logf("timeBefore is before timeLoc: %v\n", timeBefore.Before(timeLoc))
}

// time.AppendFormat 将 t 格式化为格式字符串并将结果追加到 b 中。
func TestTimeAppendFormat(t *testing.T) {
	timeLoc := time.Now()
	text := []byte("Time: ")
	text = timeLoc.AppendFormat(text, time.Kitchen)
	t.Logf("timeLoc.AppendFormat() = %v\n", string(text))
}

// time.Clock 返回一个 t 指定的日期内的小时、分钟和秒。
func TestTimeClock(t *testing.T) {
	timeLoc := time.Now()
	hour, min, sec := timeLoc.Clock()
	t.Log(hour, min, sec)
}

// time.Compare 返回 t 与 u 的比较结果。
func TestTimeCompare(t *testing.T) {
	timeLoc := time.Now()
	timeBefore := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Logf("timeLoc.Compare(timeBefore) = %v\n", timeLoc.Compare(timeBefore))
}

// time.Date 返回一个年月日
func TestTimeDate2(t *testing.T) {
	timeLoc := time.Now()
	year, month, day := timeLoc.Date()
	t.Log(year, int(month), day)
}

// time.Day 返回指定月份的第几天
func TestTimeDay(t *testing.T) {
	timeLoc := time.Now()
	day := timeLoc.Day()
	t.Logf("day = %v\n", day)
}

// time.Equal 报告 t 和 u 是否代表同一时刻，即使两个时间在不同的时区或位置。
func TestTimeEqual(t *testing.T) {
	// 构造两个不同的时区
	loc1 := time.FixedZone("GMT -1:00", -3600)
	loc2 := time.FixedZone("GMT +8:00", 28800)

	str := "2020-01-01 00:00:00"

	t1, _ := time.Parse("2006-01-02 15:04:05", str)
	t1 = t1.In(loc1)

	t2, _ := time.Parse("2006-01-02 15:04:05", str)
	t2 = t2.In(loc2)

	t.Logf("t1 = %v\n", t1)
	t.Logf("t2 = %v\n", t2)
	t.Logf("t1.Equal(t2) = %v\n", t1.Equal(t2))
	t.Logf("t1 = t2, %v\n", t1 == t2)
}

// time.Format 返回使用给定格式进行格式化的时间字符串。
func TestTimeFormat(t *testing.T) {
	timeLoc := time.Now()
	datetime := timeLoc.Format(time.DateTime)
	t.Logf("timeLoc.Format(time.Kitchen) = %q\n", datetime)
}

// time.GoString 返回一个表示该时间的Go语法格式的字符串。
func TestTimeGoString(t *testing.T) {
	t1, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	t.Logf("t1 = %v\n", t1.GoString())
}

// time.GobEncode 实现了 gob.GobEncoder 接口。
func TestTimeGobEncode(t *testing.T) {
	t1, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	b, _ := t1.GobEncode()
	t.Logf("t1.GobEncode() = %v\n", b)
}

// TIME.GobDecode 实现了 gob.GobDecoder 接口。
// 不推荐使用，将在 Go 2 中删除。
func TestTimeGobDecode(t *testing.T) {
	t1, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	b, _ := t1.GobEncode()
	t.Logf("t1.GobEncode() = %v\n", b)

	t2 := time.Time{}
	_ = t2.GobDecode(b)
	t.Logf("t2 = %v\n", t2.String())
}

// time.ISOWeek 返回 ISO 8601 格式的年份和周编号。
func TestTimeISOWeek(t *testing.T) {
	timeLoc := time.Now()
	year, week := timeLoc.ISOWeek()
	t.Logf("year = %v, week = %v\n", year, week)
}

// time.In 返回一个拥有相同年月日，但设置了新的时区的时间。
func TestTimeIn(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.In(time.UTC) = %v\n", timeLoc.In(time.UTC))
}

// time.IsDST 报告 t 是否是夏令时
func TestTimeIsDST(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.IsDST() = %v\n", timeLoc.IsDST())
}

// time.IsZero 报告 t 是否代表零时间点，
// 时间的零值是time.Time{}，也就是一个没有被赋值的time.Time类型变量
func TestTimeIsZero(t *testing.T) {
	timeNil := time.Time{}
	t.Logf("timeNil.IsZero() = %v\n", timeNil.IsZero())

	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.IsZero() = %v\n", timeLoc.IsZero())
}

// time.Local 返回一个使用本地时区的时间。
func TestTimeLocal(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.Local() = %v\n", timeLoc.Local())
}

// time.Location 返回 t 的时区。
func TestTimeLocation(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.Location() = %v\n", timeLoc.Location())
}

// time.MarshalBinary 实现了 encoding.BinaryMarshaler 接口。
// 不推荐使用，将在 Go 2 中删除。
func TestTimeMarshalBinary(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2023-09-17T15:04:05Z")

	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("t1 = %v\n", t1)
	b, _ := t1.MarshalBinary()
	t.Logf("t1.MarshalBinary() = %v\n", b)

	var t2 time.Time
	_ = t2.UnmarshalBinary(b)
	t.Logf("t2 = %v\n", t2)
}

// time.Minute 返回时间段内的整数分钟部分。
func TestTimeMinute(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.Minute() = %v\n", timeLoc.Minute())
	t.Logf("timeloc.Second() = %v\n", timeLoc.Second())
}

// time.Nanosecond 返回时间段内的整数纳秒部分。
func TestTimeNanosecond(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc.Nanosecond() = %v\n", timeLoc.Nanosecond())
}

// time.Round 返回舍入给定时间 d 到最接近给定时间段 m 的倍数的输出
func TestTimeRound2(t *testing.T) {
	t1 := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		t.Logf("t.Round(%6s) = %s\n", d, t1.Round(d).Format("15:04:05.999999999"))
	}
}

// time.String 返回一个表示该时间的字符串。
// 用于调试，稳定格式可使用 time.Format。
func TestTimeString(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc.String() = %v\n", timeLoc.String())
}

// time.Sub 返回 t 与 u 之间的时间间隔。
func TestTimeSub(t *testing.T) {
	timeLoc := time.Now()
	timeBefore := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Logf("timeLoc.Sub(timeBefore) = %v\n", timeLoc.Sub(timeBefore))
}

// time.Truncate 返回时间戳 d 与时间粒度 m 的最小整数倍对齐，然后返回对齐之后的时间戳
// 与 time.Round 不同的是，time.Truncate 只会向下舍入。
func TestTimeTruncate2(t *testing.T) {
	t1, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range trunc {
		t.Logf("d.Truncate(%6s) = %s\n", r, t1.Truncate(r).String())
	}
}

// time.UTC 返回一个使用UTC时区的时间。
func TestTimeUTC(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc = %v\n", timeLoc)
	t.Logf("timeLoc.UTC() = %v\n", timeLoc.UTC())
}

// time.Unix 返回使用秒和纳秒来表示时间的 Unix 时间。
func TestTimeUnix2(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc.Unix() = %v\n", timeLoc.Unix())
	t.Logf("timeLoc.UnixNano() = %v\n", timeLoc.UnixNano())
	t.Logf("time.UnixMilli() = %v\n", timeLoc.UnixMilli())
	t.Logf("time.UnixMicro() = %v\n", timeLoc.UnixMicro())
}

// time.Weekday 表示一周的某一天。
// type Weekday int
func TestTimeYearWeek(t *testing.T) {
	timeLoc := time.Now()
	t.Logf("timeLoc.Year() = %v\n", timeLoc.Year())
	t.Logf("timeLoc.YearDay() = %v\n", timeLoc.YearDay())
	t.Logf("timeLoc.Weekday() = %v\n", int(timeLoc.Weekday()))
	t.Logf("timeLoc.Weekday().String() = %v\n", timeLoc.Weekday().String())
}

// time.Zone 返回 t 的时区和该时区相对于UTC的时间偏移量。
func TestTimeZone(t *testing.T) {
	timeLoc := time.Now()
	name, offset := timeLoc.Zone()
	t.Logf("name = %v, offset = %v\n", name, offset)
}

// time.ZoneBounds 返回在时间 t 时有效的时区范围
func TestTimeZoneBounds(t *testing.T) {
	timeLoc := time.Now()
	start, end := timeLoc.ZoneBounds()
	t.Logf("start = %v, end = %v\n", start, end)
}

// time.AfterFunc 会在时间段 d 过后，调用 f。
func TestTimeAfterFunc(t *testing.T) {
	c := make(chan int)
	time.AfterFunc(3*time.Second, func() {
		t.Log("after func")
		c <- 1
	})
	<-c
}

// time.NewTimer 创建一个 Timer 计时器，
// 它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间。
func TestTimeNewTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	t.Logf("Start time: %v.\n", <-timer.C)
}

// time.Reset 停止计时器 timer 并将其重置为 d 时长。
// 如果计时器之前已经到期或被停止，会返回 false；如果计时器之前还未到期，会返回 true。
func TestTimeTimerReset(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	timer.Reset(5 * time.Second)
	t.Logf("Start time: %v.\n", <-timer.C)
}
