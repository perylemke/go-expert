package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Hotel")
}

// Por convenção os parametros de context devem ser passados primeiramente na funcao
func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
	fmt.Println(name)
}

// Contextos são bons para informações de métricas
