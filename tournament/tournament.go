package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type League struct {
	LeaderBoard []Team
	Teams       map[string]Team
}

func (l *League) RegisterGame(game Game) {
	if _, ok := l.Teams[game.won.name]; !ok {
		l.Teams[game.won.name] = game.won
	}
	if _, ok := l.Teams[game.lost.name]; !ok {
		l.Teams[game.lost.name] = game.lost
	}
	var (
		t1, t2 Team
	)
	if game.draw {
		t1 = l.Teams[game.lost.name]
		t1.p++
		t1.d++
		t2 = l.Teams[game.won.name]
		t2.p++
		t2.d++
	} else {
		t1 = l.Teams[game.lost.name]
		t1.l++
		t2 = l.Teams[game.won.name]
		t2.w++
		t2.p += 3
	}
	t1.mp++
	t2.mp++
	l.Teams[game.lost.name] = t1
	l.Teams[game.won.name] = t2
}

func (l *League) UpdateBoard() {
	lb := []Team{}
	for _, t := range l.Teams {
		lb = append(lb, t)
	}
	sort.Slice(lb, func(i, j int) bool {
		if lb[i].p == lb[j].p {
			return lb[i].name < lb[j].name
		}
		return lb[i].p > lb[j].p
	})
	l.LeaderBoard = lb
}

func (l *League) PrintBoard() []byte {
	l.UpdateBoard()

	var maxNameLength int
	for k := range l.Teams {
		if len(k) > maxNameLength {
			maxNameLength = len(k)
		}
	}

	nameColumn := maxNameLength + 8
	formatStr := "%-*s| %2v | %2v | %2v | %2v | %2v\n"

	tally := fmt.Sprintf(formatStr, nameColumn, "Team", "MP", "W", "D", "L", "P")

	var teams []string
	for _, team := range l.LeaderBoard {
		teams = append(teams, fmt.Sprintf(formatStr, nameColumn, team.name, team.mp, team.w, team.d, team.l, team.p))
	}

	for _, t := range teams {
		tally = fmt.Sprintf("%s%s", tally, t)
	}

	return []byte(tally)
}

type Team struct {
	name           string
	mp, w, d, l, p int
}

type Game struct {
	draw      bool
	won, lost Team
}

func Tally(reader io.Reader, writer io.Writer) error {
	input, _ := ioutil.ReadAll(reader)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var league League = League{
		Teams: make(map[string]Team),
	}

	for _, line := range lines {
		if len(line) < 1 || line[0] == '#' {
			continue //skip comment line
		}

		splitLine := strings.Split(line, ";")
		if len(splitLine) < 3 {
			return fmt.Errorf("cannot parse line")
		}

		team1, team2, result := Team{name: splitLine[0]}, Team{name: splitLine[1]}, splitLine[2]
		if game, err := createGameResult(result, team1, team2); err != nil {
			return err
		} else {
			league.RegisterGame(game)
		}
	}

	writer.Write(league.PrintBoard())

	return nil
}

func createGameResult(result string, team1 Team, team2 Team) (Game, error) {
	var game Game
	switch result {
	case "draw":
		game = Game{
			draw: true,
			won:  team1,
			lost: team2,
		}
	case "win":
		game = Game{
			draw: false,
			won:  team1,
			lost: team2,
		}
	case "loss":
		game = Game{
			draw: false,
			won:  team2,
			lost: team1,
		}
	default:
		return Game{}, fmt.Errorf("can not parse game")
	}
	return game, nil
}
