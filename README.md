# Monadic

A set of monadic functions for go  

## Showcase

- Here is classic code.

https://github.com/arcane-craft/monadic/blob/915556423b447a48127a4f5e7ed532d786b7d44e/examples/showcase/classic.go#L38-L57

- Rewrite it with monadic style, notice that we used do syntax.

https://github.com/arcane-craft/monadic/blob/10574908f265b9f8bcf586506dea5bc7a15d1f94/examples/showcase/monad.go#L15-L34

- Process code with optimize tool before building the project.

```sh
go run -mod=mod github.com/arcane-craft/monadic/tool/optimize@latest [PROJECT_ROOT_DIR]
# perform building works with build flag "monadic_production"
```

- This is the result of opitmization.

https://github.com/arcane-craft/monadic/blob/10574908f265b9f8bcf586506dea5bc7a15d1f94/examples/showcase/monad_monadic_prod.go#L15-L28
