package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"os/exec"
	ps "github.com/mitchellh/go-ps"
)

type Command struct {
	cmd		string
	args	string
}

type Shell struct {
	commands[]Command
}

func New() *Shell {
	return &Shell{}
}

func (s *Shell) GetCommands() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("shell$ ")
	line, _, _ := reader.ReadLine()
	
	for _, cmd := range strings.Split(string(line), "|") {
		command := strings.Split(string(cmd), " ")
		s.commands = append(s.commands, Command{cmd: command[0], args: strings.Join(command[1:], " ")})
	}
}

func (s *Shell) ChangeDir(cmd Command) error {
	path := strings.Split(cmd.args, " ")[0]
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
	pid, err := strconv.Atoi(strings.Split(cmd.args, " ")[0])
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
        log.Printf("%d\t%s\n",proc.Pid(),proc.Executable())
    }

	return nil
}

func (s *Shell) Exec(cmd Command) error {
	exec := exec.Command(cmd.cmd, strings.Split(cmd.args, " ")...)
	stdout, err := exec.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := exec.Start(); err != nil {
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
	if err := exec.Wait(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Shell) ExecCommands() error {
	var err error

	for _, cmd := range s.commands {
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
		default:
			err = s.Exec(cmd)
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
