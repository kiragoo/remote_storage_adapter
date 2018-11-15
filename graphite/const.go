package graphite

// metric info const
import "github.com/prometheus/common/model"

//const NameSpace  = "PAAS"

type MetricInfo struct {
	name string
	labels []model.LabelName
}

type MetricLabels map[string]MetricInfo


func MetricLabelsConst() MetricLabels {
	var mlabels MetricLabels
	mlabels = make(map[string]MetricInfo)
	mlabels["node_memory_memtotal"] = MetricInfo{"node_memory_memtotal", []model.LabelName{"instance"}}
	mlabels["node_memory_memfree"] = MetricInfo{"node_memory_memfree", []model.LabelName{"instance"}}
	mlabels["node_memory_cached"] = MetricInfo{"node_memory_cached", []model.LabelName{"instance"}}
	mlabels["node_memory_buffers"] = MetricInfo{"node_memory_buffers", []model.LabelName{"instance"}}
	mlabels["node_cpu"] = MetricInfo{"node_cpu", []model.LabelName{"instance","cpu","mode"}}
	return mlabels
}





