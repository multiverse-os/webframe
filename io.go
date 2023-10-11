package maglev

import (
	config "github.com/multiverse-os/maglev/config"
	io "github.com/multiverse-os/maglev/io"
)

func DefaultOutputs(cfg config.Settings) io.Outputs {
	return append(TerminalOutput(cfg), LogOutput(cfg)...)
}

func LogOutput(cfg config.Settings) io.Outputs {
	return io.Outputs{}
}

func TerminalOutput(cfg config.Settings) io.Outputs {
	return io.Outputs{io.TerminalOutput().Prefix(cfg.Name)}
}

func (f Framework) Log(text ...interface{})   { f.Outputs.Log(io.LOG, text...) }
func (f Framework) Info(text ...interface{})  { f.Outputs.Log(io.INFO, text...) }
func (f Framework) Warn(text ...interface{})  { f.Outputs.Log(io.WARN, text...) }
func (f Framework) Error(text ...interface{}) { f.Outputs.Log(io.ERROR, text...) }
func (f Framework) Fatal(text ...interface{}) { f.Outputs.Log(io.FATAL, text...) }

func (f Framework) Debug(text ...interface{}) {
	if !f.IsEnvironment(Production) {
		f.Outputs.Log(io.DEBUG, text...)
	}
}

func (f Framework) Output(text ...interface{}) { f.Outputs.Write(text...) }
