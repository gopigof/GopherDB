// Package dal /*
package main

import (
	"fmt"
	"os"
)

type pgnum uint64

type page struct {
	num  pgnum
	data []byte
}

// allocateEmptyPage creates a new page given a DAL object based on its page size
func (d *dal) allocateEmptyPage() *page {
	return &page{
		data: make([]byte, d.pageSize),
	}
}

type dal struct {
	file     *os.File
	pageSize int
	*freelist
}

func newDal(path string, pageSize int) (*dal, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	newDal := &dal{
		file,
		pageSize,
		newFreeList(),
	}
	return newDal, nil
}

func (d *dal) readPage(pageNum pgnum) (*page, error) {
	p := d.allocateEmptyPage()

	offset := int(pageNum) * d.pageSize
	_, err := d.file.ReadAt(p.data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, err
}

func (d *dal) writePage(p *page) error {
	offset := int(p.num) * d.pageSize
	_, err := d.file.WriteAt(p.data, int64(offset))
	return err
}

func (d *dal) close() error {
	if d.file != nil {
		err := d.file.Close()
		if err != nil {
			return fmt.Errorf("could not close file: %s", err)
		}
		d.file = nil
	}
	return nil
}
