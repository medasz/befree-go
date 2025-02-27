package lib

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/medasz/clash-kernel/config"
	C "github.com/medasz/clash-kernel/constant"
	"github.com/medasz/clash-kernel/hub"
	"github.com/medasz/clash-kernel/hub/executor"
	"github.com/medasz/clash-kernel/log"

	"go.uber.org/automaxprocs/maxprocs"
)

type Kernel struct {
	flagset            map[string]bool
	version            bool
	testConfig         bool
	homeDir            string
	configFile         string
	externalUI         string
	externalController string
	secret             string
}

func (kernel *Kernel) Run() {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	if kernel.version {
		fmt.Printf("Clash %s %s %s with %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), C.BuildTime)
		return
	}

	if kernel.homeDir != "" {
		if !filepath.IsAbs(kernel.homeDir) {
			currentDir, _ := os.Getwd()
			kernel.homeDir = filepath.Join(currentDir, kernel.homeDir)
		}
		C.SetHomeDir(kernel.homeDir)
	}

	if kernel.configFile != "" {
		if !filepath.IsAbs(kernel.configFile) {
			currentDir, _ := os.Getwd()
			kernel.configFile = filepath.Join(currentDir, kernel.configFile)
		}
		C.SetConfig(kernel.configFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}
	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	if kernel.testConfig {
		if _, err := executor.Parse(); err != nil {
			log.Errorln(err.Error())
			fmt.Printf("configuration file %s test failed\n", C.Path.Config())
			os.Exit(1)
		}
		fmt.Printf("configuration file %s test is successful\n", C.Path.Config())
		return
	}

	var options []hub.Option
	if kernel.flagset["ext-ui"] {
		options = append(options, hub.WithExternalUI(kernel.externalUI))
	}
	if kernel.flagset["ext-ctl"] {
		options = append(options, hub.WithExternalController(kernel.externalController))
	}
	if kernel.flagset["secret"] {
		options = append(options, hub.WithSecret(kernel.secret))
	}

	if err := hub.Parse(options...); err != nil {
		log.Fatalln("Parse config error: %s", err.Error())
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

func NewClashKernel(options ...Option) *Kernel {
	ck := &Kernel{}
	for _, option := range options {
		option(ck)
	}
	return ck
}

type Option func(kernel *Kernel)

// WithHomeDir set configuration directory
func WithHomeDir(homeDir string) Option {
	return func(kernel *Kernel) {
		kernel.homeDir = homeDir
	}
}

// WithConfigFile specify configuration file
func WithConfigFile(configFile string) Option {
	return func(kernel *Kernel) {
		kernel.configFile = configFile
	}
}

// WithExternalUI override external ui directory
func WithExternalUI(externalUI string) Option {
	return func(kernel *Kernel) {
		kernel.externalUI = externalUI
	}
}

// WithExternalController override external controller address
func WithExternalController(externalController string) Option {
	return func(kernel *Kernel) {
		kernel.externalController = externalController
	}
}

// WithVersion show current version of clash
func WithVersion(version bool) Option {
	return func(kernel *Kernel) {
		kernel.version = version
	}
}

// WithTestConfig test configuration and exit
func WithTestConfig(testConfig bool) Option {
	return func(kernel *Kernel) {
		kernel.testConfig = testConfig
	}
}

// WithSecret override secret for RESTful API
func WithSecret(secret string) Option {
	return func(kernel *Kernel) {
		kernel.secret = secret
	}
}
