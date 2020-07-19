# lazylit

lazylit is a code documentation tool that generates static HTML that can be
published via e.g. GitHub Pages and linked from anywhere (code comments,
READMEs, Jira tickets, etc.).

See [an example](https://dsabsay.github.io/lazylit-example/index.html) of the final output.

The main idea is that you can write extensive explanations of particularly
confusing code, tying your explanation to a specific revision (i.e. commit) of a
given source code file. You can link to this explanation from anywhere you like.
The rendered web page makes it clear which version of the source code you are
referring to so there is no expectation that the explanation be kept up-to-date.
And yet it may still help out your fellow colleagues as they navigate a complex codebase :)

I see this being particularly helpful for documenting project- and team-specific
conventions and patterns. See [here](https://dsabsay.github.io/lazylit-example/index.html) for a more complete motivation.

## Install
For macOS:
```
brew tap dsabsay/homebrew-tap
brew install lazylit
```

## Usage
Create a new repository in GitHub. Configure it to [serve static content from the
`docs/` directory](https://docs.github.com/en/github/working-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site).

Create a directory `artifacts/`.

Copy a source code file you'd like to document into a _subdirectory_ of
`artifacts/`.

```
mkdir artifacts/crazy_makefile
cp <your crazy makefile> artifacts/crazy_makefile/crazy_makefile.jul_18_2020
```

Add your documentation as comments to the file under `artifacts/`. Make sure you
add the [necessary headers](https://github.com/dsabsay/lazylit-example/blob/master/artifacts/lazylit/lazylit.jul_18_2020.go#L1).

```
lazylit
git add .
git commit -m 'new docs'
git push
```

See the [lazylit-example repo](https://github.com/dsabsay/lazylit-example) to
see what your repo should look like.
