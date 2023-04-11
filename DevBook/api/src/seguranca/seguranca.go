package seguranca

import "golang.org/x/crypto/bcrypt"

// hash recebe uma string e cloca um has nela
func Hash(senha string) ([]byte, error) {

	// gera o hash  apartir de uma senha
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha compara uma senha e um heas e retorna se elas são iguais
func VerificarSenha(SenhaComHash, SenhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(SenhaComHash), []byte(SenhaString))
}
