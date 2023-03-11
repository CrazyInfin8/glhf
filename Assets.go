package glhf

import (
	"archive/zip"
	"bytes"
	"image"
	"io"
	"io/fs"
	"path"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/crazyinfin8/glhf/driver"
)

var assets *AssetFS = NewAssetFS()

type AssetPath struct{ Mount, Path string }

func (path AssetPath) String() string { return path.Mount + ":" + path.Path }

func NewAssetPath(assetPath string) AssetPath {
	split := strings.IndexByte(assetPath, ':')
	if split < 0 {
		return AssetPath{Mount: assetPath}
	}
	cleanPath := path.Join(".", path.Join("/", assetPath[split+1:]))
	return AssetPath{assetPath[:split], cleanPath}
}

func (assetPath AssetPath) Clean() (AssetPath, error) {
	if assetPath.Mount == "" {
		if assetPath.Path == "" {
			return AssetPath{}, ErrEmptyAssetPath
		}
		return AssetPath{}, ErrEmptyMountName
	} else if assetPath.Path == "" {
		return AssetPath{}, ErrEmptyPathName
	}

	if strings.IndexByte(assetPath.Mount, ':') != -1 {
		return AssetPath{}, ErrInvalidMountName
	}

	assetPath.Path = path.Join(".", path.Join("/", assetPath.Path))
	return assetPath, nil
}

type ColoredGraphics struct {
	Width, Height int
	Color Color
}

type AssetFS struct {
	mounted map[string]fs.FS

	loadedImages map[AssetPath]*Graphic
	loadedColors map[ColoredGraphics]*Graphic
	// loadedData   map[AssetPath][]byte
}

func NewAssetFS() *AssetFS {
	assets := new(AssetFS)

	assets.mounted = make(map[string]fs.FS)

	return assets
}

func GetAssetFS() *AssetFS {
	if assets == nil {
		assets = NewAssetFS()
	}
	return assets
}

func (assets AssetFS) checkMountName(mountName string) error {
	if strings.IndexByte(mountName, ':') != -1 {
		return ErrInvalidMountName
	}

	if _, ok := assets.mounted[mountName]; ok {
		return ErrMountPointExists
	}

	return nil
}

func (assets *AssetFS) MountZipFromData(mountName string, data []byte) error {
	err := assets.checkMountName(mountName)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(data)
	return assets.mountZip(mountName, reader, reader.Size())
}

func (assets *AssetFS) MountZipFromFile(mountName string, file fs.File) error {
	err := assets.checkMountName(mountName)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	reader, ok := file.(io.ReaderAt)
	if !ok {
		return ErrFileIsNotReaderAt
	}

	return assets.mountZip(mountName, reader, stat.Size())
}

func (assets *AssetFS) mountZip(mountName string, reader io.ReaderAt, size int64) error {
	fs, err := zip.NewReader(reader, size)
	if err != nil {
		return err
	}

	assets.mounted[mountName] = fs

	return nil
}


func (assets *AssetFS) MountFS(mountName string, fs fs.FS) error {
	if fs == nil {
		return ErrFSIsNil
	}

	err := assets.checkMountName(mountName)
	if err != nil {
		return err
	}

	assets.mounted[mountName] = fs
	return nil
}

func (assets *AssetFS) NewImageFromColor(c ColoredGraphics, cache, unique bool) (*Graphic, error) {
	if !cache {
		goto skipCacheCheck
	}

	if graphic, ok := assets.loadedColors[c]; ok {
		if unique {
			return newGraphic(nil, graphic.texture.Clone()), nil
		}
	}

skipCacheCheck:
	graphic := newGraphic(nil, driver.Drivers.NewGraphic(c.Width, c.Height, driver.GraphicOptions{false, false}))

	if cache {
		if assets.loadedColors == nil {
			assets.loadedColors = make(map[ColoredGraphics]*Graphic)
		}
		assets.loadedColors[c] = graphic
	}

	return graphic, nil
}

func (assets *AssetFS) LoadImage(assetPath AssetPath, cache, unique bool) (*Graphic, error) {
	assetPath, err := assetPath.Clean()
	if err != nil {
		return nil, err
	}

	if !cache {
		goto skipCacheCheck
	}

	if graphic, ok := assets.loadedImages[assetPath]; ok {
		if unique {
			return newGraphic(nil, graphic.texture.Clone()), nil
		}
		return graphic, nil
	}

skipCacheCheck:
	fs, ok := assets.mounted[assetPath.Mount]
	if !ok {
		return nil, ErrMountPointNotExists
	}

	file, err := fs.Open(assetPath.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	graphic := newGraphic(nil, driver.Drivers.NewGraphicFromImage(img, driver.GraphicOptions{false, false}))

	if cache {
		if assets.loadedImages == nil {
			assets.loadedImages = make(map[AssetPath]*Graphic)
		}
		assets.loadedImages[assetPath] = graphic
	}

	return graphic, nil
}

// func (assets *AssetFS) LoadFont(assetPath AssetPath, cache bool) (Graphic, error)
