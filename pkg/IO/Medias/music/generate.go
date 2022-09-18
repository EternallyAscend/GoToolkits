package music

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	fifteen = 15 * 1000
	ten     = 10 * 1000
)

type Time struct {
	Hour        int64
	Minute      int64
	Second      int64
	MicroSecond int64
}

func GenerateTime(hour int64, minute int64, second int64, micro int64) *Time {
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}
}

func ReadTimeFromString(time string) *Time {
	result := strings.Split(time, ",")
	if 4 != len(result) {
		fmt.Println(errors.New("Wrong time: " + time))
		return nil
	}
	hour, err := strconv.ParseInt(result[0], 10, 64)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	minute, err := strconv.ParseInt(result[1], 10, 64)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	second, err := strconv.ParseInt(result[2], 10, 64)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	micro, err := strconv.ParseInt(result[3], 10, 64)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}
}

func (that *Time) ExportString() string {
	return fmt.Sprintf("%s,%s,%s,%s", that.HourString(), that.MinuteString(), that.SecondString(), strconv.FormatInt(that.MicroSecond, 10))
}

func (that *Time) TransferTimeToTimestamp() int64 {
	return ((that.Hour*60+that.Minute)*60+that.Second)*1000 + that.MicroSecond
}

func (that *Time) HourString() string {
	base := ""
	if that.Hour < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Hour, 10)
	return base
}

func (that *Time) MinuteWithHourString() string {
	base := ""
	if that.Hour*60+that.Minute < 100 {
		base = "0"
	}
	base += strconv.FormatInt(that.Hour*60+that.Minute, 10)
	return base
}

func (that *Time) MinuteString() string {
	base := ""
	if that.Minute < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Minute, 10)
	return base
}

func (that *Time) SecondString() string {
	base := ""
	if that.Second < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Second, 10)
	return base
}

func (that *Time) MicroHundredString() string {
	base := ""
	if that.MicroSecond < 100 {
		base = "0"
		if that.MicroSecond < 10 {
			base = "00"
		}
	}
	base += strconv.FormatInt(that.MicroSecond, 10)
	return base
}

func (that *Time) MicroTwoDecimalString() string {
	base := ""
	if that.MicroSecond < 100 {
		base = "0"
	}
	base += strconv.FormatInt(that.MicroSecond/10, 10)
	return base
}

func (that *Time) ExportSRT() string {
	return that.HourString() + ":" + that.MinuteString() + ":" + that.SecondString() + "," + that.MicroHundredString()
}

func (that *Time) ExportLRC() string {
	return "[" + that.MinuteWithHourString() + ":" + that.SecondString() + "." + that.MicroTwoDecimalString() + "]"
}

func TransferTimestampToTime(timestamp int64) *Time {
	hour := timestamp / 1000 / 60 / 60
	minute := timestamp/1000/60 - hour*60
	second := (timestamp/1000 - minute*60) - hour*60*60
	micro := timestamp % 1000
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}
}

type Line struct {
	Start   *Time
	End     *Time
	Content string
}

func GenerateLine(start *Time, end *Time, content string) *Line {
	return &Line{
		Start:   start,
		End:     end,
		Content: content,
	}
}

func ReadLineFromString(line string) *Line {
	result := strings.Split(line, ":")
	if 3 != len(result) {
		fmt.Println(errors.New("Wrong line: " + line))
		return nil
	}
	return &Line{
		Start:   ReadTimeFromString(result[0]),
		End:     ReadTimeFromString(result[1]),
		Content: result[2],
	}
}

type Sentence struct {
	Offset   int64
	Title    string
	Album    string
	Artist   string
	Author   string
	Lyrics   string
	Composer string
	Lines    *[]*Line
}

func GenerateEmptySentence() *Sentence {
	return &Sentence{
		Offset:   0,
		Title:    "",
		Album:    "",
		Artist:   "",
		Author:   "",
		Lyrics:   "",
		Composer: "",
		Lines:    &[]*Line{},
	}
}

func (that *Sentence) DealTimeOffset(time *Time) *Time {
	return TransferTimestampToTime(time.TransferTimeToTimestamp() + that.Offset)
}

