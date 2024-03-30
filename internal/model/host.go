package model

type Host struct {
	ID             int    `yaml:"-"`
	Alias          string `yaml:"alias"`
	User           string `yaml:"user"`
	HostName       string `yaml:"hostname"`
	Description    string `yaml:"description,omitempty"`
	PrivateKeyPath string `yaml:"identity_file_path,omitempty"`
}

func NewHost(id int, alias, user, hostname, desc, privateKeyPath string) Host {
	return Host{
		ID:             id,
		Alias:          alias,
		User:           user,
		HostName:       hostname,
		Description:    desc,
		PrivateKeyPath: privateKeyPath,
	}
}
