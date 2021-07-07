package unarchive

import (
	"compress/flate"
	"compress/gzip"
	"errors"
	"github.com/mholt/archiver/v3"
	"path"
)

func Save(absFileURL, decompressDestDir string) error {
	switch path.Ext(absFileURL) {
	case ".zip":
		z := archiver.Zip{
			OverwriteExisting:    true,
			MkdirAll:             true,
			SelectiveCompression: true,
			CompressionLevel:     flate.DefaultCompression,
			FileMethod:           archiver.Deflate,
		}
		err := z.Unarchive(absFileURL, decompressDestDir)
		if err != nil {
			return err
		}
	case ".gz":
		gz := archiver.TarGz{
			CompressionLevel: gzip.DefaultCompression,
			Tar: &archiver.Tar{
				OverwriteExisting: true,
				MkdirAll:          true,
			},
		}
		err := gz.Unarchive(absFileURL, decompressDestDir)
		if err != nil {
			return err
		}
	default:
		return errors.New("not support file ext:" + absFileURL)
	}
	return nil
}
