// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 000001_init.down.db.sql (82B)
// 000001_init.up.db.sql (840B)
// 000002_add_chats.down.db.sql (74B)
// 000002_add_chats.up.db.sql (599B)
// 000003_add_contacts.down.db.sql (21B)
// 000003_add_contacts.up.db.sql (381B)
// 000004_user_messages_compatibility.down.sql (33B)
// 000004_user_messages_compatibility.up.sql (928B)
// 1567112142_user_messages.down.sql (26B)
// 1567112142_user_messages.up.sql (551B)
// doc.go (377B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_initDownDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x8a\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\x42\x92\xc9\x4d\xcd\x4d\x4a\x2d\x2a\xce\xc8\x2c\x88\x2f\x2d\x48\x49\x2c\x41\x93\x4e\xce\x48\x2c\x89\x87\xaa\xb1\xe6\x02\x04\x00\x00\xff\xff\x69\x98\x5e\xa1\x52\x00\x00\x00")

func _000001_initDownDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownDbSql,
		"000001_init.down.db.sql",
	)
}

func _000001_initDownDbSql() (*asset, error) {
	bytes, err := _000001_initDownDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.db.sql", size: 82, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe8, 0x5f, 0xe0, 0x6, 0xfc, 0xed, 0xb7, 0xff, 0xb5, 0xf3, 0x33, 0x45, 0x1, 0x5b, 0x84, 0x80, 0x74, 0x60, 0x81, 0xa6, 0x8b, 0xb4, 0xd4, 0xad, 0x10, 0xa8, 0xb3, 0x61, 0x6f, 0xc5, 0x2f, 0xaa}}
	return a, nil
}

var __000001_initUpDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xb1\x6e\xc2\x30\x14\xdc\xf3\x15\x6f\x24\x12\x43\xf7\x4e\x0e\xbc\x80\xd5\xd4\x6e\x1d\xa7\xc0\x14\x19\xe2\x82\x0b\x09\x11\x36\x52\xf9\xfb\xaa\x49\x8c\x02\x45\x15\xac\xef\xce\xe7\xbb\x7b\x6f\x24\x90\x48\x04\x49\xa2\x04\x81\xc6\xc0\xb8\x04\x9c\xd3\x54\xa6\x70\xb4\xfa\x90\x97\xda\x5a\xb5\xd6\x16\x06\x01\x00\x80\x29\x20\x4a\x78\x04\x19\xa3\xef\x19\x36\x6c\x96\x25\xc9\xb0\x01\x57\x1b\xe5\x72\x53\xc0\x07\x11\xa3\x29\x11\xd7\xe8\xbe\x72\xba\x72\xb9\x3b\xd5\xda\x53\x5a\xa4\xfb\xe3\x06\xe2\xf4\xb7\x03\x89\x73\xd9\x49\xec\xf6\xab\x2d\x44\x74\x42\x59\x37\x71\xa6\xd4\xd6\xa9\xb2\xbe\x98\xfa\xaf\xbc\xa1\x9e\x82\x37\x71\x29\x5c\x1f\x97\x3b\xb3\xca\xb7\xfa\xd4\xc4\x6b\x87\x9f\x3b\xb5\xb6\x40\x99\x3c\x07\x81\x31\xc6\x24\x4b\x24\x3c\x05\xe1\x73\x10\x74\xdd\x51\x36\xc6\xb9\x0f\x6f\x81\xb3\xcb\xe6\x06\x1d\xd2\x7b\x71\xab\xed\x52\x97\x4b\x7d\xb0\x1b\x53\xe7\xc7\xba\x50\xae\x5f\xb9\x2f\xf4\x4d\xd0\x57\x22\x16\xf0\x82\x8b\xab\x72\x0b\xe5\x54\xbb\x99\x47\x56\x12\x73\x81\x74\xc2\x1a\xbd\xb3\x4d\x10\x18\xa3\x40\x36\xc2\xb4\x79\x6e\x07\xa6\x08\x83\x10\x66\x54\x4e\x79\x26\x41\xf0\x19\x1d\xff\x9f\xa5\x91\xea\x02\x75\x29\xae\x1a\x7e\xc8\xa6\x2a\x4a\x53\x41\xc4\x79\x82\x84\xfd\x5d\x46\x4c\x92\x14\x5b\xe6\xd7\xde\x54\xba\xb8\x8b\x7a\x7f\xf6\x96\xdf\x5e\xbc\x67\x0e\x7b\x81\xc2\xdf\x63\xf8\x09\x00\x00\xff\xff\x66\xab\x2d\x2f\x48\x03\x00\x00")

