{
  description = "A very dumb static site generator";

  inputs = {
    nixpkgs.url = github:nixos/nixpkgs/nixos-21.11;
    utils.url = github:numtide/flake-utils;
  };

  outputs = inputs@{ self, nixpkgs, utils, ... }:
  utils.lib.eachDefaultSystem (system:
    let pkgs = nixpkgs.legacyPackages.${system};
    in {
      defaultPackage = import ./default.nix { inherit pkgs; };
      devShell = import ./shell.nix { inherit pkgs; };
    }
  );
}
