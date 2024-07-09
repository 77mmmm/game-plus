package redis

const (
	DefaultIdleTimeSeconds = 10
	DefaultMaxActive       = 100
	DefaultMaxIdle         = 20

	DefaultExpireTimeSeconds = 60
	DefaultWatchDogSeconds   = 10
)

type ClientOptions struct {
	maxIdle            int //最大空闲连接数
	maxActive          int //最大激活连接数
	wait               bool
	idleTimeoutSeconds int

	//必填参数
	password string
	network  string
	address  string
}

type ClientOption func(c *ClientOptions)

func WithMaxIdle(maxIdle int) ClientOption {
	return func(c *ClientOptions) {
		c.maxIdle = maxIdle
	}
}

func WithMaxActive(maxActive int) ClientOption {
	return func(c *ClientOptions) {
		c.maxActive = maxActive
	}
}

func WithWaitMode() ClientOption {
	return func(c *ClientOptions) {
		c.wait = true
	}
}

func repairClient(c *ClientOptions) {
	if c.maxIdle < 0 {
		c.maxIdle = DefaultIdleTimeSeconds
	}
	if c.idleTimeoutSeconds < 0 {
		c.idleTimeoutSeconds = DefaultMaxIdle
	}

	if c.maxActive < 0 {
		c.maxActive = DefaultMaxActive
	}

}
