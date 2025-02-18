/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ColocatedVariant struct {

	DbSnpId string `json:"dbSnpId,omitempty"`

	// GnomAD Non-Finnish European Allele
	GnomadNfeAllele string `json:"gnomad_nfe_allele"`

	// GnomAD Non-Finnish European MAF
	GnomadNfeMaf string `json:"gnomad_nfe_maf"`

	// GnomAD African/African American Allele
	GnomadAfrAllele string `json:"gnomad_afr_allele"`

	// GnomAD African/African American MAF
	GnomadAfrMaf string `json:"gnomad_afr_maf"`

	// GnomAD East Asian Allele
	GnomadEasAllele string `json:"gnomad_eas_allele"`

	// GnomAD East Asian MAF
	GnomadEasMaf string `json:"gnomad_eas_maf"`
}
