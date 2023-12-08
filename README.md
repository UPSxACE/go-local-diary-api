# Requirements:
- Golang
- Godoc
- Swaggo/swag
- Makefile
- Tailwind standalone CLI
- Node for @storybook/server-webpack5, and for webpack
- Git bash or simillar (to build, and for some other commands too)

#How to:
<ol>
  <li>First time running after clone
    <ol>
      <li>make dep</li>
      <li>make storybook-dep</li>
      <li>make webpack-dep</li>
    </ol>
  </li>

  <li>Run development mode
    <ol>
      <li>In one CMD do: (tailwind, with live reload)
        <ol>
          <li>make tailwind-watch</li>
        </ol>
      </li>

      <li>In a second CMD do: (storybook)
        <ol>
          <li>make storybook</li>
        </ol>
      </li>

      <li>In a third CMD do: (webpack, with live reload)
        <ol>
          <li>make webpack-watch</li>
        </ol>
      </li>

      <li>In a fourth CMD do: (server, with live reload)
        <ol>
          <li>air</li>
        </ol>
      </li>

      <li>In a fifth CMD do: (go docs)
        <ol>
          <li>make docs</li>
        </ol>
      </li>

      <li>In a sixth CMD do: (swagger ui)
        <ol>
          <li>make swag</li>
        </ol>
      </li>
    </ol>
  </li>

  <li>Alternative for running in development mode
    <ol>
      <li>In case you don't need any chat inputs, you can run all commands above at once using:
        <ol>
          <li>go run ./dev-script</li>
        </ol>
      </li>
    </ol>
  </li>

  <li>Compile for production mode
    <ol>
      <li>make build</li>
      <li>(copy the files needed to another directory)</li>
      <li>make clean</li>
    </ol>
  </li>
</ol>