package task

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type Example struct {
	Description string
	Command     string
}

type Task struct {
	LookupPath string
	Name       string `yaml:"-"`
	Summary    string
	Command    string
	Script     string
	Exec       string
	Usage      string
	Examples   []*Example
}

func (t *Task) Run(args []string) error {
	if len(t.Exec) != 0 {
		return t.RunExec(args)
	}
	if len(t.Script) != 0 {
		return t.RunScript(args)
	}
	if len(t.Command) != 0 {
		return t.RunCommand(args)
	}
	return fmt.Errorf("nothing to run (add script, command, or exec key)")
}

func (t *Task) RunScript(args []string) error {
	path := filepath.Join(t.LookupPath, t.Script)
	args = append([]string{path}, args...)
	cmd := exec.Command("sh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (t *Task) RunCommand(args []string) error {
	args = append([]string{"-c", t.Command, "sh"}, args...)
	cmd := exec.Command("sh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (t *Task) RunExec(args []string) error {
	fields := strings.Fields(t.Exec)
	bin := fields[0]
	if path, err := exec.LookPath(bin); err != nil {
		return err
	}
	args = append(fields, args...)
	return syscall.Exec(path, args, os.Environ())
}
