package handlers

import (
	//"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AppHandler(c *gin.Context) {
	cTLen := fmt.Sprintf("%d", len(os.Getenv("CONSUL_TOKEN")))
	cACTTLen := fmt.Sprintf("%d", len(os.Getenv("CONSUL_ACL_TOKEN")))
	vTLen := fmt.Sprintf("%d", len(os.Getenv("VAULT_TOKEN")))

	b := strings.Builder{}
	b.WriteString("VAULT_HOST: " + os.Getenv("VAULT_HOST") + "\n")
	b.WriteString("VAULT_PORT: " + os.Getenv("VAULT_PORT") + "\n")
	b.WriteString("CONSUL_HOST: " + os.Getenv("CONSUL_HOST") + "\n")
	b.WriteString("CONSUL_PORT: " + os.Getenv("CONSUL_PORT") + "\n")
	b.WriteString("CLOUD_APPLICATION: " + os.Getenv("CLOUD_APPLICATION") + "\n")
	b.WriteString("CLOUD_CLUSTER: " + os.Getenv("CLOUD_CLUSTER"))
	b.WriteString("CONSUL_TOKEN: " + cTLen + "\n")
	b.WriteString("CONSUL_ACL_TOKEN: " + cACTTLen + "\n")
	b.WriteString("VAULT_TOKEN: " + vTLen)

	c.String(http.StatusOK, b.String())
}
