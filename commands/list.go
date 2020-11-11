package commands

import (
	"fmt"
	"github.com/alessio/go-plugins-ex/module"
	"io/ioutil"
	"path/filepath"
	"plugin"
	"sort"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list plugins",
	Long:  `This command lists available plugins.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		plugins, err := availablePlugins()
		if err != nil {
			return err
		}

		for _, plug := range plugins {
			fmt.Println(plug)
		}

		return nil
	},
}

type pluginT struct {
	path   string
	plug   *plugin.Plugin
	module.Interface
}

func (p *pluginT) Filename() string       { return p.path }
func (p *pluginT) Loadable() bool         { return p.plug != nil }
func (p *pluginT) String() string         { return p.Filename() }
func (p *pluginT) Plugin() *plugin.Plugin { return p.plug }

func (p *pluginT) Compatible() bool {
	moduleFunc, err := p.plug.Lookup("Module")
	if err != nil {
		return false
	}

	getter, ok := moduleFunc.(module.Getter)
	if !ok {
		return false
	}

	p.Interface = getter()

	return true
}

type Plugins []*pluginT

func (p Plugins) Len() int           { return len(p) }
func (p Plugins) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Plugins) Less(i, j int) bool { return p[i].Filename() < p[j].Filename() }

func NewPlugin(filename string) *pluginT {
	plug, err := plugin.Open(filepath.Join(pluginsDir, filename))
	if err != nil {
		return &pluginT{
			path:      filename,
			plug:      nil,
			Interface: nil,
		}
	}

	return &pluginT{path: filename, plug: plug}
}

func availablePlugins() (Plugins, error) {
	var plugins Plugins

	entries, err := ioutil.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s: %w", pluginsDir, err)
	}

	for _, entry := range entries {
		plug := NewPlugin(entry.Name())
		if plug.Loadable() {
			plugins = append(plugins, plug)
		}
	}

	sort.Sort(plugins)

	return plugins, nil
}