func _000001_initUpDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpDbSql,
		"000001_init.up.db.sql",
	)
}

func _000001_initUpDbSql() (*asset, error) {
	bytes, err := _000001_initUpDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.db.sql", size: 840, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe7, 0x27, 0x96, 0x3b, 0x72, 0x81, 0x7d, 0xba, 0xa4, 0xfb, 0xf7, 0x4, 0xd, 0x6f, 0xc8, 0x30, 0xfe, 0x47, 0xe0, 0x9, 0xf, 0x43, 0x13, 0x6, 0x55, 0xfc, 0xee, 0x15, 0x69, 0x99, 0x53, 0x3f}}
	return a, nil
}

var __000002_add_chatsDownDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x4d\xcd\x4d\x4a\x2d\x2a\xce\xc8\x2c\x88\x2f\x2d\x48\x49\x2c\x49\x2d\xb6\xe6\x42\x92\x4e\xce\x48\x2c\x89\x87\xaa\xc1\x90\x28\xb6\xe6\x02\x04\x00\x00\xff\xff\xde\x59\xf6\x29\x4a\x00\x00\x00")

func _000002_add_chatsDownDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_chatsDownDbSql,
		"000002_add_chats.down.db.sql",
	)
}

func _000002_add_chatsDownDbSql() (*asset, error) {
	bytes, err := _000002_add_chatsDownDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_chats.down.db.sql", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1565597460, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd3, 0xa7, 0xf0, 0x94, 0x7a, 0x9, 0xdc, 0x6c, 0x7b, 0xdc, 0x12, 0x30, 0x55, 0x31, 0x17, 0xf2, 0xcc, 0x6e, 0xfd, 0xbb, 0x70, 0xb9, 0xd8, 0x9f, 0x81, 0x83, 0xdc, 0x1d, 0x1c, 0x3a, 0x8d, 0xce}}
	return a, nil
}

var __000002_add_chatsUpDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x4b\x03\x31\x10\xc5\xef\xf3\x29\x06\x3c\x54\x21\x07\x3d\x88\x42\x4f\xd9\x6d\x8a\x8b\x71\x53\xd2\x54\xec\x29\xa4\xd9\xc1\x2e\xdd\x7f\x74\xb3\x95\x7e\x7b\x69\x5d\x6a\xff\x58\xf0\x98\xbc\xf7\x7b\x93\x17\x26\xd6\x82\x1b\x81\x86\x47\x52\x60\x32\xc6\x54\x19\x14\x1f\xc9\xd4\x4c\xd1\x2f\x5d\x68\xf1\x16\xf2\x0c\xdf\xb9\x8e\x5f\xb8\xc6\x89\x4e\xde\xb8\x9e\xe3\xab\x98\xa3\x4a\x31\x56\xe9\x58\x26\xb1\x41\x2d\x26\x92\xc7\x82\x41\xe5\x4a\x3a\xb8\x77\x59\xe9\x4c\x4a\x06\xbe\x2e\xea\xf5\xc5\x3d\x8e\xc4\x98\xcf\xa4\xc1\xc1\x8d\x7b\x78\x7e\xca\x1e\x07\x0c\xc2\xb6\x21\x4c\x52\x73\x04\x3b\x1f\xf2\x0d\x61\xa4\x94\x14\x3c\xbd\xa4\x8d\x9e\x09\x06\x21\x2f\xa9\x0d\xae\x6c\xce\xe8\x8c\x0a\x0a\x94\x59\x17\xac\x2f\x6a\xbf\xb2\x1b\x57\x74\xa7\x23\x0e\x49\xf7\x0c\x9a\x6e\x51\xe4\xde\xae\x68\x8b\x91\x54\x11\x83\xae\xda\xe4\xf4\x45\x99\x2d\xa9\x6d\xdd\x27\x59\x5f\x77\x55\xb8\xca\x17\xae\xfd\xdf\xa0\xbd\xf1\x37\xb3\x0a\x54\x05\xbb\x6f\xdf\x7f\xd3\xdf\x96\x2b\xea\x49\xfb\x73\xf2\xf4\x35\x0c\x4a\x2a\x17\xb4\x6e\xfb\x82\xfd\x69\x99\x37\xb6\x6b\x32\x17\xe8\x47\x80\xbb\x21\x00\x8c\xb4\x9a\xf4\xcb\x71\xe9\x1b\x1e\xcb\xbb\x6d\xb1\xbd\x67\x08\xdf\x01\x00\x00\xff\xff\x44\xe3\x59\x6f\x57\x02\x00\x00")

