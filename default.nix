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
    rev = "cd24b5f60ba0fbf902b8ddf3a6a1d4093b3722f7";
    sha256 = "sha256-e+cxat8iQe1bTLjFuoZEgdnjiqd9VFxUSXul6/VWzJM=";
  };
  # src = ../${repo}/.;
 
  vendorSha256 = "sha256-tVbMHXBv8KCM01LXTr54TlCa9Y6vyPQjwF5ywc7v5KM=";

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

