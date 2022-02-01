# fzg = [fzf](https://github.com/junegunn/fzf) + goodies

fzg is mainly a CLI tool to use [fzf (made by junegunn)](https://github.com/junegunn/fzf) with a configuration file.

fzg will also contains a collection of some shell goodies that I use everyday in my terminal.

## Project status

This project is in a very early stage. Things might break !

## But... whyyy ?!

I started to use fzf a lot in shell functions and scripts.

After I wrote a bunch of utilities and ad-hoc command lines, I had some issues.

**Issue 1 : readability**

fzf is very customizable and comes with a lot of useful options, but my
utilities ended up with very long lines that made them less readable.

I had syntax errors because of missing one `,` as separator or the `:`.

**Solution 1**

The configuration format must be flexible enough to allow both horizontal and vertical structure.

fzg use the YAML format which can allow easier reading with vertical structure.

Instead of one long value seperated by a `,`, the value can also be written
as a YAML sequence or mapping, then fzg will use the appropriate separators for the options.

**Issue 2 : outdated state**

I reuse a lot of common commands and options between my shell utilities, but when I want
to edit one command or an option in a file, I often forget to do the update for all utilities.

Also, I sometime need to edit functions, but every functions update involve
sourcing the files containing those functions.

**Solution 2**

The configuration must be contained in one place to avoid having to edit
multiple files just for one modification.

fzg use the YAML format which also avoid configuration duplication
and allow reusability with merging, anchors and aliases features.

To avoid outdated state, the configuration can be sourced easily in functions or scripts.

This way, when options are modified, each code that use the same group of options will have the update.

**Issue 3 : human memory**

I often need to write ad-hoc command lines and either for the command arguments
or fzf options, I don't always remember them.

**Solution 3**

The commands or options must be organised and retrieved by only using a configuration key.

fzg default configuration is splitted in 3 sections : `commands`, `options` and `profiles`.

`commands` and `options` sections contain subkeys that are used to identify
commands and group of options.

The `profiles` section is used to combine command and options as one group.

**NOTICE**

I agree that the issues above are very subjective, mostly due to how I manage
things and an example of my poor skills for shell scripting.

They are not criticisms of fzf !

## How does it work

The fzg CLI tool does only two things :

1) it parses the required configuration and turn it into a string ;
2) it formats the string to be sourceable/exportable (without the `-r` flag)
   or assigned to variables as a raw value (with the `-r` flag).


Configuration example : `./configs/fzg.yaml`

```yaml
commands:
  invalid_command: null

  find_files: &cmd_find_files >-
    find . -mindepth 1
    -not \( -path './.git/*' -or -path './node_modules/*' \)
    -and -type f

options:
  invalid_options: null

  default: &opts_default
    exact: false
    extended: true
    multi: false
    reverse: true

  preview: &opts_preview
    <<: *opts_default
    multi: true
    preview-window: 60%,right,wrap
    preview:
      - 'cat -n {1} 2>/dev/null'
      - '|| tree -aCFL 1 {1} 2>/dev/null'
    prompt: 'view: '

profiles:
  invalid_profile: null

  view_files:
    command: *cmd_find_files
    options: *opts_preview
```

With the configuration above, here are some usages of fzg.
To find other usages, check `./scripts/tests`.

```sh
# print the command assigned to 'find_files'
$ fzg -c find_files
export FZF_DEFAULT_COMMAND="find . -mindepth 1 -not \\( -path './.git/*' -or -path..."

# print options assigned to 'preview'
$ fzg -o preview
export FZF_DEFAULT_OPTS="--extended --multi --no-exact --preview-window=60%,right,..."

# export variables by sourcing output of fzg, then run fzf
$ source <(fzg -c find_files -o preview); fzf

# assign raw values from configuration to each variables,
# then run fzf (notice the -r flag)
$ FZF_DEFAULT_COMMAND="$(fzg -r -c find_files)" FZF_DEFAULT_OPTS="$(fzg -r -o preview)" fzf

# print the profile assigned to 'view_files'
$ fzg -p view_files
export FZF_DEFAULT_COMMAND="find . -mindepth 1 -not \\( -path './.git/*' -or -path..."
export FZF_DEFAULT_OPTS="--extended --multi --no-exact --preview-window=60%,right,..."

# export variables by sourcing output of fzg, then run fzf
$ source <(fzg -p view_files); fzf

# use with an alias
$ alias zv='FZF_DEFAULT_COMMAND="$(fzg -r -c find_files)" FZF_DEFAULT_OPTS="$(fzg -r -o preview)" fzf'

# use with a function
$ zv(){
  local FZF_DEFAULT_COMMAND FZF_DEFAULT_OPTS
  source <(fzg -p view_files)
  fzf
}
```

```sh
$ fzg -h
USAGE: fzg [-q] [-r] [-c CMD -o OPTS | -c CMD | -o OPTS | -p PROFILE]

OPTIONS:
  -c string
        Configuration key to use for the command
  -o string
        Configuration key to use for the options
  -p string
        Configuration key to use for the profile (command+options)
  -q    Fail without printing errors but with exit code > 0 (default: false)
  -r    Print raw value without variable name or quoting (default: false)
```

## Installation

You can download binaries prebuilt with [goreleaser](https://github.com/goreleaser/goreleaser/) in the releases page.

Or you can clone this repository and build the tool yourself.

I tend to use the install script `./scripts/install`.

The script will :

- build fzg and copy it in the `$HOME/.local/bin` directory ;
- create the `$HOME/.config/fzg/` directory and copy `./configs/fzg.yaml` with
  `./shell/completions.bash` into it.


## Configuration

By default the configuration is a YAML file called `fzg.yaml`.

YAML is used rather than JSON because of merging, anchors and aliases features.

As JSON is a subset of YAML, you could also use JSON files
but you would have to define `$FZG_CONFIG_FILE`.

On execution, fzg will check these conditions to determinate what file to use.

1) `$FZG_CONFIG_FILE` environment variable points to a readable file ;
2) `$HOME/.config/fzg/fzg.yaml` is a readable file ;
3) `./fzg.yaml` is a readable file.

A configuration is available in the `./configs/` directory or in the example above.

## Goodies

At this moment, fzg only have a completion script for fzg configuration keys.

Later, some utilities that I use everyday will be added.

## Thanks

I want to say a big thank you to :

- junegunn and contributors for this amazing tool that [fzf](https://github.com/junegunn/fzf) is ;
- the GoReleaser team for their awesome tool.
