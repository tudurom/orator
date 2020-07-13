with import <nixpkgs> { };
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
  ];
}
