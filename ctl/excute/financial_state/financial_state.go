package financial_state

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/spf13/cobra"
)

var miniNumber = 13

func FinancialStateCmd(_ *cobra.Command, _ []string) {
	headers, states, err := readCsv(FileStr)
	if err != nil {
		panic(err)
	}
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "财政综合",
			Subtitle: "净资产",
		}),
	)
	for _, state := range states {
		line.SetXAxis(headers).AddSeries(state.Title, state.lineDatas)
	}
	createFile, err := os.Create(OutStr)
	if err != nil {
		panic(err)
	}
	_ = line.Render(createFile)
}

func readCsv(path string) ([]string, []FinancialState, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	financialStates := make([]FinancialState, 0)
	headers := make([]string, 0, miniNumber-1)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if len(record) == 0 {
			continue
		}
		if len(headers) == 0 {
			headers = record[1:]
			continue
		}
		if len(record) < miniNumber {
			continue
		}
		financialState := FinancialState{
			Title:     record[0],
			lineDatas: generateLineDatas(record[1:miniNumber]...),
		}
		financialStates = append(financialStates, financialState)
	}
	return headers, financialStates, nil
}
