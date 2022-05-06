package main

import (
	"io/ioutil"
	"log"
	"os"
	"plugin"
	"regexp"

	"github.com/sgettys/bellman/pkg/hub"
	"github.com/sgettys/bellman/pkg/types"
)

// Search in the plugin dir for individual plugin subdirectories.
// i.e. "./plugins/pluginfoo and ./plugins/pluginbarr" would each be loaded
func listPlugins(dir, pattern string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	plugins := []os.FileInfo{}
	for _, file := range files {
		if file.IsDir() {
			plugins, err = listPlugins(file.Name(), pattern)
		}
		matched, err := regexp.MatchString(pattern, file.Name())
		if err != nil {
			return nil, err
		}
		if matched {
			plugins = append(plugins, file)
		}
	}
	return plugins, nil
}

func registerPlugins(hub *hub.Hub, plugins []os.FileInfo) error {

	d := types.InData{V: 1}
	log.Printf("Invoking pipeline with data: %#v\n", d)
	o := types.OutData{}
	for _, p := range plugins {
		p, err := plugin.Open(p.Name())
		if err != nil {
			log.Fatal(err)
		}
		pName, err := p.Lookup("Name")
		if err != nil {
			return err
		}
		log.Printf("Invoking plugin: %s\n", *pName.(*string))

		input, err := p.Lookup("Input")
		if err != nil {
			return err
		}
		f, err := p.Lookup("F")
		if err != nil {
			return err
		}

		*input.(*types.InData) = d
		f.(func())()

		output, err := p.Lookup("Output")
		if err != nil {
			return err
		}
		// Feed the output to the next plugin's input
		d = types.InData{V: output.(*types.OutData).V}
		*input.(*types.InData) = d

		o = *output.(*types.OutData)
	}
	log.Printf("Final result: %#v\n", o)
	return nil
}
