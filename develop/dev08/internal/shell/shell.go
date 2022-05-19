package shell

import (
	"bufio"
	"bytes"

	// "bytes"
	"fmt"
	"io"

	// "io"
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
	input	*io.PipeReader
	output	*io.PipeWriter
}

func New() *Shell {
	return &Shell{}
}

func (s *Shell) GetCommands() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("shell$ ")
	line, _, _ := reader.ReadLine()

	for _, cmd := range strings.Split(string(line), "|") {
		cmd = strings.Trim(cmd,  " ")
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
	fmt.Fprintln(s.output, dir)

	return nil
}

func (s *Shell) Echo(cmd Command) {
	go func() {
		defer s.output.Close()
		fmt.Fprintln(s.output, strings.Join(cmd.args, " "))
	}()
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
		fmt.Fprintf(s.output, "%d\t%s\n", proc.Pid(), proc.Executable())
	}

	return nil
}

func (s *Shell) Exec(cmd Command) error {
	defer s.output.Close()
	exec := exec.Command(cmd.cmd, cmd.args...)
	
	exec.Stdout = s.output

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

	if err := exec.Wait(); err != nil {
		return fmt.Errorf("executable failed: %v", err)
	}

	close(sigquit)
	defer signal.Reset()

	return nil
}

func (s *Shell) ExecCommands() error {
	var err error

	buffer := new(bytes.Buffer)
	for _ , cmd := range s.commands {
		s.input, s.output = io.Pipe()
		go func () {
			defer s.output.Close()
			fmt.Fprint(s.output, buffer.String())
		}()
		buffer.Reset()
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
			go s.Exec(cmd)
		}
		

		buffer.ReadFrom(s.input)
		fmt.Print(buffer.String(), "AAA")

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