func _000002_add_chatsUpDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_chatsUpDbSql,
		"000002_add_chats.up.db.sql",
	)
}

func _000002_add_chatsUpDbSql() (*asset, error) {
	bytes, err := _000002_add_chatsUpDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_chats.up.db.sql", size: 599, mode: os.FileMode(0644), modTime: time.Unix(1572543168, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0xf, 0xef, 0x31, 0x64, 0x3, 0xf0, 0x58, 0x1e, 0xed, 0x1e, 0x7e, 0xce, 0x4, 0x20, 0x51, 0x79, 0xf9, 0xca, 0xc3, 0xb9, 0xc9, 0x60, 0x25, 0x2f, 0x6b, 0x99, 0xb1, 0x92, 0x6d, 0x56, 0x31}}
	return a, nil
}

var __000003_add_contactsDownDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\xcf\x2b\x49\x4c\x2e\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x66\x64\xd9\xdd\x15\x00\x00\x00")

func _000003_add_contactsDownDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_add_contactsDownDbSql,
		"000003_add_contacts.down.db.sql",
	)
}

func _000003_add_contactsDownDbSql() (*asset, error) {
	bytes, err := _000003_add_contactsDownDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_add_contacts.down.db.sql", size: 21, mode: os.FileMode(0644), modTime: time.Unix(1565597570, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfc, 0x7e, 0xb, 0xec, 0x72, 0xcd, 0x21, 0x3e, 0xa2, 0x38, 0xe0, 0x95, 0x7e, 0xce, 0x4a, 0x17, 0xc8, 0xd0, 0x1c, 0xfa, 0xa3, 0x23, 0x5, 0xab, 0x89, 0xf9, 0xfc, 0x63, 0x7, 0x28, 0xe9, 0x93}}
	return a, nil
}

var __000003_add_contactsUpDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x41\x4b\x33\x31\x14\x45\xf7\xf3\x2b\xee\xf2\xfb\xc0\x85\x7b\x57\x99\xf1\x15\x06\x63\x52\xc6\x14\xec\x2a\xc4\xc9\xab\x06\xa7\x49\x99\xbc\x16\xfc\xf7\x52\x85\x8a\x8e\xb8\x3d\xf7\x70\x4f\x37\x90\x72\x04\xa7\x5a\x4d\x18\x4b\x96\x30\x4a\xc5\xbf\x06\x48\x11\x8e\x1e\x1d\xd6\x43\x7f\xaf\x86\x2d\xee\x68\x0b\x6b\xd0\x59\xb3\xd2\x7d\xe7\x30\xd0\x5a\xab\x8e\xae\x1a\x20\xc4\x38\x73\xad\x9f\xbe\xb1\x0e\x66\xa3\xf5\x79\xc8\x61\xcf\x4b\xca\xb9\xfa\x13\xcf\x69\x97\x38\xa2\xb5\x56\x93\x32\xb8\xa5\x95\xda\x68\x87\x95\xd2\x0f\xf4\xd3\xf2\x41\xd0\x9b\xaf\x97\x8b\x7d\xfd\x91\x9f\x52\xf8\x25\x9e\x22\x67\x49\x63\xc9\xcb\xe9\xf0\x52\xa4\x2c\xf1\x14\xaa\xf8\xe3\x21\x06\xe1\xf8\x47\xaf\xbe\x55\xe1\xbd\x97\xf0\x5c\xd1\x6a\xdb\x9e\x59\xe4\x53\x1a\xd9\xa7\xbc\x2b\x17\x26\x73\x7a\x3a\x0a\x7b\x29\x5e\xc2\xf4\xfa\xbd\xd7\xfc\xbf\x69\xde\x03\x00\x00\xff\xff\xc5\xff\x5b\xb1\x7d\x01\x00\x00")

