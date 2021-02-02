/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v1.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D e2sm_rc_pre_v01_asn1/asn1c-gen`
 */

#ifndef	_E2SM_RC_PRE_IndicationHeader_Format1_H_
#define	_E2SM_RC_PRE_IndicationHeader_Format1_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct CellGlobalID;

/* E2SM-RC-PRE-IndicationHeader-Format1 */
typedef struct E2SM_RC_PRE_IndicationHeader_Format1 {
	struct CellGlobalID	*cgi	/* OPTIONAL */;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_RC_PRE_IndicationHeader_Format1_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_RC_PRE_IndicationHeader_Format1;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_RC_PRE_IndicationHeader_Format1_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_RC_PRE_IndicationHeader_Format1_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_RC_PRE_IndicationHeader_Format1_H_ */
#include "asn_internal.h"