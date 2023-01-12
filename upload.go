import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	// Max file size allowed for uploads
	MaxFileSize
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

}