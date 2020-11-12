package registry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"plugin"

	"github.com/alessio/go-plugins-ex/module"
)

var registry map[string]registryEntry

type registryEntry struct {
	path   string
	module module.Interface
}

func LoadFromDir(pluginsDir string) error {
	registry = make(map[string]registryEntry)

	return loadPlugins(pluginsDir)
}

func Lookup(name string) (module.Interface, string) {
	m, ok := registry[name]
	if !ok {
		return nil, ""
	}

	return m.module, m.path
}

func Modules() (names []string) {
	for key := range registry {
		names = append(names, key)
	}

	return
}

func loadPlugins(pluginsDir string) error {
	entries, err := ioutil.ReadDir(pluginsDir)
	if err != nil {
		return fmt.Errorf("couldn't read %s: %w", pluginsDir, err)
	}

	for _, entry := range entries {
		filename := filepath.Join(pluginsDir, entry.Name())

		plug, err := plugin.Open(filename)
		if err != nil {
			// couldn't open filename
			continue
		}

		entryPoint, err := plug.Lookup("Loader")
		if err != nil {
			// couldn't lookup the entry point in filename
			continue
		}

		loader, ok := entryPoint.(module.Loader)
		if !ok {
			// type assertion failed for plugin
			continue
		}

		mod := loader.Load()
		if _, ok := registry[mod.Name()]; ok {
			return fmt.Errorf("duplicate module name: %s", mod.Name())
		}

		registry[mod.Name()] = registryEntry{path: filename, module: mod}
	}

	return nil
}
