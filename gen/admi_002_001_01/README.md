## admi.002.001.01

This package is generated from a XSD file from the `xsd/` directory with the same name.

## Manual Changes

We have had to modify the generated code to work with RTP systems.

1. Every `,omitempty` type needs to be nilable. That way Go's xml package will not marshal zero-value'd XML tags.
1. The `ISODateTime` struct cannot contain the timezone. RTP defaults to US Eastern.
   - We replace the generated `ISODateTime` with `dt.DateTime`.
1. Added a struct for the CDATA field AddtlData in order to define the xml attribute AND the `,cdata` definition for the data.
