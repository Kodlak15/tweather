{
  pkgs,
  lib,
}:
pkgs.buildGoModule {
  name = "tweather";
  src = ./.;
  vendorHash = "sha256-g+yaVIx4jxpAQ/+WrGKxhVeliYx7nLQe/zsGpxV4Fn4=";

  postInstall = ''
    mv $out/bin/cmd $out/bin/tweather
  '';

  meta = with lib; {
    description = "A command line interface for retrieving data from the openweathermap.org API";
    homepage = "https://github.com/Kodlak15/tweather";
  };
}
