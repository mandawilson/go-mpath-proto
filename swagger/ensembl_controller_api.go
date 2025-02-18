/*
 * Genome Nexus API
 *
 * This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type EnsemblControllerApiService service


/* EnsemblControllerApiService Retrieves Ensembl canonical gene id by Entrez Gene Id
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param entrezGeneId An Entrez Gene Id. For example 23140
 @return EnsemblGene*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET(ctx context.Context, entrezGeneId string) (EnsemblGene,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EnsemblGene
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-gene/entrez/{entrezGeneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"entrezGeneId"+"}", fmt.Sprintf("%v", entrezGeneId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves canonical Ensembl Gene ID by Entrez Gene Ids
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param entrezGeneIds List of Entrez Gene Ids. For example [\&quot;23140\&quot;,\&quot;26009\&quot;,\&quot;100131879\&quot;]
 @return []EnsemblGene*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST(ctx context.Context, entrezGeneIds []string) ([]EnsemblGene,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []EnsemblGene
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-gene/entrez"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &entrezGeneIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves Ensembl canonical gene id by Hugo Symbol
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param hugoSymbol A Hugo Symbol. For example TP53
 @return EnsemblGene*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblGeneIdByHugoSymbolGET(ctx context.Context, hugoSymbol string) (EnsemblGene,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EnsemblGene
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-gene/hgnc/{hugoSymbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"hugoSymbol"+"}", fmt.Sprintf("%v", hugoSymbol), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves canonical Ensembl Gene ID by Hugo Symbols
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param hugoSymbols List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;]
 @return []EnsemblGene*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST(ctx context.Context, hugoSymbols []string) ([]EnsemblGene,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []EnsemblGene
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-gene/hgnc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &hugoSymbols
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves Ensembl canonical transcript by Hugo Symbol
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param hugoSymbol A Hugo Symbol. For example TP53
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "isoformOverrideSource" (string) Isoform override source. For example uniprot
 @return EnsemblTranscript*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblTranscriptByHugoSymbolGET(ctx context.Context, hugoSymbol string, localVarOptionals map[string]interface{}) (EnsemblTranscript,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EnsemblTranscript
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-transcript/hgnc/{hugoSymbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"hugoSymbol"+"}", fmt.Sprintf("%v", hugoSymbol), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["isoformOverrideSource"], "string", "isoformOverrideSource"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["isoformOverrideSource"].(string); localVarOk {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves Ensembl canonical transcripts by Hugo Symbols
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param hugoSymbols List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;]
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "isoformOverrideSource" (string) Isoform override source. For example uniprot
 @return []EnsemblTranscript*/
func (a *EnsemblControllerApiService) FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST(ctx context.Context, hugoSymbols []string, localVarOptionals map[string]interface{}) ([]EnsemblTranscript,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []EnsemblTranscript
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/canonical-transcript/hgnc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["isoformOverrideSource"], "string", "isoformOverrideSource"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["isoformOverrideSource"].(string); localVarOk {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &hugoSymbols
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves the transcript by an Ensembl transcript ID
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param transcriptId An Ensembl transcript ID. For example ENST00000361390
 @return EnsemblTranscript*/
func (a *EnsemblControllerApiService) FetchEnsemblTranscriptByTranscriptIdGET(ctx context.Context, transcriptId string) (EnsemblTranscript,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EnsemblTranscript
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/transcript/{transcriptId}"
	localVarPath = strings.Replace(localVarPath, "{"+"transcriptId"+"}", fmt.Sprintf("%v", transcriptId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves Ensembl Transcripts by Ensembl transcript IDs, hugo Symbols, protein IDs, or gene IDs
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param ensemblFilter List of Ensembl transcript IDs. For example [\&quot;ENST00000361390\&quot;, \&quot;ENST00000361453\&quot;, \&quot;ENST00000361624\&quot;]&lt;br&gt;OR&lt;br&gt;List of Hugo Symbols. For example [\&quot;TP53\&quot;, \&quot;PIK3CA\&quot;, \&quot;BRCA1\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl protein IDs. For example [\&quot;ENSP00000439985\&quot;, \&quot;ENSP00000478460\&quot;, \&quot;ENSP00000346196\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl gene IDs. For example [\&quot;ENSG00000136999\&quot;, \&quot;ENSG00000272398\&quot;, \&quot;ENSG00000198695\&quot;]
 @return []EnsemblTranscript*/
func (a *EnsemblControllerApiService) FetchEnsemblTranscriptsByEnsemblFilterPOST(ctx context.Context, ensemblFilter EnsemblFilter) ([]EnsemblTranscript,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []EnsemblTranscript
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/transcript"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &ensemblFilter
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Retrieves Ensembl Transcripts by protein ID, and gene ID. Retrieves all transcripts in case no query parameter provided
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "geneId" (string) An Ensembl gene ID. For example ENSG00000136999
     @param "proteinId" (string) An Ensembl protein ID. For example ENSP00000439985
     @param "hugoSymbol" (string) A Hugo Symbol For example ARF5
 @return []EnsemblTranscript*/
func (a *EnsemblControllerApiService) FetchEnsemblTranscriptsGET(ctx context.Context, localVarOptionals map[string]interface{}) ([]EnsemblTranscript,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []EnsemblTranscript
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/transcript"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["geneId"], "string", "geneId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["proteinId"], "string", "proteinId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["hugoSymbol"], "string", "hugoSymbol"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["geneId"].(string); localVarOk {
		localVarQueryParams.Add("geneId", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["proteinId"].(string); localVarOk {
		localVarQueryParams.Add("proteinId", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["hugoSymbol"].(string); localVarOk {
		localVarQueryParams.Add("hugoSymbol", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* EnsemblControllerApiService Perform lookups of Ensembl identifiers and retrieve their external references in other databases
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accession Ensembl gene accession. For example ENSG00000169083
 @return []GeneXref*/
func (a *EnsemblControllerApiService) FetchGeneXrefsGET(ctx context.Context, accession string) ([]GeneXref,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GeneXref
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ensembl/xrefs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	localVarQueryParams.Add("accession", parameterToString(accession, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

