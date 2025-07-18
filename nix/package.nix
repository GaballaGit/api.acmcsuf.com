{
  lib,
  buildGoModule,
}:
buildGoModule {
  name = "api-acmcsuf";
  src = ../.;
  vendorHash = "sha256-PbObU1GpSa18kwp1kQvYTR3j+Vuh6dNLphwDYXxOejc=";

  postBuild = ''
    mv $GOPATH/bin/api $GOPATH/bin/api-acmcsuf
  '';

  subPackages = ["cmd/api"];

  meta = {
    description = "API created and used by CSUF's ACM chapter";
    homepage = "https://github.com/acmcsufoss/api.acmcsuf.com";
    license = lib.licenses.mit;
  };
}
