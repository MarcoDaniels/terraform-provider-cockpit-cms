let
  pkgs = import (fetchTarball {
    name = "nixpkgs-23.05-darwin";
    url = "https://github.com/NixOS/nixpkgs/archive/fc541b860a28.tar.gz";
    sha256 = "0929i9d331zgv86imvsdzyfsrnr7zwhb7sdh8sw5zzsp7qsxycja";
  }) { };

  installPlugin = pkgs.writeScriptBin "installPlugin" ''
    ${pkgs.go_1_18}/bin/go install .
  '';

  terraformPlan = pkgs.writeScriptBin "terraformPlan" ''
    ${installPlugin}/bin/installPlugin
    cd examples
    rm .terraform.lock.hcl
    ${pkgs.terraform_1}/bin/terraform plan
  '';

  terraformApply = pkgs.writeScriptBin "terraformApply" ''
    ${installPlugin}/bin/installPlugin
    cd examples
    rm .terraform.lock.hcl
    TF_LOG=debug ${pkgs.terraform_1}/bin/terraform apply --auto-approve
  '';

  terraformImport = pkgs.writeScriptBin "terraformImport" ''
    ${installPlugin}/bin/installPlugin
    cd examples
    rm .terraform.lock.hcl
    ${pkgs.terraform_1}/bin/terraform init
    TF_LOG=debug ${pkgs.terraform_1}/bin/terraform import cockpit-cms_collection.collection from-terraform
    TF_LOG=debug ${pkgs.terraform_1}/bin/terraform state show cockpit-cms_collection.collection
  '';

in pkgs.mkShell {
  buildInputs = [
    pkgs.nixfmt
    pkgs.terraform_1
    pkgs.go_1_19

    installPlugin
    terraformPlan
    terraformApply
    terraformImport
  ];

  shellHook = ''
      export GOPATH="$(pwd)/go"
      export GOCACHE=""
      export GO111MODULE='on'
      export TF_CLI_CONFIG_FILE="$(pwd)/examples/.terraformrc"
  '';

  # intellij
  # set GOROOT to: go env GOROOT
  # set GOPATH $(pwd)/go
  # enable Go Modules, disable Index entire GOPATH
}
