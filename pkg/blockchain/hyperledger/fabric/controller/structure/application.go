package structure

type Application struct {
	Organizations []string `yaml:"organizations" json:"organizations"`
	organizations []*Organization
}
