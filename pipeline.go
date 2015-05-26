// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopack

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Main interface to define an asset pipeline.
//
// Src      - Source directory to load assets from.
// Dst      - Destination directory to compile assets to.
// Matchers - Open-ended interface to define a custom pipeline.
// store    - In-memory storage by groups, patterns, file paths and contents.
//
type Pipeline struct {
	Src      string
	Dst      string
	Matchers []Matcher
	store    *assetStore
}

// Remove the destination directory (Dst). Then, compile
// all assets from the source directory (Src) into the
// destination directory (Dst).
//
// Return a map of published assets, where keys are group names
// as defined in Matchers, and values are arrays of strings,
// where each string is a path to a file.
//
func (p Pipeline) Build() map[string][]string {
	p.store = newAssetStore()
	p.clean()
	p.process()
	return p.publish()
}

// Remove the destination directory (Dst).
// Return error if any, from os.RemoveAll().
//
func (p Pipeline) clean() error {
	return os.RemoveAll(p.Dst)
}

// Walk-through the source directory (Src), calling walkFunc() for
// each file in it. Ideally, should be called after clean().
//
// Return error if any, from filepath.Walk().
//
func (p Pipeline) process() error {
	return filepath.Walk(p.Src, p.walkFunc)
}

// Apply matching processors to the given asset file,
// loading its bytes into the local, in-memory store,
// which mapping structure is allocated in runtime.
//
// Return error if any, from walkFunc().
//
func (p Pipeline) walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	for _, matcher := range p.Matchers {
		group := p.store.groups[matcher.Group]

		if group == nil {
			group = newAssetGroup()
			p.store.groups[matcher.Group] = group
		}

		for _, ePattern := range matcher.Exceptions {
			matched, err := filepath.Match(ePattern, info.Name())
			if err != nil || matched {
				return err
			}
		}

		for _, sPattern := range matcher.Patterns {
			matched, err := filepath.Match(sPattern, info.Name())
			if err != nil {
				return err
			}

			if matched {
				pattern := group.patterns[sPattern]

				if pattern == nil {
					pattern = newAssetPattern()
					group.patterns[sPattern] = pattern
				}

				filename := info.Name()
				bytes, err := ioutil.ReadFile(path)

				if err != nil {
					return err
				}

				for _, processor := range matcher.Processors {
					filename, bytes = processor.Process(filename, bytes)
				}

				basePath := strings.TrimSuffix(path, info.Name())
				relPath, _ := filepath.Rel(p.Src, basePath)
				assetPath := filepath.Join(p.Dst, relPath, filename)

				pattern.assets[assetPath] = bytes
				return nil
			}
		}
	}
	return nil
}

// Publish assets from the in-memory storage into the destination
// directory (Dst), depending on the group Concat configuration.
// Must be called after process(), so the in-memory storage is loaded.
// When Concat is set to true, assets are concatenated and published
// at the top level directory. Otherwise, they're published at the
// same directory path as the source directory (Src).
//
// Return a map of published assets, where keys are group names
// as defined in Matchers, and values are arrays of strings,
// where each string is a path to a file.
//
func (p Pipeline) publish() map[string][]string {
	result := make(map[string][]string)

	for gName, group := range p.store.groups {
		for _, sPattern := range p.findPatterns(gName) {
			pattern := group.patterns[sPattern]

			if pattern == nil {
				continue
			}

			for assetPath, bytes := range pattern.assets {
				os.MkdirAll(path.Dir(assetPath), 0775)
				ioutil.WriteFile(assetPath, bytes, 0644)
				result[gName] = append(result[gName], "/"+assetPath)
			}
		}
	}
	return result
}

// Find all patterns within a Matcher by group name.
//
// Return array of strings if found, otherwise nil.
//
func (p Pipeline) findPatterns(gName string) []string {
	for _, matcher := range p.Matchers {
		if gName == matcher.Group {
			return matcher.Patterns
		}
	}
	return nil
}
