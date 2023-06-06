package helper

func SaySomething(name string) string {
	return "Hello " + name
}

// function tidak boleh sama di satu package
// function bisa sama jika berbeda package
// jika ingin memanggil function di satu package tinggal panggil saja functionnya
// jika ingin memanggil function di beda package harus di import dulu package nya