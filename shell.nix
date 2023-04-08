{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    go
    zsh
    go-tools
    gopls
    go-outline
    gocode
    gopkgs
    gocode-gomod
    godef
    golint
    delve
    rnix-lsp
  ];
}

