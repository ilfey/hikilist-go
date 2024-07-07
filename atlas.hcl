data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./data/entities",
    "--dialect", "sqlite", // | postgres | mysql | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "sqlite://test.db"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}