func _000003_add_contactsUpDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_add_contactsUpDbSql,
		"000003_add_contacts.up.db.sql",
	)
}

func _000003_add_contactsUpDbSql() (*asset, error) {
	bytes, err := _000003_add_contactsUpDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_add_contacts.up.db.sql", size: 381, mode: os.FileMode(0644), modTime: time.Unix(1572543168, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x45, 0xfe, 0x1e, 0xf2, 0x75, 0x33, 0x37, 0x8e, 0x7f, 0x93, 0x6f, 0x16, 0xbb, 0xf8, 0xa4, 0x70, 0x6b, 0xe0, 0xc1, 0x4f, 0x99, 0x8d, 0xc8, 0x2d, 0x40, 0xf1, 0xed, 0x65, 0x90, 0xc3, 0xad, 0xc7}}
	return a, nil
}

var __000004_user_messages_compatibilityDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x8a\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\x8e\xcf\x49\x4d\x4f\x4c\xae\xb4\xe6\x02\x04\x00\x00\xff\xff\x25\xef\xa4\x66\x21\x00\x00\x00")

func _000004_user_messages_compatibilityDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_user_messages_compatibilityDownSql,
		"000004_user_messages_compatibility.down.sql",
	)
}

func _000004_user_messages_compatibilityDownSql() (*asset, error) {
	bytes, err := _000004_user_messages_compatibilityDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_user_messages_compatibility.down.sql", size: 33, mode: os.FileMode(0644), modTime: time.Unix(1565631683, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0xaf, 0x48, 0x80, 0x3d, 0x54, 0x5e, 0x53, 0xee, 0x98, 0x26, 0xbb, 0x99, 0x6a, 0xd8, 0x37, 0x94, 0xf2, 0xf, 0x82, 0xfa, 0xb7, 0x6a, 0x68, 0xcd, 0x8b, 0xe2, 0xc4, 0x6, 0x25, 0xdc, 0x6}}
	return a, nil
}

