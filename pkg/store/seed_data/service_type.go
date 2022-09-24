package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var ServiceType1 = &db.ServiceType{
	Name:        "Yerinde Tamir",
	Description: "İstanbul içinde bulunduğunuz yere geliyoruz. Mobil servis aracımızda cihazınızın tamirini gerçekleştiriyoruz. Covid-19 prosedürü kapsamında kapıdan teslim alıp, garanti belgesi ile teslim ediyoruz.",
	Price:       200,
}

var ServiceType2 = &db.ServiceType{
	Name:        "Merkezde Tamir",
	Description: "Servisimize gelmek isteyen kullanıcılarımız için, kampüsümüzde ister bahçede ister ferah kapalı bekleme salonumuzda vakit geçirme imkânı bulunmaktadır.",
	Price:       0,
}
var ServiceType3 = &db.ServiceType{
	Name:        "Kargo",
	Description: "Dilediğiniz kargo firması ile merkez servisimize arızalı telefonunuzu gönderebilirsiniz. Teslim etmeden önce, sim kartınızı ve kılıfınızı alarak, varsa kendi orijinal kutusunda, yoksa korunaklı şekilde paketleyiniz",
	Price:       0,
}

var ServiceType4 = &db.ServiceType{
	Name:        "Kurye",
	Description: "İstanbul'da kuryemizle bulunduğunuz yerden cihazınızı teslim alıyor, tamir sonrası dilediğiniz yere kurye ile teslim ediyoruz. Teslim etmeden önce, sim kartınızı ve kılıfınızı alarak, varsa kendi orijinal kutusunda, yoksa korunaklı şekilde paketleyiniz.",
	Price:       75,
}
