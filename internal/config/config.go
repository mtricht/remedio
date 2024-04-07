package config

type specification struct {
	Port            int `default:"8080"`
	Username        string
	Password        string
	NotificationURL string `required:"true" split_words:"true"`
}

var Config specification
