package util

type Config struct {
	Mysql struct {
		User     string
		Password string
	}
	Option struct {
		Domain string
		Mail   string
	}
}
