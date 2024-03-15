package main

import (
	"github.com/spf13/cobra"
)

/*
	NOTE: Por que uma váriavel, ao invés de uma função?

O comando raiz deve ser uma váriavel, para que seja possível adicionar novos
elementos (comandos) no momento da inicaialização `func init()`.
*/
var say = &cobra.Command{
	Use:     "golang-hello-cobra",
	Short:   "Linha de comando para comprimentar.",
	Long:    "Retornamos a melhor forma de dizer tchau.",
	Version: "0.1.0",
}

/*
	NOTE: Closure

Para executar uma função a partir de uma váriavel (Closure), é necessário
chamar o método Execute().
*/
func main() {
	say.Execute()
}

/*
	NOTE: Inicialização

Tudo que precisar ser executado ou adicionado ao comando base `say`, deve estar
nesse bloco.
*/
func init() {
	cobra.OnInitialize()
	say.AddCommand(SayHelloCommand())
	say.AddCommand(SayGoodBeyCommand())
}

// Comando - hello
func SayHelloCommand() *cobra.Command {
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "A melhor forma de dizer olá.",
		Run:   runHelloCommand,
	}
	return helloCmd
}

// Comando - goodbey
func SayGoodBeyCommand() *cobra.Command {
	goodbeyCmd := &cobra.Command{
		Use:   "goodbey",
		Short: "A melhor forma de dizer olá.",
		Run:   runGoodBeyCommand,
	}
	return goodbeyCmd
}

// Run - hello
func runHelloCommand(cmd *cobra.Command, args []string) {
	println("✌️ Olá, tudo jóia?")
}

// Run - goodbey
func runGoodBeyCommand(cmd *cobra.Command, args []string) {
	println("👋 Tachau tchau")
}
