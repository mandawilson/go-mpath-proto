/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NucleotideContext struct {

	Hgvs string `json:"hgvs,omitempty"`

	Id string `json:"id,omitempty"`

	Molecule string `json:"molecule,omitempty"`

	Query string `json:"query,omitempty"`

	// Nucleotide context sequence
	Seq string `json:"seq"`
}