var __000004_user_messages_compatibilityUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xcf\x6e\xe2\x3c\x14\xc5\xf7\x3c\xc5\xd9\x01\x12\xf9\xd4\x45\xd5\x4d\x57\x81\x9a\x6f\xa2\xc9\x24\x55\x30\x23\xba\xb2\x4c\xb8\x43\xac\x49\x6c\x64\xdf\x0c\x13\xa9\x0f\x3f\xa2\x84\xaa\x61\xe8\x6a\xb2\xc8\xc2\xbf\x73\xec\x73\xff\x44\x11\x12\x1e\x07\x98\xe6\xe0\x3c\x6b\xcb\xe0\x4a\x9f\x7e\x26\x80\xf5\xb6\x26\x54\x3a\xc0\xbb\xa3\xd9\x41\x07\x1c\x09\x9e\xea\x0e\xce\xc2\xf0\x28\x8a\x70\xac\xc8\x9e\xcc\x35\x35\x64\xd9\xd8\x3d\x8c\xfd\x61\xac\x61\x8a\x42\xe9\x5d\x5d\xff\x37\x5a\x14\x22\x96\x02\x32\x9e\xa7\x02\xc9\x12\x59\x2e\x21\x36\xc9\x4a\xae\xd0\x06\xf2\xaa\xa1\x10\xf4\x9e\x82\xaa\x69\xaf\xcb\x0e\x93\x11\x00\x98\x1d\xbe\xc7\xc5\xe2\x4b\x5c\xe0\xb9\x48\xbe\xc5\xc5\x0b\xbe\x8a\x17\xe4\x19\x16\x79\xb6\x4c\x93\x85\x44\x21\x9e\xd3\x78\x21\x66\x6f\xfa\x63\x65\xc2\x81\xbc\x62\xd3\x50\x60\xdd\x1c\x90\x64\x52\xfc\x2f\x8a\xb7\xf7\xb2\x75\x9a\x9e\x75\xc1\xb5\xbe\x24\x48\xb1\x91\x57\x64\x47\x81\x8d\xd5\x6c\x9c\xc5\x3c\xcd\xe7\xe7\xd3\xd2\x59\x26\xcb\xef\x61\x86\x9e\x9e\x2a\xee\x0e\xf4\x89\xe4\x54\xa3\xd5\xcd\x3b\x3e\x9f\x0e\x62\x5e\x5f\x5a\x69\x56\x1f\xea\x1f\x52\x4f\xec\x3b\x55\xba\xd6\xf2\xc0\x8b\x27\xb1\x8c\xd7\xa9\xc4\xdd\x45\x77\xa8\x3b\xc5\x6e\xf8\x6e\xdf\xec\x41\xe0\x21\x09\xac\xb9\x0d\x43\x56\xd6\xae\xfc\xa9\x7e\xe9\xba\xa5\x1b\x79\x43\xe5\x8e\x98\xe7\x79\x2a\xe2\xec\xef\x38\xb2\x58\xf7\x23\x0a\x44\xf6\x73\xdd\x32\x4e\x57\xbd\xd0\xb5\xbc\x77\xc6\xee\xaf\xb2\x8c\xa6\x8f\xa3\xcb\x36\x25\xd9\x93\xd8\xc0\xec\x7e\xab\x7e\xa0\x79\x76\x73\x9b\x26\x67\x3c\x7d\xbc\x61\x24\xed\xcb\x4a\x6d\x3b\x75\x69\x78\x9e\xe1\xf6\x25\xe7\xf8\xed\x36\xb0\x9f\x8c\xef\xfe\xf1\x1b\xe3\xf5\xf5\x63\x47\x67\x88\x1e\xee\x67\x78\xb8\x9f\x9e\x80\xd9\xcd\x2e\x0b\x70\xaa\xf7\x4f\x00\x00\x00\xff\xff\xe0\x0c\x93\xa2\xa0\x03\x00\x00")

func _000004_user_messages_compatibilityUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_user_messages_compatibilityUpSql,
		"000004_user_messages_compatibility.up.sql",
	)
}

func _000004_user_messages_compatibilityUpSql() (*asset, error) {
	bytes, err := _000004_user_messages_compatibilityUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_user_messages_compatibility.up.sql", size: 928, mode: os.FileMode(0644), modTime: time.Unix(1569474812, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0x6c, 0x7f, 0xfb, 0xb3, 0x30, 0xd8, 0xbf, 0x6, 0x4a, 0x4a, 0xec, 0x9e, 0xd7, 0xa, 0x5e, 0x8, 0xbb, 0xf8, 0x11, 0x97, 0x80, 0xd3, 0xfa, 0x2c, 0x1e, 0x82, 0xb2, 0xe2, 0xe7, 0x6e, 0xa}}
	return a, nil
}

var __1567112142_user_messagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x8a\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xa9\xe2\x72\x97\x1a\x00\x00\x00")

func _1567112142_user_messagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1567112142_user_messagesDownSql,
		"1567112142_user_messages.down.sql",
	)
}

func _1567112142_user_messagesDownSql() (*asset, error) {
	bytes, err := _1567112142_user_messagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1567112142_user_messages.down.sql", size: 26, mode: os.FileMode(0644), modTime: time.Unix(1568961800, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x8e, 0xbe, 0x63, 0x64, 0x52, 0xa3, 0x13, 0x83, 0xc7, 0x47, 0xff, 0x56, 0xa9, 0xc, 0x72, 0xb4, 0x97, 0x6, 0xc7, 0xa5, 0x68, 0xb6, 0x55, 0x6a, 0xd5, 0xb0, 0x12, 0xfb, 0x4c, 0xa5, 0x27}}
	return a, nil
}

