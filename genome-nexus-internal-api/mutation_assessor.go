/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MutationAssessor struct {

	// Codon start position
	CodonStartPosition string `json:"codonStartPosition,omitempty"`

	// Number of mutations in COSMIC for this protein
	CosmicCount int32 `json:"cosmicCount,omitempty"`

	// Functional impact
	FunctionalImpact string `json:"functionalImpact,omitempty"`

	// Functional impact score
	FunctionalImpactScore float64 `json:"functionalImpactScore,omitempty"`

	Hgvs string `json:"hgvs,omitempty"`

	// Hugo gene symbol
	HugoSymbol string `json:"hugoSymbol,omitempty"`

	// User-input variants
	Input string `json:"input"`

	// Mapping issue info
	MappingIssue string `json:"mappingIssue,omitempty"`

	// Portion of gaps in variant position in multiple sequence alignment
	MsaGaps float64 `json:"msaGaps,omitempty"`

	// Number of diverse sequences in multiple sequence alignment (identical or highly similar sequences filtered out)
	MsaHeight int32 `json:"msaHeight,omitempty"`

	// Link to multiple sequence alignment
	MsaLink string `json:"msaLink,omitempty"`

	// Link to 3d structure browser
	PdbLink string `json:"pdbLink,omitempty"`

	// Reference genome variant
	ReferenceGenomeVariant string `json:"referenceGenomeVariant,omitempty"`

	// Reference genome variant type
	ReferenceGenomeVariantType string `json:"referenceGenomeVariantType,omitempty"`

	// Refseq protein ID
	RefseqId string `json:"refseqId,omitempty"`

	// Variant position in Refseq protein, can be different from the one in Uniprot
	RefseqPosition int32 `json:"refseqPosition,omitempty"`

	// Reference residue in Refseq protein, can be different from the one in Uniprot
	RefseqResidue string `json:"refseqResidue,omitempty"`

	// Number of SNPs in dbSNP for this protein
	SnpCount int32 `json:"snpCount,omitempty"`

	// Uniprot protein accession ID
	UniprotId string `json:"uniprotId,omitempty"`

	// Variant position in Uniprot protein, can be different from the one in Refseq
	UniprotPosition int32 `json:"uniprotPosition,omitempty"`

	// Reference residue in Uniprot protein, can be different from the one in Refseq
	UniprotResidue string `json:"uniprotResidue,omitempty"`

	// Amino acid substitution
	Variant string `json:"variant,omitempty"`

	// Variant conservation score
	VariantConservationScore float64 `json:"variantConservationScore,omitempty"`

	// Variant specificity score
	VariantSpecificityScore float64 `json:"variantSpecificityScore,omitempty"`
}
