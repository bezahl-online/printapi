{ lib
, buildGoModule
, fetchFromGitHub
, nixosTests
, testers
, installShellFiles
}:
let
  version = "1.0.1";
  owner = "bezahl-online";
  repo = "printapi";
  rev = "v${version}";
  sha256 = "";
in
buildGoModule {
  pname = "printapi";
  inherit version;

  src = fetchFromGitHub {
    owner = "bezahl-online";
    repo = repo;
    rev = "afa3fbaf24d386f4526d52a788f934d79822e56c";
    sha256 = "sha256-qNooHgwhFTkfAkvgX4xEjzge6MjjaCrhfKKy2NzKoUo=";
  };
  # src = ../${repo}/.;
 
  vendorSha256 = null;

  buildPhase = ''
    runHook preBuild
    CGO_ENABLED=0 go build -o printapi .
    runHook postBuild
  '';

  installPhase = ''
    mkdir -p $out/bin
    mv printapi $out/bin
    cp localhost.crt localhost.key $out/bin
  '';

  meta = with lib; {
    homepage = "https://github.com/bezahl-online/printapi";
    description = "printapi server code";
    license = licenses.mit;
    maintainers = with maintainers; [ /* list of maintainers here */ ];
  };
}

