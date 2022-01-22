{ pkgs ? import <nixpkgs> }:
pkgs.buildGoModule rec {
  pname = "orator";
  version = "unstable";

  src = ./.;

  modSha256 = "0f1q730wp3rhznbfb67ym4gqw8xflcvqh5k8i5hc2rg7xdr41iwg";
  vendorSha256 = "05j01gyqzgg8wgdr2xhdby3chkph3lmrmlwn225zdb2bvqmiyfq0";

  meta = with pkgs.lib; {
    description = "Simple, fast and flexible static site generator written in Go.";
    homepage = "https://github.com/tudurom/orator";
    maintainers = [ maintainers.tudorr ];
    platforms = platforms.all;
  };
}
