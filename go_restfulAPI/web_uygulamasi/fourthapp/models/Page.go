//sayfalar arası transfer yaparken kullanıcı bilgileri dışında sayfa bilgilerini de göndermek istediğimizde
//Sayfanın (about ya da contact sayfası mesela) id'si, name bilgisi, description ve uri bilgisini diğer sayfaya göndermek için

package models

type Page struct {
	ID          int
	Name        string
	Description string
	URI         string
}