var __1567112142_user_messagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x51\x6f\x9b\x30\x14\x85\xdf\xf9\x15\xe7\xad\xad\xb4\x54\x7b\xef\x13\x01\x67\x43\x42\x64\x23\xce\xd6\xb7\xc8\xd8\xb7\xc5\x8a\xb1\x11\xbe\x2c\xe5\xdf\x4f\xa1\xb0\x2d\x53\x5f\xcf\x77\xfd\xdd\x73\xbd\xd9\x20\x1f\x42\x0f\xe5\x27\xf4\x03\xfd\xb2\x61\x8c\x6e\x82\x1e\x48\x31\x19\x8c\x91\x86\x53\x47\x31\xaa\x57\x8a\x60\xd5\x38\x7a\x4c\x36\x1b\xfc\x24\x98\xe0\xef\x18\x9e\xc8\x80\x03\x22\xab\x09\x8d\xd2\xe7\x8b\x1a\x0c\x74\xe8\x7a\xc5\xb6\x71\x84\x8b\xe5\x16\x96\xaf\x8f\x1a\xd2\x6a\x8c\x04\xcb\x77\x11\x3e\xf0\xd5\x6e\xae\x9b\x2f\x2d\x0d\x04\x7a\xd3\xd4\x33\x5e\xc2\x00\x6e\x09\x3a\xf8\x18\x1c\x41\x3b\x4b\x9e\x1f\x93\xbc\xde\x7f\x83\x4c\xb7\xa5\xb8\x6d\xf5\x94\x24\x59\x2d\x52\x29\x16\x58\xec\x50\xed\x25\xc4\x73\x71\x90\x87\xff\x0e\xb8\x4f\x00\xc0\x1a\x6c\xcb\xfd\x16\xc7\xaa\xf8\x7e\x14\xf3\x74\x75\x2c\xcb\x4f\x33\xd4\xad\xe2\x93\x35\xf8\x91\xd6\xd9\xd7\xb4\xfe\x43\x51\x8b\x9d\xa8\x45\x95\x89\xc3\x3c\x13\xef\xad\x79\xc0\xbe\x42\x2e\x4a\x21\x05\xb2\xf4\x90\xa5\xb9\x58\x24\xc1\x33\x79\x3e\xf1\xd4\xd3\x6a\x7a\x27\x4b\x95\x0f\x08\xd3\x1b\x43\x8a\x67\xb9\x28\x5c\xd0\x67\x6c\x8b\x2f\x45\xb5\x24\x6c\x3b\x8a\xac\xba\xfe\x26\x5d\x57\xad\xbd\xff\x31\xac\x25\x6e\xc5\xfd\xd8\x38\xab\x4f\x67\x9a\xe6\x5f\x78\x0f\x5f\x9c\x7a\x8d\x28\x2a\xf9\xf7\xde\x5c\xec\xd2\x63\x29\xf1\x39\x79\x78\x4a\x7e\x07\x00\x00\xff\xff\x07\x0e\xee\x5c\x27\x02\x00\x00")

func _1567112142_user_messagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1567112142_user_messagesUpSql,
		"1567112142_user_messages.up.sql",
	)
}

