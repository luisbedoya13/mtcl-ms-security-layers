package mtcl_ms_security_layers_poc

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("FirstFn", firstFn)
}

func firstFn(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from cloud")
}
