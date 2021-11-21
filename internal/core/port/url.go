package port

type UrlServices interface {
	Save(ourl string)(surl string,err error)
	Read(surl string)(ourl string,err error)
}
