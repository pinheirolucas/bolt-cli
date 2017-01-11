<##>

$script:ProjectName = "bolt-cli"
$script:ProjectRoot = (Join-Path "W:/md/src/github.com/pinheirolucas/" "$script:ProjectName")
$script:ProjectGoImport = "$env:CtlrSrcName/$script:ProjectName"

function BoltTest {
	param (
		$Package="./..."
	)

	go test "$Package"
}

Export-ModuleMember -Alias @("??") `
	-Function @(
		"BoltTest"
	)
