package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required,min=3"` // tag min=3 artinya minimal nama 3 karakter
	Email    string `json:"email" binding:"required,email"` // tag "required,email" artinya data wajib diisi (required) dan formatnya harus berupa format email seperti @gmail.com (email)
	Role     string `json:"role"`
	Password string `json:"password" binding:"required,min=8"` // tag min=8 artinya minimal pass 8 karakter
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Perlu menggunakan tag json agar saat pengambilan payload dengan fungsi ShouldBindJSON tidak error
// Perlu menggunakan tag binding agar saat frontend mengirim data maka atribut yang ada binding wajib diisi ke backend jika tidak maka error, pemeriksaan data dilakukan pada saat ShouldBindJSON
