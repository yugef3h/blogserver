package utils

import (
	"fmt"
	"os"
)

const ServiceNameTemplate = "com.test365.warehouse.%s.%s"

func GetMicroServiceName(name string) string {
	warehouseNS, ok := os.LookupEnv("WAREHOUSE_NS")
	if !ok || EmptyOrWhiteSpace(warehouseNS) {
		panic("WAREHOUSE_NS_NOT_FOUND")
	}
	return fmt.Sprintf(ServiceNameTemplate, warehouseNS, name)
}
