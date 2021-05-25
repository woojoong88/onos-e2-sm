/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "../v2/e2sm-rc-pre_v2_rsys.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_NRT_RCPRE_H_
#define	_NRT_RCPRE_H_


#include "asn_application.h"

/* Including external dependencies */
#include "CellGlobalID-RCPRE.h"
#include "ARFCN-RCPRE.h"
#include "Cell-Size-RCPRE.h"
#include "PCI-RCPRE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* NRT-RCPRE */
typedef struct NRT_RCPRE {
	CellGlobalID_RCPRE_t	 cgi;
	ARFCN_RCPRE_t	 dl_ARFCN;
	Cell_Size_RCPRE_t	 cell_Size;
	PCI_RCPRE_t	 pci;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} NRT_RCPRE_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_NRT_RCPRE;
extern asn_SEQUENCE_specifics_t asn_SPC_NRT_RCPRE_specs_1;
extern asn_TYPE_member_t asn_MBR_NRT_RCPRE_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _NRT_RCPRE_H_ */
#include "asn_internal.h"