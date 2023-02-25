package service

import (
	"regexp"
	"testing"
)

/**
..Le nom du fichier de test doit obligatoirement etre la concatenation du fichier que l'on veut tester avec _test.go.
Ce fichier de test doit être dans le même package que le fichier testé.
..Le nom des methodes de tests doivent obligatoirement commencer par Test
..Pour tester lancer go test dans le reportoire ou se trouve le fichier de test
/**


TestHelloName test la fonction greetings.Hello avec un nom
et verifie un retour de valeur valide
**/

func TestHelloName(t *testing.T) {
	name := "Jean"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Jean")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("jean") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

/*
*
TestHelloEmpty test la fonction greetings.Hello avec un nom vide
et verifie un message d'erreur
*
*/
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello(") = %q, %v, want "", error `, msg, err)
	}
}
