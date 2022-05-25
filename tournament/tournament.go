package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Map = map[string]int

type Stat struct {
	Name   string
	Played int
	Won    int
	Lost   int
	Draw   int
	Points int
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*Stat)

	pointsMap := map[string]int{
		"win":  3,
		"draw": 1,
		"loss": 0,
	}

	rsInverts := map[string]string{
		"win":  "loss",
		"draw": "draw",
		"loss": "win",
	}

	input, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ";")

		// skip new line.
		if line == "" {
			continue
		}

		// skip comment.
		if strings.HasPrefix(line, "#") {
			continue
		}

		if len(parts) < 3 {
			return errors.New("invalid result")
		}

		t1 := parts[0]
		t2 := parts[1]
		rs := parts[2]
		rsInvert := rsInverts[rs]

		if _, ok := pointsMap[rs]; !ok {
			return errors.New("invalid result")
		}

		if _, ok := teams[t1]; !ok {
			teams[t1] = &Stat{
				Name:   t1,
				Played: 0,
				Won:    0,
				Lost:   0,
				Draw:   0,
				Points: 0,
			}
		}

		if _, ok := teams[t2]; !ok {
			teams[t2] = &Stat{
				Name:   t2,
				Played: 0,
				Won:    0,
				Lost:   0,
				Draw:   0,
				Points: 0,
			}
		}

		if team, ok := teams[t1]; ok {
			team.Played += 1
			team.Points += pointsMap[rs]

			if rs == "win" {
				team.Won += 1
			} else if rs == "loss" {
				team.Lost += 1
			} else {
				team.Draw += 1
			}
		}

		if team, ok := teams[t2]; ok {
			team.Played += 1
			team.Points += pointsMap[rsInvert]

			if rs == "win" {
				team.Lost += 1
			} else if rs == "loss" {
				team.Won += 1
			} else {
				team.Draw += 1
			}
		}
	}

	teamList := make([]Stat, 0)
	for _, team := range teams {
		teamList = append(teamList, *team)
	}

	sort.SliceStable(teamList, func(i, j int) bool {
		// if points are equal, sort by name, alphabetically.
		if teamList[i].Points == teamList[j].Points {
			return teamList[i].Name < teamList[j].Name
		}

		// sort by points, descending.
		return teamList[i].Points > teamList[j].Points
	})

	output := ""

	output += fmt.Sprintf(
		"%-30s |%3s |%3s |%3s |%3s |%3s\n",
		"Team",
		"MP",
		"W",
		"D",
		"L",
		"P",
	)

	for _, team := range teamList {
		output += fmt.Sprintf(
			"%-30s |%3d |%3d |%3d |%3d |%3d\n",
			team.Name,
			team.Played,
			team.Won,
			team.Draw,
			team.Lost,
			team.Points,
		)
	}

	io.WriteString(writer, output)

	return nil
}