func _1567112142_user_messagesUpSql() (*asset, error) {
	bytes, err := _1567112142_user_messagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1567112142_user_messages.up.sql", size: 551, mode: os.FileMode(0644), modTime: time.Unix(1568961800, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7a, 0x11, 0x2, 0x79, 0x3b, 0xc5, 0x37, 0xe7, 0x3f, 0xa9, 0x35, 0x99, 0xa5, 0x3f, 0x32, 0xa3, 0xbe, 0xf7, 0x53, 0xf3, 0xea, 0x72, 0xbd, 0xb, 0xfc, 0xa7, 0xdc, 0x97, 0x1c, 0x6, 0x71, 0x16}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\xbb\x6e\xc3\x30\x0c\x45\x77\x7f\xc5\x45\x96\x2c\xb5\xb4\x74\xea\xd6\xb1\x7b\x7f\x80\x91\x68\x89\x88\x1e\xae\x48\xe7\xf1\xf7\x85\xd3\x02\xcd\xd6\xf5\x00\xe7\xf0\xd2\x7b\x7c\x66\x51\x2c\x52\x18\xa2\x68\x1c\x58\x95\xc6\x1d\x27\x0e\xb4\x29\xe3\x90\xc4\xf2\x76\x72\xa1\x57\xaf\x46\xb6\xe9\x2c\xd5\x57\x49\x83\x8c\xfd\xe5\xf5\x30\x79\x8f\x40\xed\x68\xc8\xd4\x62\xe1\x47\x4b\xa1\x46\xc3\xa4\x25\x5c\xc5\x32\x08\xeb\xe0\x45\x6e\x0e\xef\x86\xc2\xa4\x06\xcb\x64\x47\x85\x65\x46\x20\xe5\x3d\xb3\xf4\x81\xd4\xe7\x93\xb4\x48\x46\x6e\x47\x1f\xcb\x13\xd9\x17\x06\x2a\x85\x23\x96\xd1\xeb\xc3\x55\xaa\x8c\x28\x83\x83\xf5\x71\x7f\x01\xa9\xb2\xa1\x51\x65\xdd\xfd\x4c\x17\x46\xeb\xbf\xe7\x41\x2d\xfe\xff\x11\xae\x7d\x9c\x15\xa4\xe0\xdb\xca\xc1\x38\xba\x69\x5a\x29\x9c\x29\x31\xf4\xab\x88\xf1\x34\x79\x9f\xfa\x5b\xe2\xc6\xbb\xf5\xbc\x71\x5e\xcf\x09\x3f\x35\xe9\x4d\x31\x77\x38\xe7\xff\x80\x4b\x1d\x6e\xfa\x0e\x00\x00\xff\xff\x9d\x60\x3d\x88\x79\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 377, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xaf, 0xdf, 0xcf, 0x65, 0xae, 0x19, 0xfc, 0x9d, 0x29, 0xc1, 0x91, 0xaf, 0xb5, 0xd5, 0xb1, 0x56, 0xf3, 0xee, 0xa8, 0xba, 0x13, 0x65, 0xdb, 0xab, 0xcf, 0x4e, 0xac, 0x92, 0xe9, 0x60, 0xf1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"000001_init.down.db.sql": _000001_initDownDbSql,

	"000001_init.up.db.sql": _000001_initUpDbSql,

	"000002_add_chats.down.db.sql": _000002_add_chatsDownDbSql,

	"000002_add_chats.up.db.sql": _000002_add_chatsUpDbSql,

	"000003_add_contacts.down.db.sql": _000003_add_contactsDownDbSql,

	"000003_add_contacts.up.db.sql": _000003_add_contactsUpDbSql,

	"000004_user_messages_compatibility.down.sql": _000004_user_messages_compatibilityDownSql,

	"000004_user_messages_compatibility.up.sql": _000004_user_messages_compatibilityUpSql,

	"1567112142_user_messages.down.sql": _1567112142_user_messagesDownSql,

	"1567112142_user_messages.up.sql": _1567112142_user_messagesUpSql,

	"doc.go": docGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"000001_init.down.db.sql":                     {_000001_initDownDbSql, map[string]*bintree{}},
	"000001_init.up.db.sql":                       {_000001_initUpDbSql, map[string]*bintree{}},
	"000002_add_chats.down.db.sql":                {_000002_add_chatsDownDbSql, map[string]*bintree{}},
	"000002_add_chats.up.db.sql":                  {_000002_add_chatsUpDbSql, map[string]*bintree{}},
	"000003_add_contacts.down.db.sql":             {_000003_add_contactsDownDbSql, map[string]*bintree{}},
	"000003_add_contacts.up.db.sql":               {_000003_add_contactsUpDbSql, map[string]*bintree{}},
	"000004_user_messages_compatibility.down.sql": {_000004_user_messages_compatibilityDownSql, map[string]*bintree{}},
	"000004_user_messages_compatibility.up.sql":   {_000004_user_messages_compatibilityUpSql, map[string]*bintree{}},
	"1567112142_user_messages.down.sql":           {_1567112142_user_messagesDownSql, map[string]*bintree{}},
	"1567112142_user_messages.up.sql":             {_1567112142_user_messagesUpSql, map[string]*bintree{}},
	"doc.go":                                      {docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
