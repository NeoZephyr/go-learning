package meta

type FileMeta struct {
	Sha1 string
	Name string
	Size int64
	Location string
	DateCreated string
}

var fileMetas map[string]FileMeta

func init()  {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(meta FileMeta) {
	fileMetas[meta.Sha1] = meta
}

func GetFileMeta(sha1 string) FileMeta {
	return fileMetas[sha1]
}
