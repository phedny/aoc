package main

import (
	"aoc2024/util"
	"fmt"

	"github.com/tidwall/btree"
)

func main() {
	disk := NewDisk(util.ReadFile())
	partA(disk)
	fmt.Println(disk.Checksum())

	disk = NewDisk(util.ReadFile())
	partB(disk)
	fmt.Println(disk.Checksum())
}

func partA(disk *Disk) {
	for {
		gapPos, gapLen, _ := disk.gaps.GetAt(0)
		filePos, file, _ := disk.files.GetAt(disk.files.Len() - 1)
		if gapPos > filePos {
			return
		}
		l := gapLen
		if file.Len < gapLen {
			l = file.Len
		}
		disk.Swap(gapPos, filePos, l)
	}
}

func partB(disk *Disk) {
	filePos, file, _ := disk.files.GetAt(disk.files.Len() - 1)
	nextFile := file.Id
	for nextFile > 0 {
		disk.files.Descend(filePos, func(pos int, f File) bool {
			if f.Id == nextFile {
				filePos = pos
				file = f
				return false
			}
			return true
		})
		for gapPos, gapLen := range disk.gaps.Scan {
			if gapPos > filePos {
				break
			}
			if gapLen >= file.Len {
				disk.Swap(gapPos, filePos, file.Len)
				break
			}
		}
		nextFile--
	}
}

type File struct {
	Id  int
	Len int
}

type Disk struct {
	gaps  btree.Map[int, int]
	files btree.Map[int, File]
}

func NewDisk(spec []byte) *Disk {
	var disk Disk
	pos, id, free := 0, 0, false
	for _, v := range spec {
		length := int(v - '0')
		if free {
			disk.SetGap(pos, length)
		} else {
			disk.SetFile(id, pos, length)
			id++
		}
		free = !free
		pos += length
	}
	return &disk
}

func (disk *Disk) Checksum() int {
	var checksum int
	for pos, file := range disk.files.Scan {
		for i := range file.Len {
			checksum += (pos + i) * file.Id
		}
	}
	return checksum
}

func (disk *Disk) SetGap(p, l int) {
	if l == 0 {
		return
	}
	var pre int
	disk.gaps.Descend(p, func(p2, l2 int) bool {
		if p2+l2 == p {
			pre = l2
		}
		return false
	})
	post, _ := disk.gaps.Delete(p + l)
	p -= pre
	l += pre + post
	disk.gaps.Set(p, l)
}

func (disk *Disk) ShrinkGap(p, l int) {
	if deletedL, had := disk.gaps.Delete(p); had {
		if l < deletedL {
			disk.gaps.Set(p+l, deletedL-l)
		}
	} else {
		panic("invalid p")
	}
}

func (disk *Disk) SetFile(id, p, l int) {
	if l == 0 {
		return
	}
	var pre int
	disk.files.Descend(p, func(p2 int, f2 File) bool {
		if p2+f2.Len == p && f2.Id == id {
			pre = f2.Len
		}
		return false
	})
	p -= pre
	l += pre
	if post, has := disk.files.Get(p + l); has && post.Id == id {
		disk.files.Delete(p + l)
		l += post.Len
	}
	disk.files.Set(p, File{id, l})
}

func (disk *Disk) ShrinkFile(p, l int) {
	f, has := disk.files.Get(p)
	if !has {
		panic("ivalid p")
	}
	if f.Len == l {
		disk.files.Delete(p)
	} else {
		disk.files.Set(p, File{f.Id, f.Len - l})
	}
}

func (disk *Disk) Swap(gapPos, filePos, l int) {
	gapLen, has := disk.gaps.Get(gapPos)
	if !has {
		panic("invalid gapPos")
	}
	if gapLen < l {
		panic("invalid l")
	}
	file, has := disk.files.Get(filePos)
	if !has {
		panic("invalid filePos")
	}
	if file.Len < l {
		panic("invalid l")
	}
	// fmt.Println(gapPos, filePos, l, gapLen, file)

	disk.ShrinkGap(gapPos, l)
	disk.ShrinkFile(filePos, l)
	disk.SetFile(file.Id, gapPos, l)
	disk.SetGap(filePos+file.Len-l, l)
}
