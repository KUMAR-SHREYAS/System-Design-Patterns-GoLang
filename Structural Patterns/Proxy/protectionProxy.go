package main

import (
	"fmt"
	"strings"
)

type ProtectionProxy struct {
	fileName  string
	realImage *HighResolutionImage
	userRole  string
}

func NewProtectionProxy(fileName string, userRole string) *ProtectionProxy {
	return &ProtectionProxy{fileName: fileName, userRole: userRole}
}

func (p *ProtectionProxy) checkAccess(userRole string) bool {
	return userRole == "admin" || !strings.Contains(p.fileName, "secret")
}

func (p *ProtectionProxy) display(userRole string) {
	if !p.checkAccess(userRole) {
		fmt.Println("ProtectionProxy: Access denied for " + p.fileName)
		return
	}

	if p.realImage == nil {
		fmt.Println("ImageProxy: Loading image for authorized access...")
		p.realImage = &HighResolutionImage{fileName: p.fileName}

	}

	p.realImage.display()
}
