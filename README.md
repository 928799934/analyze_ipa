# analyze-ipa

解析IOS ipa 文件

## Example
```go
info, err := Parser("./Link_Live_1.0.ipa")
if err != nil {
	return
}
fmt.Println(info)
```