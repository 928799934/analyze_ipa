package analyze_ipa

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/928799934/go-png-cgbi"
	"howett.net/plist"
	"image/png"
	"io/ioutil"
)

func parsePListFile(f *zip.File) (*PList, error) {
	if f == nil {
		return nil, errors.New("info.plist not found")
	}
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	p := &PList{}
	if _, err = plist.Unmarshal(buf, p); err != nil {
		return nil, err
	}
	return p, nil
}

func parseIcon(f *zip.File) ([]byte, error) {
	if f == nil {
		return pngBody, nil
	}
	// 打开源文件
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	//转化还原回正常png
	img, err := cgbi.Decode(rc)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err = png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseMobileProvision(f *zip.File) ([]byte, error) {
	if f == nil {
		return nil, errors.New("embedded.mobileprovision")
	}
	// 打开源文件
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	byteData, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}
