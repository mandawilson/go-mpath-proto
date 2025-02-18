/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AlleleCount struct {

	Ac int32 `json:"ac"`

	AcAfr int32 `json:"ac_afr"`

	AcAmr int32 `json:"ac_amr"`

	AcAsj int32 `json:"ac_asj"`

	AcEas int32 `json:"ac_eas"`

	AcFin int32 `json:"ac_fin"`

	AcNfe int32 `json:"ac_nfe"`

	AcOth int32 `json:"ac_oth"`

	AcSas int32 `json:"ac_sas"`
}
