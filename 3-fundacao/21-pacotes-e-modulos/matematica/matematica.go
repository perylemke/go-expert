package matematica

// Para funções, variáveis, métodos, interfaces e structs serem acessadas fora do pacote,
// devem estar com a primeira letra em maiusculo
var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando..."
}

func Soma[T int | float64](x, y T) T {
	return x + y
}
