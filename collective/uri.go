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

// Applications can create as many domains/NISD as needed
// "agent" is the reserved domain for the agent collective supporting agent development

const (
	AgentNID = "agent" // Restricted NID/Domain

	ThingNSS    = "thing"    // urn:agent:thing.{module-package}:{type}
	AspectNSS   = "aspect"   // urn:agent:aspect.testing-aspect
	FrameNSS    = "frame"    // urn:agent:frame:testing-frame
	RuleNSS     = "rule"     // urn:agent:rule:testing-rule
	GuidanceNSS = "guidance" // urn:agent:guidance:testing-rule
)
