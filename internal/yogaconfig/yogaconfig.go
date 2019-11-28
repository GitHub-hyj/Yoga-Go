package yogaconfig

import (
	"github.com/GitHub-hyj/Yoga-Go/yogautil"
	"github.com/GitHub-hyj/Yoga-Go/yogaverbose"
	"os"
	"path/filepath"
	"runtime"
)

const (
	// EnvConfigDir 配置路径环境变量
	EnvConfigDir = "YOGA_GO_CONFIG_DIR"
	// ConfigName 配置文件名
	ConfigName = "yoga_config.json"
)

var (
	pcsConfigVerbose = yogaverbose.NewYogaVerbose("YOGACONFIG")
	configFilePath   = filepath.Join(GetConfigDir(), ConfigName)
)


// GetConfigDir 获取配置路径
func GetConfigDir() string {
	// 从环境变量读取
	configDir, ok := os.LookupEnv(EnvConfigDir)
	if ok {
		if filepath.IsAbs(configDir) {
			return configDir
		}
		// 如果不是绝对路径, 从程序目录寻找
		return yogautil.ExecutablePathJoin(configDir)
	}

	// 使用旧版
	// 如果旧版的配置文件存在, 则使用旧版
	oldConfigDir := yogautil.ExecutablePath()
	_, err := os.Stat(filepath.Join(oldConfigDir, ConfigName))
	if err == nil {
		return oldConfigDir
	}

	switch runtime.GOOS {
	case "windows":
		dataPath, ok := os.LookupEnv("APPDATA")
		if !ok {
			pcsConfigVerbose.Warn("Environment APPDATA not set")
			return oldConfigDir
		}
		return filepath.Join(dataPath, "BaiduPCS-Go")
	default:
		dataPath, ok := os.LookupEnv("HOME")
		if !ok {
			pcsConfigVerbose.Warn("Environment HOME not set")
			return oldConfigDir
		}
		configDir = filepath.Join(dataPath, ".config", "Yoga-Go")

		// 检测是否可写
		err = os.MkdirAll(configDir, 0700)
		if err != nil {
			pcsConfigVerbose.Warnf("check config dir error: %s\n", err)
			return oldConfigDir
		}
		return configDir
	}
}
