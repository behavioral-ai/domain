package test

const (
	ResiliencyThreshold = "resiliency:thing/operative/agent/threshold"
	ResiliencyInterpret = "resiliency:thing/operative/agent/interpret"
)

//test: FileName(file://[cwd]/test/test-response.txt) -> [type:*url.URL] [url:C:\Users\markb\GitHub\core\iox\test\test-response.txt]
//test: FileName(file:///c:/Users/markb/GitHub/stdlib/iox/test/test-response.txt) -> [type:string] [url:c:\Users\markb\GitHub\stdlib\iox\test\test-response.txt]

// KeyToPath - convert a name/Urn to a file path
func KeyToPath(folder, name string, version int) string {

	return ""
}
