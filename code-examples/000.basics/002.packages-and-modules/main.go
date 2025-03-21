package main

import "module-path/module-name/other_package"

func main() {
	OtherFileSamePackageFunction()
	other_package.OtherPackageFunction()
}
