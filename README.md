# Requirements:
- Golang
- Godoc
- Swaggo/swag
- Makefile
- Tailwind standalone CLI
- Node for @storybook/server-webpack5, and for webpack
- Git bash or simillar (to build, and for some other commands too)

# How to:
1. First time running after clone
1.1. make dep
1.2. make storybook-dep
1.3. make webpack-dep
2. Run development mode
2.1. In one CMD do: (tailwind, with live reload)
2.1.1. make tailwind-watch
2.2. In a second CMD do: (storybook)
2.2.1. make storybook
2.3. In a third CMD do: (webpack, with live reload)
2.3.1. make webpack-watch
2.4. In a fourth CMD do: (server, with live reload)
2.4.1. air
2.5. In a fifth CMD do: (go docs)
2.5.1. make docs
2.6. In a sixth CMD do: (swagger ui)
2.6.1. make swag
3. Alternative for running in development mode
3.1. In case you don't need any chat inputs, you can run all commands above at once using:
3.1.1. go run ./dev-script
4. Compile for production mode
4.1. make build
4.2. (copy the files needed to another directory)
4.3. make clean