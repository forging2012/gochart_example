package main

import (
	"math"
	"strconv"

	"github.com/fananchong/gochart"
	"github.com/shirou/gopsutil/mem"
)

type ChartMemory struct {
	gochart.ChartTime
}

func NewChartMemory() *ChartMemory {
	m, _ := mem.VirtualMemory()
	this := &ChartMemory{}
	this.RefreshTime = DEFAULT_REFRESH_TIME
	this.SampleNum = DEFAULT_SAMPLE_NUM
	this.ChartType = "line"
	this.Title = "内存占用"
	this.SubTitle = "内存大小: " + strconv.Itoa(int(math.Ceil(float64(m.Total)/float64(1024*1024*1024)))) + "GB"
	this.YAxisText = "memory"
	this.YMax = "100"
	this.ValueSuffix = "%"
	return this

}

func (this *ChartMemory) Update(now int64) map[string][]interface{} {
	datas := make(map[string][]interface{})
	m, _ := mem.VirtualMemory()
	datas["memory"] = []interface{}{int(m.UsedPercent)}
	return datas
}
