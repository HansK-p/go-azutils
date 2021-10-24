package azutils

import (
	"fmt"
	"regexp"
)

var (
	rgNameRE           = regexp.MustCompile(`/resourceGroups/([^/]+)/`)
	subscriptionRE     = regexp.MustCompile(`/subscriptions/([^/]+)(?:/.*)?$`)
	reservationIdRE    = regexp.MustCompile(`/microsoft.capacity/reservationOrders/[^/]+/reservations/([^/]+)(?:/.*)?$`)
	reservationOrderRE = regexp.MustCompile(`/microsoft.capacity/reservationOrders/([^/]+)(?:/.*)?$`)
	hostGroupRE        = regexp.MustCompile(`/providers/Microsoft.Compute/hostGroups/([^/]+)(?:/.*)?$`)
	vmNameRE           = regexp.MustCompile(`providers/Microsoft.Compute/virtualMachines/([^/]+)(?:/.*)?$`)
	sqlServerNameRE    = regexp.MustCompile(`/Microsoft.Sql/servers/([^/]+)(?:/.*)?$`)
)

func IDToRGName(id string) (string, error) {
	return idExtractPart(id, *rgNameRE, "Resource Group name")
}

func IDToSubscriptionID(id string) (string, error) {
	return idExtractPart(id, *subscriptionRE, "Subscription ID")
}

func IDToReservationID(id string) (string, error) {
	return idExtractPart(id, *reservationIdRE, "Reservation ID")
}

func IDToReservationOrderID(id string) (string, error) {
	return idExtractPart(id, *reservationOrderRE, "Reservation Order ID")
}

func IDToHostGroup(id string) (string, error) {
	return idExtractPart(id, *hostGroupRE, "Hostgroup name")
}

func IDToVMName(id string) (string, error) {
	return idExtractPart(id, *vmNameRE, "VM name")
}

func IDToSQLServerName(id string) (string, error) {
	return idExtractPart(id, *sqlServerNameRE, "SQL Server name")
}

func idExtractPart(id string, re regexp.Regexp, what string) (string, error) {
	matches := re.FindStringSubmatch(id)
	if len(matches) < 1 {
		return "", fmt.Errorf("unable to extract '%s' using regexp '%s' from the id '%s'", what, re.String(), id)
	}
	return matches[1], nil
}
