// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopack

// In-memory asset storage structured by
// groups, patterns, file paths and contents.
//
// Map groups by group name.
//
type assetStore struct {
	groups map[string]*assetGroup
}

// Map patterns by pattern value.
//
type assetGroup struct {
	patterns map[string]*assetPattern
}

// Map assets by filepath.
// Each asset is represented by an array of bytes.
//
type assetPattern struct {
	assets map[string][]byte
}

// Constructor: assetStore
// Initialize an empty map of groups
//
func newAssetStore() *assetStore {
	return &assetStore{groups: make(map[string]*assetGroup)}
}

// Constructor: assetGroup
// Initialize an empty map of patterns
//
func newAssetGroup() *assetGroup {
	return &assetGroup{patterns: make(map[string]*assetPattern)}
}

// Constructor: assetPattern
// Initialize an empty map of assets
//
func newAssetPattern() *assetPattern {
	return &assetPattern{assets: make(map[string][]byte)}
}
