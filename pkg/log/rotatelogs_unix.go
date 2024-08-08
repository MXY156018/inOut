//go:build !windows
// +build !windows

package log

import (
	"mall-pkg/config"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer(c config.Zap) (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(c.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(c.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if c.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
