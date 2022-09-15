package structure

type Channel struct {
	Name          string   `yaml:"name" json:"name"`
	Consortium    string   `yaml:"consortium" json:"consortium"`
	Applications  []string `yaml:"applications" json:"applications"`
	Organizations []string `yaml:"organizations" json:"organizations"`
	applications  []*Application
	organizations []*Organization
}
