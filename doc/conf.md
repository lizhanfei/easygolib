## 配置解析
实现yaml和json 的配置文件解析

### 用法
```go
//json
target := confTest{}
confJson := NewImplJson()
_ = confJson.Load("implJsonTestConf.json", &target)

//yaml
target := confTest{}
confYaml := NewImplYaml()
_ = confYaml.Load("implYamlTestConf.yaml", &target)
```