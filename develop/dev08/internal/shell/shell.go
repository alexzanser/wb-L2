package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	ps "github.com/mitchellh/go-ps"
)

type Command struct {
	cmd  string
	args []string
}

type Shell struct {
	commands []Command
	buf	 *bytes.Buffer
}

func New() *Shell {
	var buf bytes.Buffer
	return &Shell{buf: &buf}
}

func (s *Shell) GetCommands() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("shell$ ")
	line, _, _ := reader.ReadLine()

	for _, cmd := range strings.Split(string(line), "|") {
		command := strings.Split(string(cmd), " ")
		s.commands = append(s.commands, Command{cmd: command[0], args: command[1:]})
	}
}

func (s *Shell) ChangeDir(cmd Command) error {
	path := cmd.args[0]
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("error when change dir: %v", err)
	}

	return nil
}

func (s *Shell) PWD() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error when get PWD: %v", err)
	}
	fmt.Println(dir)

	return nil
}

func (s *Shell) Echo(cmd Command) {
	fmt.Println(cmd.args)
}

func (s *Shell) Kill(cmd Command) error {
	pid, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return fmt.Errorf("incorrect PID: %v", err)
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		return fmt.Errorf("can't kill process %d: %v", pid, err)
	}

	return nil
}

func (s *Shell) PS(cmd Command) error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("can't get processes list: %v", err)
	}

	for _, proc := range processes {
		log.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
	}

	return nil
}

func (s *Shell) Exec(cmd Command) error {
	exec := exec.Command(cmd.cmd, cmd.args...)
	stdout, _ := exec.StdoutPipe()

	if err := exec.Start(); err != nil {
		return fmt.Errorf("can't start executable: %v", err)
	}

	sigquit := make(chan os.Signal, 1)
	go func () {
		signal.Notify(sigquit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <- sigquit:
			exec.Process.Kill()
		}
	}()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Fprintln(s.buf, scanner.Text())
	}
	
	if err := exec.Wait(); err != nil {
		return fmt.Errorf("executable failed: %v", err)
	}

	close(sigquit)
	defer signal.Reset()

	return nil
}

func (s *Shell) ExecCommands() error {
	var err error

	for i, cmd := range s.commands {
		switch cmd.cmd {
		case "cd":
			err = s.ChangeDir(cmd)
		case "pwd":
			err = s.PWD()
		case "echo":
			s.Echo(cmd)
		case "kill":
			err = s.Kill(cmd)
		case "ps":
			err = s.PS(cmd)
		case "exit":
			os.Exit(0)
		default:
			err = s.Exec(cmd)
		}

		if s.commands[i + 1].cmd != "|" {
			out, _ := s.buf.ReadLi
			fmt.Println(string(out))
		}
		if err != nil {
			return fmt.Errorf("error when exec command: %v", err)
		}
	}

	return nil
}

func (s *Shell) Run() {
	for {
		s.GetCommands()
		if s.commands[0].cmd != "" {
			err := s.ExecCommands()
			if err != nil {
				log.Println(fmt.Errorf("error: %v", err))
			}
		}
		s.commands = nil
	}
}