func (that *Sentence) DealInfoChinese() {
	info := &[]*Line{}
	if (*that.Lines)[0].Start.TransferTimeToTimestamp() > fifteen {
		info = &[]*Line{
			GenerateLine(
				GenerateTime(0, 0, 0, 0),
				GenerateTime(0, 0, 5, 0),
				that.Title),
			GenerateLine(
				GenerateTime(0, 0, 5, 0),
				GenerateTime(0, 0, 10, 0),
				"词："+that.Lyrics),
			GenerateLine(
				GenerateTime(0, 0, 10, 0),
				GenerateTime(0, 0, 15, 0),
				"曲："+that.Composer),
		}
	} else if (*that.Lines)[0].Start.TransferTimeToTimestamp() > ten {
		info = &[]*Line{
			GenerateLine(
				GenerateTime(0, 0, 0, 0),
				GenerateTime(0, 0, 5, 0),
				that.Title),
			GenerateLine(
				GenerateTime(0, 0, 5, 0),
				GenerateTime(0, 0, 10, 0),
				"词："+that.Lyrics+" 曲："+that.Composer),
		}
	} else {
		info = &[]*Line{
			GenerateLine(
				GenerateTime(0, 0, 0, 0),
				(*that.Lines)[0].Start,
				that.Title+"词："+that.Lyrics+" 曲："+that.Composer),
		}
	}
	*that.Lines = append(*info, *that.Lines...)
}

func (that *Sentence) exportDealing() {
	// TODO 检查输出时Offset和Line第一行Start的关系
	that.DealInfoChinese()
}

func (that *Sentence) exportSaving(path string, append string, data string) {
	file, err := os.OpenFile(path+that.Title+append, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			fmt.Println(errIn.Error())
		}
	}(file)
	_, err = file.WriteString(data)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func (that *Sentence) ExportFile(path string) {
	data := ""
	var line *Line
	for i := range *that.Lines {
		line = (*that.Lines)[i]
		data += fmt.Sprintf("%s:%s:%s\n", line.Start.ExportString(), line.End.ExportString(), line.Content)
	}
	that.exportSaving(path, ".txt", data)
}

func (that *Sentence) ExportLRC(path string) {
	that.exportDealing()
	result := "[ti:" + that.Title + "]\n"
	result += "[ar:" + that.Artist + "]\n"
	result += "[al:" + that.Album + "]\n"
	for _, element := range *that.Lines {
		start := that.DealTimeOffset(element.Start)
		result += start.ExportLRC() + element.Content + "\n"
	}
	that.exportSaving(path, ".lrc", result)
}

func (that *Sentence) ExportSRT(path string) {
	that.exportDealing()
	result := ""
	for index, element := range *that.Lines {
		result += strconv.Itoa(index + 1)
		start := that.DealTimeOffset(element.Start)
		end := that.DealTimeOffset(element.End)
		result += "\n" + start.ExportSRT() + " --> " + end.ExportSRT()
		result += "\n" + element.Content + "\n\n"

		// fmt.Println(*element.Start)
		// fmt.Println(*TransferTimestampToTime(element.Start.TransferTimeToTimestamp()))
		// fmt.Println(*element.End)
		// fmt.Println(*TransferTimestampToTime(element.End.TransferTimeToTimestamp()))
	}
	that.exportSaving(path, ".srt", result)
}

func (that *Sentence) SetOffset(offset int64) {
	that.Offset = offset
}

func (that *Sentence) SetTitle(title string) {
	that.Title = title
}

func (that *Sentence) SetAlbum(album string) {
	that.Album = album
}

func (that *Sentence) SetArtist(artist string) {
	that.Artist = artist
}

func (that *Sentence) SetAuthor(author string) {
	that.Author = author
}

func (that *Sentence) SetLyrics(lyrics string) {
	that.Lyrics = lyrics
}

func (that *Sentence) SetComposer(composer string) {
	that.Composer = composer
}

func (that *Sentence) AddLine(line *Line) {
	*that.Lines = append(*that.Lines, line)
}

func (that *Sentence) LoadLines(filePath string) {
	file, err := os.Open(filePath)
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			fmt.Println(errIn.Error())
			return
		}
	}(file)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		that.AddLine(ReadLineFromString(scanner.Text()))
	}
}

