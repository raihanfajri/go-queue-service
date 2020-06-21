package helpers

type RouteYML struct {
	QueueName string `yaml:"queue_name"`
	CosumedBy map[string]struct {
		Path string `yaml:"path"`
	} `yaml:"consumed_by"`
}
