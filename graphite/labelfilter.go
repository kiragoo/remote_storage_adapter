package graphite

import (
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/log"
	"strings"
)


// filter the MetricsNameLabel just adapte to the graphite plaintext

func LabelsFilter(m model.Metric) (model.Metric ,[]model.LabelName) {
	// metric
	var metric model.Metric = m

	// metric labels wanted to save out
	var mlabels = MetricLabelsConst()
	var metriclabelvalue = strings.ToLower(string(metric[model.MetricNameLabel]))

	if _, ok := mlabels[metriclabelvalue]; ok {
		// judge the metric name "__name__" which is replaced by the config about relabel of remote_write in the metrcis labels wanted

		for label, _ := range metric {

			if contains(label, mlabels[string(metriclabelvalue)].labels) {
				delete(metric, label)
			}
		}
		return  metric, mlabels[string(metriclabelvalue)].labels
	} else {
		log.Error("there is no labels need to filter")
	}
	return metric, mlabels[string(metriclabelvalue)].labels
}

func contains (l model.LabelName ,lns []model.LabelName) bool {
	var r bool = false
	for _, i := range lns {
		if i != l {
			r = true
		}else {
			// 若匹配上需要重置 r 为 false
			r = false
			// 当发现包含的时候立马跳出 for 循环，避免与下个 lns 中的元素比较将将变量 r 重置为 true。
			break
		}
	}

	return r
}

