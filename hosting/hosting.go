package hosting

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/10gen/stitch-cli/utils"
)

// ListLocalAssetMetadata walks all files from the rootDirectory
// and builds []AssetMetadata from those files
// returns the assetMetadata, cacheData and whether or not the cacheData was altered
func ListLocalAssetMetadata(appID, rootDirectory string, assetDescriptions map[string]AssetDescription, cacheData AssetCacheDataMap) ([]AssetMetadata, AssetCacheDataMap, bool, error) {
	var assetMetadata []AssetMetadata

	hashesUpdated := false
	err := filepath.Walk(rootDirectory, buildAssetMetadata(appID, &assetMetadata, rootDirectory, assetDescriptions, &cacheData, &hashesUpdated))
	if err != nil {
		return nil, AssetCacheDataMap{}, false, err
	}

	return assetMetadata, cacheData, hashesUpdated, nil
}

func buildAssetMetadata(appID string, assetMetadata *[]AssetMetadata, rootDir string, assetDescriptions map[string]AssetDescription, cacheData *AssetCacheDataMap, hashesUpdated *bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, pathErr := filepath.Rel(rootDir, path)
			if pathErr != nil {
				return pathErr
			}
			assetPath := fmt.Sprintf("/%s", relPath)
			var assetDesc AssetDescription
			if assetDescriptions != nil {
				assetDesc = assetDescriptions[assetPath]
			}
			am, hashUpdated, fileErr := FileToAssetMetadata(appID, path, assetPath, info, assetDesc, cacheData)
			if fileErr != nil {
				return fileErr
			}

			*hashesUpdated = *hashesUpdated || hashUpdated

			*assetMetadata = append(*assetMetadata, *am)
		}
		return nil
	}
}

// FileToAssetMetadata generates a file hash for the given file
// and generates the assetAttributes and creates an AssetMetadata from these
func FileToAssetMetadata(appID, path, assetPath string, info os.FileInfo, desc AssetDescription, cacheData *AssetCacheDataMap) (*AssetMetadata, bool, error) {
	// check cache for file hash
	if cacheData.Contains(appID, assetPath) {
		acd := cacheData.Get(appID, assetPath)

		if acd.FileSize == info.Size() && acd.LastModified == info.ModTime().Unix() {
			return NewAssetMetadata(appID, assetPath, acd.FileHash, info.Size(), desc.Attrs), false, nil
		}
	}

	// file hash was not cached so generate one
	generated, err := utils.GenerateFileHashStr(path)
	if err != nil {
		return nil, false, err
	}

	cacheData.Set(appID, assetPath, AssetCacheData{
		assetPath,
		info.ModTime().Unix(),
		info.Size(),
		generated,
	})

	return NewAssetMetadata(appID, assetPath, generated, info.Size(), desc.Attrs), true, nil
}

// MetadataFileToAssetDescriptions attempts to open the file at the path given
// and build AssetDescriptions from this file
func MetadataFileToAssetDescriptions(path string) (map[string]AssetDescription, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	descs := []AssetDescription{}
	decErr := dec.Decode(&descs)
	if decErr != nil {
		return nil, decErr
	}

	descM := make(map[string]AssetDescription, len(descs))
	for _, desc := range descs {
		descM[desc.FilePath] = desc
	}

	return descM, nil
}

// CacheFileToAssetCacheData attempts to open the file at the path given
// and build a map of appID to a map of file path strings to a map of cache data
func CacheFileToAssetCacheData(path string) (AssetCacheDataMap, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	cacheData := AssetCacheDataMap{}
	decErr := dec.Decode(&cacheData)
	if decErr != nil {
		return nil, decErr
	}
	return cacheData, nil
}

// UpdateCacheFile attempts to update the file at the path given
// with the AssetCacheData passed in
func UpdateCacheFile(path string, cacheData AssetCacheDataMap) error {
	mCacheData, mErr := json.Marshal(cacheData)
	if mErr != nil {
		return mErr
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, wErr := f.Write(mCacheData)
	if wErr != nil {
		return wErr
	}

	return f.Close()
}

// DiffAssetMetadata compares a local and remote []AssetMetadata and returns a AssetMetadataDiffs
// which contains information about the difrences between the two
func DiffAssetMetadata(local, remote []AssetMetadata, merge bool) *AssetMetadataDiffs {
	var addedLocally []AssetMetadata
	var modifiedLocally []ModifiedAssetMetadata
	remoteAM := AssetsMetadata(remote).MapByPath()

	for _, lAM := range local {
		if rAM, ok := remoteAM[lAM.FilePath]; !ok {
			addedLocally = append(addedLocally, lAM)
		} else {
			modifiedAM := GetModifiedAssetMetadata(lAM, rAM)
			if modifiedAM.BodyModified || modifiedAM.AttrModified {
				modifiedLocally = append(modifiedLocally, modifiedAM)
			}
			delete(remoteAM, lAM.FilePath)
		}
	}

	var deletedLocally []AssetMetadata
	//at this point the remoteAM map only contains AssetMetadata that were deleted locally
	//if this is a merge then just ignore files deleted locally
	if !merge {
		for _, rAM := range remoteAM {
			deletedLocally = append(deletedLocally, rAM)
		}
	}

	return NewAssetMetadataDiffs(addedLocally, deletedLocally, modifiedLocally)
}

// Diff returns a list of strings representing the diff
func (amd *AssetMetadataDiffs) Diff() []string {
	var diff []string

	if len(amd.AddedLocally) > 0 {
		diff = append(diff, "New Files:")
	}
	for _, added := range amd.AddedLocally {
		diff = append(diff, fmt.Sprintf("\t+ %s", added.FilePath))
	}

	if len(amd.DeletedLocally) > 0 {
		diff = append(diff, "Removed Files:")
	}
	for _, deleted := range amd.DeletedLocally {
		diff = append(diff, fmt.Sprintf("\t- %s", deleted.FilePath))
	}

	if len(amd.ModifiedLocally) > 0 {
		diff = append(diff, "Modified Files:")
	}
	for _, modified := range amd.ModifiedLocally {
		diff = append(diff, fmt.Sprintf("\t* %s", modified.AssetMetadata.FilePath))
	}

	return diff
}
