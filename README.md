# Monadic

A set of monadic functions for go  

## Showcase

- Here is classic code.

https://github.com/arcane-craft/monadic/blob/1e3fc213838ccd8176ab4618779ee232e0ca0dad/examples/showcase/classic.go#L38-L57

- Rewrite it with monadic style, notice that we used do syntax.

https://github.com/arcane-craft/monadic/blob/1e3fc213838ccd8176ab4618779ee232e0ca0dad/examples/showcase/monad.go#L15-L34

- Process code with optimize tool before building the project.

```sh
go run -mod=mod github.com/arcane-craft/monadic/tool/optimize@latest [PROJECT_ROOT_DIR]
# perform building works with build flag "monadic_production"
```

- This is the result of opitmization.

https://github.com/arcane-craft/monadic/blob/1e3fc213838ccd8176ab4618779ee232e0ca0dad/examples/showcase/monad_monadic_prod.go#L15-L28
