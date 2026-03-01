package financial_state

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	FileStr string
	OutStr  string
)

type FinancialState struct {
	Title     string
	lineDatas []opts.LineData
}

func generateLineDatas[T comparable](items ...T) []opts.LineData {
	strs := make([]string, 0, len(items))
	for _, item := range items {
		strs = append(strs, fmt.Sprintf("%v", item))
	}
	out := make([]opts.LineData, 0, len(items))
	for _, item := range items {
		out = append(out, opts.LineData{Value: item})
	}
	return out
}
