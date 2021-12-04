package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
)

func parseToArray(text string) []string {
	lines := strings.Split(text, "\n")
	metarSlice := make([]string, 0, len(lines))
	metarStr := ""
	for _, line := range lines {
		if tafValidation.MatchString(line) {
			break
		}
		if !comment.MatchString(line) {
			metarStr += strings.Trim(line, " ")
		}
		if metarClose.MatchString(line) {
			metarSlice = append(metarSlice, metarStr)
			metarStr = ""
		}
	}
	return metarSlice
}

func extractWindDirection(metars []string) []string {
	winds := make([]string, 0, len(metars))
	for _, metar := range metars {
		if windRegex.MatchString(metar) {
			winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
		}
	}
	return winds
}

func mineWindDistribution(winds []string) {
	for _, wind := range winds {
		if variableWind.MatchString(wind) {
			for i := 0; i < 8; i++ {
				windDist[i]++
			}
		} else if validWind.MatchString(wind) {
			windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
			if d, err := strconv.ParseFloat(windStr, 64); err == nil {
				dirIndex := int(math.Round(d/45.0)) % 8
				windDist[dirIndex]++
			}

		}

	}
}
