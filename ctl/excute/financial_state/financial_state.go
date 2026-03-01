package financial_state

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/spf13/cobra"
)

func FinancialStateCmd(cmd *cobra.Command, args []string) {
	states, err := readCsv(FileStr)
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
	var x []string
	for i := 0; i < 12; i++ {
		x = append(x, strconv.Itoa(i+1)+"月")
	}
	for _, state := range states {
		line.SetXAxis(x).AddSeries(state.Title, state.lineDatas)
	}
	create, err := os.Create(OutStr)
	if err != nil {
		panic(err)
	}
	_ = line.Render(create)
}

func readCsv(path string) ([]FinancialState, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	financialStates := make([]FinancialState, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) < 13 {
			return nil, errors.New(fmt.Sprintf("数据不足13个，%v", record))
		}
		financialState := FinancialState{
			Title:     record[0],
			lineDatas: generateLineDatas(record[1:13]...),
		}
		financialStates = append(financialStates, financialState)
	}
	return financialStates, nil
}
