module k8s.agent

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.7.2
	github.com/jameskeane/bcrypt v0.0.0-20120420032655-c3cd44c1e20f
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/metrics v0.17.0
)
