# Monadic

A set of monadic functions for go  

# Showcase

- Here is classic code.

https://github.com/arcane-craft/monadic/blob/915556423b447a48127a4f5e7ed532d786b7d44e/examples/showcase/classic.go#L38-L57

- Rewrite it with monadic style, notice that we used do syntax.

https://github.com/arcane-craft/monadic/blob/915556423b447a48127a4f5e7ed532d786b7d44e/examples/showcase/monad.go#L13-L32

- Desugar code with optimize tool before building the project.

```sh
go run -mod=mod github.com/arcane-craft/monadic/optimize [PROJECT_ROOT_DIR]
# perform building works..
```

https://github.com/arcane-craft/monadic/blob/915556423b447a48127a4f5e7ed532d786b7d44e/examples/showcase/monad_desugar.go#L21-L38
