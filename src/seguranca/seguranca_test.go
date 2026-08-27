package seguranca

import "testing"

func TestVerificarSenha(t *testing.T) {
	hash, err := Hash("senha-secreta")
	if err != nil {
		t.Fatalf("Hash() retornou erro: %v", err)
	}

	if err := VerificarSenha(string(hash), "senha-secreta"); err != nil {
		t.Fatalf("VerificarSenha() rejeitou a senha correta: %v", err)
	}

	if err := VerificarSenha(string(hash), "senha-incorreta"); err == nil {
		t.Fatal("VerificarSenha() aceitou uma senha incorreta")
	}
}
