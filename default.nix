with import <nixpkgs> { };

buildGoModule rec {
  pname = "orator";
  version = "unstable";

  src = ./.;

  modSha256 = "0f1q730wp3rhznbfb67ym4gqw8xflcvqh5k8i5hc2rg7xdr41iwg";

  meta = with stdenv.lib; {
    description = "Simple, fast and flexible static site generator written in Go.";
    homepage = "https://github.com/tudurom/orator";
    maintainers = [ maintainers.tudorr ];
    platforms = platforms.all;
  };
}
