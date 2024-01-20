package lg

import "log/slog"

// Err Custom error handler
func Err(log *slog.Logger, op, msg string, err error) {
	log.Error(msg,
		slog.String("err", err.Error()),
		slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})
}
