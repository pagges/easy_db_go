package internal

type EasyDB struct {
	memIndex *MemIndex
	df       Datafile
}

func Open() (*EasyDB, error) {
	memIndexMap := MemIndex{}
	memIndexMap.Load(Config.FilePathConfig.IndexPath)
	easyDB := EasyDB{
		memIndex: &memIndexMap,
	}
	return &easyDB, nil
}

func (db *EasyDB) Put(key, value []byte) error {
	offset, n, err := db.df.Write(NewEntry(key, value, nil))
	if err != nil {
		return err
	}
	item := Item{Key: string(key), FileID: db.df.FileID(), Offset: offset, Size: n}
	db.memIndex.Put(item)
	return nil
}
