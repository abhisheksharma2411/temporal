package wideevents

import (
	"go.opentelemetry.io/otel/log"
)

const WorkerConfigEventName = "worker_config"

// WorkerConfigPayload describes the static properties of a worker: runtime,
// OS, and architecture.
type WorkerConfigPayload struct {
	Namespace         string
	TaskQueue         string
	WorkerInstanceKey string
	RuntimeType       string
	OS                string
	Architecture      string
}

func (p WorkerConfigPayload) EventName() string { return WorkerConfigEventName }

func (p WorkerConfigPayload) Attributes() []log.KeyValue {
	return []log.KeyValue{
		log.String("namespace", p.Namespace),
		log.String("task_queue", p.TaskQueue),
		log.String("worker_instance_key", p.WorkerInstanceKey),
		log.String("runtime_type", p.RuntimeType),
		log.String("os", p.OS),
		log.String("architecture", p.Architecture),
	}
}
