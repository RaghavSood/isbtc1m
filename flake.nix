{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv/1e4701fb1f51f8e6fe3b0318fc2b80aed0761914";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = { self, nixpkgs, devenv, systems, ... } @ inputs:
    let
      forEachSystem = nixpkgs.lib.genAttrs (import systems);
    in
    {
      devShells = forEachSystem
        (system:
          let
            pkgs = nixpkgs.legacyPackages.${system};
          in
          {
            default = devenv.lib.mkShell {
              inherit inputs pkgs;
              modules = [
                {
                  packages = with pkgs; [
                    air
                    flyctl
                    tailwindcss
                  ];

                  languages.go.enable = true;

                  enterShell = ''
                    echo "isbtc1m shell activated!"
                  '';
                }
              ];
            };
          });
       packages = forEachSystem
        (system:
          let
            pkgs = nixpkgs.legacyPackages.${system};
          in
          {
            default = pkgs.buildGoModule {
              name = "isbtc1m";

              src = ./.;
              vendorHash = "sha256-q/+ZEaHOEQ5lUHJGNN8iPBYy/IiWTvWHwuMSP+hOvaM=";

              subPackages = [ "cmd/isbtc1m" ];

              doCheck = false;
            };
          });
    };
}
