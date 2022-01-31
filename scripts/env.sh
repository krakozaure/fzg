#!/bin/bash
# shellcheck disable=SC1090,SC1091

unset FZF_DEFAULT_COMMAND
unset FZF_DEFAULT_OPTS
FZG_CONFIG_FILE="$PWD/scripts/fzg.yaml"

if [[ -r $FZG_CONFIG_FILE ]]; then
  export FZG_CONFIG_FILE
elif [[ -r "$HOME/.config/fzg/fzg.yaml" ]]; then
  export FZG_CONFIG_FILE="$HOME/.config/fzg/fzg.yaml"
elif [[ -r "./configs/fzg.yaml" ]]; then
  export FZG_CONFIG_FILE="$PWD/configs/fzg.yaml"
elif [[ -r "./fzg.yaml" ]]; then
  export FZG_CONFIG_FILE="$PWD/fzg.yaml"
else
  printf "\e[33mUnable to find configuration file\e[0m\n" >&2
  exit 1
fi
printf "\e[32mConfiguration file : %s\e[0m\n" "$FZG_CONFIG_FILE" >&2


if [[ -r "./shell/completions.bash" ]]; then
  source "./shell/completions.bash"
  printf "\e[32mCompletions file : %s\e[0m\n" "./shell/completions.bash" >&2
elif [[ -r "$HOME/.config/fzg/completions.bash" ]]; then
  source "$HOME/.config/fzg/completions.bash"
  printf "\e[32mCompletions file : %s\e[0m\n" "$HOME/.config/fzg/completions.bash" >&2
else
  printf "\e[33mUnable to find completions file\e[0m\n" >&2
  exit 1
fi
