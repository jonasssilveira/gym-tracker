package bot

type Command string

const (
	Start        Command = "/start"
	StartSeries  Command = "/start_series"
	FinishSeries Command = "/finish_series"
	Calculate    Command = "/calculate"
)
