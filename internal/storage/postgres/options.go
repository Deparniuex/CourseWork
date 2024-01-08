package postgres

type Option func(*Settings)

func WithPort(port string) Option {
	return func(settings *Settings) {
		settings.port = port
	}
}

func WithHost(host string) Option {
	return func(settings *Settings) {
		settings.host = host
	}
}

func WithUser(user string) Option {
	return func(settings *Settings) {
		settings.user = user
	}
}
func WithPassword(password string) Option {
	return func(settings *Settings) {
		settings.password = password
	}
}
func WithDBName(dbname string) Option {
	return func(settings *Settings) {
		settings.dbname = dbname
	}
}