func GenerateOldTime() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("年老")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.LoadLines("./src/" + sentence.Title + ".txt")
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func GenerateStillSing() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("一直这样唱")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.LoadLines("./src/" + sentence.Title + ".txt")
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func GenerateAfterKnown() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("了解之后")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.LoadLines("./src/" + sentence.Title + ".txt")
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func GenerateOurSong() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("我们的歌")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.LoadLines("./src/" + sentence.Title + ".txt")
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func GenerateEnterYourCar() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("坐上你的车")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.LoadLines("./src/" + sentence.Title + ".txt")
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func GenerateEgg() {
	filePath := "./src/"
	sentence := GenerateEmptySentence()
	sentence.SetOffset(0)
	sentence.SetAuthor("未知")
	sentence.SetTitle("给我一个痛快")
	sentence.SetAlbum("2022 乐界 Demo")
	sentence.SetArtist("张雨生")
	sentence.SetLyrics("张雨生")
	sentence.SetComposer("张雨生")
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 0, 49, 0),
		GenerateTime(0, 0, 53, 500),
		"烟雾中浮现一幅画"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 0, 54, 0),
		GenerateTime(0, 0, 58, 500),
		"有好多瑰丽的花"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 0, 59, 0),
		GenerateTime(0, 1, 3, 0),
		"我喜欢这样静静地"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 3, 500),
		GenerateTime(0, 1, 8, 0),
		"让时间从容过去"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 8, 0),
		GenerateTime(0, 1, 12, 500),
		"伤口中酥麻着快感"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 13, 0),
		GenerateTime(0, 1, 16, 500),
		"为原始的兽性震撼"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 17, 0),
		GenerateTime(0, 1, 21, 500),
		"我期待这样疯狂地"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 22, 0),
		GenerateTime(0, 1, 27, 500),
		"把空虚胡乱填平"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 29, 0),
		GenerateTime(0, 1, 33, 0),
		"总想为每件事找理由"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 34, 0),
		GenerateTime(0, 1, 37, 0),
		"每个理由都是空洞"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 39, 0),
		GenerateTime(0, 1, 41, 0),
		"总想形容我每个心痛"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 43, 0),
		GenerateTime(0, 1, 50, 0),
		"当我被诅咒的时候"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 52, 0),
		GenerateTime(0, 1, 56, 0),
		"给我一个痛快"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 1, 58, 0),
		GenerateTime(0, 2, 1, 100),
		"给我即兴摇摆"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 2, 0),
		GenerateTime(0, 2, 10, 0),
		"哪怕是烟火就只要能填充无奈"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 11, 0),
		GenerateTime(0, 2, 15, 500),
		"永无意义的光环中"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 16, 0),
		GenerateTime(0, 2, 20, 0),
		"因为我深怕那寂寞"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 20, 500),
		GenerateTime(0, 2, 25, 0),
		"回忆是治不愈的创痛"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 26, 0),
		GenerateTime(0, 2, 30, 0),
		"常令我午夜惊梦"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 30, 0),
		GenerateTime(0, 2, 34, 500),
		"仍然有纯洁的思欲"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 35, 0),
		GenerateTime(0, 2, 39, 0),
		"我问你相不相信"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 39, 500),
		GenerateTime(0, 2, 44, 0),
		"谁敢说他不戴面具"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 44, 0),
		GenerateTime(0, 2, 50, 500),
		"应酬中逃离人群"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 50, 500),
		GenerateTime(0, 2, 55, 500),
		"不是想证明我已成熟"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 2, 56, 0),
		GenerateTime(0, 3, 0, 0),
		"只是想表现的与众不同"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 1, 0),
		GenerateTime(0, 3, 5, 500),
		"当你们指责我的时候"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 5, 500),
		GenerateTime(0, 3, 13, 500),
		"希望你们能听我说"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 14, 0),
		GenerateTime(0, 3, 18, 500),
		"给我一个痛快"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 19, 0),
		GenerateTime(0, 3, 23, 500),
		"给我即兴摇摆"))
	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 24, 0),
		GenerateTime(0, 3, 33, 330),
		"哪怕是烟火就只要能填充无奈"))

	sentence.AddLine(GenerateLine(
		GenerateTime(0, 3, 33, 333),
		GenerateTime(0, 4, 4, 0),
		"... 尾奏 Solo ..."))
	sentence.ExportFile(filePath)
	sentence.ExportSRT(filePath)
	sentence.ExportLRC(filePath)
}

func Generate() {
	// GenerateEgg()
	// GenerateOldTime()
	// GenerateStillSing()
	// GenerateAfterKnown()
	// GenerateOurSong()
	// GenerateEnterYourCar()
}
