package search

import (
	"fmt"
	"strings"

	q "github.com/core-go/sql"
)

func BuildRateQuery(filter interface{}) (query string, params []interface{}) {
	query = `select * from rate`
	s := filter.(*RateFilter)
	var where []string

	i := 1
	if s.Time != nil {
		if s.Time.Min != nil {
			where = append(where, fmt.Sprintf(`time >= %s`, q.BuildDollarParam(i)))
			params = append(params, s.Time.Min)
			i++
		}
		if s.Time.Max != nil {
			where = append(where, fmt.Sprintf(`time <= %s`, q.BuildDollarParam(i)))
			params = append(params, s.Time.Max)
			i++
		}
	}
	if len(s.Id) > 0 {
		where = append(where, fmt.Sprintf(`id = %s`, q.BuildDollarParam(i)))
		params = append(params, s.Id)
		i++
	}
	if len(s.Author) > 0 {
		where = append(where, fmt.Sprintf(`author = %s`, q.BuildDollarParam(i)))
		params = append(params, s.Author)
		i++
	}
	if len(s.Rate) > 0 {
		where = append(where, fmt.Sprintf(`rate = %s`, q.BuildDollarParam(i)))
		params = append(params, s.Rate)
		i++
	}
	if len(s.Review) > 0 {
		where = append(where, fmt.Sprintf(`review ilike %s`, q.BuildDollarParam(i)))
		params = append(params, "%"+s.Review+"%")
		i++
	}
	if len(s.UsefulCount) > 0 {
		where = append(where, fmt.Sprintf(`usefulCount = %s`, q.BuildDollarParam(i)))
		params = append(params, s.UsefulCount)
		i++
	}
	if len(s.ReplyCount) > 0 {
		where = append(where, fmt.Sprintf(`commentCount = %s`, q.BuildDollarParam(i)))
		params = append(params, s.ReplyCount)
		i++
	}

	if len(where) > 0 {
		query = query + ` where ` + strings.Join(where, " and ")
	}
	return
}
