{
  description = "A very dumb static site generator";

  inputs = {
    nixpkgs.url = github:nixos/nixpkgs/nixos-21.11;
    utils.url = github:numtide/flake-utils;
  };

  outputs = inputs@{ self, nixpkgs, utils, ... }:
  utils.lib.eachDefaultSystem (system:
    let pkgs = nixpkgs.legacyPackages.${system};
    in rec {
      packages = utils.lib.flattenTree {
        orator = import ./default.nix { inherit pkgs; };
      };
      defaultPackage = packages.orator;

      apps.orator = utils.lib.mkApp { drv = packages.orator; };
      defaultApp = apps.orator;

      devShell = import ./shell.nix { inherit pkgs; };
    }
  );
}
