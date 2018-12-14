package user

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/msteinert/pam"
)

const etcPasswdPath = "/etc/passwd"

func ListUsers() ([]string, error) {
	file, err := os.Open(etcPasswdPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := make([]string, 0)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if strings.Index(line, "#") == 0 {
			continue
		}

		split := strings.Split(line, ":")
		if len(split) > 0 {
			users = append(users, split[0])
		}
	}

	return users, nil
}

func AuthenticateUser(name, passwd string) error {
	t, err := pam.StartFunc("", name, func(s pam.Style, msg string) (string, error) {
		return passwd, nil
	})
	if err != nil {
		return err
	}

	err = t.Authenticate(0)
	if err != nil {
		return err
	}

	return nil
}
