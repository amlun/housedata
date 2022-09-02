package configs

type JobConfig struct {
	Name     string
	SavePath string
	FileName string
	Regions  map[string]string
	URL      string
	URLExtra string
	Selector string
}
