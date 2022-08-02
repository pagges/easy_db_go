package internal

import (
	"hash/crc32"
	"os"
	"sync"
)

type EasyDB struct {
	mu       sync.RWMutex
	memIndex *MemIndex
	df       Datafile
	dfmap    map[int]Datafile
}

func Open() (*EasyDB, error) {
	memIndexMap := MemIndex{}
	easyDB := &EasyDB{
		memIndex: &memIndexMap,
	}
	err := easyDB.load()
	return easyDB, err
}

func (db *EasyDB) load() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	datafiles, lastID, err := loadDatafiles(
		Config.FilePathConfig.DataFilePath,
		Config.BasicConfig.MaxKeySize,
		Config.BasicConfig.MaxValueSize,
		0777)
	if err != nil {
		return err
	}
	currentDataFile, err := NewDatafile(
		Config.FilePathConfig.DataFilePath, lastID, false,
		Config.BasicConfig.MaxKeySize,
		Config.BasicConfig.MaxValueSize,
		0777,
	)
	if err != nil {
		return err
	}
	loadIndexes(db)
	db.dfmap = datafiles
	db.df = currentDataFile
	return nil
}

func (db *EasyDB) Put(key, value []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if err := db.fileAllocation(); err != nil {
		return ErrFileAllocation
	}
	offset, n, err := db.df.Write(NewEntry(key, value, nil))
	if err != nil {
		return err
	}
	item := Item{Key: string(key), FileID: db.df.FileID(), Offset: offset, Size: n}
	db.memIndex.Put(item)
	return nil
}

func (db *EasyDB) Get(key []byte) (Entry, error) {
	var df Datafile
	item, found := db.memIndex.Search(key)
	if !found {
		return Entry{}, nil
	}

	if item.FileID == db.df.FileID() {
		df = db.df
	} else {
		df = db.dfmap[item.FileID]
	}
	e, err := df.ReadAt(item.Offset, item.Size)
	if err != nil {
		return Entry{}, err
	}
	checksum := crc32.ChecksumIEEE(e.Value)
	if checksum != e.Checksum {
		return Entry{}, ErrChecksumFailed
	}
	return e, nil

}

func (db *EasyDB) fileAllocation() error {
	size := db.df.Size()
	maxSize := Config.BasicConfig.MaxDataFileSize
	if size < int64(maxSize) {
		return nil
	}
	err := db.df.Close()
	if err != nil {
		return err
	}
	id := db.df.FileID()
	df, err := NewDatafile(Config.FilePathConfig.DataFilePath,
		id,
		true,
		Config.BasicConfig.MaxKeySize,
		Config.BasicConfig.MaxValueSize,
		0777)
	if err != nil {
		return err
	}
	db.dfmap[id] = df
	id = db.df.FileID() + 1
	currDf, err := NewDatafile(
		Config.FilePathConfig.DataFilePath,
		id,
		false,
		Config.BasicConfig.MaxKeySize,
		Config.BasicConfig.MaxValueSize,
		0777)
	if err != nil {
		return err
	}
	db.df = currDf
	return nil
}

// DB size
func (db *EasyDB) Size() int {
	return db.memIndex.Size()
}

/**
 * load data files form storage path
 */
func loadDatafiles(path string, maxKeySize uint32, maxValueSize uint64, fileModeBeforeUmask os.FileMode) (datafiles map[int]Datafile, lastID int, err error) {
	fns, err := GetDatafiles(path)
	if err != nil {
		return nil, 0, err
	}

	ids, err := ParseIds(fns)
	if err != nil {
		return nil, 0, err
	}

	datafiles = make(map[int]Datafile, len(ids))
	for _, id := range ids {
		datafiles[id], err = NewDatafile(
			path, id, true,
			maxKeySize,
			maxValueSize,
			fileModeBeforeUmask,
		)
		if err != nil {
			return nil, 0, err
		}

	}
	if len(ids) > 0 {
		lastID = ids[len(ids)-1]
	}
	return datafiles, lastID, err
}

/**
 * load memIndex from config idnex file
 */
func loadIndexes(db *EasyDB) error {
	err := db.memIndex.Load(Config.FilePathConfig.IndexPath)
	if err != nil {
		return err
	}
	return nil
}

//close db
func (db *EasyDB) Close() error {
	err := db.memIndex.Save(Config.FilePathConfig.IndexPath)
	for _, df := range db.dfmap {
		if err := df.Close(); err != nil {
			return err
		}
	}
	return err
}
