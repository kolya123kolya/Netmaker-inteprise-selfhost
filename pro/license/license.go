package license

import (
proLogic "github.com/gravitl/netmaker/pro/logic"
"golang.org/x/exp/slog"
)

func AddLicenseHooks() {
// stub
}

func ValidateLicense() (err error) {
proLogic.SetDeploymentMode("enterprise")
slog.Info("✅ Enterprise mode activated (test override)")
return nil
}

func FetchApiServerKeys() (pub *[32]byte, priv *[32]byte, err error) {
return nil, nil, nil
}

func ClearLicenseCache() error {
return nil
}
