<##>

$script:ProjectName = "bolt-cli"
$script:ProjectRoot = (Join-Path "W:/md/src/github.com/pinheirolucas/" "$script:ProjectName")
$script:ProjectGoImport = "$env:CtlrSrcName/$script:ProjectName"

Export-ModuleMember -Alias @("??") `
	-Function @()
