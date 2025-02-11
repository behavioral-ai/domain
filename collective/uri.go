package collective

// https://www.rfc-editor.org/rfc/rfc8141#section-2
//   NSS           = pchar *(pchar / "/")

// https://www.rfc-editor.org/rfc/rfc3986
//  pchar         = unreserved / pct-encoded / sub-delims / ":" / "@"
//  unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
//  pct-encoded = "%" HEXDIG HEXDIG
//  sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//                  / "*" / "+" / "," / ";" / "="
//  assigned-name = "urn" ":" NID ":" NSS
// Urn syntax : "urn" : NID : NSS
// NID == Domain
