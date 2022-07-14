package feature

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

// NewFeature new Feature
func NewFeature(h *handler.Handler) (*Feature, error) {
	var err error

	f := new(Feature)
	f.handler = h

	if f.System, err = NewSystem(h); err != nil {
		return nil, err
	}

	if f.OpenAPI, err = NewOpenAPI(h); err != nil {
		return nil, err
	}

	if f.Deployment, err = NewDeployment(h); err != nil {
		return nil, err
	}

	if f.ModelPackage, err = NewModelPackage(h); err != nil {
		return nil, err
	}

	if f.Monitor, err = NewMonitor(h); err != nil {
		return nil, err
	}

	// if f.Auths, err = NewFAuths(h); err != nil {
	// 	return nil, err
	// }

	// if f.Email, err = NewFEmail(h); err != nil {
	// 	return nil, err
	// }

	return f, nil
}

// Feature represet Feature
type Feature struct {
	BaseFeature

	System       *FSystem
	OpenAPI      *FOpenAPI
	Deployment   *FDeployment
	ModelPackage *FModelPackage
	Monitor      *FMonitor
	// Auths   *FAuths
	// Email *FEmail
}
