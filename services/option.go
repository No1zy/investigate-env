package services

var (
	distPrefix = "dist"
	srv        = []service{
		service{"go", "go", distPrefix + "/go/"},
		service{"java", "java", distPrefix + "/java/src/main/"},
		service{"php", "php", distPrefix + "/php/"},
		service{"python", "py", distPrefix + "/python/"},
		service{"ruby", "rb", distPrefix + "/ruby/"},
		service{"perl", "pl", distPrefix + "/perl/"},
		service{"node", "js", distPrefix + "/node/"},
		service{"rails", "rb", distPrefix + "/rails/"},
		service{"php-apache", "php", distPrefix + "/php-apache/"},
	}
	defaultOptions = []Option{
		WithServices(srv),
	}
)

type Option func(*property) error

func WithServices(srvs []service) Option {
	return func(p *property) error {
		if p != nil {
			p.services = srvs
		}
		return nil
	}
}
