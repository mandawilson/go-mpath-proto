/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UntranslatedRegion struct {

	// UTR Type
	Type_ string `json:"type"`

	// Start position of UTR
	Start int32 `json:"start"`

	// End position of UTR
	End int32 `json:"end"`

	// Strand UTR is on, -1 for - and 1 for +
	Strand int32 `json:"strand"`
}
