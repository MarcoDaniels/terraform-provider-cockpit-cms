let
  pkgs = import (fetchTarball {
    name = "nixpkgs-23.05-darwin";
    url = "https://github.com/NixOS/nixpkgs/archive/fc541b860a28.tar.gz";
    sha256 = "0929i9d331zgv86imvsdzyfsrnr7zwhb7sdh8sw5zzsp7qsxycja";
  }) { };

  name = "terraform-provider-cockpit-cms";

  plugin = "marcodaniels.com/cockpit-cms/0.1/darwin_amd64";

  build = pkgs.writeScriptBin "build" ''
    ${pkgs.go_1_18}/bin/go build -o ${name}
  '';

  install = pkgs.writeScriptBin "install" ''
    ${build}
    mkdir -p ~/.terraform.d/plugins/${plugin}
    mv ${name} ~/.terraform.d/plugins/${plugin}
  '';

in pkgs.mkShell {
  buildInputs = [
    pkgs.nixfmt
    pkgs.terraform_1
    pkgs.go_1_18

    build
    install
  ];

  shellHook = ''
      export GOPATH="$(pwd)/go"
      export GOCACHE=""
      export GO111MODULE='on'
  '';

  # intellij
  # set GOROOT to: go env GOROOT
  # set GOPATH $(pwd)/go
  # enable Go Modules, disable Index entire GOPATH
}
