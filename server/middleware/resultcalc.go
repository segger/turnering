package middleware

import (
	"server/models"
	"sort"
)

func tournamentPoints(placement int, result models.ContestResult) float32 {
	if result.ClassNbr == 2 {
		return tournamentPointsNW2(placement, result.Points, result.Sse, result.Errors)
	} else {
		return tournamentPointsNW1(placement, result.Points, result.Sse)
	}
}

func tournamentPointsNW1(placement int, points float32, sse bool) float32 {
	tp := 3
	if points > 0 {
		tp += 7
	}
	if points > 0 && placement >= 1 && placement <= 5 {
		tp += (6 - placement)
	}
	if sse {
		tp += 5
	}
	return float32(tp)
}

func tournamentPointsNW2(placement int, points float32, sse bool, errors int32) float32 {
	tp := points
	if points == float32(25.0) && placement >= 1 && placement <= 5 {
		tp += float32(6 - placement)
	}
	if sse {
		tp += float32(5)
	}
	tp -= float32(errors)
	if tp < float32(3.0) {
		return float32(3.0)
	}
	return tp
}

func sortAndCalc(event map[string][]models.ContestResult) map[string][]models.ContestResultSorted {
	sorted := make(map[string][]models.ContestResultSorted)

	for key, val := range event {
		sort.Slice(val, func(i, j int) bool {
			if val[i].Points < val[j].Points {
				return false
			}
			if val[i].Points > val[j].Points {
				return true
			}
			if val[i].Errors < val[j].Errors {
				return true
			}
			if val[i].Errors > val[j].Errors {
				return false
			}
			return val[i].Time < val[j].Time
		})

		var resultList []models.ContestResultSorted
		for idx, res := range val {
			var result models.ContestResultSorted
			result.Placement = idx + 1
			result.FirstName = res.FirstName
			result.DogName = res.DogName
			result.Points = res.Points
			result.Errors = res.Errors
			result.Time = res.Time
			result.Sse = res.Sse
			result.TP = tournamentPoints((idx + 1), res)
			resultList = append(resultList, result)
		}
		sorted[key] = resultList
	}

	return sorted
}

func Calc(input []models.ContestResult) []map[string][]models.ContestResultSorted {

	nw1 := make(map[string][]models.ContestResult)
	nw2 := make(map[string][]models.ContestResult)
	for _, res := range input {
		if res.ClassNbr == 1 {
			nw1[res.EventName] = append(nw1[res.EventName], res)
		} else if res.ClassNbr == 2 {
			nw2[res.EventName] = append(nw2[res.EventName], res)
		}
	}

	sorted_nw1 := sortAndCalc(nw1)
	sorted_nw2 := sortAndCalc(nw2)

	// log.Printf("sorted 1 %v", sorted_nw1)
	// log.Printf("sorted 2 %v", sorted_nw2)

	result := []map[string][]models.ContestResultSorted{sorted_nw1, sorted_nw2}
	return result
}
