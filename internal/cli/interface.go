package cli

import "io"

type env = map[string]string

type Vault interface {
	Create(name string, body io.Reader) error
	Get(names []string) (env, error)
	List() ([]string, error)
}
