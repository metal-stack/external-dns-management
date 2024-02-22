// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"
)

func ValidateDomainName(name string) error {
	check := NormalizeHostname(name)
	if strings.HasPrefix(check, "_") {
		// allow "_" prefix, as it is used for DNS challenges of Let's encrypt
		check = "x" + check[1:]
	}

	var errs []string
	if strings.HasPrefix(check, "*.") {
		errs = validation.IsWildcardDNS1123Subdomain(check)
	} else if strings.HasPrefix(check, "@.") {
		// special case: allow apex label for Azure
		errs = validation.IsDNS1123Subdomain(check[2:])
	} else {
		errs = validation.IsDNS1123Subdomain(check)
	}

	if len(errs) > 0 {
		return fmt.Errorf("%q is no valid dns name (%v)", name, errs)
	}

	metaCheck := CalcMetaRecordDomainNameForValidation(check)
	if strings.HasPrefix(metaCheck, "*.") {
		errs = validation.IsWildcardDNS1123Subdomain(metaCheck)
	} else {
		errs = validation.IsDNS1123Subdomain(metaCheck)
	}
	if len(errs) > 0 {
		return fmt.Errorf("metadata record %q of %q is no valid dns name (%v)", metaCheck, name, errs)
	}

	labels := strings.Split(strings.TrimPrefix(check, "*."), ".")
	for i, label := range labels {
		if i == 0 && label == "@" {
			// special case: allow apex label for Azure
			continue
		}
		if errs = validation.IsDNS1123Label(label); len(errs) > 0 {
			return fmt.Errorf("%d. label %q of %q is not valid (%v)", i+1, label, name, errs)
		}
	}
	metaLabels := strings.SplitN(strings.TrimPrefix(metaCheck, "*."), ".", 2)
	if errs = validation.IsDNS1123Label(metaLabels[0]); len(errs) > 0 {
		return fmt.Errorf("1. label %q of metadata record of %q is not valid (%v)", metaLabels[0], name, errs)
	}

	return nil
}
