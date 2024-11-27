package bconf

import "strings"

type Config struct {
	Name    string
	Values  []string
	MainArg string
}

func Parse(input string) ([]*Config, error) {
	var configs []*Config

	root := &Config{
		Name: "root",
	}
	configs = append(configs, root)
	var currentCong = root

	for i := 0; i < len(input); i++ {
		if input[i] == '#' {
			for i < len(input) && input[i] != '\n' {
				i++
			}
		}

		if input[i] == ':' {
			pos := i
			for i > 0 && !(input[i] == ' ' || input[i] == '\n') {
				i--
			}

			currentCong = &Config{
				Name: strings.TrimSpace(input[i:pos]),
			}
			configs = append(configs, currentCong)

			i = pos
			continue
		}

		if input[i] == '+' {
			pos := i
			for i < len(input) && input[i] != '\n' {
				i++
			}

			dat := strings.TrimSpace(input[pos+1 : i])
			currentCong.Values = append(currentCong.Values, dat)
		}

		if i < len(input) && input[i] == '^' {
			pos := i
			for i < len(input) && input[i] != '\n' {
				i++
			}

			dat := strings.TrimSpace(input[pos+1 : i])
			currentCong.MainArg = dat
		}

		if i < len(input) && (input[i] == ';' || (i+1 < len(input) && input[i] == '\n' && input[i+1] == '\n') && currentCong != root) {
			currentCong = root
		}
	}

	return configs, nil
}
