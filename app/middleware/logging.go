package middleware

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"lattice-manager-grpc/config"
	"os"
	"time"
)

func NewLoggingInterceptor(cfg *config.Config) grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithTimestampFormat(zerolog.TimeFormatUnix),
		// Add any other option (check functions starting with logging.With).
	}

	filename := fmt.Sprintf("go-grpc-web_%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log logFile")
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)
	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.MultiLevelWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout},
		))
	} else {
		log.Logger = log.Output(file)
	}

	return logging.UnaryServerInterceptor(InterceptorLogger(log.Logger), opts...)
}

// InterceptorLogger adapts zero_log logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.With().Fields(fields).Logger()

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Msg(msg)
		case logging.LevelInfo:
			l.Info().Msg(msg)
		case logging.LevelWarn:
			l.Warn().Msg(msg)
		case logging.LevelError:
			l.Error().Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
