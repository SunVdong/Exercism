package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team = string

type TeamRecord struct {
	Name string
	MP   int // 比赛场次
	W    int // 赢得  +3 分
	D    int // 平局  +1 分
	L    int // 失败  +0 分
	P    int // 得分
}

func (t *TeamRecord) win() {
	t.MP += 1
	t.W += 1
	t.P += 3
}

func (t *TeamRecord) lose() {
	t.MP += 1
	t.L += 1
}

func (t *TeamRecord) draw() {
	t.MP += 1
	t.D += 1
	t.P += 1
}


type AllRecords map[Team]*TeamRecord

func (t AllRecords) getOrCreate(team Team) *TeamRecord {
	if t[team] == nil {
		t[team] = &TeamRecord{Name: team}
	}
	return t[team]
}


type Results []TeamRecord

func (s Results) Len() int {
	return len(s)
}
func (s Results) Less(i, j int) bool {
	if s[i].P == s[j].P {
		return s[i].Name < s[j].Name
	}

	return s[i].P > s[j].P
}
func (s Results) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	team_map := make(AllRecords)

	for scanner.Scan() {
		l := scanner.Text()
		l = strings.Trim(l, " ")

		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		arr := strings.Split(l, ";")
		if len(arr) != 3 {
			return errors.New("Invalid info")
		}

		team_1, team_2, result := arr[0], arr[1], arr[2]
		switch result {
		case "win":
			team_map.getOrCreate(team_1).win()
			team_map.getOrCreate(team_2).lose()
		case "loss":
			team_map.getOrCreate(team_1).lose()
			team_map.getOrCreate(team_2).win()
		case "draw":
			team_map.getOrCreate(team_1).draw()
			team_map.getOrCreate(team_2).draw()
		default:
			return errors.New("Wrong result")
		}
	}

	var allResult Results

	for _, val := range team_map {
		allResult = append(allResult, *val)
	}
	sort.Sort(allResult)

	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, val := range allResult {
		fmt.Fprintf(writer, "%-30s |%3d |%3d |%3d |%3d |%3d\n", val.Name, val.MP, val.W, val.D, val.L, val.P)
	}

	return nil
}
