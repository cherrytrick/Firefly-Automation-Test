# Blockly创建自定义块

## 自定义块

本文档面向希望在Blockly中创建新块的开发人员。假设你拥有一个可以编辑的Blockly本地副本，一个人通常熟悉Blockly的用法，并且一个人对JavaScript有基本的了解。

[![img](https://developers.google.cn/blockly/images/logo.png)](https://blockly-demo.appspot.com/static/demos/blockfactory/index.html#8rxik8)

Blockly有大量的预定义块. 从数学函数到循环结构应有尽有. 但是，为了与外部应用程序接口，必须创建自定义块以形成API。例如，创建绘图程序时，可能需要创建一个“*半径为R的绘图圆*”块。

在大多数情况下，最简单的方法是找到一个已经存在的非常相似的块，将其复制并根据需要对其进行修改。以下文档适用于需要更多帮助的人。

### 定义一个块

第一步是创建一个块；指定其形状，字段和连接点。使用Blockly Developer Tools是编写此代码的最简单方法。

→ 更多信息 [Blockly Developer Tools](https://developers.google.cn/blockly/guides/create-custom-blocks/blockly-developer-tools)...

另外，您可以在研究API之后手动编写此代码。

→ 更多信息 [Defining Blocks](https://developers.google.cn/blockly/guides/create-custom-blocks/define-blocks)...

高级块可以响应于用户或其他因素动态地改变其形状。

→ 更多信息 [Mutators](https://developers.google.cn/blockly/guides/create-custom-blocks/mutators)...

### 代码生成

第二步是创建生成器代码，以将新块导出为编程语言（例如JavaScript，Python，PHP，Lua或Dart）。

→ 更多信息 [Generating Code](https://developers.google.cn/blockly/guides/create-custom-blocks/generating-code)...

要生成既干净又正确的代码，必须牢记给定语言的操作顺序列表。

→ 更多信息 [Operator Precedence](https://developers.google.cn/blockly/guides/create-custom-blocks/operator-precedence)...

创建更复杂的块需要使用临时变量和/或实用程序功能。当输入两次使用并且需要缓存时，尤其如此。

→ 更多信息 [Caching Arguments](https://developers.google.cn/blockly/guides/create-custom-blocks/caching-arguments)...

### 使用新块

创建块后，不要忘记将其添加到工具箱或在工作区中使用它。

→ 更多信息 [Toolbox](https://developers.google.cn/blockly/guides/configure/web/toolbox)...