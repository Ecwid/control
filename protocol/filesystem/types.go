package filesystem

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/storage"
)

/*
 */
type File struct {
	Name         string                `json:"name"`
	LastModified common.TimeSinceEpoch `json:"lastModified"`
	Size         float64               `json:"size"`
	Type         string                `json:"type"`
}

/*
 */
type Directory struct {
	Name              string   `json:"name"`
	NestedDirectories []string `json:"nestedDirectories"`
	NestedFiles       []*File  `json:"nestedFiles"`
}

/*
 */
type BucketFileSystemLocator struct {
	StorageKey     storage.SerializedStorageKey `json:"storageKey"`
	BucketName     string                       `json:"bucketName,omitempty"`
	PathComponents []string                     `json:"pathComponents"`
}

type GetDirectoryArgs struct {
	BucketFileSystemLocator *BucketFileSystemLocator `json:"bucketFileSystemLocator"`
}

type GetDirectoryVal struct {
	Directory *Directory `json:"directory"`
}
