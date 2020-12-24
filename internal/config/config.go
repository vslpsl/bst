package config

type Config struct {
	HTTPPort          int    `json:"httpPort"`
	BSTSourceFilePath string `json:"bstSourceFilePath"`
}
