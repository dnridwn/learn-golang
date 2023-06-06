package helper

// jika di awali dengan huruf besar, maka bisa di akses dari luar package
// jika di awali dengan huruf kecil, maka tidak bisa di akses dari luar package

var ApplicationName string = "Golang" // bisa di akses dari luar package
var version string = "1.0.0" // tidak bisa di akses dari luar package

func GetVersion() string {
	return version
}

func getApplicationName() string {
	return ApplicationName
}