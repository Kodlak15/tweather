{
  description = "";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        goVersion = "22";
        overlays = [
          (final: prev: {
            go = prev."go_1_${toString goVersion}";
          })
        ];
        pkgs = import nixpkgs {inherit overlays system;};
      in {
        devShells = {
          default = pkgs.mkShell {
            packages = with pkgs; [
              zellij
              go
              gotools
            ];

            shellHook = ''
              exec $SHELL -c zellij
            '';
          };
        };
        packages = {
          default = {};
        };
      }
    );
}
