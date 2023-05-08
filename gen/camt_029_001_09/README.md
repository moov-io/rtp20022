## camt.029.001.09

This package is generated from a XSD file from the `xsd/` directory with the same name.

## Manual Changes

We have had to modify the generated code to work with RTP systems.

1. Every `,omitempty` type needs to be nilable. That way Go's xml package will not marshal zero-value'd XML tags.
1. The `ISODateTime` structs cannot contain the timezone. RTP defaults to US Eastern.
    - We replace the generated `ISODate` with `dt.ISODate` and `ISODateTime` with `dt.ISODateTime`.
1. Added `xml.Attr` to `ResolutionOfInvestigationV09`
