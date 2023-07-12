// This file is part of the program "yant".
// Please see the LICENSE file for copyright information.

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/noisetorch/pulseaudio"
)

const (
	loaded = iota
	unloaded
	inconsistent
)

// the ugly and (partially) repeated strings are unforunately difficult to avoid, as it's what pulse audio expects

func updateNoiseSupressorLoaded(ctx *ntcontext) {
	c := ctx.paClient
	upd, err := c.Updates()
	if err != nil {
		fmt.Printf("Error listening for updates: %v\n", err)
	}

	for {
		ctx.noiseSupressorState, ctx.virtualDeviceInUse = supressorState(ctx)
		if !c.Connected() {
			break
		}

		<-upd
	}
}

func supressorState(ctx *ntcontext) (int, bool) {
	//perform some checks to see if it looks like the noise supressor is loaded
	c := ctx.paClient
	var inpLoaded, outLoaded, inputInc, outputInc bool
	var virtualDeviceInUse bool = false
	if ctx.config.FilterInput {
		module, ladspasource, err := findModule(c, "module-ladspa-source", "source_name='Filtered Microphone")
		if err != nil {
			log.Printf("Couldn't fetch module list to check for module-ladspa-source: %v\n", err)
		}
		virtualDeviceInUse = virtualDeviceInUse || (module.NUsed != 0)
		inpLoaded = ladspasource
		inputInc = false
	} else {
		inpLoaded = true
	}

	if ctx.config.FilterOutput {
		module, ladspasink, err := findModule(c, "module-ladspa-sink", "sink_name='Filtered Headphones'")
		if err != nil {
			log.Printf("Couldn't fetch module list to check for module-ladspa-sink: %v\n", err)
		}
		virtualDeviceInUse = virtualDeviceInUse || (module.NUsed != 0)
		outLoaded = ladspasink
		outputInc = false
	} else {
		outLoaded = true
	}

	if (inpLoaded || !ctx.config.FilterInput) && (outLoaded || !ctx.config.FilterOutput) && !inputInc {
		return loaded, virtualDeviceInUse
	}

	if (inpLoaded && ctx.config.FilterInput) || (outLoaded && ctx.config.FilterOutput) || inputInc || outputInc {
		return inconsistent, virtualDeviceInUse
	}

	return unloaded, virtualDeviceInUse
}

func loadSupressor(ctx *ntcontext, inp *device, out *device) error {
	if inp.checked {
		var err error
		err = loadPipeWireInput(ctx, inp)
		if err != nil {
			log.Printf("Error loading input: %v\n", err)
			return err
		}
	}

	if out.checked {
		var err error
		err = loadPipeWireOutput(ctx, out)
		if err != nil {
			log.Printf("Error loading output: %v\n", err)
			return err
		}
	}

	return nil
}

func loadModule(ctx *ntcontext, module, args string) (uint32, error) {
	idx, err := ctx.paClient.LoadModule(module, args)

	//14 = module initialisation failed
	if paErr, ok := err.(*pulseaudio.Error); ok && paErr.Code == 14 {
		resetUI(ctx)
		ctx.views.Push(makeErrorView(ctx, fmt.Sprintf("Could not load module '%s'. This is likely a problem with your system or distribution.", module)))
	}
	return idx, err
}

func loadPipeWireInput(ctx *ntcontext, inp *device) error {
	log.Printf("Loading supressor for pipewire\n")
	idx, err := loadModule(ctx, "module-ladspa-source",
		fmt.Sprintf("source_name='Filtered Microphone for %s' master=%s "+
			"rate=48000 channels=1 "+
			"label=nt-filter plugin=%s control=%d", inp.Name, inp.ID, ctx.librnnoise, ctx.config.Threshold))

	if err != nil {
		return err
	}
	log.Printf("Loaded ladspa source as idx: %d\n", idx)
	return nil
}

func loadPipeWireOutput(ctx *ntcontext, out *device) error {
	log.Printf("Loading supressor for pipewire\n")
	idx, err := loadModule(ctx, "module-ladspa-sink",
		fmt.Sprintf("sink_name='Filtered Headphones' master=%s "+
			"rate=48000 channels=1 "+
			"label=nt-filter plugin=%s control=%d", out.ID, ctx.librnnoise, ctx.config.Threshold))

	if err != nil {
		return err
	}
	log.Printf("Loaded ladspa source as idx: %d\n", idx)
	return nil
}

func unloadSupressor(ctx *ntcontext) error {
	return unloadSupressorPipeWire(ctx)
}

func unloadSupressorPipeWire(ctx *ntcontext) error {
	log.Printf("Unloading modules for pipewire\n")

	log.Printf("Searching for module-ladspa-source\n")
	c := ctx.paClient
	m, found, err := findModule(c, "module-ladspa-source", "source_name='Filtered Microphone")
	if err != nil {
		return err
	}
	if found {
		log.Printf("Found module-ladspa-source at id [%d], sending unload command\n", m.Index)
		c.UnloadModule(m.Index)
	}

	log.Printf("Searching for module-ladspa-sink\n")
	m, found, err = findModule(c, "module-ladspa-sink", "sink_name='Filtered Headphones'")
	if err != nil {
		return err
	}
	if found {
		log.Printf("Found module-ladspa-sink at id [%d], sending unload command\n", m.Index)
		c.UnloadModule(m.Index)
	}
	return nil
}

// Finds a module by exactly matching the module name, and checking if the second string is a substring of the argument
func findModule(c *pulseaudio.Client, name string, argMatch string) (module pulseaudio.Module, found bool, err error) {
	lst, err := c.ModuleList()

	if err != nil {
		return pulseaudio.Module{}, false, err
	}
	for _, m := range lst {
		if m.Name == name && strings.Contains(m.Argument, argMatch) {
			return m, true, nil
		}
	}

	return pulseaudio.Module{}, false, nil
}
